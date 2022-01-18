---
title: "Mattermost 6.3の新機能"
emoji: "🎉"
published: true
date: 2022-01-18T18:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/01/14 に Mattermost の新しいバージョン `v6.3.0` がリリースされました。  
本バージョンは[Extended Support Release(ESR)](https://docs.mattermost.com/upgrade/extended-support-release.html)としてリリースされており、2022/10/15までセキュリティFixなどのサポートが実施される予定です。(通常リリースのサポート期間は3ヶ月間)

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.3 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v6-3-is-now-available/)
- [Mattermost Self\-Hosted Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-3-extended-support-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## (Enterprise) Playbook: 詳細な権限スキーム

Playbookについての権限設定を詳細に行えるようになりました。

**システムコンソール > ユーザー管理 > 権限**で設定可能な権限一覧に、Playbooksに関する権限項目が表示されるようになりました。(権限設定について詳しくは[Advanced Permissions](https://docs.mattermost.com/onboard/advanced-permissions.html)を参照)

![playbooks permission](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.3/playbooks-permission.png)

また、Playbookに関する新たなロールである `Playbook Administrator`(`プレイブック管理者`) が追加され、Playbookに対して可能な操作をユーザー単位で設定することができるようになりました。

![playbooks administrator](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.3/playbooks-administrator.png)

ユーザーに対するロールは、Playbook編集画面で設定することができます。

![playbooks role](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.3/playbooks-role.png)

Playbooksの権限やロールについては、[Roles and Permissions](https://docs.mattermost.com/playbooks/playbook-permissions.html)を参照してください。

## Playbooks: 翻訳

Playbook機能もメッセージの翻訳が行われており、すでに12を超える言語で翻訳がスタートしています。日本語への翻訳も既に一部対応しています。

![playbook-translations](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.3/playbooks-translations.png)

## Playbooks: 通知の改善

通知機能のアップデートにより、チャンネルのノイズを減らし、重要な通知のみを受け取れるようになりました。  
チャンネル内に投稿される通知は全て削除され、タスクが割り当てられた場合や実行オーナーに設定された場合などの重要な通知はPlaybooks Botからダイレクトメッセージとして通知されるようになりました。

Playbooksの通知について詳しくは[Notifications and Updates](https://docs.mattermost.com/playbooks/notifications-and-updates.html)を参照してください。

## Boards: General Availability

Board機能が正式にGeneral Availability(GA)な機能となりました。  
2021年にオープンソースプロジェクトとして([Focalboard](https://github.com/mattermost/focalboard))開始されて以降、159人のコントリビューターと7,200を超えるGitHubスターを獲得しており、Trello/Notion/Asanaなどのプロプライエタリなプロジェクト管理ツールの競合となるオープンソースプロジェクトの中で最もポピュラーなものの一つになっています。

## Boards: 更新通知

カードの更新を見逃さないようにするため、カードをフォローすることが出来るようになりました。  
カードをフォローすることで、そのカードに対する変更内容をMattermostメッセージとして受け取ることができます。自分が作成したカードや@メンションされたカードは自動でフォローされます。

(ローカル環境のMattermostに[`v0.12.0`](https://github.com/mattermost/focalboard/releases/tag/v0.12.0)のPlugin版をインストールしてみましたが、うまく動作しませんでした。)

## Boards: アバター表示

カード内の**人物**プロパティを選択する際、名前の横にアイコン画像が表示されるようになりました。アイコンはコメント入力画面で`@`を入力した際も表示されます。

![boards-avatar](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.3/boards-avatar.png)

## その他の変更
特になし

## その他のトピック

### Desktop App Contributor Event

[Join the Mattermost Desktop App Contributor Event](https://mattermost.com/blog/mattermost-desktop-app-contributor-event/)

[Mattermost デスクトップアプリ](https://github.com/mattermost/desktop)のコントリビューションイベントが 2022/2/4 まで開催されています。  
[Playwright](https://playwright.dev/)や[Robot.js](http://robotjs.io/)を使ったE2Eテストに関するコントリビュートを増やすことが主な目的のようで、これらのツールを使ったDesktopアプリのE2Eテストについて興味がある場合は、Mattermostのテスト方法を知ったり、Mattermostの開発者にテストの書き方をレビューをしてもらうチャンスかもしれません。

### Mattermost Community Awards
[The 2021 Mattermost Community Awards \- Mattermost](https://mattermost.com/blog/2021-mattermost-community-awards/)

2021年中にMattermostへ貢献をした3,400名のうち、目立った貢献をしたコントリビューターを表彰するMattermost Community Awardsに **Top Overall Mattermost Contributors** として選出していだだきました。

> Yusuke Nemoto
> 
> Yusuke is proof that you can have a huge impact on an open source community without contributing large volumes of code. To start, he’s our unofficial Japanese diplomat and has been at the center of our community growth in Japanese audiences. He dedicates considerable time to maintaining Japanese translations of Mattermost release announcements on his personal blog and is highly active in our localization community for Japanese translations. He’s also active on social media, where he maintains our community-run Japanese Twitter account. All the while, Yusuke contributes code to multiple Mattermost repos regularly.

## おわりに
次の`v6.4`のリリースは 2022/02/16(Web)を予定しています。
