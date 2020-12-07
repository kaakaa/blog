package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

type SpotifyClient struct {
	clientId     string
	clientSecret string
	accessToken  string
}

func NewSpotifyClient(clientId, clientSecret string) (*SpotifyClient, error) {
	client := &SpotifyClient{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
	client.getAuthToken()
	return client, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *SpotifyClient) getAuthToken() error {
	if len(c.clientId) == 0 || len(c.clientSecret) == 0 {
		return fmt.Errorf("client id and (or) client secret is not specified")
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"https://accounts.spotify.com/api/token",
		strings.NewReader("grant_type=client_credentials"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	token := base64.StdEncoding.EncodeToString([]byte(c.clientId + ":" + c.clientSecret))
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to request to get access token")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to request: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response access token")
	}
	var tokenResp TokenResponse
	if err := json.Unmarshal(b, &tokenResp); err != nil {
		return errors.Wrap(err, "failed to unmarshal token response")
	}
	c.accessToken = tokenResp.AccessToken

	// Set empty string to accesstoken when expired
	go func(c *SpotifyClient, expiresIn int) {
		time.Sleep(time.Duration(expiresIn) * time.Second)
		log.Println("access token is expired")
		c.accessToken = ""
	}(c, tokenResp.ExpiresIn)

	return nil
}

func (c *SpotifyClient) request(req *http.Request) (*http.Response, error) {
	if len(c.accessToken) == 0 {
		if err := c.getAuthToken(); err != nil {
			return nil, errors.Wrap(err, "failed to get access_token")
		}
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request: %s", resp.Status)
	}
	return resp, nil
}

func (c *SpotifyClient) findRandomPlaylist(term string) (*Playlist, error) {
	// first request is for getting total number of playlists
	result, err := c.searchPlaylists(term, 0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get first playlist")
	}
	log.Printf("total: %d", result.Playlists.Total)
	// search playlists randomly
	maxOffset := result.Playlists.Total
	if maxOffset > 2000 {
		// Maximum offset is 2000: https://developer.spotify.com/documentation/web-api/reference/search/search/
		maxOffset = 2000
	}
	result, err = c.searchPlaylists(term, rand.Intn(maxOffset)-1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get random playlist")
	}
	return &result.Playlists.Playlist[0], nil
}

func (c *SpotifyClient) searchPlaylists(term string, offset int) (*SearchResult, error) {
	log.Printf("SearchTerm: %s, Offset: %d", term, offset)

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("type", "playlist")
	q.Add("limit", "1")
	q.Add("q", url.PathEscape(term))
	q.Add("offset", strconv.Itoa(offset))
	req.URL.RawQuery = q.Encode()

	resp, err := c.request(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request search playlist")
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response")
	}
	var result SearchResult
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall search results")
	}
	if len(result.Playlists.Playlist) == 0 {
		return nil, fmt.Errorf("No playlists found")
	}

	return &result, nil
}

type SearchResult struct {
	Playlists struct {
		Playlist []Playlist `json:"items"`
		Total    int        `json:"total"`
	} `json:"playlists"`
}

type Playlist struct {
	Description string `json:"description"`
	ExternalUrl struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Id     string `json:"id"`
	Images []struct {
		Height int    `json:"height`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	}
	Name  string `json:"name"`
	Owner struct {
		DisplayName  string `json: "display_name"`
		ExternalUrls struct {
			Spotify string `json: "spotify"`
		} `json: "external_urls"`
	} `json: "owner"`
	Tracks struct {
		Total int `json: "total"`
	} `json: "tracks"`
}

func (p *Playlist) ToMessage() []*model.SlackAttachment {
	// Find url with most large height
	var maxHeight int
	var imageUrl string
	for _, image := range p.Images {
		if maxHeight < image.Height {
			maxHeight = image.Height
			imageUrl = image.Url
		}
	}
	return []*model.SlackAttachment{{
		Title:      p.Name,
		TitleLink:  p.ExternalUrl.Spotify,
		Text:       fmt.Sprintf("**Title**: %s\n**Description**: %s", p.Name, p.Description),
		AuthorName: p.Owner.DisplayName,
		AuthorLink: p.Owner.ExternalUrls.Spotify,
		ImageURL:   imageUrl,
		Footer:     fmt.Sprintf("%d tracks", p.Tracks.Total),
	}}
}

type PlaylistTracks struct {
	Items []struct {
		Track struct {
			Artists []struct {
				Name string `json: "name"`
			} `json: "artists"`
			Name string `json: "name"`
		} `json: "track"`
	} `json: "items"`
}

func (c *SpotifyClient) getTracks(id string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PlaylistTracks
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	var tracks []string
	for _, item := range result.Items {
		var artists []string
		for _, artist := range item.Track.Artists {
			artists = append(artists, artist.Name)
		}
		tracks = append(tracks, fmt.Sprintf("♫ %s / %s", item.Track.Name, strings.Join(artists, ", ")))
	}
	return tracks, nil
}
