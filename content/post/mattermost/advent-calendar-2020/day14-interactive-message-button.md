---
title: "[Mattermost Integrations] Interactive Message Button"
date: 2020-12-14T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "interactive_message", "button"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第14日目の記事です。

本記事では、Mattermostの投稿にユーザーが操作できるボタンを追加するInteractive Message Buttonの機能について紹介します。

## Interactive Message Button概要

Interactive Message Buttonは、Mattermostの投稿にボタンを表示し、ボタンがクリックされるとどのボタンが押されたのかという情報がサーバーへ送信される機能です。

![official example](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day14/official-example.png) 
(画像は[公式ドキュメント](https://docs.mattermost.com/developer/interactive-messages.html)から)

Interactive Messageに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/interactive-messages.html

### 作成

Interactive Message Buttonを含む投稿を作成するには、Message Attachments機能と同様`attachments`にButtonに関する情報を追加します。

Incoming Webhook(内向きのウェブフック)を使って作成する例は下記のようになります。

```bash
BODY='{
  "attachments": [{
    "title": "Echo text",
    "actions": [{
	  "id": "echo",
      "name": "Echo",
      "integration": {
        "url": "http://localhost:8080/actions/echo",
        "context": {
          "text": "sample text"
		}
	  }
	}, {
      "name": "Reject",
	  "style": "danger",
      "integration": {
        "url": "http://localhost:8080/actions/reject"
	  }
	}]
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

上記のリクエストを実行すると、下記のようなボタンを含む投稿が作成されます。

![first example](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day14/example-button.png)

リクエスト内容と見ると、`actions`フィールドに含まれる下記の一つのオブジェクトが一つのボタンに関連していることが分かります。

```bash
{
  "id": "echo",
  "name": "Echo",
  "integration": {
    "url": "http://localhost:8080/actions/echo",
    "context": {
      "text": "sample text"
    }
  }
}
```

`id`はボタンのID、`name`は投稿上に表示されるボタンのラベルを指定します。`integration`内には、ボタンを押した時に送信されるリクエストの情報を記述しており、`url`にはリクエスト送信先のURLを、`context`にはリクエストに含まれる任意の情報を指定することができます。

ボタンが押された時に何か処理を行うには、リクエストを受け取るサーバーを用意する必要がありますが、それについては後述します。

#### パラメータ
`actions`フィールドに指定できるパラメータは、Mattermostのサーバー側のコードに`PostAction`構造体として宣言されています。
https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L37

* `id`: 1投稿内でユニークとなるボタンのIDを指定します。指定しない場合、Mattermost側で一意のIDが付与されます。
* `type`: Interactive Messageの種別を指定します。`select`か`button`を指定でき、何も指定しないと`button`が指定されます。
* `name`: Mattermost内でボタンのラベルとして表示される文字列を指定します。
* `disabled`: ボタンを無効化するオプションです。`true`を指定するとボタンが押せなくなります。
* `style`: ボタンの色を指定できます。`default`, `primary`, `success`, `good`, `warning`, `danger`, もしくは任意のHex color(`"#FF9900"`など)を指定できます。
Outgoing WebHookを利用するには、**システムコンソール > 統合機能管理 > 外向きのウェブフックを有効にする** の設定が`有効`になっている必要があります。
* `data_source`: Interactive Message **Button**では使用しません。Interactie Message **Menu**でのみ使用します。
* `options`: Interactive Message **Button**では使用しません。Interactie Message **Menu**でのみ使用します。
* `default_option`: Interactive Message **Button**では使用しません。Interactie Message **Menu**でのみ使用します。
* `integration`: アクションを実行した際に送信されるリクエストに関する情報を指定します。
  * `url`: リクエスト送信先のURLを指定します。
  * `context`: 送信されるリクエストに含まれる追加情報を指定します
  * `integration`に指定した情報はリクエスト送信時にしか読み出されないため、(httpsなどで通信経路のセキュリティが担保されていれば)アクセストークンのような秘密情報なども含めることができます。


### 実行

Interactive Message Buttonを機能させるには、ボタンが押された時に送信されるリクエストを処理するサーバーを立てておく必要があります。今回は、slash commandの引数に指定したテキストをEchoするようなInteractive Message Buttonについて考えます。

![video](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day14/interactive-button.gif)

slash command、Interactive Message Buttonからのリクエストを受け取るサーバーアプリのサンプルコードは下記のようになります。

(今回の例ではslash command、Interactive Message Button共に、Mattermostから`localhost`に対してリクエストが送信されるため、**システムコンソール > 開発者 > 信頼されていない内部接続を許可する**の設定に、リクエスト送信先のサーバーを記述しておく必要があります。詳しくは[公式ドキュメント](https://docs.mattermost.com/administration/config-settings.html#allow-untrusted-internal-connections-to)を参照ください。)

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

func main() {
	// (1) Interactive Message Button作成用slash command
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Attachments: []*model.SlackAttachment{{
				Title: "Echo text",
				Actions: []*model.PostAction{{
					// (2) Echoボタン
					Id:   "echo",
					Name: "Echo",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/echo",
						Context: map[string]interface{}{
							"text": r.Form.Get("text"),
						},
					},
				}, {
					// (3) Rejectボタン
					Name:  "Reject",
					Style: "danger",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/reject",
					},
				}},
			}},
		}

		// Need to set header. if not, just json string will be posted.
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})

	// (4) Echo Buttonが押されたときの処理
	http.HandleFunc("/actions/echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (5) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.PostActionIntegrationRequest
		json.Unmarshal(b, &payload)
		text, ok := payload.Context["text"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['text'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		// (6) レスポンスの構築
		response := &model.PostActionIntegrationResponse{
			Update: &model.Post{
				Message: text,
				Props:   model.StringInterface{},
			},
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	// (7) Rejectボタンが押されたときの処理
	http.HandleFunc("/actions/reject", func(w http.ResponseWriter, r *http.Request) {
		// (8) レスポンスの構築
		response := &model.PostActionIntegrationResponse{
			Update: &model.Post{
				Props: model.StringInterface{},
			},
			EphemeralText: "Echoing was rejected.",
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	http.ListenAndServe(":8080", nil)
}
```

`(1)`はslash commandのリクエストを処理するロジックです。`http://localhost:8080/command`へリクエストを送信するslash commandを事前に登録しておく必要があります。slash commandの登録方法については6日目の記事を参照ください。
`(2)`では、slash command実行時の引数情報をEchoする**Echo**ボタンを定義しています。`Integration.URL`にリクエスト送信先として`http://localhost:8080/actions/echo`を、`Integration.Context`に、slash commandの引数として指定された文字列を保持しています。`(3)`では、Echo処理を注意するための**Reject**ボタンを定義しています。リクエスト送信先は`http://localhost:8080/actions/reject`にしてあります。また、**Reject**ボタンには`"style": "danger"`をしているため、ボタンが赤く表示されます。

slash commandを実行すると、下記のような投稿が作成されます。

![example](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day14/example-button.png) 

`(4)`で、**Echo**ボタンが押された時に送信されるリクエストを処理しています。まず、`(5)`で送信されたリクエストを読み出しています。Interactive Message Buttonから送信されるリクエストはMattermostのコードでは[`PostActionIntegrationRequest`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L173)として定義されています。

```go
type PostActionIntegrationRequest struct {
	UserId      string                 `json:"user_id"`
	UserName    string                 `json:"user_name"`
	ChannelId   string                 `json:"channel_id"`
	ChannelName string                 `json:"channel_name"`
	TeamId      string                 `json:"team_id"`
	TeamName    string                 `json:"team_domain"`
	PostId      string                 `json:"post_id"`
	TriggerId   string                 `json:"trigger_id"`
	Type        string                 `json:"type"`
	DataSource  string                 `json:"data_source"`
	Context     map[string]interface{} `json:"context,omitempty"`
}
```

今回は`Context`内に`text`をキーとしてslash command実行時の引数が格納されているため、その文字列を読み出しています。
そして、`(6)`で読み出した文字列を使ってレスポンスを構築しています。レスポンスの形式はMattermostのコードで[`PostActionIntegrationResponse`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L187)として定義されています。

```go
type PostActionIntegrationResponse struct {
	Update           *Post  `json:"update"`
	EphemeralText    string `json:"ephemeral_text"`
	SkipSlackParsing bool   `json:"skip_slack_parsing"` // Set to `true` to skip the Slack-compatibility handling of Text.
}
```

`PostActionIntegrationResponse`の`Update`フィールドに指定した`Post`のインスタンスで、Interactive Message Buttonを置き換えることができます。今回はslash commandの引数として与えられた文字列をメッセージに持つ投稿に置き換えています。`Props: model.StringInterface{},`で`Post`の`Props`を空にしていますが、これをしないとInteractive Message Buttonがずっと残り続けてしまうため、ボタンを削除したい場合は必ず空に設定する必要があります。Mattermost v5.10以前では、明示的に`Props`を空にしなくてもInteractive Message Buttonが削除されていましたが、Mattermost v5.11以上では、明示的に空に設定しないと削除されないようになったようです。(https://docs.mattermost.com/developer/interactive-messages.html#how-do-i-manage-properties-of-an-interactive-message)


`(7)`は、**Reject**ボタンが押されたときの処理です。`(8)`でレスポンスを構築していますが、ボタンが押された時にInteractive Message Buttonを含む投稿を削除したい場合は、このように`Props`を空に設定しただけの`Post`インスタンスを指定します。`EphemeralText`に指定した文字列は、ボタンを押した人だけに見えるメッセージになります。


## さいごに

本日は、Interactive Message Buttonの使い方について紹介しました。
明日は、Interactive Message Menuについて紹介します。
