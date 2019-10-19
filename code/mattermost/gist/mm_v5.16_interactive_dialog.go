package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

const (
	ServerURL     = "http://localhost:8080"
	MattermostURL = "http://localhost:8065"
	AccessToken   = "oqpjxgucnjypz817th363rcpjr"
)

var client *model.Client4

func main() {
	client = model.NewAPIv4Client(MattermostURL)
	client.SetOAuthToken(AccessToken)
	http.HandleFunc("/dialog", handleDialogRequest)
	http.HandleFunc("/callback", handleCallBack)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleDialogRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	log.Printf("%#v", r)

	dialog := model.OpenDialogRequest{
		TriggerId: r.Form["trigger_id"][0],
		URL:       fmt.Sprintf("%s/callback", ServerURL),
		Dialog: model.Dialog{
			Title:            "SampleDialog",
			IntroductionText: "ここには **Markdown** 形式のテストが書けます。 [Google](https://www.google.com)",
			SubmitLabel:      "Submit",
			Elements: []model.DialogElement{
				{
					DisplayName: "パスワード",
					Name:        "password",
					Type:        "text",
					SubType:     "password",
				},
				{
					DisplayName: "チェックボックス",
					Name:        "check",
					Placeholder: "チェック1",
					Type:        "bool",
				},
				{
					DisplayName: "ラジオボタン",
					Name:        "radio",
					Type:        "radio",
					Options: []*model.PostActionOptions{{
						Text:  "オプション1",
						Value: "option1",
					}, {
						Text:  "オプション2",
						Value: "option2",
					}, {
						Text:  "オプション3",
						Value: "option3",
					}},
				},
			},
		},
	}

	if ok, resp := client.OpenInteractiveDialog(dialog); !ok {
		log.Fatalf("%#v", resp)
	}
}

func handleCallBack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	request := model.SubmitDialogRequestFromJson(r.Body)
	text := request.Submission["text"].(string)
	channel := request.Submission["channel"].(string)

	client.CreatePost(&model.Post{
		ChannelId: channel,
		Message:   fmt.Sprintf("```\n%s\n```", text),
	})
}
