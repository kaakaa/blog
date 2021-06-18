---
title: "Mattermostプラグイン用のリポジトリテンプレート"
date: 2019-07-13T15:05:21+09:00
draft: false
toc: true
tags: ["mattermost", "plugin"]
---

## Mattermost Plugin用テンプレートリポジトリ

Mattermostにはプラグイン機能があり、サーバーサイド/フロントエンド共に独自の拡張機能を追加することができます。
Mattermostプラグインの開発方法は[Mattermostプラグインの作り方 · kaakaa blog](https://blog.kaakaa.dev/post/mattermost/plugin_tutorial/)にも書きましたが、一から作るとなると準備するファイルも多く大変です。

そこで、Mattermostプラグインの開発を楽にするため、Mattermostチームが[GitHubのテンプレートリポジトリ機能](https://help.github.com/ja/articles/creating-a-template-repository)を利用して作成しているテンプレートリポジトリを使ったプラグイン開発方法について紹介します。

* [mattermost/mattermost\-plugin\-starter\-template: Build scripts and templates for writing Mattermost plugins\.](https://github.com/mattermost/mattermost-plugin-starter-template)

このテンプレートリポジトリを使用することで、動作するプラグインがある状態から開発を始めることができ、必要な機能を開発することに集中することができるようになります。

## テンプレートリポジトリを使用してMattermostプラグインを開発する

### 0. Mattermostプラグイン開発環境について

プラグイン開発を行うには下記の開発環境が必要です。

* サーバー側のプラグイン実装を行う場合
  * Go: 1.12
* webapp (フロントエンド) 側のプラグイン実装を行う場合
  * Node 10.15.3+
  * npm 6.4.1+
* どちらでも必要
  * GNU Make

プラグイン開発向けの開発環境はまとめられていなようですが、Mattermost本体の推奨環境に沿うのが安全だと思います。

* 参考
  * https://developers.mattermost.com/contribute/server/developer-setup/
  * https://developers.mattermost.com/contribute/webapp/developer-setup/

### 1. テンプレートからリポジトリを作成する
まず、テンプレートリポジトリを元に、作成したいプラグイン用のリポジトリを作成します。

[mattermost/mattermost\-plugin\-starter\-template](https://github.com/mattermost/mattermost-plugin-starter-template)をブラウザで開くと、`Use this template`というボタンがあります。

![スクリーンショット 2019-07-13 11.15.37.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/25b49b1c-3096-f135-219f-352cd741f2cf.png)

このボタンを押すとリポジトリ作成画面へ遷移するので、そこでリポジトリ名を入力してリポジトリを作成するだけでテンプレートからリポジトリを作成できます。

テンプレートリポジトリの動作はリポジトリのForkと似ていますが、下記の点でForkとは異なります。

> * 新しいフォークは、親リポジトリのコミット履歴すべてを含んでいますが、テンプレートから作成されたリポジトリには、最初は 1 つのコミットしかありません。
> * フォークへのコミットはコントリビューショングラフに表示されませんが、テンプレートから作成されたリポジトリへのコミットはコントリビューショングラフに表示されます。
> * フォークは、既存のプロジェクトにコードをコントリビュートするための一時的な方法となります。テンプレートからリポジトリを作成することは、新たなプロジェクトを初めから素早く始める方法です。

_引用元: [テンプレートからリポジトリを作成する \- GitHub ヘルプ](https://help.github.com/ja/articles/creating-a-repository-from-a-template)_

#### Security Alert

テンプレートからリポジトリを作成したときに、セキュリティ脆弱性に関するアラートが出ることがあります。

![スクリーンショット 2019-07-13 10.58.36.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/c2a17096-6677-6a08-3429-61280efac613.png)


これは、リポジトリが依存しているライブラリに脆弱性が含まれる可能性があることを示しており、本体のテンプレートリポジトリ側でも随時ライブラリのアップデートは行われいますが、リポジトリを作成するタイミングによってはこの通知が出てしまうことがあります。

[リポジトリ内の脆弱な依存関係を表示・更新する \- GitHub ヘルプ](https://help.github.com/ja/articles/viewing-and-updating-vulnerable-dependencies-in-your-repository)を参考にするなどして、ライブラリのアップデートを行ってください。

#### テンプレートリポジトリの内容について

テンプレートリポジトリに含まれるファイルは下記の通りとなります。

```
.
├── .circleci/      - CircleCI設定ファイル(プロジェクト依存の情報がないのでこのまま使用できます)
├── assets/         - プラグインから利用するファイル(アイコン画像など)の格納先
├── build/          - Makefileから呼ばれるビルド用のスクリプト格納ディレクトリ
                      基本的にこのディレクトリのファイルは編集しません。
├── public/         - (assetsと同じ用途？)
├── server/         - サーバー用のプラグインコードが格納ディレクトリ
├── webapp/         - フロントエンド用のプラグインコードディクトり
├── .editorconfig   - editorconfigの設定ファイルです https://editorconfig.org/
├── .gitignore
├── CHANGELOG.md    - 更新履歴を記述ファイル（手動更新）
├── go.mod          - Go(サーバープラグイン)の依存ライブラリ管理ファイル
├── go.sum          - 同上
├── LICENSE         - Apache License v2のライセンスファイル
├── Makefile        - ビルド用のMakefile
├── plugin.json     - プラグインのメタデータを記述するファイル
┗── README.md 
```

### 2. プラグインをビルドする

作成したプラグインをcloneして、ビルドしてみましょう。

#### 2.1. 前準備

まず、リポジトリをクローンします。

```
git clone --depth 1 https://github.com/mattermost/mattermost-plugin-starter-template
```

注意: `https://github.com/mattermost/mattermost-plugin-starter-template`の部分は作成したリポジトリのURLに変更してください。


テンプレートリポジトリは、ほぼ何も変えなくてもそのままでビルド・デプロイを行うことができます。ただ、`plugin.json` に書かれているプラグインのIDなど必ず変更しておくべき箇所が何点かあります。

* **plugin.json**
    * https://github.com/mattermost/mattermost-plugin-starter-template/blob/master/plugin.json#L2
    * `id`は必ず変更してください。併せて`name`、`description`なども変えておくと良いと思います。
    * このファイルの内容については [Manifest Reference](https://developers.mattermost.com/extend/plugins/manifest-reference/) を参照してください、
* **LICENSE**
    * ライセンスの著作権者の欄がプレースホルダーとなっているため、自身の情報に変更しておきましょう
    * https://github.com/mattermost/mattermost-plugin-starter-template/blob/master/LICENSE#L189
        * `Copyright [yyyy] [name of copyright owner]` => `Copyright 2019 Yusuke Nemoto`
* **README.md**
    * 内容がテンプレートリポジトリのものになっているため更新しておきましょう

上記を変更したら、まずはプラグインをビルドしてみましょう。

#### 2.2. ビルド

```
$ make
```

上記コマンドを実行するだけでビルドが完了します。
ビルドが正常に終了すると、`dist/`というディレクトリが作成され、その中に `.tar.gz` のファイルがあるはずです。これがMattermostプラグインファイルになります。

この `.tar.gz` ファイルをアップロードすることでMattermostプラグインを有効にすることができます。プラグインのアップロード方法などは下記記事で紹介しているので参照ください。

* https://blog.kaakaa.dev/post/mattermost/plugin_tutorial/

#### 2.3. 動作確認

このテンプレートリポジトリのプラグインが持つ機能は、[プラグイン用に新しいエンドポイントを作成する](https://github.com/mattermost/mattermost-plugin-starter-template/blob/master/server/plugin.go#L24)のみのため、動作確認もこれで行います。

```
$ curl http://localhost:8065/plugins/${PLUGIN_ID}/
Hello, world!
```

`http://localhost:8065`はMattermostのSiteURLに、${PLUGIN_ID} は `plugin.json` に記述したプラグインIDに置き換えて実行してください。

#### 便利なプラグインアップロード方法

プラグインの開発を始めると、コード修正 -> アップロード -> 動作確認というフローを繰り返すことになります。その度に画面からアップロードするのは面倒です。
そこで、テンプレートリポジトリのMakefileにはコマンドラインからプラグインのビルド -> アップロードを行ってくれる`make deploy`というタスクがあります。

このタスクを実行するには、下記のように環境変数にMattermostサーバーのURLとアップロードするMattermostユーザーの認証情報を設定し、タスクを実行するだけです。

```
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
export MM_ADMIN_USERNAME=admin
export MM_ADMIN_PASSWORD=password
make deploy
```

もし、すでに同じIDのプラグインがアップロード済みでも新しいプラグインに上書きを行ってくれます。この方法だと、プラグインのバージョンが下がっている場合でも強制的に上書きされてしまうので注意が必要です。

このアップロード方法は開発時のみの利用が推奨されています。

### 3. プラグインの開発を開始する

プラグインのアップロードが完了し、動作することまで確認できたらあとは機能の開発を進めていくだけです。

ここで、もしサーバー側の機能のみを実装する場合は、`webapp`ディレクトリは削除しても構いません。逆にフロントエンド側のみの機能を実装する場合は `server` ディレクトリを削除できます。この辺りはMakefile内でディレクトリの存在確認を行い、よしなに動作してくれます。


プラグインの開発については下記のサイトの情報が参考になると思います。

* [Plugins \(Beta\)](https://developers.mattermost.com/extend/plugins/)
* [mattermost/mattermost\-plugin\-demo: A demo of what Mattermost plugins can do\.](https://github.com/mattermost/mattermost-plugin-demo)
* [Mattermostプラグインの作り方 · kaakaa blog](https://blog.kaakaa.dev/post/mattermost/plugin_tutorial/)

また、Mattermostコアチームによるプラグイン開発に関する会話は下記のチャンネルで行われています。

* https://community.mattermost.com/core/channels/developer-toolkit


