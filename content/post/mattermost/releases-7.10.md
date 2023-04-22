---
title: "Mattermost 7.10の新機能"
emoji: "🎉"
published: true
date: 2023-04-22T13:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/04/14 に Mattermost のアップデートとなる `v7.10.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- (今回のリリースに関する公式ブログでの紹介記事は無いそうです)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-10-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体

## (Channels) 投稿のリマインド

Mattermost上の投稿を一定期間後にリマインドする機能が追加されました。  
メッセージの内容を忘れずに後で確認したい時や、土日を挟んだ翌月曜向けのタスクを設定しておきたい時などに便利な機能です。

リマインドしたい投稿のメニューにある **Remind** からリマインドしたい時刻を選択することで、リマインドを設定することができます。

![channels-remind-menu](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.10/channels-remind-menu.png)

リマインド時刻を指定すると、以下のようなリマインドが実行される時刻を含む予告投稿が行われます。

![channels-remind-notice](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.10/channels-remind-notice.png)

指定した時刻になると、`system-bot`という名前のBotからのダイレクトメッセージとして以下のようなリマインドが送信されます。

![channels-remind-post](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.10/channels-remind-post.png)

リマインド時刻は `30分後`、`1時間後`、`2時間後`、`明日`、`カスタム`から選択することができ、`カスタム`を選択すると、30分単位で任意の日時を指定することができます。

![channels-remind-custom](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.10/channels-remind-custom.png)

## (Calls) 通話画面を別ウィンドウで開けるように (Desktop)

2023/03/30にリリースされた[Mattermost Desktop App](https://mattermost.com/apps/) v5.3から、Calls画面を別ウィンドウとして開くことができるようになりました。  
これによりCalls画面を開きながら、いつも通りMattermostでチャットができるようになります。

![calls-window](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.10/calls-window.png)

## その他のトピック

### Mattermostリポジトリのモノレポ化

Mattermostのソースコードはサーバーサイドの [mattermost-server](https://github.com/mattermost/mattermost-server) とフロントエンドの [mattermost-webapp](https://github.com/mattermost/mattermost-webapp) に分けて管理されていましたが、次回のv7.11リリースから [mattermost-server](https://github.com/mattermost/mattermost-server) に集約され、モノレポ管理となります。  
また、[Playbooks](https://github.com/mattermost/mattermost-plugin-playbooks)機能と[Boards](https://github.com/mattermost/focalboard)機能も、今まではMattermost Pluginとして開発されていましたが、これらも[mattermost-server](https://github.com/mattermost/mattermost-server)に集約され、Mattermostのコア機能の一つとして開発が進められることになります。

Mattermostのコア機能に関するリポジトリがモノレポ化されたことで、開発者向けのセットアップ手順が変更されています。 [Developer setup](https://developers.mattermost.com/contribute/developer-setup/)

### ChatGPT Bot

最近話題のChatGPTをMattermostから使う方法について、Mattermost公式ブログで紹介がありました。  
[Community Spotlight: Mattermost ChatGPT bot by Sebastian Müller](https://mattermost.com/blog/community-spotlight-mattermost-chatgpt-bot-by-sebastian-muller/)

{{< youtube Hx4Ex7YZZiA >}}

これはコミュニティメンバによって作られたツールで、Mattermostのメッセージを受け取ってOpenAI APIとやり取りをし、結果をMattermostに投稿するBotサーバーを構築するものです。  
[yGuy/chatgpt\-mattermost\-bot: A very simple implementation of a service for a mattermost bot that uses ChatGPT in the backend\.](https://github.com/yGuy/chatgpt-mattermost-bot)

この他にも、MattermostとChatGPTを連携したという日本語の記事もいくつか公開されています。

* [お手軽！ Mattermost に ChatGPT Bot を実装してみた！ \- ディーメイクブログ](https://www.d-make.co.jp/blog/2023/03/06/%E3%81%8A%E6%89%8B%E8%BB%BD%EF%BC%81-mattermost-%E3%81%AB-chatgpt-bot-%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%97%E3%81%A6%E3%81%BF%E3%81%9F%EF%BC%81/)
* [ChatGPTにRedmineの起票を任せてみた \- Qiita](https://qiita.com/IShun/items/8fb2501c8ae6388798bb)

## おわりに
次の`v7.11`のリリースは 2023/05/16(Tue)を予定しています。
