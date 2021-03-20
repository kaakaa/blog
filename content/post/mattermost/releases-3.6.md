
---
title: "Mattermost 3.6の新機能"
date: 2017-01-31T15:05:21+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに
2017/01/16にMattermost 3.6.0がリリースされましたので、追加された機能などについて簡単に紹介します。
(その後1/19にアップグレードにまつわる問題などを改善した[v3.6.1](https://about.mattermost.com/mattermost-3-6-1/)がリリースされています)
(その後2/3にTeamEdition, EnterpriseEdition向けのセキュリティアップデートを含んだ[v3.6.2](https://about.mattermost.com/mattermost-3-6-2/)がリリースされています)

詳細については公式のリリースを確認ください。
[Mattermost 3\.6 improves multi\-team deployment, performance, emoji reactions, CLI and much more – Mattermost](https://about.mattermost.com/mattermost-3-6/)

# changelog

## チームサイドバー

複数のチームに所属している場合に、画面左端にチームサイドバーが表示されるようになりました。
他のチームでの新しい投稿やメンションの有無が、一画面で確認しやすくなります。
![mattermost_teambar.png](https://qiita-image-store.s3.amazonaws.com/0/9891/2c0b7fd9-9666-a75b-1eeb-3ce0cdd4a3b0.png)


## Emoji リアクション

投稿に対してEmojiでのリアクションを付けられるようになりました。
チャンネルの最新の投稿、もしくは返信スレッドの末尾の投稿に対してリアクションすることができます。
投稿欄に`+:thumbsup:`のように絵文字コマンドの前に`+`をつけて投稿することでEmojiリアクションとなります。Emojiにはカスタム絵文字を指定することも可能です。
![mattermost-reaction.png](https://qiita-image-store.s3.amazonaws.com/0/9891/575df51f-347f-865b-4029-ba1aa4d6f4e8.png)

他のユーザーがEmojiリアクションをクリックすることでリアクション数をカウントアップできるため、投票機能としても利用することができます。

v3.6時点でEmojiリアクションはチャンネルまたは返信スレッドの最新の投稿にしか付与できず、Slackのように任意の投稿を選んでリアクションをつけることはまだできないようです。

## コマンドラインインタフェースの改善

今までは`platform`コマンドのオプションとして実行するコマンドを指定していましたが、v3.6からは用途ごとのサブコマンドを指定する形式になりました。
[Command Line Tools — Mattermost 3\.6 documentation](https://docs.mattermost.com/administration/command-line-tools.html)
## グローバル対応

CJK(Chinese, Japanese and Korean)を含む英数字以外のユニコード文字でのハッシュタグに対応しました。
![Off-Topic_-_test_Mattermost.jpg](https://qiita-image-store.s3.amazonaws.com/0/9891/435a3914-8d57-e782-ab97-e9e54620329e.jpeg)

また、各言語へのローカライゼーションも引き続き行われています。
日本語のローカライゼーションに興味のある方はいつでもご参加ください。
[Localization — Mattermost 3\.6 documentation](https://docs.mattermost.com/developer/localization-process.html)
[MattermostのLocalizationフロー \- Qiita](http://qiita.com/kaakaa_hoe/items/ac79289d65f7a3e3d9ea)


## (Enterprise Edition E20) パフォーマンスモニタリング

有償版限定ですが、PrometheusやGrafanaと統合されたパフォーマンスモニタリング機能が付いたようです。  
私はTeamEditionしか使っていないので詳細は分かりません。

## New Integrations

公式ページでも紹介されている統合機能は引き続き追加されています。
[Share your Mattermost projects | Mattermost](https://www.mattermost.org/share-your-mattermost-projects/)

### GitLab Auto Deploy with Mattermost

GitLab 8.15で追加されたAuto Deploy機能をMattermostのスラッシュコマンドから起動する手順の紹介です。
[Mattermost slash commands \- GitLab Documentation](https://docs.gitlab.com/ee/project_services/mattermost_slash_commands.html)

### AWS AMI with Mattermost, Gitlab and Kanboard
Mattermost, Gitlab, Kanboard全部入りのAMIの紹介です。
[AWS Marketplace: Best Project Management Package](https://aws.amazon.com/marketplace/pp/B01N6M8DP6?qid=1484222369100&sr=0-7&ref_=srh_res_product_title)

### Microsoft System Center Operations Manager Alerts via Mattermost
Microsoft SCOM（や[SolarWinds](http://www.solarwinds.com/ja/), [OMS](https://www.microsoft.com/ja-jp/cloud-platform/operations-management-suite), [datadog](https://www.datadoghq.com/lpg/?utm_source=Advertisement&utm_medium=GoogleAdsNon1stTierBrand&utm_campaign=GoogleAdsNon1stTierBrand-NonENES&utm_content=Datadog&utm_keyword=%7Bkeyword%7D&utm_matchtype=%7Bmatchtype%7D&gclid=COjTwaHC5tECFUVwvAod-0cG1w)など）の監視ツールからの通知をMattermostへ送信する機能の紹介です。

[SCOM Alerts to Microsoft Teams and Mattermost \| adatum](http://adatum.no/operationsmanager/scom-alerts-to-microsoft-teams-and-mattermost)

### PHP Driver for Mattermost
内向きのウェブフックを通じてMattermostを操作できるPHP製ドライバーの紹介です。
[GitHub \- ThibaudDauce/mattermost\-php: Mattermost PHP driver to send incoming webhooks](https://github.com/ThibaudDauce/mattermost-php)

# Coming Soon

## 次世代Mobile Apps

現在のWebViewベースのモバイルアプリからReact Nativeベースのアプリに移行する予定です。
Q1でのEarly Releaseを予定しています。

## Desktop App 3.6

Windows, Linux, Mac向けのデスクトップアプリのリリースも予定しています。

* inline spell checker
* team sidebar
* performance improvements
* (Enterprise Edition) SAML OneLogin / Google authentication

# おわりに

Mattermostは隔月16日リリースを行っているため、次回のv3.7のリリースは2017/3/16を予定しています。
[Roadmap – Mattermost](https://about.mattermost.com/direction/)

