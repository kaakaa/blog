---
title: "Mattermost5.6の新機能"
date: 2018-12-16T15:17:06+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---


# はじめに

2018/12/14にMattermost 5.6.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost Changelog — Mattermost 5\.6 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-6)
* [Mattermost 5\.6: Interactive message dialogs, new admin tools, Ukrainian language support and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-6-interactive-message-dialogs-new-admin-tools-ukrainian-language-support-and-more/)

----

2018/11/16にMattermost 5.5.0がリリースされていましたが、事前の予告通りBug Fixと小さな改善しかなかったため記事は書きませんでした。
今後も目ぼしい昨日がなければ奇数バージョンのリリースについては記事を書かないつもりです。

* [Mattermost Changelog — Mattermost 5\.6 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-5)
* [Mattermost 5\.5: Web conferencing integration, Hacktoberfest contributions and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-5-web-conferencing-integration-hacktoberfest-contributions-and-more/)

---

# アップデート内容

## インタラクティブダイアログ

インタラクティブボタンやスラッシュコマンドなどの連携機能への操作に対して、ダイアログを表示することができるようになりました。

[Interactive Dialogs — Mattermost 5\.6 documentation](https://docs.mattermost.com/developer/interactive-dialogs.html#dialog-submission)
![dialog.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/91475afb-5f52-6ffc-a758-6cdfb0599983.gif)

スラッシュコマンドからダイアログを呼び出す場合のサンプルコードは下記のようになります。
https://gist.github.com/kaakaa/1f87869f0d6520a2f645f8af014be1a5

## ユーザー体験の向上

### ダイレクトメッセージ作成画面で、グループメッセージを検索可能に

ダイレクトメッセージの追加画面で、グループメッセージ（複数人でのダイレクトメッセージ）を検索できるようになりました。
![dm.png](https://qiita-image-store.s3.amazonaws.com/0/9891/c1fcf831-d768-9ed6-e281-697c50157694.png)

### ファイルアップロード時のプログレスバー

ファイルをアップロードした際に、アップロード完了までのプログレスバーが表示されるようになりました。

![image.png](https://qiita-image-store.s3.amazonaws.com/0/9891/d21fcb32-d4b7-5871-dab6-e1cef60b7303.png)

### 管理者バッジ

投稿者のアイコンをクリックすることで開くプロフィールポップオーバーに、管理者であることを示すバッジが表示されるようになりました。

![profile.png](https://qiita-image-store.s3.amazonaws.com/0/9891/de1012a2-f10e-d97e-21d1-c050710e78a8.png)

### その他の改善

#### 管理者向けコマンドラインツール改善

`mattermost`コマンドにウェブフックやカスタムスラッシュコマンドを操作するコマンドなどが追加されています
https://docs.mattermost.com/administration/changelog.html#command-line-interface-cli

#### スラッシュコマンドに複数レスポンス
スラッシュコマンドを実行した際、複数のレスポンスを一度に返せるようになりました。  
サンプルコード: https://gist.github.com/kaakaa/b040bcd0a6fdccc18b3daaea214c6112

#### WebRTC機能の廃止
ベータ版として提供されていたWebRTC機能が廃止されました。代わりに[Zoom](https://docs.mattermost.com/integrations/zoom.html), [BigBlueButton](https://github.com/blindsidenetworks/mattermost-plugin-bigbluebutton), [Kopano](https://kopano.com/releases/a-first-look-at-the-new-kopano-web-meetings/)などのプラグインの使用が推奨されています。
https://forum.mattermost.org/t/built-in-webrtc-video-and-audio-calls-removed-in-v5-6-in-favor-of-open-source-plugins/5998


## サポート言語の追加

ウクライナ語がサポートされ、サポート言語が16言語になりました。また、ルーマニア語がベータ版として追加されています。
その他にもチェコ語、セルビア語、スウェーデン語の翻訳作業も始まっているようです。

また、シンガポール国立大学がMattermostを採用したというニュースもありました。
下記の記事ではコストとスケーラビリティの点でMattermostが採用されるまでの経緯が綴られています。
[The National University of Singapore chooses Mattermost for student\-faculty collaboration \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/the-national-university-of-singapore-chooses-mattermost-for-student-faculty-collaboration/)

## モバイルアプリの改善

* メンションのハイライト
* ピン留め機能のサポート
* ジャンボ絵文字のサポート
* ネットワーク問題が発生した際の自動再接続機能によるネットワーク接続の改善

その他のアップデート内容については下記を参照ください
https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#1150-release

## デスクトップアプリ

Mattermostデスクトップアプリ v4.2.0のリリースも行われています。
https://mattermost.com/download/#mobile

* Ctrl/CMD+F でMattermostのメッセージが検索できるようになりました
* 英語・ポルトガル語・スペイン語(スペイン/メキシコ)のスペルチェック機能が追加されました
* コンピューターの電源をONにした際に自動で起動する設定がデフォルトで有効になりました（アプリの設定で無効にすることもできます）
* MacOS向けの .dmg インストーラーが追加されました

その他のアップデート内容については下記を参照ください
https://github.com/mattermost/desktop/blob/master/CHANGELOG.md#release-v420


# おわりに

次回の`v5.7`のリリースは2019/1/16(Fri)を予定しています。
そして機能追加が行われる`v5.8`は恐らく2019/2/15(Fri)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。

