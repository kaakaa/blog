package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	WebHookToken = "8w7foap4ufrsfczda8uez51yxo"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// (1) Tokenをチェック
		if r.Form.Get("token") != WebHookToken {
			log.Printf("received an invalid request with token: %s", r.Form.Get("token"))
			return
		}

		// (2) レスポンスとしてMattermostへ投稿を作成
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"text": "Request is recieved. Results will return later."}`)

		// (3) 5秒後に "response_url" に対してリクエストを送信
		go func(url string) {
			time.Sleep(5 * time.Second)
			body := strings.NewReader(`{"text": "This is the result.", "response_type": "in_channel"}`)
			http.Post(url, "application/json", body)
		}(r.Form.Get("response_url"))
	})
	http.ListenAndServe(":8080", nil)
}
