---
title: "Mattermost 5.37の新機能"
emoji: "🎉"
published: true
date: 2021-07-17T13:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/07/16 に Mattermost v5.37.0 がリリースされました。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.37 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v5-37/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v5-37-extended-support-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/administration/changelog.html#important-upgrade-notes)

* 本バージョンより、返信機能の折りたたみ機能が利用可能になりましたが、まだベータ版のためバグが残っている可能性があり、[既知の問題](https://docs.mattermost.com/messaging/organizing-conversations.html#known-issues)等を把握した上で利用することが推奨されています
* 本バージョンより、Emoji v13.0に基づく絵文字が利用できるようになりますが、新しく追加される絵文字と同名のカスタム絵文字がすでに登録されていた場合、アップデートにより既存のカスタム絵文字の内容が上書きされます
* Incident Collaborationプラグインを全エディションで利用可能とするために、最新のIncident Collaborationプラグインの最低動作バージョンが v5.37 に引き上げられています。Incident Collaborationプラグインをアップデートする際は注意してください。
* 長期サポートバージョンであるMattermost v5.31は2021/10/15にサポート終了となります。v5.31を利用しているユーザーは、v5.37以降のバージョンへアップグレードする必要があります。

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## Collapsed Reply Threads (ベータ版)

[Get early access to the Mattermost Collapsed Reply Threads beta](https://mattermost.com/blog/collapsed-reply-threads-beta/)

Mattermostの[Feature Requestフォーラムでも最も人気のあった](https://mattermost.uservoice.com/forums/306457-general/suggestions/19572469-make-threads-collapsible)返信スレッドの折りたたみ機能(Collapsed Reply Threads)がベータ版として利用できるようになりました。  
本機能を有効にすることで、投稿への返信がチャンネル内に表示されなくなり、チャンネル内の会話をトピックごとに把握しやすくなります。

実際の動作の様子は、[公式ドキュメント](https://docs.mattermost.com/messaging/organizing-conversations.html)で確認できます。

![collapsed-threads-demo](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/collapsed-threads-demo.webp)
(上記画像は[公式ブログ](https://mattermost.com/blog/mattermost-v5-37/#collapsed)から)

---

本機能を有効にするには、まず、**システムコンソール > 実験的機能 > 機能 > 返信スレッドの折りたたみ** が **有効(デフォルトOff)** となっている必要があります。

![collapsed-threads-system-console](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/collapsed-threads-system-console.png)

上記設定を有効にした上で、**アカウント設定 > 表示 > 返信スレッドの折りたたみ (ベータ版)** を **On** に設定することでアカウントごとに機能が有効になります。

![collapsed-threads-account-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/collapsed-threads-account-setting.png)

本機能はまだベータ版の段階のため、既知の問題を含め、問題が発生する可能性を考慮して利用した方が良さそうです。

https://docs.mattermost.com/messaging/organizing-conversations.html#known-issues

## Incident Collaborationの改善

今月も[Mattermost Incident Collaboration Plugin](https://github.com/mattermost/mattermost-plugin-incident-collaboration)の改善があります。

### 全エディションで利用可能に

今までMattermost CloudとEnterprise E20プランでしか利用できなかったIncident CollaborationがOSS版のTeam Edition(Enterprise E0)を含む全エディションで利用可能になりました。
Team EditionではPlaybookが1つしか作成できないなど制限はありますが、機能の使用感の確認などはできるようになります。

https://mattermost.com/pricing-self-managed/

![incident-all-edition](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-all-edition.png)

Incident Collaborationプラグインは、**メインメニュー > マーケットプレース** から簡単にインストールできます。

![incident-install](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-install.png)

インストールが完了すると、上記画像中の`Install`ボタンが`Configure`ボタンに変わるので、`Configure`ボタンをクリックして設定画面へ移動し、**プラグインを有効にする** を **有効** にすることでプラグインが利用可能になります。

![incident-system-console](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-system-console.png)

プラグインを有効にすると **メインメニュー > Incident Collaboration**からIncident Collaborationプラグインを利用することができるようになります。

### プレイブックキーワードの監視

プレイブックに関するキーワードを設定することができるようになり、Mattermost上で設定されたキーワードを含むメッセージが投稿された際にプレイブックの開始を促すメッセージを自動で表示できるようになりました。

キーワードを設定するには、**メインメニュー > Incident Collaboration > Playbooksタブ**からキーワードを設定したいプレイブックを選び、**Edit > Actions > Prompt to run the playbook when a user posts a message > Containing any of these keywords**にキーワードを入力します。

![incident-keyword-setting](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-keyword-setting.png)

設定したキーワードを含む投稿を行うと、Playbook Botがプレイブックの開始を促すメッセージを投稿します。

![incident-keyword-react](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-keyword-react.png)

`Yes, run playbook`を選択すると、作成するプレイブックの内容を指定するモーダルが開きます。

![incident-keyword-modal](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-keyword-modal.png)

内容を入力し、`Start run`ボタンを押すとインシデント作成のきっかけとなった投稿とともにプレイブックを開始できます。

![incident-keyword-creation](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-keyword-creation.png)

### (E10/E20/Cloud) Retrospectiveレポート

作成したプレイブックに関するRetrospective(振り返り)レポートを作成できるようになりました。RetrospectiveレポートをPublishすると、チャンネルに内容が投稿されます。

![incident-retro](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-retro.png)

Playbookの設定から、Retrospectiveレポートのテンプレート等を設定できます。

![incident-retro-template](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-retro-template.png)

### (E20/Clout) Playbookダッシュボード
プレイブックの過去の実行結果を把握するためのダッシュボードを表示できるようになりました。インシデントの発生頻度やインシデント対応の参加者数、対応にかかった時間などを俯瞰することで、インシデント対応方針の改善やリソース割り当ての見直し等に役立てることができます。

![incident-dashboard](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/incident-dashboard.png)

## 絵文字のスキントーン選択

Emoji 13.0に基づく絵文字を利用できるようになり、絵文字ピッカーから絵文字を選択する際にスキントーン（肌の色）を選択できるようになりました。

![emoji-skin-tone](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/emoji-skin-tone.webp)

(画像は[公式ブログ](https://mattermost.com/blog/mattermost-v5-37/#emoji)から)

## Focalboard Plugin

[Focalboard](https://www.focalboard.com/)にも改善があります。

現在、Mattermost Pluginとして利用できるFocalboardの最新バージョンは `v0.7.0` で、このバージョンのインストール方法については以下の公式ドキュメントを参照してください。　　
https://www.focalboard.com/download/mattermost/latest-plugin/

Focalboard自体は先日[`v0.8.0`](https://github.com/mattermost/focalboard/releases/tag/v0.8.0)がリリースされています。

### プロパティ値によるテーブルのグルーピング
作成したタスクをテーブル表示する際に、タスクに設定したプロパティの値ごとにグループ化して表示することができるようになりました。
以下の例では、`Priority`プロパティをグルーピングの対象とした場合に、**Priority無し**、**High**、**Medium**、**Low** でグループ化されたタスクを表示しています。

![focalboard-grouping](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/focalboard-grouping.png)

### プロパティタイプの追加
タスクに設定するプロパティのタイプに、マルチセレクト、人物、チェックボックスが追加されました。

![focalboard-property-type](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/focalboard-property-type.png)

## その他の変更

### チャンネル切替ダイアログの改善

Mattermost上で `Ctrl(Cmd) + K` を入力することで開く **チャンネル切替** ダイアログに、最近開いたチャンネルが表示されるようになりました。

![channel-switcher](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/channel-switcher.png)

### カスタムステータスの有効期限設定

カスタムステータスを設定する際に、ステータスの有効期限を設定できるようになりました。

![custom-status](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.37/custom-status.png)

### `platform` バイナリの廃止

以前、Mattermostサーバーは `mattermost/platform` というリポジトリで管理されており、当時の名残でMattermostサーバー管理用のCLIツールとして`platform`バイナリというものが残っていました。今回のリリースでこの`platform`バイナリや`--platform`オプションが利用できなくなりました。
現在では、リポジトリは[`mattermost/mattermost-server`](https://github.com/mattermost/mattermost-server)に移行され、Mattermostサーバー管理用のCLIツールとして`mattermost`バイナリが利用可能になっているため、今後は`mattermost`バイナリを使用することが推奨されています。

## その他のトピック

### Mattermost Dockathon

Mattermostの公式ドキュメントサイトである [https://docs.mattermost.com](https://docs.mattermost.com) の構成などを改善する作業を開始しているようで、それに伴い、今月下旬から **Mattermost Docathon** というドキュメント改善のためのイベントを開催するようです。

[Join Us for our First Mattermost 'Docathon' and win swag and more\!](https://mattermost.com/blog/docathon-2021/)

このイベント期間中のコントリビュートが多かった上位5名にMattermostロゴ入りAirPod Proがプレゼントされるようです。1件のみのコントリビュートでもグッズがもらえるようです。

### Mattermost v6.0

先月の記事で少し触れたMattermot v6.0について、公式ブログで紹介がありました。

[Looking ahead to Mattermost v6\.0, which ships Fall 2021](https://mattermost.com/blog/looking-forward-to-mattermost-v6-0/)

ベータ版からGA(Generally Available)に昇格予定の機能や、廃止予定の機能などが紹介されています。

## おわりに

次の`v5.38`のリリースは 2021/08/16(Mon)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
