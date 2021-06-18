
---
title: "esa.ioからMattermostへのWebhook送信を試してみた"
date: 2018-07-07T23:19:23+09:00
draft: false
toc: true
tags: ["esa.io","Mattermost"]
---

## はじめに

esa.ioのWebhook送信先としてMattermostが選べるようになったので試してみました。
[ReleaseNotes/2018/06/21/Discord・Microsoft Teams・MattermostへのWebhookを追加しました \- docs\.esa\.io](https://docs.esa.io/posts/273)

使用しているMattermostのバージョンはMattermost v5.1.0-rc1になります。

## Mattermostの設定

まず、Mattermost側にesa.ioからのWebhookを受け取るための内向きのウェブフックを作成する必要があります。

Mattermostのメインメニューから**統合機能**を選択します。
<img alt="スクリーンショット 2018-07-07 22.39.24.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/258af11d-5a3c-43e2-250c-7985ecb64efc.png" width="20%">


統合機能のメニューから**内向きのウェブフック**を選択し、**内向きのウェブフックを追加する**ボタンを押します。
<img alt="スクリーンショット 2018-07-07 22.39.47.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/269e7bec-a002-deaf-742c-314c8a97963f.png" width="70%">

ウェブフックを作成する際に必須の項目は**タイトル**と、通知を投稿する**チャンネル**のみですので、この２つは必ず指定してください。その他の情報は必要に応じて指定してください。
入力か完了したら**保存する**ボタンを押します。

<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/cb8bbc85-1c2f-5b2e-3cba-e265868eccab.png" width="70%">

すると、新しい内向きのウェブフックが作成されるため、表示されているURLをコピーしておきます。
<img alt="スクリーンショット 2018-07-07 23.00.50.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/3888ac41-6e30-a06f-b9ea-4cdabe0789f8.png" width="70%">

これでMattermost側の設定は完了です。

## esa.ioの設定
次にesa.ioの設定を行います。

`https://{team}.esa.io/team/webhooks`よりウェブフックのメニューを開き、**Add webhook**ボタンを押します。
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/8f5bb3d1-7175-5352-98eb-09576bab04bd.png" width="70%">

ウェブフックのメニューから**Mattermost**を選択し、**Incoming Webhook URL**に、先ほどコピーしたMattermostの内向きのウェブフックのURLをペースとします。
Labelやチェックボックスなどは必要に応じて変更してください。今回は投稿を作成した際に通知が飛ぶようにするため、**on post create (only ShipIt)**にチェックをしています。
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/d0fe802e-a680-43e4-c95f-3fc731a4046f.png" width="50%">

**Save**ボタンを押してウェブフックが作成されたらesa.ioの設定も完了です。

<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/eb4d663f-ec6d-804a-f251-e0c0df8e980e.png" width="70%">

## 通知の確認

esa.io側のウェブフックの設定を行なった段階で、Mattermostへ通知が飛びます。
Mattermostで内向きのウェブフックを作成した時に選択したチャンネルを開いてみましょう`(\( ⁰⊖⁰)/)`。
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/ce0af763-27bc-b89d-3f3e-1c6eb494e707.png" width="100%">

次に、esa.io側で新しい投稿を作成し、`ShipIt`します。
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/03168f23-5f3f-37f0-ecfa-d2e1d9758345.png" width="70%">

すると、Mattermostの方に通知が飛びます。
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/46b29361-b612-49ea-57ca-388d6d3eb7a5.png" width="100%">

投稿されるメッセージはesa.io側の`Change log`の内容のようです。
リンクになっている`esa.ioへの通知テスト`をクリックすると、esa.ioで作成された投稿へ飛ぶことができます。



