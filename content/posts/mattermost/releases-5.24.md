---
title: "Mattermost5.24の新機能"
date: 2020-06-20T12:56:12+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/6/17にMattermost v5.24.0がリリースされました。  
(2020/06/20現在、[バグFixを含む v5.24.1 のリリースが計画されています](https://docs.mattermost.com/administration/changelog.html#release-v5-24-feature-release))。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost Release v5\.24 is now available \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-release-v5-24/)
* https://docs.mattermost.com/administration/changelog.html#release-v5-24-feature-release

---

## Mattermostアプリ使用時にセッション維持期間を自動で延長する

Mattermostにはセッション維持期間の設定項目があり、ログイン成功時からそのセッション維持期間で指定した期間が経過すると自動でログアウトされ、ユーザーは再度ログイン処理を行わなければなりませんでした。セキュリティ向上のためにセッション維持期間は短くしたいが、短くしてしまうとユーザーに再認証を要求するまでの期間も短くなってしまい、その結果、セッション切れから再認証までの間の通知を見落としてしまうというジレンマがありました。

それを解決するため、ユーザーがMattermost上でチャンネル切替やタイピングなどの行動を行った場合に、その時点から再度セッションが開始される設定が追加されました。

この設定は、**システムコンソール > セッション維持期間 > 活動中はセッション維持期間を延長する** から行うことができます。

![extend session](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/extending_session.png)

詳しくは、下記の公式ブログで説明されています。
[Automatically extend Mattermost sessions for less frequent logouts and happier end users \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/session-expiry-experience/)


## 検索フィルタの改善

Mattermostヘッダ部分にある検索画面で、検索フィルタが利用しやすくなりました。
今までのバージョンでは空の検索画面をクリックすると検索フィルタについての説明が表示されるだけでしたが、今回のバージョンから検索フィルタの入力方法が改善され、検索フィルタの↑、↓、Enterキーでの選択入力や、日付指定フィルタ入力時のカレンダー表示などが行えるようになりました。

![search filter](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/search_filter.gif)

## 絵文字リアクション操作が簡単に (モバイル)

モバイルアプリ v1.32から、投稿を長押しすると出てくるメニューによく使われるリアクションが表示されるようになり、今までより簡単に絵文字リアクションが付けられるようになりました。

<img src="https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/emoji_mobile.jpg" height=480/>

## 自動補完を使用したスラッシュコマンド入力 (BETA)

今までのバージョンでも、メッセージ入力欄で `/` を入力すると利用可能なスラッシュコマンドの自動補完を行うことができましたが、今回のバージョンから、さらに選択したスラッシュコマンドで利用可能なオプション部分についても自動補完が可能となりました。

<img src="https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/slash_autocomplete.gif" height=480/>
(画像は[公式リリースブログ](https://mattermost.com/blog/mattermost-release-v5-24/)より転載)

## Bleveを利用した全文検索で日本語検索も可能に！（実験的機能）

今まで、Mattermost内の日本語メッセージを検索するためには、有償版のElasticsearch機能を使うか、DBに対して設定を追加する必要がありました。今回のバージョンから、Mattermost内のメッセージの全文検索に[Bleve](http://blevesearch.com/)を利用できるようになり、Bleveを使用することで、無償版でもDBへの設定なしに日本語を含む文字列を検索できるようになります。

Bleveの利用設定は **システムコンソール > 実験的機能 > Bleve** から行うことができます。
**Bleveのインデックス付与を有効にする** を有効にし、**IndexDir** にインデックスファイルを配置するディレクトリを指定する。（ここで指定したディレクトリの親ディレクトリが無いと、インデックス作成時にエラーが発生します）

![bleve_settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/bleve_settings.png)

上記の設定を保存したら、**一括インデックス付与 > インデックス付与をいますぐ開始する**をクリックすることで、Bleveによるインデックス付与を開始することができます。メッセージ数が多い場合、この処理には時間がかかる可能性があります。インデックス付与はバックグラウンド処理として行われるため、処理中もMattermostは利用可能ですが、検索結果に不整合が発生する可能性があります。

![bleve_indexing](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/bleve_indexing.png)

インデックス付与が完了したら、**Bleveの検索クエリを有効にする**、**Bleveの自動補完クエリを有効にする**を有効に設定することで、Bleveを利用した検索が行えるようになります。設定完了後は、新しいメッセージも自動でインデックス付与が行われるため、インデックス最新かのための操作は必要ありません。

設定完了後、Mattermost上で検索を行うことで、日本語一文字だけでも検索できるようになります。

<img src="https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/bleve_search.png" height=480/>

Bleveによる全文検索機能は実験的な機能のため、予期せぬ事象が発生する可能性があります。また、私も数件のメッセージがある環境でしか試していないため、大量のメッセージが存在する環境でどのような事象が発生するかは確認できていません。

## Enterprise版の改善

### Office365カレンダーの同期 (E20)

Mattermost上でOffice365カレンダーの予定を管理できるようになりました。

* その日のカレンダーイベントの確認
* Office365カレンダーのstatusとMattermostのステータスを同期 (会議中は"離席中"になる、等)
* カレンダーイベントの承認/拒否がMattermrost上で可能に

![office365_calendar](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/office365_calendar.png)
(画像は[公式リリースブログ](https://mattermost.com/blog/mattermost-release-v5-24/)より転載)

### AD/LDAP グループメンション機能 (BETA) (E20)

AD/LDAPグループに対して @メンション が行えるようになり、これにより複数の人に対して簡単にメンションを飛ばすことができるようになりました。

![group_mention](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/group_mention.png)
(画像は[公式リリースブログ](https://mattermost.com/blog/mattermost-release-v5-24/)より転載)

モバイルアプリでは、自動補完ドロップダウンでLDAPグループメンションのセクションは見えないようです。

https://docs.mattermost.com/administration/changelog.html#breaking-changes
> Breaking Changes
> * On mobile apps, users will not be able to see LDAP group mentions (E20 feature) in the autocomplete dropdown. Users will still receive notifications if they are part of an LDAP group. However, the group mention keyword will not be highlighted.

### チーム/チャンネルのメンバーシップ管理がシステムコンソールから可能に (E20)

今までチームやチャンネルのメンバーシップ管理は、そのチーム/チャンネルの画面からしか行うことが行うことができませんでしたが、今回のバージョンからシステムコンソール上で行うことができるようになりました。

![team_setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/team_setting.png)
(画像は[公式リリースブログ](https://mattermost.com/blog/mattermost-release-v5-24/)より転載)

### プロフィール画像をAD/LDAPと同期 (E10/E20)

Mattermost上でのプロフィール画像をAD/LDAPのプロフィール画像と同期することができるようになりました。

## その他の変更

* チャンネル内のピン留めされた投稿の数がヘッダ部分に表示されるようになりました
![count_pin](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.24/count_pin.png)
* その他、多くの機能追加がありますが、詳細は[Changelog](https://docs.mattermost.com/administration/changelog.html#release-v5-24-feature-release)からご確認ください

# おわりに

次の`v5.25`のリリースは2020/7/15(Thu)を予定しています。
そして機能追加が行われる`v5.26`は恐らく2020/8/14(Fri)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
