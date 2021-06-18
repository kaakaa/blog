---
title: "Mattermost5.18の新機能"
date: 2019-12-21T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2019/12/16にMattermost v5.18.0がリリースされました。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.18: ID\-only option for push notifications \(E20\), one\-click plugin updates, mark posts as unread, and more \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-18-id-only-option-for-push-notifications-one-click-plugin-updates-mark-posts-unread-and-more/)
* https://docs.mattermost.com/administration/changelog.html#release-v5-18-feature-release

---

## モバイルアプリのクラッシュ

Mattemrost 5.18と同日にアップデートされたモバイル向けアプリ(iOS, Android)でMattermostに接続できない問題が発生しているようです。

問題の一つとして、新しいバージョンのアプリでMattermost 5.9より前のバージョンのMattermostサーバーへ接続しようとするとクラッシュするという問題があるようです。

現在、Mattermost 5.9より前のバージョンに接続してもクラッシュしないようにする修正を加えた新しいバージョン(**v1.26.1**)がリリースされているようです。(AppStoreで確認しましたが、AppStoreにはまだリリースされていないようです)
https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#mattermost-mobile-apps-changelog

---

# アップデート内容

## (E20) モバイルアプリ向けプッシュ通知の改善

現在、モバイルアプリへのプッシュ通知は、通知内のメッセージ内容を含めてApple Push Notification Service (APNS) もしくは Google Firebase Cloud Messaging (FCM) を経由してデバイスへ送信されています。これは、非常に厳密なセキュリティやコンプライアンスが求められる組織では問題となり得ます。

E20ライセンスに追加された **ID-only** オプションを利用すると、プッシュ通知内のペイロードにはメッセージIDのみが含まれるようになり、メッセージ内容がAPNSやFCMへ送信されることがなくなります。通知内のメッセージ内容については、デバイスから直接Mattermostサーバーへ取得しにいくようになります。

## 1クリックでプラグインをアップデート

インストールされているプラグインのアップデートが簡単になりました。

[Mattermost v5.16で **プラグインマーケットプレース** の機能](https://blog.kaakaa.dev/post/mattermost/releases-5.16/#%E3%83%97%E3%83%A9%E3%82%B0%E3%82%A4%E3%83%B3%E3%83%9E%E3%83%BC%E3%82%B1%E3%83%83%E3%83%88%E3%83%97%E3%83%AC%E3%83%BC%E3%82%B9)が追加されましたが、その **プラグインマーケットプレース** のページでインストール済のプラグインのアップデートを行えるようになりました。

インストール済のプラグインの新しいバージョンがリリースされている場合、**アップデートする**というリンクが表示されます。

![plugin_update](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/plugin_update_1.png)

**アップデートする**ボタンをクリックとすると確認ダイアログが表示されます。リリースノートへのリンクも表示されているため、ここから更新内容を確認しにいくこともできます。

![plugin_update](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/plugin_update_2.png)

## 未読としてマーク
新しいメッセージは、それが後で読み返したい重要なメッセージであっても、一度表示してしまうと既読としてマークされてしまい、他のメッセージに埋もれてしまっていました。
しかし、今回のアップデートにより追加された **未読としてマークする** 機能を使うことで、後で見返したいメッセージを再度未読状態に戻すことができます。

以下のサンプルは[公式ブログ](https://mattermost.com/blog/mattermost-5-18-id-only-option-for-push-notifications-one-click-plugin-updates-mark-posts-unread-and-more/)からの転載です。

![mark_as_unread](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/mark_as_unread.gif)


## アーカイブされたチャンネルの参照

アーカイブされたチャンネルを簡単に参照できるようになりました。

まだベータ版の機能ですが、Mattermostサーバー内の `config.json` にある `ExperimentalViewArchivedChannels` の設定を `true` にすることで、参加可能チャンネルの一覧にアーカイブされたチャンネルを表示できるようになります。
([公式ドキュメント](https://docs.mattermost.com/administration/config-settings.html#allow-users-to-view-archived-channels-beta)では設定名が `ViewArchivedChannels` となっていますが、おそらく誤りだと思われるため[修正のPR](https://github.com/mattermost/docs/pull/3278)を出しています。)

![view_archived_channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/view_archived_channel_1.png)

アーカイブチャンネルは普通のチャンネルと同じように開くことができますが、書き込みはできません。

![view_archived_channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/view_archived_channel_2.png)

また、この設定を有効にすると、アーカイブチャンネル内のコメントを検索できるようにもなります。

![view_archived_channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/view_archived_channel_3.png)


## その他の変更点
* JIRAプラグインがアップデートされ、１チャンネルで複数のJIRAプロジェクトの通知を受けれるようになったり、通知のフィルタリングができるようになっているようです
  * 今回のリリースとは関係ないとは思いますが、JIRAプラグイン用のドキュメントサイトもあるようです
  * https://mattermost.gitbook.io/plugin-jira/
* リモート端末からMattermostを操作できる新しいCLIツール `mmctl` が公開されています
  * https://github.com/mattermost/mmctl
* (E10/E20) v5.16で追加された[ゲストアカウント](https://blog.kaakaa.dev/post/mattermost/releases-5.16/#e10-20-%E3%82%B2%E3%82%B9%E3%83%88%E3%82%A2%E3%82%AB%E3%82%A6%E3%83%B3%E3%83%88)機能に、AD/LDAPとSAMLから直接ゲストアカウントを作成する機能が追加されました
* (E20)コンプライアンス用のActiance Export Toolで投稿/ファイルの削除/修正イベントを扱えるようになりました
* (E20) AD/LDAPグループ同期機能が実験的な機能(Experimental)からベータ版(Beta)となりました

### コードブロックに行番号

シンタックスハイライト付きでコードブロックを投稿した際に、行番号が表示されるようになりました。

![line_number](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.18/line_number.png)

# MISC

## Mattermost Hackathon
11/22-26の間、Mattermostコミュニティのヴァーチャルハッカソンが開催されました。
以下の記事で、ハッカソンの成果が紹介されています。数が多すぎて見切れていませんが、Calendar PluginやAir Quality Index Pluginなど気になる物がいくつもあります。
https://mattermost.com/blog/mattermost-hackathon-2019-highlights/

## NRI x Mattermost社イベント動画

10月に行われた野村総合研究所主催のMattermostイベントの動画がYouTubeで見られるようです。
https://www.youtube.com/channel/UC-FNtllB7J7nAgxU5FXxyVQ

## LINC Biz - Mattermostベースサービス？

SHARPの子会社よりMattermostをベースとしていると思われるサービスが開始されました。（公式にMattermostを利用していると発表している訳ではないので、Mattermostを利用していないかもしれません）

https://corporate.jp.sharp/news/191128-b.html

# おわりに

次の`v5.19`のリリースは2020/1/16(Thu)を予定しています。
そして機能追加が行われる`v5.20`は恐らく2020/2/17(Mon)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
