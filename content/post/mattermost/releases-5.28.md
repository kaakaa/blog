---
title: "Mattermost5.28の新機能"
date: 2020-10-23T23:50:12+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2020/10/15にMattermost v5.28.0がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

変更内容の詳細については公式のリリースを確認してください。

* [Mattermost Release v5\.28 is now available \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-release-v5-28/)
* [Release v5.28 - Feature Release](https://docs.mattermost.com/administration/changelog.html#release-v5-28-feature-release)

また、`mmctl` コマンドとコンプライアンスエクスポート機能のバグFixが行われた v5.28.1 がリリースされています。
* [Mattermost 5\.28\.1 and 5\.27\.1 released \- Mattermost \- Open Source, On\-Prem or Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-28-1-and-5-27-1-released/)

---

### 新しく導入された機能のお知らせ
Mattermostを新しいバージョンに上げた後にログインすると、新しいバージョンで利用可能なった機能などがMattermost上で通知されるようになりました。(v5.27からv5.28に上げた時は出てこなかったような? v5.26で導入されたチャンネルカテゴリの機能では表示されましたが。）

![in production](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.28/in-production-notice.png)
(画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-28/#notifications)のもの)

この設定は **システムコンソール > 通知 > エンドユーザー通知を有効にする** からオフにすることもできます。  
(システムコンソールのサイドメニュー上に２つ「通知」の項目が存在してしまったため、下の方の設定（今回追加された機能に関する設定）の項目名を「お知らせ」に変更しました。v5.29から反映されると思います。)

![settings in production](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.28/settings-in-production-notice.png)


### その他の新機能
#### bold / italic の入力簡易化
メッセージ入力欄で `Ctrl + I` を入力するとMarkdownのItalic記法が、`Ctrl + b` を入力するとMarkdownのBold記法が挿入されるようになりました。

![shortcut](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.28/shortcut-markdown.gif)

#### Bleve検索でのワイルドカードのサポート
[v5.24でBleveを利用した全文検索機能](https://qiita.com/kaakaa_hoe/items/574972591f6b0b0f642f#bleve%E3%82%92%E5%88%A9%E7%94%A8%E3%81%97%E3%81%9F%E5%85%A8%E6%96%87%E6%A4%9C%E7%B4%A2%E3%81%A7%E6%97%A5%E6%9C%AC%E8%AA%9E%E6%A4%9C%E7%B4%A2%E3%82%82%E5%8F%AF%E8%83%BD%E3%81%AB%E5%AE%9F%E9%A8%93%E7%9A%84%E6%A9%9F%E8%83%BD)がベータ版としてリリースされましたが、Bleveを利用した検索を行う際にワイルドカードを利用することができるようになりました。

#### 検索除外文字のみを指定した場合でも検索結果を返す
[過去にIssueも立てていましたが](https://github.com/mattermost/mattermost-server/issues/14641)、チャンネルと検索**除外**ワードのみを指定した場合に、検索結果が空になってしまう問題が解決しました。チャンネル内の検索除外ワードを含む投稿を除いた投稿が検索結果として返されるようになります。

関連記事: [Mattermostの検索機能についてちょっと調べた · kaakaa blog](https://blog.kaakaa.dev/post/mattermost%E3%81%AE%E6%A4%9C%E7%B4%A2%E6%A9%9F%E8%83%BD%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%A1%E3%82%87%E3%81%A3%E3%81%A8%E8%AA%BF%E3%81%B9%E3%81%9F/)

#### 通知音の切り替え

デスクトップ通知の通知音を変更できるようになりました。

![notification sound](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.28/notification-sound.png)


### (E20) 管理者タスクを委譲するための新たな管理者ロールの導入
[Enterprise E20](https://mattermost.com/pricing/) 限定の機能です。

管理者タスクを委譲しやすくするため、それぞれ細かい粒度の管理者権限を持つ新たな管理者ロール **System Manger**、 **UserManage**、 **Read-only Admin**が導入されました。  
以前のバージョンでは、システム管理者の権限を持つユーザーのみがシステムコンソールにアクセスしたり、管理者向けのコマンドを実行することができました。そのため、一部の管理者権限を他のユーザーに渡したい場合でも、全ての管理者権限を渡す必要がありました。
今回のバージョンから、以下のような管理者タスクを他のユーザーに委譲したい場合に、全ての管理者権限を渡すことなく権限を委譲することができるようになります。

* 管理タスクやユーザー管理の委譲
* プラグイン管理の委譲
* システムコンソールの読み取り専用権限
* AD/LDAPやSAMLなどの認証管理やアクセスコントロールシステムへのアクセス権限
* コンプライアンスレポートやコンプライアンス管理へのアクセス権限

今回新たに追加された管理者ロールが持つ権限について、詳しくは公式ブログや公式ドキュメントを確認してください。
[Granular administrator roles for delegated administration (E20 Edition)](https://mattermost.com/blog/mattermost-release-v5-28/#admin)
[Additional System Admin Roles (E20)](https://docs.mattermost.com/deployment/admin-roles.html)

今回新たに追加された管理者ロールは、システム管理者がコマンドラインツールである`mmctl`を実行することで付与できます。


### (E10) AD/LDAP 証明書ベースの認証
[Enterprise E10](https://mattermost.com/pricing/) 限定の機能です。

AD/LDAPと連携している場合に、クライアント証明書を利用してAD/LDAPディレクトリに対する認証を行うことができるようになりました。この追加の認証メカニズムにより、認証情報の誤用に対する保護を強化できます。

証明書ベースの認証を有効にするには、**システムコンソール > AD/LDAP**から公開証明書と秘密鍵をアップロードしてください。

## その他の変更

### GitHub Plugin v2.0 リリース
Mattermost公式チームが開発を進めている GitHub Plugin v2.0 がリリースされました。
[Github Plugin 2\.0 Release \- Mattermost](https://mattermost.com/blog/github-plugin-2-0-release/)

GitHubプラグインでは、GitHubリポジトリのアクティビティを関連するMattermostへ通知したり、MattermostからIssueを作成するなど、MattermostとGitHubを使っている場合に便利な機能が利用可能になります。

その他、プラグインを含むMattermostとの連携機能については [Mattermost Integrations \- Powerful integrations to help your team do better work, faster\.](https://integrations.mattermost.com/) にまとまっています。

### mmctl の TLS1.0, 1.1 サポートについて
2021/1/16リリース予定の Mattermost v5.31 から、Mattermostの管理者向けコマンドラインツールである`mmctl` を用いてTLS1.0, 1.1 を利用しているMattermostサーバーに対して操作を行おうとすると、エラーを返すようになります。オプションを指定することで利用し続けることも可能ですが、TLS1.0, 1.1は多くのブラウザでdeprecatedとなっているため、TLS1.2以降へのアップグレードを推奨します。

### [破壊的変更](https://docs.mattermost.com/administration/changelog.html#breaking-changes)
* Mattermostがクラッシュした時に coredump が生成されるようになりました

# おわりに

次の`v5.29`のリリースは2020/11/16(Mon)を予定しています。
そして機能追加が行われる`v5.30`は恐らく2020/12/16(Wed)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。
