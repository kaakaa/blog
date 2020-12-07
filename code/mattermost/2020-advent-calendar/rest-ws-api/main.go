package main

import (
	"log"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	MattermostWebSocketURL = "ws://localhost:8065"
	AccessToken            = "3uhox91m6bdbt8pqsczouy9mny"
)

func main() {
	client, appErr := model.NewWebSocketClient4(MattermostWebSocketURL, AccessToken)
	if appErr != nil {
		log.Fatal(appErr.Error())
	}

	if appErr = client.Connect(); appErr != nil {
		log.Fatal(appErr.Error())
	}

	client.Listen()

	for {
		select {
		case event := <-client.EventChannel:
			log.Printf("Received: %v", event)
		}
	}
}
