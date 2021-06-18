---
title: "Mattermost5.31の新機能"
date: 2021-01-17T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/01/16にMattermost v5.31.0がリリースされました。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

今回のリリースから毎バージョン新機能が追加されることになったため、新機能紹介記事も毎月書いていきます。
(v5.30までは、新機能リリースと機能改善リリースを交互にリリースしていたため、新機能紹介記事も隔月で書いていました)


変更内容の詳細については公式のリリースを確認してください。

* [Mattermost v5\.31 is now available \- Mattermost Extended Support Release](https://mattermost.com/blog/mattermost-release-v5-31/)
* [Mattermost Changelog — Mattermost 5\.31 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-31-esr)

---

## (E20) Incident Managementプラグインの改善

v5.30から導入されたMattermost上でインシデント管理ができるIncident Mangementプラグインについて、機能改善が行われました。(Incident Managementプラグインについては、[前回の記事](https://blog.kaakaa.dev/post/mattermost/releases-5.30/#e20-mattermost%E3%81%A7%E3%82%A4%E3%83%B3%E3%82%B7%E3%83%87%E3%83%B3%E3%83%88%E7%AE%A1%E7%90%86)で無償版でも使えそうと書いていましたが、やはりEnterprise版限定のプラグインのようです)

新しいバージョンのプラグインでは、インシデント対応のテンプレートを定義するPlaybookを作成する際の設定項目が追加されています。
* インシデントに対する更新があった場合に、更新内容を通知するチャンネルを指定
* インシデントに関するリマインダを通知する時間を指定
* メッセージのテンプレートを追加

![incident plugin](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.31/update-incident-plugin.png)

## プラグインの紹介
### GitLabプラグインのアップデート

MattermostでGitLabの更新内容を受け取れるGitLabプラグインがアップデートされ、セットアップが簡単にできるようになったようです。

[Mattermost/GitLab Integration \- GitLab Plugin](https://mattermost.gitbook.io/plugin-gitlab/)

### Icebreakerプラグイン

コロナ禍のコミュニケーション対策として、アイスブレイクとなる質問を投稿するBotを追加するプラグインの紹介です。
[monsdar/mattermost\-icebreaker\-plugin: A Mattermost plugin that asks Icebreaker questions](https://github.com/monsdar/mattermost-icebreaker-plugin)

### GIF コマンドプラグイン

Mattermost上に[GIPHY](https://media.giphy.com/headers/2020-09-10-58-1599746331/NFL_BANNER_HP.gif)からGifアニメーションを投稿するプラグインの紹介です。過去にGiphyプラグインとして公開されていたプラグインの名前が変わったもののようです。
[moussetc/mattermost\-plugin\-giphy: A Giphy/Gfycat/Tenor plugin for Mattermost \(/gif&/gifs commands\)](https://github.com/moussetc/mattermost-plugin-giphy#mattermost-gif-commands-plugin-ex-giphy-plugin-)

## その他の更新

### 実験的なチャンネルサイドバーの改善

[Mattermost v5.26より実験的な機能としてチャンネルサイドバーに独自のカテゴリを作成できるようになりました](https://blog.kaakaa.dev/post/mattermost/releases-5.26/#%E5%AE%9F%E9%A8%93%E7%9A%84%E3%81%AA%E6%A9%9F%E8%83%BD-%E3%83%81%E3%83%A3%E3%83%B3%E3%83%8D%E3%83%AB%E3%81%AE%E3%82%AB%E3%83%86%E3%82%B4%E3%83%AA%E5%8C%96%E3%81%A8%E9%A0%86%E5%BA%8F%E3%81%AE%E5%85%A5%E3%82%8C%E6%9B%BF%E3%81%88)が、作成したカテゴリ単位でチャンネルのミュートを指定できるようになりました。

![mute category](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.31/mute-category.png)

また、ドラッグ＆ドロップでチャンネルを移動する際に、`Ctrl + クリック`や`Shift + クリック`でチャンネルを複数選択することで、同時に複数のチャンネルを移動できるようになりました。

## 今後のリリース予定

### 返信の折畳み機能
以前より開発が進められていた、返信スレッドをチャンネル上に表示しないようにできるCollapsed Reply Thread(返信スレッドの折畳み)機能が2021年1~3月中に追加される予定です。Mattermost v5.29.1で、Collapsed Reply Thread向けのデータベースの更新が行われているため、このバージョンより前のバージョンを使用している場合は、アップデートしておくことが推奨されています。
[Dev Sneak Peek: Collapsed Reply Threads in Mattermost](https://mattermost.com/blog/dev-sneak-peek-collapsed-reply-threads/)

### `mmctl`コマンドのTLS 1.0, 1.1接続の無効化

リモート環境からMattermostを操作するコマンドラインツール[`mmctl`](https://docs.mattermost.com/administration/mmctl-cli-tool.html)を使用する際、接続先のMattermostサーバーがTLS 1.0, 1.1を使用している場合、明示的にオプションを指定しない限りエラーとなるようになります。

### `platform`コマンドがdeprecatedに

Mattermostサーバーを操作するコマンドラインツールとして`platform`コマンドがありましたが、`platform`コマンドは過去のコマンドであり、今後のリリースでdeprecatedとなる予定のため使用しないことが推奨されています。代わりに[`mattermost`コマンド](https://docs.mattermost.com/administration/command-line-tools.html)を使用してください。

## おわりに
次の`v5.32`のリリースは2021/02/16(Tue)を予定しています。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
