---
title: "Mattermost 7.4の新機能"
emoji: "🎉"
published: true
date: 2022-10-15T16:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2022/10/14 に Mattermost のアップデートとなる `v7.4.0` がリリースされました。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v7\.4 is now available \- Mattermost](https://mattermost.com/blog/mattermost-v7-4-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-4-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(OSS 版)でも利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- Platform: 上記機能を含むMattermost全体

## Calls: キーボードショートカットが利用可能に

Calls機能の操作を行うためのキーボードショートカットが追加されました。
ミュート解除やスクリーン共有などの操作をする際にマウスを操作する必要がなくなるので、キーボードから手を離す必要がなくなります。

現在サポートされているキーボードショートカットは以下の通りです。

![calls-shortcut](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/calls-shortcut.png)

[Keyboard shortcuts for Channels](https://docs.mattermost.com/channels/keyboard-shortcuts-for-channels.html#calls-shortcuts)

## Boards: Commenter/Viewerロールの追加

Boardsの権限ロールに、コメントの追加が可能な**Commenter**と、Boardおよびその内容の閲覧のみが可能な**Viewer** が追加されました。これらの権限が追加されたことで、Board内容の意図せぬ変更の可能性を減らすことができるようになります。

Boardsに対する権限ロールの設定は、Boards画面右上の **Share** ボタンから行うことができます

![boards-new-role](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-new-role.png)

---

現在、**Commenter** の日本語訳が **コメントした人** となっていますが、次バージョンでは**コメンター**に修正されているかと思います。

また、 **Viewer**(閲覧者)以外のロールが設定されたユーザーのロールを変更しようとすると、**Viewer**(閲覧者)の表示位置がずれてしまいます。

![boards-new-role-misaligned](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-new-role-misaligned.png)

こちらについても既に修正済みのため、次バージョンで修正されているかと思います。  
[Fix misaligned 'viewer' role on share board/template dialog by vivekkj123 · Pull Request \#3926 · mattermost/focalboard](https://github.com/mattermost/focalboard/pull/3926)


## Boards: ゲストアカウントでもBoards機能が利用可能に

Mattermostには昔から組織外のユーザー用に[**ゲストアカウント**](https://docs.mattermost.com/onboard/guest-accounts.html)を作成する機能がありましたが、本バージョンからゲストアカウントもBoardにアクセスできるようになりました。　 
ただし、正規のユーザーとは異なり「新しいBoardの作成は行えない」や「Boardの管理者権限は付与できない」などの制限があります。詳しくは[公式ドキュメント](https://docs.mattermost.com/boards/share-and-collaborate.html#guest-accounts)を参照してください。

## Boards: メンバー追加時に自動補完リストが表示されるように

Cardの内容やコメントを追加する際に `@` を入力することで、ユーザーの自動補完リストが表示されるようになりました。  

![boards-autocomplete](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-autocomplete.png)

もし、Boardのユーザーではないユーザーの場合、自動補完リストのユーザー名の横に **(not board member)** と表示されます。  
また、このユーザーに対する@メンションを行おうとすると、そのユーザーの権限を設定するダイアログが表示されます。

![boards-autocomplete-permission](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-autocomplete-permission.png)

これにより、@メンションの対象となったユーザーをBoardに追加するとともに、そのユーザーの権限も同時に設定することができます。

この自動補完リストは、Cardのプロパティでも動作します。

![boards-autocomplete-property](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-autocomplete-property.png)

## Boards: Boardのリンク/リンク解除操作がチャンネルに通知されるように

ChannelsのAppBar(右サイドバー)からチャンネルをBoardにリンク、もしくはチャンネルとのリンクの解除を実行すると、チャンネルに通知メッセージが投稿されるようになりました。

![boards-channel-notification](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-channel-notification.png)

## Boards: Multi-personプロパティの追加

Cardのプロパティに複数のユーザーを指定できる **Multi-person** というプロパティ種別が追加されました。

![boards-multi-person](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/boards-multi-person.png)

今までのバージョンでは、一人のユーザーのみ指定できる **Person** タイプしか選択できませんでした、そのため、例えばCardに対する担当者(Assignee)を複数人設定したい場合は、複数のプロパティ(**Assignee1**, **Assignee2**, ...)を作成する必要がありました。しかし、**Multi-person**が追加されたことで、一つのプロパティで複数人の担当者を扱えるようになります。

## その他の変更

### メッセージ転送機能における日本語入力時の不具合改善

Mattermost v7.2で追加されたメッセージ転送(Message Forwarding)機能にて転送メッセージを日本語で入力する際、入力確定の `Enter` を押下すると、意図せずメッセージ転送が実行されてしまうという不具合がありました。  
[Mattermost 7.2の新機能 - 注意: 日本語入力時の不具合について](https://blog.kaakaa.dev/post/mattermost/releases-7.2/#%E6%B3%A8%E6%84%8F-%E6%97%A5%E6%9C%AC%E8%AA%9E%E5%85%A5%E5%8A%9B%E6%99%82%E3%81%AE%E4%B8%8D%E5%85%B7%E5%90%88%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)

この不具合はv7.4で解消されました。  
[Unable to add Japanese comments correctly in Message Forwarding Dialog · Issue \#20838 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/20838)


## その他のトピック

### Hacktoberfest

今年もMattermostは[Hacktoberfest](https://hacktoberfest.com/)に力を入れています。  
[Hacktoberfest is here again\! \- Mattermost](https://mattermost.com/blog/hacktoberfest-2022/)

[Hacktoberfest](https://hacktoberfest.com/)の期間中(10月中)にMattermostへのコントリビューを行うと、オリジナルのステッカーが貰えるそうです。また、期間中のトップコントリビューターに選ばれると、Tシャツやマグカップ等も貰えるようです。

![event-hacktoberfest-swags](https://blog.kaakaa.dev/images/posts/mattermost/releases-7.4/event-hacktoberfest-swags.webp)

### Mattermost関連の記事紹介

#### Mattermost構築

* [Mattermost を Docker でローカルサーバにインストール \- Qiita](https://qiita.com/nanbuwks/items/b20e2df483f6806909ab)
  * [mattermost/docker](https://github.com/mattermost/docker)を利用したMattermost構築方法について
* [MattermostサーバをGCPパッケージを使って構築する \- Qiita](https://qiita.com/Power-mind74/items/6e8cbe5cb42d7094ca57)
  * GCPにMattermostを構築する手順について
* [森　崇さんはTwitterを使っています: 「仕事でmattermostに大量のユーザ/ポストを登録必要があり、バッチ・スクリプト群を一般公開ししました。 負荷テストやるときに便利と思います〜。 https://t\.co/0ErNlJhI2w \#mattermost」 / Twitter](https://twitter.com/kanetugu2020/status/1577185142896230401)
  * Mattermostに任意のデータを一括登録するスクリプトの紹介

#### Tech

* [Mattermost サンプルデータの作成](https://zenn.dev/kiyasu7028/articles/20295d293aa0ae)
  * 公式CLIツール([mmctl](https://github.com/mattermost/mmctl))を使ったMattermostサンプルデータの自動生成について
* [Mattermost 投稿内容の出力](https://zenn.dev/kiyasu7028/articles/f2a59351495c58)
  * 公式CLIツール([mmctl](https://github.com/mattermost/mmctl))を使ったMattermost投稿データの取得について
* [Mattermost 投稿数順に集計して通知](https://zenn.dev/kiyasu7028/articles/83b38aff54f4b6)
  * 投稿数ランキングをMattermostで通知する方法について (Python)
* [Mattermot APIのPHPドライバを Laravel で実行する手順 \- Qiita](https://qiita.com/kanetugu2018/items/bf0e22a58d2ecd8f4062)
  * Mattermost PHPドライバ (コミュニティ製) の実行方法の紹介

## おわりに
次の`v7.5`のリリースは 2022/11/16(We)を予定しています。

