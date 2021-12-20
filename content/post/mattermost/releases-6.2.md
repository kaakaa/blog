---
title: "Mattermost 6.2の新機能"
emoji: "🎉"
published: true
date: 2021-12-20T12:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2021/12/16 に Mattermost のメジャーバージョンアップである `v6.2.0` がリリースされました。
また、2021/12/17にセキュリティFixを含むv6.2.1がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.2 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v6-2-is-now-available/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6.2-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/install/self-managed-changelog.html#important-upgrade-notes)

本バージョンより、チャンネル補完機能の対象に非公開チャンネルが追加されます。
それに伴い、 BleveやElasticsearchを利用した自動補完機能を有効にしている場合、インデックスの再構築が必要になります。この処理は時間がかかる可能性があるため、アップグレードする際は自動補完機能を一時的に無効にし、バックグラウンド処理によるインデックス構築が完了した後に、再度自動補完を有効にすることをお勧めします。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))


## Channels: 非公開チャンネルの自動補完

Mattermostでメッセージを投稿する際、メッセージ入力欄に`~`を入力するとチャンネル一覧が表示され、このチャンネル一覧から選択することでチャンネルへのリンクを簡単に入力できるという機能がありました。
今までのバージョンでは、この一覧に表示されるチャンネルは公開チャンネルのみでしたが、本バージョンからは自分が所属している非公開チャンネルも一覧に表示されるようになります。

![Autocomplete Private Channel](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/channels-autocomplete-private.gif)

## Channels: 返信スレッドに関する操作性改善

[返信スレッドの折り畳み機能](https://docs.mattermost.com/messaging/organizing-conversations.html#organizing-conversations-using-collapsed-reply-threads-beta)を有効にしている場合、`スレッド`ビューにて表示するスレッドを`↑`、`↓`キーで切り替えることができるようになりました。

![Move thread by key](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/channels-thread-move.png)

また、返信スレッドが存在する場合、メッセージのどの部分をクリックしても返信スレッドが表示されるようになりました。この機能は **設定 > 表示 > クリックしてスレッドを開く**の設定から無効にすることもできます。

![Click to open reply](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/channels-click-to-open-reply.png)

## Playbooks: 実行中のPlaybookのフォロー

実行中のPlaybookをフォローする機能が追加されました。
Playbookの実行をフォローすることで、気にかけている作業の開始や完了、ステータス更新、期日の超過、レトロスペクティブ(振り返り)の公開などが通知されるようになります。また、Playbookの全ての実行を自動でフォローするよう設定することもできます。

![Follow run](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/playbooks-auto-follow-run.png)

その他にも、Playbook検索機能やPlaybook共有URL、実行中のPlaybookのフィルタ条件の追加などの更新が行われています。

## Boards: カレンダー表示

Board上のカードをカレンダー上に表示できるようになり、締め切り管理がやりやすくなりました。

![Calendar View](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/boards-calendar-view.png)

デフォルトでは、カードの作成日をもとにカレンダー表示が行われますが、表示プロパティを切り替えることで、カードに設定されたプロパティの情報を元にカードを表示することができます。

![Calendar property](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.2/boards-switch-calendar-property.png)

## その他のトピック

### Mattermostのインスタンス統合により得た知見

弊社で2つのMattermostインスタンスを統合した事例について、アドベントカレンダーの記事として公開しました。

[Lessons learned from merging Mattermost instances](https://zenn.dev/kaakaa/articles/mattermost-merge-instances)

2つのMattermostインスタンスの統合は公式ではサポートされていなかったため、統合の際に問題となった点や対処方法、また失う可能性のあるデータ等についてまとめ増田。

### 開発者の生産性に関するガイドの発行

Mattermostが、約300名の開発者に対する調査結果等をまとめた [Unblocking Workflows: The Guide to Developer Productivity in 2022](https://mattermost.com/guide-to-developer-productivity/) という開発者の生産性に関するガイドを発行しています。

## おわりに
次の`v6.3`のリリースは 2022/01/14(Fri)を予定しています。
