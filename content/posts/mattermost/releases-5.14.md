---
title: "Mattermost5.14の新機能"
date: 2019-08-29T10:15:49+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2019/8/16にMattermost v5.14.0がリリースされました。
その後すぐにv5.14.1のリリースが予告され、2019/8/28にキーボードアクセシビリティのバグ修正を含むv5.14.1がリリースされました。
この記事では、アップデートの内容について簡単に紹介していきます。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.14: Accessibility improvements, enhanced Jira integration, LDAP group sync upgrade, and more \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-14-accessibility-improvements-enhanced-jira-integration-ldap-group-sync-upgrade-and-more/)
* [Mattermost 5\.14\.1 released \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-14-1-released/)
* [Mattermost Changelog — Mattermost 5\.14 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-14-feature-release)

---

# アップデート内容


## アクセシビリティの改善
Mattermostの画面上をキーボードで移動できるようになりました。

デスクトップアプリ上では `F6`、Webブラウザ上では `Ctrl + F6` で、サイドバー、ヘッダー、投稿表示部分などにカーソルを移動することができます。
また、`TAB`ボタンで画面上のボタンやリンクにカーソルを移動することができ、`ENTER`キーで操作することができます。

![mm5.14_accessibility.gif](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/a4f30ce1-ea97-075e-8396-d55299ed4611.gif)


また、スクリーンリーダーとの相性がさらに良くなりました。

視覚的な障碍のある方にもMattermostを効果的に利用してもらえるよう、ボタンやリンク、アプリ表示部分の読み上げがより正確になりました。

