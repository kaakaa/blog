---
title: "Mattermost 9.4の新機能"
emoji: "🎉"
published: true
date: 2024-01-17T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/01/16 に Mattermost のアップデートとなる `v9.4.0` がリリースされました。  
(同日にバグ修正版の`v9.4.1`がリリースされています)  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v9\.4: IP filtering, bring your own key & more\!](https://mattermost.com/blog/mattermost-v9-4-is-now-available/)
- [Mattermost changelog — Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-4-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube bEMp4vYLi6c >}}

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


## Channels: チャンネルの通知設定UIの改善

チャンネルメニューからアクセスできるチャンネル毎の通知設定モーダルのUIが改善されました。以前より少ないクリック数で設定変更が出来るようになっています。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.4/channels-channel-notification.png)

チャンネルの通知設定は、**チャンネル名をクリック > 通知の設定**から開くことができます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.4/channels-channel-notification-menu.png)


## (Cloud Enterprise) Platform: IPフィルタ機能

Mattermost CloudのEnterprise版限定ですが、Mattermostへのアクセス元IPアドレスを指定できるようになりました。  
Company Proxyを利用している場合など、Mattermostユーザーのアクセス元IPアドレスが限定されている組織において、環境をよりセキュアにできるようになります。

詳しくは以下のドキュメントを参照ください。  
[Cloud IP Filtering — Mattermost documentation](https://docs.mattermost.com/manage/cloud-ip-filtering.html)

## (Cloud Enterprise) Platform: BYOKによる暗号化

Mattermost CloudのEnterprise版限定ですが、AWSの各サービスを利用したBring Your Own Key (BYOK)による暗号化がサポートされました。  

詳しくは以下のドキュメントを参照ください。  
[Cloud Dedicated Bring Your Own Key — Mattermost documentation](https://docs.mattermost.com/manage/cloud-byok.html)

## Platform: MySQL v5.7のサポート終了について

MySQL v5.7がEoLを迎えたため、次のバージョンであるMattermost v9.5 ([ESR](https://docs.mattermost.com/upgrade/extended-support-release.html)) から、MySQL v5.7がサポートされなくなります。  
MattermostではMySQL v8以降の使用が推奨されています。

## その他の変更

その他、Administrator関連の機能追加などが行われています。    
詳しくは[公式のChangelog](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-3-feature-release)を参照ください。

## その他のトピック

特になし。

## おわりに
次の`v9.5`のリリースは 2024/2/16(Fri)を予定しています。  
