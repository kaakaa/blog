---
title: "[Mattermost Integrations] Incoming WebHook 基本編"
date: 2020-12-02T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "webhook", "incoming"]
---


Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第2日目の記事です。

本記事では、Mattermostの最も簡易な統合機能であるIncoming WebHook(内向きのウェブフック)機能について紹介します。

## Incoming WebHook

Incoming WebHookは外部アプリケーションからMattermostの投稿を作成するための機能です。

![overview](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/overview.drawio.png)

Incoming WebHookを使うことで、Mattermostの外側で発生したイベントを簡単にMattermost上に通知することができます。

Incoming WebHookに関する公式ドキュメントは下記になります。

* https://docs.mattermost.com/developer/webhooks-incoming.html
* https://developers.mattermost.com/integrate/incoming-webhooks/

一つ目の公式ドキュメントはIncoming WebHookの概要について記述しており、二つ目のDeveloper Documentはより細かい開発者向けの情報が書かれています。

### 設定

Incoming WebHookを利用するには、**システムコンソール > 統合機能管理 > 内向きのウェブフックを有効にする** の設定が`有効`になっている必要があります。

![system console](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/config-incoming-webhook.png)

また、統合機能はデフォルトではシステム管理者とチーム管理者しか作成することができませんが、誰でも作成できるようにしたい場合、**システムコンソール > 統合機能管理 > 統合機能の管理を管理者のみに制限する**の設定を`無効`してください。

### 作成

**メインメニュー > 統合機能**から統合機能の画面を開き、

![main menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/main-menu.png)

**内向きのウェブフック > 内向きのウェブフックを追加する**から、新たなIncoming WebHookを追加します。

![menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/integration-menu.png)

ウェブフックの作成画面で入力する情報は下記の通りです。

![incoming webhook](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/create-incoming-webhook.png)

* **タイトル**: ウェブフックの一覧ページに表示されるタイトルです
* **説明**: ウェブフックの説明です
* **チャンネル**: チャンネル名を指定せずにIncoming WebHookを実行した場合、ここで指定したチャンネルに投稿が作成されます
* **このチャンネルに固定する**: このIncoming WebHookでは、**チャンネル**で指定したチャンネルにしか投稿を作成できなくなります

Incoming WebHookの作成が完了すると、WebHook実行時に指定するURLが表示されます。このURLに対してHTTPリクエストを送信することで、外部アプリケーションからMattermostに投稿を作成することができます。

![complete](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/complete-incoming-webhook.png)

作成したIncoming WebHookは後で編集することも可能です。

![list](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/list-incoming-webhook.png)

### 実行

今回は、`curl`コマンドを使って、定期的にマシンのDISK容量を通知するスクリプトを作ってみます。


```bash
DISK=`df -h /System/Volumes/Data`

curl \
  -H "Content-Type: application/json" \
  -d "{\"text\": \"${DISK}\"}"  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

先ほど生成されたWebHookのURLに`text`フィールドを持つJSONを送信するだけで投稿を作成することができます。

![post](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day2/execute-incoming-webhook.png)

`curl`コマンドを実行するスクリプトを`cron`などで毎日定刻に実行することで、毎日ディスクの状況をチェックしたりすることができます。

## さいごに

Incoming WebHookの基本的な使い方について紹介しました。
明日は、Incoming WebHook実行時のパラメータについて紹介します。