---
title: "Mattermost 5.35の新機能"
date: 2021-05-18T07:30:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

(最近Qiitaでの体験があまり良くないので、来月の記事からZennに移行しようと思っています。既存の記事も移行し、Qiitaの記事内容をZennの記事へのリンクに更新する予定です。https://zenn.dev/kaakaa)

2021/05/14 に Mattermost v5.35.0 がリリースされました。 

v5.35へのマイグレーション時に、MySQL を使用している場合にデータベースに対する`readTimeout`時間が短すぎることが原因で、MySQL へのコネクションが正常に貼れなくなるという問題が確認されています。アップグレードを行う場合は、まもなくリリース予定のv5.35.1の利用をおすすめします。

https://docs.mattermost.com/administration/changelog.html#release-v5-35-feature-release
> v5.35.1, release day TBD
> Fixing an issue where 5.35.0 migration is failing on MySQL installations with an “invalid connection” error due to an issue with the readTimeout parameter in SqlSettings.DataSource (default is 30 seconds). To mitigate this, readTimeout can be either removed or increased to a high enough value to allow the migration to happen. MM-35767

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。　　
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.35 is now available \- Upgrade today to try the new features](https://mattermost.com/blog/mattermost-release-v5-35/)
- [Mattermost Changelog](https://docs.mattermost.com/administration/changelog.html#release-v5-35-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/administration/changelog.html#important-upgrade-notes)

v5.35へのアップグレードに関連するいくつかの注意事項があります。

* Mattermost v5.35では、今後リリース予定の共有チャンネル機能や、返信スレッドの折り畳み機能に必要なデータベースのマイグレーションが行われており、**アップグレードに時間がかかる**ことがあります。データベースのサイズにもよりますが、**数分から、MySQL 5.x系を利用している場合は数時間かかる**こともあります。パフォーマンスへの影響については[5\.35 Migration](https://gist.github.com/streamer45/9aee4906639a49ebde68b2f3c0f924c1)も確認ください。
* v5.35から導入されたファイル検索機能は、デフォルトではすでにアップロード済みのファイルに対してはファイル名のみの検索しか動作しません。**アップロード済みのファイルの内容で検索を行いたい場合、[`mattermost extract-documents-content`](https://docs.mattermost.com/administration/command-line-tools.html#mattermost-extract-documents-content)コマンドを実行**する必要があります。また、**検索に`Elasticsearch`もしくは`Bleve`を利用している**場合は、`mattermost extract-documents-content`が完了した後で、**検索インデックスの再構築**を行う必要があります。ファイル検索機能については、[Search for files and document contents in Mattermost](https://mattermost.com/blog/file-search/)でも説明されています。
* **Bulk Import**により追加されたユーザーに設定されるパスワードは、**比較的脆弱なパスワード**となっていたため、Bulk Import後に一度もパスワードを変更していないユーザーがいた場合、そのユーザーのパスワードをリセットすることをおすすめします。
* v5.38にて、`config.json`ファイルを監視し、変更を自動でリロードする `config watcher`**の機能が廃止**されます。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。


## ファイル/ドキュメント内容検索

Mattermostにアップロードされたファイルの内容を検索できるようになりました。

![search_file](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/search_file.gif)


検索対象となるファイルの形式は、`.pdf`, `.pptx`, `.odt`, `.html` と plain textドキュメントです。
[Cloud版](https://mattermost.com/mattermost-cloud/)では、さらに`.docx` も検索対象になります。
Team Edition(OSS版)でも、下記の依存ライブラリをインストールすることで、`.docx`, `.rtf`, `.pages`ファイルを検索対象とすることもできるようです。(Linuxのみ？)

```
$ sudo apt-get install poppler-utils wv unrtf tidy
$ go get github.com/JalfResi/justext
```
[sajari/docconv: Converts PDF, DOC, DOCX, XML, HTML, RTF, etc to plain text](https://github.com/sajari/docconv#dependencies)

また、Team Edition(OSS版)においてMattermsotサーバーの設定ファイル`config.json`の`FileSettings.ArchiveRecursion`を`true`に設定することで、ZIPファイルの内容も検索できるようになるようです。

記事冒頭の`Import Upgrade Notes`にもあるように、ファイル内容についてのインデックスが作成されるまではファイル名での検索しか動作しません。v5.35にアップデートする前にアップロード済みのファイルの内容に対するインデックスを作成するには [`mattermost extract-documents-content`](https://docs.mattermost.com/administration/command-line-tools.html#mattermost-extract-documents-content)コマンドを実行する必要があります。

ファイル検索について、詳細は以下の公式ブログを参照ください。
[Search for files and document contents in Mattermost](https://mattermost.com/blog/file-search/)

## (E20/Cloud) Incident Collaborationの改善

### Ad hoc tasks

インシデント対応時のタスクは、事前にPlaybookに定義されているものしか使えませんでしたが、新しいバージョンではインシデント対応中にアドホックでタスクの追加、編集、削除ができるようになりました。  

![incident_adhoc](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_adhoc.webp)
(※画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-35/)から)

### Stakeholder overview

Incident Collaborationサイドバーの下部に表示されている `Overview` ボタンを押すことで、インシデント対応に当たっているユーザー等を確認できるようになっているようです。

![incident_overview1](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_overview1.png)
![incident_overview2](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_overview2.png)

### 詳細なアクセス管理

インシデントに対するアクセス権限をMattermostチーム単位?に設定できるようになりました。Mattermostインスタンスにユーザー用のチームと管理者用のチームに分けていた場合、管理者側のチームの人だけがインシデントにアクセスできるようにするなどの対応を取れるようになったのだと思います。
![incident_access_control](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_access_control.png)

また、Incident Collaborationの設定の中に `Experimental Features`に関する設定も見えるようです。これを有効にすると、インシデント対応に関するメトリクス一覧の画面を表示できるようになったりするらしいです。

![incident_stats](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_stats.png)

### インシデント開始に係る処理の自動化

[先月のリリース時点と比較すると](https://blog.kaakaa.dev/post/mattermost/releases-5.34/#cloude20-incident-collaboration-%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%B7%E3%83%87%E3%83%B3%E3%83%88%E9%96%8B%E5%A7%8B%E6%99%82%E6%93%8D%E4%BD%9C%E3%81%AE%E8%87%AA%E5%8B%95%E5%8C%96)、Playbookの設定に`Announce in another channel`と`Send a webhook`という項目が追加されているようです。
この設定により、(おそらく)Incidentに関するイベントを他のチャンネルに通知したり、Webhook経由で別のサーバーへ通知して他サービスと連携するなどができるようになっているようです。

![incident_settings](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/incident_settings.png)


## (E20/Cloud) システムコンソールページに対する詳細な権限設定

システムコンソールの各設定ページに対する編集/閲覧権限に関する設定項目が細分化されました。
新たに **Experimental**(実験的な機能), **About**(エディションとライセンス),**Environment**(環境),**Site Configuration**(サイト設定),**Authemtication**(認可),**Integration**(統合機能),**Compliance**(コンプライアンス)のセクションが追加され、それぞれのサブセクションごとの権限設定が可能になっています。

![granular_access](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/granular_access.png)

## (E20/Cloud) 共有チャンネル

Mattermostインスタンス間でチャンネルを共有できる機能が実験的な機能として追加されました。Enterprise E20ライセンスが必要です。また、共有チャンネル機能はデフォルトでは無効化されています。

**システムコンソール > 実験的機能 > 機能** から有効にでき、スラッシュコマンド `/share` で共有チャンネルを管理するようです。

![shared_channel_setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/shared_channel_setting.png)

共有するインスタンスを持っていないので実際の動作は確かめられませんでしたが、共有チャンネルは以下のような見た目になるそうです。(※画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-35/)から)

![Shared-channels](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/Shared-channels.webp)

## Apps Framework (Developer Preview)

Mattermostの統合機能の新たな形式として、Mattermost Appsが利用できるようになります。(現時点ではDeveloper Previewのため、プロダクション環境では利用できません。)

Mattermost Appsはどんな言語でも記述することができ、モバイルアプリやデスクトップアプリでも動作します。

[公式リリースブログ](https://mattermost.com/blog/mattermost-release-v5-35/)では、AWS LambdaによるAppsのサーバーレスホスティング、ServiceNowとの連携、Zendeskとの連携、について紹介されています。

- [mattermost/mattermost\-app\-servicenow: Service Now app for Mattermost](https://github.com/mattermost/mattermost-app-servicenow#configuration)
- [mattermost/mattermost\-app\-zendesk: Zendesk App for Mattermost](https://github.com/mattermost/mattermost-app-zendesk#setting-up)

Mattermost Appsについて詳しくは、公式ドキュメント[Apps \(Developers Preview\)](https://developers.mattermost.com/integrate/apps/)を参照してください。

また、日本語で書かれた記事もいくつかあります。

- [Mattermost Apps Framework をJava \(JAX\-RS\)で試してみた – maruTA\(Bis5\)'s Weblog – Side T:echnology](https://tech.bis5.net/2021/05/09/248.html)
- [Mattermost Apps Frameworkを触ってみた](https://zenn.dev/kaakaa/articles/mattermost-apps-sample)


## その他の機能

### (E20) Entrepriseトライアルの改善

Enterprise E20トライアルの開始時、終了３日前、終了日にシステム管理者に対してバナー警告が表示されるようになりました。

### 履歴からカスタムステータスを設定

最近設定したカスタムステータスを選択することができるようになりました。

![recent_custom_status](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/recent_custom_status.png)

### リンクプレビュー表示を無効化するドメインの設定

URLを含むメッセージが投稿された場合、MattermostはリンクのOGPプレビューを表示しようとしますが、プレビューの表示をドメインによってに抑制する設定ができるようになりました。特定のサイトのプレビューを表示したくない場合などに利用できます。

![restrict_preview](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/restrict_preview.png)

### サイドバーにプロフィール画像が表示されるように

チャンネルサイドバーのダイレクトメッセージのセクションに、DM相手のプロフィール画像が表示されるようになりました。

![dm_sidebar_icon](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.35/dm_sidebar_icon.png)

## その他のトピック

### Mattermost-Focalboard 連携

Mattermost が開発している `a self-hosted alternative to Trello, Notion, and Asana` な[Focalboard](https://www.focalboard.com/)とMattermostとの連携機能の開発が進められています。

[Focalboard\-Mattermost Integration · Discussion \#118 · mattermost/focalboard](https://github.com/mattermost/focalboard/discussions/118)

まだEarly Preview段階であり、かつ誰でも利用できるというわけではなく、Mattermostコミュニティサーバーの`Focalboard`チャンネルからしか利用できませんが、以下のGitHub Discussionsで意見を募集しています。

[Mattermost\-Focalboard \- Early Preview · Discussion \#349 · mattermost/focalboard](https://github.com/mattermost/focalboard/discussions/349)

Focalboardに関するYouTube動画も公開されています。

[Focalboard \- YouTube](https://www.youtube.com/watch?v=v6hG91_WvhY)
[Focalboard \- YouTube](https://www.youtube.com/watch?v=eO63Owne-XI)

また、Focalboard の開発者がMattermost社のPodcastである [What Matters](https://mattermost.libsyn.com/)に出演しています。

[What Matters, Episode 18: Meet the Focalboard Team](https://mattermost.com/blog/what-matters-episode-18-meet-the-focalboard-team-with-chen-i-lim-and-jesus-espino/)

### E2E Cypress Test Automation Hackfest
5月中、MattermostではCypressでE2Eのテストを記述するHackfestを開催しているようです。Top 3になったコントリビューターはSpecial Awardを受け取れるようです。

[E2E Test Automation Hackfest 2021 🚀 · Issue \#17555 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/17555)
[Join the 2021 Mattermost E2E Cypress Test Automation Hackfest\!](https://mattermost.com/blog/mattermost-e2e-cypress-test-automation-hackfest-2021/)

## おわりに

次の`v5.36`のリリースは 2021/06/16(Wed)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
