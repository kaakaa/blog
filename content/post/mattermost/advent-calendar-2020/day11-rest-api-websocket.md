---
title: "[Mattermost Integrations] REST API (WebSocket)"
date: 2020-12-11T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "api", "rest", "websocket"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第11日目の記事です。

本記事では、MattermostのWebSocket APIついて紹介します。

## WebSocket APIの概要
MattermostはWebSocket APIを利用することで、Mattermost上でやり取りされているWebSocketのイベントを取得することもできます。

WebSocket APIについてはREST APIのドキュメントに説明があります。
https://api.mattermost.com/#tag/WebSocket

## WebSocketイベントの取得

MattermostのWebSocket Eventを取得するために、Mattermost公式のWebSocketクライアント [websocket_client.go](https://github.com/mattermost/mattermost-server/blob/master/model/websocket_client.go) を利用します。


```go
package main

import (
	"log"

	"github.com/mattermost/mattermost-server/v5/model"
)

const (
  // (1) Mattermost WebSocket Endpoint
  MattermostWebSocketURL = "ws://localhost:8065"
  // (2) Access Token
	AccessToken            = "4cacdozwn3fndnzobbpha3nnhy"
)

func main() {
  // (3) WebSocket Clientの構築
	client, appErr := model.NewWebSocketClient4(MattermostWebSocketURL, AccessToken)
	if appErr != nil {
		log.Fatal(appErr.Error())
	}

  // (4) 接続
	if appErr = client.Connect(); appErr != nil {
		log.Fatal(appErr.Error())
	}
	client.Listen()

	for {
		select {
    case event := <-client.EventChannel:
      // (5) WebSocket Eventについての処理
			log.Printf("Received: %v", event)
		}
	}
}
```

`(1)`でMattermost WebSocketのURLを宣言しています。プロトコルが`http(s)`ではなく`ws(s)`になります。また、WebSocket APIのエンドポイントは`/api/v4/websocket`ですが、この部分はMattermostのWebSocket Driverが付与してくれます。
`(2)`ではWebSocket APIにアクセスするために使用するアクセストークンを宣言しています。REST API実行用と同じTokenです。

`(3)`で、MattermostのWebSocket Driverを利用してWebSocketクライアントを生成し、このクライアントを使って`(4)`でWebSocket APIへ接続しています。

接続が問題なく完了すると、 WebSocketクライアントの`EventChannel`フィールドにWebSocket Eventが流れてきます。`(5)`で、流れてきたイベントを出力しています。


上記コードを実行しておくと、Mattermost上で何か操作をした時にWebSocket Eventがコンソールに表示されるようになります。

![video](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day11/example-websocket-api.gif)

`Ctrl + c`を入力するとプログラムを終了します。

今回は単に受け取ったイベント情報を出力するだけでしたが、受け取った[WebSocketEvent](https://github.com/mattermost/mattermost-server/blob/c54ab6da211af861b25af1364518e549d3a1ed91/model/websocket_message.go#L109)の内容に応じた処理を実装することもできます。
`WebSocketEvent`構造体の`Data`フィールドの内容はイベントの種類によって異なり、イベントの種類は下記から確認できます。
https://github.com/mattermost/mattermost-server/blob/c54ab6da211af861b25af1364518e549d3a1ed91/model/websocket_message.go#L14


## WebSocket APIの実行
MattermostではWebSocket APIを通じてリクエストを送信することもできるようですが、現在利用可能なのが下記3種類のみのようで、ちょっと用途が今のところ分かりません。

* user_typing
* get_statuses
* get_statuses_by_id


## さいごに

Mattermost WebSocket APIの使い方について紹介しました。
明日は、Botアカウントの使い方を紹介します。
