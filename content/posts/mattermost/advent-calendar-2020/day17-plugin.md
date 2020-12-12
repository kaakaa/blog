---
title: "[Mattermost Integrations] Plugin基本編"
date: 2020-12-17T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第17日目の記事です。

本記事では、Mattermostのプラグイン機能について紹介します。Mattermost Pluginは機能が多いため、本日はMattermostの概要と使い方、Mattermost公式チームが用意しているMattermost Plugin開発のためのテンプレートプロジェクトの紹介までを行います。

## Mattermost Plugin

Mattermost Pluginは、Mattermostの様々な箇所に存在する拡張ポイントを使ってより幅広く、コアに近い部分に処理を差し込むことができる機能です。今まで紹介してきた統合機能に比べて、下記２点で優れていると個人的に思っています。

1. 投稿以外のユーザーの行動に対して独自の処理を追加することができる
2. Mattermostと同じサーバー上で動かすことができ、プロセス管理もMattermostに任せられる

`1.`の点については、今まで紹介してきた統合機能は、メッセージの投稿やSlash Commandの実行などMattermost上で投稿操作が行われた際に独自の処理を実行することができるというものでした。しかし、Plugin機能では、ユーザーによる投稿操作の他にもチャンネル/チームへの参加/脱退やチャンネル作成時など、Mattermost上の様々な操作に反応して独自の処理を実行させることができます。また、投稿操作についても、投稿がデータベースに保存される前後で処理を実行することができるため、投稿が作成される前にRejectするようなこともできます。

`2.`の点についてはWebSocket APIを使うことで実現できたかもしれませんが、WebSocket APIを始めMattermostの統合機能を使うには、Mattermostとは別に統合機能からのリクエストを受け取るサーバーアプリケーションを構築しておく必要がありました。しかし、Mattermost Pluginは、そのプロセス管理をMattermost自身が行うため、Mattermost以外のサーバーを管理する必要はありません。これが`2.`の利点です。さらに、Mattermost PluginではPlugin用のKey-Valueストアも用意されているため、簡単なデータの保持であればMattermostに任せることができます。

Mattermost Pluginは、サーバー側で動作する**Server**サイドのプラグインとブラウザ上で動作する**Webapp**サイドのプラグインがあります。一つのPluginに両方の機能を持たせても良いですし、どちらか一方の機能だけのPluginを実装することもできます。

## Mattermost Pluginの使用

まず、Mattermost Pluginを使ってみるところから紹介します。

今回は、私も開発に参加しているMattermost上で投票機能を作成することができる[Matterpollプラグイン](https://github.com/matterpoll/matterpoll)を手動でインストールしてみます。

https://github.com/matterpoll/matterpoll

![matterpoll](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day17/matterpoll.png)

### 設定

まず、Mattermost Pluginを使うには、**システムコンソール > プラグイン管理 > プラグインを有効にする**が`有効`になっている必要があります。

![config plugin](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day17/config-plugin.png)

また、手動でプラグインのインストールを行う場合、**システムコンソール > プラグイン管理 > プラグインをアップロードする**の`ファイルを選択する`が有効になっている必要があります。もし、このボタンが無効化されている場合は、Mattermostが起動しているサーバーのMattermostインストールディレクトリにある`config/config.json`の`EnableUploads`を`true`にし、Mattermostを再起動する必要があります。

```config.json
    ...
    "PluginSettings": {
        "Enable": true,
        "EnableUploads": true,
    ...
```

### インストール

Pluginを使用する設定が完了したら、以下のページから`com.github.matterpoll.matterpoll-1.4.0.tar.gz`をダウンロードして、**システムコンソール > プラグイン管理 > プラグインをアップロードする**からアップロードします。

https://github.com/matterpoll/matterpoll/releases/tag/v1.4.0

(Matterpoll v1.4.0は**Mattermost v5.20**以上を対象としているため、それより古いバージョンを使用している場合は、Matterpollも古いバージョンのものをお使いください)

インストールが正常に終了すると、`インストール済みプラグイン`のところにMatterpollが表示されます。Matterpollプラグインのところにある`有効にする`のリンクをクリックすることでプラグインを有効にできます。

![install plugin](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day17/install-plugin.png)

### 実行

Matterpollプラグインが問題なく起動できれば、Matterpollの機能が利用できるようになっています。MatterpollはSlash Commandで投票を作成するため、チャンネルに戻り `/poll help` と打つと、Matterpollのヘルプメッセージが出てきます。

![execute plugin](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day17/execute-plugin.png)


## Mattermost Pluginの開発
先ほども見たように、Mattermost Pluginの本体は`.tar.gz`ファイルになります。この`.tar.gz`ファイルには、プラグインのマニフェストファイルや、Serverサイド機能のバイナリ、Webappサイドの`.js`ファイルなどが含まれています。

毎回一からこれらを用意するのは大変なため、MattermostチームはMattermost Plugin開発用のテンプレートリポジトリを公開しています。

https://github.com/mattermost/mattermost-plugin-starter-template

このテンプレートリポジトリを使ったMattermost Pluginの開発方法は、以下の記事で既に紹介しているため、ここでの説明は割愛します。

[Mattermostプラグイン用のGitHubリポジトリテンプレートを使ってみる \- Qiita](https://qiita.com/kaakaa_hoe/items/6f3d1aa0a126f2e94e01)

## Mattermost Demo Plugin

Mattermost PluginではMattermostの様々な箇所を拡張できますが、Pluginで拡張できる箇所を紹介するDemo用のプラグインが公開されています。

https://github.com/mattermost/mattermost-plugin-demo

このプラグインをインストールすることで、全てではないですがMattermost Pluginで拡張可能な箇所を確認できます。

## さいごに

本日は、Mattermost Pluginの概要や使い方について紹介しました。
明日からは、Mattermost Pluginで実現できる機能について紹介していきます。
