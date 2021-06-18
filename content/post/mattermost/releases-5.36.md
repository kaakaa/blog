---
title: "Mattermost 5.36の新機能"
emoji: "🎉"
published: true
date: 2021-06-18T07:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/06/16 に Mattermost v5.36.0 がリリースされました。 

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。　　
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.36 is now available \- Upgrade today to try the new features](https://mattermost.com/blog/mattermost-release-v5-36/)
- [Mattermost Changelog](https://docs.mattermost.com/administration/changelog.html#release-v5-36-feature-release)

## MatterCon 2021

[MatterCon 2021 \- Learn about the latest and greatest from Mattermost\!](https://mattermost.com/events/mattercon-2021/)

MattermostのVirtualコミュニティイベントであるMatterConが来週6/22-24の間で開催されます。
誰でも無料で参加できるイベントのようで、上記のサイトから参加登録できます。

MatterConは、今までもコミュニティメンバ限定で年に一回オフラインで開催されていたイベントでしたが、コロナ禍ということもありVirtual開催でパブリックなイベントになったようです。([過去のMatterCons](https://handbook.mattermost.com/contributors/mattercon#past-mattercons))

Mattermostはフルリモートの会社のため、社員やコミュニティメンバーが同期的なコミュニケーションを行う場というのがMatterConの主な目的ですが、その他にもフルリモートの会社ならではの知見の紹介や、新機能のデモ、ロードマップ、Mattermost内のOpen Source Fridayプロジェクトについての紹介などもあるようです。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## Focalboard Plugin

以前から開発が進められているNotion, Trello, AsanaクローンのKanbanツールである[Focalboard](https://www.focalboard.com/)が、Mattermost Pluginとして利用できるようになりました。

Focalboard Mattermost Pluginについては、以下の記事に使ってみた感想を書いています。
[FocalboardのMattermostプラグインについて](https://zenn.dev/kaakaa/articles/mattermost-focalboard-plugin)

上記の記事では、プラグインをインストールする方法として`.tar.gz`ファイルをアップロードする方法を紹介していましたが、Marketplaceから数クリックでインストールする手順が公式のYouTubeチャンネルで1分程度の動画として紹介されています。

https://youtu.be/V_pLVc4K-1g

MattermostへのプロキシとしてnginxやApacheを使用している場合は、追加の設定が必要になります。

[How to configure the NGINX or Apache web proxy to enable the Focalboard plugin · Discussion \#566 · mattermost/focalboard](https://github.com/mattermost/focalboard/discussions/566)

## (E20/Cloud) Incident Collaborationの改善

今月も[Mattermost Incident Collaboration Plugin](https://github.com/mattermost/mattermost-plugin-incident-collaboration)の改善があります。

### チーム全体に対するPlaybookアクセス権設定

作成されたPlaybookに対するアクセス権を設定する際に、メンバー単位のアクセス権設定だけでなく、チーム全体に対するアクセス権を設定できるようになりました。

![incident-edit-team](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/incident-edit-team.png)

### Playbook作成権限の管理

Playbookを作成できる権限をメンバー単位で設定できるようになりました。

![incident-create-permission](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/incident-create-permission.png)

### Welcomeメッセージ

PlaybookにWelcomeメッセージを設定できるようになりました。
Welcomeメッセージは、Playbookを元に作成されたチャンネルに始めた参加した時に表示されるメッセージです。インシデント対応時に意識しておいて欲しいことや、見ておいてほしいサイトへのリンクなどを指定しておくことで、無駄なコミュニケーションを削減することが期待できます。

![incident-welcome-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/incident-welcome-setting.png)
![incident-welcome-message](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/incident-welcome-message.png)


## 破壊的変更

### 高可用モードの通信がゴシッププロトコルに

商用版を利用しているユーザーで、[可用性を高めるためのクラスタリング機能](https://docs.mattermost.com/deployment/cluster.html#high-availability-cluster-e20)を利用している場合に、クラスタ間の通信が[ゴシッププロトコル](https://ja.wikipedia.org/wiki/%E3%82%B4%E3%82%B7%E3%83%83%E3%83%97%E3%83%97%E3%83%AD%E3%83%88%E3%82%B3%E3%83%AB)で行われるようになりました。今まではゴシッププロトコルの使用はオプションでしたが、そのオプションが廃止され、ゴシッププロトコルのみが用いられるようになります。クラスタ構成では、全てのノードが同じプロトコルを利用して通信している必要があるため、現在ゴシッププロトコルを利用していない場合は、ゴシッププロトコルを使用したアップグレードを行うか、一度全てのノードをシャットダウンしてからアプグレードする必要があるそうです。

https://docs.mattermost.com/administration/changelog.html#important-upgrade-notes

## その他のトピック

### Collapsed Reply Thread

要望の多い返信の折り畳み機能(Collapsed Reply Thread)が、来月リリースされるバージョンからベータ版の機能として利用可能になる予定です。また、クラウド版では今月のリリースから利用可能になる予定です。

![collapsed-thread-announce](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/collapsed-thread-announce.png)

Mattermostコミュニティ用のチャットではすでに利用可能になっています。**アカウント設定 > 表示 > 返信スレッドの折りたたみ** から有効にすることができます。

https://community.mattermost.com

![collapsed-thread-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.36/collapsed-thread-setting.png)


### Mattermost 6.0?

MattermostのJIRAを見ると、`v6.0(Sep 15)`というバージョンが見えます。9月にメジャーバージョンアップが行われるのかもしれません。  

https://mattermost.atlassian.net/browse/MM-36539?jql=project%20%3D%20%22MM%22%20AND%20fixVersion%20%3D%20%22v6.0%20(Sep%2015)%22

## おわりに

次の`v5.37`のリリースは 2021/07/16(Fri)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
