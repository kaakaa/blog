---
title: "[Mattermost Integrations] Outgoing WebHook 基本編"
date: 2020-12-04T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "webhook", "outgoing"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第4日目の記事です。

本記事では、Mattermostから外部アプリケーションへ投稿内容を送信できるOutgoing WebHook（外向きのウェブフック）について紹介します。

## Outgoing WebHook概要

Outgoing WebHookは、Mattermostに投稿が作成された時にその投稿情報を外部アプリケーションへ送信するための機能です。

![overview](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/overview.drawio.png)

Outgoing WebHookを使うことで、Mattermostで発生した投稿イベントを外部アプリケーションに通知することができます。

Outgoing WebHookに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/webhooks-outgoing.html
* https://developers.mattermost.com/integrate/outgoing-webhooks/

一つ目の公式ドキュメントはOutgoing WebHookの概要について記述しており、二つ目のDeveloper Documentにはより細かい開発者向けの情報が書かれています。

### 設定

Outgoing WebHookを利用するには、**システムコンソール > 統合機能管理 > 外向きのウェブフックを有効にする** の設定が`有効`になっている必要があります。

![system console](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/config-outgoing-webhook.png)

また、統合機能はデフォルトではシステム管理者とチーム管理者しか作成することができませんが、誰でも作成できるようにしたい場合、**システムコンソール > 統合機能管理 > 統合機能の管理を管理者のみに制限する**の設定を`無効`してください。

### 作成

**メインメニュー > 統合機能**から統合機能の画面を開き、

![main menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/integration-menu.png)

**外向きのウェブフック > 外向きのウェブフックを追加する**から、新たなOutgoing WebHookを追加します。

![menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/outgoing-webhook-menu.png)

ウェブフックの作成画面で入力する情報は下記の通りです。

![outgoing webhook](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/create-outgoing-webhook.png)

* **タイトル**: ウェブフックの一覧ページに表示されるタイトルです
* **説明**: ウェブフックの説明です
* **コンテントタイプ**: 外部アプリケーションに送信するデータの形式を`application/x-www-form-urlencoded` と `application/json` から選択できます
* **チャンネル**: ここで指定したチャンネルに投稿が作成されると、外部アプリケーションへのリクエストが送信されます。トリガーワードを指定した場合、チャンネルの指定は必須ではなくなります。非公開チャンネルやダイレクトメッセージチャンネルは指定できません。
* **トリガーワード**: ここで指定したキーワードで始まる投稿が外部アプリケーションへ送信されます。チャンネルとトリガーワードを両方指定した場合、指定したチャンネル内のトリガーワードで始まる投稿のみが外部アプリケーションに送信されます。
* **トリガーとなる条件**: トリガーワードと一致する条件を指定します。スペース区切りで最初の単語がトリガーワードと完全に一致するか、トリガーワードで始まるかを設定するものであるため、日本語の投稿を対象とする場合、`最初の単語がトリガーワードで始まる`に設定しておくと良いかと思います。
* **コールバックURL**: リクエスト送信先のURLです。

Outgoing WebHookの作成が完了すると、トークンが表示されます。このトークンは、外部アプリケーションに対して送信されるリクエストに含まれる値であり、リクエストがMattermostから送信されたことを検証するために使われます。

![complete](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/complete-outgoing-webhook.png)

### 実行

Outgoing WebHookを実行する前に、リクエスト送信先サーバーを立ち上げておく必要があります。
今回は、送信されたリクエストのJSONボディをログ出力する簡単なサーバーを立ち上げておきます。

```go:main.go
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
		// (1) `model.OutgoingWebhookPayload`によるリクエストの読み出し
		var payload model.OutgoingWebhookPayload
		json.Unmarshal(b, &payload)

    	log.Printf("Request: %#v", payload)
  })
  
	http.ListenAndServe(":8080", nil)
}
```

`(1)` Outgoing WebHookにより送信されるJSONデータは、Mattermost本体の [`model.OutgoingWebhookPayload`](https://github.com/mattermost/mattermost-server/blob/master/model/outgoing_webhook.go#L35) の形式で送信されるため、この構造体を利用してデータを変換しています。

上記コードを`main.go`として保存し、`go run main`を実行してサーバーを立ち上げます。MattermostサーバーとOutgoing WebHookリクエスト受付用のサーバを同じマシン上で起動し、Outgoing WebHook作成時の**コールバックURL**に`localhost`のサーバーを指定している場合、**システムコンソール > 開発者 > 信頼されていない内部接続を許可する**に`localhost`を追加しておく必要があります。

![config allow internal](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/config-allow-internal.png)

上記の設定が完了した後、Outgoing WebHook作成時に**チャンネル**に指定したチャンネルに投稿を作成すると、

![trigger](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day4/trigger-outgoing-webhook.png)

Outgoing WebHookがサーバに送信され、下記のようなリクエストの情報がコンソールに出力されます。

```bash
$ go run main.go 
2020/11/03 00:10:30 Request: model.OutgoingWebhookPayload{Token:"9mgpebi9a7bq3qjedd1kt6mwtr", TeamId:"9d1xf4gg7fnibxs8fdw6fo5fre", TeamDomain:"test", ChannelId:"9eexapjuabd89fzbwfajdqhwta", ChannelName:"outgoing-webhook", Timestamp:1604329830865, UserId:"87x93uo8pfnzdro9ktcmobpa1r", UserName:"kaakaa", PostId:"au6tf4hoebyeffiaw9h1w6rpaw", Text:"こんにちは、テスト。", TriggerWord:"こんにちは、", FileIds:""}
```

リクエストに含まれるトークンの値(`Token:"9mgpebi9a7bq3qjedd1kt6mwtr"`)がOutgoing WebHook作成時に表示されていたトークンと同じものであることがわかります。


このように、Mattermostの投稿情報を外部アプリケーションへ送信することができます。外部アプリケーション側で、送信されたリクエストに応じた処理を実装することで、Mattermostから処理を起動したりすることができます。また、ウェブフックを受け取った外部アプリケーションから、再びMattermostへ投稿を作成したりすることもできますが、その辺りは明日の記事で紹介します。

## さいごに

本日は、Outgoing WebHookの基本的な使い方について紹介しました。
明日は、Outgoing WebHookに反応して外部アプリケーションからMattermostへ投稿を作成する方法について紹介します。