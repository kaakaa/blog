---
title: "Mattermost 9.7の新機能"
emoji: "🎉"
published: true
date: 2024-04-16T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/04/16 に Mattermost のアップデートとなる `v9.7.0` がリリースされました。  
同日にノイジーなログの削除 [PR](https://github.com/mattermost/mattermost/pull/26713) という軽微な修正を加えた `v9.7.1` がリリースされています。また、現在 `v9.7.2` のリリース作業も進んでいるようです。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog \- Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-7-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube ObfVwIYDXAQ >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。


## Mattermost AI Plugin

> ・Added Mattermost AI plugin to pre-packaged plugins.

Mattermost AI Copilot Pluginがpre-packagedなプラグインとして提供されるようになったようです。  
pre-packagedなプラグインは、Mattermostをv9.7アップデートするだけで自動でインストールされるはずですが、もし**システムコンソール > プラグイン > プラグイン管理 > インストール済みプラグイン** にAI Copilotプラグインが表示されていない場合、Mattermostのアプリマーケットプレースからインストールすることが可能です。

[![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/ai-copilot-marketplace.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/ai-copilot-marketplace.png)

AI Copilotプラグインをインストールし、有効化したのち、使用するAIサービスの設定等を行うことで利用できるようになります。

[![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/ai-copilot-setting.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/ai-copilot-setting.png)

AIサービスの設定については、公式の [configuration-guide](https://github.com/mattermost/mattermost-plugin-ai/blob/master/docs/configuration-guide.md)を参照ください。

## チーム設定モーダルのUI更新

> ・Updated the user interface of Team Settings modal.

チーム設定モーダルのUIが更新されました。  
設定可能な項目に変更はないようです。

[![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/team-setting.png)](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.7/team-setting.png)

## その他の変更

特になし

## その他のトピック

### Mattermostカスタマイズガイド

Mattermostをカスタマイズするためのガイドとなる資料が公開されているようです。(以前からあった?)

{{< tweet user="mattermost" id="1778776725112988040" >}}

サイドバーの構成方法や、テーマ等のUI設定、表示言語の設定などの基本的な機能からMattermost App/Pluginなどの拡張機能の開発方法、Mattermostへコントリビュートを行う方法など、広範なトピックがドキュメントとしてまとめられています。

[The Guide to Customizing Mattermost for Technical Teams](https://mattermost.com/customizing-mattermost-for-technical-teams/#)

## おわりに
次の`v9.8`のリリースは 2024/5/16(Thu)を予定しています。  
