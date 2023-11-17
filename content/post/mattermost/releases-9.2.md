---
title: "Mattermost 9.2の新機能"
emoji: "🎉"
published: true
date: 2023-11-17T22:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/11/16 に Mattermost のアップデートとなる `v9.2.2` がリリースされました。  

今までCloud版のMattermostとセルフホスト版のMattermostは異なるリリースサイクルを採用していましたが、本バージョンからリリースサイクルを揃えることになった影響で、セルフホスト版のリリースバージョンが`v9.2.2`となっています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog](https://docs.mattermost.com/deploy/mattermost-changelog.html)

Changelogのページも新しくなっており、`v9.1`以前のバージョンのChangelogは[Mattermost legacy self\-hosted changelog](https://docs.mattermost.com/deploy/legacy-self-hosted-changelog.html)に移動されています。

本バージョンでの主な変更点に関する3分程度の動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube udC2OCTGooc >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体

## (Channels) 複数行のヘッダーを編集した際のシステムメッセージ改善

複数行のチャンネルヘッダーを編集する際、ヘッダーを編集したことを通知するシステムメッセージでは改行が削除された形で書き込まれていましたが、本バージョンから改行を維持したまま編集内容を表示するようになりました。

* Mattermost v9.1 (以前のバージョン)  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.2/channels-sysmsg-header-v91.png)

* Mattermost v9.2 (今回のバージョン)  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.2/channels-sysmsg-header-v92.png)

## その他の変更

特になし  
(今回のバージョンアップでは、インターフェースの変更はあまりなく、バックグラウンドのパフォーマンス改善などが主だったようです)

## その他のトピック

### Mattermost MVP

今回のリリースで、3度目のMattermost MVPに選出いただきました。  
[Mattermost MVP](https://developers.mattermost.com/contribute/more-info/mvp/)

あまり主だった活動はしていなかった認識だったので面食らいましたが、貰えるものはありがたく頂戴したいと思います。

## おわりに
次の`v9.3`のリリースは 2023/12/15(Fri)を予定しています。  
