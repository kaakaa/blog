---
title: "[Mattermost Integrations] Plugin (Server API)"
date: 2020-12-19T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "server", "hooks"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第19日目の記事です。

本記事では、Mattermost Server Pluginの開発に利用できるAPIについて紹介します。

Mattermost Pluginについての公式ドキュメントは下記になります。
https://developers.mattermost.com/extend/plugins/overview/

## Mattermost Plugin API

Mattermost Server Pluginの開発に利用できるAPIの一覧は下記にあります。

https://developers.mattermost.com/extend/plugins/server/reference/

ユーザーやチャンネル、投稿などの操作についてはREST APIとほぼ同様の機能を有しています。数が多いため全ては紹介しませんが、Server Plugin特有の処理に関するAPIを紹介していこうと思います。

### GetConfig

`GetConfig`は、Mattermost Serverのシステムコンソールの設定情報を取得します。

```go
siteURL := p.API.GetConfig().ServiceSettings.SiteURL
```

`p.API.GetConfig()`下記の`OnConfigurationChange`内で実行し、Plugin構造体のフィールドとして保持しておくのが良いそうです。
https://github.com/mattermost/mattermost-plugin-starter-template/blob/master/server/configuration.go


### OpenInteractiveDialog
`OpenInteractiveDialog`は、第16日目の記事でも紹介したInteractive DialogをPluginから開くためのAPIです。Interactive Dialogを開くには`TriggerId`というパラメータが必須であり、`TriggerId`はSlash Command実行時、もしくはInteractive Message Button/Menu実行時に送信されるリクエストからしか取得できません。

以下の例はSlash Command実行時にInteractive Dialogを開き、入力された情報を整形して投稿を作成するコードです。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day19/open-interactive-dialog.gif)

```go:plugin.go
func (p *SamplePlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	appErr := p.API.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: args.TriggerId,
		URL:       fmt.Sprintf("%s/plugins/com.github.kaakaa.mattermost-plugin-starter-template/callback", *p.API.GetConfig().ServiceSettings.SiteURL),
		Dialog: model.Dialog{
			Title:       "Sample Plugin Dialog",
			SubmitLabel: "Submit",
			Elements: []model.DialogElement{
				{
					DisplayName: "タイトル",
					Name:        "title",
					Placeholder: "Title",
					Type:        "text",
				},
				{
					DisplayName: "Snippet",
					Name:        "snippet",
					Type:        "textarea",
				},
			},
		},
	})
	if appErr != nil {
		return nil, appErr
	}
	return &model.CommandResponse{}, nil
}
```

Interactive Dialogのリクエスト送信先を

```go
		URL:       fmt.Sprintf("%s/plugins/com.github.kaakaa.mattermost-plugin-starter-template/callback", *p.API.GetConfig().ServiceSettings.SiteURL),
```

のようにすることで、Pluginから開いたInteractive DialogのリクエストをPluginで処理することもできます。このリクエストの受信先として、Mattermost Plugin Server Hookである`ServeHTTP`を使って`/callback`のエンドポイントを受け取る処理を追加します。

```go:plugin.go
func (p *SamplePlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/callback" {
		defer r.Body.Close()
		request := model.SubmitDialogRequestFromJson(r.Body)
		t := request.Submission["title"]
		s := request.Submission["snippet"]
		p.API.CreatePost(&model.Post{
			UserId:    request.UserId,
			ChannelId: request.ChannelId,
			Message:   fmt.Sprintf("## %s\n\n```\n%s\n```", t, s),
		})
		return
	}
	fmt.Fprint(w, "Hello, world!")
}
```

このようにすることで、Interactive Dialogの起動からリクエストの処理までをPluginだけで完結させることができます。

### PublishWebSocketEvent
`PublishWebSocketEvent`は、プラグイン独自のWebSoketイベントを送信するAPIです。

`PublishWebSocketEvent`は、3つの引数を取ります。

* `event string`: WebSocketイベント名を指定します。実際に送信されるWebSocketイベントには接頭辞として`custom_<PluginID>_`が付与されます
* `payload map[string]interface{}`: 送信されるデータの内容を指定します
* `broadcast *model.WebsocketBroadcast`: WebSocketを送信する対象を指定します
  * [`model.WebsocketBroadcast`](https://pkg.go.dev/github.com/mattermost/mattermost-server/v5/model#WebsocketBroadcast)で、ユーザー、チャンネル、チームを指定したり、受信対象から外すユーザーを指定したりできます

Webapp Plugin APIの[`registerWebSocketEventHandler`](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerWebSocketEventHandler)と組み合わせることで、プラグイン内でWebSocketイベントに関する処理を完結させることができます。
この辺りの使用例は第22日目以降の記事で紹介する予定です。

### SendMail

`SendMail`は、Mattermost PluginからHTMLメールを送信するAPIです。システムコンソールでSMTPの設定が完了している必要があります。

Slash Commandを実行した際にメールを送信する例は以下のようになります。

```go:plugin.go
func (p *SamplePlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	appErr := p.API.SendMail(
		"test@example.com",
		"Sample Title",
		"<h1>Mail from plugin</h1><div><p>Sample Mail</p></div>")
	if appErr != nil {
		return nil, appErr
	}

	return &model.CommandResponse{}, nil
}
```

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day19/send-mail.jpg)

### KVGet / KVSet

