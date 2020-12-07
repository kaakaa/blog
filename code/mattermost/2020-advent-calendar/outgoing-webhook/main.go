package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

func main() {
	http.HandleFunc("/outgoing", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.OutgoingWebhookPayload
		json.Unmarshal(b, &payload)
		log.Printf("Request: %#v", payload)
	})
	http.ListenAndServe(":8080", nil)
}
