---
title: "Mattermost 9.0の新機能"
emoji: "🎉"
published: true
date: 2023-09-16T00:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/09/15 に Mattermost のメジャーアップデートとなる `v9.0.0` がリリースされました。  
2023/07/14に`v8.0.0`が出たばかりで、2ヶ月ほどでのメジャーバージョンアップですが、これはいくつかの主要機能を削除したことによる破壊的な変更という意味でのメジャーバージョンアップの位置付けかと思います。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v9\.0 offers new secure collaboration solutions in AI, Dev/Sec/ChatOps, and Zero Trust from partners](https://mattermost.com/blog/mattermost-v9-0-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v9-0-major-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体


## (Channels) チャンネル閲覧モーダルにチャンネル内のユーザー数が表示されるように

左サイドバーの **チャンネルを追加する > チャンネルを閲覧する** を選択することで開くチャンネル一覧を閲覧する画面が更新され、チャンネルに所属するユーザー数やチャンネルの説明が見れるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.0/channels-browse-channels.png)

## (Channels) カテゴリ全体を既読にするメニューの追加

チャンネルカテゴリ内に未読投稿が残っているチャンネルが存在する場合に、一括でカテゴリ内の全チャンネルを既読にするメニュー `Mark category as read` が追加されました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.0/channels-mark-category-1.png)

`Mark category as read` メニューをクリックすると、以下のようなダイアログが表示されます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.0/channels-mark-category-2.png)

`Mark as read`を選ぶと、選択したカテゴリ内の全チャンネルが既読状態になります。

## GIF画像サービスがGifycatからGIPHYに

Mattermostの投稿にGIF動画を貼る機能では、今までGifycatが使われていましたが、[Gifycatが2023/09/01にサービス終了する](https://techcrunch.com/2023/07/05/gfycat-shuts-down-on-september-1/)ことを受け、GIFサービスとして新たに[GIPHY](https://giphy.com/)が使われるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.0/channels-giphy.png)

過去のGifycatを採用していたバージョンのMattermostを使用している場合、GIF画像が何も表示されない状態となっています。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.0/channels-gfycat.png)


## その他の変更

### Boards Pluginを始めとした多くのプラグインのコミュニティへの移行

前バージョン(v8.1)リリース時も予告がありましたが、Mattermost公式チームによりメンテナンスされていたBoardsプラグインを含む多くのプラグインがコミュニティサポートに移行しました。引き続きプラグイン自体は利用できますが、Mattermost公式チームがバグフィックスを含むプラグインの更新に携わらなくなります。

コミュニティサポートに移行するプラグイン一覧は以下のページで確認できます。
[Upcoming product changes to Boards and various plugins \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/upcoming-product-changes-to-boards-and-various-plugins/16669)

Mattermost v6.0リリース時にMattermostのコア機能として追加されたBoards(Focalboard)も、コミュニティサポートに移行するプラグインに含まれます。  
Boardsが削除される理由としては、Boards機能のユースケースが限定的で、プロジェクト管理はJiraのようなツールで行われるのが一般的であり、さらにPlaybooks機能を使うことでよりよく解決できるはずだと説明されています。  
Boards機能についてはUIの自由度が高いために細かなバグが多い印象があり、メンテナンスコストも高かったのかなと想像しています。

### インサイト機能の削除

Mattermost v7.1でベータ版として追加されたインサイト機能が削除されました。  
[Platform: インサイト(ベータ版) - Mattermost 7\.1の新機能](https://blog.kaakaa.dev/post/mattermost/releases-7.1/#professionalenterprise-platform-%E3%82%A4%E3%83%B3%E3%82%B5%E3%82%A4%E3%83%88-%E3%83%99%E3%83%BC%E3%82%BF%E7%89%88)

DBへの負荷が高いのが原因だと見た記憶があります。

## その他のトピック

### MySQL非推奨化に向けた活動

Mattermostでは、バックエンドのDBとしてMySQLとPostgreSQLのどちらかを選択することができますが、Mattermost v11.0にMySQLの利用を非推奨とする予定があるそうです。  
以下のIssueの中にで一部説明がありますが、大規模な環境ではMySQLに信頼性とパフォーマンスの問題が起こりがちで、サポートコストがかかることが判断の一因とのことです。
[Support MySQL · Issue \#19 · mattermost/mattermost\-plugin\-ai](https://github.com/mattermost/mattermost-plugin-ai/issues/19)

先日、MySQLからPostgreSQLへの移行ガイドラインのドラフトが公開され、フィードバックを募集しています。
[Migration guidelines from MySQL to PostgreSQL](https://docs.mattermost.com/deploy/postgres-migration.html)

## おわりに
次の`v9.1`のリリースは 2023/10/16(Mon)を予定しています。  
