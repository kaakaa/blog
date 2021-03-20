---
title: "Mattermost5.26の新機能"
date: 2020-08-16T00:06:12+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/8/14にMattermost v5.26.0がリリースされました。  

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost Release v5\.26 is now available \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-release-v5-26/)
* [Mattermost Changelog — Mattermost 5\.26 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-26-feature-release)

---

## (実験的な機能) チャンネルのカテゴリ化と順序の入れ替え

Mattermostにもチャンネルのカテゴリ化の機能が追加されました！ 🎉

Mattermostでは最近チャンネルサイドバーの改善が行われていますが、ついにチャンネルをカテゴリ毎にまとめる機能が（実験的な機能としてではありますが）使えるようになりました。

### 設定方法

**システムコンソール > 実験的機能 > 機能 > 実験的なサイドバー機能** を `有効化` に設定する必要があります。**実験的なサイドバー機能**はユーザー毎にOn/Offの切り替えが可能であり、システムコンソールで`有効化 (デフォルト On)` に設定するとサーバー上の全てのユーザーに対して実験的なサイドバー機能を `On` にすることができ、`有効化 (デフォルt Off)` にするとユーザー自身で有効化しない限り **実験的なサイドバー機能** は Onになりません。

![settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_settings.png)

システムコンソールで **実験的なサイドバー機能** が `有効化 (デフォルト On)` もしくは `有効化 (デフォルト Off)` に設定されている場合、ユーザーは **アカウント設定 > サイドバー > 実験的なサイドバー機能** から、機能のOn/Offを切り替えることができます。

![account settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_account_settings.png)


機能を有効にすると、Mattermostの画面にカテゴリの作成方法に関するポップアップが表示されます。

![popup](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_popup.png)

### カテゴリ作成

**実験的なサイドバー機能**を有効にしてから、サイドバーの `チャンネル`セクションの横にある `...` をクリックし、`新しいカテゴリを作成する`をクリックすることでカテゴリを作成することができます。

![create](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_create.png)

カテゴリを作成したら、サイドバーにある各チャンネルをドラッグ＆ドロップすることで、カテゴリにチャンネルを追加していくことができます。（カテゴリに追加するだけでなく、ドラッグ＆ドロップ操作でチャンネルの順序を入れ替えることもできるようになっています）

![activity](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_activity.gif)
（画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-26/#sidebar)から）


カテゴリ名の部分をクリックすることで、カテゴリに追加されたチャンネルの表示・非表示を切り替えることができ、チャンネルを非表示にした状態でも未読の投稿があるとチャンネル名が表示されるようです。（未読であってもチャンネル名を表示したくない場合もある気がしますが、その場合は優先度の低いチャンネルのカテゴリを作って、そのカテゴリを一番下に持っていくなどの対応が必要そうです）

その他、チャンネル名の横にあるメニューからも、様々な操作が可能になっているようです。

![collapse](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/category_collapse.gif)


### ダイレクトメッセージの表示順変更

また、ダイレクトメッセージチャンネルの表示順を「新しい順」と「アルファベット順」で切り替えることができるようになっています。

![dm order](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/dm_order.gif)

### チャンネルサイドバー改善作業の今後

チャンネルサイドバー改善については公式ブログの下記エントリにまとめられています。
https://mattermost.com/blog/upcoming-channel-sidebar-features/

今後の改善としては、2020年中にカテゴリのミュート機能とダイレクトメッセージの新しい順でのソート機能（今回追加されたもの？）が予定されており、さらに2020年中に実験的な機能から正式な機能への格上げが予定されているようです。

> * Phase 1 (v5.22): Unreads filtering, collapsible favorites, public, private, and direct message categories
> * Phase 2 (v5.26): User-defined custom categories with drag & drop of channels and categories, quick access menus
> * Phase 3 (Q4 2020): Muting categories, direct message recency sorting
> * Phase 4 (Q4 2020): General availability (i.e., remove the experimental config setting and release these features to all users) 



## Enterprise版の改善

（Enterprise版の機能は動作確認が取れる環境がないので伝聞調の書き方になります...）

### (E20) チャンネルのアーカイブ・復元がシステムコンソールから可能に

チャンネルのアーカイブ、アーカイブされたチャンネルの復元が、システムコンソールから実行可能になったようです。

![channel configuration](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/channel-configuration.png)
![channge configuration2](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/channel-configuration-2.png)

（画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-26/#archive)から）


### (E20) システムコンソールでのメンバー・チャンネルの

チームメンバ管理、チャンネルメンバ管理ページで、ユーザーを役割(ゲスト、システム管理者、等)でフィルタリング可能になったようです。
![filter](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/filters.png)

また、チャンネル管理画面では、チャンネルのタイプ(公開チャンネル、非公開チャンネル、AD/LDAPグループと同期したチャンネル、等)でフィルタリング可能になったようです。
![filter2](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/filters-2.png)

（画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-26/#filters)から）

### (E20) ログ設定のカスタマイズ

Mattermostが出力するログについて、出力先をコンソール、ファイル、syslog、 TCPから選択できるようになり、さらにそれぞれの出力先に対して異なるログレベルを設定するなど柔軟なログ出力の設定が可能になったようです。

設定方法については公式ドキュメントを参照ください。
https://docs.mattermost.com/administration/config-settings.html#output-logs-to-multiple-targets

### (E20・ベータ版) モバイルアプリからもAD/LDAPグループへのメンションが利用可能に

モバイルアプリからも、メンションするAD/LDAPグループの補完、グループメンバーメンションのハイライト、5人以上が属するグループへメンションする際の警告ダイアログが利用可能になったようです。

![groups](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/auto-suggest-groups.png)


## その他の変更

### `Ask the community` リンクの追加
Mattermostのヘッダ部分にヘルプアイコンが追加され、そこからMattermostチームが利用している公式インスタンスの `Peer-to-peer Help` チャンネルへ移動できるようになりました。

![ask](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.26/ask_community.png)

(画像中に翻訳ミスの部分がありますが修正しておきます...)

`Peer-to-peer Help`チャンネルでは、やりとりは基本的に英語ですが、開発チームも見ているチャンネルのため素早い回答が期待できるのではないかと思います。

また、Mattermostは質問用のフォーラムも運用してしおり、こちらからも質問を投稿することができます。
https://forum.mattermost.org/

上記フォーラムも基本的に英語でやりとりが行われていますが、一応、日本語のフォーラムもあります。（あまり利用されてはいないですが...）
https://forum.mattermost.org/c/international/japanese/29


### PostgreSQL 10のサポート
今までサポート対象だったPostgreSQL 9.4のLTSが2020年2月に終了したことを受け、Mattermostも5.26から公式にPostgreSQL 10をサポートするようになりました。今後もMattermost v5系はPostgreSQL 9.4でも動作可能となりますが、Mattermost v6.0(リリース時期未定)ではPostgreSQL 9.4はサポート対象外となる予定です。

### 破壊的変更
* Mattermost v5.26では、Elasticsearchのインデックスを再作成する必要があるそうです
* Gossipプロトコルを使用したクラスター間の通信を暗号化するオプションが追加されたそうです

# おわりに

次の`v5.27`のリリースは2020/9/16(Wed)を予定しています。
そして機能追加が行われる`v5.28`は恐らく2020/10/16(Fri)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
