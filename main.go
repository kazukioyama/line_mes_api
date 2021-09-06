package main

import (
	"fmt"
	"go_line_bot/line"
	"go_line_bot/models"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/auth", authHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from Docker container!")
	fmt.Println("Hello World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	// ToDo: ファイルの変更内容がdockerコンテナ上にリアルタイムで反映されない問題の解決
	admin := models.Find(1)
	// admin := entity.Admin{Id: 1, LineMessagingId: 1656383139, LineMessagingSecret: "3b636209f942b66dd8910510d8eb0406", LineMessagingToken: "yS8yY1Y4nvSOV36RQcHyca1V4wRzKjgitcGTX9/C0b4p0tkKspTxY363HnlawAq9EFk0+0RtilUCgIrPTALQx+blLFHoKK9LyN5u+03GoNOdzlBxazjzTBrdeC+pIZVGtIC4emg5hrIQVdHEMGgCnAdB04t89/1O/w1cDnyilFU=", LineLoginId: 1656382872, LineLoginSecret: "ab954a407e9bcd398dd1132cd04c8fb1"}
	oauth := line.New(&admin)
	http.Redirect(w, r, oauth.AuthUri("ddd"), http.StatusMovedPermanently) // http.StatusMovedPermanently→301
}
