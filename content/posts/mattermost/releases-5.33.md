---
title: "Mattermost 5.33の新機能"
date: 2021-03-20T10:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/03/17 に Mattermost v5.33.0 がリリースされました。  
また、近いうちに TLS 接続をしている際に WebSocket 通信が失敗する問題の修正を加えた v5.33.1 がリリースされるようです。  
[\[MM\-34000\] Websockets fail with TLS connections \- Mattermost](https://mattermost.atlassian.net/browse/MM-34000)

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise 版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.33 is now available \- Upgrade to the new release today](https://mattermost.com/blog/mattermost-release-v5-33/)
- [Mattermost Changelog — Mattermost 5\.33 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-33-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## カスタムステータス

Mattermost 上に表示されるステータスを自由に設定できるようになりました。  
今までのバージョンで設定できるステータスは、システム固有の`オンライン(Online)`、`離席中(Away)`、`取り込み中(Do Not Disturb)`、`オフライン(Offline)`の 4 種別でしたが、 これらのステータスとは別に、自由にメッセージや絵文字を指定できるカスタムステータスを設定できるようになりました。

Mattermost 画面左上のアイコン部分をクリックするとステータス設定メニューが表示され、その中にカスタムステータス設定メニューがあります。  
![status menu](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.33/status-menu.png)

カスタムステータス設定メニューでは、絵文字とメッセージを指定することができます。絵文字にはユーザーが登録したカスタム絵文字も指定することができます。  
![set custom status](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.33/set-status.png)

設定されたステータスは、投稿に表示される名前の横などに現れます。  
![show custom status](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.33/show-status.png)

また、スラッシュコマンドを使ってカスタムステータスを設定することもできます。  
![slash command](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.33/status-command.png)

カスタムステータスを設定しても、システム固有のステータス(`オンライン`等)は別に設定することができるため、システム固有ステータスによる通知のコントロール機能は引き続き使用することができます。

カスタムステータスについて、詳しくは下記のエントリでも紹介されています。
[Introducing custom statuses for Mattermost users \| Mattermost](https://mattermost.com/blog/custom-statuses/)

## (Cloud/E20) OpenID Connect

OpenID Connect 仕様に基づいた全ての OAuth 2.0 プロバイダを Mattermost の認証に使用できるようになりました。

今までのバージョンでは、GitLab、Google Apps、Office 365 による OAuth 2.0 認証が使用できましたが、OpenID Connect のサポートにより、Keycloak、Atlassian Crows、Apple、Microsoft、Salesforce、Auth0、Ory.sh、Facebook、Okta、OneLogin、Azure AD などが認証に使用できるようになります。

## (Cloud/E20) インシデントタイムライン

ここ数ヶ月開発が活発な[Incident Management Plugin](https://github.com/mattermost/mattermost-plugin-incident-collaboration)にタイムライン機能が追加されました(呼び方が`Incident Collaboration Plugin`に変わった?)。

タイムライン機能では、インシデント対応におけるステータスを更新や重要なイベントの記録を時系列で確認できる機能です。タイムライン機能により、インシデント対応後のレポート作成や、対応中インシデントの状況確認などがやりやすくなります。さらに、タイムラインに表示されたイベントをクリックすることで、そのイベントが行われた際の投稿にジャンプできるため、イベントに関連する会話を見返すのも容易になります。

![incident timeline](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.33/incident-timeline.png)

Incident Collaboration 機能のアップデートの詳細については、以下のエントリで詳しく紹介されています。

[Mattermost Incident Collaboration for faster incident response](https://mattermost.com/blog/mattermost-incident-collaboration-for-incident-response-teams/)

## (Cloud/E10/E20) サポートパケット生成機能

Mattermost インスタンスの設定の詳細や、ログ、その他のデプロイ情報などをダウンロードできるようになりました。これは、Mattermost の商用サポートチームが、トラブルシューティングを円滑に行うための機能です。

## その他のトピック

### Focalboard

Mattermost が開発している `a self-hosted alternative to Trello, Notion, and Asana` な[Focalboard](https://focalboard.com)が Reddit や HackerNews に取り上げられ、話題になっています。

- https://www.reddit.com/r/selfhosted/comments/m72mes/focalboard_open_source_selfhosted_project/
- https://news.ycombinator.com/item?id=26499062

注目されたことで開発も活発になり、[Docker サポート](https://github.com/mattermost/focalboard/pull/105)や[翻訳サポート](https://github.com/mattermost/focalboard/commit/8f31d14a304a5a46d8f7f197a85361d04a196fff)などが進んでいます。

ちょうど私も先週 Focalboard を動かしていて、その時の感想を Zenn の方に書いていました。

- [Notion のような UI の Trello っぽい Kanban ツールの OSS の Focalboard を触ってみた](https://zenn.dev/kaakaa/articles/mattermost-focalboard-first)

まだバージョンが v0.6.1 ということもあり開発段階ではありますが、面白そうなプロジェクトのため今後もチェックしていきたいと思います。

### Mattermost リポジトリのスター数が 20,000 に

Mattermost Server の GitHub リポジトリのスター数が 20,000 を超えたようです。

[The Mattermost server repo surpasses 20,000 stars on GitHub](https://mattermost.com/blog/mattermost-server-surpasses-20000-stars-on-github/)

### International Woman's Day

3/8 の International Woman's Day に合わせ、Mattermost に携わる女性にフィーチャーしたエントリが投稿されています。

[Happy International Women’s Day from Mattermost\!](https://mattermost.com/blog/international-womens-day-2021/)

## おわりに

次の`v5.34`のリリースは 2021/04/16(Fri)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
