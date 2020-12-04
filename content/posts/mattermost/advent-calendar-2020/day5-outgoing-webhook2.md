---
title: "[Mattermost Integrations] Outgoing WebHook 発展編"
date: 2020-12-05T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "webhook", "outgoing"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第5日目の記事です。

前回の記事では、Mattermostに投稿が作成された際に、その投稿情報を外部アプリケーションに送信することができるOutgoing WebHook(外向きのウェブフック)の機能について紹介しましたが、本記事はその続きです。Outgoing WebHookが送信された外部アプリケーションからMattermostへ投稿を作成する方法について紹介します。

## Outgoing WebHookのレスポンスとして投稿を作成する

昨日の記事では、Outgoing WebHookを受け取ってコンソールへ出力する例を紹介しましたが、受け取ったリクエストへのリクエストとしてJSONデータを返却することで、外部アプリケーションからMattermostへ投稿を作成することもできます。

以下のコードは昨日のコードにレスポンスを返すコードを追加したものです。

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
	WebHookToken = "9mgpebi9a7bq3qjedd1kt6mwtr"
)

func main() {
	http.HandleFunc("/outgoing", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.OutgoingWebhookPayload
		json.Unmarshal(b, &payload)

		// (1) Tokenのチェック
		if payload.Token != WebHookToken {
			log.Printf("received an invalid request with token: %s", payload.Token)
			return
		}

		// (2) レスポンスとしてMattermostへ投稿を作成
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf(`{"text": "Echo: %s"}`, payload.Text))
	})
	http.ListenAndServe(":8080", nil)
}
```

`(1)` では、トークンを利用したリクエストの検証を行っています。
`(2)` で、Outgoing WebHookのレスポンスとして投稿を作成しています。投稿を作成するには、ヘッダーに`Content-Type: application/json`を指定し、レスポンスボディに`text`フィールドを含むJSONデータを指定します。ここでは、受け取ったメッセージをそのまま返しています。

このサーバに対してOutgoing WebHookを送信した時の画面が下記になります。

![response](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day5/response.png)

Outgoing WebHookが送信されると、ユーザー名の横に`BOT`というラベルのついた投稿が作成されます。

レスポンスボディのJSONに含めることができるパラメータは、3日目の記事で紹介したIncoming WebHook作成時のパラメータと多くが共通します。異なる点は下記です。

* `icon_emoji`は使用できない
* `channel`による投稿先の設定はできない
* `response_type`が指定できる

このうち、`response_type`に`comment`を指定した場合、Outgoing WebHookのきっかけとなった投稿への返信という形で投稿が作成されます。
先ほどのコードのレスポンス部分を下記のように変更し、

```go
		io.WriteString(w, fmt.Sprintf(`{"text": "Echo: %s", "response_type": "comment"}`, payload.Text))
```

Outgoing WebHookによるリクエストを送信すると、下記の通り元の投稿への返信として投稿が作成されます。

![as comment](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day5/response-as-comment.png)

## さいごに

本日は、Outgoing WebHookの詳細な使い方について紹介しました。
明日からは、Slash Commandの使い方について紹介していきます。
