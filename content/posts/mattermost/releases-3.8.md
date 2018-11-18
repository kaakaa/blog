
---
title: "Mattermost3.8の新機能"
date: 2017-04-19T11:26:07+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

2017/4/16にMattermost 3.8.0がリリースされましたので、追加された機能などについて簡単に紹介します。
(現在、v3.8.1のリリースを検討しているようなので、安定稼働を求める方はもう少し待ったほうが良いかもしれません)

今回のリリースから隔月リリースから毎月16日のリリースサイクルに戻したようです。
(2016年9月から隔月リリースに変更していました)
[Mattermost moves to monthly releases](https://about.mattermost.com/mattermost-moves-to-monthly-releases/)

詳細については公式のリリースを確認ください。
[Mattermost 3.8: Deploy with next generation iOS and Android Apps in Beta – Now Faster than Ever](https://about.mattermost.com/mattermost-3-8/)
[Mattermost Changelog - Release v3.8.0](https://docs.mattermost.com/administration/changelog.html#release-v3-8-0)

## iOS/Android アプリ

ReactNativeで開発されたiOS/Androidアプリのベータ版がリリースされています。
Android版はGoogle Play Storeから、iOSの場合は下記`Sign up`でのユーザー登録と[TestFlight](https://itunes.apple.com/jp/app/testflight/id899247664?mt=8)のインストールが必要です。

> To test on Android: Search for “Mattermost 2017 – Beta” in the Google Play store, or opt-in here.
> To test on iOS: Sign up to become a beta tester.

## 3.8 New End User Features

### Emoji機能の拡充 - Emoji Picker

絵文字選択機能がプリリリースされました。
今までのバージョンでは :emoji: の形式でしかEmoji入力入力・リアクションの付与ができませんでしたが、3.8からは画面上で入力できるようになります。

絵文字選択機能を使用するには`アカウント設定` > `詳細` > `プリリリース機能をプレビューする` から行うことができます。

![mm3.8_emoji_picker1.png](https://qiita-image-store.s3.amazonaws.com/0/9891/6685280a-c5d9-b7f8-5550-048bfe4c095d.png)

上記の設定を有効にすると、メッセージ入力ボックスに絵文字選択ボタンが出現します。
![mm3.8_emoji_picker2.png](https://qiita-image-store.s3.amazonaws.com/0/9891/8d86b429-d025-9e7c-4ca3-80c27372775c.png)

また、絵文字リアクションを付与する際にも絵文字選択機能を使用することができるようになっています。
![mm3.8_emoji_picker3.png](https://qiita-image-store.s3.amazonaws.com/0/9891/1023daf7-f6f6-9ebf-187f-dfde8c5c64d1.png)



### Pinned Post

チャンネルごとにピン止めされた投稿を設定することができるようになりました。
投稿の横にある`[...]`メニューから`チャンネルにピン止めする`を選択することで、ピン止めすることができます。

![mm3.8_pinned1.png](https://qiita-image-store.s3.amazonaws.com/0/9891/ff5d8c7c-a223-6830-1638-5780fec83ae4.png)

ピン止めされた投稿は画面上部の :pushpin: アイコンをクリックすることで確認できます。

![mm3.8_pinned2.png](https://qiita-image-store.s3.amazonaws.com/0/9891/a29ee696-f1a1-fe23-30ed-55a7df4edf89.png)


## 3.8 New System Admin Tools

### システムコンソールのユーザーリスト/チーム管理

今までのバージョンでは、ユーザー一覧はチームごとにしか確認できませんでしたが、システムコンソールで全てのユーザーの一覧を確認・管理できるようになりました。
また、新しくチーム統計画面が追加されています。

`チームの統計`画面
![mm3.8_team.png](https://qiita-image-store.s3.amazonaws.com/0/9891/d86dd799-76f6-ab4b-853e-6f3d4e84a059.png)

`ユーザーリスト`画面
![mm3.8_userlist.png](https://qiita-image-store.s3.amazonaws.com/0/9891/5168e3f9-17fe-4ba5-5783-08692f17da6a.png)


## クラウド自動デプロイがよりシンプルに

様々な環境でMattermostを動作させるため、設定情報を環境変数から指定できるようになりました。
`config.json`に書かれた設定情報を環境変数によって上書きすることができるため、多様なクラウド環境へのデプロイ設定が容易になります。

例えば`config.json`内で`.ServiceSetting.ListenAddress`をキーとする設定は、環境変数`MM_SERVICESETTINGS_LISTENADDRESS`によって上書きすることができます。

## Community Integration


### Siba Bot Integration

チャットボットツールである[Siba](https://siba.ai)のMattermostアダプターです。
[SibaのDocumentationページはまだ準備中らしく](https://siba.ai/resources/docs/)、Sibaがどのようなものかは分かりませんでした。


### Conduct a Quick Emoji Poll

v3.6から追加されたリアクション機能を利用した投票メッセージの作成を簡単にするサーバーアプリを作りました。
詳細は[Mattermostで投票機能もどき](http://qiita.com/kaakaa_hoe/items/b2605ce3816cfc517ecd)に書いています。

### Simple CLI Utility for Posting Text in PHP

Mattermostの内向きのウェブフックを利用して、コマンドラインからメッセージを投稿できるPHPツールです。
[joho1968/mattermostsendphp](https://github.com/joho1968/mattermostsendphp)


### Simple Mattermost /slash Service Dispatcher and Responder in PHP

PHP製のスラッシュコマンドのディスパッチャです。
[joho1968/mmsrvdispatch](https://github.com/joho1968/mmsrvdispatch)

## Removed and Deprecated Features

* v3.5以前に使われていた古いCLIツールが削除されました
* 下記のAPIが削除されました
  * GET at /channels/more (replaced by /channels/more/{offset}/{limit})
  * POST at /channels/update_last_viewed_at (replaced by /channels/view)
  * POST at /channels/set_last_viewed_at (replaced by /channels/view)
  * POST at /users/status/set_active_channel (replaced by /channels/view)

[Removed and deprecated features for Mattermost](https://about.mattermost.com/deprecated-features/)

## おわりに

次回のv3.9.0のリリースは2017/5/16を予定しています。

v3.9.0向けに現在開発中の機能として下記が挙げられています。
* チャンネルのアーカイブとアーカイブからの復元
* 公開/非公開チャンネルの切り替え
* 参加/脱退システムメッセージの結合

[Roadmap - Mattermost](https://about.mattermost.com/direction/)
> Coming Soon (In Progress by Community)
> 
> Community members are currently working on the features listed below. They are not scheduled for a specific release, but will be added to the product when work is complete.
> 
> * Unarchive channels
> * Converting public channels to private
> * Combining consecutive join/leave messages

