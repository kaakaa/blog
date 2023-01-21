---
title: "Mattermost 7.7の新機能"
emoji: "🎉"
published: true
date: 2023-01-20T23:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/01/16 に Mattermost のアップデートとなる `v7.7.0` がリリースされました。  
また、2023/01/20に`v7.7.1`がリリースされています。`v7.7.1`では、`v7.7.0`からMattermost使い始めた場合に最初のシステム管理者のユーザーアカウントを作成できない問題や、インサイト機能でトップチャンネルの結果が表示されない問題に対応されています。

2022/12/16にリリース予定だった`v7.6.0`は、パフォーマンスに関する問題を理由にリリースが中止されていたため、2ヶ月ぶりのリリースです。

> The Mattermost v7.6 release has been cancelled as we are working on investigating performance issues. The next scheduled release is v7.7 in January 16th, 2023.

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.7 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v7-7-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-7-feature-release)

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


## Calls: General Availability

2022/06にリリースされたMattermost v7.0でベータ版として追加されたCalls機能がGeneral Availabilityとなりました。

先々月のリリースからのアップデートとして、通話中にリアクションを送ることができるようになりました。リアクションをクリックすると、ユーザー名と共に選択したリアクションが画面左に表示され、一定時間経過すると消えていきます。

![calls-reaction](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/calls-reaction.png)

