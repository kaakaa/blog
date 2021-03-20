
---
title: "Mattermostチームお気に入り・10の生産性向上ハック"
date: 2019-05-26T00:03:36+09:00
draft: false
toc: true
tags: ["OSS","Mattermost"]
---
[The Mattermost team’s 10 favorite productivity hacks](https://mattermost.com/blog/the-mattermost-teams-10-favorite-productivity-hacks/)

先日、Mattermostチームがよく使っている時間節約のためのMattermostのショートカットやハックに関する記事を公開されていましので、内容を紹介します。


## 1. /rent-a-carlos

MattermostチームのCarlosが開発した、同僚の手を借りるのを簡単にするプラグイン。
`/rent-a-carlos`のスラッシュコマンドを使って、他のユーザー向けのタスクをダイレクトメッセージとして作成できます。

![rent-a-carlos.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/b0b4b7e8-72bb-804e-0146-440d31d76b07.gif)


## 2. タイムスタンプを右クリックして投稿へのリンクを取得する

Mattermostの投稿へのリンクをは、投稿のドロップダウンメニューから取得できます。しかし、リンクを取得するまでには何回かボタンをクリックする必要があります。

この投稿へのリンクは、投稿のタイムスタンプをクリックすることでも取得できるため、タイムスタンプを右クリックしてリンクのアドレスをコピーすることで少し時間を短縮することができます。


## 3. TODOリストとしてフラグを利用する

TODO管理のアプリケーションを導入する代わりに、Mattermost上の重要なメッセージにフラグを付けることでTODOリストを代用することができます。

投稿のタイムスタンプの横にあるフラグアイコンをクリックすることでフラグを立てることができ、Mattermost画面右上のフラグアイコンをクリックすることで、今までフラグを立てた投稿の一覧を表示することができます。
フラグを立てた投稿は自分だけが表示可能なため、自分用のTODOリストとして活用することができます。

![flag.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/650afd45-7978-0bf3-370a-44a2f0e4d213.gif)


## 4. キーボードショートカット

キーボードショットカットではチャンネルの移動やファイルのアップロード、情報の表示などを素早く行うことができます。
`CTRL + /` (Macでは`CMD + l`)を入力してショートカットメニューを表示してみてください。

![shortcut.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/ec2940ec-77e4-f681-186c-3eac80560ad9.gif)


## 5. `SHIFT+UP`でスレッドへの返信

`SHIFT + UP` (Macでは`OPTION + SHIFT + UP`)のショートカットを使うことで特定のスレッドに素早く返信できるようになります。

![thread.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/e2a57917-214c-bd3c-22cc-1979ebe76f03.gif)


## 6. 検索機能

組織の知識はMattermostに蓄積されているため、多くの人がGoogleを利用するようにMattermostの検索機能を使用しています。
特定の結果を検索するために`from:`, `in:`, `on:`, `before:`, `after:` のなどのパラメーターを使って、検索機能を便利にすることができます。

## 7. 重要なメッセージをチャンネルにピン留めする

フラグは重要なメッセージを内緒でキープすることができる機能です。一方、投稿のピン留めは重要なメッセージをチャンネル内の全員に対してキープする機能です。
投稿のドロップダウンメニューから投稿のピン留めを実行することで、メッセージをチャンネルにピン留めすることができます。ピン留めされたメッセージを確認するにはMattermost画面の右上にある押しピンのアイコンをクリックしてください。

![pin.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/d0d89b9a-23b9-a907-2b0b-93aa76a17a08.gif)


## 8. サイドバー構成（実験的な機能）

まだ実験的な機能ではありますが、**アカウント設定 > サイドバー > チャンネルのグループ化とソート**から設定できるサイドバー構成オプションはMattermost上での作業を効率化すると確信しています。
未読メッセージのあるチャンネルをグループ化したり、投稿がもっとも最近行われた順にチャンネルをそーとしたり、全てのチャンネル種別（公開チャンネル・非公開チャンネル・ダイレクトメッセージ）を一つのリストの結合したりできます。

![sidebar.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/20bc2983-1a38-3c24-b6e4-da080fbe4eb9.gif)


## 9. `+:絵文字名:`による絵文字リアクション

Mattermostを使えば使うほど、絵文字の名前に詳しくなっていくはずです。
メッセージ入力欄で`+:絵文字名:`と入力してメッセージを送信すると、一つ前の投稿に対して絵文字リアクションを付けることができます。少しずつ時間を節約して会話をもっと楽しみましょう。

![emoji.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/b633f475-3531-0ecb-bd42-f8a0814a1bbb.gif)

## 10. `CTRL + K` / `CMD + K`によるチャンネル切り替え

とても重要なもう一つのショートカットは、実は全てのMattermostインスタンスに目立つように表示されていました。 `チャンネル切替 - CTRL + K`

![switch.mov.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/8a3d9b54-923e-8e54-cd37-dd4a417f3027.gif)




