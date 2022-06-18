---
title: "Mattermost 7.0の新機能"
emoji: "🎉"
published: true
date: 2022-06-18T12:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/06/15 に Mattermost のメジャーバージョンアップデートとなる `v7.0.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.0 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-0-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-0-major-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))

## Channels: 返信スレッドの折りたたみ機能がGA(General Availability)へ昇格

[Feature Request Forumで最も多くの要望を集めていた](https://mattermost.uservoice.com/forums/306457-general/suggestions/19572469-make-threads-collapsible)返信スレッドの折りたたみ機能が正式にリリースされました。

返信スレッドの折りたたみ機能の概要は、以下の動画で確認することができます。  
[Tutorial: Using Collapsed Reply Threads on Mattermost \- YouTube](https://www.youtube.com/watch?v=_jR2KWGtgEU)

正式リリースとなったことにより、全てのユーザーが**設定 > 表示 > 返信スレッドの折りたたみ**から本機能のON/OFFを選択できるようになりました。(以前バージョンではベータ版としてリリースだったため、返信スレッドの折りたたみ設定を表示するには**システムコンソール > 実験的機能 > 機能 > 返信スレッドの折りたたみ**から設定を有効にする必要がありました)

![channels-crt-ga-us](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-crt-ga-us.png)

(翻訳の更新があたっていないのか、英語以外の多くの言語で "(ベータ版)" の表記が残ってしまっていますが、正式版です。この表記は次のリリースで修正されていると思います。)

もし、返信スレッドの折りたたみ機能をMattermostインスタンス全体として無効にしたい場合などは、**システムコンソール > サイト設定 > 投稿 > 返信スレッドの折りたたみ**から設定することができます。

![channels-crt-ga](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-crt-ga.png)

## Channels: Mobile v2.0 (Beta)

[Feature Request Forumで二番目に多くの要望を集めていた](https://mattermost.uservoice.com/forums/306457-general/suggestions/10975938-ios-and-android-apps-should-allow-multiple-server)モバイルアプリで複数サーバーへ接続する機能を含む、Mattermostモバイルアプリのメジャーアップデートとなるv2.0のリリースが今年の後半に予定されています。それに先立ち、v2.0のベータ版を試用してくれるユーザーを募集しているようです。

ベータ版の利用方法については、公式のリリースブログを参照してください。  
[Channels: Mobile v2.0 (Beta)](https://mattermost.com/blog/mattermost-v7-0-is-now-available/#mobile)

複数サーバーへ接続する機能は、以下のような操作性になるようです。  
Mattermost画面左上にあるサーバーのようなアイコンをクリックすると、接続するサーバーを選択するメニューが表示されます。


<img alt="channels-mobile-v2" src="https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-mobile-v2.jpg" style="display: block; margin: auto; width: 400px;">

## Channels: 音声通話と画面共有 (Beta)

Slackのハドルミーティングのような音声通話機能がMattermostにも追加されました。  
チャンネルごとに音声通話を開始することができ、通話中でもMattermostの機能を通常通り使用することができます。

以下に自環境でブラウザを2つ横に並べて音声通話機能を動作させた時の様子を載せています。

![channels-call](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-call.gif)

音声通話/画面共有機能はDesktopアプリ、Mobileアプリ、およびブラウザ上で利用することができます。  
Mobileアプリでも共有された画面を綺麗に見ることができました。共有画面のサイズが大きい場合は厳しそうですが。また、前述のMobileアプリv2.0では利用できませんでした。

<img alt="channels-call-mobile" src="https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-call-mobile.jpg" style="display: block; margin: auto; width: 400px;">



本機能はWebRTCをベースに構築されており、Mattermost PluginプロジェクトとしてGitHubで開発が進められています。 (`mattermost-plugins-calls`プラグインは、コンパイル済みのプラグインを利用する場合はMITライセンスの条件の下で利用することが可能ですが、ソースコードを改変して利用する場合は特殊な条件があるようなので注意が必要です。)

https://github.com/mattermost/mattermost-plugin-calls

[Mattermost公式のクラウドサービス](https://customers.mattermost.com/cloud/connect-workspace)を利用している場合、音声通話機能の設定が完了した状態で提供されるため、すぐに使い始めることができますが、セルフホスト版を利用している場合は、以下のドキュメントを参考にプラグインの設定を行う必要があります。

[Start a call \(beta\)](https://docs.mattermost.com/channels/make-calls.html)

## Channels: メッセージ書式設定ツールバー

Mattermostのメッセージ入力欄にWYSIWYG風なボタンが追加され、ボタンによるメッセージの書式設定を行えるようになりました。

![channels-message-format](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/channels-message-format.png)

書式設定のボタンは追加されていますが、メッセージ入力自体は今まで通りMarkdown形式で入力することができ、追加されたボタンはMarkdown形式の入力を補助する機能という位置付けのようです。  
動作している様子については公式ブログを参照ください。

[Channels: Message Formatting Toolbar](https://mattermost.com/blog/mattermost-v7-0-is-now-available/#toolbar)

## Playbooks: インラインPlaybookエディタ

Playbookの編集画面が新しくなり、各セクションごとに分割して表示されているのでは無く、1枚のレポートのように各セクションがシームレスに結合されて表示されるようになりました。

![playbooks-inline-editor](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/playbooks-inline-editor.png)

## Playbooks: Playbooksの利用統計

Playbooksの利用統計情報がシステムコンソールから確認できるようになり、インスタンス内でPlaybooksがどの程度利用されているかを把握しやすくなりました。。**システムコンソール > レポート > サイトの使用統計**から確認することができます。

![playbooks-stats](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/playbooks-stats.png)

## Playbooks: 実行中のPlaybookに対するアクションとトリガーの追加

実行中のPlaybookのステータスが更新された際に、その更新内容の配信先となるチャンネルと外向きのウェブフックを実行ごとに指定できるようになりました。今までもPlaybook単位で配信先を指定することはできましたが、本バージョンからPlaybookの実行ごとに配信先を追加で設定することができます。

![playbooks-trigger-action](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/playbooks-trigger-action.gif)

## Integrations: Apps Bar (ベータ版)

Mattermost統合機能のアイコンを表示する領域として、新たに画面右端にApps Barという領域が追加されました。

![integrations-appsbar](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/integrations-appsbar.png)

今までのバージョンでも、MattermostにはPluginやApps等の統合機能によってチャンネルヘッダ部分に統合機能専用のアイコンを表示することができました。ただ、追加される統合機能が多くなると全てのチャンネルのヘッダ部分に統合機能のアイコンが数多く並ぶようになり、チャンネルヘッダのテキストなど、現在表示しているチャンネルに特化した情報が埋もれるようになっていました。Mattermost v6.0で刷新された新しいUIでは、チャンネルヘッダにはそのチャンネルに特化した情報を表示すべきという明確な目的が与えられたため、v7で統合機能のアイコンを表示するための専用の場所(Apps Bar)が作成されたようです。

Apps Bar機能は**システムコンソール > 実験的機能 > 機能 > Enable App Bar**から有効/無効を設定できます。

## Platform: デスクトップアプリがmacOS App Storeからインストール可能に

MattermostデスクトップアプリがmacOS App Storeから取得できるようになりました。　 

[Mattermost Desktop on the Mac App Store](https://apps.apple.com/app/mattermost-desktop/id1614666244)

App Storeからインストールした場合、新しいバージョンがリリースされると自動でデスクトップアプリがアップデートされるようになります。

## Platform デスクトップアプリのDEB/RPMパッケージリリース

Linux向けにもMattemrostデスクトップアプリのDEB/RPMパッケージがリリースされました。  
インストール方法については公式ブログを参照ください。

[Platform: Linux DEB and RPM packages for Desktop App](https://mattermost.com/blog/mattermost-v7-0-is-now-available/#packages)

## その他の変更

### ログイン画面の更新

ログイン画面のデザインが更新されました。

![platform-login](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/platform-login.png)


## アップグレード時の注意事項

### ログインセッション維持期間に関する設定変更への対応

Mattermostのログインセッションの維持期間に関する設定が、日(day)単位から時間(hour)単位に変更されました。`config.json` を使ったファイル形式で設定を管理している場合、セッション維持期間に関する設定値は自動的に新しい形式へ変換されますが、環境変数を通じて設定している場合は設定する環境変数を変更する必要があります。環境変数による設定の変更方法について詳しくは以下の公式ドキュメントを参照してください。

[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

### MySQLを使用している場合、アップグレード時のマイグレーションに時間がかかる可能性有

セルフホストのMySQLを使ってMattermostを構築している場合、`FileInfo`テーブルに格納されているデータの数によっては、アップグレード時のマイグレーション処理に時間がかかる可能性があります。  
参考情報として、CPU: Intel i7 6cores@2.6GHz、Memory: 16GBの環境で、`FileInfo`テーブルに70万行のデータが格納されている場合、マイグレーションに19秒かかったとのことです。PostgreSQLの場合、マイグレーションの処理時間は無視できるほどです。

### プラグインにより追加されたチャンネルヘッダーアイコンがAppBar領域へ

v7より追加されたApp Bar機能を**システムコンソール > 実験的機能 > 機能 > Enable App Bar**から有効にした場合、プラグインによってチャンネルヘッダーに追加された全てのアイコンが、AppBar領域に移動されます。セルフホスト版では、AppBar機能はデフォルトで無効になっています。

### `TrustedProxyIPHeader`の設定値の確認

以前のバージョンでは、特定の状況下で`ServiceSettings.TrustedProxyIPHeader`のデフォルト設定値が空配列にならないバグが存在しましたが、本バージョンからデフォルト値が空配列になるよう修正されました。以前のバージョンからv7へアップデートする場合、この設定値を確認し、意図しない値が設定されている場合は `nil` を設定する必要があります。

当該設定項目に関する説明は[Configuration settings](https://docs.mattermost.com/configure/configuration-settings.html#trusted-proxy-ip-header)を参照してください。


## その他のトピック

### チャンネル名コンテスト

MattermostのTwitterで、実際に運用されているユニークなMattermostチャンネル名を表彰するコンテストが開催されています。募集期限は6/22 PM12:00(ET) まで募集しており、最もユニークなチャンネル名を投稿した5名の勝者にはグッズが送られるようです。

応募方法など詳しくは、以下のツイートツリーを参照ください。

https://twitter.com/Mattermost/status/1537080481384124419

### Happy Pride Month
6月は、アメリカやカナダ、オーストラリアを始め様々な国でLGBTQ+の権利についての啓発を行うイベントが開催されるPride Monthとされています。Mattermost社は仕事をする上でもPrideを尊重することは生産性と幸福を高めると考えており、Pride Monthに関する意識を向上させるために、6月の毎週金曜日にPride Monthに関するクイズを出し、最初に正解した人にグッズを贈呈するというイベントを開催しています。

![etc-hpm](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/etc-hpm.png)  
https://community-daily.mattermost.com/core/channels/off-topic-pub/9hwibadgfp8jxj8sdfur5qt7kw

### Hackathon

Mattermost社のプロダクト(Channels, Playbooks, Boardsなどを指しているものだと思います)を、うまく活用したソリューションのアイデアを集めるためのハッカソンが6月下旬に予定されています。  
Mattermost社のR&Dチームが主体となって開催しているものですが、コミュニティメンバーの参加も歓迎されています。

[Mattermost Hackathon 2022 \- Google スライド](https://docs.google.com/presentation/d/1ZBlYdwEShXjNTG_6EeqrqxuS-ceMmO27otg3F2KO56Q/edit#slide=id.ge9b7e537ec_0_0)

![etc-hackathon](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.0/etc-hackathon.png)  



## おわりに
次の`v7.1`のリリースは 2022/07/15(Fri)を予定しています。

---
