---
title: "[Mattermost Integrations] Message Attachments"
date: 2020-12-09T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "messageattachments"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第9日目の記事です。

本記事では、Mattermostの統合機能から投稿を作成する際に、視覚的にリッチな投稿を作成することができるMessage Attachmentsの機能について紹介します。

https://docs.mattermost.com/developer/message-**attachments.html

## Message Attachmentsの概要
Message AttachmentsはMattermost上に下記のような投稿を作成するための機能です。（画像は[公式ドキュメント](https://docs.mattermost.com/developer/message-attachments.html)より）

![message attachment](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/official-example.png)

MattermostではデフォルトでMarkdown記法(設定によってはkatex記法も)をサポートしているため、簡単な見出しやリンクなどは普通の投稿で表すことができますが、多くの情報を通知したい場合、Markdownだと情報量の分だけ投稿が縦に伸びていってしまうため視認性が悪いという問題があります。
そのような場合にMessage Attachmentsを使うことで、より構造化された情報をMattermost上に投稿できるようになります。

Message Attachmentsは、今まで紹介してきたIncoming WebHookやSlash Command、今後紹介予定のREST APIやプラグインなど、Mattermostへ投稿を作成する多くの場合で使用することができます。今回は、Incoming WebHookを利用してMessage Attachmentsを作成する例を紹介していきます。

## Message Attachmentsの作成

まず、Incoming WebHook(内向きのウェブフック)のおさらいです。
Incoming WebHookでは、ウェブフックのエンドポイントに下記のようなJSONデータを送信することで、Mattermostに投稿を作成することができました。

```bash
BODY='{
  "text": "Hello, Mattermost"
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

Message Attachmentsを使う場合、送信するJSONデータに`attachment`というフィールドを追加します

```bash
BODY='{
  "text": "Hello, Mattermost",
  "attachments": [{
    "color": "#FF6600",
    "text": "message attachments",
    "title": "This is title",
    "title_link": "https://github.com/mattermost"
  },{
    "color": "#0066FF",
    "text": "message attachments2",
    "title": "This is title",
    "title_link": "https://github.com/mattermost"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

上記のコマンドを実行することで、通常のメッセージの下にMessage Attachmentsのメッセージが表示されます。

![example](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/first-example.png)

`attachments`は複数指定することができ、複数指定した場合は上記のように縦に並んで表示されます。

以降は、`attachmenst`に指定できるオプションと、その効果を紹介していきます。

### Attachments Options

#### fallback

メッセージが作成された際の通知メッセージを指定します。

![use fallback](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-fallback1.png)

```bash
BODY='{
  "attachments": [{
    "fallback": "This is fallback"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

何も指定しない場合はシステムデフォルトのメッセージが表示されます。

![default fallback](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-fallback2.png)

#### color

Message Attachmentsの左側に表示されるラインの色を指定します。

![use color](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-color.png)

```bash
BODY='{
  "attachments": [{
    "color": "#FF6600"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

#### pretext

Message Attachments領域の上に表示されるテキストを指定します。Markdownや絵文字も使用できます。

![use pretext](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-pretext.png)

```bash
BODY='{
  "attachments": [{
    "pretext": "This is `pretext`"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

`attachments`と同時に`text`フィールドを指定した場合、両方のテキストが表示されます。
また、`attachments`を複数指定していた場合、`pretext`はMessage Attachmentsの前にそれぞれ表示されます。

```bash
BODY='{
  "text": "This is `text`",
  "attachments": [{
    "pretext": "This is first `pretext`"
  },{
    "pretext": "This is second `pretext`"
  }]
}'
```

![use pretext](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-pretext2.png)

#### text

Message Attachments領域内に表示されるテキストです。こちらもMarkdownや絵文字を利用可能です。

![use text](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-text.png)

```bash
BODY='{
  "attachments": [{
    "text": "## This is `text` :smiley: \n**Markdown** can be used."
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

### Author Details

#### author_name / author_link / author_icon
Message Attachmentsの作成者情報を指定します。

![use author](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-author.png)

```bash
BODY='{
  "attachments": [{
    "author_name": "kaakaa",
    "author_link": "https://github.com/kaakaa",
    "author_icon": "https://blog.kaakaa.dev/images/avatar.png",
    "text": "Author Details"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

* `author_name`: `attachments`の上部に表示される作成者の名前を指定します
* `author_link`: `attachments`の上部に表示される作成者の名前をクリックしたときのリンクを指定します
* `author_icon`: `attachments`の上部に表示される作成者のアイコン画像のURLを指定します

### Titles

#### title / title_link

![use title](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-title.png)

```bash
BODY='{
  "attachments": [{
    "title": "Attachment Title",
    "title_link": "https://github.com/kaakaa",
    "text": "Titles"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

* `title`: `attachments`の上部に表示されるタイトルを指定します
* `title_link`: `attachments`の上部に表示されるタイトルをクリックしたときのリンクを指定します

Author Detailsと同時に指定した場合は、下記のような見た目になります。

![use title](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-title2.png)

### Fields

#### title / value / short
Message Attachmentsの内容を指定します。

![use fields](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-fields.png)

```bash
BODY='{
  "attachments": [{
    "text": "Fields"
    "fields": [{
      "title": "Short Field :one:",
      "value": "This is first **short** field. This is first **short** field. This is first **short** field. This is first **short** field. This is first **short** field. ",
      "short": true
    }, {
      "title": "Short Field :two:",
      "value": "This is second **short** field. This is second **short** field. This is second **short** field. This is second **short** field. This is second **short** field. ",
      "short": true
    }, {
      "title": "Long Field :one:",
      "value": "This is first **long** field. This is first **long** field. This is first **long** field. This is first **long** field. This is first **long** field. ",
      "short": false
    }, {
      "title": "Long Field :two:",
      "value": "This is second **long** field. This is second **long** field. This is second **long** field. This is second **long** field. This is second **long** field. ",
      "short": false
    }],
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

* `title`: Fieldsのタイトル部分を指定します。絵文字を利用できます。
* `value`: Fieldsの`title`に続く文字を指定します。Markdown形式で記述できます。
* `short`: Fieldの長さをshortにするかどうかを指定します。`"short": true`を指定すると、１行に２つのFieldを表示できます。

### Images

#### image_url / thumb_url

Message Attachmentに表示される画像のURLを指定します.

![use images](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-images.png)

```bash
BODY='{
  "attachments": [{
    "text": "Images",
    "fields": [
      {"title": "Long Field", "value": "This is a long field", "short": false},
      {"title": "Short Field", "value": "This is a short field", "short": true},
      {"title": "Short Field", "value": "This is a short field", "short": true},
      {"title": "Short Field", "value": "This is a short field", "short": true}
    ],
    "image_url": "https://blog.kaakaa.dev/images/avatar.png",
    "thumb_url": "https://blog.kaakaa.dev/images/avatar.png"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

* `image_url`: `text`フィールドと`fields`フィールドの間に表示される画像のURLを指定します。バナーやチャートのような横長の画像を表示するのに適してします。
* `thumb_url`: Message Attachmentsの右上端に表示される画像のURLを指定します。

### Footer

#### footer / footer_link

Message Attachmentsのフッター情報を指定します。

![use footer](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day9/example-footer.png)

```bash
BODY='{
  "attachments": [{
    "text": "Footer",
    "footer": "sample footer",
    "footer_icon": "https://blog.kaakaa.dev/images/avatar.png"
  }]
}'

curl \
  -H "Content-Type: application/json" \
  -d "$BODY"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```


## 既知の問題
2020/11月時点では、下記の制限があります。
* `Footer`フィールドのタイムスタンプ(`ts`)に対応していない
* Message Attachmentの内容はMattermostの検索機能で検索できない

## さいごに

Message Attachmentsの基本的な使い方について紹介しました。
明日は、MattermostのREST APIについて紹介します。
