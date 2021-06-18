---
title: "[Mattermost Integrations] Matterpoll Plugin (Server)"
date: 2020-12-22T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "matterpoll"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第22日目の記事です。

以前から開発に参加している[Matterpoll](https://github.com/matterpoll/matterpoll)がMattermost公式のPlugin MarketplaceでCommunity Pluginとして公開されました。

公式のPlugin Marketplaceで公開されたことにより、プラグインをアップロードする必要なく**メインメニュー > プラグインマーケットプレース**からボタン一つでインストールできるようになりました。

![main menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/marketplace-menu.png)

![install button](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/marketplace-install-button.png)

MatterpollはMattermost上で投票を行うことができるPluginであり、Mattermost社の開発者である Hanzei と私を中心に開発が進められています。

![screenshot of MatterPoll](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll.png)

以前よりMattermostに投票機能を付けるという提案は[MattermostのFeature Request Forumでも票数を集めて](https://mattermost.uservoice.com/forums/306457-general/suggestions/11721996-polls-like-doodle-or-for-decision-making)いたため、以前は絵文字で投票を促すような統合機能を開発していました。

[https://github.com/matterpoll/matterpoll-emoji](https://github.com/matterpoll/matterpoll-emoji)

ただその当時はMattermostにPlugin機能がなかったため、この統合機能を利用するにはMattermostとは別にサーバーアプリを起動しておく必要があるのが難点でした。

その後、MattermostにPlugin機能が付いたことで、Mattermostのプロセスの上で統合機能を動作させることができるようになり、`matterpoll-emoji` もPluginとして開発し直すこととなりました。そこで出来たプラグインが、今回Plugin Marketplaceで公開されたMatterpoll Pluginになります。

この記事では、まだ日本語ではあまり情報のないMattermost Pluginについて、Matterpoll Pluginの構造を元に紹介していきたいと思います。

本記事では、MattermostプラグインのGetting Started的な部分についてはあまり触れません。Mattermost Pluginの基本的な部分については、[アドベントカレンダーの他の記事](https://qiita.com/advent-calendar/2020/mattermost-integrations)や、公式ドキュメントを参照いただければと思います。

[https://developers.mattermost.com/extend/plugins/](https://developers.mattermost.com/extend/plugins/)

また、実際に開発を始める場合はGitHubで公開されているMattermost Plugin用のテンプレートリポジトリ [https://github.com/mattermost/mattermost-plugin-starter-template](https://github.com/mattermost/mattermost-plugin-starter-template) を使用するのがおすすめです。テンプレートリポジトリを使った開発の始め方については、下記の記事で紹介しています。（少し情報が古いかもしれません...）

[Mattermostプラグイン用のリポジトリテンプレート · kaakaa blog](https://blog.kaakaa.dev/post/mattermost/plugin-template/)

## 概要

Matterpoll PluginはServer側の機能とWebapp側の機能が連携して動作しています。

Server側の機能は投票作成コマンド(`/poll`)を処理したり、投票が実行のHTTPリクエストを受け取って処理をしたり、投票に関するデータを管理したりしています。Webapp側ではユーザー毎に表示を変えたい場合など、クライアント側の処理を実装しています。Server側はGo言語、WebApp側はReact.js(JavaScript/TypeScript)で書かれています。

Mattermostの基本的な処理の流れは下記のようになります。

![architecture](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll.dio.png)

ユーザーがMattermost上でスラッシュコマンド `/poll` を実行する **(1)** と、そのコマンドをMatterpoll Server側が受け取り **(2)**、投票データ(`Poll`)を作成します **(3)**。作成された `Poll` のデータは、プラグイン用のKey-Value Storeに格納されます **(4)**。そして、Mattermostの投稿形式 (`Post`)へ変換され **(5)**、通常の投稿と同様にMattermostのデータベースに格納され **(6)**、WebappへWebSocketを通じてPublishされます **(7)**。(実際には、諸事情により投稿が作成 **(6)** されてから投票データをKey-Valueストア に格納 **(4)** していますが、話を簡単にするため上記の順で話をしています)

Matterpoll Pluginにより作成されたメッセージは、Mattermost Webapp側の処理によってユーザー毎に見た目が変わります **(8)**。具体的には、自分が投票した回答のボタンの色が変わったり、投票管理系のボタンが権限のないユーザーから見えなくなったりします。このWebApp側の機能は、現在実験的な機能として提供されており、デフォルトではOffになっています。MattermostのSystem ConsoleのMatterpollプラグインの設定からOnにすることが出来ます。(設定がOffでも、投票機能自体は利用可能)

## Server側の処理

MatterpollプラグインはServer側で投票の作成、実行、データの管理などの投票に関わる基本的な処理を実装しています。ここでは、Mattermost Pluginの機能により、以下の処理をどのように実装しているかについて紹介していきます。

- プラグイン専用のスラッシュコマンドを追加し、投票作成コマンドを実行できるようにする
- プラグイン専用のエンドポイントを追加し、ユーザーからの投票を受け付ける
- プラグイン専用のKey-Valueストアで投票データを管理する

#### 投票作成用の独自スラッシュコマンドの追加

![matterpoll-command.dio.png](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll-command.dio.png)

Matterpollは専用のスラッシュコマンド `/poll` を実行することで投票を作成します。この `/poll` コマンドは、プラグイン起動時に登録されます。

Mattermostプラグインでは、プラグインAPIの [RegisterCommand]([https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand](https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand)) を実行することで新たなスラッシュコマンドを登録することができます。

実際のコードでは下記のようになります。

```go:server/plugin/configuration.go
...
// OnConfigurationChange loads the plugin configuration, validates it and saves it.
func (p *MatterpollPlugin) OnConfigurationChange() error {
  ...
	// This require a loaded i18n bundle
	if p.isActivated() {
		command, err := p.getCommand(configuration.Trigger)
		if err != nil {
			return errors.Wrap(err, "failed to get command")
		}
    ...
		if err := p.API.RegisterCommand(command); err != nil {
			return errors.Wrap(err, "failed to register new command")
		}
    ...
	}
  ...
}
...
```

[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/configuration.go#L43](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/configuration.go#L43)

`RegisterCommand`の引数には、`getCommand`メソッドの戻り値 `command` を指定していますが、`getCommand`はSlash Commandのモデルである`model.Command`を生成するための内部メソッドです。

```go:server/plugin/command.go
func (p *MatterpollPlugin) getCommand(trigger string) (*model.Command, error) {
  ...
	return &model.Command{
		Trigger:              trigger,
		AutoComplete:         true,
		AutoCompleteDesc:     p.LocalizeDefaultMessage(localizer, commandAutoCompleteDesc),
		AutoCompleteHint:     p.LocalizeDefaultMessage(localizer, commandAutoCompleteHint),
		AutocompleteIconData: iconData,
	}, nil
}
```
[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L215](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L215)

また、`getCommand`の引数 (`configuration.Trigger`)はSlash Commandのトリガーワード(スラッシュコマンド名)ですが、Matterpollではスラッシュコマンド名を設定から変更できるようにしてあるため、設定画面で入力された値を使用するために`configuration.Trigger`を指定しています。

![設定画面](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll-settings-trigger.png)

そのため、設定が変更された際に実行される [OnConfigurationChange]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnConfigurationChange](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnConfigurationChange)) Hook の中で、[RegisterCommand]([https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand](https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand)) を実行しています。


---

プラグインにより登録されたスラッシュコマンド(`/poll`)が実行された場合の処理は、[ExecuteCommand]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand)) Hookとして実装します。

```go:server/plugin/command.go
// ExecuteCommand parses a given input and creates a poll if the input is correct
func (p *MatterpollPlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	msg, appErr := p.executeCommand(args)
	if msg != "" {
		p.SendEphemeralPost(args.ChannelId, args.UserId, args.RootId, msg)
	}
	return &model.CommandResponse{}, appErr
}
```

[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L84](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L84)

[ExecuteCommand]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand))の第二引数である [*model.CommandArgs]([https://pkg.go.dev/github.com/mattermost/mattermost-server/v5/model#CommandArgs](https://pkg.go.dev/github.com/mattermost/mattermost-server/v5/model#CommandArgs)) 型の変数にコマンド実行時の引数や、コマンド実行したユーザーのユーザーID、コマンドが実行されたチャンネルのIDなどの情報が格納されているため、これらの値を使ってスラッシュコマンドが実行された時の処理を実装していきます。

`executeCommand` メソッドで引数を解析して投票を作成する処理が実行されますが、細かい処理が多いためこれ以上の説明は割愛します。

#### 投票リクエストを受け付けるための独自エンドポイントの追加

Matterpollにより作成された投稿では、メッセージに埋め込まれたボタンをクリックすることで投票を行うことができます。

![Message Button](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll-sample.png)

このボタンはMattermostの [Interacitve Message]([https://docs.mattermost.com/developer/interactive-messages.html](https://docs.mattermost.com/developer/interactive-messages.html)) 機能を利用して実装されています (Interactive Messageについては[第14日目](https://blog.kaakaa.dev/post/mattermost/advent-calendar-2020/day14-interactive-message-button/)、[第15日目](https://blog.kaakaa.dev/post/mattermost/advent-calendar-2020/day15-interactive-message-menu/)の記事で紹介しています)。Matterpollでは [https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L186](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L186) のあたりで実装されています。）

Matterpollが作成する投稿に表示されているボタンは、クリックするとそれぞれ異なるURLへHTTPリクエストが送信されます。このHTTPリクエストをMatterpollのServer側が受け取り、どのユーザーがどの回答に投票したかを解析し、データベースに保存されている投票のデータを更新することで投票処理を実行しています。

![matterpoll-command.dio.png](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll-vote.dio.png)

リクエストを受け取るためにはMatterpollプラグイン用のエンドポイントをMattermost上に定義する必要がありますが、これはMattermostプラグインの機能である [ServeHTTP]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP)) Hookにより実現しています。[ServeHTTP]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP))を実装すると、Mattermostに `/plugins/{pluginid}` というエンドポイントが生成され、このエンドポイントに送信されたリクエストをプラグインで処理することができるようになります。例えば、Mattermostに `https://example.com:8065` というURLでアクセスできるとすると、`https://example.com:8065/plugins/com.github.matterpoll.matterpoll`がMatterpollプラグイン専用のエンドポイントになります。

Matterpollでは、さらに[gorilla/mux]([https://github.com/gorilla/mux](https://github.com/gorilla/mux)) を使用して Router を生成し、[ServeHTTP]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP)) の引数をそのままRouterへ流し込むことで、投票の作成/削除/終了や投票の管理など様々なリクエストに対応するエンドポイントを生成して処理を実装しています。

```go:server/plugin/api.go
...
// InitAPI initializes the REST API
func (p *MatterpollPlugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", p.handleInfo).Methods(http.MethodGet)
	r.HandleFunc("/"+iconFilename, p.handleLogo).Methods(http.MethodGet)

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(checkAuthenticity)
	apiV1.HandleFunc("/configuration", p.handlePluginConfiguration).Methods(http.MethodGet)

	apiV1.HandleFunc("/polls/create", p.handleSubmitDialogRequest(p.handleCreatePoll)).Methods(http.MethodPost)
	pollRouter := apiV1.PathPrefix("/polls/{id:[a-z0-9]+}").Subrouter()
	pollRouter.HandleFunc("/vote/{optionNumber:[0-9]+}", p.handlePostActionIntegrationRequest(p.handleVote)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/votes/reset", p.handlePostActionIntegrationRequest(p.handleResetVotes)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/option/add/request", p.handlePostActionIntegrationRequest(p.handleAddOption)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/option/add", p.handleSubmitDialogRequest(p.handleAddOptionConfirm)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/end", p.handlePostActionIntegrationRequest(p.handleEndPoll)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/end/confirm", p.handleSubmitDialogRequest(p.handleEndPollConfirm)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/delete", p.handlePostActionIntegrationRequest(p.handleDeletePoll)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/delete/confirm", p.handleSubmitDialogRequest(p.handleDeletePollConfirm)).Methods(http.MethodPost)
	pollRouter.HandleFunc("/metadata", p.handlePollMetadata).Methods(http.MethodGet)
	return r
}

func (p *MatterpollPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("New request:", "Host", r.Host, "RequestURI", r.RequestURI, "Method", r.Method)
	p.router.ServeHTTP(w, r)
}
...
```

[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/api.go#L70](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/api.go#L70)

#### Key-Valueストア による投票データの管理


ここまで、スラッシュコマンドによる投票の作成、Interactive Messageによる投票処理について説明してきましたが、投票機能を実現するには、これらの投票に関するユーザーの操作を記録しておく必要があります。Mattermostプラグインでは各プラグイン毎に専用のKey Valueストアが用意されており、Matterpollでは投票に関するデータの保存にこのKey-Valueストアを利用しています。

![matterpoll-kv.dio.png](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day22/matterpoll-kv.dio.png)

投票データのKeyは `poll_` を接頭辞にもつ投票IDであり、投票IDは投票が作成されるごとに Mattermost本体の[`model.NewID`]([https://github.com/mattermost/mattermost-server/blob/00ed2b138b0099be64a5f6b829ad916c220e1d6a/model/utils.go#L160](https://github.com/mattermost/mattermost-server/blob/00ed2b138b0099be64a5f6b829ad916c220e1d6a/model/utils.go#L160))メソッドを使って乱数を生成しています。Mattermost PluginのKey Valueストアには`[]byte`型のデータのみ格納できるため、投票状態を持つ[Poll]([https://github.com/matterpoll/matterpoll/blob/22dc39da19825ef10c59e35d28359c02c220debb/server/poll/poll.go#L23](https://github.com/matterpoll/matterpoll/blob/22dc39da19825ef10c59e35d28359c02c220debb/server/poll/poll.go#L23))構造体をJSON形式の `[]byte` に変換した値を格納しています。

```go:server/plugin/poll.go
...
// Poll stores all needed information for a poll
type Poll struct {
	ID            string
	PostID        string `json:"post_id,omitempty"`
	CreatedAt     int64
	Creator       string
	Question      string
	AnswerOptions []*AnswerOption
	Settings      Settings
}
...
// EncodeToByte returns a poll as a byte array
func (p *Poll) EncodeToByte() []byte {
	b, _ := json.Marshal(p)
	return b
}
...
```

[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/poll/poll.go#L23](https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/poll/poll.go#L23)

Key Valueストアは、単純に[KVSet]([https://developers.mattermost.com/extend/plugins/server/reference/#API.KVSet](https://developers.mattermost.com/extend/plugins/server/reference/#API.KVSet)) で値の格納、[KVGet]([https://developers.mattermost.com/extend/plugins/server/reference/#API.KVGet](https://developers.mattermost.com/extend/plugins/server/reference/#API.KVGet))で値の取得を行うことができますが、Matterpollでは同時に投票操作が行われた場合などに不整合が起きないようにするため、データを格納する際はAtomicオプションを付けて実行しています。

```go
...
// Update updates an existing a poll in the KV Store.
func (s *PollStore) Update(prev *poll.Poll, new *poll.Poll) error {
	opt := model.PluginKVSetOptions{
		Atomic:   true,
		OldValue: prev.EncodeToByte(),
	}
	ok, err := s.api.KVSetWithOptions(pollPrefix+prev.ID, new.EncodeToByte(), opt)
	if err != nil {
		return err
	}
  ...
}
...
```

また、Matterpollでは、[Key Valueストア への処理は全てinterfaceを用意しています]([https://github.com/matterpoll/matterpoll/blob/22dc39da19825ef10c59e35d28359c02c220debb/server/store/store.go](https://github.com/matterpoll/matterpoll/blob/22dc39da19825ef10c59e35d28359c02c220debb/server/store/store.go))が、これはテスト用に [mockery]([https://github.com/vektra/mockery](https://github.com/vektra/mockery)) でモックオブジェクトを自動で作成するためです。


## さいごに

本日は、Mattermost上で投票を行えるようにするMatterpoll Pluginの概要とServer側の実装について紹介しました。
明日は、Matterpoll PluginのWebapp側の実装について紹介します。
