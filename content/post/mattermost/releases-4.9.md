
---
title: "Mattermost4.9の新機能"
date: 2018-04-17T22:24:28+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/04/16にMattermost 4.9.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4\.9: Muted channels, team icons, Global Relay and more – Mattermost](https://about.mattermost.com/releases/mattermost-4-9/)
[Mattermost Changelog — Mattermost 4\.9 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-9)


# アップデート内容

## チャンネルのミュート

デスクトップ通知やメール通知の有無をチャンネルごとに設定できるミュート機能が追加されました。

スラッシュコマンド `/mute` を実行することで、現在のチャンネルをミュートすることができます。
`/mute ~[チャンネル名]`とすることで、指定したチャンネルをミュートすることもできます。

<img width="218" alt="スクリーンショット 2018-04-17 21.11.32.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/ff70bc13-97e8-c7fd-5d7c-676116633c73.png">

ミュートされたチャンネルはチャンネル名の文字色が薄くなります。(上記画像では`タウンスクウェア`チャンネル)
また、メンションがあった場合のチャンネル横の通知数は表示されたままになります（上記画像では`ミュート機能のテスト`チャンネル）

## チームメイトの名前の表示方法

チームメイトの名前の表示方法を柔軟に設定できるようになりました。
ユーザー名/ニックネーム/氏名から選ぶことができます。

<img width="797" alt="スクリーンショット 2018-04-17 21.13.59.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/facca489-692c-0dd7-cdc0-bda3a56f98bc.png">

システム管理者は `システムコンソール` > `全般` > `ユーザーとチーム` > `チームメイトの名前の表示` から、この設定のデフォルト値を設定することができます。

## チームアイコンのカスタマイズ

左サイドバーに表示されるチームアイコンを設定することができるようになりました。
`チームの設定` > `チームアイコン` からアップロードすることができます。

<img width="286" alt="スクリーンショット 2018-04-17 21.19.40.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/926187e3-5bd9-1246-596c-9f64caa30f20.png">

## Global Relayコンプライアンスの自動エクスポート(E20)

E20プランの顧客を対象に、組織内のポリシーを遵守するためのGlobal Relay形式でのメッセージの自動エクスポートができるようになりました。
(v4.5からActiance XML形式のエクスポートには対応していました)

## モバイルアプリのエンハンス

先日リリースされたモバイルアプリ 1.7.0 versionでは、下記のような機能追加が行われました。

* 他のアプリからMattermostへの添付ファイルを共有できるようになりました
* Markdown形式のテーブルを確認することができるようになりました
* (投稿の?)パーマリンクをアプリ内で開くことができるようになりました(今までは別ブラウザが立ち上がっていました)
* アナウンスバナーが表示されるようになりました
* アプリが今までより素早く立ち上がるようになりました

## Coming Soon: Advanced Permissions(E10, E20)

顧客からの要望を受け、より細かな権限管理が行えるようになります。
今回のリリースではバックエンドの変更のみのため、ユーザーから行える変更はまだありません。

今後、顧客の要望に合わせた様々な権限設定を予定しており、2018年夏リリース予定のMattermost5.0で次のフェーズを公開するつもりです。

# その他

## 公式Docker compose

今回のリリースに伴う変更により、Breaking Changeとなりそうな変更が加えられています。（未確認）
[Release 4\.9\.0 · mattermost/mattermost\-docker](https://github.com/mattermost/mattermost-docker/releases/tag/4.9.0)

ホスト上のディレクトリをマウントしている場合に、UID/GIDを適切に設定しなくてはならなくなったようです。
また、Mattermostアプリケーションの起動ポートについて、以前は80番ポートを使用していましたが、今回のリリースから8000番ポートを利用するように変更されたため、設定変更が必要となる場合があります。

公式のdocker-compose.ymlを利用している場合は注意してください。

## Mattermost 5.0で廃止となる機能

2018年夏リリース予定のMattermost v5.0で廃止となる機能についての議論が開始されています。
以前からアナウンスされていたAPI v3の廃止のほか、管理用のCLIツールである`platform`コマンドの名称変更などが予定されています。
[Upcoming Deprecated Features in Mattermost v5.0](https://docs.mattermost.com/administration/changelog.html#upcoming-deprecated-features-in-mattermost-v5-0)

## Twitter: mattermost_jp

１年ほど更新の無かったMattermostの日本語版Twitterアカウントですが、Mattermostに関する日本語のツイートがほぼ毎日あるのに更新がないのは勿体無いと思い、Mattermostチームからログイン権限をもらいMattermostに関する情報の発信を始めました。
Mattermostに関する日本語の情報をツイートしているので、フォローいただければと思います。
[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja)

ベストエフォートで運営しています。

# おわりに

次回の`v4.10.0`のリリースは2018/5/16を予定しています。

