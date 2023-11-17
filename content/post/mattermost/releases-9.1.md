---
title: "Mattermost 9.1の新機能"
emoji: "🎉"
published: true
date: 2023-10-17T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/10/16 に Mattermost のアップデートとなる `v9.1.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v9-1-feature-release)

本バージョンでの主な変更点に関する3分程度の動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube dbHg-63J9dA >}}

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

## (Channels) サイドバーの幅が変更可能に

Mattermostの左サイドバーと右サイドバーの幅を調整することができるようになりました。  
左サイドバーについては、チャンネル名が長すぎて省略表示されてしまっていた場合などに表示文字数を拡大することができそうです。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-resize-sidebar.gif)


## (Channels) グループメッセージを非公開チャンネルに変換可能に

Mattermostでは、複数人が参加するダイレクトメッセージを「グループメッセージ」と呼びますが、この「グループメッセージ」を非公開チャンネルに変換できるようになりました。

この操作を行うには、まず、グループメッセージのチャンネルメニューから **Convert to Private Channel** を選択します。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-gm-convert-menu.png)

ダイアログが開くので、変換先となる非公開チャンネルのチャンネル名を入力し、**Convert to Private Channel**ボタンを押します。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-gm-convert-dialog.png)

すると、チャンネルの投稿内容を引き継いだ形で非公開チャンネルに変換することができます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-gm-convert-after.png)

## (Channels) グループメッセージへの投稿の通知設定が変更に

グループメッセージの通知設定が変更され、明示的にメンションされていなくてもメンションされた時と同様、チャンネル名の横に投稿数を示すバッジが表示され、デスクトップ通知が実施されるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-gm-notification.png)

Mattermost v9.1以降でグループメッセージを表示すると、通知設定の変更に関する以下のようなポップアップが表示されます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-gm-notification-popup.png)


## (Channels) セルフホストMarketplaceが再び利用可能に

Mattermost v9.0にて、Mattermost UI上のマーケットプレースで管理可能なプラグインが[大幅に絞られる](https://blog.kaakaa.dev/post/mattermost/releases-9.0/#boards-plugin%E3%82%92%E5%A7%8B%E3%82%81%E3%81%A8%E3%81%97%E3%81%9F%E5%A4%9A%E3%81%8F%E3%81%AE%E3%83%97%E3%83%A9%E3%82%B0%E3%82%A4%E3%83%B3%E3%81%AE%E3%82%B3%E3%83%9F%E3%83%A5%E3%83%8B%E3%83%86%E3%82%A3%E3%81%B8%E3%82%B5%E3%83%9D%E3%83%BC%E3%83%88%E3%81%AE%E7%A7%BB%E8%A1%8C)こととなり、Mattermost公式によってサポートされていないコミュニティ管理のプラグインについては、今後、プラグインの更新があった場合に手動でのプラグインアップデートを実施しなくてはならなくなりました。

今後も引き続きMattermost UI上のマーケットプレース画面でプラグインの管理を実施するための機能として、セルフホストのMarketplaceと接続する設定が復活しました。  
[mattermost/mattermost\-marketplace](https://github.com/mattermost/mattermost-marketplace)を起動し、そのURLを`PluginSettings.MarketplaceURL`に設定し、`PluginSettings.EnableRemoteMarketplace`を`true`とすることで、セルフホストのマーケットプレースに接続できるようになり、Mattermost公式によってサポートされていないプラグインもMattermost UI上のマーケットプレースから管理できるようになります。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.1/channels-marketplace.png)

## その他の変更

特になし

## その他のトピック

### Hacktoberfest 2023

もう10月も下旬に差し掛かってきましたが、今年もMattermostはHacktoberfestへの参加を奨励しています。

[Hacktoberfest 2023: Celebrate open source, make an impact on the planet](https://mattermost.com/blog/mattermost-hacktoberfest-2023/)

Hacktoberfestへの参加登録を行うと、[Holopin](https://www.holopin.io/)からBadgeが送られてきますが、上記のブログエントリの末尾にMattermostからのHolopin Badgeがもらえるリンクが記載されています。

## おわりに
次の`v9.2`のリリースは 2023/11/16(Thu)を予定しています。  
