
---
title: "Mattermost4.6の新機能"
date: 2018-01-27T23:22:52+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/01/16にMattermost 4.6.0がリリースされたので、追加された機能などについて簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4\.6: Faster channels, enhanced 508 compliance – Mattermost](https://about.mattermost.com/releases/mattermost-4-6/)
[Mattermost Changelog — Mattermost 4\.6 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-6)


# 新機能

今回のリリースでは年末年始の休暇を挟んだこともあり、大きな新機能のリリースは無く、改善系の更新が多かった模様です。

## Webアプリのパフォーマンス改善

MattermostのWebアプリのパフォーマンスが改善されました。

* 同じチーム内でチャンネルを切り替える際の時間が、最大で45%ほど短縮されたようです
* チャンネルを切り替えた際のメモリ使用量が最大で85%削減されたようです
* アイコンやロゴのサイズも70-80%ほど削減されたようです

## 通知設定のデフォルト値の更新

デフォルト通知設定がMattermostユーザー・コミュニティによって推奨された設定に変更されたようです。

* デスクトップ通知はメンションとダイレクトメッセージがあった時のみ通知されます
* モバイルプッシュ通知はユーザーが離席中かオフラインの時のみ通知され、オンラインの際は通知されません
* ファーストネームでのメンションでは通知されません

上記の通設定は、アカウント設定よりいつでも変更できます。

## APIv4が公式APIに

今回のリリースより、APIv3がdeprecatedになり、APIv4が公式のAPIとなりました。
APIv3はサポート対象外となり、Mattermost5.0で削除されます。

APIv3からAPIv4へのマイグレーションについてのTIPSが下記にあります。
[Mattermost API Reference](https://api.mattermost.com/#tag/schema)

## オープンソースプロジェクト

### Mattermost Poll Integration with Firebase

Firebaseを利用したMattermost上での投票機能のようです。
[jedfonner/MattermostOnFire: Firebase Realtime Database and Firebase Cloud Function code to power a Mattermost slash command for creating a poll using Mattermost Interactive Buttons\.](https://github.com/jedfonner/MattermostOnFire)

### CloudApp Integration
デスクトップのキャプチャやGIF動画などを作成できるサービスである[CoundApp](https://www.getcloudapp.com/)とMattermostとの連携機能のようです。
[Mattermost Integration I CloudApp Video Screen Recorder](https://www.getcloudapp.com/integrations/mattermost)

# MVP

私事ですが、今回のv4.6リリースのMVPに選ばれていました。
大きな貢献はまだできていないですが、日本語の翻訳や、細々したコントリビュートなどが評価されたものと思っています。
また、年末年始の休暇中に時間を割くようなワーカホリックな国民性の賜物かと思っています。

![DUO-M8QVoAAI-7r.jpg](https://qiita-image-store.s3.amazonaws.com/0/9891/a0db2679-9999-d4e1-cf5c-bdbc6788e80f.jpeg)


# おわりに

次回の `v4.7.0`のリリースは2018/2/16を予定しています。

