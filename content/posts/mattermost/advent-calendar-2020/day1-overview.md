---
title: "Mattermost Integrationsアドベントカレンダーについて"
date: 2020-12-01T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration"]
---


Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第1日目の記事です。

今日からMattermostの統合機能の開発方法について書いていきます。

## Mattermost について
Mattermostは、米国のMattermost Inc.が開発しているSlack Alternativeなチャットツール[Mattermost](https://mattermost.com)です。OSSとして開発が進められており、誰でも開発に参加することができます。

https://github.com/mattermost

Atlassian HipChat亡き後、オンプレミスで運用できるビジネスチャット基盤として勢力を伸ばしており、昨年には投資額が $50Mに達し、Y CombinatorのSeries B史上最大の投資額を得るなど、その規模を拡大しています。
[YC leads $50M Series B in Mattermost as open source Slack alternative \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/yc-leads-50m-series-b-in-mattermost-as-open-source-slack-alternative/)


国内でも(自分の観測範囲だけでも)様々な企業・組織が採用しており、また、Mattermost Inc.と提携し、サービスを提供している企業もいくつもあります。
* [Mattermost \| aslead/アスリード \| 野村総合研究所\(NRI\)](https://aslead.nri.co.jp/products/mattermost.html)
* [Educhat](https://www.castalia.co.jp/educhat)
* [Mattermost\(企業向けコラボレーション チャットツール\)｜ 製品 ｜ リックソフト](https://www.ricksoft.jp/mattermost/)

---

Mattermostはビジネスチャットの機能だけでも便利なのですが、さらに組織の要望に合わせたカスタマイズを行うことができます。このカスタマイズの方法は多岐に渡るのですが、まとまった情報が公式の英語ドキュメントしかなさそうなので情報源を増やすことを目的に記事を書いていこうと思います。各機能の解説というよりは、その機能を使ってどのようなことが出来るのかを実例と合わせて書いていきたいと思っています。

また、今回書いていく内容は、現在リリースされているMattermostの最新版 v5.28 で動作確認を行っているものであり、異なるバージョンを使っている場合は公式ドキュメントを参照することをオススメします。

* https://docs.mattermost.com/guides/integration.html
* https://developers.mattermost.com/integrate/getting-started/

## 概要

１日目の記事は、Mattermostの統合機能の全体像を紹介します。

![overview](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day1/overview.drawio.png)

まず、Mattermost本体の機能は上図青い四角のMattermost ServerとMattermost Webappから成り、Mattermost Serverはデータの管理などのコアの機能を果たしており、Go言語で書かれています。Mattermost WebappはブラウザやデスクトップアプリのUIの部分であり、React/Reduxで書かれています。

Mattermost ServerとMattermost Webappの間は、初回こそHTTPでやりとりされますが、その後の通信は基本的にWebSocketで行われています。

このMattermostのコア機能に対して独自の処理を追加するための拡張ポイントとして、赤字で書かれたMattermostの統合機能があります。

#### (1) Incoming WebHook (HTTP)
Incoming WebHook（内向きのWebHook）は、外部アプリケーションからMattermostへ投稿をするためのエンドポイントを新たに作成する機能です。この新たに作成したエンドポイントに対して既定のJSONをHTTP POSTすることで、Mattermost外のアプリケーションから投稿が作成できます。

Incoming Webhookについては第2,3日目の記事で紹介します。

#### (2) Outgoing WebHook (HTTP)
Outgoing WebHook（外向きのWebHook）は、Mattermost上でユーザーによって作成された投稿の情報を外部アプリケーションへ送信する機能です。Mattermost上での会話に応じて処理を行いたい場合などに使用します。

Outgoing Webhookについては第4,5日目の記事で紹介します。

#### (3) Slash Command (HTTP)
Slash Commandは、(2) Outgoing WebHookと動作は似ていますが、ユーザーが`/`で始まる特殊なコマンドを投稿することで動作させることができる機能です。ユーザーの任意のタイミングで処理を実行させることができるため、Outgoing WebHookよりも使い勝手が良いです。

Slash Commandについては、第6, 7, 8日目の記事で紹介します。

#### (4) REST API (HTTP)
REST API は、Mattermost上で管理されているリソースを取得、作成、更新、削除する機能です。一般的なREST APIです。HTTPによるリクエストだけでなく、Mattermost ServerとMattermost Webapp間でやりとりされるWebSocketを購読することもできます。

REST APIについては、第10,11日目の記事で紹介します。

#### (5) Plugin (Server) (RPC)
Plugin(Server) はMattermost Serverの機能を拡張するための機能です。今まで紹介してきた機能は、基本的にHTTPを通じてMattermostと外部アプリケーションのやりとりを行う機能でしたが、Plugin(Server)は、Mattermostと同一サーバー上でRPCを通じてやりとりを行う機能です。今まで紹介してきた機能より、より多様なMattermost上のイベントとやりとりを行うことができます。

Plugin (Server)については、第18, 19日目の記事で紹介します。

#### (6) Plugin (Webapp)
Plugin(Webapp)は、Mattermost Webappの機能を拡張するための機能です。ブラウザなどのクライアントアプリ上のMattermost画面の見え方を変えることができます。また、Plugin(Server)はMattermost Server上の全ユーザー共通の情報を扱う機能でしたが、Plugin(Webapp)により、ユーザーごとに異なる見え方を提供したりできます。

Plugin (Webapp)については、第20,21日目の記事で紹介します。

---

Mattermostの統合機能は大きくわけで以上の５つになるかと思います。また、統合機能とは別に統合機能から作成する投稿の見た目を変える`Message Attachments`や`InteractiveMessage`などの機能もあるため、それらの機能についても触れていきます。
