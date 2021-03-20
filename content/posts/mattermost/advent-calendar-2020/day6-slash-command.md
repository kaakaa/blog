---
title: "[Mattermost Integrations] Slash Command 基本編"
date: 2020-12-06T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "command", "slashcommand"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第6日目の記事です。

本記事では、Mattermost上で任意のタイミングで特定の処理を実行できるSlash Command機能について紹介します。

## Slash Command概要

Slash Commandは、Mattermostの投稿入力欄で `/` で始まるコマンドを入力することで、特定の処理を実行できる機能です。
Slash Commandには、Mattermostにデフォルトで搭載されている**Built-in Slash Command（内蔵スラッシュコマンド）**と、ユーザーが独自にカスタマイズできる**Custom Slash Command（カスタムスラッシュコマンド）**があります。

投稿入力欄に何も入力していない状態で`/`を入力すると、利用できるスラッシュコマンドの一覧が表示されます。

![slash command](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/slash-command.png)

Slash Commandに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/slash-commands.html
* https://developers.mattermost.com/integrate/slash-commands/

一つ目の公式ドキュメントはSlash Command機能の概要について記述しており、二つ目のDeveloper Documentにはより細かい開発者向けの情報が書かれています。


## Built-in Slash Command
Built-in Slash Commandとして使用できるスラッシュコマンドは、下記のドキュメントに記載されています。

https://docs.mattermost.com/developer/slash-commands.html#built-in-commands

数が多いので全ては紹介しませんが、個人的には `/shortcut` と `/shrug` をよく使います。

`/shortcut`はMattermost上で使用できるショートカットを表示するコマンドです。ショートカットの中では、`Ctrl + k`によるチャンネル切り替えをよく使います。

![shortcut](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/shortcut.png)

`/shrug`は、`¯\_(ツ)_/¯` を簡単に入力するためのコマンドです。なんだかよく分からないけど上手くいかなかった時などに使います。

## Custom Slash Command
Custom Slash Commandは、ユーザーが独自のSlash Commandを作成できる機能です。Custom Slash Commandが実行されると、Custom Slash Command作成時に指定したURLへコマンド実行情報が送信されます。

![overview](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/overview.drawio.png)

Custom Slash Commandは、外部アプリケーションにリクエストを送信するという点でOutgoing WebHookと似ていますが、ユーザーが明示的にコマンドを実行できる点や、Outgoing WebHookでは利用できなかった非公開チャンネルやダイレクトメッセージチャンネル内でも利用できるという点が異なります。

### 設定

Custom Slash Commandを利用するには、**システムコンソール > 統合機能管理 > カスタムスラッシュコマンドを有効にする** の設定が`有効`になっている必要があります。

![system console](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/config-slash-command.png)

また、統合機能はデフォルトではシステム管理者とチーム管理者しか作成することができませんが、誰でも作成できるようにしたい場合、**システムコンソール > 統合機能管理 > 統合機能の管理を管理者のみに制限する**の設定を`無効`してください。

### 作成

**メインメニュー > 統合機能**から統合機能の画面を開き、

![main menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/integration-menu.png)

**スラッシュコマンド > スラッシュコマンドを追加する**から、新たなSlash Commandを追加します。

![menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/slash-command-menu.png)

Slash Commandの作成画面で入力する情報は下記の通りです。

* **タイトル**: ウェブフックの一覧ページに表示されるタイトルです
* **説明**: ウェブフックの説明です
* **コマンドトリガーワード**: ここで指定したキーワードがスラッシュコマンド名として使われます。日本語も使用できます。
* **リクエストURL**: リクエスト送信先のURLです
* **リクエストメソッド**: 送信されるリクエストを`GET`、`POST`から選択できます
* **返信ユーザー名**: Slash Command実行によって作成される投稿のユーザー名部分に表示される名前を指定できます
* **応答アイコン**: Slash Command実行によって作成される投稿のユーザーアイコン部分に表示されるアイコン画像を指定できます
* **自動補完**: 投稿入力欄に`/`を入力した時に表示されるSlash Command一覧画面にこのコマンドを表示するかどうかを選択できます。**自動補完**にチェックをした場合、Slash Command一覧に表示される**自動補完のヒント**と**自動補完の説明**を入力できるようになります。

