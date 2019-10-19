package main

import (
	"io"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

const (
	ServerURL     = "http://localhost:8080"
	MattermostURL = "http://localhost:8065"
	AccessToken   = "yj4xzckhiig5urrdzm4ydduccc"
)

var client *model.Client4

func main() {
	client = model.NewAPIv4Client(MattermostURL)
	client.SetOAuthToken(AccessToken)
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
				Text:         "Sample Response1",
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
