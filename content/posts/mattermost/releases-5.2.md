
---
title: "Mattermost5.2の新機能"
date: 2018-08-23T13:28:39+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/08/16にMattermost 5.2.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost 5\.2: Upgraded plugin system, search archived channels, Romanian language support and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-2-upgraded-plugin-system-search-archived-channels-romanian-language-support-and-more/)
* [Mattermost Changelog — Mattermost 5\.2 documentation](https://docs.mattermost.com/administration/changelog.html)

---

v5.2.0でアーカイブ後のチャンネル内容を検索できる機能が追加されましたが、この機能に不具合が発生している為、この機能をデフォルトでOffとする変更を加えたv5.2.1のリリースに向けた作業が進んでいます。もし、同様の事象が発生している場合は修正版のリリースをお待ちください。

# アップデート内容

## プラグインシステムのアップデート
Mattermost v4.4からベータ版として利用可能だったプラグイン機能のアーキテクチャが刷新されました。

このアップデートにより、下記のような機能が利用可能になっています。

* **プラグインからのログ出力** - 新しいロギングAPIを使ってMattermost本体のサーバーログにプラグインからログを出力できるようになりました
* **カスタムWebSocketイベント** - サーバーサイドでプラグイン独自のWebSocketイベントを送信できるようになり、また、クライアントサイドでそのイベントを扱う機能が追加されました
* **APIの拡張** - サーバーサイド・クライアントサイド共に利用できるAPIが追加されました
* **CLIでの管理** - Mattermostのコマンドラインインターフェースからプラグインを直接管理できるようになりました

このアップグレードにより、以前のアーキテクチャー向けに開発されたプラグインは動作しなくなっています。以下のマイグレーションガイドを参考に新しいプラグインアーキテクチャーに対応してください。
[Migrating Plugins from Mattermost 5.1 and earlier](https://developers.mattermost.com/extend/plugins/migration/)

今回の変更の詳細については、Mattermostフォーラムの投稿にまとめられています。
[Plugin System Upgrade in Mattermost 5\.2 \- Announcements \- Mattermost Peer\-to\-Peer Forum](https://forum.mattermost.org/t/plugin-system-upgrade-in-mattermost-5-2/5498)

かいつまむと...

* MattermostのUIコンポーネントを拡張する際、今までは単一の(React)コンポーネントを複数のプラグインから拡張することはできませんでしたが、今回のアップグレードにより可能になりました
* 今まではサーバーサイドのプラグインにGo言語のプラグイン機構を使用していましたが、今回のアップデートから[hashicorp/go\-plugin](https://github.com/hashicorp/go-plugin)を使用するようになりました


より詳しく知りたい場合は、[デモプラグイン](https://github.com/mattermost/mattermost-plugin-demo)　や　[サンプルプラグイン](https://github.com/mattermost/mattermost-plugin-sample)、または[開発者向けドキュメント](https://developers.mattermost.com/extend/plugins/)などを参考にしてください。

---

(モバイルアプリはプラグインにまだ対応していないようなので、モバイルアプリをよく利用する環境の場合はプラグインインストール時に注意してください)

## アーカイブされたチャンネルの検索
Mattermostの画面からアーカイブされたチャンネルを閲覧・検索できるようになりました。
詳細については[公式ドキュメント](https://docs.mattermost.com/help/getting-started/organizing-conversations.html#archiving-a-channel)を確認してください。

<img width="926" alt="スクリーンショット 2018-08-22 20.30.09.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/78c9dc29-df12-810d-7474-4d692bfc5f6a.png">


## 埋め込みMattermost(ベータ版)
OAuth2.0を通じて、Mattemostを別のアプリやウェブサイトに埋め込むことができるようになりました。
詳しく知りたい場合は、サンプルの[Chromeエクステンション](https://github.com/mattermost/mattermost-chrome-extension)や、[公式ドキュメント](https://docs.mattermost.com/integrations/embedding.html)を参照してください。

---

(サイト内チャットのような埋め込みではなく、Mattermost起動ボタンを配置できるようになったということですかね？)

## ルーマニア語のサポート
今バージョンよりMattermostの表示言語にルーマニア語が選択可能になりました。
現在、Mattemostでは15言語をサポートしています。

日本語の翻訳ももちろんあり、下記サイトから翻訳作業を行うことができますので気になる点があれば報告をお願いします。
[Welcome \| Mattermost Translation Server](https://translate.mattermost.com/)

## モバイルアプリでのディープリンク
パーマリンクよりMattermostのモバイルアプリが自動で開くようになりました。
Android、iOSアプリ共に対応しています。

詳しくは[公式ドキュメント](https://docs.mattermost.com/mobile/mobile-faq.html#how-do-i-configure-deep-linking)を参照してください。

## その他

### アンチウイルスプラグイン(ベータ版)
[mattermost/mattermost\-plugin\-antivirus: Antivirus plugin for scanning files uploaded to Mattermost](https://github.com/mattermost/mattermost-plugin-antivirus)
[ClamavNet](https://www.clamav.net/assets/Ill-01.png)を利用して、Mattermostへアップロードしたファイルのウイルススキャンを行うプラグインのようです。

### GitHubプラグイン(ベータ版)
[mattermost/mattermost\-plugin\-github: Experimental GitHub plugin for Mattermost Written for Mattercon Hackathon 2018](https://github.com/mattermost/mattermost-plugin-github)
GitHubと連携するプラグインです。GitHub Enterpriseで使おうとしましたが、まだバグがあるようで動きませんでした。


### プラグインハッカソン'18
プラグインアーキテクチャーの刷新を記念して、オンラインカンファレンスサービスであるZoomを利用したVirtual Hackasonが8/16,17に開催されました。
[Virtual Hackathon\! \| Meetup](https://www.meetup.com/ja-JP/mattermost/events/253346351/?eventId=253346351&rv=ea1_v2&rv=ea1_v2)

ハッカソン向けのプラグインアイデア集は下記に記載されています。
https://docs.google.com/spreadsheets/d/1Xxy4J7wtchtCMXVBxl0TL2h0uHOet96w6KIGFKuwYvs/edit#gid=0

後日、ハッカソンで開発したプラグインを披露する場が設けられていたのですが、そちらの模様は下記にアップロードされています
https://drive.google.com/drive/folders/1cJJZVUYax33GFoE6ckRT9TSL3kLfScya


# おわりに

次回の`v5.3`のリリースは2018/9/14(Fri)を予定しています。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。


