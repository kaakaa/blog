---
title: "Mattermost 7.9の新機能"
emoji: "🎉"
published: true
date: 2023-03-17T23:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/03/16 に Mattermost のアップデートとなる `v7.9.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.9 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-9-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体

## (Boards) システム管理者/チーム管理者がBoardsにアクセス可能に

Mattermostインスタンスのシステム管理者とチーム管理者は、Boardの作成者が招待しなくとも、作成されたBoardにアクセスできるようになりました。  
これにより、使われなくなったBoardの削除等、Boardsの管理作業を行いやすくなります。  

システム管理者/チーム管理者は、公開Boardに`Boardを探す`メニューから、もしくはBoardのURLへの直接アクセスにより、そのBoardに参加することができます。非公開Boardには、BoardのURLに直接アクセスすると、Boardに参加するかどうかを確認するダイアログが表示されます。

---

まず、ユーザーがBoard作成を行うと、**非公開Board**として作成されます。  
作成されたBoardの**共有**メニューより、チームの全員に対する権限を設定することで、そのBoardは**公開Board**となります。

![boards-to-public](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.9/boards-to-public.png)

Mattermostのシステム管理者、もしくは作成されたBoardが属するチームのチーム管理者は、`Boardを探す`メニューから公開Boardを検索できるようになります。

![boards-search-public](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.9/boards-search-public.png)

検索されたBoardにアクセスすると、そのBoardの**管理者**として自動的に参加することができます。また、公開Boardには、BoardのURLを直接指定して参加することもできます。

非公開のBoardの場合、システム管理者であっても`Boardを探す`からは検索できず、BoardのURLから直接アクセスする必要があります。非公開のBoardにURLから直接アクセスすると、以下のようなダイアログが表示されます。

![boards-join-private](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.9/boards-join-private.png)

ここで、**Join**を選択すると、非公開BoardにそのBoardの**管理者**として参加することができます。

![boards-joined-private](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.9/boards-joined-private.png)

## (Channels) メッセージの編集履歴が確認可能に

投稿済みのメッセージを後から編集した場合、`編集済`の表示をクリックすることで、投稿者がメッセージの修正履歴を確認できるようになりました。ある特定のバージョンのメッセージを復元することもできます。

![channels-edit-history](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.9/channels-edit-history.png)

## アップグレード時の注意事項

アップグレード時の注意事項について、詳しくは公式ドキュメントを確認ください。　 
[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

### 新たなインデックス作成によるアップグレード時のDBテーブルのロック

今回のアップグレードで、DBの新たなインデックスとして `Posts(OriginalId)` が追加されます。  
1000万以上の投稿を持つMySQLを利用したMattermostサーバーの場合、i7-11800H(8コア/16スレッド)、32GBメモリの環境でインデックス作成に約100秒かかるそうです。

アップグレード時のテーブルロックを避けたい場合は、[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)を確認ください。


## おわりに
次の`v7.10`のリリースは 2023/04/14(Fri)を予定しています。
