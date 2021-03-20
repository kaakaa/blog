---
title: "[Mattermost Integrations] Incoming WebHook 発展編"
date: 2020-12-03T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "webhook", "incoming"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第3日目の記事です。

前回の記事では、外部アプリケーションからMattermostの投稿を作成することができるIncoming WebHook(内向きのウェブフック)の機能について紹介しましたが、本記事はその続きです。Incoming WebHookを作成する際に利用できるパラメータについて紹介します。

## Incoming WebHookパラメータ

### `text`

Incoming WebHookを通じて投稿するメッセージであり、後述の`attachments`を指定しない場合は必須のパラメータです。Markdown記法や、`@`によるメンションを投稿することもできます。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"text": "## Test\n@channel\n* foo\n* bar"}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![use text](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/example-text.png)

### `channel`

投稿先のチャンネルを指定することができます。ここで指定するチャンネル名は、サイドバーに表示されているチャンネル表示名でなく、URLとして使われているチャンネル名になる点に注意が必要です。

![channel name](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/channel-name.png)

また、`@kaakaa`のように、`@`+ユーザー名を指定するとDM(ダイレクトメッセージ)として投稿することも可能です。`channel`を指定しない場合、Incoming WebHook作成時に指定したデフォルトチャンネルに投稿されます。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"text": "Sample Message", "channel": "incoming-webhook"}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![use channel](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/example-channel.png)

Incoming WebHook作成時に`このチャンネルに固定する`にチェックをしていた場合、デフォルトのチャンネル以外に投稿しようとすると403エラーとなります。

```bash
$ curl \
>   -H "Content-Type: application/json" \
>   -d '{"text": "Sample Message", "channel": "incoming-webhook"}'  \
>   http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr

{"id":"web.incoming_webhook.channel_locked.app_error","message":"This webhook is not permitted to post to the requested channel.","detailed_error":"","request_id":"qe3eyht563d1b8xani3g93dgfh","status_code":403}
```

### `username`
### `icon_url`

`username`と`icon_url`のパラメータを指定することで、Incoming WebHookを実行して作成される投稿のユーザー名とアイコンの表示を変更することができます。デフォルトでは、ウェブフックを作成したユーザーになります。

このパラメータを使用する場合、システムコンソールから下記の設定をそれぞれ`有効`にする必要があります。
* **システムコンソール > 統合機能管理 > 統合機能によるユーザー名の上書きを許可する**
* **システムコンソール > 統合機能管理 > 統合機能によるプロフィール画像アイコンの上書きを許可する**

![config overwrite](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/config-overwrite.png)

（この設定を有効にすると、`username`と`icon_url`を指定しない場合の投稿の作成者がデフォルトのアイコン、ユーザー名(`webhook`)に変更されます）

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"text": "Sample Message", "username": "new-bot", "icon_url": "http://www.mattermost.org/wp-content/uploads/2016/04/icon_WS.png"}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![user overwrite](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/example-overwrite.png)

ユーザー名とプロフィール画像アイコンの上書きを許可すると、Incoming WebHook作成画面にも`ユーザー名`、`プロフィール画像`の指定画面が現れるようになります。

![create with overwrite](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/create-with-overwrite.png)


### `icon_emoji`

`icon_url`では、アイコン画像に使用する画像のURLを指定しましたが、`icon_emoji`を使うことでアイコン画像に絵文字を指定することもできます。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"text": "Sample Message", "username": "new-bot", "icon_url": "http://www.mattermost.org/wp-content/uploads/2016/04/icon_WS.png", "icon_emoji": "smiley"}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![use icon-emoji](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/example-icon-emoji.png)

`icon_url`と`icon_emoji`を同時に指定すると、`icon_emoji`の方が採用されるようです。

### `attachments`

`attachments`は、Mattermostにリッチな投稿を作成することができる[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)を指定できる機能です。[MessageAttachments](https://docs.mattermost.com/developer/message-attachments.html)については、第9日目の記事で紹介します。

### `type`
`type`は、Mattermostの投稿の種別を指定するパラメータですが、基本的に使用することはありません。Plugin(Webapp)にてプラグイン独自の投稿種別を指定した場合のみに使用するものかと思います。Plugin(Webapp)については、第20日目以降の記事で紹介します。

### `props`
`props`は投稿のメタデータを格納する場所であり、他の統合機能から利用されることを想定して用意されており、基本的に決まったデータ構造などはありません。`card`というパラメータのみ特別で、`card`に格納されたテキストデータは、ウェブフックにより作成された投稿のコンテキストデータとしてMattermostの画面上で確認することができます。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"text": "Sample Message", "props": {"card": "## Sample\n* foo\n* bar"}}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![use card](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day3/example-card.png)

`props.card`を指定して作成された投稿には、メッセージのヘッダ部分に `i` ボタンが表示されるようになり、`i` ボタンをクリックすると右サイドバーに`props.card`に指定したテキストが表示されます。`props.card`にはMarkdownを使用することもできますが、`@`によるメンションは機能しません。(ユーザープロファイルの確認はできます)
 
Incoming WebHookによる投稿が長くなってしまう場合などに有効だと思います。

ただし、`card`による投稿はモバイルアプリからは確認できない点に注意が必要です。また、`card`はMattermost v5.14から利用可能です。

## さいごに

Incoming WebHookの詳細な使い方について紹介しました。
明日からは、Outgoing WebHookについて紹介していきます。
