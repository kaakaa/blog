---
title: "Mattermost5.4の新機能"
date: 2018-10-27T15:30:06+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/10/16にMattermost 5.4.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost Changelog — Mattermost 5\.4 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-4)
* [Mattermost 5\.4: User\-focused features, new data export tool, mobile enhancements and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-4-user-focused-features-new-data-export-tool-mobile-enhancements-and-more/)

---

# アップデート内容

## ユーザーの生産性向上のための改善

### メッセージ作成中の表示
メッセージ作成中に別のチャンネルに移動した際、チャンネルの表示アイコンが変更されるようになり、送信前のメッセージが残っていることが一目でわかるようになりました。
サイドバーとチャンネル切替画面で確認できます。

<img width="1082" alt="スクリーンショット 2018-10-26 20.23.31.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/db9d0e3a-10a2-70bb-c61c-a97077bc1334.png">

### ジャンボ絵文字
絵文字だけの投稿を行った際、絵文字が大きく表示されるようになりました。
大きさは `### :+1:` と同じぐらいなので、Markdown形式で `# :+1:` とすればより大きな絵文字を表示することもできます。

<img width="274" alt="スクリーンショット 2018-10-27 16.28.30.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/3e4597db-9728-2056-f68b-4f3823fe58a2.png">

### ダイレクトメッセージ・グループメッセージを指定した検索
以前より、メッセージ検索ボックスで `in:` と入力することで検索対象のチャンネルを指定して検索することができましたが、v5.4から `in:` でのチャンネル指定にダイレクトメッセージチャンネルとグループメッセージチャンネルを指定できるようになりました。

### メッセージ入力ボックスの拡張
長いメッセージを入力する際、メッセージ入力ボックスの高さが今まで以上に拡張されるようになりました。
v5.3以前は7行程度までの表示でしたが、v5.4以降は13行ほど表示できるようになっています。

<img width="857" alt="スクリーンショット 2018-10-26 23.58.35.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/6a454a71-5933-9e8d-c0ce-5fcf05bad985.png">


### Interactive Messageにセレクトメニューが追加

[Interactive Message](https://docs.mattermost.com/developer/interactive-messages.html#message-menus)のアクションとして、ボタンだけでなくセレクトメニューを指定できるようになりました。

<img width="608" alt="スクリーンショット 2018-10-27 16.30.25.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/61531d68-666a-41cd-03f9-d67b496dc95a.png">

### コードブロックのCTRL + ENTERでの送信
**アカウント設定** > **詳細** > **CTRL + ENTER でメッセージを投稿する**をオフにしている場合に限り、**アカウント設定** > **詳細** > **CTRL + ENTERでコードブロックメッセージを送信します**というメニューが現れます。
これを `オン` にすると、”\`\`\`” で始まるコードブロック入力中はENTERキーで改行を入力でき、その状態で`CTRL + ETNER`を押すとコードブロック終了の ”\`\`\`” を入力しなくてもコードブロックとしてメッセージを送信することができます。

<img width="787" alt="スクリーンショット 2018-10-27 16.31.15.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/e2089920-df87-8130-b491-c6609e9ef872.png">


### ピン留めされたメッセージの更新
`ピン留めされた投稿`のサイドバーを開いた状態でチャンネルを切り替えると、`ピン留めされた投稿`のサイドバーの内容も切り替え先のチャンネルの内容に切り替わるようになりました。

## Mattermostサーバーの統合
新たなデータエクスポート用のCLIツールが追加され、複数のMattermostサーバーをマージできるようになりました。
このツールで現在エクスポート可能なデータは、基本的なデータである投稿・ユーザー・返信・チーム・チャンネルのみですが、コミュニティによってより多くのデータが対象になるよう改善が進められています。

[Bulk Export Tool — Mattermost 5\.4 documentation](https://docs.mattermost.com/administration/bulk-export.html)
[Bulk Loading Data — Mattermost 5\.4 documentation](https://docs.mattermost.com/deployment/bulk-loading.html)

10月中に開催されている[Hacktober](https://hacktoberfest.digitalocean.com/)を利用して多くの機能拡張が行われている最中です。[Search · bulk export](https://github.com/mattermost/mattermost-server/search?q=bulk+export&type=Issues)

## 管理機能
### 参加/脱退メッセージの非表示オプション

**アカウント設定** > **詳細** > **参加/脱退メッセージを有効にする** を`オフ`にすることで、ユーザーがチャンネルに参加・脱退した時に表示されるシステムメッセージを非表示にすることができるようになりました。

<img width="791" alt="スクリーンショット 2018-10-27 15.47.22.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/1ef217b6-5240-8445-840c-b460bb6b91dd.png">

数百人が参加するようなチームだとタウンスクエアなどのチャンネルが参加通知ばかりになり会話をしにくいこともあるため、そのような場合に有効な設定です。

<img width="540" alt="スクリーンショット 2018-10-27 15.51.32.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/96aeba1b-75ed-d47a-0859-089e5027beeb.png">


### チーム管理者に他ユーザーの投稿の編集権限を付与するかの設定
**システムコンソール** > **全般** > **ユーザーとチーム** > **チーム管理者が他者の投稿を編集できるようにする** からチーム管理者に他ユーザーの投稿を編集する権限を付与するかどうかを設定することができるようになりました。
この設定に関わらずシステム管理者は常に他ユーザーの投稿を編集することができます。

## コンプライアンス
### E20: カスタム利用規約
E20版限定ですが、利用者独自の利用規約を設定し、その規約に同意しないと利用できないよう設定することができるようになりました。
企業・組織内で運用する場合に必要になる機能かと思います。

**システムコンソール** > **カスタマイズ** > **法的事項とサポート** > **カスタム利用規約を有効にする(ベータ版)** から設定することができます。

## モバイルアプリの改善

* 投稿に対するハッシュタグを利用することができるようになりました
* リアクション絵文字を長押しすることで、誰がリアクションしたかを確認できるようになりました
* iPhone XR、 XS、 XS Maxのサポートが追加されました

# おわりに

次回の`v5.5`のリリースは2018/11/16(Fri)を予定しています。

次回のリリースから、Mattermostは新しいリリースサイクルを採用し、**奇数月**のリリースはバグフィックスなどの品質向上に関する修正のみを加えたリリース、**偶数月**に新規機能の追加を含むリリースを行うそうです。

[Document minimum server version for plugin API methods by hanzei · Pull Request \#9616 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/pull/9616#issuecomment-429815195)

> Mattermost is moving to a different release cycle starting in October, where
> * every second month we have a feature release (October, December, February, ..), and
> * every second month we have a quality/bug fix release (November, January, ..).
> So given this is not a bug fix, it's shipped in v5.6 by default.


---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。

