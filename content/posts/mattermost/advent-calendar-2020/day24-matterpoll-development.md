---
title: "[Mattermost Integrations] Matterpoll Plugin (Development)"
date: 2020-12-24T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "matterpoll", "testing", "ci"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第24日目の記事です。

以前から開発に参加している[Matterpoll](https://github.com/matterpoll/matterpoll)がMattermost公式のPlugin MarketplaceでCommunity Pluginとして公開されました。

昨日までの記事でMatterpoll PluginのServer側・Webapp側の実装について紹介しましたが、本日の記事ではMatterpoll PluginのテストやCIなどの開発周りの話を紹介します。

## Mattermost Pluginのテスト

### Server側のテスト

Mattermost PluginのServer側の機能をMattermostを実際に起動することなくテストするには、[`plugintest`]([https://github.com/mattermost/mattermost-server/tree/master/plugin/plugintest](https://github.com/mattermost/mattermost-server/tree/master/plugin/plugintest))パッケージを使用します。

[`plugintest`]https://github.com/mattermost/mattermost-server/tree/master/plugin/plugintest](https://github.com/mattermost/mattermost-server/tree/master/plugin/plugintest)パッケージを使用してテストを記述していく方法について、Mattermost本体にある[プラグインテストのサンプルコード](https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go)を元に紹介します。Matterpollプラグインも大部分はこれと同じ方式でテストが記述されています。

```go:plugin/plugintest/example_hello_user_test.go
...
type HelloUserPlugin struct {
	plugin.MattermostPlugin
}

func (p *HelloUserPlugin) ServeHTTP(context *plugin.Context, w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")
	user, err := p.API.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.API.LogError(err.Error())
		return
	}

	fmt.Fprintf(w, "Welcome back, %s!", user.Username)
}
...
```

[https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go#L21](https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go#L21)

Mattermostプラグインの本体は、`plugin.MattermostPlugin`がembeddedされた構造体です。ここでは、`HelloUserPlugin` 構造体がそれにあたります。この構造体を通じてプラグイン用のAPIを実行したり、Hookとなるメソッドを実装したりすることで、サーバー側の動作を定義することができます。`HelloUserPlugin` では、プラグイン独自のエンドポイントを追加する [`ServeHTTP`]([https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP))  Hooksが実装され、その処理の中で `GetUser` API を実行してエンドポイントへアクセスしたユーザーの情報を取得し、そのユーザー名をレスポンスとして返しています。

この`ServeHTTP` Hooksをテストする場合、単に`HelloUserPlugin`構造体を生成して`ServeHTTP`メソッドを実行しただけだと、プラグインAPI (`GetUser`) の実行部分で実際のMattermostサーバーの機能を呼び出そうとしてしまいエラーとなってしまいます。

```go
	user, err := p.API.GetUser(userID)
```

`HelloUserPlugin` の `API` フィールドは、`plugin.MattermostPlugin` 構造体が持っていたフィールドであり、このフィールドを通して呼び出されるメソッドはMattermostサーバーの処理に依存しているからです。Mattermostサーバーがない状態でテストを実行する場合、この`API`フィールドを入れ替える必要があります。

ここで、プラグイン用のテストメソッドとして定義されている [`Example`](https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go#L37) メソッドをみてみましょう。

```go
...
func Example() {
	t := &testing.T{}
	user :&model.User{                           ... (1)
		Id:       model.NewId(),
		Username: "billybob",
	}

	api := &plugintest.API{}                       ... (2)
	api.On("GetUser", user.Id).Return(user, nil)   ... (3)
	defer api.AssertExpectations(t)                ... (4)

	helpers := &plugintest.Helpers{}
	defer helpers.AssertExpectations(t)

	p := &HelloUserPlugin{}
	p.SetAPI(api)                                  ... (5)
	p.SetHelpers(helpers)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Add("Mattermost-User-Id", user.Id)
	p.ServeHTTP(&plugin.Context{}, w, r)           ... (6)
	body, err := ioutil.ReadAll(w.Result().Body)
	require.NoError(t, err)
	assert.Equal(t, "Welcome back, billybob!", string(body))
}
```
[https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go#L37](https://github.com/mattermost/mattermost-server/blob/5122b9e2929dbf84e22f496ee97d007fa18f2d2e/plugin/plugintest/example_hello_user_test.go#L37)


**(2)**で `plugintest.API` という構造体のインスタンスを生成し、生成されたインスタンスをテスト対象のプラグイン構造体 `HelloUserPlugin` へ `SetAPI` 関数を通じてセット **(5)** しています。 `plugintest.API`自体は、そのままでは何も処理を行わないため、`api.On` で特定の引数が与えられた時の処理をテスト側で実装しています **(3)**。

```go
    ...
	user :&model.User{                           ... (1)
		Id:       model.NewId(),
		Username: "billybob",
	}
	...
    api.On("GetUser", user.Id).Return(user, nil)   ... (3)
	defer api.AssertExpectations(t)                ... (4)
	...
```

上記のコードの場合、`user.Id` を引数として `GetUser` メソッドが呼び出された場合に **(1)** で定義された `user` を返却するようモックを定義しています。また、`defer api.AssertExpectations(t)` **(4)** を書いておくことで、テスト実行が終了した時に `api.On` で定義したモックが実行されていなかった場合にテストを失敗させることができます。

`HelloUserPlugin` では [`Helpers`]([https://developers.mattermost.com/extend/plugins/server/reference/#Helpers](https://developers.mattermost.com/extend/plugins/server/reference/#Helpers))関数を利用していませんでしたが、`Helpers`のメソッドを利用している場合も `API` と同様に関数をモックすることができます。

最後に `httptest` パッケージを使って `ServeHTTP` Hook を呼び出し、レスポンスが想定通りであることをチェックしています。

```go
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Add("Mattermost-User-Id", user.Id)
	p.ServeHTTP(&plugin.Context{}, w, r)           ... (6)
```

このように `plugintest` パッケージを使うことで 、Mattermostサーバーがない状況でもMattermostプラグインの処理をテストできるようになります。あとは、とにかくパターンを網羅するケースを書き出すだけです。筋力です。

### Server側のテスト(Store)

Matterpollでは Mattermost PluginのKey Valueストアへのアクセスを抽象化した `store` パッケージを用意しています。テスト対象のメソッド内で Key Value ストアへのアクセスが必要な処理があった場合、プラグインAPIの `KVSet`、`KVGet` などのメソッドを `plugintest` パッケージでモックすることも可能ですが、テスト記述が煩雑になるため、Matterpollでは [vektra/mockery]([https://github.com/vektra/mockery](https://github.com/vektra/mockery)) を使って `store.Store` インタフェースからモックを生成してテストを記述しています。この  [vektra/mockery]([https://github.com/vektra/mockery](https://github.com/vektra/mockery)) は、Mattermost本体の `plugintest` パッケージを生成するのにも使われているものです。

### Webapp側のテスト

ここについては、Mattermost Plugin特有のトピックというものはなく、開発したReact Componentに対して[`Jest`]([https://jestjs.io/](https://jestjs.io/ja/)) を使ったsnapshotテストを実装しています。

### CI

CIはCircleCIを使っており、Mattermostプラグイン用のCircleCI Orbを使って lint / build/ coverage / deploy を行っています。

[https://github.com/matterpoll/matterpoll/blob/master/.circleci/config.yml](https://github.com/matterpoll/matterpoll/blob/master/.circleci/config.yml)

カバレッジは [Codecov](https://codecov.io/gh/matterpoll/matterpoll) によるPRへのコメントフィードバックを実施しています。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day24/codecov/png)

## Checks

静的解析系のツールはMattermostプラグインのテンプレートリポジトリ [mattermost/mattermost\-plugin\-starter\-template: Build scripts and templates for writing Mattermost plugins\.](https://github.com/mattermost/mattermost-plugin-starter-template) で定義されているものとベースは同じで、採用しているlinterやルールが少し異なるという感じです。

サーバー側は [golangci-lint](https://github.com/golangci/golangci-lint) を使っており設定ファイルは下記の通りです。

```yml:.golangci.yml
run:
  timeout: 5m
  modules-download-mode: readonly

linters-settings:
  goconst:
    min-len: 2
    min-occurrences: 2
  gofmt:
    simplify: true
  goimports:
    local-prefixes: github.com/matterpoll/matterpoll
  golint:
    min-confidence: 0.0
  govet:
    check-shadowing: true
    enable-all: true
  misspell:
    locale: US
  maligned:
    suggest-new: true
  
linters:
  disable-all: true
  enable:
    - bodyclose
    - deadcode
    - dogsled
    - errcheck
    - goconst
    - gocritic
    - gofmt
    - goimports
    - golint
    - gosec
    - gosimple
    - govet
    - ineffassign
    - interfacer
    - maligned
    - misspell
    - nakedret
    - scopelint
    - staticcheck
    - structcheck
    - stylecheck
    - typecheck
    - unconvert
    - unparam
    - unused
    - varcheck
    - whitespace

issues:
  exclude-rules:
    # Exclude some linters from running on tests files.
    - path: _test\.go
      linters:
        - dupl
        - goconst
        - scopelint # https://github.com/kyoh86/scopelint/issues/4
```

[https://github.com/matterpoll/matterpoll/blob/d4ffdbfd6dcdea359b7419e0baa3ab8aaa32e420/.golangci.yml](https://github.com/matterpoll/matterpoll/blob/d4ffdbfd6dcdea359b7419e0baa3ab8aaa32e420/.golangci.yml)

クライアント側は [ESLint](https://eslint.org/) を使っており、設定ファイルは下記の通りです。（長いためリンク先参照）

[https://github.com/matterpoll/matterpoll/blob/0b797025cfbf319c43464fcf457d4cdfe5086188/webapp/.eslintrc.json](https://github.com/matterpoll/matterpoll/blob/0b797025cfbf319c43464fcf457d4cdfe5086188/webapp/.eslintrc.json)

## 翻訳

Matterpollのメッセージは翻訳可能な形式で管理されており、現在、下記の言語が利用可能です。

- English
- France
- German
- Japanese
- Korean
- Polish
- Simplified Chinese
- Spanish
- Traditional Chinese

言語の切り替えは、Mattermost本体の設定に応じて実施されます。

Mattermost Pluginにおける翻訳処理の詳細については、Matterpollの共同開発者であるHanzei による下記の記事で紹介されています。

[https://developers.mattermost.com/blog/localizing-matterpoll/](https://developers.mattermost.com/blog/localizing-matterpoll/)

### Server側の翻訳

Server側の翻訳機能は[go-i18n](https://github.com/nicksnyder/go-i18n)を使用しています。

翻訳対象のメッセージは、コード上では以下のように`go-i18n`の`i18n`パッケージの構造体として書かれています。

```go:server/plugin/command.go
...
HelpText: p.LocalizeWithConfig(l, &i18n.LocalizeConfig{
	DefaultMessage: &i18n.Message{
		ID:    "dialog.createPoll.setting.multi",
		Other: "The number of options that an user can vote on.",
	}}),
...
```
[https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L265](
https://github.com/matterpoll/matterpoll/blob/45f095875a98fb1d4f3f166851c86f41b987493e/server/plugin/command.go#L265)

実際の翻訳を行う場合は、`go-i18n`のコマンドラインツールを使って、上記のような`i18n`の構造体として宣言されたメッセージをjsonファイルに集約します。その辺りの手順については以下のドキュメントにまとめられています。

https://github.com/matterpoll/matterpoll/blob/master/CONTRIBUTING.md#translating-strings

`go-i18n`によって集約されたメッセージを各国のコントリビュータにローカライズしてもらうことで、翻訳されたメッセージが表示されるようになっています。
https://github.com/matterpoll/matterpoll/tree/45f095875a98fb1d4f3f166851c86f41b987493e/assets/i18n

### Webapp側の翻訳
Matterpoll Pluginでは、まだWebapp側の翻訳機能は実装されていませんが、Mattermost Pluginの機能としては実装できるようになっています。

Webapp側の翻訳は、[`react-intl`](https://www.npmjs.com/package/react-intl)を使用しています。

```jsx
...
import {FormattedMessage} from 'react-intl';
...
        <FormattedMessage
            id='rootModal.message'
            defaultMessage='Root Modal2'
		/>
...
```

上記のようにコード内で使用されている翻訳対象のメッセージは、`make i18n-extract`で集約することができます。

```Makefile
...
## Extract strings for translation from the source code.
.PHONY: i18n-extract
i18n-extract:
ifneq ($(HAS_WEBAPP),)
ifeq ($(HAS_MM_UTILITIES),)
	@echo "You must clone github.com/mattermost/mattermost-utilities repo in .. to use this command"
else
	cd $(MM_UTILITIES_DIR) && npm install && npm run babel && node mmjstool/build/index.js i18n extract-webapp --webapp-dir $(PWD)/webapp
endif
endif
...
```
[https://github.com/matterpoll/matterpoll/blob/master/Makefile#L205](https://github.com/matterpoll/matterpoll/blob/master/Makefile#L205)

集約されたメッセージファイルをMattermost Webapp Plugin APIの[`registerTranslations`](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerTranslations)で登録することで、Webapp側のメッセージの翻訳ができるようになります。

## さいごに

本日は、Matterpoll PluginのテストやCIなどの開発に関する事柄について紹介しました。
Mattermost Integractionsに関する紹介は本日の記事で終了です。

明日は、本記事を執筆する中で見つかった問題に対するIssue/PRについて紹介します。