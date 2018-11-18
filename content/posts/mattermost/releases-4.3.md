
---
title: "Mattermost4.3の新機能"
date: 2017-10-23T15:41:03+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

2017/10/15にMattermost 4.3.0がリリースされたので、追加された機能などについて簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4\.3: Tablet support, E20 data retention\. Plus: join us for Hacktoberfest\! \- Mattermost](https://about.mattermost.com/blog/mattermost-4-3/)
[Mattermost Changelog — Mattermost 4\.3 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-3-0)

## セキュリティアップデート

[Mattermost Changelog — Mattermost 4\.3 documentation](https://docs.mattermost.com/administration/changelog.html#security-update)
Mattermost v4.3.0は複数のセキュリティアップデートを含んでいるため、早期のアップデートを推奨しています。
Mattermostのポリシーに基づき、脆弱性の内容についてはリリースから14日後に [Security Updates \- Mattermost](https://about.mattermost.com/security-updates/) で公開されるそうです。

## Team Edition
(Team EditionはOSS版のことです)

### モバイルアプリ v1.3 の改善点

#### タブレット向け横向き画面

タブレットデバイスに適した横向きのビューが追加されました（ベータ版）

手元のiPhone7 Plusで試したところ、こんな感じになりました。
![image.png](https://qiita-image-store.s3.amazonaws.com/0/9891/4a525d84-feb7-0bdb-15f8-43f62668f829.png)


#### リンクプレビュー

画像やGIF、YouTubeのプレビュー表示が追加されました

#### 通知

Androidユーザーは通知に関するライト、振動、サウンドの設定ができるようになりました。
また、最新の通知が一番最初に表示されるよう改善が行われました。

## Enterprise Edition (E20)

### データ保持期間設定の追加

Mattermost上のメッセージやファイルを保持する期間を設定できるようになりました。
[Data Retention Policy Beta \(E20\) — Mattermost 4\.3 documentation](https://docs.mattermost.com/administration/data-retention.html)

また、次のバージョン(v4.4.0)では、メッセージを削除する前にサードパーティのアーカイブシステム向けにエクスポートする機能が追加される予定だそうです。
これにより、監査のためにメッセージを保存しておくこともできるとのことです。

### タイムアウト設定の追加

ユーザーが指定した時間(分単位)、システムを利用していない場合に自動でログアウトさせることができるようになりました。

## New Community Integrations

### Spotify-Deezer

[Kaylleur/mattermost\-integration\-deezer\-spotify\-link](https://github.com/Kaylleur/mattermost-integration-deezer-spotify-link)
SpotifyとDeezer上のミュージックトラックへのリンクをMattermostへ投稿するBotアプリのようです。

## Hacktoberfest

毎年恒例になってきた「10月中にGitHub上でコントリビュート行うとTシャツがもらえるイベント」の紹介です。
[Hacktoberfest 2017 \- DigitalOcean](https://hacktoberfest.digitalocean.com/)

MattermostではHacktoberfest向けに、HelpWanted的なIssueに`Hacktoberfest`のラベルを付けて参加を奨励しています。
[Hacktoberfest 2017 \- DigitalOcean](https://hacktoberfest.digitalocean.com/)

昨年はこんな感じのTシャツがもらえました。
![image.png](https://qiita-image-store.s3.amazonaws.com/0/9891/27bc0e10-2001-24c0-9462-a8821da2e095.png)


## おわりに

次回のv4.4.0のリリースは2017/11/16を予定しています。


