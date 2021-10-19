---
title: "Mattermost 6.0の新機能"
emoji: "🎉"
published: true
date: 2021-10-19T22:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/10/13 に Mattermost のメジャーバージョンアップである `v6.0.0` がリリースされました。  
また、10/18にMedium Level Security Fixを含む `v6.0.1` がリリースされています。今後、v6へアップグレードする場合は `v6.0.1` を使用してください。

本バージョンから開発におけるコラボレーションを加速させる**Playbook**と**Boards**という機能が追加されています。今回のメジャーバージョンアップは、MattermostがSlack alternativeなチャットツールから、開発におけるコラボレーションを扱うプラットフォームへと舵を切る分岐点となるリリースとなりそうです。  
Mattermost 6.0リリースの背景については、以下のMattermost CEOによるエントリで語られています。
[Introducing Mattermost 6\.0: A New Approach to Developer Collaboration](https://mattermost.com/blog/mattermost-6-0-future-of-developer-collaboration/)

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v6\.0 is Now Available \- Mattermost](https://mattermost.com/blog/mattermost-v6-0-is-now-available/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v6-0-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/install/self-managed-changelog.html#important-upgrade-notes)

今回のアップグレードではデータベーススキーマの変更が行われており、この処理に(特にMySQLを使用している場合は)時間がかかることが予想されます。データベースのサイズによっては、長い時間Postテーブルへのロックがかかり、メッセージの投稿や受信に影響が出る場合があります。Mattermost上の投稿が1000万件以上ある環境や、クラスタ構成を採用している環境では特に注意が必要です。アップグレード作業時の注意点については以下のエントリにまとめられています。

[How to Upgrade to Mattermost v6\.0](https://mattermost.com/blog/upgrade-to-mattermost-6-0/)

また、アップグレード時に実施されるスキーマ変更の内容と、1000万/7000万件の投稿を持つデータベースに対するアップグレード処理にかかる時間については、以下のドキュメントに参考情報がまとめられています。

* https://gist.github.com/streamer45/59b3582118913d4fc5e8ff81ea78b055
* https://gist.github.com/streamer45/868c451164f6e8069d8b398685a31b6e

クラスタ構成を組んでいる場合、v6の一部データ形式が以前のバージョンと非互換となっている箇所があるため、クラスタ内でv5系とv6系を混在させることができなくなっています。クラスタ構成を組んでいる場合、またはMattermost Kubernetes operatorを利用している場合は[ChangelogのImportant Upgrade Notes](https://docs.mattermost.com/install/self-managed-changelog.html#important-upgrade-notes)を参照してください。

v6.0へのアップグレードのリスクを感じる場合、Mattermost 5.37がESR(Extended Support Release)として2022/4月までサポートされる予定のため、このリリースを使い続けるのが良さそうです。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [`Professional`](https://mattermost.com/pricing/)
- [`Enterprise`](https://mattermost.com/pricing/)

見出しの前に何もない場合、Starter(無償版)でも利用可能な機能です。

## プラットフォームの再設計と新たなサブスクリプションプラン

`v6.0`からはチャットベースのメッセージングだけでなく、コラボレーションプラットフォームとしての価値を高めていく方にシフトするようです。
それに伴い、以前に`Incident Collaboration`と呼ばれていたプラグインが`Playbook`として、`Focalboard`と呼ばれていた機能が`Boards`として、それぞれMattermostプラットフォームに統合され、デフォルトで利用できるようになりました。

Mattermostプラットフォームの全体像については、以下の動画で紹介されています。
[Mattermost Platform Overview \- Mattermost](https://mattermost.wistia.com/medias/lr85qzv51e)

また、v6.0リリースに合わせてサブスクリプションプランの内容が変更されています。  
https://mattermost.com/pricing/

今までは、無償のOSSセルフホスティング版はTeam Edition、有償プランは`E10`/`E20`と呼ばれていましたが、今後は無償版を`Starter`、`E10`に当たるプランを`Professional`、`E20`に当たるプランを`Enterprise`と呼ぶよう変更され多様です。また、以前は[Mattermost Cloud](https://mattermost.com/platform-overview/)独自のサブスクリプションプランがありましたが、本バージョンからセルフホスティング版のプランと統合され、それぞれのプランでセルフホスト/Cloudを選択できるようになったようです。(`Starter`プランでCloudを選択した場合のみ(?)、年額$149の追加費用が必要)

また、メジャーバージョンアップにあたり、いくつかの機能が下位のプランでも利用可能になっています。詳細については以下のChangelogを参照してください。
https://docs.mattermost.com/install/self-managed-changelog.html#packaging-changes

## グローバルナビゲージョンの追加
Mattermostに関する各種メニューを表示するためのグローバルナビゲーションヘッダーが新たに追加されました。グローバルナビゲーションメニューから、本バージョンで追加された**Playbook**や**Boards**へ簡単に移動することができます。
アカウント設定などのメニューは、今までは画面左上からアクセスできましたが、本バージョンからは画面右上に移動しているので注意が必要です。

![Global Navigation](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/Global-Header-Feature-Animation.gif)
(画像は[公式ブログ](https://mattermost.com/blog/mattermost-v6-0-is-now-available/)より)

## ベータ版機能のGA化

以前のバージョンでベータ版として含まれていた以下の機能が正式版(GA, Generally Available)となりました。(機能名末尾の文字列は、当該機能が利用できるプランを示しています)

* [アーカイブされたチャンネルの閲覧](https://docs.mattermost.com/configure/configuration-settings.html#allow-users-to-view-archived-channels)
* [コンプライアンスエクスポート](https://docs.mattermost.com/comply/compliance-export.html) (`Enterprise`)
* [カスタム利用規約](https://docs.mattermost.com/comply/custom-terms-of-service.html) (`Enterprise`)
* [ゲストアカウント](https://docs.mattermost.com/onboard/guest-accounts.html) (`Enterprise`, `Professional`)
* [`mmctl` コマンド](https://docs.mattermost.com/manage/mmctl-command-line-tool.html)
* [追加のシステム管理者ロール](https://docs.mattermost.com/onboard/system-admin-roles.html) (`Enterprise`)
* [プラグイン機能](https://developers.mattermost.com/integrate/admin-guide/admin-plugins-beta/)
* [タイムゾーン設定](https://docs.mattermost.com/messaging/managing-account-settings.html?highlight=timezone#timezone)

## メッセージリンクのプレビュー表示
Mattermost内の投稿のリンクを同じMattermost内に投稿した際、投稿のプレビューが自動で表示されるようになりました。プレビューは、リンク先の投稿があるチャンネルに参加している場合のみ表示されます。  
[Sharing Messages](https://docs.mattermost.com/messaging/sharing-messages.html?highlight=preview)

![Preview Post](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/preview_post.png)

プレビューが表示されるのは、メッセージのコンテキストメニューから **リンクをコピーする** を選択した時にコピーされるURL形式 (`${SITE_URL}/${TEAM_NAME}/pl/${MESSAGE_ID}`) の場合のようです。投稿日付の部分をクリックした時も投稿へのリンクを取得することができますが、この時のURL形式は`${SITE_URL}/${TEAM_NAME}/channels/${CHANNEL_NAME}/${MESSAGE_ID}` となり、この形式のリンクではプレビューが表示されませんでした。

![Copy Link](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/copy_link.png)

## Playbook
以前、Incident Collaboration Pluginとして開発されていた機能が **Playbook** としてMattermostプラットフォームに統合されました。

* https://docs.mattermost.com/guides/playbooks.html
* https://github.com/mattermost/mattermost-plugin-playbooks

**Playbook** は、繰り返し行われるプロセスのワークフロー/チェックリストをPlaybookとして事前に定義し、作業が発生するたびにPlaybookを元にチャンネルを作成することで、作業ごとの記録をチャンネルにまとめておくことのできる機能です。

リリース直後(約1年前)ぐらいに触った際の記録を以下に書いています。
[Mattermost5.30の新機能 - (E20) Mattermostでインシデント管理](https://blog.kaakaa.dev/post/mattermost/releases-5.30/#e20-mattermost%E3%81%A7%E3%82%A4%E3%83%B3%E3%82%B7%E3%83%87%E3%83%B3%E3%83%88%E7%AE%A1%E7%90%86)

その後、数多くの機能追加が行われているため、現在の機能について詳しくは公式ドキュメントを参照しください。また、Mattermost v6.0公開にあたり、更新通知の再設計やリマインダー設定機能の追加などの機能追加が行われています。
[Mattermost Playbooks](https://docs.mattermost.com/guides/playbooks.html)

## Boards

以前、Focalboardとして開発されていた機能が **Boards** としてMattermostプラットフォームに統合されました。

* https://docs.mattermost.com/guides/boards.html
* https://www.focalboard.com/

Focalboardは、Mattermost社が開発しているTrello・Notion・Asana Alternativeを謳うKanbanツールです。Focalboardについては、以下の記事で紹介しています。  
[NotionのようなUIのTrelloっぽいKanbanツールのOSSの Focalboard を触ってみた](https://zenn.dev/kaakaa/articles/mattermost-focalboard-first)

**Boards**はチャンネルごとにボードを管理することできますが、Mattermostで**Boards**を開くとチャンネルごとにBoardがどの程度利用されているか俯瞰できるようになっています。

![Board Overview](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/board_overview.png)

その他、v6リリースに伴うアップデートとして、カラムにsum/rangeなどの解析関数を設定可能になったり、日付プロパティに期間を指定する機能が追加されています。  
また、shared boardsの機能により、誰でもBoardにアクセスできるリンクを生成できるようになりました。Boardはチャンネルごとに作成されるため、参加していないチャンネルに紐づくBoardや、そもそもMattermostにアカウントの無いユーザーはBoardにアクセスできませんが、そのようなユーザーでもBoardを閲覧できるようになります。

![Board Share](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/board_share.png)


## Desktop App v5.0

デスクトップアプリもv5系へメジャーバージョンアップされました。
[Desktop Application Changelog](https://docs.mattermost.com/install/desktop-app-changelog.html#id1)

Mattermost Desktop App v5.0はMattermost v5系もサポートしていますが、Minimum Server Versionが`5.37`となっているので注意が必要です。

#### グローバルナビゲーションのサポート
デスクトップアプリもグローバルナビゲーションに対応しており、Channels/Playbooks/Boardsを素早く切り替えることができます。
今までのバージョンとは違い、タブではChannels/Playbooks/Boardsを表示し、サーバーの切り替えはドロップダウンメニューから行うようになっているようです。

![Desktop Nav](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/desktop_nav.png)

`Ctrl + Shift + [数字キー]` のショートカットででサーバーを切り替えることができるようになっているので、多くのサーバーを登録していた場合は、このショートカットを覚えておいた方が良さそうです。また、サーバーの登録順はドラッグ＆ドロップで変更できます。（自分はMattermostだけでなくSlackなどもサーバー登録していたので、このショートカットは必須になりそうです）

#### スペルチェック言語の多言語指定
スペルチェック機能の対象言語として複数の言語を指定できるようになりました。しかし、残念ながら日本語の設定は無いようです。

![Desktop SpellCheck](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/desktop_spellcheck.png)

## Mobile App v1.47.0

モバイルアプリの方は、マイナーバージョンアップのようです。  
[Mattermost Mobile Apps Changelog](https://docs.mattermost.com/deploy/mobile-app-changelog.html#)

こちらもMattermost v5.37以上のサポートになっているようなので、注意が必要です。また、自己署名証明書を使用しているサーバーへ接続する場合は、デバイスへ証明書をインストールしなくてはならなくなったようです。

> Server Versions Supported: Server v5.37.0+ is required. Self-Signed SSL Certificates are not supported unless the user installs the CA certificate on their device.

## 破壊的変更

### `mattermost`コマンド

`mattermost`コマンドの多くのサブコマンドが、`mmctl`へ移行されたため、Mattermostインスタンス管理用のCLIツールのデフォルトは`mmctl`になったようです。  
[Command Line Tools](https://docs.mattermost.com/manage/command-line-tools.html)

ただし、`import/export`サブコマンドと、それに関連するサブコマンドは`mattermost`コマンドに残っています。これは、ユーザー権限に基づいてMattermostインスタンスを操作する`mmctl`コマンドでは、インスタンス移行のためのコマンドを実行するには権限が足りないためだと思われます。  
インスタンスの移行等を検討する場合でない限りは、`mmctl`コマンドを使用することをお勧めします。

### SlackデータをMattermost上からアップロードする機能の廃止
以前のバージョンでは、SlackからエクスポートしたデータをMattermost画面上からアップロードすることができましたが、この機能が廃止されました。
[Migration Guide](https://docs.mattermost.com/onboard/migrating-to-mattermost.html#migrating-from-slack-using-the-mattermost-web-app)

代わりに[`mmetl`](https://github.com/mattermost/mmetl)ツールと`mmctl import`を使用して、コマンドラインからインポートする方法が推奨されるようになったようです。
[Migration Guide](https://docs.mattermost.com/onboard/migrating-to-mattermost.html#migrating-from-slack-using-the-mattermost-mmetl-tool-and-bulk-import)

### その他の破壊的変更

Mattermost v6リリースに伴い、以下のバージョンのサポートが廃止されました。
* MySQL 5.7.12以前のバージョン
* Elasticsearch 7以前のバージョン
* Windows 7

また、Mattermostの設定項目の内、いくつかの項目が削除されています。詳しくは以下のドキュメントを参照ください。
https://docs.mattermost.com/install/self-managed-changelog.html#deprecations


---

## その他の変更

### デフォルトテーマの変更

デフォルトのテーマが `Denim` に変更され、シックな印象となりました。
公式サイトも`Denim`をテーマカラーとして刷新されています。
https://mattermost.com/

その他のテーマも、名称といくつかのカラーが変更されているため、Mattermostサーバーをv6.0にアップグレードした後は、少し見た目が変わっている場合があります。
既存のテーマカラーに戻したい場合は、カラーテーマが公式ドキュメントで公開されているので、カスタムテーマから設定してください。
https://docs.mattermost.com/messaging/customizing-theme-colors.html#custom-theme-examples

## その他のトピック

### Hacktoberfest

[Join Mattermost for Hacktoberfest 2021](https://mattermost.com/blog/hacktoberfest-2021/)

今年もMattermostはHacktoberfestのPartnerとして参加しています。 
https://hacktoberfest.digitalocean.com/

Hacktoberfestの景品とは別に、Mattermost独自のSwagを用意しています。
10月中に一つでもコントリビュートを行なった場合は、限定ステッカーがもらえるようです。このステッカーは、日本のJapanese Maple Treeをモチーフにしているそうです。

![Fall Sticker](https://blog.kaakaa.dev/images/posts/mattermost/releases-6.0/fall_sticker.png)

さらに、各カテゴリ(**Focalboard**, **Mattermost App**, **Writing Program**, **Translation**, **QA Testing**)でトップコントリビューターに選ばれた場合、オリジナルのメカニカルキーボードがプレゼントされるようです。

## おわりに

次の`v6.1`のリリースは 2021/11/16(Tue)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
