---
title: "Mattermost 9.6の新機能"
emoji: "🎉"
published: true
date: 2024-03-25T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/03/15 に Mattermost のアップデートとなる `v9.6.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog \- Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-6-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube hHd05qSnWDI >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。


## 外向きのOAuth接続

統合機能の一つとして、外向きのOAuth接続(Outgoing OAuth Connection)が追加されました。  
詳しくは、公式ドキュメント[Outgoing OAuth connections](https://developers.mattermost.com/integrate/slash-commands/outgoing-oauth-connections/)を参照ください。

正しく理解できているか自信がないですが、スラッシュコマンドを使って認証が必要なシステムにリクエストを送信する際に、認証に必要なOAuthトークンを取得してくれる中間者として動作する統合機能だと理解しています。

[![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/outgoing-oauth-connection.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/outgoing-oauth-connection.png)

## システムコンソール > ユーザー管理画面の改善

システムコンソールのユーザー管理画面が刷新され、ユーザーに関する様々な情報を一覧表示することができるようになりました。  

[![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/user-management.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/user-management.png)

チーム・ロールによるフィルタリング、表示項目の選択、集計範囲の設定等を行うことができます。

[![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/user-management-settings.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.6/user-management-settings.png)

また、Professional/Enterprise版限定機能として、ユーザー情報のCSV Exportを実行することもできます。


## その他の変更

### Webp画像のプレビューが可能に

本バージョンからWebp画像のプレビューが可能になりました。

## その他のトピック

### Mattermost AI Copilot

Mattermost AI Pluginとして開発されていたプラグインが、現在は**Mattermost AI Copilot** Pluginと名前を変え、開発が進められているようです。

[Mattermost AI Copilot: Accelerating the conversation with LLMs \- Mattermost](https://mattermost.com/blog/mattermost-ai-copilot-accelerating-the-conversation-with-llms/)

以前、[紹介記事を書いた頃](https://zenn.dev/kaakaa/articles/mattermost-plugin-ai)から機能もいくつか追加されているようです。  
[mattermost\-plugin\-ai/docs/usage\.md at master · mattermost/mattermost\-plugin\-ai](https://github.com/mattermost/mattermost-plugin-ai/blob/master/docs/usage.md)


## おわりに
次の`v9.7`のリリースは 2024/4/16(Tue)を予定しています。  
