---
title: "Mattermost5.30の新機能"
date: 2020-12-20T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/12/18にMattermost v5.30.0がリリースされました。また、現在v5.30.1もリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost release v5\.30 is now available](https://mattermost.com/blog/mattermost-release-v5-30/)
* [Mattermost Changelog — Mattermost 5\.30 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-30)

---

## Mattermostで投票機能
Mattermost上で投票機能を実現する[`Matterpoll Plugin`](https://github.com/matterpoll/matterpoll)が、公式のプラグインマーケットプレースからインストールできるようになりました。

![movie](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/matterpoll.gif)

Matterpollプラグインについては、過去に自分が作り始めたものがプラグイン形式に変わって開発が進められているものです。
[Mattermostで投票機能もどき \- Qiita](https://qiita.com/kaakaa_hoe/items/b2605ce3816cfc517ecd)

Matterpollプラグインについては、[Mattermost Integrations Advent Calendar 2020 \- Qiita](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第22日目以降の記事で内部構造等を紹介します。

## (E20) Mattermostでインシデント管理
Mattermost上でインシデントの管理を行うことができる機能が追加されたようです。
元々、[Incident Management Plugin](https://github.com/mattermost/mattermost-plugin-incident-management)と[Mattermost Channel Export Plugin](https://mattermost.gitbook.io/channel-export-plugin/)として公開されていたプラグインをMattermost本体に内蔵したということのようです。（E20版限定の機能として紹介されていますが、Team Editionでも使用できている気がします）

### インシデント管理方法
[Incident Management Plugin](https://github.com/mattermost/mattermost-plugin-incident-management)が有効になっている場合、**メインメニュー > Playbook & Incidents**からインシデント管理画面を表示できます。
![](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/incident-mainmenu.png)

インシデント管理画面では、インシデント対応の手順をまとめた`Playbook`と、インシデント一覧を管理することができます。
![](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/incident-management.png)

`Playbook`タブで`+ New Playbook`をクリックするとPlaybookの作成画面が表示され、インシデント対応プロセスとしてステージと各ステージにおけるタスクを入力することができます。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/incident-playbook.png)

また、各タスクにはSlash Commandを指定することができるようですが、これは、例えば[Mattermost Jira Plugin](https://mattermost.gitbook.io/plugin-jira/)のコマンドを指定してタスク完了時にJiraなどの外部サービスへの通知が必要なことを表示するための機能ではないかと思われます。Playbook内のタスクに存在しないSlash Commandを指定しても、新たなSlash Commandが作成されるわけではなさそうです。

ここで作成したPlaybookは、チャンネルヘッダーボタンの`Incidents`メニューなどから`+ Start Incident`を実行した際に、インシデント対応プロセスとして指定することができます。

![movie](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/incident.gif)

インシデントが作成されると、インシデント対応専用のチャンネルが作られ、右サイドバーではインシデントの進行状況を確認することができます。対応担当を指定したり、完了したタスクをクローズするなどしてインシデントを進めていくことになります。

タスク完了後は[Mattermost Channel Export Plugin](https://mattermost.gitbook.io/channel-export-plugin/)で、チャンネル内でのやりとりをCSV形式でエクスポートなどもできるようです。

Incident Managementプラグインには、インシデント管理用のSlash Commandなどもあるようで、詳しくは下記の公式ドキュメントを参照ください。

https://docs.mattermost.com/administration/devops-command-center.html

## (E20) Microsoft Teams会議プラグイン (Beta版)

Mattermostのチャンネルヘッダーボタンから、Microsoft Teamsの会議を作成し、会議へ参加するためのボタンを含む投稿をMattermostチャンネルに作成するプラグインのようです。こちらは有償版Mattermost限定の昨日のようです。

![公式ブログから](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/ms-teams-plugin.webp)
（※画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-30/)から）

## (E20) 新たに追加された管理者ロールの権限が設定可能に
Mattermos v5.28で、システム管理者の一部の権限のみを持つ`System Manager`、`User Manager`、`Read-only Admin`という役割が新たに追加されましたが、今回のv5.30から、これらの役割の権限を任意に設定できるようになりました。

![公式ブログから](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.30/admin-roles.webp)
（※画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-30/)から）

## その他のトピック

### Mattermost Cloud
今までオンプレミス専門でやってきたMattermostのクラウドサービスが始まりました。

[Introducing Mattermost Cloud \- Open Source, Private Cloud Messaging](https://mattermost.com/blog/introducing-mattermost-cloud/)

ユーザーとワークスペースを作成すると専用のMattermostのインスタンスが作成され、自分が作成したワークスペース内ならシステムコンソールにも普通にアクセスできるようです。  
10ユーザーまでは無償で利用でき、無償版でもセルフホスティング版の有償版相当の機能が使え流ように見えます。（システムコンソールで有償版の設定にもアクセスできる）

https://mattermost.com/pricing-cloud/

Slackのようにメッセージ履歴や統合機能数などでの制限ではなく、ユーザー数での制限のみのため、その辺りが一つの差異化ポイントのようです。

### リリースサイクルの見直し

これまで、Mattermostは偶数バージョンは新機能を追加、奇数バージョンはバグ修正などの品質改善のみを追加、という"tick-tock"なリリースサイクルを採用していました。

しかし、前記のMattermost Cloudの開始などを受け、リリースサイクルを早めるべきとの結論に至ったようで、2021/01/16リリースのv5.31からは各バージョンでの新機能の追加が行われるようになります。

[Mattermost moves away from alternating release cycle](https://mattermost.com/blog/discontinuing-alternating-release-cycle/)

### What Matters - A Podcast from Mattermost
少し前になりますが、Mattermost主催のPodcastの配信が始まっています。
Mattermost社のメンバーや、外部からエンジニアを迎えて技術的な話や開発に関するトークなどを不定期で配信しているようです。

* https://mattermost.libsyn.com/
* [Announcing What Matters: A podcast from the folks at Mattermost](https://mattermost.com/blog/what-matters-podcast/)

### Mattermost Dev Sneak Peek (YouTube)
最近、MattermostのYouTubeチャンネルで**Dev Sneak Peek**として、開発中の新機能の紹介動画をよくアップロードしているようです。

* https://www.youtube.com/channel/UCNR05H72hi692y01bWaFXNA

### Mattermost Integrations Advent Calendar
Mattermost Integrationsに関する記事をQiita Advent Calendarに書いています。

[Mattermost Integrations Advent Calendar 2020 \- Qiita](https://qiita.com/advent-calendar/2020/mattermost-integrations)

## おわりに
次の`v5.31`のリリースは2021/01/15(Fri)を予定しています。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
