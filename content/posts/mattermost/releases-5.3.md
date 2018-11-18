
---
title: "Mattermost5.3の新機能"
date: 2018-09-20T22:06:57+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/09/14にMattermost 5.3.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost Changelog — Mattermost 5\.3 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-3)
* [Mattermost 5\.3: Enhanced search on desktop and mobile, plugin hackathon highlights and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-3-enhanced-search-on-desktop-and-mobile-plugin-hackathon-highlights-and-more/)

---

v5.3.0のリリース作業でミスがあり翻訳が一部壊れています。英語以外の翻訳版を利用されている場合はv5.3.1を利用してください。
[Mattermost 5\.3\.1 released \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-3-1-released/)

# アップデート内容

## 日付指定検索

Mattermostメッセージ検索機能で投稿日付を指定できるようになりました。

* `before:YYYY-MM-DD`: 指定日付より前の投稿を検索
* `on:YYYY-MM-DD`: 指定日付に行われた投稿を検索
* `after:YYYY-MM-DD`: 指定日付より後の投稿を検索

また、右上の検索欄に`after:`、`on:`、`before:`を入力することで、日付選択モーダルが現れます。

<img width="627" alt="スクリーンショット 2018-09-19 0.29.52.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/1eb95c86-7eea-7103-8a61-99e7f447a3e1.png">


## SAMLユーザーへのID設定

ID ProviderとしてSAMLを設定している場合、過去に無効化されたユーザーと同じメールアドレスを持つユーザーを新たに追加すると、無効化されたユーザーの履歴にアクセスできてしまうという問題がありました。
この問題を解決するために、メールアドレスの代わりとなるIDを指定することができるようになりました。
IDを指定することができるようになったことにより、結婚による姓の変更などに起因するメールアドレスの変更にも対応できるようになりました。

## プラグインハッカソンの成果

2018/8/16,17に行われたMattermostプラグインハッカソンで開発された~~プラグイン~~機能の紹介です。

* [**RSS feed**](https://github.com/jespino/matterfeed): ~~RSSのフィードを行えるプラグインのようです~~ RSSのフィードを購読する連携機能のようです。この機能はプラグインではないですが、プラグインハッカソン中に開発された成果として紹介されていたようです
* [**Webcome bot**](https://github.com/mattermost/mattermost-plugin-welcomebot/tree/master): 新しくチームに加入したユーザーにウェルカムメッセージを投稿できます
* [**Matterpoll**](https://github.com/matterpoll/matterpoll): Mattermost上で投票を行える機能です
* [**/remind**](https://github.com/scottleedavis/mattermost-plugin-remind): 指定した時間が経過した後にリマインドをしてくれるプラグインです。
* **Crosspost**: つのチャンネルを複数のMattermostサーバーとリンクすることができ流ようです
* [**Auto-translator**](https://github.com/mattermost/mattermost-plugin-autotranslate): Amazon Translateを使って様々な言語を英語に自動変換するプラグインです

その他にもリッチなエディタ機能を追加するプラグインや、[表示時間をユーザーのローカルタイムゾーンに合わせてくれるプラグイン](https://github.com/mattermost/mattermost-plugin-walltime)などの開発が行われています。
また、Mattermost CTOによるハッカソンの[振り返り](https://mattermost.com/blog/plugin-hackathon/)やプラグインの[ビデオチュートリアル](https://www.youtube.com/watch?v=Cx2EBkGkz00)などが公開されています。

## その他

### リックソフトがMattermostとパートナー契約を締結

[ニュースリリース：20180920 リックソフト Slackライクなオンプレミス型チャットツール Mattermostのパートナー契約を締結](https://www.ricksoft.jp/news/n20180920.html)

アトラシアン社のパートナー企業として有名なリックソフト社がHipChat Server販売終了を受け、Mattermostとのパートナー契約を締結したそうです。
私の知る限りでは、国内初のパートナー企業となるかと思います。（他にあったらゴメンナサイ）
[Mattermost Partner Programs — Mattermost 5\.3 documentation](https://docs.mattermost.com/process/partner-programs.html)

### Matterpoll
現在、Mattermostプラグイン機能を活用した投票機能である`Matterpoll`というOSSを開発しています。
まだ安定バージョンではないため今後も破壊的変更が入る可能性は高いですが、もしご興味がある方がいらっしゃいましたら導入・Contributeをお願いします。
[matterpoll/matterpoll: Poll command for Mattermost](https://github.com/matterpoll/matterpoll)

<img width="848" alt="スクリーンショット 2018-09-20 22.04.19.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/dd7dd94a-78cf-c6cb-6b2d-4a1bce1e13aa.png">


# おわりに

次回の`v5.4`のリリースは2018/10/16(Tue)を予定しています。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。


