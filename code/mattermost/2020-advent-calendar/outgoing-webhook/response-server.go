package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	WebHookToken = "9mgpebi9a7bq3qjedd1kt6mwtr"
)

func main() {
	http.HandleFunc("/outgoing", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.OutgoingWebhookPayload
		json.Unmarshal(b, &payload)

		// (1) Tokenのチェック
		if payload.Token != WebHookToken {
			log.Printf("received an invalid request with token: %s", payload.Token)
			return
		}

		// (2) レスポンスとしてMattermostへ投稿を作成
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf(`{"text": "Echo: %s", "response_type": "comment"}`, payload.Text))
	})
	http.ListenAndServe(":8080", nil)
}
