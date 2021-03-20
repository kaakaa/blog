---
title: "[Mattermost Integrations] REST API"
date: 2020-12-10T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "api", "rest"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第10日目の記事です。

本記事では、MattermostのREST APIについて紹介します。

## REST APIの概要
Mattermostには統合機能を作成するための様々な機能がありますが、一般的なREST APIを利用することもできます。
Mattermostで利用可能なREST APIは、下記の公式ドキュメントにまとまっています。

https://api.mattermost.com/

## Access Token (Personal Access Token)

REST APIを実行するには、API実行権限を保持していることを示すためのTokenが必要になります。MattermostではいくつかTokenを取得する方法がありますが、ここでは最も手軽に使えるPersonal Access Tokenを生成する方法について紹介します。

Personal Access Tokenを利用するには、**システムコンソール > 統合機能 > 統合機能管理 > パーソナルアクセストークンを有効にする**が有効になっている必要があります。

![config pat](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/config-pat.png)

### Personal Access Token生成

**メインメニュー > アカウントの設定 > セキュリティ > パーソナルアクセストークン > トークンを生成する**を実行し、生成するトークンの説明を入力することでPersonal Access Tokenを生成できます。

![create pat](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/create-pat.png)

Personal Access Tokenの生成に成功すると、画面にトークンが表示されます。ここで**アクセストークン**と表示されている文字列を使ってREST APIの呼び出しを行います。(**トークンID**は、Mattermost内部でトークンを一意に識別するためのIDであり、REST APIの実行に利用することはできません)

![complete pat](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/complete-pat.png)

### REST API実行

先ほど生成したトークンをAuthorizationヘッダーに指定し、`/api/v4/users/me`のAPIを実行することで、ユーザーの情報を取得するAPIを実行することができます。

```bash
$ curl -i \
>   -H 'Authorization: Bearer 4cacdozwn3fndnzobbpha3nnhy' \
>   http://localhost:8065/api/v4/users/me
HTTP/1.1 200 OK
Content-Type: application/json
Etag: 5.30.0.87x93uo8pfnzdro9ktcmobpa1r.1606000641083..0.true.true.0
Expires: 0
Vary: Accept-Encoding
X-Request-Id: qws4oo9dyjy8pxs48kaqd86goh
X-Version-Id: 5.30.0.dev.3fbba2b2e9e21536fd2bdc547afe31fd.false
Date: Sun, 22 Nov 2020 05:33:58 GMT
Content-Length: 774

{"id":"87x93uo8pfnzdro9ktcmobpa1r","create_at":1598680540414,"update_at":1606000641083,"delete_at":0,"username":"kaakaa","auth_data":"","auth_service":"","email":"kaakaa@example.com","nickname":"","first_name":"","last_name":"","position":"","roles":"system_user system_admin","allow_marketing":true,"notify_props":{"auto_responder_active":"false","auto_responder_message":"ただいま外出中のため返信できません。","channel":"true","comments":"never","desktop":"all","desktop_notification_sound":"Bing","desktop_sound":"true","email":"true","first_name":"false","mention_keys":"","push":"mention","push_status":"away"},"last_password_update":1598680540414,"locale":"ja","timezone":{"automaticTimezone":"","manualTimezone":"","useAutomaticTimezone":"true"}}
```

## REST API実行

REST APIは種類が多いため、一部のみ実行方法を紹介します。

### REST APIから投稿を作成する

REST APIを使ってMattermostに投稿を作成するには`/posts`を利用します。また、リクエストパラメータはJSON形式のため`Content-Type`ヘッダーに`application/json`を指定し、下記のようにAPIを実行することで投稿を作成できます。

```bash
BODY='{
  "channel_id": "uoxmk8819pyftybx6zqkij37ce",
  "message": "Create post by REST API"
}'

curl -i \
  -H 'Authorization: Bearer 4cacdozwn3fndnzobbpha3nnhy' \
  -H 'Content-Type: application/json' \
  -d "$BODY" \
  http://localhost:8065/api/v4/posts
```

![use rest api](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/example-rest.png)

以下のように`props`フィールドに`attachments`を指定することでMessage Attachmentsを利用することもできます。

```bash
BODY='{
  "channel_id": "uoxmk8819pyftybx6zqkij37ce",
  "message": "Create post by REST API",
  "props": {
    "attachments": [{
      "text": "hoge"
    }]
  }
}'
```

![use rest api](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/example-rest2.png)

### 統計情報を取得する

`/analytics/old`ではMattermostの統計情報を取得することができます。

```bash
$ curl \
>   -H 'Authorization: Bearer 4cacdozwn3fndnzobbpha3nnhy' \
>   http://localhost:8065/api/v4/analytics/old
[
  {
    "name": "channel_open_count",
    "value": 9
  },
  {
    "name": "channel_private_count",
    "value": 1
  },
  {
    "name": "post_count",
    "value": 3696
  },
  {
    "name": "unique_user_count",
    "value": 3
  },
  {
    "name": "team_count",
    "value": 1
  },
  {
    "name": "total_websocket_connections",
    "value": 3
  },
  {
    "name": "total_master_db_connections",
    "value": 7
  },
  {
    "name": "total_read_db_connections",
    "value": 0
  },
  {
    "name": "daily_active_users",
    "value": 1
  },
  {
    "name": "monthly_active_users",
    "value": 2
  },
  {
    "name": "inactive_user_count",
    "value": 0
  }
]
```

nameクエリにより取得する統計情報を指定することもできます。例えば、`?name=post_counts_day`では、ここ一週間の日ごとの投稿数を取得することができます。（取得する期間の指定などには対応していないようです）

```bash
$ curl \
>   -H 'Authorization: Bearer 4cacdozwn3fndnzobbpha3nnhy' \
>   http://localhost:8065/api/v4/analytics/old?name=post_counts_day
[
  {
    "name": "2020-11-19",
    "value": 11
  },
  {
    "name": "2020-11-07",
    "value": 15
  },
  {
    "name": "2020-11-04",
    "value": 29
  },
  {
    "name": "2020-11-03",
    "value": 41
  },
  {
    "name": "2020-11-02",
    "value": 5
  },
  {
    "name": "2020-11-01",
    "value": 17
  },
  {
    "name": "2020-10-31",
    "value": 25
  }
]
```

## その他

### API実行回数の制限
大量のAPIリクエストの処理による高負荷状態を避けるために、REST APIの頻度制限(Rate Limit)をかけることもできます。**システムコンソール > 環境 > 投稿頻度制限**から設定を行えます。

![config rate limit](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day10/config-rate-limit.png)

### 各言語向けDrivers

Mattermost REST APIには各言語向けのDriverが存在します。GoとJavascriptについては、Mattermostが公式にサポートしています。
https://api.mattermost.com/#tag/drivers

* Go(公式): https://github.com/mattermost/mattermost-server/blob/master/model/client4.go
* Javascript(公式): https://github.com/mattermost/mattermost-redux/blob/master/src/client/client4.ts
* Python: https://github.com/Vaelor/python-mattermost-driver
* PHP: https://github.com/gnello/php-mattermost-driver
* Java: https://github.com/maruTA-bis5/mattermost4j

## さいごに

Mattermost REST APIの使い方について紹介しました。
明日は、WebSocket APIの使い方を紹介します。
