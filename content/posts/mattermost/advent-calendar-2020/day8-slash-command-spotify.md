---
title: "[Mattermost Integrations] Slash Command 応用編"
date: 2020-12-08T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "slashcommand", "command"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第8日目の記事です。

本日の記事は、昨日まで紹介してきたSlash Commandを使ったサンプルアプリの紹介です。

## 背景

コロナ影響によりWFH(Work From Home)が定着してきましたが、毎日自宅で作業していると同じ景色しか見えないため気持ちが滅入ってくることがあります。自分は少しでも日常に変化を付けようと音楽を流しながら作業していることが多いので、MattermostからSpotifyのプレイリストを探すことができる統合機能を作ってみました。（普段はAppleMusic派ですが...）

## 動作例

![movie](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day8/example-music-command.gif)

`/music [search_term]`で、`[search_term]`を検索キーとしてプレイリストを検索します。`[search_term]`を指定しない場合、`workfromhome`を検索キーとして検索します。
コマンドを実行すると、見つかったプレイリストの情報をリンク付きでMattermostに投稿します。この時、`card`機能を使って投稿からプレイリストの内容を確認できるようにしています。

## コード

ソースコードは下記のリポジトリにコミットしてあります。

https://github.com/kaakaa/blog/tree/master/code/mattermost/2020-advent-calendar/slash-command-spotify

Mattermost連携部分の動作コードは下記になります。これとは別にSpotifyクライアント用のコードもありますが、そちらの紹介は割愛します。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	// (1) 定数の設定
	MattermostToken     = "arzoyh4i57dw5x57yumenuktqw"
	SpotifyClientID     = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	SpotifyClientSecret = "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ"
)

func main() {
	// (2) Spotify クライアントの初期化
	client, err := NewSpotifyClient(SpotifyClientID, SpotifyClientSecret)
	if err != nil {
		log.Fatalf("failed to set up spotify client: %s", err.Error())
	}
	http.HandleFunc("/spotify", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// (3) トークンによるリクエストの検証
		if MattermostToken != r.Form.Get("token") {
			log.Printf("received token is invalid: %s", r.Form.Get("token"))
			return
		}

		// (4) serach_termの取得
		term := "workfromhome"
		if t := r.Form.Get("text"); len(t) > 0 {
			term = t
		}

		// (5) プレイリストの検索
		playlist, err := client.findRandomPlaylist(term)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("failed to serach playlist: %s", err.Error()))
			return
		}

		// (6) プレイリスト内のトラック情報の取得
		tracks, err := client.getTracks(playlist.Id)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("failed to get tracks: %s", err.Error()))
			return
		}
		contextInfo := []string{"Tracks"}
		for _, t := range tracks {
			contextInfo = append(contextInfo, fmt.Sprintf("1. %s", t))
		}

		// (7) Mattermostへのレスポンスの構築
		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Username:     "Music Provider",
			Props: map[string]interface{}{
				"card": strings.Join(contextInfo, "\n"),
			},
			Attachments: playlist.ToMessage(),
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})

	// Start server
	http.ListenAndServe(":8080", nil)
}
```

### (1) 定数の設定
プログラム内で使用する定数を宣言しています。

Mattermost関係はSlash Commandから送信されるリクエスト検証用のトークン文字列。  
Spotifyは、アプリケーションのクライアントID/クライアントシークレットです。これらの値の取得方法については [Spotify Developer Platform: Spotify APIアクセスしてデータ取得してみてみた \- Qiita](https://qiita.com/shirok/items/ba5c45511498b75aac27) などを参照ください。[Client Credential Flow](https://developer.spotify.com/documentation/general/guides/authorization-guide/)で認証するため、Redirect URIなどの設定は必要無いかと思います。

### (2) Spotify クライアントの初期化
SpotifyのAPIを実行するためのインスタンスを構築しています。内容については割愛します。

### (3) トークンによるリクエストの検証
MattermostのCustom Slash Commandによるリクエストを受信したら、まずはトークンによる検証を行います。

### (4) serach_termの取得
Custom Slash Commandのリクエストから、Slash Commandの引数として指定されたテキストを取得します。これがSpotifyへの検索キーワードになります。
Custom Slash Commandのリクエスト内容については公式ドキュメントを参照ください。
* https://developers.mattermost.com/integrate/slash-commands/#basic-usage

### (5) プレイリストの検索
(2) で初期化したSpotifyクライアントを使ってプレイリストの検索を行います。詳細は割愛しますが、[Search for an Item \| Spotify for Developers](https://developer.spotify.com/documentation/web-api/reference/search/search/)のAPIを実行し、ランダムに選んだ一つのプレイリストを返すようになっています。

### (6) プレイリスト内のトラック情報の取得
こちらもSpotifyクライアントを使って(5)で取得したプレイリストのトラックを取得しています。内部では[Get a Playlist's Items \| Spotify for Developers](https://developer.spotify.com/documentation/web-api/reference/playlists/get-playlists-tracks/)のAPIを実行しています。
一度のAPIリクエストで取得できる最大のトラックス数が100であり、100トラック以上あるプレイリストについては、最初の100トラックだけを取得しています。

### (7) Mattermostへのレスポンスの構築
今まで取得してきた情報を使ってMattermostへのレスポンスを構築しています。`attachments`フィールドは[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)の機能を使っていますが、[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)については後日紹介します。

---

## さいごに

以上のように、Custom Slash Commandによるリクエストを受信したサーバーから、外部のAPIを実行し、結果をMattermostに返すようなアプリケーションを構築することができます。外部APIの実行に時間がかかる恐れがある場合は、[Delayed Response](https://developers.mattermost.com/integrate/slash-commands/#delayed-and-multiple-responses)を利用するとSlash Commandを実行したユーザーへフィードバックをすぐに返せるため便利です。[Delayed Response]の機能については昨日の記事で紹介しました。

明日は、Message Attachmentsについて紹介します。
