
---
title: "Mattermost4.4の新機能"
date: 2017-12-24T13:38:39+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

2017/11/16にMattermost 4.4.0がリリースされたので、追加された機能などについて簡単に紹介します。(投稿日現在(2017/12/24)の最新版は`v4.5.0`です)

詳細については公式のリリースを確認ください。
[Mattermost 4.4: Plugin support in beta, plus new “Do Not Disturb” mode - Mattermost](https://about.mattermost.com/blog/mattermost-4-4/)
[Mattermost Changelog — Mattermost 4.5 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-4-5)

`v4.4.0`リリース以降、セキュリティアップデートやバグ修正が加えられ、4.4系の現在の最新リリースは`v4.4.5`となっています。
この記事では`v4.4.0`リリース時に加えられた新機能のみについてのみ記述しています。その他の修正内容については、上記のChangelogを確認してください。

## プラグイン機能

Mattermostのサーバーサイド、クライアントサイドのそれぞれに独自の機能を追加するためのプラグイン機能が使えるようになりました。

**クライアントプラグイン**では、プロフィールポップオーバーやサイドバーなど、画面上の要素を上書きすることができます。
コアメンバーによるサンプル実装が下記で公開されています。Reactのコンポーネントを上書きしているようです。
[jwilander/hovercardexample: Example hovercard webapp plugin for Mattermost](https://github.com/jwilander/hovercardexample)

**サーバープラグイン**では、JIRAやGitLab,Jenkinsなどのサードパーティーシステムと容易に連携できるようになります。
現在のMattermostでは、[JIRA Webhook Plugin (Beta)](https://docs.mattermost.com/integrations/jira.html#)がプリインストールされています。

作成したプラグインは`tar.gz`の形式で圧縮し、画面UIからアップロードすることもできます。
[Plugins (Beta) — Mattermost 4.5 documentation](https://docs.mattermost.com/administration/plugins.html)

## 取り込み中(Do Not Disturb)ステータス

今まであったステータスである`オンライン(Online)`、`離席中(Away)`、`オフライン(Offline)`に加え、集中したい時間など通知を受け取りたくない場合のステータスである`取り込み中(Do Not Disturb / DND`が追加されました。

画面左上のユーザーアイコンかスラッシュコマンド `/dnd` より設定できます。

<img width="355" alt="スクリーンショット 2017-12-24 13.07.56.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/11f85fd3-b515-63c8-e9bb-cae450e335b4.png">

## E20: SAML AuthenticationによるAD/LDAP同期

AD/LDAPによる認証を有効にしている場合に、AD/LDAP上で無効になったユーザーを自動でMattermostでも無効にできるようになりました。
この機能はSlackにはないMattermost独自の機能だとのことです。

## New Open Source Projects Built on Mattermost

オープンソースのMattermostをベースとしたコミュニティによるアプリ・システムが100を超えたとのことです。
公開されている連携機能は下記のページから探すことができます。
[Apps and Integrations - Mattermost](https://about.mattermost.com/community-applications/)

Mattermostと連携する機能を作成した場合、下記より報告すると上記のページで紹介されます。
[Share your Mattermost projects | Mattermost](https://www.mattermost.org/share-your-mattermost-projects/)

拙作の[投票機能もどき](https://qiita.com/kaakaa_hoe/items/b2605ce3816cfc517ecd)も掲載してもらっています。

### Mattersend Python Integration

Mattermostの内向きのWebhookを通じてテキストを送信するための Python3 製CLIツールです。
[joho1968/mattermostsendpy: Simple CLI utility for posting text to a Mattermost Incoming Webhook, written in Python 3](https://github.com/joho1968/mattermostsendpy)

### Sensu Monitoring Service
[Sensu](https://sensuapp.org/)によるヘルスチェックやKPIなどのモニタリングをMattermost上で行える連携アプリです。
[mattronix/handler-mattermost: mattermost plugin for sensu](https://github.com/mattronix/handler-mattermost)

### Spy Bot Integration
あるユーザーがオンラインになったかどうかを監視するBotのようです。どのような場合に利用するのだろう...？
[prabhu43/mattermost-spybot: Spybot for mattermost](https://github.com/prabhu43/mattermost-spybot)


## おわりに

次回の`v4.5.0`のリリースは2017/12/16に行われました。

---

拙作の紹介ばかりな気がしますが、[Interactive Message Buttons](https://docs.mattermost.com/developer/interactive-message-buttons.html)の機能を利用したサンプルアプリの紹介を下記のページで行なっているので興味があればご覧ください。
[打ち合わせ みんな多忙で 草生えるwww \- Qiita](https://qiita.com/kaakaa_hoe/items/7df3789f1fc57ffa02fc)

