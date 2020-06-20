---
title: "Mattermost5.22の新機能"
date: 2020-04-17T22:56:12+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/4/16にMattermost v5.22.0がリリースされました。

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost 5\.22: Read\-only channels & channel moderation settings in E20 \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-22-read-only-channels-channel-moderation-settings-in-e20/)
* https://docs.mattermost.com/administration/changelog.html#release-v5-22-feature-release

---


## チームサイドバーの改善

Mattermostは一つのサーバーに複数のチームを作成でき、複数のチームに所属していると画面の左側にチーム切り替え用のサイドバーが表示されます。

今回のリリースで、ショートカットキーによるチームの切り替えができるようになりました。
**Ctrl + Alt** (Windows) / **⌘ + ⌥** (Mac) を押しながら **↑/↓** キー、もしくはチームアイコンの脇に表示された数字のキーを押すことでチームを切り替えることができます。また、ドラッグ&ドロップでチームアイコンの順番を変えることができます。

![switch team](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/switch-team.gif)

## アーカイブしたチャンネルの復元

今まで一度アーカイブしたチャンネルを復元するには [サーバー管理者向けのコマンドラインツール](https://docs.mattermost.com/administration/command-line-tools.html#mattermost-channel-restore) を使用するか、 [API](https://api.mattermost.com/#tag/channels/paths/~1channels~1%7Bchannel_id%7D~1restore/post)からリクエストを送信する必要がありました。

今回のバージョンから、MattermostのWeb画面からチャンネルを復元するメニューが使用可能になりました。
メッセージ検索結果などからアーカイブされたチャンネルに移動し、チャンネルメニューからチャンネル復元のメニューを選択することで復元可能になります。
(注意: 翻訳ミスのため、メニュー名が**「チャンネルをアーカイブする」**と表示されてしまっています...次回リリースの`v5.23.0`で修正されているはずです。)

![unarchive](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/unarchive.png)

## Confluenceプラグインの改善

Atlassian Confluenceを利用しているユーザー向けに、Cofluenceプラグインが新しくなりました。
この改善により、複数のConfluenceサーバーへ接続できるようになり、また様々なConfluence上のイベントを関連するMattermostチャンネルに流すことができるようになります。

イベントを購読したいMattermostチャンネルで、`/cofluence subscribe`というコマンドを実行することで設定することができます。

詳しくは、下記を参照してください。
https://mattermost.gitbook.io/plugin-confluence/


## (実験的な機能) チャンネルサイドバーの改善

現在、Mattermostではチャンネルサイドバーの改善が進んでいます。
今回のリリースでは、実験的な機能としてチャンネルサイドバーに新たな機能が追加されています。

* サイドバーのカテゴリ(お気に入り、公開チャンネル、非公開チャンネル、ダイレクトメッセージ)が折り畳み可能に
* 未読チャンネルのみが表示される未読フィルダボタンの設置
* チャンネル閲覧履歴に基いたチャンネル移動が可能なボタンの設置

![sidebar](https://mattermost.com/wp-content/uploads/2020/04/5.22-1.gif)

(画像は[公式ブログ](https://mattermost.com/blog/mattermost-5-22-read-only-channels-channel-moderation-settings-in-e20/)から転載)

この実験的な機能を有効にするには、まずシステム管理者によってシステムコンソールから **実験的なサイドバー機能** を有効にします。

![system_console](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/sidebar-system-console.png)

その後、ユーザーごとに **アカウント設定 > サイドバー > 実験的なサイドバー機能** をオンにすることで有効になります。

![account_settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/sidebar-account-settings.png)


## チャンネルモデレーション設定 (E20)

（[Enterprise E20](https://mattermost.com/pricing/) 限定の機能です）

システム管理者がチャンネルごとに権限設定を行えるようになりました。
この機能により、以下のようなことが可能になります。

* チャンネル管理者が重要事項を通知するための読み込み専用アナウンスチャンネルの作成
* チャンネル全体へのメンション (@channel, @here, @all) の無効化
* チャンネルメンバーの変更を許可されたチャンネル管理者のみに制限

チャンネル設定画面から設定できます。

![moderation](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/channel-moderation.png)

詳細について公式ドキュメントを参照してください。
https://docs.mattermost.com/deployment/advanced-permissions.html#channel-moderation-beta-e20


## Extended Support Rellease(ESR)

Mattermostは毎月16日にリリース日が設定されていますが、セキュリティアップデートなどがサポートされるのはリリースから3ヶ月後までです。毎月、もしくは3ヶ月ごとのアップデートが負担になるユーザー向けに昨年11月から9ヶ月サポートが続く Extended Support release (ESR) バージョンが半年に一度リリースされることになりました。

今では`v5.9.0`がESRとして設定されていましたが、そのサポートが終了し、今後は2020/01/16にリリースされた`v5.19.0`が ESR リリースとして 2020/10/16 までサポートされます。

詳細は下記を参照してください。

* https://forum.mattermost.org/t/upcoming-extended-support-release-updates/8526
* https://docs.mattermost.com/administration/extended-support-release.html

---

## 破壊的変更

### 絵文字リアクションポップアップのパフォーマンス改善
カスタム絵文字を多く登録すると、絵文字リアクションのポップアップ表示が重くなることがありました。今回のリリースで、パフォーマンス改善のために絵文字リアクション関連のスキーマ変更が行わレました。しかし、多くのカスタム絵文字が登録されている場合、スキーマ変更のアップデート処理に時間がかかることがあるため、利用者の少ない時間帯でのアップデートが推奨されています。

### チャンネルモデレーション設定のモバイルアプリ対応
今回のリリースで追加されたチャンネルモデレーションの設定は、モバイルアプリ v1.30 以降でサポートされています。それ以前のバージョンを使用している状態で、チャンネルモデレーション設定により投稿/絵文字リアクションが制限されているチャンネルで禁止された行為を行った場合、エラーが表示されてしまします。

### `model.Post`の`Props`フィールドへの直接アクセスの制限
[mattermost/mattermost-server](https://github.com/mattermost/mattermost-server/) の [`model.Post`](https://github.com/mattermost/mattermost-server/blob/master/model/post.go#L68)を利用した統合機能・プラグインを開発しているユーザー向けの注意事項です。
`model.Post`内の `Props` フィールドへの直接アクセスが禁止されたため、今後は`GetProps()`・`SetProps()`メソッドを利用してアクセスする必要があります。

## キーボードショートカットの追加

* **Ctrl/⌘ + .**: 右サイドバーのopen/close
* **Ctrl/⌘ + shift + /**: 直前のメッセージにリアクション

キーボードショートカットは **Ctrl / ⌘ + /**で確認することができます。

![shortcut](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.22/shortcut.png)



# おわりに

次の`v5.23`のリリースは2020/5/15(Fri)を予定しています。
そして機能追加が行われる`v5.24`は恐らく2020/6/16(Tue)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
