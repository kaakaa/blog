---
title: "Mattermost 7.5の新機能"
emoji: "🎉"
published: true
date: 2022-11-19T10:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/11/16 に Mattermost のアップデートとなる `v7.5.0` がリリースされました。
また、同日にUbuntu 18.04環境でアップグレード時に発生する問題の修正を含む `v7.5.1`がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.5 is now available \- Upgrade your deployment today](https://mattermost.com/blog/mattermost-v7-5-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-5-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- Platform: 上記機能を含むMattermost全体

## Calls: メッセージスレッド

Calls機能を利用中に、スレッド形式でメッセージをやり取りできるようになりました。  
Callsのメッセージスレッドでは、通常のメッセージスレッドと同様に絵文字リアクションや@メンションを利用することができます。

メッセージスレッドを利用するには、Callへ参加し、Windowsモードにし、コメントアイコンをクリックすることで、メッセージスレッドを開くことができます。

![calls-message-open](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/calls-message-open.png)

このスレッドでのやり取りは、通常のメッセージスレッドと同様、チャンネルに残ります。

![calls-message-channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/calls-message-channel.png)

---

[Mattermost Cloud](https://mattermost.com/sign-up/)の方でCallsプラグインを [v0.10.0](https://github.com/mattermost/mattermost-plugin-calls/releases/tag/v0.10.0)にアップデートしたところ、Webブラウザ/Decksopアプリからはコメントアイコンが表示されず、モバイルアプリ v2 (Beta)の方では `チャットスレッド` のメニューが表示されるという状態でした。

一方、公式の[Community向けMattermost](https://community.mattermost.com)の方で試したみたところ、Webブラウザ/Desktopアプリではコメントアイコンが表示され、モバイルアプリの方では `チャットスレッド` のメニューが現れないという状態でした。

一応ブラウザのキャッシュ削除等は試してみたのですが解消せず、ちょっと原因がわかりませんでした。


## Boards: ボードテンプレートの追加

ボードを新規に作成する際に選択できるテンプレートに、新たに以下の6つが追加されました。

![boards-new-templates](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-new-templates.png)

**Company goals & OKRs**

組織のゴールや、OKRの管理を行うためのテンプレートです。

![boards-template-okr](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-okr.png)

Objective(目標)ごとにタスク(カード)を作成し、Target/Actual/Competion等のプロパティで達成度合いを管理することができます。また、QuarterやDepartment(部署)ごとにタスクを一覧化できるビューもデフォルトで作成されます。

**Competitive analysis**

競合分析用のテンプレートです。

![boards-template-competitive](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-competitive.png)

競合企業の情報 (Website, Location, Revenue,...) や各企業のMarket Positionなどを一覧化することができます。Market Positionごとにカードを並べたビューもデフォルトで作成されます。

**Sales pipeline CRM**

セールスの状況を管理するためのテンプレートです。

![boards-template-sales](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-sales.png)

各顧客に関する情報や商談の状況などを一覧化することができます。商談の状況(Open Deals)や商談フェーズ(Pipeline Tracker)ごとにタスクを一覧化できるビューもデフォルトで作成されます。

**Sprint planner**

スプリント計画やリリースに向けたタスク管理を行うためのテンプレートです。

![boards-template-sprint](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-sprint.png)

スプリントごとのタスク一覧を管理することができます。タスクの状況(Status)やタスクの種別(Type)ごとにタスクを一覧化できるビューもデフォルトで作成されます。

**Team retrospective**

うまく行ったことや改善点など、将来に向けた振り返りを行うためのテンプレートです。

![boards-template-retrospective](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-retrospective.png)

**User research sessions**

ユーザーリサーチの状況を管理するためのテンプレートです。

![boards-template-user-research](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-template-user-research.png)

ユーザーへのインタビュー状況を一覧で管理することができます。インタビュー予定(Date)や状況(Status)ごとにタスクを一覧化するビューもデフォルトで作成されます。

## Boards: テキストプロパティでのフィルタリング

ボード内のカードをフィルタリングをする際に、以下のプロパティの値でフィルタリングできるようになりました。

* Text
* Email
* Phone
* URL

![boards-text-filtering](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-text-filtering.png)

## Boards: ボードの統計情報

**システムコンソール > サイトの使用統計**に、ボードに関する統計情報が表示されるようになりました。ボード数とカード数が表示されるようになります。

![boards-statistics](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/boards-statistics.png)

Mattermost CloudのFreeプランではカード数の上限が **10,000** までに設定されているため、Freeプランの上限を超えないためにカード数の総数を把握しやすくなったのは良いことだと思います。

## Channels: 最終活動日時の表示

ユーザープロフィールを表示するポップアップに、そのユーザーが最後に活動した時間が表示されるようになりました。
ステータスがアクティブでないユーザーのみ表示されます。

![channels-last-activity](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/channels-last-activity.png)

この機能は **システムコンソール > サイト設定 > ユーザーとチーム > 最終活動日時を有効にする** から無効化することもできます。

![channels-last-activity-console](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/channels-last-activity-console.png)


また、**設定 > 表示 > 最終活動日時を共有する**から、ユーザーごとにこの情報を表示しないよう設定することもできます。

![channels-last-activity-settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/channels-last-activity-settings.png)


## Platform: Desktop v5.2

Mattermost Desktopアプリv5.2 がリリースされました。

### ファイルダウンロード状況の表示

ファイルダウンロードの状況を表示するダイアログが追加されています。

![desktop-download](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/desktop-download.png)

### メッセージの翻訳

デスクトップアプリのメニュー等がローカライズされるようになりました。日本語にも対応しています。

![desktop-localize](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/desktop-localize.png)

## Platform: Mobile App v1.55.1 (ESR)

9月にリリースされたv1.55.1が、モバイルアプリとしては初めてのESR(Extended Support Release)として設定されました。このバージョンは2023/06までサポートされます。  
モバイルアプリのメジャーバージョンアップであるv2.0が12月にリリース予定となりますが、モバイルアプリのカスタムビルドを利用しているユーザーのための経過措置になります。


## アップグレード時の注意事項

### DBマイグレーション

`Posts`テーブル内の`ParentID`カラムを削除するDBマイグレーションが実行されますが、MySQLを利用している場合、`Posts`テーブルのサイズによってはこの処理によってCPU使用率が上昇する場合があります。
また、このマイグレーション実行中は書き込みが制限されます。

### Webappプラグイン

[`PluginRegistry.registerCustomRoute`](https://developers.mattermost.com/integrate/plugins/components/webapp/reference/#registerCustomRoute)を利用したプラグイン開発を行なっている場合、CSS`grid-area: center`が設定されていないとコンポーネントが適正な位置に表示されなくなります。

## その他の変更

### `Starter`プランの名称が`Free`に変更

今まで、Mattermostの無償プランは`Starter`プランと呼ばれていましたが、これからは`Free`プランと呼ぶようになりました。

### 未読チャンネルをフィルタするためのショートカットが追加

未読のチャンネルのみを表示するためショートカットとして `Ctrl/Cmd + Shift + U`が追加されました。

![channels-shortcut-unreads](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.5/channels-shortcut-unreads.png)

## その他のトピック
### Mattermost関連の記事紹介
  
#### Mattermost構築
* [【Tool】Mattermostのセルフホスティングを導入してみる \| willserver for tech\-future](https://tech.willserver.asia/2022/10/21/tool-mattermost_self_setup/)
  * AWS LightsailにMattermostを構築する手順について
* [【Tool】Mattermostクライアントインストール手順\(Windows、macOS、Linux\) \| willserver for tech\-future](https://tech.willserver.asia/2022/10/24/tool-mattermost-client/)
  * Mattermost Desktopアプリ(Windows/Linux/macOS) のインストール方法について
* [Slackの投稿データをMattermostに移行してみる \- Qiita](https://qiita.com/showchan33/items/e9a3860b1d1becdca0ed)
  * SlackのデータをMattermostへ移行する方法について

#### Mattermostの紹介
* [Mattermost〜OSSのチャットツール〜 \| OSSのデージーネット](https://www.designet.co.jp/ossinfo/mattermost/)
  * Mattermostの概要・機能紹介

## おわりに
次の`v7.6`のリリースは 2022/12/15(Thu)を予定しています。