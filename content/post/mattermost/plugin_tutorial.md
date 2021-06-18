---
title: "Mattermostプラグインの作り方"
date: 2019-01-01T15:05:21+09:00
draft: false
toc: true
tags: ["mattermost", "plugin"]
---

**(追記 2020/04/12)** 

本記事でも使用している`dep`がもうメンテナンスされておらず、本記事で紹介している内容でMattermostの最新版を利用したプラグインの開発が困難になっているようです。

今後Mattermostプラグイン開発を始める場合、[Mattermostプラグイン用のリポジトリテンプレート](https://blog.kaakaa.dev/post/mattermost/plugin-template/) でも紹介していますが、Mattermost公式チームがメンテナンスしているテンプレートリポジトリ https://github.com/mattermost/mattermost-plugin-starter-template を利用することを推奨します。

**(追記おわり)**

# はじめに

2017/11/16にリリースされたMattermost v4.4.0で、Mattermost自体の機能拡張を行うためのプラグイン機能が追加されました。
その後、2018/08/16にリリースされたMattermost v5.2.0で、破壊的変更となるプラグインアーキテクチャの一新が行われ、現在も新たな拡張ポイントやAPIが追加され続けています。

Mattermostプラグインでは、サーバーサイド(Go言語)とフロントエンド(React.js)両方の拡張が行えます。サーバーサイドを拡張したい場合はGo言語、フロントエンドを拡張したい場合はJavascripnt(React.js)を書くことになります。
Mattermostプラグインを実装することで、新たなユーザーがチーム/チャンネルに参加した際の処理を追加したり、投稿を表すReactコンポーネントを自作のものに置き換えるなど、今まで手を入れることができなかった部分にまで機能追加を行うことができるようになります。

今回は、簡単な実装例を交えながらMattermostプラグインの作成方法について紹介していきたいと思います。

この記事で紹介しているコードは、下記のリポジトリで公開しています。
https://github.com/kaakaa/mattermost-plugin-first
(Mattermost v5.6を対象に動作確認を行なっています）

# Step 1: プラグイン開発・アップロード

まず、何も機能を持たない空のプラグインを作成し、そのプラグインをアップロードするまでのフローを見ていきます。

## マニフェストファイル

Mattermostプラグインの実体は`.tar.gz`形式のファイルとなります。
`.tar.gz`ファイルの中身として最低限必要なファイルは**マニフェストファイル**だけです。

JSON形式(`plugin.json`)かYAML形式(`plugin.yaml`)で、プラグインのID、名前、バージョンや実行バイナリのパスなどのメタ情報を記述するファイルです。

マニフェストファイル記述できる内容については、下記のリファレンスを参照してください。
https://developers.mattermost.com/extend/plugins/manifest-reference/

まず、下記のようにマニフェストファイルを作成します。

```json:plugin.json
{
    "id": "org.kaakaa.mattermost-plugin-first",
    "name": "Mattermost Plugin Sample",
    "description": "このプラグインはQiita記事用のMattermostプラグインです。",
    "version": "0.0.1"
}
```

## プラグイン作成

上記のマニフェストファイルを含む`.tar.gz`ファイルを作成するため、下記のようなMakefileを作成します。（makeの環境がない場合は、手動でコマンドを実行しても構いません）

```Makefile
PLUGIN_ID ?= org.kaakaa.mattermost-plugin-first
PLUGIN_VERSION ?= 0.0.1
BUNDLE_NAME ?= $(PLUGIN_ID)_$(PLUGIN_VERSION).tar.gz

build:
	rm -rf dist/
	mkdir -p dist/$(PLUGIN_ID)
	cp plugin.json dist/$(PLUGIN_ID)/

	cd dist && tar -cvzf $(BUNDLE_NAME) $(PLUGIN_ID)

	@echo plugin built at: dist/$(BUNDLE_NAME)
```

`make build`を実行することで`dist/org.kaakaa.mattermost-plugin-first_0.0.1.tar.gz`というファイルが作成されます。

## プラグインアップロード用の設定

作成したプラグインをアップロードしていきます。
まず、プラグインをシステムコンソールからアップロードできるようにするために、Mattermostサーバーの設定ファイル(`config.json`)を編集します。

```json

...
    },
    "PluginSettings": {
        "Enable": true,
        "EnableUploads": true,
        "Directory": "./plugins",
...
```

Mattermost設定ファイル内の`PluginSettings`にある`Enable`と`EnableUploads`を`true`にし、Mattermostサーバーを再起動します。

## プラグインのアップロード
Mattermostの **システムコンソール** > **プラグイン(ベータ版)** > **管理**画面からプラグインをアップロードします。

![スクリーンショット 2019-01-01 10.36.28.png](https://qiita-image-store.s3.amazonaws.com/0/9891/9d56287e-75d6-02a1-5ee8-529c54a95823.png)

**ファイルを選択する** ボタンを押し、先ほど作成した`dist/org.kaakaa.mattermost-plugin-first_0.0.1.tar.gz`を選択し、**アップロード**ボタンを押すことでプラグインがアップロードされます。
プラグインが正常にアップロードされると**インストール済みプラグイン**に表示されます。

![スクリーンショット 2019-01-01 10.42.33.png](https://qiita-image-store.s3.amazonaws.com/0/9891/2de7b8f6-3280-0468-0236-7c6726de1c60.png)

アップロードされたプラグインの**有効にする**ボタンを押すことでプラグインを有効にすることができますが、現在はプラグインの実装が行われていないため、エラーメッセージ「このプラグインを起動できませんでした。エラーに関するシステムログを確認してください。」が表示されるはずです。

![スクリーンショット 2019-01-01 10.43.14.png](https://qiita-image-store.s3.amazonaws.com/0/9891/35f8c473-079d-bdc6-fa72-87226bc1a216.png)


# Step 2: サーバーサイドプラグインの実装

Step 1で作成したプラグインにサーバーサイドの実装を加えていきます。
`server`というディレクトリを作成し、そこにGoコードを書いていきます。

```
root/
  ├ dist/
  │  └ org.kaakaa.mattermost-plugin-first_0.0.2.tar.gz
  ├ server/
  │  ├ main.go
  │  ├ plugin.go
  │  └ Gopkg.toml
  └ plugin.json
```

依存ライブラリの管理には[golang/dep](https://github.com/golang/dep)を使用しています。

## main.go

まずGo言語のmainメソッドでMattermostプラグインを認識させるために`plugin.ClientMain`メソッドを呼び出し、自作のPlugin用の構造体(`FirstPlugin`)を読み込ませます。

```go:server/main.go
package main

import "github.com/mattermost/mattermost-server/plugin"

func main() {
	plugin.ClientMain(&FirstPlugin{})
}
```

この自作のFirstPlugin構造体にプラグインの実装を追加していきます。

## plugin.go

`FirstPlugin`構造体に最低限求められるのは、`plugin.MattermostPlugin`を継承することです。
`plugin.MattermostPlugin`を継承することでMattermostプラグインとして認識され、この構造体の持つ`API`というフィールドを通じて[プラグイン用の各種API](https://developers.mattermost.com/extend/plugins/server/reference/#API)を実行できるようになります。
実際のAPIの呼び出し方については、次の`hooks.go`のセクションで説明しています。

```go:server/plugin.go
package main

import "github.com/mattermost/mattermost-server/plugin"

type FirstPlugin struct {
	plugin.MattermostPlugin
}
```

これでサーバーサイドのプラグインを実装することはできましたが、このプラグインはまだ何の機能も持っていません。
Mattermost本体の動作に影響を及ぼすには、いずれかのHooksを実装する必要があります。

https://developers.mattermost.com/extend/plugins/server/reference/#Hooks

## hooks.go

ここでは、`MessageWillBePosted`のHooksを実装し、Mattermost上にQiitaのリンクを含む投稿が行われた際に自動で`#Qiita`というハッシュタグを付与する機能を実装していきます。

```go:server/hooks.go
package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

func (p *FirstPlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	if strings.Index(post.Message, "https://qiita.com/") == -1 {
		return post, ""
	}

	p.API.LogDebug("Qiita link is detected.")
	post.Message = fmt.Sprintf("%s #Qiita", post.Message)
	post.Hashtags = fmt.Sprintf("%s #Qiita", post.Hashtags)
	return post, ""
}
```

`MessageWillBePosted`メソッドを実装することで、ユーザーから投稿されたメッセージがDBに保存される直前に実行される処理を記述できます。
ここでは、投稿内容に`https://qiita.com/`というURLを含む場合に、自動で`#Qiita`というハッシュタグを付与するよう実装になっています。

途中で`p.API.LogDebug`というメソッドを利用してログ出力を行なっています。このメソッドを通じて出力されたログは下記のように **システムコンソール** > **ログ** に出力されます。

> 2018-12-31T21:42:59.988+0900	debug	app/plugin_api.go:623	Qiita link is detected.	{"plugin_id": "org.kaakaa.mattermost-plugin-first"}

## Gopkg.toml

ここまで書けたら、依存ライブラリを取得するために`dep`コマンドを使用します。

`server`ディレクトリ配下で`dep init`を実行することで、依存ライブラリの検出・取得を行うことができます。

```
$ cd server
$ dep init
```

## ビルド

サーバーサイドの実装を含むプラグイン `.tar.gz`ファイルを作成していきます。

まず、`plugin.json`にサーバーサイドのプラグインを認識させるために`server`というセクションを追加します。

```json:plugin.json
{
    "id": "org.kaakaa.mattermost-plugin-first",
    "name": "Mattermost Plugin Sample",
    "description": "このプラグインはQiita記事用のMattermostプラグインです。",
    "version": "0.0.2",
    "server": {
        "executables": {
            "linux-amd64": "server/dist/plugin-linux-amd64",
            "darwin-amd64": "server/dist/plugin-darwin-amd64",
            "windows-amd64": "server/dist/plugin-windows-amd64.exe"
        },
        "executable": ""
    }
}
```

[`server`セクション](https://developers.mattermost.com/extend/plugins/manifest-reference/#server)には、.tar.gzファイル内のGoバイナリの位置を指定します。`executable`セクションでは各OSごとのバイナリを指定することもできます。

`plugin.json`に書かれた構成通りに`.tar.gz`ファイルが作成されるよう、`Makefile`を下記のように修正します。

```
PLUGIN_ID ?= org.kaakaa.mattermost-plugin-first
PLUGIN_VERSION ?= 0.0.2
BUNDLE_NAME ?= $(PLUGIN_ID)_$(PLUGIN_VERSION).tar.gz

build:
	mkdir -p server/dist
	cd server && env GOOS=linux GOARCH=amd64 GO build -o dist/plugin-linux-amd64
	cd server && env GOOS=darwin GOARCH=amd64 GO build -o dist/plugin-darwin-amd64
	cd server && env GOOS=windows GOARCH=amd64 GO build -o dist/plugin-windows-amd64.exe

	rm -rf dist/
	mkdir -p dist/$(PLUGIN_ID)
	cp plugin.json dist/$(PLUGIN_ID)/

	mkdir -p dist/$(PLUGIN_ID)/server/dist
	cp -r server/dist/* dist/$(PLUGIN_ID)/server/dist

	cd dist && tar -cvzf $(BUNDLE_NAME) $(PLUGIN_ID)

	@echo plugin built at: dist/$(BUNDLE_NAME)
```

`make build`を実行することでサーバーサイドの実装を含むプラグインファイルが作成されます。

## プラグインアップロード・有効化

先ほどと同様、**システムコンソール** > **プラグイン(ベータ版)** > **管理** よりプラグインファイルをアップロードし、同画面から**有効にする**リンクをクリックしてプラグインを有効化します。今度は「このプラグインは起動中です。」というメッセージが表示されると思います。

![スクリーンショット 2019-01-01 10.59.34.png](https://qiita-image-store.s3.amazonaws.com/0/9891/94b65442-26e0-1f1b-a765-980db0599500.png)

## 動作確認

![plugin.mov.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/548309ab-c5cf-b368-4384-92c5b64a888b.gif)

## まとめ

以上のように、Mattermostプラグインのサーバーサイドの実装を行うことができます。
今回は触れませんが、サーバーサイドの拡張ポイント(Hooks)は他にも数多くあるのでユースケースに合致する組み合わせを探して見てください。

https://developers.mattermost.com/extend/plugins/server/reference/#Hooks

* [OnActivate](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnActivate) / [OnDeactivate](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnDeactivate)
  - プラグインを有効化/無効化した時に呼ばれるメソッド
  - このメソッド内で[RegisterCommand](https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand)/[UnregisterCommand](https://developers.mattermost.com/extend/plugins/server/reference/#API.UnregisterCommand)のAPIを実行し、スラッシュコマンドを追加するようなプラグインで利用することができます (`ExecutedCommand`のHooksも実装する必要がある)
* [ExecutedCommand](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand)
  * プラグインを通じて追加されたコマンドが実行された時に呼ばれるメソッド
* [ServeHTTP](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ServeHTTP)
  * プラグイン専用のエンドポイントを追加することができます
  * 追加されたエンドポイントのルートは `${MATTERMOST_SITEURL}/plugins/${PLUGIN_ID}/` になります
* [OnConfigurationChange](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.OnConfigurationChange)
  * プラグインの設定が変更された時に呼び出されるメソッドです
  * マニフェストファイルに[setting_schema](https://developers.mattermost.com/extend/plugins/manifest-reference/#settings_schema)セクションを追加することで、プラグインの設定画面が出現します


# Step 3: フロントエンドプラグインの実装

最後にフロントエンドプラグインの実装について見ていきます。

フロントエンドの拡張ポイントは下記に挙げられています。
https://developers.mattermost.com/extend/plugins/webapp/reference/#registry

今回は[`registerChannelHeaderButtonAction`](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerChannelHeaderButtonAction)を使い、チャンネルヘッダーボタンから投稿を作成するサンプルを作っていきます。


```
root/
  ├ dist/
  │  └ org.kaakaa.mattermost-plugin-first_0.0.3.tar.gz
  ├ server/
  │  ├ ...
  ├ webapp
  │  ├ src/
  │  │  ├ index.js
  │  │  └ action.js
  │  ├ package.json
  │  └ webpack.config.js
  └ plugin.json
```

## index.js

プラグインの読み込みと実装を`index.js`というファイルに書いていきます。

```javascript:webapp/src/index.js
const React = window.React;
const {Glyphicon} = window.ReactBootstrap;

import {createPluginPost} from './action'

window.registerPlugin("org.kaakaa.mattermost-plugin-first", new Plugin());

class Plugin {
    initialize(registry, store) {
        registry.registerChannelHeaderButtonAction(
            <Glyphicon glyph="pencil" />,
            (channel) => {
                createPluginPost(channel.id)(store.dispatch, store.getState)
            },
            'First Plugin',
            'This is plugin tooltip',
        )
    }
}
```

フロントエンドプラグインは`window.registerPlugin`メソッドにプラグインIDとプラグインを実装したクラスのインスタンスを与えることで読み込まれます。
プラグイン実装クラス(`Plugin`)は`initialize`メソッドを持ち、このメソッドの[第一引数である`registry`が持つメソッド](https://developers.mattermost.com/extend/plugins/webapp/reference/#registry)を呼び出すことで機能を追加することができます。

今回はチャンネルヘッダーボタンを追加する部分を実装するため、[registry.registerChannelHeaderButtonAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerChannelHeaderButtonAction)を呼び出しています。

このメソッドの第１引数には、ボタンのアイコンとなるReact要素を指定します。

![スクリーンショット 2019-01-01 11.13.12.png](https://qiita-image-store.s3.amazonaws.com/0/9891/7bb6c0cc-c8ab-4346-4a0a-dc1e5a81083a.png)

MattermostプラグインではReact Bootstrapがexportされているため、`const {Glyphicon} = window.ReactBootstrap;`と書くことでReact Bootstrapの機能をimportすることができます。(exportされているライブラリ一覧は[Exported Libraries and Functions](https://developers.mattermost.com/extend/plugins/webapp/reference/#exported-libraries-and-functions))。

第２引数にボタンが押された時の動作を記述します。今回は投稿を作成する`createPluginPost`という処理をactionsとして実装しました。
第３引数には、複数のプラグインが`registerChannelHeaderButtonAction`を実装していた場合に、ドロップダウンメニューとして使われるテキストです。
第４引数にはチャンネルヘッダボタンのツールチップテキストです。

![スクリーンショット 2019-01-01 11.14.30.png](https://qiita-image-store.s3.amazonaws.com/0/9891/592b31ef-45d3-ffd0-9904-5f1e6ece8518.png)

## action.js

チャンネルヘッダーボタンを押した時に行われる処理を`action.js`に書いています。
MattermostのJavascriptドライバーである`mattermost-redux`を利用して、簡単なメッセージを投稿する機能を下記のように実装しています。

```javascript:webapp/src/action.js
import {createPost} from 'mattermost-redux/actions/posts';
import {getCurrentUserId} from 'mattermost-redux/selectors/entities/users';

export function createPluginPost(channelId) {
    return async (dispatch, getState) => {
        const state = getState();
        const userId = getCurrentUserId(state)
        const post = {
            channel_id: channelId,
            user_id: userId,
            message: "Post from webapp plugin",
        }
        return await dispatch(createPost(post));
    }
}
```

## ビルド

まず、`plugin.json`にフロントエンドのプラグインを認識させるために`webapp`というセクションを追加します。

```json:plugin.json
{
    "id": "org.kaakaa.mattermost-plugin-first",
    "name": "Mattermost Plugin Sample",
    "description": "このプラグインはQiita記事用のMattermostプラグインです。",
    "version": "0.0.3",
    "server": {
        "executables": {
            "linux-amd64": "server/dist/plugin-linux-amd64",
            "darwin-amd64": "server/dist/plugin-darwin-amd64",
            "windows-amd64": "server/dist/plugin-windows-amd64.exe"
        },
        "executable": ""
    },
    "webapp": {
        "bundle_path": "webapp/dist/main.js"
    }
}
```

実装してきた`.js`ファイルを、マニフェストファイルで指定したした`webapp/dist/main.js`にバンドルされるように`webpack.config.js`を用意します。
（フロントエンドに明るくないので[公式のサンプル](https://github.com/mattermost/mattermost-plugin-demo/blob/master/webapp/webpack.config.js)をベースにしています）

```javascript:webapp/webpack.config.js
var path = require('path');

module.exports = {
    entry: [
        './src/index.js',
    ],
    resolve: {
        modules: [
            'src',
            'node_modules',
        ],
        extensions: ['*', '.js', '.jsx'],
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env','@babel/preset-react'],
                        plugins: [
                            'transform-class-properties',
                            '@babel/plugin-proposal-object-rest-spread'
                        ],
                    },
                },
            },
        ],
    },
    externals: {
        react: 'React',
        redux: 'Redux',
        'react-redux': 'ReactRedux',
    },
    output: {
        path: path.join(__dirname, '/dist'),
        publicPath: '/',
        filename: 'main.js',
    },
};
```

下記のような`package.json`を用意し、`npm install`してから`npm run build`を実行し、正常にコマンドが完了すれば準備は完了です。

```json:webapp/package.json
{
  "name": "webapp",
  "version": "0.0.3",
  "description": "Post today's metal from channel header button",
  "main": "src/index.js",
  "scripts": {
    "build": "webpack --mode=production",
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC",
  "dependencies": {
    "mattermost-redux": "^5.6.2"
  },
  "devDependencies": {
    "@babel/core": "^7.2.2",
    "@babel/plugin-proposal-object-rest-spread": "^7.2.0",
    "@babel/preset-env": "^7.2.3",
    "@babel/preset-react": "^7.0.0",
    "babel-loader": "^8.0.4",
    "babel-plugin-transform-class-properties": "^6.24.1",
    "babel-plugin-transform-object-rest-spread": "^6.26.0",
    "webpack": "^4.28.3",
    "webpack-cli": "^3.1.2"
  }
}
```

最後にルートディレクトリにあるMakefileを下記のように編集し、で`make build`を実行することで、フロントエンドプラグインを含むプラグイン`.tar.gz`ファイルが`dist`ディレクトリ内に作成されるはずです。

```Makefile
PLUGIN_ID ?= org.kaakaa.mattermost-plugin-first
PLUGIN_VERSION ?= 0.0.3
BUNDLE_NAME ?= $(PLUGIN_ID)_$(PLUGIN_VERSION).tar.gz

build:
	mkdir -p server/dist
	cd server && env GOOS=linux GOARCH=amd64 GO build -o dist/plugin-linux-amd64
	cd server && env GOOS=darwin GOARCH=amd64 GO build -o dist/plugin-darwin-amd64
	cd server && env GOOS=windows GOARCH=amd64 GO build -o dist/plugin-windows-amd64.exe

	cd webapp && npm run build;

	rm -rf dist/
	mkdir -p dist/$(PLUGIN_ID)
	cp plugin.json dist/$(PLUGIN_ID)/

	mkdir -p dist/$(PLUGIN_ID)/server/dist
	cp -r server/dist/* dist/$(PLUGIN_ID)/server/dist
	mkdir -p dist/$(PLUGIN_ID)/webapp/dist;
	cp -r webapp/dist/* dist/$(PLUGIN_ID)/webapp/dist/;

	cd dist && tar -cvzf $(BUNDLE_NAME) $(PLUGIN_ID)

	@echo plugin built at: dist/$(BUNDLE_NAME)
```

## プラグインアップロード・有効化

先ほどと同様、**システムコンソール** > **プラグイン(ベータ版)** > **管理** よりプラグインファイルをアップロードし、プラグインを有効化します。

## 動作確認

![webapp.mov.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/7098d243-a8db-3af1-2295-bfa528be7d21.gif)

## まとめ

以上のように、Mattermostプラグインのフロントエンドの実装を行うことができます。
他にも[多くの拡張ポイント](https://developers.mattermost.com/extend/plugins/webapp/reference/#registry)があるので、実装方法については公式のサンプルリポジトリの実装を参照してみてください。
https://github.com/mattermost/mattermost-plugin-demo

# おわりに

以上がMattermost severプラグインの開発方法になります。

今まではMattermostを他のシステムと連携する場合、連携機能専用のサーバーを立ち上げなければならず、連携機能が増えるたびにサーバーを運用する手間が増えていましたが、プラグイン機能を使うことで運用の手間をMattermostインスタンスに一元化することができます。

また、プラグイン機能では今まで以上に細かな動作にまで手を加えることができるようになりました。
例えば、[mattermost/mattermost\-plugin\-antivirus](https://github.com/mattermost/mattermost-plugin-antivirus)のようにファイルがアップロードされた際に、そのファイルに対してセキュリティのチェックをかけたりできるようになります。

Slackとは違い、オンプレミスでの運用もメインターゲットとなるMattermostでは、プラグイン機能によるインフラ運用コストの削減や、企業特有のポリシーを実装するための細かなカスタマイズ性は非常に大きなメリットとなります。
プラグイン実装にまつわる本体のアップデートの追従などのメンテナンスコストは今後問題となる可能性がありますが、その辺りをどのようにクリアしていくかを含め、開発に協力しながらウォッチしていければと思います。


## 参考資料

### 開発者向けドキュメント
* https://developers.mattermost.com/
* https://github.com/mattermost/mattermost-developer-documentation

このサイトでは、プラグイン開発についてだけでなく、その他の統合機能であるウェブフックやスラッシュコマンドなどの開発方法や、Mattermostへのコントリビュート手順などについてもまとめられています。

### サンプルプラグイン
* https://github.com/mattermost/mattermost-plugin-demo

公式で開発されているデモ用のプラグインです。
プラグイン向けに提供されている拡張ポイントの多くを利用しているので、実装する際の参考になると思います。

* https://github.com/mattermost/mattermost-plugins

このリポジトリでは、Official/Unofficial問わずMattermostプラグイン開発プロジェクトのGitHubリポジトリがまとめられています。
文書よりも動くコードのほうが分かりやすいという方には参考になると思います。
