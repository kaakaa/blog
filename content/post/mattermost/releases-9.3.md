---
title: "Mattermost 9.3の新機能"
emoji: "🎉"
published: true
date: 2023-12-18T17:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/12/15 に Mattermost のアップデートとなる `v9.3.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog — Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-3-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube eXA8emM97Bo >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体


## Channels: 設定モーダルのデザイン変更

**設定**や**プロフィール**メニューを開いた時に表示されるモーダルのデザインが変更されました。

左が旧バージョン(Mattermost v9.2)で、右が新バージョン(Mattermost v9.3)になります。  

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-modal-design.png)

## Channels: 最新の投稿に移動するボタンの表示

チャンネルの投稿を過去に遡った際に、最新の投稿まで1クリックで戻ることができるボタンが表示されるようになりました。  

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-jump-to-recent.png)

## Channels: 1投稿に対するリアクション種別の上限設定

一つの投稿に対して付与できるリアクション種別の上限が設定できるようになりました。  
リアクション種別が多すぎることによるパフォーマンス低下を防ぐための設定です。

この設定は、リアクション"**数**"ではなく、リアクション"**種別**"に対する上限値です。  
例えば、上限値を`5`に設定した場合、5種類のリアクションに対して計10個のリアクションを付与するのは問題ないですが、6種類目のリアクションを付与しようとするとエラーとなります。

**システムコンソール > サイト設定 > 投稿 > Unique Emoji Reaction Limit**から設定できます。初期値は`50`に設定されています。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-unique-emoji-settings.png)

上限値を超える種別のリアクションを付与しようとすると、以下のモーダルが表示され、リアクションを付与できません。(以下の例は**Unique Emoji Reaction Limit**を`5`に設定し、6種類目のリアクションを付与しようとした際に表示される画面です。)  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-unique-emoji-error.png)

## (Professional/Enterprise) Channels: 通知なしのキーワードハイライト

今まで、**通知のトリガーとなるキーワード**に設定された単語は、その単語を含む投稿が行われた際にMattermostから通知が送信されると共にMattermost画面上でハイライト表示されていましたが、今回のバージョンから、通知を行うことなくMattermost画面上でのハイライト表示のみを行うキーワードを設定できるようになりました。  
例えば`"AI"`のような、その単語を含む投稿は注目はしておきたいものの、そのような投稿が数多く行われることが予想されるために一々通知はして欲しくない、というような場合に使用できる機能です。

**設定 > 通知 > ハイライトされるキーワード（通知はされません）** から設定できます。  

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-keyword-highlight-settings.png)

設定したキーワードは、以下のようにハイライトされます。  

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-keyword-highlight.png)

この機能は、有償版(Professional/Enterprise)限定の機能です。


## その他の変更

> Removed all uses of the `ExperimentalTimezone` setting. The Timezone feature is now always enabled and no longer behind a configuration setting.

今までのバージョンですは、**システムコンソール > 実験的機能 > 機能 > タイムゾーン**を`有効`に設定した場合に限り、各ユーザーの設定画面で自身のタイムゾーンを設定することができましたが、v9.3から正式な機能となっため、システムコンソールの**タイムゾーン**設定が削除され、すべての環境でユーザーがタイムゾーンを設定できるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.3/channels-timezone.png)

> Added support for previewing WebVTT attachments.

WebVTT形式のファイルプレビューに対応しました。

---

その他、パフォーマンス改善やログ出力改善等の変更が数多くあります。  
詳しくは[公式のChangelog](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-3-feature-release)を参照ください。

## その他のトピック

### Hacktoberfest 2023 

10月に行われていた[Hacktoberfest 2023](https://hacktoberfest.com/)のWrap upが以下の記事にまとめられています。  
期間中、Mattermostとしては80人のコントリビューターによる159のPull Requestがマージされたようです。

[Hacktoberfest 2023: Incredible community contributions, digital rewards & a healthier planet \- Mattermost](https://mattermost.com/blog/hacktoberfest-2023-incredible-community-contributions-digital-rewards-a-healthier-planet/)

### MySQL v5.7のサポート停止

2023/10月に End of Lifeを迎えたMySQL 5.7について、2024/02月リリース予定のMattermost v9.5からサポートを停止するため、MySQL v8系へのアップグレードを推奨しています。

MySQL v8にアップグレードすることの影響等は以下の記事で説明されています。  
[MySQL 5\.7 reached EOL\. Upgrade to MySQL 8\.x today \- Mattermost](https://mattermost.com/blog/mysql-5-7-reached-eol-upgrade-to-mysql-8-x-today/)

### Mattermost Enterprise版の紹介

Mattermostの有償版であるEnterprise Editionの機能を使うことで、どのようにワークフローを改善できるかについて以下の記事で紹介されています。  

* [Enterprise Collaboration with Advanced Workflows \- Mattermost](https://mattermost.com/blog/enterprise-collaboration-with-advanced-workflows/)
* [Mattermost Playbooks for Enterprise Workflows \- Mattermost](https://mattermost.com/blog/mattermost-playbooks-for-enterprise-workflows/)



## おわりに
次の`v9.4`のリリースは 2024/1/16(Tue)を予定しています。  