Slash Commandの作成が完了すると、トークンが表示されます。このトークンは、Outgoing WebHookのトークンと同様、リクエストがMattermostから送信されたことを検証するために使われます。

![complete](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/complete-slash-command.png)

### 実行

Slash Commandを実行する前に、リクエスト送信先サーバーを立ち上げておく必要があります。
Outgoing WebHookの時と同様、送信されたリクエストのJSONボディをログ出力する簡単なサーバーを立ち上げておきます。Outgoing WebHookの時との違いは、送信されるリクエストボディの形式が異なるのと、エンドポイントを変更している点です。

```go:main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		// (1) Formデータの読み込み
		r.ParseForm()
		log.Printf("Request: %#v", r.Form)
	})
	http.ListenAndServe(":8080", nil)
}
```

`(1)` Slash Commandにより送信されるリクエストは、`application/x-www-form-urlencoded`形式で送信されるため、Formデータとして読み出しています。

上記コードを`main.go`として保存し、`go run main.go`を実行してサーバーを立ち上げます。もし、MattermostサーバーとSlashCommandリクエスト受付用のサーバを同じマシン上で起動し、Outgoing WebHook作成時の**コールバックURL**に`localhost`のサーバーを指定している場合、**システムコンソール > 開発者 > 信頼されていない内部接続を許可する**に`localhost`を追加しておく必要があります。

![config allow internal](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/config-allow-internal.png)

上記の設定が完了した後、Slash Commandを実行すると下記のようなリクエストの情報がコンソールに出力されます。

![execute](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day6/execute-slash-command.png)

```bash
$ go run main.go 
2020/11/03 15:04:24 Request: url.Values{"channel_id":[]string{"xxyu9xoref8mjgy3s9i5y7776y"}, "channel_name":[]string{"slash-command"}, "command":[]string{"/サンプルコマンド"}, "response_url":[]string{"http://localhost:8065/hooks/commands/5xthz8jf67ggx8heopn1ay1tqe"}, "team_domain":[]string{"test"}, "team_id":[]string{"9d1xf4gg7fnibxs8fdw6fo5fre"}, "text":[]string{""}, "token":[]string{"8w7foap4ufrsfczda8uez51yxo"}, "trigger_id":[]string{"bm1udGhmOGs3ZmJmeGpmb3dnODZnY2NzaWU6ODd4OTN1bzhwZm56ZHJvOWt0Y21vYnBhMXI6MTYwNDM4MzQ2NDI4MTpNRVVDSUZ5Vjc0NmwrWWt3UVUrUkwrUzFaYWlTMStnYTZPa2ZsSzJtSTBBL2wzNThBaUVBOFpTY2hjblZLa05scVU0MVNmc2l0cEdGanpXWE9tREdKK2NTeUFiRURHYz0="}, "user_id":[]string{"87x93uo8pfnzdro9ktcmobpa1r"}, "user_name":[]string{"kaakaa"}}
```

このように、Slash Commandを実行することで外部アプリケーションへリクエストを送信することができます。外部アプリケーション側で、送信されたリクエストに応じた処理を実装することで、Mattermostから処理を起動したりすることができます。動作的にはOutgoing WebHookと似ていますが、Slash Commandはチーム内のどのチャンネルからも実行することができ、チャンネルやトリガーワードの設定も存在しないため作成者でなくても使いやすいという利点があります。
また、レスポンスによる投稿作成時に利用できる機能もOutgoing WebHookよりも多いですが、その辺りは明日の記事で紹介します。

## さいごに

本日は、Slash Commandの基本的な使い方について紹介しました。
明日は、Slash Commandのレスポンスとして利用できるパラメータについて紹介します。
