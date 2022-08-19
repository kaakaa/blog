---
title: "Mattermost 7.2の新機能"
emoji: "🎉"
published: true
date: 2022-08-19T10:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/08/16 に Mattermost のアップデートとなる `v7.2.0` がリリースされました。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.2 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-2-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-2-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- Platform: 上記機能を含むMattermost全体

## Channels: メッセージ転送 (Message Forwarding)

Mattermost内の投稿を他のチャンネルへ転送するメニューが追加されました。

[Forward messages](https://docs.mattermost.com/channels/forward-messages.html)

メッセージ転送機能を使わなくても、投稿のURLをMattermostへ投稿することで[メッセージを別のチャンネルでプレビュー表示する機能はMattermost v6.0の時点で利用可能](https://blog.kaakaa.dev/post/mattermost/releases-6.0/#%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8%E3%83%AA%E3%83%B3%E3%82%AF%E3%81%AE%E3%83%97%E3%83%AC%E3%83%93%E3%83%A5%E3%83%BC%E8%A1%A8%E7%A4%BA)でした。しかし、このメッセージプレビューによる投稿内容の共有を行う場合、元の投稿のURLを取得・コピーし、転送先のチャンネルへ移動してURLを投稿する必要がありました。本バージョンで転送機能が実装されたことで、この作業が簡略化され、有益なメッセージを簡単に共有することができるようになります。

---

転送機能を利用するには、投稿のメニューから`転送`を選択します。

![channels-forwarding-menu](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/channels-forwarding-menu.png)

表示されるダイアログに転送先チャンネルとコメントを入力し、転送を実行することで、転送先に指定したチャンネルに記入したコメントとともに投稿内容が転送されます。

![channels-forwarding-dialog](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/channels-forwarding-dialog.png)

転送先チャンネルには、Mattermostインスタンス内の自分が参加しているチャンネルならばどのチャンネルでも指定することができます。また、以前にダイレクトメッセージ(DM)をやり取りしたことのあるユーザーに対しては、そのDMチャンネルを転送先として選択することもできます(手元の環境では、DMを送ったことのないユーザーとのDMチャンネルは転送先として指定できないようでした)。また、非公開チャンネルやDM/GM(グループメッセージ)チャンネルの投稿でも転送メニューを選択することはできますが、転送先チャンネルは指定することができず、その投稿の存在するチャンネルにしか転送することができません。

また、コメント入力欄は1行テキスト欄に見えますが、`Shift+Enter`を押すことで改行を行うことができ、Markdown記法も使用することができます。

![channels-forwarding-content](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/channels-forwarding-content.png)

転送が完了すると、ダイアログで選択した転送先チャンネルに以下のような投稿が作成されます。

![channels-forwarding-post](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/channels-forwarding-post.png)

元の投稿へのリンクURLは自動で挿入されます。  

### 注意: 日本語入力時の不具合について

現在、コメントを追加する前に転送先チャンネルを選択してしまうと、コメント入力中の`Enter`キー押下によってメッセージ転送が実行されてしまうバグがあるようです。日本語でコメントを入力する場合は、`Enter`キーを押さずに確定を行う（カーソルを押下することが確定するなど）か、転送先チャンネルを選択する前に入力を完了することで回避できそうです。

[Unable to add Japanese comments correctly in Message Forwarding Dialog · Issue \#20838 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/20838)


## (Enterprise) Platform: Audit Log v2 Beta

システム管理者向けの監査ログの内容がアップデートされるようです。

[Audit log v2 \(experimental\)](https://docs.mattermost.com/comply/audit-log.html)

今回のアップデートでは、監査対象イベントの前後の状態の変化により焦点を当て、以前より多くのイベントのデータが記録できるようスキーマの変更が行われているようです。
以前のバージョンの監査ログを対象にアラート設定などを行なっている場合、今後Audit Log v2に合わせた修正が必要になります。

## その他の変更

### Channels: 未読チャンネルを表示する際のスクロールポジション設定

未読の投稿が存在するチャンネルを開いた時のスクロール位置に関する設定が追加されました。**設定 > 詳細 > 未読チャンネルを表示したときのスクロール位置** から設定を変更することができます。

![channels-scroll-position](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/channels-scroll-position.png)

`Start me where I left off` (以前表示した位置から開始する)を選択すると、以前チャンネルを開いた時に表示していた位置に遡ってチャンネルが表示されます。`Start me at the newest message` (最新のメッセージから開始する)を選択すると、以前開いていた位置とは関係なく、常に最新のメッセージを開始位置としてチャンネルを表示します。

## アップグレード時の注意事項

### スキーマ変更に係る所要時間

今回のバージョンアップでは、格納されるデータをより厳密なものにするためにデータベースのスキーマ変更が実行されます。スキーマ変更にかかる時間はチャンネル数に依存するようですが、10万チャンネル以上あるインスタンスでも10秒超ほどで完了するため、大きな問題にはならないかと思います。

スキーマの変更内容と所要時間に関するテスト結果については以下のリンク先を参照してください。

[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

### Slack料金改訂に伴いMattermostへの言及が増加

2022/07/18に[Slackが料金改訂を行うことを発表しました](https://slack.com/intl/ja-jp/blog/news/pricing-and-plan-updates)。  
中でもフリープランのメッセージ上限が10,000件から90日間に変更されたことの影響が大きいようで、フリープランを利用していたSlackユーザーの方々が代わりのチャット基盤となり得るチャットツールを探し、SNS等でそれらのツールに対する言及するといったことが多く行われていました。  
それに合わせ、ここ１ヶ月ほどでMattermostを紹介する記事がいくつも公開されていたため、見つけた範囲でまとめてみました。

Mattermostnの紹介およびセットアップ手順に関する記事
* [第726回　Slackの代替として、ハドルミーティングも実装されたオープンソース版のMattermostはいかが？ \| gihyo\.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0726)
* [チャットツールMattermostとは？使い方や評判、無料で出来ることを紹介 - TravelWork](https://travelwork.jp/2022/07/mattermost/)
* [Slack フリープランが改悪したので Mattermost をインストールしてみた \- Qiita](https://qiita.com/bezeklik/items/a3882af7f4303bdc3cb6?utm_campaign=post_article&utm_medium=twitter&utm_source=twitter_share)
* [WSL2 Ubuntu で Mattermost Preview を試す](https://zenn.dev/fehde/articles/a8919efbe2efa7)
* [Mattermost を Docker で試験導入する \(Docker Image 版\)](https://zenn.dev/fehde/articles/8d530e1118f8a6)

SlackからMattermostへの移行に関する記事
* [SlackからMattermostへの移行 \- Qiita](https://qiita.com/mitsurupj/items/7f002a696bab66b4259e?utm_campaign=post_article&utm_medium=twitter&utm_source=twitter_share)
* [Slack から Mattermost へ移行](https://blog2.issei.org/2022/07/27/slack-to-mattermost/?utm_source=dlvr.it&utm_medium=twitter)

Mattermost利用者向けサポートサービスに関する記事
* [【インターネットにつながってなくても使える】NASAがSlackの代わりに使うビジネスチャット「Mattermost」を知ってほしい。｜リックソフト：Jira,Slack, WorkatoなどBtoB「仕事を楽にする」ITツール導入支援](https://note.ricksoft.jp/n/n1124f6b5ee87)
* [SlackからMattermostの切り替えサポートについて｜PressWalker](https://presswalker.jp/press/2127)

### Mattermost Pricingページの記載内容に一部誤解を招く表現があることについて

MattermostのPricingページの記述を確認すると、Self-hostedのStarterプラン(無償利用)でも、アクセスできるメッセージ履歴の数に制限があるように読めてしまうようです。

PricingページのStarterプランの説明に、メッセージ履歴は最新の10,000件しかアクセスできないというような表記があります。この制限は、**Cloud版**のStarterプラン利用時のみの制約であるはずですが、PricingページはSelf-hosted版とCloud版について記載を分けていないようなので、この制限がSelf-hosted版にも適用されるように見えてしまいます。

![topic-pricing](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/topic-pricing.png)

また、Pricingページ末尾のFAQにも、「Self-hosted版とCloud版で料金の面で差異は無い」と書かれているため、Pricingページの記載だけを読むと、やはりSelf-hosted版にもデータ制限があるように読めてしまいます。

![topic-pricing-faq](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/topic-pricing-faq.png)

この点について[Mattermost公式チームに問い合わせ](https://community-daily.mattermost.com/core/pl/nnhn63q8c3gpukow1xqe6f5x5o)てみましたが、やはりデータ制限はCloud版のみの話で、Self-hosted版には特にMattermostとしてデータ制限は無いはずとの回答をもらいました。

![topic-pricing-confirmation](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/topic-pricing-confirmation.png)

Pricingページの記載内容についても見直してくれるそうです。

## Roadmap

### Channels: 今後追加される予定の機能について

毎月、[Mattermost公式コミュニティサーバーのRoadmapチャンネル](https://community-daily.mattermost.com/core/channels/roadmap)にて、Mattermostプロダクトの方向性についての情報共有が行われています。その中から、Mattermostのチャット機能である`Channels`に追加される予定の機能を紹介します。

参考にしたのは以下の資料で、誰でもみれる形で公開されています。

[July 2022 \- Channels Roadmap Update \- Viewer Copy \- Google スライド](https://docs.google.com/presentation/d/1gORMxIpa87z0R2vWZuOmJURVB8K2HFjQn0wXPeuzM9Y/edit#slide=id.ge9b7e537ec_0_0)

まず、`Global Drafts`です。

![topic-roadmap-channels-drafts](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/topic-roadmap-channels-drafts.png)

`Global Drafts`では、Mattermost上でメッセージを投稿しようとして途中で入力を中断した場合などに、その書きかけのメッセージを集約してくれる機能です。  
今までのバージョンでも、書きかけのメッセージがある場合、左サイドバーでそのチャンネルのアイコンが鉛筆アイコンに変わることで目視することはできましたが、参加するチャンネルが多い場合は見つけづらいなどの難点がありました。`Global Drafts`によって、急な対応が必要になった場合でも、すぐに書きかけのメッセージに戻ることができそうです。

もう一つは`Message Priority & Acknowledgement`です。

![topic-roadmap-channels-priority](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.2/topic-roadmap-channels-priority.png)

この機能により、Mattermost内の投稿に対する注目度をさまざまな形式で設定できるようになりそうです。

`Message Priority`では、メッセージの重要度を`Standard`、`Important`、`Urgent`などの形式で設定できるようになるようです。  
また、`Request acknowledgement`では、メッセージの既読チェックに使えるボタンが付与されるようです。今までは絵文字リアクションで行なうことが多いかと思いますが、公式の機能としてサポートされるようです。  
最後に、`Message Priority`を`Urgent`に設定した投稿は、`acknowledge`か返信が行われるまで5分ごとに通知が行われるようにすることもできるようです。

これらの機能について、実装されたらまた紹介したいと思います。



## おわりに
次の`v7.3`のリリースは 2022/09/16(Fri)を予定しています。

