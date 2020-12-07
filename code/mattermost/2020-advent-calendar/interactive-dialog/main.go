package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	MattermostAccessToken = "71t1n3pykbnstkr4c66mgwxaja"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// (1) Slash Command実行時に送信されるリクエストから "Trigger ID" を取得
		triggerId := r.Form.Get("trigger_id")

		// (2) Interactive Dialogを起動するためのリクエストを構築
		request := model.OpenDialogRequest{
			TriggerId: triggerId,
			URL:       "http://localhost:8080/actions/dialog",
			Dialog: model.Dialog{
				Title: "Sample Interactive Dialog",
				Elements: []model.DialogElement{{
					DisplayName: "Title",
					Name:        "title",
					Type:        "text",
				}, {
					DisplayName: "Message",
					Name:        "message",
					Type:        "textarea",
				}},
			},
		}

		// (3) Interactive Dialogを開く
		b, _ := json.Marshal(request)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8065/api/v4/actions/dialogs/open", bytes.NewReader(b))
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)

		commandResp := model.CommandResponse{}
		if err != nil {
			commandResp.ResponseType = model.COMMAND_RESPONSE_TYPE_EPHEMERAL
			commandResp.Text = err.Error()
		}
		if resp.StatusCode != http.StatusOK {
			commandResp.ResponseType = model.COMMAND_RESPONSE_TYPE_EPHEMERAL
			commandResp.Text = fmt.Sprintf("failed to request: %s", resp.Status)
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, commandResp.ToJson())
	})

	http.HandleFunc("/actions/dialog", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (4) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.SubmitDialogRequest
		json.Unmarshal(b, &payload)

		title, ok := payload.Submission["title"].(string)
		if !ok {
			resp := model.SubmitDialogResponse{Error: "failed to get title"}
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, string(resp.ToJson()))
			return
		}
		msg, ok := payload.Submission["message"].(string)
		if !ok {
			resp := model.SubmitDialogResponse{Error: "failed to get message"}
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, string(resp.ToJson()))
			return
		}

		// (5) Message Attachmentsインスタンス作成
		post := &model.Post{
			ChannelId: payload.ChannelId,
			Props: model.StringInterface{
				"attachments": []*model.SlackAttachment{{
					Title: title,
					Text:  msg,
				}},
			},
		}

		// (6) REST APIによるメッセージ投稿
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8065/api/v4/posts", strings.NewReader(post.ToJson()))
		req.Header.Add("Authorization", "Bearer "+MattermostAccessToken)
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)

		// (7) エラー処理
		dialogResp := model.SubmitDialogResponse{}
		if err != nil {
			dialogResp.Error = err.Error()
		}
		if resp.StatusCode != http.StatusCreated {
			dialogResp.Error = fmt.Sprintf("failed to request: %s", resp.Status)
		}
		dialogResp.Errors = map[string]string{
			"title":   "title error",
			"message": "message error",
		}
		dialogResp.Error = "error"
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(dialogResp.ToJson()))
	})

	http.ListenAndServe(":8080", nil)
}
