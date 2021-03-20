
---
title: "Mattermost5.1の新機能"
date: 2018-07-23T08:31:18+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/07/16にMattermost 5.1.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost 5\.1: New GIF selector, auto\-linking plugin, subpath support and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-1-new-gif-selector-auto-linking-plugin-subpath-support-and-more/)
* [Mattermost Changelog — Mattermost 5\.1 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-1)

# アップデート内容

## Gfycat

MattermostでGIF動画を簡単に入力できるようになりました。
絵文字選択メニューからGfycatのGIF動画を検索・選択し、入力することができます。

![gifycat.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/f43660c3-535b-34ca-96f6-d9305f292d6c.gif)


**システムコンソール** > **カスタマイズ** > **GIF(Beta)** > **GIF選択機能を有効にする**から、この機能の有効/無効を切り替えられます。
GfycatのAPIキーを設定する必要がありますが、最初から利用可能なAPIキーが入力されているため、自分のAPIキーを使いたいなどの理由がない限り変更する必要はありません。

GIFアニメーションの表示をOFFにしたい場合、`/collapse`というメッセージを投稿することでGIFの自動表示をOFFにすることができます。
自動表示を再度ONにする場合は`/expand`というメッセージを投稿してください。

## Autolink

投稿がデータベースに保存される直前に処理を追加できるプラグインAPIを利用したAutoLinkプラグインの紹介です。

[mattermost/mattermost\-plugin\-autolink: Automatically rewrite text matching a regular expression into a markdown link\.](https://github.com/mattermost/mattermost-plugin-autolink)

これにより、ある特定の文字列を入力した場合に自動でリンク文字列に変換することができます。

![v5.1_autolink.mov.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/243b77a7-984c-1e72-d544-84e9bdca754b.gif)

AutoLinkプラグインの最新リリースバージョンは`v0.3.0`ですが、このバージョンはMattermost v5.2で導入される新しいプラグイン機構に基づいて作成されているため、Mattermost v5.1で利用する場合はAutoLinkプラグイン v0.2.0を利用してください。

## サブパスでのMattermostのホスト

`https://mattermost.example.com/mattermost`のようなサブパスでのMattermostのホスティングが可能になりました。
**システムコンソール** > **全般** > **設定**  > **サイトURL** でサブパスを含むURL(``https://mattermost.example.com/mattermost``)を設定し、サーバーを再起動することで有効になります。

## モバイルアプリ

* アナウンスバナーでのMarkdown記法がサポートされました
* Android向けにドロワーレイアウトの方法を変更し、チャンネルの切り替えがし易くなりました
* プッシュ通知から投稿を取得する際の読み込み速度の向上

そのほか、多くの改善が行われています。
[mattermost\-mobile/CHANGELOG\.md at master · mattermost/mattermost\-mobile](https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#v1100-release)


また、先月リリースされたv1.9.0で正常に日本語入力ができないという問題がありましたが、2018/7/4にリリースされた[v1.9.3](https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#193-release)にて解消されています。

## (E20) コンプライアンスレポートのアップグレード

参加、脱退、ファイルアップロードなどのイベントのCSV形式でのエクスポートが可能となりました。
コンプライアンスエクスポート機能はE20ライセンスのみの機能になります。

CSVエクスポートのサンプルファイルが下記よりダウンロードできます。
https://github.com/mattermost/docs/blob/master/source/samples/csv_export.zip




# おわりに

次回の`v5.2`のリリースは2018/8/16を予定しています。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。

