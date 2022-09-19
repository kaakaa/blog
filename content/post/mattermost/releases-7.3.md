---
title: "Mattermost 7.3の新機能"
emoji: "🎉"
published: true
date: 2022-09-19T10:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---


Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/09/16 に Mattermost のアップデートとなる `v7.3.0` がリリースされました。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.3 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-3-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-3-feature-release)

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



## Boards: Roleベースの権限設定

今までBoards機能の各ボードの権限はMattermostのチャンネルに基づいて管理されていましたが、今回のバージョンからRoleベースの権限管理に変更されました。

ボード画面の右上に表示される **共有** ボタンから、ボードに対する権限設定を行うことができます。

![boards-share](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-share.png)

チーム全体に対する権限と、チャンネル・ユーザーごとの権限を設定できるようになっています。

ユーザーに対する権限は、ボードの内容を編集できる「**編集者(Editor)**」と、それに加えてボードの権限設定まで行うことができる「**管理者(Admin)**」から選択できます。また、今後のリリースでは、コメントの追加のみ可能な「**Commenter**」とBoardの閲覧のみ可能な「**Viewer**」の権限が追加される予定です。

Mattermostチャンネルのメンバー全員に権限を与えるには、ボードをMattermostチャンネルに**リンク**する必要があります。ボードとの**リンク**を行うと、そのMattermostチャンネルのメンバー全員にボードに対する **編集者(Editor)** 権限が付与されます。ボードは一つのMattermostチャンネルとリンクすることができ、Mattermostチャンネルとボードをリンクすると、チャンネル画面の右サイドバーにリンクされたボードの一覧が表示されるようになります。

![boards-rhs](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-rhs.png)

ボードの共有メニューでは、リンク済みのチャンネルに対する **リンク解除(Unlink)** を実行することができます。既にMattermostチャンネルにリンクされてるボードを別のMattermostチャンネルにリンクしようとすると、元々リンク済みだったチャンネルとのリンクが解除されるという警告メッセージが表示されます。

![boards-link-warning](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-link-warning.png)

