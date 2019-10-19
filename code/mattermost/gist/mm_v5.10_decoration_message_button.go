// v5.10.0からMessage Attachmentsのタイトルに絵文字とリンクを使用できるようになったため、その機能を試すサンプルコード

package main

import (
	"fmt"
	"log"

	"github.com/mattermost/mattermost-server/model"
)

const (
	MattermostURL = "http://localhost:8065"
	ChannelId     = "zyxwvutsrqponmlkjihgfedcba"
	AccessToken   = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	client := model.NewAPIv4Client(MattermostURL)
	client.AuthToken = AccessToken
	client.AuthType = model.HEADER_BEARER
	post := &model.Post{
		Type:      model.POST_SLACK_ATTACHMENT,
		ChannelId: ChannelId,
		Props: model.StringInterface{
			"attachments": []*model.SlackAttachment{
				{
					Text: "Emoji :+1:\n[Link](https://www.google.com)",
					Actions: []*model.PostAction{
						{
							Type:       model.POST_ACTION_TYPE_BUTTON,
							Name:       "Button :+1:",
							DataSource: "channels",
							Integration: &model.PostActionIntegration{
								URL: fmt.Sprintf("%s/deco", ""),
							},
						},
					},
				},
			},
		},
	}
	if _, resp := client.CreatePost(post); resp.Error != nil {
		log.Fatalf("Failed to create post: %#v", resp.Error)
	}
}
