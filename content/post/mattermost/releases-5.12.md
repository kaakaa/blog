---
title: "Mattermost5.12の新機能"
date: 2019-06-16T10:15:49+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---
# はじめに

2019/6/13にMattermost v5.12.0がリリースされました。
この記事では、アップデートの内容について簡単に紹介していきます。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.12: Infinite scrolling, bot accounts, deeper Jira integration, and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-12-infinite-scrolling-bot-accounts-deeper-jira-integration-and-more/)
* [Mattermost Changelog — Mattermost 5\.12 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-12-feature-release)

---

# アップデート内容


## Mattermostへのフィードバック

Mattermostの改善を目的とした、ユーザーへのアンケート機能が追加されました。
これにより、Mattermostの使い勝手に関するアンケートメッセージがBotから送信されるようになります。

![image.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/bc07b62b-7687-bc8e-5195-4e6c9ee2eb17.png)

Mattermostを5.12にアップデートすると、実際にアンケートが送信される前に、システム管理者に対してアンケートが予定されていることを伝える電子メールかダイレクトメッセージが送信されると思います。アンケートは3週間後に設定されているはずです。

<img width="321" alt="スクリーンショット 2019-06-16 21.30.29.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/61664110-5024-9def-ea3a-ea8e5a7da5ab.jpeg">

現在、アンケートは英語で送信されるため、アンケートが送信されるとユーザーが混乱してしまう場合など、アンケートをを実施したくない場合は、**システムコンソール > プラグイン > Net Promoter Score > Enable Net Promoter Score Survey**を`無効`に設定してください。


## 無限スクロール

チャンネル内の古い投稿を見るとき、`もっとメッセージを読み込む`というリンクをクリックする必要がなくなりました。チャンネルの一番上までスクロールすると、自動で古いメッセージの読み込みが開始されます。

(この機能はIE11では利用できません)

## Botアカウント

MattermostでBot用のアカウントを作成できるようになりました。Botアカウントは電子メールやパスワード、またはその他の認証方法を必要とせずにユーザーとやり取りができるアカウントです。これにより、統合機能用のダミーのアカウントを作る必要がなくなります。

さらに、Botアカウントは下記のような特徴を持ちます。

* 非公開チーム、非公開チャンネル、ダイレクトメッセージなど、どのチャンネルにも投稿することができます
* 統合機能をBotアカウントに関連づけておくことで、その統合機能を開発したユーザーが退職した後も統合機能を使い続けることができるようになります
* Botアカウントはサーバー上の誰もがやり取り可能で、どのチーム、チャンネルにも追加することができます

![image.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/0435acc1-99bf-4483-4a1e-f339f580bb71.png)

また、Enterprise版においてBotアカウントはライセンス上のユーザーとしてカウントされません。

Botアカウントについて詳しくは下記のドキュメントを参照してください。
https://docs.mattermost.com/developer/bot-accounts.html

## Jiraプラグインのアップデート

MattermostにバンドルされているJiraプラグインのバージョンが2.0になりました。このアップデートにより、MattermostとJiraをより深く統合することができるようになりました。

* Issue作成、更新、コメント追加をMattermostチャンネルに通知
* Jira Issueの作成や、Mattermost上のコメントのJira Issueへの反映、また、Issueの状態の更新などをMattermostから実行

より詳しい情報は、下記を参照してください。
https://github.com/mattermost/mattermost-plugin-jira#jira-20-features

## デフォルトプラグインの追加

Mattermost 5.12から、さまざまなプラグインがMattermostにデフォルトでバンドルされるようになりました。

* [**GitHub Plugin**](https://github.com/mattermost/mattermost-plugin-github): GitHubの通知やリマインドをMattermostに投稿することができるプラグイン。SaaS版やオンプレミスのGitHub Enterpriseにも対応しています。
* [**Autolink Plugin**](https://github.com/mattermost/mattermost-plugin-autolink): あるテキストを自動でハイパーリンク化することができるプラグイン。(例えば、Issue TrackerのIssue番号を投稿すると、自動でリンクに変換してくれます)
* [**Custom Attributes Plugin**](https://github.com/mattermost/mattermost-plugin-custom-attributes): ユーザープロフィールのポップオーバーに属性を追加するプラグイン
* [**Welcome Bot Plugin**](https://github.com/mattermost/mattermost-plugin-welcomebot): 新規メンバーが追加された時にメッセージを通知するプラグイン
* [**Amazon SNS CloudWatch plugin**](https://github.com/mattermost/mattermost-plugin-aws-sns): [Amazon CloudWatch](https://aws.amazon.com/jp/cloudwatch/)のアラートをAWS SNS経由でMattermostに通知できるようになるプラグイン。

他にも下記サイトから様々なMattermostの拡張機能を検索することができます。
https://integrations.mattermost.com/

## システムコンソール画面の再構築

Mattermost 5.12からシステムコンソール画面の構成が変わりました。

この変更は、Mattermost Private Cloudなど、システムコンソールを使わず[環境変数でMattermostの設定を指定する](https://docs.mattermost.com/administration/config-settings.html#configuration-settings)際に、よりロジカルにしたいという意図があるそうです。

<img width="161" alt="スクリーンショット 2019-06-16 21.30.29.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/545f675b-73b1-aa16-4f71-fb6a7078b138.png">


## E20版のみの追加機能

* 非公開チーム/チャンネルに関する権限の管理を個別に操作することなく、AD/LDPグループを通じて管理できるようになりました。
* Elasticsearchの自動補完機能により検索機能を利用する際にユーザー/チャンネルが入力しやすくなりました

## モバイルアプリのタブレット向け改善

モバイルアプリ v1.20もリリースされています。

* チャンネルサイドバーが表示され続けるようになりました
* iOSでキーボードが表示されている際、スワイプダウンすることでキーボードを閉じることができるようになりました

そのほかの変更については下記を参照ください。
https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#1200-release


## 破壊的変更

Mattermost v5.10以前のバージョンからアップデートをする場合、統合機能に関する下記の破壊的変更があります。
* (v5.12~) プラグインAPI `DeleteEphemeralMessage` の引数が`(userId string, post *model.Post)`から`(userId, postId string)`に変更
* (v5.11~) 統合機能で既に投稿されている投稿の内容を更新する際、投稿のProps属性を`Update.Props == nil` のように`nil`を指定してクリアする方法は利用できなくなりました。代わりに`Update.Props == {}`のように、空の要素を指定する必要があります。

# おわりに

次の`v5.13`のリリースは2019/7/16(Tue)を予定しています。
そして機能追加が行われる`v5.14`は恐らく2019/8/16(Fri)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
