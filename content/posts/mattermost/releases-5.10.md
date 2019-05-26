---
title: "Mattermost5.10の新機能"
date: 2019-04-21T15:17:06+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2019/4/16にMattermost v5.10.0がリリースされました。
この記事では、アップデートの内容について簡単に紹介します。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.10: Single sign\-on for mobile, richer integrations and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-10-single-sign-on-for-mobile-richer-integrations-and-more/)
* [Mattermost Changelog — Mattermost 5\.10 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-10)

---

# アップデート内容

## Interactive ephemeral messages

今まで、Mattermostの[Interactive Messages](https://docs.mattermost.com/developer/interactive-messages.html)(ユーザーが操作可能なボタンやメニューを含む投稿)は、全ユーザーが閲覧可能な形式でしか作成することができませんでした。
今回のバージョンアップにより、自分だけが閲覧可能なInteractive Messagesを作成可能になりました。

これにより、例えば個人単位で機能のOn/Offを切り替えられるプラグインなどで自分だけが見れるメッセージにOn/Off切り替えボタンを表示するなど、よりパーソナライズされたメッセージを作成できるようになります。

この機能はプラグイン用のAPIから使用できるようです。（[Go言語ドライバ](https://github.com/mattermost/mattermost-server/blob/master/model/client4.go)からは作成することができないようでした。）
プラグイン用のサンプルコードは下記になります。

https://gist.github.com/kaakaa/b040bcd0a6fdccc18b3daaea214c6112#file-mm_v5-10_ephemeral_interactive_message-go

![スクリーンショット 2019-04-21 14.44.54.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/c2349a35-f7ab-d23b-f7e5-04dc0cc9d13b.png)


## (実験的な機能) Mattermost設定のデータベース化

最近、MattermostではKubernetes環境向けのデプロイやコンテナ化に注力しており、そのために、サーバー間で安定した設定のやり取りを行う方法が求められています。今まで、MattermostのサーバはJSONファイルによって管理されてきましたが、このバージョンからMattermostの設定をデータベースに格納することができるようになりました。

まだドキュメント化はされていないようですが、`mattermost config migrate`コマンドで設定できるのだと思われます。
https://github.com/mattermost/mattermost-server/blob/master/cmd/mattermost/commands/config.go#L71

ドキュメント化されてないことからも分かるように、この機能はまだ実験的な機能なため、強い要望が無い限りは使用しないほうが良いと思います。

## その他の変更

### Message Attachmentのタイトルに絵文字/ハイパーリンクが利用可能に


[Message Attachments](https://docs.mattermost.com/developer/message-attachments.html)のタイトル部分に絵文字とリンクを使用できるようになりました。
また、Message Attachmentsのボタンには絵文字を使用できるようになりました。

![スクリーンショット 2019-04-21 22.33.57.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/c936c847-2a86-181b-7738-c3d297eb442b.png)

上記投稿を作るサンプルコードは下記になります。
https://gist.github.com/kaakaa/b040bcd0a6fdccc18b3daaea214c6112#file-mm_v5-10_decoration_message_button-go

Message Buttonでは絵文字のみ利用可能に。

### メッセージ入力ボックスにチャンネル名が表示されるように

メッセージ入力ボックスに何も入力していない状態の時に表示されるメッセージに、チャンネル名が含まれるようになりました。

![スクリーンショット 2019-04-21 22.37.20.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/8a45d410-c0ed-1730-8e41-8eea2d7fcbe1.png)


### Excel から Markdownテーブルへの変換

Excelのセルをコピーし、そのままMatterostのメッセージ入力欄に貼り付けるとMarkdownテーブル形式に変換して貼り付けてくれる機能が追加されました。
Excelだけでなく、Googleスプレッドシートからでも同様のことが行えます。

![image.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/5fcdb63a-6208-e967-780b-77af821491c7.png)

### 実験的機能がシステムコンソールから設定可能に

**システムコンソール** > **詳細** > **実験的機能**から、実験的な機能の設定を行えるようになりました。
細かな設定含めて多くの項目がありますが、**サイドバー構成**を有効にすることで使えるようになる、「未読チャンネルのグループ化」などは便利です。

ただ、「実験的な機能」とあるように、動作が保証されている機能ではない点に注意が必要です。

## モバイルアプリのアップデート
v5.10.0のリリースと同日の4/16にMattermostのモバイルアプリの最新版v1.18.0もリリースされました。
主な変更内容は下記になります。

* Office365 SSO / Integrated Windows Authentication(IWA) のサポート
* [Message Attachment](https://docs.mattermost.com/developer/message-attachments.html)のタイトルに絵文字やハイパーリンクが使用可能に
* 投稿を長押しした時のメニューに「返信する」が追加

詳細な変更内容については[mattermost\-mobile/CHANGELOG\.md at master · mattermost/mattermost\-mobile](https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#1180-release)を参照してください。

* * *

# その他
## Mattermost Plugin Bounty Program

要望の多いMattermostプラグインについて、報奨金を出してコミュニティに開発してもらおうというMattermost Plugin Bounty Programが公開されています。（４つの募集プラグインについて、すでにすべてコミュニティメンバがアサインされている模様）

https://forum.mattermost.org/t/mattermost-plugin-bounty-program/6857

募集されているのは下記4つのサービス/OSSのプラグインです。

* Cisco WebEx Meetings Server
* Skype for Business
* Jenkins
* GitLab

開発したプラグインをGitHubにApache 2.0 or MITライセンスで公開することが条件となっているため、これらのプラグインは誰でも利用できる形になるかと思います。

## Write a Review on Mattermost

「Mattermostについてのレビューを送って$20のギフトカードをもらおう」という旨のキャンペーンが行われています。
（現在でも受け付けているのか、日本でもギフトカードを受け取れるのか、等の詳細については不明です。すいません。）

[Capterra \| Review Software](https://review.capterra.com/SS-Mattermost-170524?utm_source=twitter&utm_medium=organic&utm_campaign=reviews)

## Software Design
3/18発売のSoftware Deisign4月号から、4か月間の短期連載でMattermostに関する紹介記事を書かせていただいています。
書店などで見かけたらぜひ手に取ってみてください。

https://gihyo.jp/magazine/SD


# おわりに

次の`v5.11`のリリースは2019/5/16(Fri)を予定しています。
そして機能追加が行われる`v5.12`は恐らく2019/6/16(Tue)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。

