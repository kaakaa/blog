---
title: "[Mattermost Integrations] Interactive Message Menu"
date: 2020-12-15T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "interactive_message", "menu"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第15日目の記事です。

本記事では、Mattermostの投稿にユーザーが操作できるセレクトボックスを追加するInteractive Message Menuの機能について紹介します。

## Interactive Message Menu概要

Interactive Message Menuは、昨日紹介したInteractive Message Buttontと同様の機能です。

Interactive Message Buttonは、Mattermostの投稿にボタンを表示する機能でしたが、Interactive Message Menuでは、Mattermostの投稿にセレクトボックスを表示することができます。

Interactive Messageに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/interactive-messages.html

### 作成

Interactive Message Menuを含む投稿を作成する方法はInteractive Message Buttonの時とほぼ同じです。

Incoming Webhook(内向きのウェブフック)を使って作成する例は下記のようになります。

```bash
BODY='{
  "attachments": [{
    "title": "Echo text",
    "actions": [{
	  "id": "echo",
      "name": "Would you like to echo?",
      "integration": {
        "url": "http://localhost:8080/actions/echo",
        "context": {
          "text": "sample text"
		}
	  },
	  "type": "select",
	  "options": [{
		  "text": "Echo",
		  "value": "echo"
	  }, {
		  "text": "Reject",
		  "value": "reject"
	  }]
	}]
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

上記のリクエストを実行すると、下記のようなセレクトボックスを含む投稿が作成されます。

![first exmple](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day15/first-example.png)

`options`に指定したオプションがセレクトボックスから選べるようになっています。

![first example2](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day15/first-example2.png)

セレクトボックスから要素を選択すると、対応するオプションの`value`に指定した値を含むリクエストが`Integration`に指定したURLへ送信されることになります。

#### パラメータ
`actions`フィールドに指定できるパラメータは、Interactive Message Buttonと同じ`PostAction`構造体として宣言されています。
https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L37

昨日の記事にも載せましたが、再掲します。

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

この中で、Interactive Message Menuのみで使用できるパラメータは`data_source`, `options`, `default_option`になります。

`options`は、先ほどの例にも示したようにセレクトボックスで選択する要素を定義するフィールドです。
`default_option`は、デフォルトで選択されているオプションを指定することができます。`options`内のいずれかのオプションの`value`の値を指定します。

`data_source`は特殊なオプションです。`users`か`channels`を指定でき、これを指定するとセレクトボックスから選択できるオプションがそれぞれMattermost上の**ユーザー**かMattermost上の**公開チャンネル**になります。`data_source`を指定した場合、`options`に指定されたオプションは無視されます。

![data source](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day15/data-source.png)

```bash
BODY='{
  "attachments": [{
    "title": "Echo text",
    "actions": [{
	  "id": "echo",
      "name": "Would you like to echo?",
      "integration": {
        "url": "http://localhost:8080/actions/echo",
        "context": {
          "text": "sample text"
		}
	  },
	  "type": "select",
	  "data_source": "channels"
	}]
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

### 実行

今回は、slash commandの引数に指定したテキストをセレクトボックスで選択されたチャンネルへEchoするようなInteractive Message Menuを紹介します。

![video](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day15/interactive-menu.gif)

slash command、Interactive Message Menuからのリクエストを受け取るサーバーアプリのサンプルコードは下記のようになります。

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
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		response := model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Attachments: []*model.SlackAttachment{{
				Title: "Echo server",
				Actions: []*model.PostAction{{
					// (1) Echoメニュー
					Name: "Echo",
					Integration: &model.PostActionIntegration{
						URL: "http://localhost:8080/actions/echo",
						Context: map[string]interface{}{
							"text": r.Form.Get("text"),
						},
					},
					Type: model.POST_ACTION_TYPE_SELECT,
					Options: []*model.PostActionOptions{{
						Text:  "Echo",
						Value: "echo",
					}, {
						Text:  "Reject",
						Value: "reject",
					}},
				}},
			}},
		}

		// Need to set header. if not, just json string will be posted.
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, response.ToJson())
	})

	// (2) Echo Buttonが押されたときの処理
	http.HandleFunc("/actions/echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (3) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.PostActionIntegrationRequest
		json.Unmarshal(b, &payload)

		text, ok := payload.Context["text"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['text'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		selected, ok := payload.Context["selected_option"].(string)
		if !ok {
			resp := &model.PostActionIntegrationResponse{EphemeralText: "invalid request. Context['selected_option'] is not found."}
			fmt.Fprint(w, resp.ToJson())
			return
		}

		// (4) レスポンスの構築
		response := &model.PostActionIntegrationResponse{}
		switch selected {
		case "echo":
			response.Update = &model.Post{
				Message: text,
				Props:   model.StringInterface{},
			}
		case "reject":
			response.Update = &model.Post{Props: model.StringInterface{}}
			response.EphemeralText = "Echoing was rejected"
		default:
			response.EphemeralText = "invalid operation"
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(response.ToJson()))
	})

	http.ListenAndServe(":8080", nil)
}
```

`(1)`ではInteractive Message Menuの内容を定義しています。`Integration.URL`にリクエスト送信先として`http://localhost:8080/actions/echo`を、`Integration.Context`に、slash commandの引数として指定された文字列を保持しています。また、`Echo`と`Reject`のオプションを`options`に定義しています。

`(2)`で、セレクトボックスからオプションが選択された時に送信されるリクエストを処理しています。`(3)`で送信されたリクエストを読み出しています。送信されるリクエストの形式は、Interactive Message Menuから送信されるリクエストもInteractive Message Buttonと同じく[`PostActionIntegrationRequest`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L173)として定義されています。

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

今回は`Context`から`text`をキーとして格納されているslash command実行時の引数と、`selected_option`をキーとして格納されている選択されたオプションを読み出しています。
そして、`(4)`で読み出した`selected_option`の値に応じてレスポンスを構築しています。レスポンスの形式もInteractie Message Buttonと同じく[`PostActionIntegrationResponse`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L187)として定義されています。

```go
type PostActionIntegrationResponse struct {
	Update           *Post  `json:"update"`
	EphemeralText    string `json:"ephemeral_text"`
	SkipSlackParsing bool   `json:"skip_slack_parsing"` // Set to `true` to skip the Slack-compatibility handling of Text.
}
```


## さいごに

本日は、Interactive Message Menuの使い方について紹介しました。
明日は、Interactive Dialogについて紹介します。
