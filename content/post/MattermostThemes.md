---
title: "Mattermostのテーマ集"
date: 2019-12-15T00:15:49+09:00
draft: false
toc: true
tags: ["mattermost", "theme"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost

# Mattermost テーマ

Mattermostではアカウント設定から画面の表示テーマを変更することができます。
ちょっと気分を変えたい時や、複数のインスタンスに参加している場合、自分が今どのインスタンスにいるかを識別しやすくするためにテーマを変えたりするかと思います。

Mattermostはデフォルトで4種類のテーマ (**Mattermost**, **Organization**, **Mattermost Dark**, **Windows Dark**) を選択することができ、さらにカスタムテーマを利用することで自由にカラーを設定できます。

しかし、カスタムカラーを設定する場合、カラーを設定できる箇所が20箇所以上もあるため、一からテーマを作成するのはなかなか難儀かと思います。

そこで、今回はMattermostのテーマを紹介しているサイトをいくつか紹介していきます。


# カスタムテーマの設定方法
テーマ紹介サイトを紹介する前に、カスタムテーマの設定方法について紹介します。

カスタムテーマの設定方法は非常に簡単で、Mattermostテーマを表すJSON文字列をコピーして、アカウント設定のカスタムテーマの部分にペーストするだけです。

![custom_theme_settings](https://blog.kaakaa.dev/images/posts/mattermost/MattermostThemes/custom_theme_settings.png)

ペーストした時点で画面表示が変わるので、問題なさそうなら **保存** ボタンを押して変更を完了します。

Mattermostテーマを表すJSON文字列は、下記のような形式をした文字列です。

```json
{"sidebarBg":"#146b3a","sidebarText":"#f7f7f7","sidebarUnreadText":"#eba828","sidebarTextHoverBg":"#de161a","sidebarTextActiveBorder":"#175b33","sidebarTextActiveColor":"#ffa900","sidebarHeaderBg":"#175b33","sidebarHeaderTextColor":"#f7f7f7","onlineIndicator":"#21e4a4","awayIndicator":"#eba828","dndIndicator":"#bb2528","mentionBj":"#de161a","mentionColor":"#ffa900","centerChannelBg":"#ffffff","centerChannelColor":"#333333","newMessageSeparator":"#ffa900","linkColor":"#016341","buttonBg":"#1c8c4c","buttonColor":"#ffffff","errorTextColor":"#bb2528","mentionHighlightBg":"#bb2528","mentionHighlightLink":"#ffa900","codeTheme":"solarized-light","mentionBg":"#de161a"}
```

また、Slackで設定しているテーマカラーを、そのまま流用することもできます。

![custome_theme_slack](https://blog.kaakaa.dev/images/posts/mattermost/MattermostThemes/custom_theme_slack.png)

# カスタムテーマ集

## [Mattermost Themes](https://avasconcelos114.github.io/mattermost-themes/)

![gh_mattermost_themes](https://blog.kaakaa.dev/images/posts/mattermost/MattermostThemes/gh_mattermost_themes.png)

Mattermostのコミュニティメンバーが作成したGitHub Pagesサイトです。
現時点では、このサイトで最も多くのテーマが紹介されているようです。(Light Theme 16種類、 Dark Theme 14種類)

このサイトでは、テーマのサンプル画像をクリックするとJSON文字列がコピーされるので、そのままMattermostのカスタムテーマ設定の部分にペーストするだけでテーマを適用することができます。

また、このプロジェクトはGitHub上で公開されているため、自分のテーマを追加することもできます。
https://github.com/avasconcelos114/mattermost-themes

## [公式ドキュメント](https://docs.mattermost.com/help/settings/theme-colors.html#custom-theme-examples)

![official_docs_themes](https://blog.kaakaa.dev/images/posts/mattermost/MattermostThemes/official_docs_themes.png)

公式ドキュメントの **Theme Color** のページにいくつかテーマが紹介されています。現時点(2019/12/14現在)で紹介されているのは、以下の5つのテーマです。

* **GitHub theme**
* **Monokai theme**
* **Solarized Dark theme**
* **Gruvbox Dark theme**
* **One Dark**

このサイトでは、JSON文字列が直接表示されているため、これを選択してコピー、カスタムテーマ設定画面でペーストという流れになります。

## [公式フォーラム](https://forum.mattermost.org/t/share-your-favorite-mattermost-theme-colors/1330)

![forum_themes](https://blog.kaakaa.dev/images/posts/mattermost/MattermostThemes/forum_themes.png)

公式フォーラムでもテーマを共有するトピックが立っています。名前のついていないテーマもありますが、現時点(2019/12/14事典)で、14種類ほどのテーマが紹介されています。(**One Dark**等、他で紹介されているものと重複しているものもありそうですが...）

