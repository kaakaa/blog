---
title: "[Mattermost Integrations] Slash Command 発展編"
date: 2020-12-07T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "command", "slashcommand"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第7日目の記事です。

前回の記事では、Mattermostで`/`で始まるコマンドを実行することで、特定の処理を実行することができるSlash Commandの機能について紹介しました。また、Custom Slash Commandにより外部アプリケーションにリクエストを送信する方法を紹介しました。本記事では、Custome Slash Commandのリクエストが送信された外部アプリケーション側からMattermost側へ返却するレスポンス内容について紹介します。


## Slash Commandのレスポンスとして投稿を作成する

Slash Commandによるリクエストを受け付けた外部アプリケーションから、レスポンスとしてMattermostに投稿を作成する方法について紹介します。

以下のコードは昨日のコードにレスポンスを返すコードを追加したものです。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	WebHookToken = "8w7foap4ufrsfczda8uez51yxo"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%#v", r.Form)
		r.ParseForm()

		log.Printf("token: %v", r.Form.Get("token"))
		// (1) Tokenをチェック
		if r.Form.Get("token") != WebHookToken {
			log.Printf("received an invalid request with token: %s", r.Form.Get("token"))
			return
		}

		// (2) レスポンスとしてMattermostへ投稿を作成
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf(`{"text": "**%s**"}`, r.Form.Get("text")))
	})
	http.ListenAndServe(":8080", nil)
}
```

Outgoing WebHookの時とほぼ同じですが、`(1)` でトークンを利用したリクエストの検証を行い、`(2)` でSlash Commandのレスポンスとして投稿を作成しています。レスポンスヘッダーへの`Content-Type: application/json`の設定は必須です。`text`フィールドに指定したメッセージがMattermostに投稿されます。

このサーバに対して`/サンプルコマンド TEST`というSlash Commandを送信した時の画面が下記になります。


![response](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/response-slash-command.png)

作成された投稿には、`BOT`というラベルに加え、`(あなただけが見ることができます)`というメッセージがついています。Slash Commandのレスポンスとして作成した投稿は、デフォルトではSlash Commandを実行したユーザーしか見ることができない投稿になり、他のユーザーからは見ることができません。チャンネルの全員が閲覧可能な投稿を作成するには、後述の`response_type`パラメータを設定します。


## Slash Commandレスポンスパラメータ

### `text`

投稿されるメッセージであり、後述の`attachments`を指定しない場合は必須のパラメータです。Markdown記法や、`@`によるメンションを投稿することもできます。

```go
io.WriteString(w, `{"text": "**TEST**"}`)
```

### `response_type`

Slash Commandのレスポンスにより作成される投稿の種別を指定します。デフォルトでは`ephemeral`であり、Slash Commandを実行したユーザーにしか見えない投稿になります。チャンネル全員が見えるようにするには`response_type`に`in_channel`を指定します。

```go
io.WriteString(w, `{"text": "**TEST**", "response_type": "in_channel"}`)
```

### `username`, `icon_url`

Slash Commandのレスポンスとして作成される投稿のユーザー名とアイコンの表示を変更することができます。デフォルトでは、ウェブフックを作成したユーザーになります。

このパラメータを使用する場合、システムコンソールから下記の設定をそれぞれ`有効`にする必要があります。
* **システムコンソール > 統合機能管理 > 統合機能によるユーザー名の上書きを許可する**
* **システムコンソール > 統合機能管理 > 統合機能によるプロフィール画像アイコンの上書きを許可する**

（この設定を有効にすると、`username`と`icon_url`を指定しない場合の投稿の作成者がデフォルトのアイコン、ユーザー名(`webhook`)に変更されます）

```go
io.WriteString(w, `{
	"text": "**TEST**",
	"username": "new-bot", 
	"icon_emoji": "http://www.mattermost.org/wp-content/uploads/2016/04/icon_WS.png"
}`)
```
![use overwrite](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/example-overwrite.png)

### `channel_id`

投稿を作成するチャンネルを指定できます。ここで指定するチャンネルIDは、**チャンネルメニュー > 情報を表示する**から確認できます。

![channel info view](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/view-channel-id1.png)

![channel info view](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/view-channel-id2.png)

```go
io.WriteString(w, `{
	"text": "**TEST**",
	"channel_id": "u6ijofzg8bnsie5u8c1eumtc1h"
}`)
```

### `goto_location`

Slash Command特有のパラメータです。`goto_location`フィールドにURLを指定しておくと、Slash Commandが実行された際に指定されたURLに移動します。この時`text`などで指定した投稿も作成されます。

```go
io.WriteString(w, `{
	"text": "**TEST**",
	"goto_location": "https://github.com/mattermost"
}`)
```

![movie](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/example-goto-location.gif)

`goto_location`には`https`だけでなく、`ssh`や`mailto`のプロトコルも使用できるため、例えばスラッシュコマンドを実行した場合に「メーラーでメール作成画面を開く」というような動作をさせることも可能です。

```go
io.WriteString(w, `{
	"text": "**TEST**",
	"goto_location": "mailto:test@example.com"
}`)
```

### `attachments`

`attachments`は、Mattermostにリッチな投稿を作成することができる[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)を指定できる機能です。[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)については、別途紹介記事を用意する予定です。

### `type`
`type`は、Mattermostの投稿の種別を指定するパラメータですが、基本的に使用することはありません。Plugin(Webapp)にてプラグイン独自の投稿種別を指定した場合のみに使用するものかと思います。Plugin(Webapp)については、別途紹介記事を用意する予定です。


### `extra_responses`

`extra_responses`配列としてレスポンスのJSONデータを記述することで複数の投稿を作成することができます。
`extra_responses`が導入されるMattermost v5.6以前は、Slash Commandのレスポンスとして複数の投稿を作成したい場合はREST APIを実行するなどをしなくてはなりませんでしたが、`extra_responses`によってREST APIを実行するためのコードなどが必要なくなりました。

Slash Commandのリクエストに対して下記のレスポンスを返却すると、「チャンネル全員が閲覧できる投稿を作りつつ、コマンド実行ユーザーのみにしか見えない投稿も合わせて作成し、さらに、別のチャンネルにコマンドが実行されたことを通知する」ということができるようになります。

```go
io.WriteString(w, `{
	"text": "**TEST**",
	"response_type": "in_channel",
	"extra_responses": [
		{
			"text": "**Secret Info**"
		},
		{
			"text": "Notify in other channel",
			"channel_id": "u6ijofzg8bnsie5u8c1eumtc1h"
		}
	]
}`)
```

`extra_responses`内にさらに`extra_responses`を指定することはできず、また、`goto_location`も使用できません。

### `skip_slack_parsing`

Mattermostの統合機能はSlack記法との互換性を意識して作成されていますが、`skip_slack_parsing`に`true`を指定すると、レスポンスとして返却した`text`をSlack互換の記法として解析しないようになります。あまり使うことは無いかと思いますが、レスポンスが想定どおりに投稿されない場合などに利用を検討してみても良いかもしれません。

このパラメータが追加されたきっかけとなったのは下記のIssueです。
[slack\-compatibility\-layer seems to break results of MM\-only integration · Issue \#12702 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/12702)

### `props`
Incoming WebHookなどと同様、投稿のメタデータを格納するフィールドです。Incoming WebHookと同様、`card`も利用することができます。

## Slash Commandで時間のかかる処理を実行する場合

単にSlash Commandのレスポンスとして投稿を作成できるだけだと、外部アプリケーションで時間のかかる処理を実行し、その結果を返したいという場合、Mattermost上のユーザーはコマンドを実行してから結果が返るまで何の反応も得ることができません。それを回避するため、Slash Command実行時に送信されるリクエストには、レスポンスを一度返した後で投稿を作成ことができるURLである`response_url`フィールドを持っています。

`response_url`を使ったコードは下記のようになります。

```go:main.go
package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	WebHookToken = "8w7foap4ufrsfczda8uez51yxo"
)

func main() {
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("token") != WebHookToken {
			log.Printf("received an invalid request with token: %s", r.Form.Get("token"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"text": "Request is recieved. Results will return later."}`)

		// (1) 5秒後に "response_url" に対してリクエストを送信
		go func(url string) {
			time.Sleep(5 * time.Second)
			body := strings.NewReader(`{"text": "This is the result.", "response_type": "in_channel"}`)
			http.Post(url, "application/json", body)
		}(r.Form.Get("response_url"))
	})
	http.ListenAndServe(":8080", nil)
}
```

今までと同じようにレスポンスを返した後、5秒Sleepした後に`response_url`のURLに対してリクエストを送信しています

![movie](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day7/example-delayed-response.gif)

`response_url`に送信するリクエストには、`goto_location`は使えませんでしたが、`response_type`や`extra_responses`などは使えるようです。

また、`response_url`が有効なのは、元のSlash Commandが実行されてから30分間のみで、5つまでの投稿を作成することができます。処理が30分以上かかる場合や、5つ以上の投稿を作成したい場合は、REST APIを利用するか、Pluginを利用する必要がありそうです。

## さいごに

本日は、Slash Commandの詳細な使い方について紹介しました。
明日は、REST APIの使い方について紹介していきます。