詳しくは [Keyboard Accessibility — Mattermost 5\.14 documentation](https://docs.mattermost.com/help/getting-started/accessibility.html) を参照ください。

Mattermostは米国リハビリテーション法第508条の遵守に取り組んでおり、継続的にアクセシビリティの改善を行なっています。
[Voluntary Product Accessibility Template \(VPAT\) — Mattermost 5\.14 documentation](https://docs.mattermost.com/overview/vpat.html)

## JIRA連携機能の改善

Mattermostに同梱されているJIRAプラグインがアップデートされました。
より効率的に作業を行える多くの機能が追加されています。

* 全ての更新、割り当て、コメントを自動で通知する専用のJIRAチャンネル
  * JIRAの情報を一箇所に集めることで、全ての更新を把握するためにメールに頼る必要が無くなります
* あるJIRAプロジェクトから特定のMattermostチャンネルに通知を簡単に送信する機能
  * チャンネル管理者は、JIRAプロジェクトやイベントに基づいてチャンネルに送信される通知を設定できルため、手動でWebHookを設定する必要が無くなります
* JIRAのIssueを作成、管理、閲覧でき、通知の設定を管理できるスラッシュコマンド

詳しくは [mattermost/mattermost\-plugin\-jira: JIRA plugin for Mattermost 🔌](https://github.com/mattermost/mattermost-plugin-jira#jira-21-features) を参照してください。

## LDAPグループ管理 (E10/E20)

LDAPグループ同期機能によるチームとチャンネルの管理に関して新たなオプションが追加されました。

今回追加されたオプションでは、Mattermost v5.12で追加されたCLIのグループコマンドに加え、システムコンソールに追加された新しいチーム/チャンネルのページからLDAPグループ同期によるメンバーシップの管理を行うことができるようになりました。

詳しくは [Using AD/LDAP Synchronized Groups to Manage Team or Private Channel Membership — Mattermost 5\.14 documentation](https://docs.mattermost.com/deployment/ldap-group-constrained-team-channel.html) を参照してください。


## 未読メッセージが見つけやすく

チャンネルを開いたときに、もっとも古い未読メッセージが自動で表示されるようになりました。これにより未読メッセージを追うために投稿を遡ってスクロールする必要が無くなりました。


## その他

### IEサポートの廃止

2019/10/16リリース予定のMattermost v5.16からInternet Explorerのサポートが廃止されます。
[End of support for Internet Explorer 11 coming in October](https://mattermost.com/blog/mattermost-5-13-community-plugins-devops-integrations-series-b-announcement-and-more/#ie11)

Mattermostがサポートしているブラウザは [Software and Hardware Requirements](https://docs.mattermost.com/install/requirements.html#pc-web) から確認できます。


### 破壊的変更
* 内向き/外向きのウェブフック編集画面で、他のユーザーが作成したウェブフックを見ることができなくなりました。システム管理者は全てのウェブフックを見ることができます
* Google Single Sign-On (E20) の機能を利用している場合、ログインに利用するサービスが Google+ から Google People に変更されたため、設定の変更が必要です。詳しくは [Google Single Sign\-On \(E20\)](https://docs.mattermost.com/deployment/sso-google.html) を参照してください。

### インストール済みプラグインの追加
Jenkins, Antivirus, GitLabのプラグインがデフォルトインストール済みプラグインとして追加されました。
それぞれのプラグインの機能については下記を参照してください。

* [mattermost/mattermost\-plugin\-jenkins: A Mattermost plugin to interact with Jenkins](https://github.com/mattermost/mattermost-plugin-jenkins)
* [mattermost/mattermost\-plugin\-antivirus: Antivirus plugin for scanning files uploaded to Mattermost](https://github.com/mattermost/mattermost-plugin-antivirus)
* [mattermost/mattermost\-plugin\-gitlab: GitLab plugin for Mattermost](https://github.com/mattermost/mattermost-plugin-gitlab)


### ウェブフック投稿のアイコンに絵文字

内向きのウェブフックのペイロードに `icon_emoji` というフィールドが追加され、ここに絵文字名を指定してウェブフックを飛ばすと、ウェブフックにより作成された投稿のプロフィール画像に絵文字が表示されるようになりました。

```
curl -v -XPOST -H 'Content-Type: application/json' -d '{"text": "hello", "icon_emoji":"+1"}' https://example.com:8065/hooks/XXXXXXXXXXXXXXXXXXXXXXXXXX
```

↑ のようなリクエストを送ると ↓ のような投稿が作成されます。

<img width="288" alt="スクリーンショット 2019-08-30 0.13.37.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/ad7d122e-d755-8dc3-ee94-bd7fac642580.png">


## misc

以下、最近あまり追えてないですがMattermostチームでの気になる動きについて紹介します。

### Plugin Marketplace

最近、新しいバージョンが増えるごとにMattermost本体に同梱されるプラグインが増えるなど、MattermostチームはPlugin機能に力を入れているようです。
その中の大きな動きの一つとして、Mattermost が Plugin Marketplace の開設があります。
https://github.com/mattermost/mattermost-marketplace

全容はわかりませんが、システムコンソール画面でプラグインのURLを指定するだけで、プラグインをインストールできる仕組みを開発しているようです。（一時期masterブランチから起動すると画面が見れましたが、今は見れなくなっているようです）

### Private Cloud (k8s)

もう一つの気になる動きとして、MattermostコアチームのチャットでPrivate Cloudというチャンネルが立っており、k8s関係のツールの開発が進んでいるようです。

* [mattermost/mattermost\-operator: Mattermost Operator for Kubernetes](https://github.com/mattermost/mattermost-operator)
* [mattermost/mattermost\-helm: Mattermost Helm charts for Kubernetes](https://github.com/mattermost/mattermost-helm)

こちらも全容は掴めておりませんが、OSSプロジェクトらしくドキュメントが公開されているため、興味のある方はそちらを参照してください。
https://drive.google.com/drive/folders/1iayrTYRjQXpdqUAB4pbrnn9QQh_g8oNC

## おわりに

次の`v5.15`のリリースは2019/9/16(Mon)を予定しています。
そして機能追加が行われる`v5.16`は恐らく2019/10/16(Wed)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
