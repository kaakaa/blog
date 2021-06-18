---
title: "Mattermostへの貢献"
date: 2020-12-25T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "oss", "contribute"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第25日目の記事です。

最終日の記事は、これまでの記事執筆にあたって見つけた問題を修正した記録について紹介します。と言っても、簡単なドキュメントの修正が主です。（昨日までのMattermost Pluginの記事が重かったため力尽きました...。）

## Mattermostへのコントリビュート

MattermostはOSSとして開発が進められており、そのコードやドキュメントなどはGitHubで公開されているため、誰でもPullRequestを送ることができます。

https://github.com/mattermost

コントリビューションの仕方などは本日の記事では紹介しませんが、以下のページにMattermostとしてのコントリビューションのやり方などについてまとめてあります。コントリビューションを続けることで[様々なギフト](https://blog.kaakaa.dev/post/mattermost/advent-calendar-2020/day13-swags/)を貰うことができるので、興味のある方は是非コントリビューションしてみてください。

https://mattermost.com/contribute/

## コントリビューション記録

### ドキュメントのフォーマット修正
[Specify required language in code blocks by kaakaa · Pull Request \#726 · mattermost/mattermost\-developer\-documentation](https://github.com/mattermost/mattermost-developer-documentation/pull/726)

開発者向けのドキュメントサイトで、フォーマットが正しくされていない点がいくつかあったため修正しました。

### Tooltipテキストの修正
[Fix tooltip text for copying incoming webhook url by kaakaa · Pull Request \#6996 · mattermost/mattermost\-webapp](https://github.com/mattermost/mattermost-webapp/pull/6996)

Incoming WebHook作成時のコピーアイコンに設定されているTooltipテキストに誤った値が設定されていたのを修正しました。

### ドキュメントのパラメータ型誤りの修正
[Fix incorrect type of \`notify\_on\_cancel\` by kaakaa · Pull Request \#4193 · mattermost/docs](https://github.com/mattermost/docs/pull/4193)

これもドキュメントの修正で、パラメータの型が誤っていたのを修正（併せて簡単なインデントの修正も）。

### デッドリンクの修正
[fix invalid link for js websocket client by kaakaa · Pull Request \#596 · mattermost/mattermost\-api\-reference](https://github.com/mattermost/mattermost-api-reference/pull/596)

APIドキュメントのJavaScript WebSocket Clientへのリンクが404だったので修正。

## さいごに

これまで25日に渡ってOSSのSlack Alternativeなチャットツール Mattermostの外部連携を行うための統合機能について紹介してきました。
チャットと言えばユーザー数の面からSlackやMicrosoft Teamsが有名ですが、オンプレミスなどのプライベートな環境にも構築可能なチャットツールとして、Mattermostも着実に進化を続けています。Mattermostは、最近では[DevOps Command Center](https://mattermost.com/devops/)という方針を打ち出しており、ただのチャットだけではなくDevOpsを進める上でのコミュニケーションの中心となるべく機能が追加されています。2020/12/16にリリースされたMattermost v5.30に同梱されている、Mattermost上でインシデント管理を行える [Mattermost Incident Management Plugin](https://mattermost.gitbook.io/mattermost-incident-management/) などはその最たるものだと思います。
これらの外部サービスと連携するコミュニケーション基盤は、今回のアドベントカレンダーで紹介してきたMattermost Integrationsの機能をベースに構築されるため、DevOpsに限らずMattermostのユースケースを広げようとする際に参考になる情報になるのではないかと思います。
