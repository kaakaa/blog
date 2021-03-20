---
title: "Mattermost5.20の新機能"
date: 2020-02-17T22:56:12+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/2/14にMattermost v5.20.0がリリースされました。また、2/16にバグFixを含んだ `v5.20.1` がリリースされています。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.20: New mobile editor, desktop dark theme, and more \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-20-new-mobile-editor-desktop-dark-theme/)
* https://docs.mattermost.com/administration/changelog.html#release-v5-20-feature-release

---

## モバイルアプリのメッセージ編集画面の改善

Mattermost v5.20.0と同時にリリースされているMattermostモバイルアプリv1.28.0で、@メンションやスラッシュコマンド呼び出しなど、よく使う機能にすぐにアクセスするためのボタンが追加されました。

![mobileapp](https://mattermost.com/wp-content/uploads/2020/02/New-mobile-editor.png)
(公式ブログに掲載されている画像です)

## 最新のメッセージへ移動するためのバナー

未読メッセージが多くあるチャンネルに、クリックすると最新の投稿へ移動するためのバナーが表示されるようになりました。

![banner](https://mattermost.com/wp-content/uploads/2020/02/Jump-to-recent-posts.png)
(公式ブログに掲載されている画像です)

## `mmctl` のリリース

リモートサーバーにあるMattermostを操作するためのCLIツールである `mmctl` がリリースされました。

https://github.com/mattermost/mmctl

## 同梱プラグイン
セキュリティの問題などでMattermostサーバーをインターネットへ接続できないように設定していた場合、いくつかの同梱プラグインはプリインストールされた状態で利用することはできましたが、プラグインマーケットプレースを通じたプラグインの管理を行うことができませんでした。
しかし、Mattermost 5.20からはインターネットから切断されたサーバーであっても、Mattermost本体に同梱されたバイナリについてはマーケットプレースを通じて管理できるようになりました。同梱プラグインは公式チームが開発しているものに限ります。

また、マーケットプレース画面で、「公式プラグイン」「コミュニティプラグイン」「ベータ版」などを識別しやすくなりました。

![marketplace](https://mattermost.com/wp-content/uploads/2020/02/Plugins.png)
(公式ブログに掲載されている画像です)

さらに、独自で開発したMattermostプラグインがある場合、下記フォームからプラグインの情報を送ることで[Mattermost Integrations](https://integrations.mattermost.com/)に追加することができます。
[Mattermost Integrations and Installers](https://spinpunch.wufoo.com/forms/mattermost-integrations-and-installers/)

## デスクトップアプリのUI/UX改善
* 新しいダークテーマ
* タブをドラッグ＆ドロップで並べ替え可能に
* タブでメンションと未読を表すデザインを更新
![darktheme](https://mattermost.com/wp-content/uploads/2020/02/Desktop-app-dark-theme.png)
(公式ブログに掲載されている画像です)

## 未読メッセージがある場合のfavicon

未読のメンションがある場合、faviconに赤いドットが付与されるようになりました
![favicon](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.20/favicon.png)

# おわりに

次の`v5.21`のリリースは2020/3/16(Mon)を予定しています。
そして機能追加が行われる`v5.22`は恐らく2020/4/17(Thu)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
