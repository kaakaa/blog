
---
title: "GrowiからMattermostへWebhook通知を試してみた"
date: 2018-07-08T00:02:04+09:00
draft: false
toc: true
tags: ["Mattermost","Growi"]
---

## はじめに

[Growi](https://growi.org/)ではイベントの通知先としてSlack Incoming Webhookを指定することができます。
Slack代替OSSであるMattermostはWebhookの仕様もSlackに準拠しているため、Slack Incoming Webhook通知の機能を利用してMattermostへ通知を送れるか試してみました。

* 動作環境
  * Growi 3.1.12
  * Mattermost 5.1.0-rc1

## Mattermostの設定

まず、Mattermost側にGrowiからのWebhookを受け取るための**内向きのウェブフック**を作成します。
別記事にて内向きのウェブフックの作成手順について書いたものがありますので、こちらの手順については下記の記事を参考にしてください。
参考: [esa.ioからMattermostへのWebhook送信を試してみた](https://qiita.com/kaakaa_hoe/items/44b4122cd38e3fd0d262#mattermost%E3%81%AE%E8%A8%AD%E5%AE%9A)

作成したMattermostの内向きのウェブフックのURLをコピーしておきます。

## Growiの設定
次にGrowiの設定を行います。

**管理** > **通知設定**より通知設定画面を開き、**Webhook URL**にMattermostで作成した内向きのウェブフックのURLを入力します。
<img alt="スクリーンショット 2018-07-07 23.33.30.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/b1c969f0-862e-10d7-baad-f719142031a5.png" width="70%"<img width="895" alt="スクリーンショット 2018-07-07 23.34.28.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/a62dc70a-12bd-23fa-70d7-174018c04489.png">

**Save**ボタンを押すことでGrowi側の設定は完了です。

## 通知の確認

Growi側で新しい投稿を作成します。

ここで、Mattermostへの通知を送信するには、投稿作成画面の右下にある**Slack通知**のチェックボックスをチェックし、その隣の通知先欄に通知を送信したいMattermostのチャンネルのIDを入力する必要があります。
MattermostのチャンネルIDはチャンネルのURLから確認できます: `https://${SITE_URL}/${TEAM}/channels/${チャンネルID}`。

<img width="895" alt="スクリーンショット 2018-07-07 23.34.28.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/f79bfb1e-5107-c0ca-3ddd-e52f9befcdfa.png">

上記の通知先の指定を行なってからGrowiの投稿を作成すると、Mattermostへ下記のような通知が送信されます。

<img width="837" alt="スクリーンショット 2018-07-07 23.37.24.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/3e5a820e-67ae-aaf3-6710-796a3965031a.png">




