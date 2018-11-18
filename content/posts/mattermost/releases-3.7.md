
---
title: "Mattermost3.7の新機能"
date: 2017-03-30T22:58:23+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに
2017/03/16にMattermost 3.7.0がリリースされました。
リリースから２週間ほど経ってしまいましたが、追加された機能などについて簡単に紹介します。
現在、セキュリティfixなどを含んだ v3.7.3 が最新バージョンとなっています

詳細については公式のリリースを確認ください。
[Mattermost 3.7: Auto-deploy to cloud, multi-person direct messaging and much more](https://about.mattermost.com/mattermost-3-7/)

その他、細かい変更や修正は下記にまとめられています。
[Mattermost Changelog](https://docs.mattermost.com/administration/changelog.html)

# 変更点

## Cloud Platform Auto-deploy

[Deployment Solutions Programs](https://docs.mattermost.com/guides/orchestration.html)が公開されています。
(Mattermostのデプロイメントを体系的にしたもの？）

また、[Bitnami Mattermost Team Edition Stack](https://bitnami.com/stack/mattermost)によってMattermostのTeam Edition([OSS版のこと](https://about.mattermost.com/pricing/))の主要パプリッククラウドサービスへのデプロイが容易になりました。
Azureのプライベートクラウドプラットフォームである[Azure Goverment Cloudにも適用できるようです。](https://about.mattermost.com/open-source-slack-alternative-for-azure-government-plus-aws-gcp-and-ocp-via-bitnami-mattermost-team-edition-stack/)

BitnamiからAzureへMattermostをデプロイし、Slackからデータをインポートする手順についての[動画](https://www.youtube.com/watch?v=AKqHWqrAgpk)が公開されています。

## [Group Messaging for Quick, Direct Chats](https://about.mattermost.com/mattermost-3-7/#group-messaging)

素早く複数人でのメッセージングを行う **グループメッセージ** の機能が追加されました。
**ダイレクトメッセージ** の機能拡張のようなものです。

以前のダイレクトメッセージ機能では１対１の会話のみが可能でしたが、**グループメッセージ**では最大8人での会話が可能です。
8人を超える場合は **非公開グループ** の作成が推奨されています。

## [More Flexibility](https://about.mattermost.com/mattermost-3-7/#more-flexibility)

### [API Version 4 (early beta release)](https://about.mattermost.com/mattermost-3-7/#api-v4)

Mattermost APIの新バージョンのベータリリースです。
[Mattermost API Reference (v4.0.0)](https://api.mattermost.com/v4/)

現在のAPI v3は **2017/9/16** までサポートされます。

### [New Website Link Previews](https://about.mattermost.com/mattermost-3-7/#link-previews)

Website Linkを投稿すると、リンク先の内容を表示するようになります。
v3.6までのプレビュー機能では日本語が文字化けしていましたが、v3.7からは文字化けしていないようです。

![mm3.7_preview_1.png](https://qiita-image-store.s3.amazonaws.com/0/9891/e1c89b1a-6fb5-c3fd-6ca2-64998abae69b.png)

プレビュー機能を有効にするには、システムコンソールでプレビュー機能を有効にした上で、各アカウントごとに設定します。

![mm3.7_preview_2.png](https://qiita-image-store.s3.amazonaws.com/0/9891/6724d0c2-09e5-b135-4004-9b5be14e511d.png)

### [Channel Push Notification Preferences](https://about.mattermost.com/mattermost-3-7/#channel-mobile-notif)

モバイルプッシュ通知の設定をチャンネルごとに行えるようになりました。
(モバイルプッシュ通知の設定が行われている場合のみ)

![notification.png](https://qiita-image-store.s3.amazonaws.com/0/9891/f1d05b58-b506-e910-6c8f-2573737abd9d.png)

## スケーラビリティの改善

### [Faster Server-side Processing](https://about.mattermost.com/mattermost-3-7/#performance)

サーバーサイドのパフォーマンス改善が行われました。  
キャッシュ・クエリ最適化によるCPU/DBのボトルネック低減や、不要なWebSocketの呼び出しを削除したとのことです。
詳しい改善内容は[Mattermost Changelog - Performance](https://docs.mattermost.com/administration/changelog.html#performance)にまとめられています。

また、パフォーマンス改善について、[動画 Mattermost Developer Brown Bag - Performance](https://www.youtube.com/watch?v=pJsfXAHxAjM&t)で語られています。

### [Simplified Import of Bulk Data](https://about.mattermost.com/mattermost-3-7/#bulk-import)

他のシステムからチーム、チャンネル、ユーザー、投稿をインポートするツールが使えるようになりました。
明示的なリンクはないのですが、おそらく[Bulk Loading Data](https://docs.mattermost.com/deployment/bulk-loading.html)の事だと思います。

## セキュリティ関係の新機能

### [New “Channel Administrator” Role (Enterprise Edition E10 & E20)](https://about.mattermost.com/mattermost-3-7/#channel-admin)

チャンネル名の変更とチャンネルの削除を行うことができる**チャンネル管理者**のロールが追加されました。
Enterprise版(E10,E20)のみの機能となります。

### [OneLogin Support (Enterprise Edition E20)](https://about.mattermost.com/mattermost-3-7/#onelogin)

v3.6までは[Okta](https://docs.mattermost.com/deployment/sso-saml-okta.html)か[Microsoft ADFS](https://docs.mattermost.com/deployment/sso-saml-adfs.html)を利用したSAML SSOを利用できましたが、v3.7から[OneLogin](https://docs.mattermost.com/deployment/sso-saml-onelogin.html)を使用したSAML SSOも行うことが可能になりました。
Enterprise版(E20)のみの機能になります。

## [Second Generation Mobile Apps: Beta Release in March](https://about.mattermost.com/mattermost-3-7/#react-native-app)

React Native製のiOS/Androidアプリのβ版が今月中にリリースされる~~そうです。~~ 3/29にリリースされました。 => [A Native Mobile Experience – Second Generation Mobile Apps Released in Beta](https://about.mattermost.com/a-native-mobile-experience-second-generation-mobile-apps-released-in-beta/)

[Android版](https://play.google.com/apps/testing/com.mattermost.react.native)、[iOS版](https://mattermost-fastlane.herokuapp.com)それぞれβ版を試してみることができます。
(iOS版はTestFlightアプリのインストールが必要でした)

開発は[mattermost/mattermost-mobile](https://github.com/mattermost/mattermost-mobile#how-to-contribute)で行われています。

## 新しいコミュニティ連携

### [Mattermost-Tuleap Bot Plug-in](https://about.mattermost.com/mattermost-3-7/#mattermost-tuleap)

OSSの開発サポートツールである[Tuleap](https://www.tuleap.org)との連携機能です。
TuleapへのGitプッシュとScrumレポートのLattermostへの通知を行えるようです。
[Mattermost connectors for Tuleap Scrum reports and Git now available](https://www.tuleap.org/mattermost-connectors-tuleap-scrum-reports-and-git-now-available)

### [New Simple GitHub Integration](https://about.mattermost.com/mattermost-3-7/#github-integration)

Haskell製のGithub連携ツールです。
[UlfS/ghmm](https://github.com/UlfS/ghmm)

python/flask製のGithub連携ツール [softdevteam/mattermost-github-integration](https://github.com/softdevteam/mattermost-github-integration) のメンテも継続されています。

# おわりに

次のバージョンであるMattermost v3.8.0は2017/4/17リリースを予定しています。
[Pinned Post](https://github.com/mattermost/platform/pull/4217)や[Emoji Picker](https://github.com/mattermost/platform/pull/5157),[Reaction Picker](https://github.com/mattermost/platform/pull/5816)などの機能が追加される予定です。
[Core Team向けMattermost](https://pre-release.mattermost.com/core)では、すでに利用可能です。

