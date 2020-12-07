package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

func main() {
	// (1) Interactive Message Button作成用スラッシュコマンド
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Attachments: []*model.SlackAttachment{{
				Title: "Echo server",
				Actions: []*model.PostAction{{
					// (2) Echoボタン
					Name: "Echo",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/echo",
						Context: map[string]interface{}{
							"text": r.Form.Get("text"),
						},
					},
				}, {
					// (3) Rejectボタン
					Name:  "Reject",
					Style: "danger",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/reject",
					},
				}},
			}},
		}

		// Need to set header. if not, just json string will be posted.
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})

	// (4) Echo Buttonが押されたときの処理
	http.HandleFunc("/actions/echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (5) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.PostActionIntegrationRequest
		json.Unmarshal(b, &payload)
		text, ok := payload.Context["text"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['text'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		// (6) レスポンスの構築
		response := &model.PostActionIntegrationResponse{
			Update: &model.Post{
				Message: text,
				Props:   model.StringInterface{},
			},
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	// (7) Rejectボタンが押されたときの処理
	http.HandleFunc("/actions/reject", func(w http.ResponseWriter, r *http.Request) {
		// (8) レスポンスの構築
		response := &model.PostActionIntegrationResponse{
			Update: &model.Post{
				Props: model.StringInterface{},
			},
			EphemeralText: "Echoing was rejected.",
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	http.ListenAndServe(":8080", nil)
}
