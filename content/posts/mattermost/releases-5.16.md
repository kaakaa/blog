---
title: "Mattermost5.16の新機能"
date: 2019-10-19T10:15:49+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2019/10/16にMattermost v5.16.0がリリースされました。
既にv5.16.1のリリースが予定されています。v5.16.1に含まれる修正内容については[Changelog](https://docs.mattermost.com/administration/changelog.html#release-v5-16-feature-release)を確認ください。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.16: Guest accounts, a new Plugin Marketplace, faster installation on desktops, and more \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-16-guest-accounts-a-new-plugin-marketplace-faster-installation-on-desktops-and-more/)
* [Mattermost Changelog — Mattermost 5\.16 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-16-feature-release)


---

# アップデート内容


## (E10/20) ゲストアカウント

Enterprise版限定の機能です。
組織外のユーザーとのコラボレーションのために、限定的なアクセス権のみを持つゲストアカウントを作成できるようになりました。

ゲストアカウントは、招待されたチャンネルのみにアクセスでき、その他の公開チャンネルを検索したり、他のチャンネルに参加したりすることはできません。また、ダイレクトメッセージについても、招待されたチャンネルに参加しているユーザーのみにしか送信できず、その他のユーザーにはメッセージを送ることはできません。

![guest_account](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/guest_account_01.png)

ゲストアカウントの投稿や、プロフィールポップアップ画面には `GUEST` というバッチが付与されます。(画像は[公式ブログ](https://mattermost.com/blog/mattermost-5-16-guest-accounts-a-new-plugin-marketplace-faster-installation-on-desktops-and-more/)から)

### ゲストアカウントの設定・招待

ゲストアカウントはシステムコンソール画面から有効にできます。また、ゲストアカウントとして追加できるメールアドレスのドメインを制限することもできます。

![configuration](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/guest_systemconsole_01.png)


ゲストアカウント機能を有効にしたのち、メインメニューの **招待する** メニューをクリックすると、下記のようなメニューが表示されます。
ゲストアカウントを追加したい場合は、「ゲストを招待する」メニューを選択します。

![invite](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/guest_invite_01.png)

既に他のチームに参加しているユーザーをゲストアカウントとして追加したい場合はユーザー名を、まだMattermostアカウントを作成していない場合はメールアドレスを指定します。
ゲストとして追加するチャンネルは、公開チャンネル・非公開チャンネル共に選択することができます。

![invite](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/guest_invite_02.png)

メールアドレスを入力してゲストを招待すると、このようなメールが送信されます。

![invite](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/guest_invite_03.png)

また、既に存在するアカウントをゲストアカウントへ降格させることもできます。

![demote](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/demote_guest_01.png)

## プラグインマーケットプレース
Mattermost統合機能のマーケットプレースにメインメニューからアクセスできるようになりました。
これによりシステム管理者は、現在使用しているバージョンのMattermostで利用可能なプラグインを簡単にインストールできるようになります。

メインメニューの **プラグインマーケットプレース** メニューからマーケットプレースを開けます。
プラグインマーケットプレースから、インストールしたいプラグインの横にある **インストール** ボタンをクリックするだけでプラグインをインストールできます。

![plugin_marketplace](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/plugin_marketplace_02.png)


現在は、Mattermostチームによって検証されたプラグインのみが表示されます。コミュニティメンバによって開発されたプラグインは https://integrations.mattermost.com/ から探すことができます。

また、このマーケットプレース用のサーバーをセルフホストすることもできます。

https://github.com/mattermost/mattermost-marketplace

上記リポジトリをcloneし、`plugins.json`を編集して `make run-server` を実行するだけで、独自のマーケットプレースサーバーを構築できます。（ただ、ローカルマシンで試してみたところ、プラグインのインストールが完了してもインストール中の表示が消えなかったりしたので、実際に利用する場合は注意してください）

マーケットプレースサーバーをセルフホストした場合、そのサーバーURLをシステムコンソールから設定することにより、セルフホストしたマーケットプレースに接続できるようになります。

![plugin_marketplace](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/plugin_marketplace_03.png)


## その他の変更点

### Internet Explorer 11のサポート終了

v5.16.0より Internet Explorer 11(IE11) のサポートが終了しました。今後、Internet Explorer での Mattermost の動作は保証されません。
この決定の背景等については[フォーラムの投稿](https://forum.mattermost.org/t/mattermost-is-dropping-support-for-internet-explorer-ie11-in-v5-16/7575)にまとめられています。

今後は、IEの代わりに下記のクライアントを使用することが推奨されています。

> * Mattermost Desktop App for Windows, Mac and Linux
> * Chrome for Windows, Mac and Linux
> * Firefox for Windows, Mac and Linux
> * Safari for Mac, usually available on all Macs by default
> * Microsoft Edge for Windows, usually available on all Windows 10 machines by default

Mattermostがサポートしているクライアントのバージョンについては https://docs.mattermost.com/install/requirements.html#pc-web で確認できます。

### MSIインストーラーによるDesktopアプリインストールの効率化

Mattermost Desktop v4.3.0から、MSIインストーラーによるMattermost Desktopアプリのインストールが可能になりました。
また、Windowsのグループポリシーを設定することで、下記を制御することが可能になりました。

* アプリの自動アップデート機能の有効/無効
* サーバーの追加/削除機能の有効/無効
* デフォルトサーバーリストの有効/無効

詳しくは https://docs.mattermost.com/install/desktop-msi-gpo.html を参照してください。

### Mattermost Desktopアプリのデスクトップ通知方式変更
[Mattermost Desktopアプリ v4.3.0](https://github.com/mattermost/desktop/releases/tag/v4.3.0)から、セキュアでない`http`プロトコルからのデスクトップ通知の送信方式が変更されました。
Mattermost Desktopアプリ v4.3.0を使用している場合は、Mattermostサーバーを v5.16.0+, v5.15.1, v5.14.4, v5.9.5のいずれかのアップデートする必要があります。

### インタラクティブダイアログの改善

Mattermostの[インタラクティブダイアログ](https://docs.mattermost.com/developer/interactive-dialogs.html)で使用できる要素が追加されました。

* ダイアログの上部にMarkdown形式の説明文を追加することができるようになりました
* 画面上で表示されないパスワード型のテキストを使用できるようになりました
* チェックボックス型の要素を利用できるようになりました
* ラジオボタンを利用できるようになりました

![interactive_dialog](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.16/interactive_dialog_01.png)

上記のダイアログを生成するためのコードは https://gist.github.com/kaakaa/a159cdc51a96b2daa41d84b2d1e2c218#file-mm_v5-16_interactive_dialog-go-L34 にあります。

## misc

### Hacktoberfest

[Hacktoberfest returns\! Contribute to open source projects and win free swag \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/hacktoberfest-returns-contribute-to-open-source-projects-and-win-free-swag/)

Hacktoberfestは、10月中に4つGitHub上でPullRequestsがMergeされるとTシャツなどのノベルティがもらえるオンライン上のイベントです。

Mattermostでは今年もHacktober向けのIssueを多く用意しています。興味のある方は、まず https://hacktoberfest.digitalocean.com/ で参加登録をしてから、下記のIssueリストから手をつけられそうなものを探してみてください。

https://github.com/mattermost/mattermost-server/labels/Hacktoberfest

### NRI x Mattermost社 イベント

先日、Mattermost社が参加する初の日本でのイベントが、野村総合研究所主催で行われました。
[イベント・セミナー \| aslead/アスリード \| 野村総合研究所\(NRI\)](https://aslead.nri.co.jp/event/index.html#about)

イベントに合わせて野村総合研究所によるMattermostの紹介記事が公開されています。
[NRIが全社利用するビジネスチャットツールMattermostのご紹介](https://www.slideshare.net/daisukekato13/nrimattermost)


## おわりに

次の`v5.17`のリリースは2019/11/15(Fri)を予定しています。
そして機能追加が行われる`v5.18`は恐らく2019/12/16(Mon)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
