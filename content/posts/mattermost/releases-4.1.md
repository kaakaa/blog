
---
title: "Mattermost4.10の新機能"
date: 2018-05-19T20:32:34+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/05/16にMattermost 4.10.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost 4\.10: Up to 90% faster first load time – Mattermost](https://about.mattermost.com/releases/mattermost-4-10/)
* [Mattermost Changelog — Mattermost 4\.10 documentation](https://docs.mattermost.com/administration/changelog.html)

# アップデート内容

## メッセージ読み込みの高速化
ページを更新した後、メッセージを読み込む速度が以前より90%以上も高速化されたようです

## 公開チャンネルを非公開チャンネルに

既存の公開チャンネルを非公開チャンネルに変更できるようになりました。
途中から協力会社の方々がチームに加入し、今までの会話を見えないようにしくてはいけない状況などに有効です。

![スクリーンショット 2018-05-19 17.50.50.png](https://qiita-image-store.s3.amazonaws.com/0/9891/787d9eea-850f-2996-0d6b-cc25a243409c.png)

非公開チャンネルに変更しても、今まで投稿されたメッセージは残ります。
ただし、チャンネルに貼られたファイルへのリンクURLなどは、非公開チャンネルに変更してもアクセス可能なままとなる点に注意が必要です。
また、非公開チャンネルを公開チャンネルに変更することはできません。

## モバイルアプリのアップデート

* 画像ダウンロード高速化
* チャンネルミュート(v4.9.0~)のサポート
* フラグを付けた投稿と@でメンションされた投稿へサイドバーのボタンからアクセス可能に
* 検索画面で投稿が日付で分割して表示されるように
* 無効化されたユーザーがメンバーリストから除外されるように

その他のアップデート内容については下記リンクから
https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#v180-release

## その他


### Mattermost v5.0

次回のリリース(2018/6/15)はMattermost v5.0のリリースになるようです。

![スクリーンショット 2018-05-19 20.30.50.png](https://qiita-image-store.s3.amazonaws.com/0/9891/d16b200f-55e7-cd3a-0925-751a0929a23e.png)

v5.0で行われる破壊的変更については、今の所下記にまとめられています。

[Upcoming Deprecated Features in Mattermost v5.0](https://docs.mattermost.com/administration/changelog.html#upcoming-deprecated-features-in-mattermost-v5-0)

* APIv3の廃止
* CLIツールとして配布されている`platform`コマンドの名称変更
* サイトURL設定の必須化
* デスクトップ通知持続時間設定の廃止
* スラッシュコマンド実行時のGETリクエスト送信方式の変更
* 自動リンク設定の追加(リンクとして認識されるプロトコル(現在は`http`,`mailto`など)をシステムコンソールから指定可能に)
* チーム削除APIのpermanentパラメータ無効化設定の追加
* Channelモデルから`ExtraUpdateAt`フィールドの削除
* [E20] CSVコンプライアンスエクスポートの刷新

### Mattercon

Mattermostの社員やコントリビューターを集め、ポルトガル・リスボンで行われたMattercon2018の報告が上がっています。

ハイライト
[MatterCon 2018 Community Meetup Highlights – Mattermost](https://about.mattermost.com/community/mattercon-2018-community-meetup/)

ハッカソン
[MatterCon 2018 Hackathon Highlights – Mattermost](https://about.mattermost.com/community/mattercon-2018-hackathon-highlights/)

パネルディスカッション
https://www.youtube.com/watch?v=3NWCIoys6vM&t=645s

### Mattercoin

毎月6日あたりから開始される新バージョンのRelease Candidateのテストに協力するとMattercoinというコインが貰えるようです。

[Earn a Mattermost Bug Hunter Coin \- YouTube](https://www.youtube.com/watch?v=7D6FJsdE_aY)

# おわりに

次回の`v5.0`のリリースは2018/6/15を予定しています。

