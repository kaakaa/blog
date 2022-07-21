---
title: "Mattermost 6.4の新機能"
emoji: "🎉"
published: true
date: 2022-02-19T17:40:00+09:00
toc: true
tags: ["mattermost", "releases"]
---
Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/02/16 に Mattermost の新しいバージョン `v6.4.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.4 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v6-4-is-now-available/)
- [Mattermost Self\-Hosted Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-4-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

本バージョンからバージョンアップ時のDBスキーママイグレーション方法が新しくなったため、アップグレード前にバックアップを取っておくことを強く推奨します。
新しいマイグレーション方法では、マイグレーションの記録をDBに保存しておくために、すべてのマイグレーションが実行されます。マイグレーションの実行結果は `db_migrations` テーブルに記録されます。
また、マイグレーション処理が並列に行われていないようにするために `db_lock` テーブルを利用します。

もしマイグレーション処理にエラーが発生した場合は、この2つのテーブルを確認してください。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## Channels: Mattermost Cloudからの新しい移行方法

Mattermost Cloud上のデータをセルフホストのMattermostインスタンスへ移行する場合の新しいプロセスが公開されました。`mmctl`コマンドを利用したデータ移行方法となっており、セルフホストのMattermostインスタンスのデータをMattermost Cloudへ移行する方法も今後同じページで公開されるようです。

[Mattermost Workspace Migration](https://docs.mattermost.com/manage/cloud-data-export.html)

また、Mattermost Cloudへの移行をサポートするツールとして、新たにAWAT (= Automatic Workspace Archive Translator) というツールが公開されています。  
https://github.com/mattermost/awat

これは、セルフホストのMattermostだけでなく、Slackのワークスペース情報などもMattermost Cloudのデータ形式に翻訳した上で移行するなど、Mattermostインスタンスの移行をサポートするツールのようです。

## Playbooks: Team/Starter EditionでもPlaybookを無制限に作成可能に

Mattermost v6.0でMattermostの機能として提供されるようになったPlaybooksの機能ですが、無償版のユーザーはPlaybookを一つしか作成できないよう制限されていました。
しかし、ユーザーからのフィードバックを受け、v6.4からは無償版のユーザーでもPlaybooksを無制限に作成できるようになりました。

この変更を受けてか、Playbooks機能を利用した事例に関する記事執筆を求めるWritingキャンペーンを開始しています。

[Get Paid to Write About Mattermost Playbooks \- Mattermost](https://mattermost.com/blog/write-about-mattermost-playbooks/)

## Boards: テンプレートプレビュー

Boardを新規に作成する際にBoardのテンプレートを選択することができますが、このテンプレートを選択する画面が新しくなり、テンプレートの内容をプレビューできるようになりました。

![boards-template](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.4/boards-template.png)

また、デフォルトで選択できるBoardテンプレートの内容が刷新され、さらに新しいテンプレートも追加されています。

## Boards: 画像を含む新たなアーカイブ形式

Boardの内容をエクスポートした際に、カードに添付された画像データなどもエクスポートファイルに含まれるようになりました。
また、エクスポートファイルの拡張子が `.focalboard` から `.boardarchive` に変更されました。古い `.focalboard` という拡張子でエクスポートされたデータも現時点ではインポートすることが可能ですが、今後のバージョンでインポートできなくなる予定です。

## Boards: カードバッジ

カード内に説明やコメント、チェックリストなどが存在する場合、Board画面でカードバッジとして表示されるようになり、カードを開くことなく概要を確認することができるようになりました。  
![boards-card-badge](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.4/boards-card-badge.png)

カードバッジはBoardの**プロパティ > Comments and Description** で表示/非表示を切り替えることができます。  
![boards-show-badge](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.4/boards-show-badge.png)

## Boards: URLプロパティの改善

URLプロパティを設定した場合、そのURLリンクを開くためには小さなリンクアイコンをクリックしなけれなりませんでしたが、本バージョンからURL文字列全体がリンクとしてクリックできるようになりました。  
![boards-url](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.4/boards-url.png)

## Boards: カードの説明にGIFが利用可能に

カードの概要にアニメーションGIFを利用することできるようになりました。

![boards-agif](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.4/boards-agif.gif)

## その他の変更
* **アカウント設定**というメニューは無くなり、**設定**という名前になりました

## その他のトピック

### ChatOps導入ガイドの公開

Mattermostから、エンタープライズチームにChatOpsを導入するための7つのステップが公開されました。

[7 Steps to ChatOps for Enterprise Teams \- Mattermost](https://mattermost.com/chatops-guide/#step-6-build-smarter-bots-and-workflows)

Mattermostの機能や実際の事例と絡めながら、ChatOpsを組織に導入するステップが説明されています。

## おわりに
次の`v6.5`のリリースは 2022/03/16(Web)を予定しています。

