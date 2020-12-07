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
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Attachments: []*model.SlackAttachment{{
				Title: "Echo server",
				Actions: []*model.PostAction{{
					// (1) Echoメニュー
					Name: "Echo",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/echo",
						Context: map[string]interface{}{
							"text": r.Form.Get("text"),
						},
					},
					Type: model.POST_ACTION_TYPE_SELECT,
					Options: []*model.PostActionOptions{{
						Text:  "Echo",
						Value: "echo",
					}, {
						Text:  "Reject",
						Value: "reject",
					}},
				}},
			}},
		}

		// Need to set header. if not, just json string will be posted.
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})

	// (2) Echo Buttonが押されたときの処理
	http.HandleFunc("/actions/echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (3) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.PostActionIntegrationRequest
		json.Unmarshal(b, &payload)

		text, ok := payload.Context["text"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['text'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		selected, ok := payload.Context["selected_option"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['selected_option'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		// (4) レスポンスの構築
		response := &model.PostActionIntegrationResponse{}
		switch selected {
		case "echo":
			response.Update = &model.Post{
				Message: text,
				Props:   model.StringInterface{},
			}
		case "reject":
			response.Update = &model.Post{Props: model.StringInterface{}}
			response.EphemeralText = "Echoing was rejected"
		default:
			response.EphemeralText = "invalid operation"
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	http.ListenAndServe(":8080", nil)
}
