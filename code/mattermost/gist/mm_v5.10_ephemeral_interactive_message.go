// v5.10.0で追加された、自分のみ閲覧可能なInteractive Messageを投稿する投稿するサンプル
// Go言語SDKのCreatePostEphemeralメソッドからは作成でkなかったため、現在はプラグインの`SendEphemeralPost`からしか作成できない？

package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func main() {
	plugin.ClientMain(&Plugin{})
}

func (p *Plugin) OnActivate() error {
	return p.API.RegisterCommand(&model.Command{
		Trigger:      "mm510-eim",
		AutoComplete: true,
	})
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	post := &model.Post{
		Type:      model.POST_EPHEMERAL,
		ChannelId: args.ChannelId,
		Props: model.StringInterface{
			"attachments": []*model.SlackAttachment{
				{
					Text: "test",
					Actions: []*model.PostAction{
						{
							Type:       model.POST_ACTION_TYPE_SELECT,
							Name:       "TEST",
							DataSource: "channels",
							Integration: &model.PostActionIntegration{
								URL: fmt.Sprintf("%s/eim", ""),
							},
						},
					},
				},
			},
		},
	}
	post.Message = "SendEphemeralPost"
	p.API.SendEphemeralPost(args.UserId, post)
	return &model.CommandResponse{
		Text: "test",
	}, nil
}
