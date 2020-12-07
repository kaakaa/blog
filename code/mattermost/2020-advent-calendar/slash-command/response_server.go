package main

import (
	"io"
	"log"
	"net/http"
)

const (
	WebHookToken = "8w7foap4ufrsfczda8uez51yxo"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%#v", r.Form)
		r.ParseForm()

		log.Printf("token: %v", r.Form.Get("token"))
		// (1) Tokenをチェック
		if r.Form.Get("token") != WebHookToken {
			log.Printf("received an invalid request with token: %s", r.Form.Get("token"))
			return
		}

		// (2) レスポンスとしてMattermostへ投稿を作成
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{
			"text": "**TEST**",
			"response_type": "in_channel",
			"extra_responses": [
				{
					"text": "**Secret Info**"
				},
				{
					"text": "Notify in other channel",
					"channel_id": "u6ijofzg8bnsie5u8c1eumtc1h"
				}
			]
		}`)
	})
	http.ListenAndServe(":8080", nil)
}
