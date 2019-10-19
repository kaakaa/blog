// v5.8.0より有効になった、スラッシュコマンド実行時のレスポンスメッセージを別のチャンネルに投稿するサンプル
// L31で別のチャンネルのIDを指定しています。

package main

import (
        "io"
        "log"
        "net/http"

        "github.com/mattermost/mattermost-server/model"
)

const (
        ServerURL      = "http://localhost:8080"
        MattermostURL  = "http://localhost:8066"
        OtherChannelID = "csg7u95tipn63ciybec6whb8ao"
)

func main() {
        http.HandleFunc("/multi", handleMultiResponse)
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatal(err)
        }
}

func handleMultiResponse(w http.ResponseWriter, r *http.Request) {
        resp := model.CommandResponse{
                ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
                Text:         "Multi Response Sample",
                ExtraResponses: []*model.CommandResponse{
                        {
                                ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
                                ChannelId:    OtherChannelID,
                                Text:         "Sample Response1 to other channel",
                        },
                        {
                                ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
                                Text:         "Sample Response2(Ephemeral)",
                        },
                },
        }
        w.Header().Add("Content-Type", "application/json")
        io.WriteString(w, resp.ToJson())
}