Boardの権限や共有の設定について、詳しくは公式ドキュメントを参照してください。  
[Sharing | What’s new in Mattermost Boards](https://docs.mattermost.com/welcome/whats-new-in-v72.html#sharing)

## Boards: 新しいサイドバーナビゲーション

Boardsのサイドバーが刷新されました。

以前のバージョンでは、チャンネルごとにボードが設定されており、別のチャンネルにリンクされているボードを開くには、左上のメニューからチャンネルを切り替える必要がありました。  

![boards-sidebar-old](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-sidebar-old.png)

今回のバージョンからボードを開く際にチャンネルを選択する必要はなくなり、Mattermostチーム内の全てのボードを一つの画面で閲覧できるようになりました。また、Channelsのサイドバーと同様、`Ctrl+K` / `Cmd + K`のショートカットでチーム内のボードを検索することができ、さらに、カスタムカテゴリを作成してユーザー独自のグルーピングを行うことができるようになっています。

![boards-sidebar-new](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-sidebar-new.png)

新しいサイドバーについて、詳しくは公式ドキュメントを参照してください。  
[Navigation | What’s new in Mattermost Boards](https://docs.mattermost.com/welcome/whats-new-in-v72.html#navigation)

## Boards: カスタムテンプレートが簡単に共有可能に

以前のバージョンでは、Boardsのカスタムテンプレートを他のチームに共有しようとすると、一度Exportしてから別のチームでImportする必要がありました。

今回のバージョンから、Boardsテンプレート編集画面の **共有** メニューから、Mattermost Workspace(インスタンス)内全員に対してテンプレートを公開することができるようになりました。また、ユーザーを指定して公開することもできます。(チャンネルは選択できないようなので、プレースホルダのメッセージは間違っている気がします)

![boards-share-template](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/boards-share-template.png)

## Playbooks: サイドバーの刷新

Playbooks機能もサイドバーナビゲーションが刷新され、Mattermost Channels(チャット機能)と同様の見た目になりました。

以前のバージョン(下画像・右)では、プレイブックと実行(Run)がそれぞれ全画面で表示されてるため、それぞれの内容を確認するためにはタブを切り替える必要がありました。今回のバージョン(下画像・左)からは、全てのプレイブックと実行(Run)をサイドバーから一覧することができるようになり、プレイブック・実行(Run)へのアクセスが容易になりました。

![playbooks-sidebar](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/playbooks-sidebar.png)

今後、カスタムカテゴリによるグルーピング機能もリリース予定だそうです。


## Platform: インサイト機能の改善

インサイトに表示される項目が追加されました。

![platform-insight](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.3/platform-insight.png)


「**新しいチームメンバー**」では、Mattermostチームに新たに参加したユーザーを確認することができます。`チームインサイト`では「**新しいチームメンバー**」が表示されますが、自分に関するインサイトである`私のインサイト`画面では、同じ場所に「**もっともアクティブなダイレクトメッセージ**」が表示され、DMのやり取りの回数の多いユーザーが表示されます。  
「**最もアクティブでないチャンネル**」では、活動の少ないチャンネルを確認することができます。活動の少ないチャンネルの情報は、再び投稿を作成して注意をひくよう計画したり、チャンネルを脱退するなどの判断に利用することができます。  
「**トップPlaybooks**」では、最も実行回数の多いプレイブックを確認することができます。


## その他の変更

### (Enterprise) Channels: Call機能専用のCallサーバーが利用可能に

Calls機能専用の[`rtcd`(read-time communication daemon)](https://github.com/mattermost/rtcd)サーバーが公開されました。MattermostサーバーからCalls機能の機能を`rtcd`サーバーに移譲することで、よりスケーラブルな通話環境を構築することができるようになります。

詳しくは、公式ドキュメントを参照してください。  
[Calls self\-hosted deployment](https://docs.mattermost.com/configure/calls-deployment.html#calls-self-hosted-deployment)

## アップグレード時の注意事項

### Boardsのマイグレーション

本バージョンから、Boards機能はチャンネルベースの権限管理からRoleベースの権限管理に移行します。Boardsごとの権限のマイグレーションはMattermostアップデート時に自動で行われますが、アップデート前にバックアップを実施することをお薦めします。

Boards機能のアップデート内容について、詳しくは[公式ドキュメント](https://docs.mattermost.com/welcome/whats-new-in-v72.html)を参照してください。

## その他のトピック

### Mattermost関連の記事

先月に引き続き、Slackの料金改定に端を発するMattermostへの移行の流れの中で、Mattermostに関する記事を見かけることが多かったため、目にした範囲で紹介します。

#### SlackからMattermostへの移行
* [MattermostをUbuntu環境にインストール \- Qiita](https://qiita.com/N-H-Shimada/items/0ee996112340030e7959)
* [Slack から Mattermost へデータ移行 \- Qiita](https://qiita.com/takaoS/items/121700e671e6e7710efa)
* [GCP無料枠を使ってSlackからMattermostへ移行してみた｜あらB｜note](https://note.com/arkb/n/nedecd6c170f5)

#### SlackからMattermostへの移行サポート

* [SlackからMattermostへの移行作業のサポートについて｜PressWalker](https://presswalker.jp/press/3088)

#### Mattermost構築
* [【Ubuntu】Mattermost をインストール \- Qiita](https://qiita.com/takaoS/items/48f440b3c11f2e18a661)
* [MattermostをGCP無料枠で試す \- Qiita](https://qiita.com/data-dokata/items/bc8b73ccca1fce82a6cf)
* [MattermostをUbuntu20\.0\.4 サーバにインストールする手順 \- Qiita](https://qiita.com/kanetugu2018/items/51cdab279d81ae06aa70)

#### Mattermost Integrations

* [GitHub ActionsからMattermostにCommit内容を通知する方法 \| UmentuLab\(Shintaro Tachibana\)](https://weblog.umentulab.com/post/tech/sre/notify_mattermost_from_github_actions/)

## おわりに
次の`v7.4`のリリースは 2022/10/14(Fri)を予定しています。

