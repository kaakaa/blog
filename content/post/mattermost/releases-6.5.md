---
title: "Mattermost 6.5の新機能"
emoji: "🎉"
published: true
date: 2022-03-19T12:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/03/16 に Mattermost の新しいバージョン `v6.5.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.5 is now available \- Upgrade your deployment today](https://mattermost.com/blog/mattermost-v6-5-is-now-available/)
- [Mattermost Self\-Hosted Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-5-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

CLIコマンド`mattermost version`の実行結果にデータベースのバージョンが含まれなくなりました（今までのバージョンでは実行結果に`DB Version: 6.5.0`のようにDBのバージョンも出力されていました）。また、`mattermost version`がDBとのやり取りを行わなくなり、DBマイグレーションも実行しなくなったため、DBマイグレーションを行うための新たなコマンド `mattermost db migrate`が追加されました。[(PR)](https://github.com/mattermost/mattermost-server/pull/19364)


---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## (Enterprise/Professional) [Channels: カスタムグループ (ベータ版)](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#groups)

複数のユーザーにメンションを送信する際に利用できる**カスタムグループ**を作成できるようになりました。例えば、`developer`というグループを作成し、そのグループにユーザーを追加しておくと、`@developer`というメンションで、グループ内の全員にメンションすることができます。　  
今までのバージョンでもAD/LDAP連携を設定している場合は[LDAP上で設定されたグループに対してメンションを送る機能](https://docs.mattermost.com/channels/mention-people.html#groupname)はありましたが、今回追加されたカスタムグループの機能によりAD/LDAP連携を行なっていなくてもグループメンションが利用できるようになりました。

カスタムグループの作成はMattermost UIのProduct Menuから行うことができます。操作概要については、公式リリースブログの `Channels: Custom groups (Beta)` セクションの動画を参照ください。

[Mattermost v6\.5 is now available \- Upgrade your deployment today](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#groups)

また、スラッシュコマンド`/invite`で、ユーザーをグループに追加することもできます。

カスタムグループ機能について、詳しくは以下のドキュメントを参照ください。

[Manage Custom Groups](https://docs.mattermost.com/welcome/manage-custom-groups.html)


## [Channels: チーム横断的なチャンネル移動が可能に](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#cross)

`Ctrl + k` のショートカットで開くことができるチャンネル検索ダイアログで、他のチームのチャンネルも横断的に検索できるようになりました。
今までのバージョンでは、チャンネル検索ダイアログで検索できるチャンネルは現在アクセスしているチームのチャンネルのみに限定されていましたが、今回のバージョンより参加したことのあるチャンネルであればチームに関係なく検索することができるようになります。

![channels-cross-team](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/channels-cross-team.png)

検索結果のチャンネル名の右側にチーム名が表示されるため、複数のチームに同名のチャンネルが存在しても見分けることができます。

## [Playbooks: 複製、インポート、エクスポート](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#playbooks)

既存のPlaybookを複製することができるようになりました。  

また、Playbook内容をJSON形式でエクスポートできるようになり、Playbookのバックアップや、インポート機能と組み合わせて別のMattermostインスタンスへのPlaybookの移行などが実施できるようになりました。

![playbooks-duplicate](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/playbooks-duplicate.png)

`複製`メニューをクリックすると、`Copy of ${Playbook名}` という名前のPlaybookが新たに作成されます。


## [Boards: 共有機能のUI改善](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#share)

Boardsの共有リンクを取得する際のUIが改善されました。

まず、Board画面に共有（Share）ボタンが表示されるようになりました。（今までは、オプションメニュー内に表示されていました）  

![boards-share](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/boards-share.png)

また、リンクを生成する際に、Mattermostユーザー向けの内部リンクを作成する画面と、MattermostユーザーアカウントがなくてもBoardの内容が確認できる読取専用の公開リンクを作成する画面が別画面になりました。

![boards-share-internal](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/boards-share-internal.png)

![boards-share-publish](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/boards-share-publish.png)

## [Boards: チャンネルイントロにBoardへのリンクが表示されるように](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#intro)

チャンネル作成時に一番最初に表示されるテキストにBoardへのリンクが含まれるようになりました。

![boards-intro-link](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/boards-intro-link.png)

## [Boards: インポート機能のドキュメントへのリンク](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#help)

Trello・Jira・NotionなどのツールからBoardsへデータをインポートする際の手順に関するドキュメントへのリンクが追加されました。  
Boardsの設定メニューから表示することができます。

![boards-import-link](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/boards-import-link.png)

他ツールからBoardsへのインポート機能について、詳しくは以下のリンク先を参照してください。

[Import your data](https://docs.mattermost.com/boards/data-and-archives.html#import-your-data)


## Integration: 統合機能

Mattermostと他機能を連携させる統合機能について、いくつかのアップデートがあります。

### [Atlassian Bitbucket Cloud Plugin](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#bitbucket)

Atlassian社のGitリポジトリ管理ツールであるBitbuket CloudとMattermostを連携させる統合機能がリリースされています。

毎日Mattermostへログインした際にBitbucket　Cloud内で行われた重要な活動が通知される機能や、Bitbucket Cloud内で自分がメンションされたことをMattermostのDMで通知してくれる機能などが利用できるようになります。  また、PRのリストが常にMattermostのサイドバーに表示されるようになります。

インストール方法など、詳しくは以下のリンク先を参照してください。  
[Mattermost/Bitbucket Plugin \- BitBucket Plugin](https://mattermost.gitbook.io/bitbucket-plugin/)

### [Configuration Wizard](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#configuration)

いくつかの統合機能は、インストール後にセットアップ手順を実行する必要があります。
いくつかの統合機能について、このセットアップ手順がプラグインインストール時にMattermostのDMとして投稿されるようになりました。これにより、外部サイトを参照しながらセットアップを行う必要がなくなります。([Mattermost GitHub Plugin v2.1.0](https://github.com/mattermost/mattermost-plugin-github/releases/tag/v2.1.0)で試してみましたが、うまく動作しませんでした)


### [GitHub Plugin v2.1](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#github)

Mattermost GitHub Plugin v2.1がリリースされました。

MattermostのメッセージからGitHubにIssueを作成する際に利用できる `Create Issue in GitHub` メニューなどが追加されています。

![integrations-github](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/integrations-github.png)

変更点について詳しくはリリースノートを参照してください。

[Release v2\.1\.0 · mattermost/mattermost\-plugin\-github](https://github.com/mattermost/mattermost-plugin-github/releases/tag/v2.1.0)


## Platform

Mattermostインスタンス全体 (Platform) についても、いくつか更新点があります。

### [ワークスペースの最適化](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#workspace)

運用中のMattermostインスタンスが最適な設定で運用されているかどうかを判定し、最適でない場合にどのように変更すべきかを提案するダッシュボード機能が追加されました。
**システムコンソール > ワークスペースの最適化**から確認することができます。（手元の環境ではメニューが表示されませんでした...）

詳しくは、以下の公式ドキュメントを参照してください。
[Optimize Your Mattermost Workspace](https://docs.mattermost.com/configure/optimize-your-workspace.html?highlight=optimization)

## [Onboarding tourの改善](https://mattermost.com/blog/mattermost-v6-5-is-now-available/#onboarding)

初めてMattermostにアクセスした際に表示されるオンボーディングが、Board/Playbookでも表示されるようになりました。  
Boards/Playbooks機能の利用方法をチュートリアル形式で確認していくことができます。

![platform-onbording](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.5/platform-onbording.png)


## その他の変更
* Firefox の利用推奨最低バージョンが`v78`から`v91`に変更されました [（PR）](https://github.com/mattermost/mattermost-server/pull/19271)
* Safari の利用推奨最低バージョンが`v12`から`v14.1`に変更されました [（PR）](https://github.com/mattermost/mattermost-server/pull/19564)

## その他のトピック

### Go Conference JP 

2022/04/23に日本で(オンライン)開催される Go Conference Online 2022 SpringにMattermost社の[Jesús Espino](https://github.com/jespino/)氏が登壇します。

[Go Conference 2022 Spring \| Dissecting Slices, Maps and Channels in Go](https://gocon.jp/2022spring/sessions/a3-l/)

Jesus氏はMattermostの開発者ブログでも、Goに関するエントリをいくつか書いています。

- [Maintaining Consistency in Codebases with Go vet](https://developers.mattermost.com/blog/maintaining-consistency-in-codebases-with-go-vet/)
- [Layered Store and Struct Embedding in Go](https://developers.mattermost.com/blog/layered-store-and-struct-embedding/)

今回の発表もGo言語のランタイムにおいて、Slice, Map, Channelがどのように動作しているかというディープな話になりそうです。

### Boardsのコメント欄に日本語を入力できない問題

Boardsのカード内のコメントに日本語を入力しようとすると、Enterキーで入力を確定したタイミングで入力内容が消えてしまう事象が発生しています。  

この問題については、以下のIssueで対応が進められています。
[Bug: Unable to enter Japanese text into comments field · Issue \#2343 · mattermost/focalboard](https://github.com/mattermost/focalboard/issues/2343)

### ロシア/ベラルーシを輸出禁止国に追加

米国政府による制裁に応じ、Mattermost社もロシア・ベラルーシを輸出禁止国のリストに追加したそうです。

[Mattermost policy changes due to conflict in Ukraine \- Mattermost](https://mattermost.com/blog/mattermost-policy-changes-due-to-conflict-in-ukraine/)

## おわりに
次の`v6.6`のリリースは 2022/04/14(Thu)を予定しています。