`KVGet`, `KVSet`は、プラグインごとに割り当てられたKey Valueストアから値を取得、または値を設定するAPIです。Key Valueストアには`[]byte`型のデータを格納できるため、格納用データの構造体を作成し、`json.Marshal`で`[]byte`型のデータ化したものを出し入れするのが主な使い方になるのではないかと思います。

Key Valueストアを使用してカウンターを実装した例が以下になります。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day19/kvget-kvset.gif)

```go:plugin.go
type counter struct {
	Count     int    `json:"count"`
	CreatedAt string `json:"created_at"`
}

func (p *SamplePlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// read data
	b, appErr := p.API.KVGet("counter")
	if appErr != nil {
		return nil, appErr
	}
	var ct counter
	if b == nil {
		// init if data is not found
		ct = counter{
			Count:     0,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
	} else {
		if err := json.Unmarshal(b, &ct); err != nil {
			return nil, model.NewAppError("where", "id", nil, err.Error(), http.StatusInternalServerError)
		}
	}
	// count up
	ct.Count = ct.Count + 1
	b, err := json.Marshal(ct)
	if err != nil {
		return nil, model.NewAppError("where", "id", nil, err.Error(), http.StatusInternalServerError)
	}
	// set data
	if appErr := p.API.KVSet("counter", b); appErr != nil {
		return nil, appErr
	}

	return &model.CommandResponse{Text: fmt.Sprintf("counter: %d", ct.Count)}, nil
}
```

Key Valueストアを操作するAPIは`KVGet`, `KVSet`以外にも数多くあります。データを削除する`KVDelete`や、プラグイン内で保存済みのデータの`key`を取得する`KVList`、期間を指定して値を格納できる`KVSetWithExpiry`などがあります。また、同時に`KVSet`が実行された場合に不整合が起きないようにするため、AtomicなKey Valueストアの操作を強制するための`KVCompareAndSet`などもあります。

APIの一覧については公式ドキュメントを参照してください。
https://developers.mattermost.com/extend/plugins/server/reference/


## Helpers
Mattermost Plugin APIは数多くあるため、Mattermostに対するほとんどの操作を実行することはできますが、より簡単にPlugin APIを実行するためのHelper関数がいくつか存在します。

ここでは、Helper関数のうち一部を紹介します。

Helper関数の一覧は下記の公式ドキュメントにあります。
https://developers.mattermost.com/extend/plugins/server/reference/#Helpers

### KVGetJSON / KVSetJSON
先ほどの`KVGet`、`KVSet`のコードでは、構造体を`[]byte`に変換するために`KVSet`を呼ぶ前に`json.Marshal`を、`KVGet`を呼んだ後に`json.Unmarshal`を呼んでいましたが、`KVGetJSON`、`KVSetJSON`を使用することで、その必要がなくなります。

そのため、先ほどの処理を多少すっきりと書くことができます。

```go:plugin.go
type counter struct {
	Count     int    `json:"count"`
	CreatedAt string `json:"created_at"`
}

func (p *SamplePlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// read data
	var ct counter
	exists, err := p.Helpers.KVGetJSON("counter", &ct)
	if err != nil {
		return nil, model.NewAppError("where", "id", nil, err.Error(), http.StatusInternalServerError)
	}
	if !exists {
		ct = counter{
			Count:     0,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
	}

	// count up
	ct.Count = ct.Count + 1

	// set data
	if appErr := p.Helpers.KVSetJSON("counter", ct); appErr != nil {
		return nil, model.NewAppError("where", "id", nil, err.Error(), http.StatusInternalServerError)
	}

	return &model.CommandResponse{Text: fmt.Sprintf("counter: %d", ct.Count)}, nil
}
```

### CheckRequiredServerConfiguration

`CheckRequiredServerConfiguration`は、システムコンソールの設定をチェックするためのHelper関数です。

引数に指定した[`model.Config`](https://pkg.go.dev/github.com/mattermost/mattermost-server/v5/model#Config)と同じ設定になっているかをチェックし、異なる設定だった場合は`false`を返します。

**システムコンソール > Botアカウント > Botアカウントの作成を有効にする**の設定が有効になっていない場合にプラグインの起動を停止するような例は以下のようになります。

```go:plugin.go
func toPtr(v bool) *bool {
	return &v
}

func (p *SamplePlugin) OnActivate() error {
	b, err := p.Helpers.CheckRequiredServerConfiguration(&model.Config{
		ServiceSettings: model.ServiceSettings{
			EnableBotAccountCreation: toPtr(true),
		},
	})
	if err != nil {
		return err
	}
	if !b {
		return errors.New("EnableBotAccountCreation must be true.")
	}

    ...
```

プラグイン起動時、サーバーのログには以下のようなログが出力されます。

```bash
2020-12-13T00:16:01.409+0900    error   mlog/log.go:229 Unable to restart plugin on upgrade.    {"path": "/api/v4/plugins", "request_id": "u31bzwuyhbyibratzznjxgxzrr", "ip_addr": "::1", "user_id": "87x93uo8pfnzdro9ktcmobpa1r", "method": "POST", "err_where": "installExtractedPlugin", "http_code": 500, "err_details": "EnableBotAccountCreation must be true."}
```

## さいごに

本日は、Mattermost PluginのServer APIについて紹介しました。
明日からは、Mattermost Pluginの**Webapp**サイドの実装について紹介します。
