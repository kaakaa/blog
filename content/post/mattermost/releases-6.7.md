---
title: "Mattermost 6.7の新機能"
emoji: "🎉"
published: true
date: 2022-05-17T15:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/05/16 に Mattermost の新しいバージョン `v6.7.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.7 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v6-7-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-7-feature-release)

---

## アップグレード時の注意事項

今回のアップグレードではデータベースのスキーマ変更が行われるため、アップグレードに時間がかかる可能性があります。MySQLを利用していて900万投稿程度のデータが存在する場合、アップグレードに2分程度の時間がかかるようです(インスタンスは[db.r5.xlarge](https://aws.amazon.com/jp/rds/instance-types/)を使用)。

アップグレード時のダウンタイムを0にしたい場合、アップグレード前に以下のSQLを手動で実行し、事前にスキーマ変更を済ませておくこともできます。

* For MySQL: `CREATE INDEX idx_posts_create_at_id on Posts(CreateAt, Id) LOCK=NONE;`
* For Postgres: `CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_posts_create_at_id on posts(createat, id);`

これらのSQLによる変更は、テーブルロックを取得せず、既存のオペレーションに対する影響はないため、Mattermost起動中でも実行することができます。

詳しくは公式ドキュメントを参照してください。

[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html#important-upgrade-notes)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## (Professional/Enterprise) Playbooks: タスクの期限を設定可能に

Playbookに設定したチェックリスト内の各タスクに対して期限を設定できるようになりました。期限は 今日/明日/次週を選択するか、`At 5:00 pm in three days` のような自然言語(英語のみ)を使って設定することができます。

![playbooks-set-task-due](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.7/playbooks-set-task-due.png)

Playbooksの[日次ダイジェスト通知](https://docs.mattermost.com/playbooks/notifications-and-updates.html#daily-digest)には期限の迫ったタスクと期限の過ぎたタスクのみが通知されるようになるため、タスクの期限を設定することで通知ノイズを減らすことができます。

この機能は、**Professional**/**Enterprise**プラン限定の機能です。

## その他の変更

### Mattermost Desktop App: 自動アップデート機能

Mattermost Desktop Appに自動アップデートの機能が追加されました。  

Windows向けのMattermost Desktop Appを使用している場合、新しいバージョンが利用可能になると、アプリ上で更新を促すアラートが表示され、1クリックでDesktop Appのアップデートを完了することができます。

macOSでは、[GitHubのReleasesページ](https://github.com/mattermost/desktop/releases)からバイナリをインストールして利用している場合、自動アップデートは機能しないそうです。App Storeからインストールした場合のみ、新しいバージョンが自動で利用可能になるようですが、まだApp StoreではMattermost Desktop Appは公開されていないようで、公式ドキュメントからもApp Storeからのインストール手順が削除されています。App Storeで公開され次第、再度ドキュメントが更新されるのだと思います。

[Desktop App Automatic & Manual Updates by cwarnermm · Pull Request \#5628 · mattermost/docs](https://github.com/mattermost/docs/pull/5628/files#r872856733)

Linux OSでは、自動アップデートはサポートされていないようです。

詳しくは下記のドキュメントを参照ください。  

[Desktop App install guides](https://docs.mattermost.com/install/desktop-app-install.html)

### ChannelsのUI変更

ChannelsのUIにいくつか変更があります。

#### 右サイドバーのチャンネル情報に表示項目が追加

v6.6からチャンネル情報を右サイドバーで確認できるようになりましたが、表示項目に**チャンネルメンバー**、**ファイル**と**ピン留めされた投稿**が追加されました。

![channels-channel-info](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.7/channels-channel-info.png)

#### チャンネル作成ダイアログの変更

チャンネル作成ダイアログのUIが変更されました。

![channels-create-channel-dialog](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.7/channels-create-channel-dialog.png)

### チャンネル名が1文字のチャンネルが作成可能に

Mattermostのチャンネル名は2文字以上にする必要がありましたが、今回のバージョンから1文字のチャンネル名も作成できるようになりました。

[MM\-41909: Allow 1 char channel name by KevinSJ · Pull Request \#19845 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/pull/19845)

## その他のトピック

### Roadmap

Mattermost公式により運用されているコミュニティサーバー内の[**Roadmap**チャンネル](https://community-daily.mattermost.com/core/channels/roadmap)にて、今後のMattermostの方向性に関する情報が公開されています。

Mattermostのメインとなるチャット機能であるChannelsのロードマップについては、以下の資料が公開されています。  
[May 2022 \- Channels Roadmap Update \- Viewer Copy \- Google スライド](https://docs.google.com/presentation/d/1hsPHnB_Xsrc8mAq0T3VLoK22E9k3bS8-NtKSOxQ8oOY/edit)

Roadmapの資料によると、2022年6月にメジャーバージョンアップとなるv7.0のリリースを予定しているようです。  
v7.0では、現在ベータ版として提供されている**返信スレッドの折り畳み**機能がGA(Generic Available)機能に昇格するほか、WYSIWYG形式のメッセージ入力欄の導入等が行われる模様です。

![roadmap-v7](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.7/roadmap-v7.png)

また、モバイルアプリv2も開発中となっており、こちらはMattermost v7がリリースされる2022年6月にベータ版がリリースされる予定となっているようです。

![roadmap-mobile-v2](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.7/roadmap-mobile-v2.png)

モバイルアプリv2からは、一つのアプリに複数のMattermostサーバーを登録しておける機能などが追加される予定となっています。複数Mattermostサーバー登録機能については、以下のJIRAチケットのFigmaで作られたデザインイメージが参考になりそうです。(実際にリリースされる機能とは異なる可能性がありますが)

[\[MM\-41961\] Multi Server: All servers are logged out \- Mattermost](https://mattermost.atlassian.net/browse/MM-41961)


## おわりに
次の`v7.0`のリリースは 2022/06/15(Wed)を予定しています。

---
