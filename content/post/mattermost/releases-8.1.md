---
title: "Mattermost 8.1の新機能"
emoji: "🎉"
published: true
date: 2023-08-26T14:45:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/08/25 に Mattermost のアップデートとなる `v8.1.0` がリリースされました。  
本バージョンは[ESR(Extended Support Release)](https://docs.mattermost.com/upgrade/extended-support-release.html)に設定されており、2024/05/15までセキュリティ対応/バグ修正等のサポートが行われる予定です。(ESRでないバージョンのサポート期間は3ヶ月間です)

現在、いくつかの修正を加えた`v8.1.1`のリリース作業を進めているようなので、近日中にリリースされると思います。8/10以降に加えられた変更が、`v8.1.1`に含まれるものだと思います。  
https://github.com/mattermost/mattermost/commits/release-8.1/

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v8-1-extended-support-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体



## (Channels) チャンネル閲覧モーダルのUI更新

左サイドバーの **チャンネルを追加する > チャンネルを閲覧する** を選択することで開くチャンネル一覧を閲覧する画面の内容が更新されました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.1/channels-browse-channels.png)

参加済みのチャンネルには、チャンネル名の下に `Joined` 表示されるようになり、また、右上の`Hide Joined`のチェックボックスから参加済みのチャンネルを非表示に設定できるようになりました。

## (Platform) パスワード復元リンクのカスタマイズ

Mattermostのログイン画面には、パスワードを忘れた場合にユーザー自身でパスワードの再設定を実施するための`パスワードをお忘れですか?`というリンクが表示されていますが、このリンクをカスタマイズできるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.1/platform-forgot-password.png)

まず、**システムコンソール > サイト設定 > カスタマイズ > Forgot Password Custom Link(パスワード再設定カスタムリンク)** から、`パスワードをお忘れですか?`のリンクを自由に設定することができます。これは、パスワード再設定を組織内の定められたプロセスで行いたい場合に、そのプロセスの案内ページへのリンクを設定するなどのユースケースで利用できます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.1/platform-forgot-custom-link.png)

また、**システムコンソール > 認証 > パスワード > Enable Forgot Password Link(パスワード再設定リンクを有効にする)** から、`パスワードをお忘れですか?`のリンクを非表示にすることもできます。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.1/platform-forgot-password-disabled.png)

ただ、パスワード再設定リンクを非表示にした場合、ユーザーがパスワードを忘れると復元する手段がなくなってしまうため、**システムコンソール > サイト設定 > カスタマイズ > 独自ブランド機能を有効にする**を有効にし、**独自ブランドテキスト**に管理者の連絡先を記載するなど、復旧手段も合わせて検討する必要があるかと思います。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.1/platform-forgot-password-brandtext.png)

## その他の変更

### (Channels) [解決] チャンネル名に日本語を指定する場合の注意点

Mattermost v8.0リリース時のエントリにも書きましたが、新しくチャンネルを作成する際、チャンネル名に日本語を指定すると意図しないタイミングでチャンネル作成が実施されることがあり、その件についてIssueを立てていました。  
* [Unexpectedly fire the handleEnterKeyPress in NewChannelModal with Japanese IME input mode · Issue \#23967 · mattermost/mattermost](https://github.com/mattermost/mattermost/issues/23967)  
* [(Channels) チャンネル名に日本語を指定する場合の注意点 - Mattermost 8\.0の新機能](https://blog.kaakaa.dev/post/mattermost/releases-8.0/#channels-%E3%83%81%E3%83%A3%E3%83%B3%E3%83%8D%E3%83%AB%E5%90%8D%E3%81%AB%E6%97%A5%E6%9C%AC%E8%AA%9E%E3%82%92%E6%8C%87%E5%AE%9A%E3%81%99%E3%82%8B%E5%A0%B4%E5%90%88%E3%81%AE%E6%B3%A8%E6%84%8F%E7%82%B9)  

Mattermost v8.1でも再現するかいくつかの環境で試してみたところ、再現しませんでした。Mattermostに対する変更を見てみたところ、ちょうどv8.1リリースに向けてKorean IMEのEnterキーの問題を解決するPRが取り込まれており、おそらくこの影響でチャンネル作成時の日本語の問題も解決されていたのではないかと思います（詳細までは理解しきれていないですが）。  
[MM\-51676 \- korean chars create duplicate categories by pvev · Pull Request \#23839 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/23839)

もし、v8.1以降でも再現する方がいれば、上記のIssueに追加でコメントするか、新規にIssueを立てていただければと思います。

## その他のトピック

### Mattermost v9.0

来月(2023/09/15)にリリース予定のMattermostの次バージョンは、メジャーアップデートとなるMattermost v9.0になる予定だそうです。  
これは、大規模な新機能があるというより、いくつかの機能を削減する必要に迫られてという感じのようです。  

大きな話としては、Mattermost Board(Focalboard)の公式サポートが終了するようです。また、その他にも多くのプラグインが公式のサポートを外れ、限られたプラグインのみをサポートしていく方針だと以下のForumエントリで語られています。  
[Upcoming product changes to Boards and various plugins \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/upcoming-product-changes-to-boards-and-various-plugins/16669)

その他、Insight機能やGfycatと連携する機能が削減予定のようです。Mattermost v9.0で削減される機能については以下を参照ください。  
[Removed and deprecated features for Mattermost](https://docs.mattermost.com/install/deprecated-features.html#removed-features-in-upcoming-versions)

### Mattermost AI plugin

先月のMattermost v8.0のリリースエントリでも触れましたが、最近Mattermost AIプラグインの活動が活発なようです。  
[mattermost/mattermost\-plugin\-ai: Mattermost plugin for LLMs](https://github.com/mattermost/mattermost-plugin-ai)

Mattermost公式YouTubeチャンネルの方でも、Mattermost AI関連の動画がよくアップロードされています。  
[Mattermost \- YouTube](https://www.youtube.com/@MattermostHQ/videos)

自分も、少し触ってみた感触に関する記事を書きました。  
[Mattermost AI Pluginを試してみた](https://zenn.dev/kaakaa/articles/mattermost-plugin-ai)

まだ安定版という感じではないですが、Mattermost上の様々な場面でAIを使えるようにしていこうとしているようなので、AIのユースケースの参考として少し触ってみるのも面白いかと思います。

## おわりに
次の`v9.0`のリリースは 2023/09/15(Fri)を予定しています。  
