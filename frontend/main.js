new Vue({
  el: "#app",
  vuetify: new Vuetify(),
  data: {
    message: "this is the best of food.",
    textmessage: "",
    chip1: true,
    // データテーブル用のヘッダーと値
    headers: [{ text: "値", value: "val" }],
    numbers: [{ val: 124 }, { val: 3553 }],
    // シンプルテーブルのデータ
    dummys: [{ name: 1 }, { name: 2 }],
  },
  methods: {
    onChipClick: async function () {
      this.message = "YOSHINOYA is better.";
    },
    onClose: function () {
      this.message = "good night.";
    },
    onBtnClick: async function () {
      let mes = await callbackMessage("test"); // Call Go functin
      this.message = mes;

      let dummyTable = await callbackTable(); // Call Go function
      // データテーブルに行を追加
      for (let i = 0; i < dummyTable.length; ++i) {
        this.numbers.push({ val: dummyTable[i] });
      }
    },
    onClick: async function () {
      let mes = await callbackMessage(this.textmessage); // Call Go function
      this.message = mes;

      let dummylist = await callbackList(); // Call Go function
      //シンプルテーブルに行を追加　配列の場合ループ処理必要
      for (let i = 0; i < dummylist.length; ++i) {
        this.dummys.push({ name: dummylist[i] });
      }
    },
  },
});
