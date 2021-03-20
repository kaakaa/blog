---
title: "Mattermost Plugin Zoom"
date: 2020-04-10T21:26:17+09:00
draft: false
toc: true
tags: ["mattermost", "zoom", "plugin"]
---

# MattermostからZoomのミーティングに参加する

Mattemrostではテキストでの議論が主になるため、テキストでは伝わり切らないニュアンスを伝えたい場合などに会話による調整が必要になる場面もあると思います。Mattermost社ではミーティングにZoomを利用しているというのもあり、Zoomミーティングを開催しやすくするプラグインを開発しています。

本記事はMattemrost Zoomプラグインの使い方について紹介します。

注記: Work From Homeの流れで必要性が増した気がしたので紹介しようとしましたが、Zoomのセキュリティ対応により現在は上手く動作しないようです。
[April 2020: Setting updates for free accounts and single Pro users – Zoom Help Center](https://support.zoom.us/hc/en-us/articles/360041408732-April-2020-Setting-updates-for-free-accounts-and-single-Pro-users)

具体的には、Zoom参加ボタンを押すとZoomミーティング参加のためのパスワードの入力が求められるようになりましたが、そのパスワードがどこからも取得できないため、開催者以外がZoomミーティングに参加できないという問題です。既にIssueが立っているので、利用される方は下記Issueの動向を確認してください。
https://github.com/mattermost/mattermost-plugin-zoom/issues/99


## 環境
* Mattermost v5.21.0 

## インストール

**プラグインのインストールにはシステム管理者の権限が必要です**

### プラグインマーケットプレース

Mattermost v5.16からプラグインのマーケットプレース機能が使えるようになり、UIからプラグインをダウンロードできます。

**メインメニュー > プラグインマーケットプレース**からマーケットプレース画面を開ます。(プラグインマーケットプレースメニューは**システム管理者**にしか見えません)

表示されたプラグインからZoomプラグインを探し、 **インストール** のボタン押し、少し待てばZoomプラグインがインストールされます。

![](/blog/images/posts/mattermost/mattermost-plugin-zoom/marketplace.png)

インストールが完了すると、**インストール**ボタンが**設定**ボタンに変わるはずです。

### 手動インストール

プラグインマーケットプレースメニューが表示されない、もしくはマーケットプレースへ接続できない、という場合は手動でプラグインをインストールします。手動でのプラグインのインストールにも、**システム管理者**の権限が必要です。

手動でインストールするには、まず、GitHubから現時点(2020/4/10)での最新版である `zoom-1.3.0.tar.gz` をダウンロードします。`v1.3.0` は、Mattermost v5.12 以上で使用できます。

https://github.com/mattermost/mattermost-plugin-zoom/releases

**システムコンソール > プラグイン(ベータ版) > プラグイン管理** を開き、 **プラグインをアップロードする**メニューから先ほどダウンロードした `zoom-1.3.0.tar.gz` をアップロードします。

(もし、**ファイルを選択する**ボタンが非活性になっている場合、Mattermostサーバーの設定ファイル`config.json`の **PluginSettings.EnableUploads**設定を `true` に手動で修正し、サーバーを再起動する必要があります。)

アップロードが完了すると、**プラグイン管理**の設定画面下部の**インストール済みプラグイン**に Zoomプラグインが表示されているはずです。

![](/blog/images/posts/mattermost/mattermost-plugin-zoom/running.png)

## 設定
プラグインを起動する前にZoomの接続設定を行う必要があります。

### Zoom API Key/Secret取得
Zoomプラグインの設定を行うにはZoomのAPI KeyとAPI Secretが必要です。

1. https://marketplace.zoom.us/ にアクセスし、ログインする
2. 右上のメニューから**Develop > Bulid App**をクリックする
3. JWTの **Create** ボタンをクリックし、任意のApp Nameを入力して **Create** ボタンを押します
![](/blog/images/posts/mattermost/mattermost-plugin-zoom/jwt.png)
4. 必須項目を入力し **Continue** ボタンを押すと、API KeyとAPI Secretが表示された画面に遷移するためコピーしておきます

### Mattermost Zooom Plugin設定

**システムコンソール > プラグイン(ベータ版) > Zoom**を開き、設定を入力していきます。

* **Zoom URL**: (パブリックのZoomサービスを使う場合は空欄)
* **Zoom API URL**: (パブリックのZoomサービスを使う場合は空欄)
* **API Key**: (先ほどコピーしておいたAPI Key)
* **API Secret**: (先ほどコピーしておいたAPI Secret)
* **Webhook Secret**: (**再生成する**ボタンを押してランダム文字列を生成しておく)

以上の設定が完了したら、再び**システムコンソール > プラグイン(ベータ版) > プラグイン管理 > インストール済みプラグイン**の設定へ移動し、Zoomプラグインの**有効にする**ボタンを押し、プラグインを起動してください。(もし、エラーとなっている場合は一度**無効にする**ボタンで停止してから、もう一度起動し直してください)

`このプラグインは起動中です。`のメッセージが表示されていれば、起動完了です。もし、起動に失敗した場合は、**システムコンソール > レポート > サーバーログ**からZoomプラグインのエラーログを確認してください。

## 実行

プラグインの起動が完了すると、Mattermostのチャンネルヘッダ部分にビデオカメラのアイコンが表示されているはずです。

![](/blog/images/posts/mattermost/mattermost-plugin-zoom/meeting-button.png)

このアイコンをクリックすることで、チャンネルにZoomミーティングへの参加ボタンが投稿されます。

![](/blog/images/posts/mattermost/mattermost-plugin-zoom/meeting-post.png)

Mattermostのチャンネル参加者は、この参加ボタンをクリックすることでZoomミーティングに簡単に参加できるようになります。

## おわりに

MattermostにもZoom以外にも、ビデオカンファレンスサービスと連携するための統合機能がコミュニティベースで公開されています。
https://integrations.mattermost.com/category/directory/video-conferencing/#main