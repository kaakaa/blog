---
title: "Mattermost 7.8の新機能"
emoji: "🎉"
published: true
date: 2023-02-19T14:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/02/16 に Mattermost のアップデートとなる `v7.8.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.8 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-8-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html)

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


## Extended support release (ESR)

今回リリースされたMattermost v7.８は、Extended Support Release (ESR)として2023/11/15までサポートされる予定です。

Mattermosthは毎月16日に新しいバージョンがリリースされていますが、次のバージョンがリリースされた後でも、リリースから3ヶ月間はバグ修正やセキュリティ修正がおこなわれます。Extended Support Release (ESR)に設定されたバージョンは、このサポート期間が3ヶ月から9ヶ月に延長されたバージョンになります。毎月バージョンアップを行うだけの時間が割けない場合は、半年ごとにリリースされるExtended Support Release (ESR)ごとにアップデートをかけることをお勧めします。

[Extended Support Release](https://docs.mattermost.com/upgrade/extended-support-release.html)

## Boards: 人物または日付によるカードのフィルタ

`日付(Date)`、`人物(Person)`、または`複数人(Multi Person)`のプロパティタイプによるフィルタリングを行えるようになりました。

例えば、あるボード内のカードにプロパティタイプ`複数人(Multi Person)`の`関係者`という属性を追加し、自分を設定したとします。

![boards-multi-person](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.8/boards-multi-person.gif)

すると、Boardの`フィルター`メニューからプロパティタイプ`複数人(Multi Person)`で作成したプロパティを選択してカードをフィルタすることができるようになります。

![boards-filter-by-person](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.8/boards-filter-by-person.gif)

同様に、プロパティタイプ `日付(Date)`、`人物(Person)`で作成したプロパティも`フィルター`の項目として選択できるようになります。

## Boards: 人物によるカードのグループ化

プロパティタイプ `人物(Person)` によるカードのグループ化を行うことができるようになりました。

これにより、担当ごとのタスク一覧等を確認しやすくなります。

![boards-group-by-person](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.8/boards-group-by-person.png)

## アップグレード時の注意事項

アップグレード時の注意事項について、詳しくは公式ドキュメントを確認ください。　 
[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

### メッセージの優先度設定と既読確認機能がデフォルトで有効に

Mattermost v7.7で追加された[メッセージの優先度設定と既読確認](https://blog.kaakaa.dev/post/mattermost/releases-7.7/#channels-%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8%E3%81%AE%E5%84%AA%E5%85%88%E5%BA%A6%E8%A8%AD%E5%AE%9A%E3%81%A8%E6%97%A2%E8%AA%AD%E7%A2%BA%E8%AA%8D)機能がデフォルトで有効になりました。システムコンソールから無効化することもできます。

### アップグレード中、Boardがロックされないようにする

新しいバージョンのBoardsを起動する際に、DBの`focalboard_category_boards`テーブルに重複するデータが存在する場合、Boards起動中にユーザーがBoardsを利用できなくなることがあるようです。アップグレードを行う前に、当該テーブルに重複がないか確認し、重複データが存在する場合は事前に削除することが推奨されています。

重複データの確認・削除のクエリは、以下のドキュメントで公開されています。

[Important Upgrade Notes — Mattermost documentation](http://mattermost-docs-preview-pulls.s3-website-us-east-1.amazonaws.com/6187/upgrade/important-upgrade-notes.html)

## おわりに
次の`v7.9`のリリースは 2023/03/16(Thu)を予定しています。
