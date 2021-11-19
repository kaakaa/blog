---
title: "Mattermost 6.1の新機能"
emoji: "🎉"
published: true
date: 2021-11-19T17:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2021/11/16 に Mattermost のメジャーバージョンアップである `v6.1.0` がリリースされました。  
また、10/18にMedium Level Security Fixを含む `v6.0.1` がリリースされています。今後、v6へアップグレードする場合は `v6.0.1` を使用してください。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.1 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v6-1-is-now-available/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-1-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/install/self-managed-changelog.html#important-upgrade-notes)


* v6.1へのアップグレード処理でデータベーススキーマの変更が行われます。v6へのメジャーアップデートほど大掛かりではないですが、多少時間がかかります。
  * スキーマ変更の内容と、所要時間の目安については[Mattermost v6\.1\.0 schema migration analysis](https://gist.github.com/streamer45/997b726a86b5d2a624ac2af435a66086)に記述があります
* オプションの検索エンジンとして搭載されているBleveの検索インデックスがScorch indexに変更されます。Scorch indexを採用することにより、検索にまつわるディスク使用量の削減を行うことができます。
  * Mattermostサーバーのアップグレード後、システムコンソールからBleveインデックスの破棄と再生成を行うことでScorch indexに切り替えることができます
  * 過去のバージョンで作成された検索インデックスでも動作はするため、インデックス種別を切り替えたくない場合、何も操作は必要ありません

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [`Professional`](https://mattermost.com/pricing/)
- [`Enterprise`](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(無償版)でも利用可能な機能です。

## Channels: チーム横断でのメンション検索

あなたへのメンションを含む投稿や、保存された投稿を一覧する機能は以前よりありましたが、Mattermost上のチームごとに管理されており、別のチームで行われたメンションや、保存された投稿を確認するには、まずチームの切り替えを行う必要がありました。本バージョンから、所属している全チームのメンションと保存された投稿が一つの画面に表示されるようになったため、チーム切り替えを行う必要がなくなりました。

![ch-cross-team](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-cross-team.png)

## Channels: リアクションへのクイックアクセス

投稿メニューに最近使った絵文字が表示されるようになり、1-clickで簡単にリアクションをつけられるようになりました。

![ch-recent-reaction](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-recent-reaction.png)

**設定 > 表示 > 1クリックでメッセージに反応する**から、この機能のON/OFFをユーザーごとに設定することができます。

![ch-recent-reaction-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-recent-reaction-setting.png)

## Channels: `取り込み中`状態の有効期間設定

Mattermostでは、オンラインや離籍中など、ユーザーの現在の状態を設定することができますが、この状態表示のうち`取り込み中`を選択する際、`取り込み中`が継続する期間を設定することができるようになりました。
`取り込み中`を設定した場合、Mattermostからの通知が送信されなくなりますが、これが指定時間経過後に自動で解除され、通知が送信されるようになります。

![ch-do-not-disturb](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-do-not-disturb.png)

## Playbooks: プレビュー表示

本バージョンからPlaybook作成時に指定した内容をプレビュー表示できるようになりました。これにより、Playbookの概要を簡単に把握できるようになります。

![pl-preview](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/pl-preview.png)


## Playbooks: TODO通知

Playbookの参加者とオーナーへ、割り当てられているタスクの一覧や期限超過の状態などを含む通知メッセージが届くようになりました。

![pl-todo](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/pl-todo.webp)
(画像は[公式エントリ](https://mattermost.com/blog/mattermost-v6-1-is-now-available/)から)

`/playbook`コマンドを使い、通知をオフにすることや、任意のタイミングでTODOメッセージを受信することもできます。`/playbook`コマンドについては以下のドキュメントを参照してください。
[Notifications and Updates](https://docs.mattermost.com/playbooks/notifications-and-updates.html)

## Boards: ボード作成UIの改善

ボード作成ボタンがサイドバーの上部に移動しました（今まではサイドバー下部だった）。また、ボード作成画面も全画面形式に変更されています。

![bo-create-board](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/bo-create-board.png)

## Boards: メトリクス選択

以前のバージョンでは、Board内のカラム名の横にはそのカラムに属するカードの数が表示されるだけでしたが、この数値の集計方法を以下のように指定できるようになりました。

* 特定のプロパティが空のカードの数を表示する (重要な情報が未入力なカードを見つけやすくする)
* 機能の開発にかかる時間の合計を表示する
* 見積もり期間の範囲を表示する

![bo-calculate](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/bo-calculate.png)

Sum, Averageなどのメニューは、カードに数値のプロパティが存在する場合のみ表示・選択できるようになります。

## Boards: @mention notification

カードのコメントや説明部分で@メンションが行われると、その内容がChannels内にBoard BotからのDMとして通知されるようになりました。

![bo-mention](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/bo-mention.png)

## Board: カードプレビュー

カードへのリンクをチャット部分に投稿すると、その内容をプレビュー展開できるようになりました。プレビュー部分をクリックすることで、Boardが開き、該当のカードの詳細を確認することができます。


![bo-card-preview](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/bo-card-preview.png)

---

## その他の変更

### 編集済マークの表示位置変更

投稿後にメッセージ内容を編集した際に表示される`編集済`マークの表示位置が変更され、メッセージの最後に表示されるようになりました。

![ch-edit-indicator](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-edit-indicator.png)

### インラインLaTeX

今までもコードブロックとしてLaTeX書式を記述できましたが、本バージョンから文中にもLaTeX書式を記述できるようになりました。`$`マーク内にLaTeX形式の数式を記述することで、数式として表示できます。

![ch-inline-latex](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.1/ch-inline-latex.png)

詳しくは以下のドキュメントを参照ください。
[Formatting Text - Math Formulas](https://docs.mattermost.com/messaging/formatting-text.html#math-formulas)

### ブラウザサポートバージョンの変更
サポートされているブラウザ/OSの最低バージョンが更新されています。

* Chrome updated from `61+` to `89+`.
* Firefox updated from `60+` to `78+`.
* MacOS updated from `10.9+` to `10.14+`.

## その他のトピック

### Mattermost Cloud
10ユーザー以内なら無料として利用できたMattermost Cloudですが、2022/2/15に無償プランが廃止され、有償のStarter plan(月額$12.41 or 年額 $149)への移行が必要となるようです。
2022/2/15までにクレジットカードの情報を入力するか、何もアクションがない場合は利用継続できなくなってしまうようです。

### Hacktoberfest

[Join Mattermost for Hacktoberfest 2021](https://mattermost.com/blog/hacktoberfest-2021/)

今年もMattermostはHacktoberfestのPartnerとして参加しています。 

https://hacktoberfest.digitalocean.com/

Hacktoberfestの景品とは別に、Mattermost独自のSwagを用意しています。
10月中に一つでもコントリビュートを行なった場合は、限定ステッカーがもらえるようです。このステッカーは、日本のJapanese Maple Treeをモチーフにしているそうです。

![Fall Sticker](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/fall_sticker.jpeg)

さらに、各カテゴリ(**Focalboard**, **Mattermost App**, **Writing Program**, **Translation**, **QA Testing**)でトップコントリビューターに選ばれた場合、オリジナルのメカニカルキーボードがプレゼントされるようです。

## おわりに

次の`v6.2`のリリースは 2021/12/16(Thu)を予定しています。
