---
title: "Mattermost 5.34の新機能"
date: 2021-04-17T10:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/04/15 に Mattermost v5.34.0 がリリースされました。  
その後、MySQL を使用している場合に発生する可能性のある問題に対応したパッチバージョンをリリースしており、MySQL を利用している場合は注意が必要そうです。以下の Changelog を確認し、最新のリリースを利用するようにしてください。

https://docs.mattermost.com/administration/changelog.html#release-v5-34-feature-release

v5.34.0 へのアップグレード時に実行される DB マイグレーションがタイムアウトすることがあり、この問題については v5.34.1 で解消されています。v5.34.1 にアップグレードする場合、大規模なインスタンスだとマイグレーションに時間がかかる可能性があります。

> v5.34.1, released 2021-04-15
> Fixed an issue where upgrading to v5.34.0 runs a migration that could cause timeouts on MySQL installations. Upgrading to v5.34.1 may also execute missing migrations that were scheduled for v5.32.0. These additions can be lengthy on very big MySQL (version 5.x) installations.

MySQL のパスワードに特殊文字を使用している場合、Mattermost の開始に失敗する問題への対処が v5.34.2 にて加えられています。

> v5.34.2, released 2021-04-17
> Fixing an issue where installs with some special characters in the MySQL password will break and fail to start.

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise 版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.34 is now available \- Upgrade to the new edition today](https://mattermost.com/blog/mattermost-release-v5-34/)
- [Mattermost Changelog — Mattermost 5\.34 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-34-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## (Cloud/E20) Incident Collaboration におけるインシデント開始時操作の自動化

ここ数ヶ月開発が活発な[Incident Collaboration Plugin](https://github.com/mattermost/mattermost-plugin-incident-collaboration)に、インシデント開始時の操作を自動化する機能が追加されました。

Incident Collaboratoin Plugin では、インシデント対応手順を予め Playbook として登録しておき、インシデント発生時に Playbook を元にしたインシデント専用チャンネルを作成することで、インシデント対応の様子を 1 つのチャンネルにまとめることができます。

今回追加されたのは、インシデント専用チャンネル作成時に自動でそのチャンネルに追加されるメンバを指定しておける機能です。

![incident automation](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.34/incident-automation.png)

## その他の機能

### 新たなプラグイン (CircleCI / Dice Roller)

新たなプラグインが紹介されています。

#### CircleCI Plugin

[Mattermost/CircleCI Plugin \- Circle CI Plugin](https://mattermost.gitbook.io/circle-ci-plugin/)

Mattermost と CircleCI を連携するプラグインです。こちらは Mattermost 公式チームによって開発されているようです。

- パイプライン/ワークフローの実行
- Mattermost への通知
- 環境変数の管理
- ワークフローメトリクスの取得

などが Mattermost 上から行えるようです。

#### Dice Roller

[Dice Roller Plugin \- Mattermost \- Open\-source collaboration, self\-managed or SaaS](https://mattermost.com/marketplace/dice-roller-plugin/)

Dice Roller プラグインは、Mattermost 上でダイスを振ることのできるプラグインです。
![dice roller](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.34/dice-roller.png)

### システム管理者向け非公開チャンネル参加確認ダイアログ

システム管理者が非公開チャンネルにそのリンクから参加しようとした際、確認のためのダイアログが表示されるようになりました。

![confirm-to-join-private](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.34/confirm-to-join-private.png)

### チームサイドバーのテーマカラー

Mattermost UI のテーマカラー指定において、チームサイドバーの部分にもカラー設定が可能になりました。`sidebarTeamBarBg`という値でサイドバーの色を指定できるようです。

![theme teambar](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.34/theme-teambar.png)

以下で Mattermost のテーマ集が公開されています。ここで公開されているテーマは公式テーマではないため、今回追加された`sidebarTeamBarBg`が設定されていない可能性があります。

[Mattermost Themes](https://avasconcelos114.github.io/mattermost-themes/)

### pprof によるパフォーマンスモニタリング

今まで有償版(E20)限定の機能だった[パフォーマンスモニタリング](https://docs.mattermost.com/deployment/metrics.html)の機能のうち、[pprof によるモニタリング機能](https://docs.mattermost.com/deployment/metrics.html#standard-go-metrics)を Team Edition(OSS 版)でも利用できるようになりました。

pprof によるモニタリングを行うには、**システムコンソール > パフォーマンスモニタリング > パフォーマンスモニタリングを有効にする** を **有効** に設定する必要があります。

![performance monitoring](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.34/performance-monitoring.png)

設定を変更した後、同じ設定画面で指定してあるポート番号に`pprof`を使ってアクセスすることで、Mattermost インスタンスのメトリクスを確認できるようになります。

```
$ go tool pprof http://localhost:8067/debug/pprof/profile
Fetching profile over HTTP from http://localhost:8067/debug/pprof/profile

channel: open channel: no such file or directory
Fetched 1 source profiles out of 2
Saved profile in /Users/test/pprof/pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Apr 16, 2021 at 5:30pm (JST)
Duration: 30s, Total samples = 50ms ( 0.17%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 50ms, 100% of 50ms total
Showing top 10 nodes out of 13
      flat  flat%   sum%        cum   cum%
      20ms 40.00% 40.00%       30ms 60.00%  runtime.scanobject
      10ms 20.00% 60.00%       10ms 20.00%  github.com/gorilla/websocket.(*Conn).write
      10ms 20.00% 80.00%       10ms 20.00%  runtime.greyobject
      10ms 20.00%   100%       10ms 20.00%  runtime.netpoll
         0     0%   100%       10ms 20.00%  github.com/gorilla/websocket.(*Conn).WriteMessage
         0     0%   100%       10ms 20.00%  github.com/gorilla/websocket.(*messageWriter).flushFrame
         0     0%   100%       10ms 20.00%  github.com/mattermost/mattermost-server/v5/app.(*WebConn).Pump.func1
         0     0%   100%       10ms 20.00%  github.com/mattermost/mattermost-server/v5/app.(*WebConn).writePump
         0     0%   100%       40ms 80.00%  runtime.gcBgMarkWorker
         0     0%   100%       40ms 80.00%  runtime.gcBgMarkWorker.func2
```

## その他のトピック

### Mattermost CEO による振り返り

Mattermost の公式ブログで、今までの Mattermost の活動をどのような思想で進めてきたかについて、Mattermost CEO によって綴られています。

[How We’ve Built an Open Source Community at Mattermost](https://mattermost.com/blog/building-open-source-community/)

こちらの記事では、Mattermost がどのようにオープンソースコミュニティを築くために、どのような思想を持って活動してきたのかについて語られています。

[Building a Go\-to\-Market Strategy for Developer Tools](https://mattermost.com/blog/go-to-market-strategy-for-developer-tools/)

こちらの記事では、Mattermost を成長させるための戦略について語られています。

## おわりに

次の`v5.35`のリリースは 2021/05/14(Fri)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
