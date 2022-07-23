---
title: "Mattermost 7.1の新機能"
emoji: "🎉"
published: true
date: 2022-07-20T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/07/15 に Mattermost のアップデートとなる `v7.1.0` がリリースされました。また、同日に[Marketplace関連の問題](https://mattermost.atlassian.net/browse/MM-45731)を修正した `v7.1.1` がリリースされています。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.1 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v7-1-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-1--extended-support-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- Platform: 上記機能を含むMattermost全体

## (Professional/Enterprise) Platform: インサイト (ベータ版)

ワークスペース内の利用状況や活動状況を一目で確認できる[**インサイト機能**](https://docs.mattermost.com/welcome/insights.html)が追加されました。(有償版限定の機能です。)

![channels-insight](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.1/channels-insight.png)

インサイトは、自分が参加しているチャンネル・スレッドのみの情報を集めたインサイト(**私のインサイト**)と、チーム全体に関するインサイト (**チームインサイト**)の2つから選んで表示することができます。

![channels-insight-menu](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.1/channels-insight-menu.png)

また、インサイトの集計期間は、**今日**、**直近7日間**、**直近28日間**から選ぶことができます。

![channels-insight-timeframes](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.1/channels-insight-timeframes.png)

各ウィジェットのタイトル部分をクリックすると、トップスレッドは上位5スレッド、トップボードとトップリアクションは上位10項目まで確認することができます。

![channels-insight-widget](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.1/channels-insight-widget.png)

**チームインサイト**の**トップスレッド**には、自分が参加していないチャンネルのスレッドも表示されるため、会話が盛り上がっているチャンネルを見つけることなどができるようにもなりそうです。

## Channels: Mobile v2.0(ベータ版)の開発に伴う各種更新

今回のリリースには、先月紹介した[Mobileアプリ v2.0のベータ版](https://blog.kaakaa.dev/post/mattermost/releases-7.0/#channels-mobile-v20-beta)の開発に伴う更新が含まれています。Mobile v2.0(Beta)は、機能追加やバグ修正が毎週行われており、ユーザーからのフィードバックも歓迎しています。正式版のリリースは今年末を予定しているそうです。

Mobileアプリ v2.0のベータテストに参加する方法は、以下のブログエントリに説明があります。  
[Join the Mattermost mobile beta program: v2\.0 is now live\! \- Mattermost](https://mattermost.com/blog/mobile-beta-program-v2/)


## アップグレード時の注意事項

### 新たに追加された`MaxImageDecoderConcurrency`の設定値について

新たに追加された画像デコード同時実行数に関する設定 MaxImageDecoderConcurrency にはデフォルトで -1 が設定されており、これはCPUと同数であることを表しています。

この設定は、Mattermostサーバーのメモリ消費量に影響があり、設定値の決定に注意が必要です。  
まず、Mattermostで単一の画像を処理する際の最大メモリ消費量は、最大画像解像度の設定MaxImageResolutionの設定値を使用して MaxImageResolution * 24バイトと計算することができます。さらに、本バージョンより追加されたMaxImageDecoderConcurrencyによる画像処理の同時実行数を考慮すると、MaxImageResolution * MaxImageDecoderConcurrency * 24バイトが、Mattermostに割り当てられたメモリより小さくなるよう設定値を決定することをお勧めします。

### スキーマ変更にかかる所要時間

Mattermost 7.1にはデータベースのスキーマ変更が存在します。アップグレード時にスキーマ変更にかかる時間については以下のテスト結果を参考にしてください。

* MySQL, 投稿数1200万, リアクション数250万 → ~1分34秒 (8 core CPU / 16GB RAM)
* PostgreSQL, 投稿数1200万, リアクション数250万 → ~1分18秒 (db.r5.2xlarge)

以下のSQLを実行することでアップグレード前に手動でスキーマ変更を行うことができますが、このクエリは`Reaction`テーブルへのロックを取得するため、SQL実行中にユーザーが投稿したリアクションはデータベースに反映されません。

* MySQL

```sql
ALTER TABLE Reactions ADD COLUMN ChannelId varchar(26) NOT NULL DEFAULT "";

UPDATE Reactions SET ChannelId = COALESCE((select ChannelId from Posts where Posts.Id = Reactions.PostId), '') WHERE ChannelId="";

CREATE INDEX idx_reactions_channel_id ON Reactions(ChannelId) LOCK=NONE;
```

* PostgreSQL

```sql
ALTER TABLE reactions ADD COLUMN IF NOT EXISTS channelid varchar(26) NOT NULL DEFAULT '';

UPDATE reactions SET channelid = COALESCE((select channelid from posts where posts.id = reactions.postid), '') WHERE channelid='';

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_reactions_channel_id on reactions (channelid);
```

## おわりに
次の`v7.2`のリリースは 2022/08/16(Tue)を予定しています。

