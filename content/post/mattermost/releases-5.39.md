---
title: "Mattermost 5.39の新機能"
emoji: "🎉"
published: true
date: 2021-09-17T11:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/09/16 に Mattermost `v5.39.0` がリリースされました。
`v5.39`はいつものような **Feature Release** ではなく、**Quality Release**となっているため、新機能の追加は無く、改善系の変更のみのリリースとなっています。

先月の記事で、今月はメジャーバージョンアップであるv6.0のリリース日であると書きましたが、**v6.0のリリースは来月(10/13)だった**ようです。（リリース日については、いつもMattermost公式チャットの[Release: Self-Managed](https://community-daily.mattermost.com/core/channels/release-discussion)チャンネルのチャンネルヘッダーを見て書いているのですが、見間違えたか変更されたかしていたようです）

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.39 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v5-39/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v5-39-quality-release)

---


## 返信スレッドの折り畳み機能の改善

2021年7月にリリースされたMattermost v5.37にてベータ版として導入され、先月のリリースからモバイルアプリでも利用可能になった返信スレッドの折り畳み機能のバグ修正がいくつか行われています。  
この機能を利用している場合は、バグ修正を適用するために v5.39 へアップデートすることをお勧めします。

現在ベータ版の本機能は、2021/4Q (2021年10月~12月) にベータ版ではなくGeneric Availableとしてリリースされる予定だそうです。
[Looking ahead to general availability of Collapsed Reply Threads](https://mattermost.com/blog/collapsed-reply-threads-ga/)


## Mattermost v6.0が10月にリリース

冒頭でも触れましたが、来月10/13(Wed)にMattermostのメジャーバージョンアップである v6.0 がリリースされます。
`v6.0`では今までベータ版としてリリースされていた以下のような機能が、Generic Availableとしてリリースされます。
 
- [Archived Channels](https://docs.mattermost.com/configure/configuration-settings.html#allow-users-to-view-archived-channels-beta): アーカイブされたチャンネルの内容を検索する機能
- [Compliance Exports(E20)](https://docs.mattermost.com/comply/compliance-export.html): コンプラインアンス向けデータエクスポート機能
- [Custom Term of Service(E20)](https://docs.mattermost.com/comply/custom-terms-of-service.html): 独自の利用規約設定機能
- [Guest Accounts(E20)](https://docs.mattermost.com/onboard/guest-accounts.html): ゲストアカウント機能
- [mmctl](https://docs.mattermost.com/manage/mmctl-command-line-tool.html): 管理者向けMattermost管理CLIツール
- [Additional System Admin Roles(E20)](https://docs.mattermost.com/onboard/system-admin-roles.html): システム管理ロールの追加
- [Plugin](https://developers.mattermost.com/integrate/admin-guide/admin-plugins-beta/): Mattermostプラグイン機能

その他 `v6.0` で変更が予定されている点は、以下の公式ブログエントリにまとめられています。

[Looking forward to the next big Mattermost product milestone: Mattermost v6.0](https://mattermost.com/blog/looking-forward-to-mattermost-v6-0/)

---

Mattermost v6.0へのアップグレード時にDBのマイグレーションが実行されますが、このマイグレーション処理に時間がかかることが予想されます。マイグレーションにかかる時間に関する分析が [Mattermost v6.0 DB Schema Migrations Analysis](https://gist.github.com/streamer45/59b3582118913d4fc5e8ff81ea78b055)で公開されているため、アップデート作業時の参考にご利用ください。

---

デスクトップアプリ、モバイルアプリに対しては後方互換性が保たれているため、現在のバージョンのままでもMattermost v6.0を利用できるようです。ただし、どちらもMattermost v6.0と同時期にリリースされる最新バージョンへのアップデートが推奨されています。

---

Extended Support Release (ESR) となっている Mattermost v5.37 は、引き続き2021年年4月までサポートされます。

## プラグインの更新

Mattermost本体の機能追加がないためか、最近 (?) 更新のあったプラグインが紹介されています。どちらのプラグインもプラグインマーケットプレースから簡単にインストールすることができます。

### Webex v1.2

[Cisco Webex](https://www.webex.com/ja/video-conferencing.html)と連携するプラグインについて、Linuxクライアントを使用している場合のMeeting URLの改善をはじめ、いくつかの改善が加えられた新たなバージョンがリリースされています。

https://github.com/mattermost/mattermost-plugin-webex/releases/tag/v1.2.0


### Gif Commander v2.1.0

アニメーションGifをMattermostに投稿するプラグインについてもいくつか改善が行われています。

https://github.com/moussetc/mattermost-plugin-giphy/releases/tag/v2.1.0

## Mattermostコミュニティイベント

Mattermostのメンバーが今後数週間のうちにいくつかのイベントに登壇する予定のようです。

Mattermostのイベント登壇を含むPublic向けのイベント情報は以下のカレンダーから確認できます。
[Mattermost Public Events Calendar](https://mattermost.com/events/public/)

また、今年も[Hacktoberfest](https://hacktoberfest.digitalocean.com)への参加を予定しているそうです。

---

## おわりに

3年半ぶりのメジャーバージョンアップとなる、次の`v6.0`のリリースは 2021/10/13(Web) を予定しています。

[Mattermost公式コミュニティチャット](https://community-daily.mattermost.com/)の方では、一足先にMattermost v6.0を触ってみることができます。まだ開発途中のため、実際にリリースされるものとは異なりますが、雰囲気を感じることができると思います。
個人的には、メインメニューが画面上部にヘッダとしてまとめられてスッキリとしている印象があったのと、デフォルトのテーマが `Denim` になっておりシックになった感じがしました。

![v6](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.39/mattermost-v6.jpeg)

また、画面左上のメニューからは、最近開発に力を入れている[Playbook](https://github.com/mattermost/mattermost-plugin-playbooks)や[Focalboard](https://www.focalboard.com/)へ移動するメニューが追加されており、v6では、チャットをベースとしたDevOpsプラットフォームとしての発展に力を入れていくのかなと思いました。

![menu](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.39/mattermost-menu.jpeg)


---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
