---
title: "Mattermost 6.6の新機能"
emoji: "🎉"
published: true
date: 2022-04-23T13:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/04/16 に Mattermost の新しいバージョン `v6.6.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.6 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v6-6-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-6-feature-release)


---

## アップグレード時の注意事項

Mattermostと他システムを連携する仕組みの一つである[Mattermost Apps Framework](https://developers.mattermost.com/integrate/apps/)のプロトコルの一部に破壊的な変更があります。今までのバージョンで`binding`や`form`のリクエストを送信する際に`call`要素として設定していた値が、`submit`,`form`,`refresh`,`lookup`要素に分割されました。以前のバージョンで動作するよう開発していたAppsは、新しいプロトコル向けに修正する必要があります。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## Channels: トリガーとアクション

ユーザーがMattermost上で特定の行動(トリガー)を行った場合に、自動で実行されるアクションをMattermost上で定義できるようになりました。対象となるのは後述する2つのトリガー/アクションのみですが、コードを書くことなく、Mattermost上の操作だけで設定することができます。

(本機能はMattermost Plugin Playbooks v1.27により追加される機能だと思われます。そのため、Playbooksプラグインのバージョンが古い、Playbooksプラグインが有効になっていない等の場合は後述の**Channel Actions**を追加するメニューが表示されません。メニューが表示されない場合は、**システムコンソール > プラグイン管理 > インストール済みプラグイン**からPlaybooks v1.27が有効になっていることを確認してみてください。)

### 設定方法

トリガーとアクションはチャンネルごとに **チャンネルメニュー > Channel Actions** から設定することができます。設定するにはチャンネル管理者の権限が必要です。

![channels-channel-actions](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-channel-actions.png)


### トリガー1: ユーザーがチャンネルに参加した時

"**ユーザーへ一時的なウェルカムメッセージを送信する**"を設定することで、新たにユーザーがチャンネルに参加した際に、そのユーザーのみが閲覧できるメッセージを表示することができます。新たにチャンネルに参加したユーザーに目を通しておいて欲しい情報や、チャンネルの運用ルール等を自動で伝えることができます。  
また、"**ユーザーのサイドバーカテゴリにチャンネルを追加する**"から、参加したチャンネルを左サイドバーのどのカテゴリに追加するかを指定することもできます。参加したユーザーの左サイドバーに追加先のカテゴリが存在しない場合は自動で作成されます。

![channels-actions-join](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-actions-join.png)

このチャンネルアクションを設定したチャンネルに参加すると、以下のように左サイドバーの指定したカテゴリにチャンネルが表示され、ウェルカムメッセージが表示されます。(左サイドバーのカテゴリに移動するまでに多少時間がかかることがあります。)

![channels-actions-join-example](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-actions-join-example.png)

### トリガー2: 特定のキーワードを含むメッセージを投稿した時

"**これらのキーワードを含むメッセージが投稿された場合**"で設定したキーワードを含むメッセージが投稿された際に、Playbookを開始するかどうかを尋ねるダイアログを自動で表示することができます。

![channels-actions-post](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-actions-post.png)

この機能を使うことで、例えば「システムのモニタリングツールが`priority: high;`というキーワードを含むメッセージをMattermostに投稿すると、SREチームがインシデント対応手順を即座に開始する」というようなフローを組むことができるようになります。

![channels-actions-post-prompt](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-actions-post-prompt.png)


## Channels: メッセージアクションの表示位置変更

PluginやAppsによりメッセージのコンテキストメニューに独自のメニューを追加することができますが、このメニューの表示位置が変更されました。  
今までのバージョンでは、Plugin/Appsによって追加されたメニューもMattermostデフォルトのメッセージアクション（編集、削除など）と並列に表示されていましたが、本バージョンからはカスタムアクション専用のボタンが表示されるようになりました。

![channels-message-action](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-message-action.png)


## (Enterprise) Playbooks: レトロスペクティブメトリクス

Playbookの各実行完了後に、任意のメトリクスを4つまで入力できるようになりました。入力したメトリクス値の統計情報を自動で計算・表示してくれるため、インシデント対応のパフォーマンス指標などに利用することができます。  

(Enterprise限定機能であるかのように説明されていますが、Team Editionでも利用できそう？)

### 設定
まず、Playbookに関するメトリクスを収集するために、収集対象のキーメトリクスを設定する必要があります。キーメトリクスとは、Playbook編集画面の`レトロスペクティブ`タブから設定することができます。

![playbooks-metrics-settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/playbooks-metrics-settings.png)

### メトリクス値の入力
Playbookにキーメトリクスを設定しておくと、Playbookの実行が完了した後に入力を促されるレトロスペクティブ編集画面で、設定したキーメトリクスの入力欄が表示されるようになります。

![playbooks-metrics-input](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/playbooks-metrics-input.png)

### メトリクス統計情報の確認

キーメトリクスを入力してレトロスペクティブを発行すると、Playbook概要ページの`キーメトリクス`タブから、入力したメトリクスの統計値や推移を確認できるようになります。

![playbooks-metrics-view](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/playbooks-metrics-view.png)

## その他の変更

### Integrations: App FrameworkがGeneral Availabilityに

開発者プレビュー版として公開されていた[Mattermost Apps Framework](https://developers.mattermost.com/integrate/apps/)がGeneral Availabilityになりました。
App Frameworkは、Plugin機能に近い拡張性を持ち、どんなプログラミング言語でも記述でき、さらにサーバーレス技術を使ったホスティングも可能なMattermostの統合機能です。

今までのバージョンでは、App Frameworkを利用するには[mattermost-plugin-apps](https://github.com/mattermost/mattermost-plugin-apps)というプラグインを自身でアップロードする必要がありましたが、このv6.6からは全てのMattermostサーバーで自動で有効になります。

Mattermost Apps Frameworkについて日本語で書かれている記事は、以下のようなものがあります。  
(Mattermost Apps Frameworkの開発者向けプレビュー版が公開された当時に書かれたもののため、記述内容が古くなっている可能性があります。)

* [Mattermost Apps Frameworkを触ってみた](https://zenn.dev/kaakaa/articles/mattermost-apps-sample)
* [Mattermost Apps Framework をJava \(JAX\-RS\)で試してみた – maruTA\(Bis5\)'s Weblog – Side T:echnology](https://tech.bis5.net/2021/05/09/248.html)

### チャンネル情報が右サイドバーに

チャンネル情報を表示するためのボタンがチャンネルヘッダ部分に追加されました。  ボタンをクリックすると、右サイドバーにチャンネル情報が表示されます。

![channels-channel-info](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.6/channels-channel-info.png)

## おわりに
次の`v6.7`のリリースは 2022/05/16(Mon)を予定しています。

---
