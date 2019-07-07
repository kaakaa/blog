
---
title: "Mattermost4.2の新機能"
date: 2017-09-23T15:04:04+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

前回の更新([v3.8](https://qiita.com/kaakaa_hoe/items/1ae33a780a4a17f4374b))からだいぶバージョンが飛んでしまいましたが、2017/9/15にMattermost v4.2.0がリリースされたので、追加された機能などについて簡単に紹介します。

詳細については公式のリリースを確認ください。  
[Mattermost 4.2: Interactive message buttons, AppConfig for Mobile App - Mattermost](https://about.mattermost.com/blog/mattermost-4-2/)
[Mattermost Changelog — Mattermost 4.3 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-2-0)


## インタラクティブメッセージボタン

Mattermostの統合機能を通じて、ユーザーが選択可能なボタンを付けた投稿を行えるようになりました。
この機能に関する[ドキュメント](https://docs.mattermost.com/developer/interactive-message-buttons.html)を見るに、ボタン毎に異なるHTTPリクエストを送信するもののようです。

この機能を利用した例として、投票機能に関するアプリが紹介されています。
[mattermost/mattermost-interactive-post-demo: Interactive post demo for Mattermost](https://github.com/mattermost/mattermost-interactive-post-demo)

![poll.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/5c7cb180-2aa0-d9c3-9f23-d548c9dc39cb.gif)

(Mattemost上での投票機能については、拙作となりますが絵文字リアクションを利用したGoアプリを以前から公開しています。ご参考まで。 [kaakaa/matterpoll-emoji: Poll server for Mattermost](https://github.com/kaakaa/matterpoll-emoji) )

## カスタムスラッシュコマンドによるワークフローの改善

こちらは新機能ではないですが、スラッシュコマンドの使いやすさの紹介です。

[Slash Commands — Mattermost 4.3 documentation](https://docs.mattermost.com/developer/slash-commands.html#slash-commands)

カスタムスラッシュコマンドを作成・実行すると、任意のエンドポイントにリクエストを送信できるため、AWS API Gateway / AWS Lambdaと組み合わせることで効果を発揮します。
ここでは、Standupミーティングの状態をハッシュタグ付きで書き込んでくれるスラッシュコマンドの例が紹介されています。
[Custom Slash Commands in Mattermost with AWS Lambda and API Gateway · Grundleborg's Cave](https://grundleborg.github.io/posts/mattermost-custom-slash-command-aws-lambda/)

## モバイルアプリのAppconfig対応

MattermostのモバイルアプリがEMMの[Appconfig](https://www.appconfig.org)に対応しました。

[AppConfig for EMM Solutions with Mattermost Mobile Apps — Mattermost 4.3 documentation](https://docs.mattermost.com/mobile/mobile-appconfig.html)

私はAppconfigを利用したことがないので知識があまりないですが、エンタープライズでのモバイルアプリを利用する際の初期設定やセキュリティの管理などを行なってくれるサービス?のようです。  


## コミュニティによる連携アプリ

### Listen for Webhooks and Post Them to Your Mattermost Server

[PromoFaux/Matterhook.NET](https://github.com/PromoFaux/Matterhook.NET)

DiscourseやGitHub,DockerHubのWebHookを受け取り、整形してMattermostのIncoming Webhook経由で投稿を行なってくれるアプリのようです。C#。

### Report TeamCity Build Progress

JetBrains社のCIサーバーである[TeamCity](https://www.jetbrains.com/teamcity/)のビルド結果を通知してくれる連携機能のようです。
他にもBitBucket、Sentryと連携できるようです。
[CGenie/mattermost-openresty: Bridge between Mattermost and various services using the Openresty platform](https://github.com/cgenie/mattermost-openresty#teamcity-integration)

上記はOpenRestyに対応しており、OpenRestyとは`ngx_lua`を含んだ形で提供されているnginxのディストリビュート?のようです。

参考
- [OpenResty® - Official Site](https://openresty.org/en/)
- [３つのnginxをうまく使い分けよう〜nginx、OpenResty、Tengine〜 - Mercari Engineering Blog](http://tech.mercari.com/entry/2016/05/25/170108)

### Unsplash – Inspire a Colleague!

ストックフォトサイト(写真画像共有サイト?)[Unsplash](https://unsplash.com/)と連携するアプリのようです。

[Markus Busche / image_quote_to_mattermost · GitLab](https://gitlab.com/m-busche/image_quote_to_mattermost#imagequote)

### Logback Appender

Javaのロギングライブラリである[Logback](https://logback.qos.ch/)と連携し、Incoming Webhookを通じてログ情報をMattermostへ送信するLogBack拡張のようです。

[elexis-server/bundles/at.medevit.logback.mattermost at master · elexis/elexis-server](https://github.com/elexis/elexis-server/tree/master/bundles/at.medevit.logback.mattermost)

## その他の更新

その他、更新していなかった間に追加されたトピックについて、ピックアップして紹介します。

### Personal Access Token

[Mattermost API](https://api.mattermost.com/)を利用する際の認証に使用できるアクセストークンが発行できるようになりました。(v4.1 ~)
[Personal Access Tokens — Mattermost 4.3 documentation](https://docs.mattermost.com/developer/personal-access-tokens.html)


### New Mobile App

β版だった新しいモバイルアプリが正式版としてiTunesStore、Google Playからダウンロードできるようになりました。(v4.0〜)
[Mattermost on the App Store](https://itunes.apple.com/us/app/mattermost/id1257222717?mt=8)  
[Mattermost - Google Play の Android アプリ](https://play.google.com/store/apps/details?id=com.mattermost.rn)

### リポジトリの分割

今までMattermostアプリは `mattermost/platform` 上で開発が行われていましたが、先月あたりからサーバーサイド(`mattermost/mattermost-server`)とフロントエンド(`mattermost/mattermost-webapp`)にリポジトリが分割されました。

[mattermost/mattermost-server: Open source Slack-alternative in Golang and React - Mattermost](https://github.com/mattermost/mattermost-server)
[mattermost/mattermost-webapp: Webapp of Mattermost server](https://github.com/mattermost/mattermost-webapp)

## その他

### Uber
Mattermostでは新バージョンのリリース毎に最も大きな貢献をしてくれたコントリビュータをMVPとして表彰していますが、今回のMVPにはUberの開発チームが選ばれていました。

Uberでは社内のコミュニケーションツールとして、Slack/HipChatなどを利用していましたが、スケーラビリティやセキュリティの面で折り合いがつかず、現在ではMattermostをベースとした`uChat`というコミュニケーションプラットフォームを開発しているそうです。

[Why Uber switched from Slack to Mattermost for enterprise collaboration - Mattermost](https://about.mattermost.com/blog/how-uber-uses-mattermost-to-enhance-enterprise-wide-communications/)
[The Road to uChat: Building Uber’s Internal Chat Solution](https://eng.uber.com/uchat/?utm_term=3SPQcbSS7Rv3yei2CY1JZxJnUkmxuwWZuXO5V40&adg_id=218769&cid=10078&utm_campaign=affiliate-ir-Skimbit%20Ltd._1_-99_national_D_all_ACQ_cpa_en&utm_content=&utm_source=affiliate-ir)

### Deprecated Mattermot API v3

過去に利用されていたMattermost API v3は2018/1/16に削除される予定です。

> API v3 endpoints are supported until January 16, 2018. To learn more about migrating to APIv4 endpoints, see https://api.mattermost.com/.


## おわりに

次回のv4.3.0のリリースは2017/10/16を予定しています。


