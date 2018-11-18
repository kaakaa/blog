
---
title: "Mattermost4.7の新機能"
date: 2018-02-17T23:08:43+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/02/16にMattermost 4.7.0がリリースされたので、追加された機能などについて簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4\.7: Enhanced image preview and thumbnails, faster load times, upgraded desktop app – Mattermost](https://about.mattermost.com/releases/mattermost-4-7/)
[Mattermost Changelog — Mattermost 4\.7 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-7)

また、Compliance Export機能に問題があったようで、近々`v4.7.1`のリリースが予定されています。

# 新機能

## 画像サムネイルの改善

画像サムネイルの見た目が改善されました。

添付画像ファイルが一つの場合、サムネイルが自動で拡大表示されるようになっています。

![single_image.png](https://qiita-image-store.s3.amazonaws.com/0/9891/d743cb90-8d27-2999-0148-aa3f4b07213c.png)


複数ファイルが添付された場合のサムネイル表示も若干変更されています。
ダウンロードボタンが付与されたことで、プレビューを開かずにダウンロードができるようになっています。

![multi_image.png](https://qiita-image-store.s3.amazonaws.com/0/9891/46f1ddaf-22f0-05fd-d7e1-dfae5c63526b.png)


## ロード時間の短縮

* カスタム絵文字のロードが非同期処理となり、最初のページンのロード時間が高速化されました
* チャンネル名の補完処理が最適化されました
* 不要なメタデータを削除したことで、ほとんどの画像のサイズが25%以上削減され、パフォーマンスが改善されました

## 画像プロキシ

クライアントからの画像のリクエストをプロキシサーバーへ直接送信することができるようになりました。
これにより、画像はプロキシサーバーにキャッシュされるためパフォーマンスが向上します。また、画像がキャッシュされるため、リンク先の画像が消失することを心配する必要がなくなりました。

不要な場合にフルサイズの画像をダウンロードする必要もなくなりました。
さらに、[willnorris/imageproxy](https://github.com/willnorris/imageproxy)を利用すると、画像のリサイズなどが可能となる`ImageProxyOption`の設定ができるようになります。

画像プロキシは **システムコンソール** > **ファイル** > **ストレージ** から設定できます。

## Desktopアプリ 4.0

デスクトップアプリのメジャーアップデート版である`v4.0`がリリースされています。
[desktop/CHANGELOG\.md at beta\-4\.0 · mattermost/desktop](https://github.com/mattermost/desktop/blob/beta-4.0/CHANGELOG.md#release-v400)

* **UX改善**: 予期せずアプリが終了した際に再度アプリを開くためのダイアログや、ビデオ通話のためにマイクやあカメラへのアクセスを求めるダイアログなどが追加されました。また、ページをロードしている間、Mattermostアイコンのアニメーションが表示されるようになりました。
* **ディープリンク**: `mattermost://`を通じたディープリンクプロトコルをサポートしました。ユーザーが`mattermost://`で始まるリンクをクリックすることで、デスクトップをアプリを開いたり、特定のチャンネルへジャンプすることができます。
* **デプロイの簡易化**: 管理者がエンドユーザーへMattermostデスクトップアプリをサイレントデプロイするのがとても簡単になりました。

これらの新機能の他に[アーキテクチャーの変更](https://github.com/mattermost/desktop/blob/master/CHANGELOG.md#architectural-changes)も行われています。

## '未読' サイドバー設定

未読メッセージのあるチャンネルをサイドバーの一番上にまとめて表示する実験的な機能を追加しました。
[Account Settings — Mattermost 4\.7 documentation](https://docs.mattermost.com/help/settings/account-settings.html?highlight=unreads#group-unreads-channels)

Mattermostの設定ファイル`config.json`の`ExperimentalGroupUnreadChannels`の設定値を`default_on`もしくは`default_off`とすると、**アカウント設定** > **サイドバー** > **未読チャンネルのグループ化** より未読チャンネルをグループ化するかの設定を行うことができます。
`ExperimentalGroupUnreadChannels`の設定値が`disabled`だと、**サイドバー**のメニューが表示されません。

その他のサイドバーに関する改善にも取り組んでいるそうです。

## その他(APIv3について)

先月のリリースでAPIv3がdeprecatedとなりましたが、APIv3はMattermost 5.0で削除される予定です。
APIv3からAPIv4へのマイグレーションについては下記を参考にしてください。
[Mattermost API Reference](https://api.mattermost.com/#tag/APIv3-Deprecation)

# おわりに

次回の`v4.8.0`のリリースは2018/3/16を予定しています。

