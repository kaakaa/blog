---
title: "Mattermost 9.5の新機能"
emoji: "🎉"
published: true
date: 2024-02-17T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/02/16 に Mattermost のアップデートとなる `v9.5.0` がリリースされました。  
(同日にバグ修正版の`v9.5.1`がリリースされています)  

本バージョンは[ESR(Extended Support Release)](https://docs.mattermost.com/upgrade/extended-support-release.html)に設定されており、2024/11/15までセキュリティ対応/バグ修正等のサポートが行われる予定です。(ESRでないバージョンのサポート期間は3ヶ月間です)

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog — Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-5-extended-support-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。（本バージョンでの目に見える機能追加はあまりないようです。）

{{< youtube b1M2BGGF578 >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。


## (Professional/Enterprise) スレッドを別チャンネルに移動する機能のリリース

有償版(Professional/Enterprise)限定ですが、メッセージをスレッドごと別のチャンネルへ移動する機能がリリースされました。今まで[Mattermost Wrangler Plugin](https://github.com/gabrieljackson/mattermost-plugin-wrangler)として開発されていたものを、Mattermost本体に取り込んだ機能になります。  

動作画面については、以下のIssueにスクリーンショットがあります。(実際の画面とは異なる可能性があります)  
[Feature: Wrangler by nickmisasi · Pull Request \#23602 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/23602)

本機能は現在は有償版限定機能としてリリースされていますが、機能が成熟した後に無償版でも利用可能とする計画があるようです。  
[Lock wrangler behind enterprise by nickmisasi · Pull Request \#25703 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/25703#issuecomment-1879030136)  


## その他の変更

### MySQL v5.7のサポート終了

本バージョンから、MySQL v5.7がサポート対象外となりました。MySQLを利用する場合はMySQL v8以降の利用が推奨されています。

## その他のトピック

### Mattermost Trustcenter

Mattermostのセキュリティプログラムに関する情報を掲載したMattermost Trustcenterが公開されました。  
[Mattermost Trustcenter](https://trust.mattermost.com/)  
[Announcing the Mattermost Trustcenter \- Mattermost](https://mattermost.com/blog/announcing-mattermost-trustcenter/)

Mattermost Trustcenterでは、Mattermost自体のセキュリティに関する情報が確認でき、具体的にはSOC2等のコンプライアンスレポートやPenTestの結果、セキュリティに関するナレッジ集などが公開されています。(一部、アクセスするためにTrustcenterへの登録を求められるドキュメントもあります)  

### Mattermost Academy

Mattermostのオンライン学習サイトであるMattermost Academyの紹介がありました。  
[Free Mattermost training courses: Mattermost Academy \- Mattermost](https://mattermost.com/blog/free-mattermost-training-courses/)

英語コンテンツのみですが、Mattermostの基本的な使い方から、セルフホストする場合のパフォーマンスに関するTipsなど様々なコースが公開されています。無料で見れるコースも多くあります。（支払い情報を登録する機能はありますが、支払いは実施していないため有償コースが存在するかはよくわかっていません）  

[Mattermost Academy](https://academy.mattermost.com/)

## おわりに
次の`v9.6`のリリースは 2024/3/15(Fri)を予定しています。  
