---
title: "Mattermost 9.8の新機能"
emoji: "🎉"
published: true
date: 2024-05-17T22:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/05/16 に Mattermost のアップデートとなる `v9.8.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog \- Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-8-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube o-OWdIkBYLk >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。


## 通知設定のDesktop/Mobile共通化

> ・Combined Desktop and Mobile notifications in the user settings modal.

今までのバージョンでは通知の対象となるイベントに関する設定は、「Web/デスクトップアプリ」と「モバイルアプリ」で別々に管理されていましたが、v9.8からユーザー設定画面から一括で管理できるようになりました。  
**Send notifications for** の設定を変更することで、Web/デスクトップアプリ/モバイルアプリに対して共通の設定を適用することができます。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.8/channels-notification-settings.png)

モバイルアプリでは異なる設定を使用したい場合、**Use different settings for my mobile devices**のチェックを入れると、モバイルアプリ専用の設定画面が現れます。

**Send notifications for**の設定については、モバイルアプリからも設定を変更することができます。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.8/channels-notification-mobile.png)


## チャンネル初期画面の変更

> ・Enhanced the user interface for channel introductions.

チャンネルを作成したときに最初に表示される画面が刷新され、各種情報が分かりやすく表示されるようになりました。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.8/channels-channel-introduction.png)

## 投稿メッセージ内絵文字のツールチップ表示

> ・Added emoji tooltips on hover in post message.

投稿内で使用された絵文字をツールチップとして表示できるようになりました。  
文字を詰め込んだ絵文字などは投稿内で使用すると表示が小さすぎて読めないことがありますが、そのような場合でも絵文字の内容が確認できるようになります。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.8/channels-emoji-tooltip.png)

## 右サイドバーのスレッド表示改善

> ・Updated the right-hand side Thread view to use relative timestamps to be more consistent with the global Threads view.  
> ・Added a total reply count to the right-hand side thread view.

メッセージスレッドを確認する際に開く右サイドバーに、スレッド内のメッセージ数が表示される用になりました。  
また、スレッド内の各メッセージの投稿時間が相対時間で表示されるようになりました。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.8/channels-message-thread.png)

## その他の変更


### チャンネルブックマーク

> ・Added Channel Bookmarks (disabled by default).

チャンネルヘッダーの下にリンクやファイル添付ができる**チャンネルブックマーク**機能が一部使えるようになったようです(?)。  
(デフォルトで無効化されており、チャンネルブックマーク用の権限があるとChangelogには説明がありましたが、v9.8のシステムコンソールを見ても当該設定は見当たりませんでした)

今回のリリースではサーバー側の機能のみ実装された状態のようで、UIからこの機能が利用できるようになるのは、以下のPRがマージされてからのようです。  
[Channel Bookmarks UI by calebroseland · Pull Request \#25889 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/25889)

PRの説明文の中にチャンネルブックマーク機能の動作イメージの動画も添付されています。  

### ERROR_SAFETY_LIMITS_EXCEEDED

> Added safety limit error message in compiled Team Edition and Enterprise Edition deployments when enterprise scale and access control automation features are unavailable, and message count exceeds 5 million posts. ERROR_SAFE_LIMITS_EXCEEDED.

Free版のMattermostでは、登録ユーザーは1万ユーザーまでというSafety Limitがあり、それを超えると警告とともにいくつか機能が制限される可能性があるそうです。  
[Mattermost error codes \- Mattermost documentation](https://docs.mattermost.com/manage/error-codes.html)

このSafety Limitに「500万投稿まで」という制限が新たに追加されたようです。  
ただ、この新たに追加された制限を削除するPRが投稿されているようで、もしかしたら「500万投稿まで」という制限は削除されるのかもしれません。（背景が分かっていないので断定的なことは言えませんが）  
[Removed message post limits warning by harshilsharma63 · Pull Request \#27036 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/27036)


## その他のトピック

今月は、Mattermost公式ブログの方で技術系の記事がいくつか公開されていたため、その内の2つを紹介します。

### cgoに対するMattermostの貢献

MattermostのSecurity Researchにより発見されたcgoの任意コード実行の脆弱性 (CVE-2024-24787)に関する解説記事が公開されました。  
[Go fixes its 7th code execution bug in the same feature \- Mattermost](https://mattermost.com/blog/go-fixes-its-7th-code-execution-bug-in-the-same-feature/)

今までcgoで発見された7つの任意コード実行の内、4つがMattermostによって発見されているそうです。

### Postgresクエリを1000倍高速化した話

大規模なMattermostのデータベース (投稿メッセージ1億件) に対するElasticsearchのIndexingに非常に時間がかかるという事象から、対象のクエリを発見し、そのクエリを高速化させるための分析や対策等について述べられています。

[Making a Postgres query 1,000 times faster \- Mattermost](https://mattermost.com/blog/making-a-postgres-query-1000-times-faster/)

## おわりに
次の`v9.9`のリリースは 2024/6/14(Fri)を予定しています。  
