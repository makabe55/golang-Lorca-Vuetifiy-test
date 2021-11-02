package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

//go:embed frontend
var fs embed.FS

type collback struct {
	message   string
	numbers   []string
	dataTable []int
}

func (c *collback) GetTable() []int {
	c.dataTable = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return c.dataTable
}

func (c *collback) Getlist() []string {
	c.numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	return c.numbers
}

func (c *collback) ReturnMessage(mes string) string {
	c.message = mes
	return c.message
}

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 1080, 640, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	// Create and bind Go object to the UI
	c := &collback{}
	ui.Bind("callbackMessage", c.ReturnMessage)
	ui.Bind("callbackList", c.Getlist)
	ui.Bind("callbackTable", c.GetTable)

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/frontend", ln.Addr()))

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
	console.log("Hello, world!");
`)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
