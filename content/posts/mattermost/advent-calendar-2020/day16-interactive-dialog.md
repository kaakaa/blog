---
title: "[Mattermost Integrations] Interactive Dialog"
date: 2020-12-16T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "interactive_dialog"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第16日目の記事です。

本記事では、Mattermostでユーザーの入力を受け付けるダイアログを表示するInteractive Dialogの機能について紹介します。

## Interactive Dialogの概要

Interactive Dialogは、Slash CommandやInteractive Messageなどのアクションを起点に、Mattermost上にユーザー入力を受け付けるダイアログ(モーダルウィンドウ)を表示する機能です。

![official example](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day16/official-example.png)
(画像は[公式ドキュメント](https://docs.mattermost.com/developer/interactive-dialogs.html)から)

Interactive Dialogに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/interactive-dialogs.html

Interactie Dialogは、何度かMattermostとインタラクションをしながら動作するもののため、動作が複雑になります。今までのように`curl`だけで動作させることは難しいため、Goのコードで書いたものを断片的に紹介していきます。

今回は、Interactive Dialogの入力内容からMessage Attachmentsのメッセージを作成するような例を考えてみます。

### Trigger IDの取得

Interactive Dialogを起動するには、まず、Mattermost内部で生成されるTrigger IDというものが必要です。Trigger IDはSlash CommandやInteractive Messageのアクションを実行した時に、Mattermostから送信されるリクエストに含まれています。Slash Command実行時のリクエストからTrigger IDを取得する場合、Slash Command実行時に送信されるリクエストを処理するサーバーで、以下のようにTrigger IDを取得することができます。

```go
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// (1) Slash Command実行時に送信されるリクエストから "Trigger ID" を取得
		triggerId := r.Form.Get("trigger_id")

		...
```

Interactive Message Buttonなどのアクションから取得する際は、[`PostActionIntegrationRequest.TriggerId`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L173)からTrigger IDを取得できます。

### Interactive Dialogの起動

先ほど取得したTrigger IDを使って、MattermostへInteractive Dialog起動のリクエストを投げます。
Trigger IDを取得するコードに続けて、[`/api/v4/actions/dialogs/open`](https://api.mattermost.com/#tag/integration_actions/paths/~1actions~1dialogs~1open/post)に[`OpenDialogRequest`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L224)で定義されるリクエストを送信することでInteractive Dialogを起動することができます。

```go
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// (1) Slash Command実行時に送信されるリクエストから "Trigger ID" を取得
		triggerId := r.Form.Get("trigger_id")

		// (2) Interactive Dialogを起動するためのリクエストを構築
		request := model.OpenDialogRequest{
			TriggerId: triggerId,
			URL:       "http://localhost:8080/actions/dialog",
			Dialog: model.Dialog{
				Title: "Sample Interactive Dialog",
				Elements: []model.DialogElement{{
					DisplayName: "Title",
					Name:        "title",
					Type:        "text",
				}, {
					DisplayName: "Message",
					Name:        "message",
					Type:        "textarea",
				}},
			},
		}

		// (3) Interactive Dialogを開く
		b, _ := json.Marshal(request)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8065/api/v4/actions/dialogs/open", bytes.NewReader(b))
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)

		...
```

`(2)`で構築している`OpenDialogRequest`にどのようなダイアログを表示するかという情報も指定するのですが、詳しくは後述します。
`(3)`で[`/actions/dialogs/open`](https://api.mattermost.com/#tag/integration_actions/paths/~1actions~1dialogs~1open/post)にリクエストを送信していますが、ここではAccessTokenなどが必要ありません。これはTrigger ID自体の利用可能期限が3秒と短く、悪用の心配がないためだと思われます。この点は、Trigger IDを取得してからダイアログを開く前に時間のかかる処理などを入れないよう注意する必要があるということも意味します。

`/actions/dialogs/open`へのリクエストが正常に完了すると、Mattermost上でInteractive Dialogが表示されます。

![video](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day16/interactive-dialog.gif)

#### Interactive Dialog起動時のパラメータ
Interactive Dialogを起動する際に送信する[`OpenDialogRequest`](https://github.com/mattermost/mattermost-server/blob/master/model/integration_action.go#L224)に与えることができるパラメータは下記の通りです。

* `TriggerId`: Slash CommandやInteractive Messageのアクションを実行した時にMattermost内部で生成されるInteractive Dialog起動用のIDを指定します
* `URL`: Interactive Dialogに入力された情報の送信先URLを指定します
* `Dialog`: Interactive Dialog上に表示される要素を指定します
  * `CallbackId`: 統合機能で設定されるIDです。Slash Commandの場合は`CommandArgs.RootId`、Interactive Messageの場合は`PostActionIntegrationRequest.PostId`を指定している気がしますが、何に使われているかはいまいちわかりません。
  * `Title`: Interactive Dialogのタイトル部分に表示されるテキストを指定します
  * `IntroductionText`: `Title`の下に表示されるダイアログの説明文を指定します
  * `IconURL`: ダイアログに表示されるアイコンのURLを指定します
  * `SubmitLabel`: ダイアログの決定ボタンのラベルを指定します
  * `NotifyOnCancel`: ダイアログのキャンセルボタンが押された時に、サーバーにその旨を通知するかを選択します。`true`の場合、キャンセル通知がサーバーに送信されます
  * `State`: 統語機能によって処理の状態を管理したい場合に設定される任意のフィールドです
  * `Elements`: ダイアログ上の入力フィールドを指定します。利用可能な`Element`については[公式ドキュメント](https://docs.mattermost.com/developer/interactive-dialogs.html#elements)を参照してください。

### Interactive Dialogからのリクエスト受信

Interactive Dialogの送信ボタンが押されると、`OpenDialogRequest`の`URL`フィールドに指定したURLへリクエストが送信されます。

```go
		// (2) Interactive Dialogを起動するためのリクエストを構築
		request := model.OpenDialogRequest{
			TriggerId: triggerId,
			URL:       "http://localhost:8080/actions/dialog",
			...
```

送信されるリクエストはMattermostのコードでは[`SubmitDialogRequest`](https://github.com/mattermost/mattermost-server/blob/dc1b42390b9bca393d03e2ccdbb16d66cd866431/model/integration_action.go#L230)として定義されています。

```go
type SubmitDialogRequest struct {
	Type       string                 `json:"type"`
	URL        string                 `json:"url,omitempty"`
	CallbackId string                 `json:"callback_id"`
	State      string                 `json:"state"`
	UserId     string                 `json:"user_id"`
	ChannelId  string                 `json:"channel_id"`
	TeamId     string                 `json:"team_id"`
	Submission map[string]interface{} `json:"submission"`
	Cancelled  bool                   `json:"cancelled"`
}
```

ユーザーがInteractive Dialog上で入力したデータは `Submission` に格納されています。`Submission`は`OpenDialogRequest`内の`DialogElement`の`Name`をkey、入力データをvalueとしたmap形式のデータです。

今回のInteractive Dialogでは、`title`と`message`という`Name`を持つ`DialogElement`を指定しているため、`Submission`からはこれらの値をキーとするValueが格納されています。

```go
...
Elements: []model.DialogElement{{
	DisplayName: "Title",
	Name:        "title",
	Type:        "text",
}, {
	DisplayName: "Message",
	Name:        "message",
	Type:        "textarea",
}},
...
```

以上より、Interactive Dialogからのリクエストを受信し、入力内容からMessage Attachmentのメッセージを作るアプリケーションは以下のようになります。

```go
...
	http.HandleFunc("/actions/dialog", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// (4) リクエストデータの読み出し
		b, _ := ioutil.ReadAll(r.Body)
		var payload model.SubmitDialogRequest
		json.Unmarshal(b, &payload)

		title, ok := payload.Submission["title"].(string)
		if !ok {
			resp := model.SubmitDialogResponse{Error: "failed to get title"}
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, string(resp.ToJson()))
			return
		}
		msg, ok := payload.Submission["message"].(string)
		if !ok {
			resp := model.SubmitDialogResponse{Error: "failed to get message"}
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, string(resp.ToJson()))
			return
		}

		// (5) Message Attachmentsインスタンス作成
		post := &model.Post{
			ChannelId: payload.ChannelId,
			Props: model.StringInterface{
				"attachments": []*model.SlackAttachment{{
					Title: title,
					Text:  msg,
				}},
			},
		}

		// (6) REST APIによるメッセージ投稿
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8065/api/v4/posts", strings.NewReader(post.ToJson()))
		req.Header.Add("Authorization", "Bearer "+MattermostAccessToken)
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)

		// (7) エラー処理
		dialogResp := model.SubmitDialogResponse{}
		if err != nil {
			dialogResp.Error = err.Error()
		}
		if resp.StatusCode != http.StatusCreated {
			dialogResp.Error = fmt.Sprintf("failed to request: %s", resp.Status)
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(dialogResp.ToJson()))
	})
...
```

Interactive Dialogからのリクエストを受け取ったら、`(4)`でリクエストを [`SubmitDialogRequest`](https://github.com/mattermost/mattermost-server/blob/dc1b42390b9bca393d03e2ccdbb16d66cd866431/model/integration_action.go#L230)形式で読み込みます。そして、`SubmitDialogRequest`の`Submission`から`title`、`message`をキーに持つ値を取得します。`Submission`のValueは`interface{}`型なので、文字列の場合はキャストが必要です。

データを読み出せたら `(5)` で、読み出したデータを使ってMessage Attachmentsを含む`Post`インスタンスを作成し、`(6)`でREST API経由で投稿を作成しています。REST APIを実行するため、Mattermostのアクセストークン(`MattermostAccessToken`)を事前に取得しておく必要があります。

最後に `(7)` でREST APIの実行結果をチェックし、エラーが発生している場合は[`SubmitDialogResponse`](https://github.com/mattermost/mattermost-server/blob/dc1b42390b9bca393d03e2ccdbb16d66cd866431/model/integration_action.go#L242)形式のデータを返却します。

```go
type SubmitDialogResponse struct {
	Error  string            `json:"error,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}
```

`SubmitDialogResponse`の`Error`にはInteractive Dialog全体のエラーとして表示される文字列、`Errors`には`DialogElement`の要素ごとのエラーメッセージを指定します。`Errors`は`Submission`と同じく`DialogElement`の`Name`をkeyとするmap形式でエラーメッセージを指定します。

試しに、以下のような`SubmitDialogResponse`を返したときの結果を紹介します。

```go
dialogResp.Errors = map[string]string{
	"title":   "title error",
	"message": "message error",
}
dialogResp.Error = "error"
w.Header().Add("Content-Type", "application/json")
io.WriteString(w, string(dialogResp.ToJson()))
```

![error](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day16/dialog-error.png)

以上のようにInteractive Dialogからのリクエストを処理できます。

## さいごに

本日は、Interactive Dialogの使い方について紹介しました。
明日からは、Mattermostのプラグイン機能について紹介していきます。
