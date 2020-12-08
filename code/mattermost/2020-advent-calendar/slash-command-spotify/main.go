package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	MattermostToken     = "arzoyh4i57dw5x57yumenuktqw"
	SpotifyClientID     = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	SpotifyClientSecret = "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
)

func main() {
	client, err := NewSpotifyClient(SpotifyClientID, SpotifyClientSecret)
	if err != nil {
		log.Fatalf("failed to set up spotify client: %s", err.Error())
	}
	http.HandleFunc("/spotify", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if MattermostToken != r.Form.Get("token") {
			log.Printf("received token is invalid: %s", r.Form.Get("token"))
			return
		}

		// Read query
		term := "workfromhome"
		if t := r.Form.Get("text"); len(t) > 0 {
			term = t
		}

		// Search playlists
		playlist, err := client.findRandomPlaylist(term)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("failed to serach playlist: %s", err.Error()))
			return
		}

		// Get tracks in playlist
		tracks, err := client.getTracks(playlist.Id)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("failed to get tracks: %s", err.Error()))
			return
		}
		contextInfo := []string{"Tracks"}
		for _, t := range tracks {
			contextInfo = append(contextInfo, fmt.Sprintf("1. %s", t))
		}

		// Response
		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Username:     "Music Provider",
			Props: map[string]interface{}{
				"card": strings.Join(contextInfo, "\n"),
			},
			Attachments: playlist.ToMessage(),
		}

		// Need to set header. if not, just json string will be posted.
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})
	http.ListenAndServe(":8080", nil)
}