また、Mattermost Enterpriseプラン限定の機能ですが、通話の記録を動画ファイルとして残す機能も追加されています。  
[Record calls](https://docs.mattermost.com/channels/make-calls.html#record-calls)

Calls機能は[Mattermost Cloud](https://mattermost.com/sign-up)を使えばすぐにCalls機能を利用することができますが、Cloud Freeプランでは参加者の上限が8名に制限されています。Cloud Professional/Enterprプランでは200名が上限です。

## Channels: Mattermost モバイルアプリ v2

Mattermostのモバイルアプリのメジャーバージョンアップである v2.0 がリリースされました。

このリリースには、モバイルアプリに対して[最も要望の多かった](https://mattermost.uservoice.com/forums/306457-general/suggestions/10975938-ios-and-android-apps-should-allow-multiple-server)複数サーバーの管理機能が追加されています。

YouTubeで新しいモバイルアプリの紹介動画が公開されています。

https://www.youtube.com/watch?v=YPFfXISvydk

Mattermost Mobile v2.0にアップデートする場合、サポート対象のMattermost Serverがv7.1以降に設定されているため、事前にMattermost Serverのバージョンを確認しておいたほうが良いかと思います。  
[Preparing for Mobile v2\.0 \- Mattermost](https://mattermost.com/blog/preparing-for-mobile-v2-0/)

> What server versions are supported by v2.0?   
> 
> Mattermost Server v7.1+ is required.

## Channels： メッセージの優先度設定と既読確認

Mattermostにメッセージを投稿する際、優先度を設定できるようになりました。  
**システムコンソール > サイト設定 > 投稿 > メッセージの優先度**から有効にすることができます。

![channels-priority-system-console](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-priority-system-console.png)

メッセージの優先度機能を有効にすると、メッセージ入力欄の左部に優先度を設定するためのアイコンが表示されます。

![channels-priority-set](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-priority-set.png)

設定できる優先度は、`標準(Standard)`、`重要(Important)`、`緊急(Urgent)`の3種類です。(通常のメッセージが`標準(Standard)`です)

![channels-priority-type](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-priority-type.png)

`標準`以外の優先度を設定すると、投稿の上部に選択した優先度が色付きで表示されます。

![channels-priority-post](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-priority-post.png)

また、Professionalプランでは、メッセージに対する確認を要求することもできます。

![channels-acknowledge](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-acknowledge.png)

`確認を要求する`をONにしてメッセージを投稿すると、他のユーザーには　**確認** ボタンが表示されるようになります。

![channels-acknowledge-ok](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-acknowledge-ok.png)

**確認** ボタンを押すと、確認済みのユーザー数がカウントアップされます。マウスを当てると、誰が確認済みかを確認することができます。

![channels-acknowledge-show](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-acknowledge-show.png)

## Playbooks: 既存チャンネルでのPlaybookの実行

今までのバージョンでは、Playbookを実行すると新しいチャンネルが自動で作成されていましたが、今回のバージョンから既存のチャンネル内でPlaybookを実行できるようになりました。

まず、チャンネル右部に表示されているPlaybooksのアイコンをクリックすると、現在、このチャンネルで実行中のPlaybookを確認することができます。(**システムコンソール > 実験的機能 > 機能 > App Barを有効にする**が無効になっている場合、右サイドバーのメニューは表示されず、チャンネルヘッダ部分にPlaybookアイコンが表示されます)  
ここで、右サイドバー上部に表示されている `実行開始` ボタンをクリックすると、Playbookを開始することができます。

![playbooks-rhs-run](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-rhs-run.png)

開始したいPlaybookを選択すると、

![playbooks-select-playbook](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-select-playbook.png)

実行の詳細を入力する画面が表示されます。  
この時、`既存のチャンネルとリンクする`を選択し、Playbookを実行したいチャンネルを選択してから `実行開始` を押すことで、新規にチャンネルを作成することなく既存のチャンネル内でPlaybookを開始することができます。

![playbooks-select-channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-select-channel.png)

既存のチャンネル内にPlaybookの実行が追加され、このチャンネルの中で実行の状況を追跡することができます。

![playbooks-running](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-running.png)

## Playbooks: タスクリスト

Playbooks画面の右上にあるタスクアイコンをクリックすることで、自分が関係しているタスクの一覧を確認することができるようになりました。

![playbooks-task-inbox](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-task-inbox.png)

`フィルター`メニューから、以下のタスクの表示/非表示を切り替えることができます。

* **所有する実行からすべてのタスクを表示する**: 自分がオーナーとして参加しているPlaybook実行のタスクをすべて表示する
* **チェック済みのタスクを表示する**: 既に完了としてチェックしたタスクを表示する

![playbooks-task-filter](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/playbooks-task-filter.png)

## Boards: ファイル添付

Cardにファイルを添付することができるようになりました。  
Card画面の上部に表示されている `Attach` メニューからファイルを選択することで添付を行うことができます。

![boards-file-attachment](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/boards-file-attachment.png)

アプロードされたファイルは、Card内の**Attachment**セクションに表示され、ここからダウンロードや削除を行うことができます。

![boards-attached-files](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/boards-attached-files.png)

## Integrations: ServiceNow Channel integrations, GitLab Playbooks integration

今回のリリースブログでは、ServiceNow連携機能とGitLab連携機能の紹介が行われています。

[ServiceNowr連携機能](https://mattermost.com/marketplace/servicenow/)では、ServiceNowのITSM機能と連携し、特定のイベントが発生した際にMattermostに通知したり、MattermostからITSMレコードへのコメント追加等を行うことができるようです。また、MattermostのPlaybooks機能と連携することで、インシデント対応時のコミュニケーションを円滑に行うことができます。  
また、[ServiceNow Virtual Agent](https://mattermost.com/marketplace/servicenow-virtual-agent/)を使うことで、ServiceNowのAIチャットボット機能であるVirtual AgentをMattermostに組み込むことができます。  
ServiceNow連携についてはYouTubeで紹介動画が公開されています。

https://www.youtube.com/watch?v=OVHng0IjLT0

GitLab Pluginは、最新バージョン(v1.6)でGitLab Pipelineとの連携が強化され、MattermostのPlaybooksk機能と連携することで、プロダクトのリリース管理をより効率的に行うことができるようになったようです。

[GitLab Plugin \- Mattermost](https://mattermost.com/marketplace/gitlab-plugin/)

## アップグレード時の注意事項

アップグレード時の注意事項について、詳しくは公式ドキュメントを確認ください。　 
[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

### Webappコンポーネントを含むPlugin

今回のリリースから依存するReactのバージョンがReact 17になった影響で、もしPluginがReact 16を使用してビルドされている場合、バージョン不整合によりMattermostがクラッシュする恐れがあります。Mattermost公式のPluginテンプレートである[mattermost/mattermost-plugin-starter-template](https://github.com/mattermost/mattermost-plugin-starter-template)を利用している場合、この問題の影響を受ける可能性があります。

回避策としては、[このPR](https://github.com/mattermost/mattermost-plugin-playbooks/pull/1489)を参考に修正を行なって再度Pluginをビルドし直すか、一時的な回避策として**システムコンソール > 実験的機能 > 機能 > プラグインが使用するReact DOMにパッチを当てる**を有効にすることで回避することができます。

`Posts`テーブル内の`ParentID`カラムを削除するDBマイグレーションが実行されますが、MySQLを利用している場合、`Posts`テーブルのサイズによってはこの処理によってCPU使用率が上昇する場合があります。
また、このマイグレーション実行中は書き込みが制限されます。

![warning-reactdom](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/warning-reactdom.png)

### Calls機能の最低動作バージョン

Calls機能専用サーバーである[rtcd](https://github.com/mattermost/rtcd)を利用してCalls機能を動作させている場合、今回のリリースからrtcdの最低動作バージョンが設定されるようになりました。Mattermost v7.7に組み込まれているCalls機能ではrctd v0.8.0を要求し、これ以下のバージョンで動作している場合はCalls機能が動作しません。

### MySQLでのシステムタイムゾーン設定時の問題

Mattermostが利用しているMySQLがシステムのタイムゾーンを参照しており、かつ、システム側で設定されているタイムゾーンがMySQLでサポートされていなかったばい、エラーが発生する可能性があります。この問題はタイムゾーンテーブルを入力することで回避することができます。

### Threadsテーブルの非正規化

`Threads`テーブルに`TeamId`が追加される非正規化が行われました。これによるデメリットは無いとされています。

### Webappプラグイン

[`PluginRegistry.registerCustomRoute`](https://developers.mattermost.com/integrate/plugins/components/webapp/reference/#registerCustomRoute)を利用したプラグイン開発を行なっている場合、CSS`grid-area: center`が設定されていないとコンポーネントが適正な位置に表示されなくなります。

## その他の変更

### 下書きの保存

編集途中のメッセージを管理する `下書き` 機能が追加されました。

![channels-draft](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-draft.png)

下書きは、チームごとに管理されます。

また、**システムコンソール > サイト設定 > 投稿 > 下書きのサーバーへの同期を有効にする**を有効にすることで、下書きの情報をサーバーに保管できるようになり、編集を開始した端末とは別の端末からでも下書きにアクセスできるようになります。(ただし、モバイルアプリでは下書き機能自体がないようです)  
システムコンソールでこの設定を有効にした場合でも、**設定 > 詳細 > メッセージの下書きをサーバーと同期する**からユーザーごとにOn/Offを切り替えることができます。

![channels-draft-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-draft-setting.png)


### フリープランでもMy Insightが利用可能に

今までは有償版でのみ利用可能であったインサイト機能のうち、自分のアクティビティに関係する `私のインサイト` 機能がフリープラン(無償版)でも利用可能になりました。  
自分がよく活動するチャンネルやPlaybook、よく使うリアクション、また、活動の少ないチャンネル等を確認することができます。

![channels-insight](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/channels-insight.png)

## その他のトピック

### Roadmap

Mattermostの各機能のRoadmapは以下の公式チャットで公開されています。

https://community.mattermost.com/core/channels/roadmap

Mattermostのチャット機能では、再来月(2023/03)リリース予定の v7.9 で、メッセージに対する確認を定期的に要求する機能が追加されるようです。  
この機能は今回リリースされたメッセージへの確認要求に付随する機能になるため、Professionalプラン以上で利用可能になるようです。

![roadmap-persistent-notification](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.7/roadmap-persistent-notification.png)

[Channels: January 2023 Roadmap Update](https://docs.google.com/presentation/d/1jExa80sdoylFbUnD8SyO7irGveTUCfZDXSI4zRCnm0A/edit)

### Mattermost関連の記事紹介
* [Zapierを使ってQiitaに投稿したらMattermostに通知を飛ばす \- Qiita](https://qiita.com/honma-h/items/86fc2ebeeaf928b836da)
  * MattermostのOAuth 2.0アプリケーション機能を使った連携方法の紹介です
* [ansibleでディスク使用率監視とMattermostへの通知 \- Qiita](https://qiita.com/horonium/items/cb3c0258606f2db2df32)
  * MattermostのWebHookを使った連携方法の紹介です
* [Growiの更新をIFTTTを使ってMattermostに通知する \- Qiita](https://qiita.com/honma-h/items/fb68204d3a3a6511f252)
  * MattermostのWebHookを使った連携方法の紹介です

## おわりに
次の`v7.8`のリリースは 2023/02/16(Thu)を予定しています。

元々、今回の`v7.7`が[ESR(Extended Suport Release)](https://docs.mattermost.com/upgrade/extended-support-release.html)となる予定でしたが、`v7.6`のリリースがスキップされた影響で多くの機能が`v7.7`でリリースされることになったため、次のESRは`v7.8`とするよう決定されました。
