
---
title: "Mattermost5.0の新機能"
date: 2018-06-18T08:14:58+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/06/15にMattermost 5.0.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost 5\.0: Intercept and modify posts, advanced permissions, longer posts and more – Mattermost](https://about.mattermost.com/releases/mattermost-5-0/)
* [Mattermost Changelog — Mattermost 5\.0 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-0)

### メジャーバージョンアップについて
[Upcoming Changes with Mattermost v5\.0 \- Announcements \- Mattermost Peer\-to\-Peer Forum](https://forum.mattermost.org/t/upcoming-changes-with-mattermost-v5-0/5119)
今回メジャーバージョンアップが行われた理由として、Enterprise版向けの高度な権限設定機能を追加するにあたり、データベースへ大きな変更が必要となったことが挙げられています。
また、データベースへの影響の他にも今回のメジャーバージョンアップにより以下の変更が行われています。

* `サイトURL`の設定が必須に
* Mattermost APIv3の廃止
* 管理者向けCLIツールの名称が`platform`から`mattermost`に変更

その他のアップデートにおける注意点について、詳しくは[Important Upgrade Notes](https://docs.mattermost.com/administration/important-upgrade-notes.html#important-upgrade-notes)を参照してください。

# アップデート内容

## プラグインによる投稿の検閲

Mattermostのプラグインに、投稿内容がデータベースに保存される直前・直後に処理を差し込めるAPIが追加されました。
[Hooks.MessageWillBePosted](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.MessageWillBePosted)

この機能を使うことで、有害な言葉を含む投稿をできなくしたり、個人情報のような重要な文字列を別の言葉に置き換えることなどができるようになります。

以下のリンクから、Mattermost CTOによるこの機能の解説動画を見ることができます。

* [Coming soon: \(APIv4\) Mattermost Post\-Intercept \- YouTube](https://www.youtube.com/watch?v=YgMqFZARzfs)
* [Coming soon: \(APIv4\) Mattermost Post\-Intercept \- Feature Idea Discussion / Design Feedback Request \- Mattermost Peer\-to\-Peer Forum](https://forum.mattermost.org/t/coming-soon-apiv4-mattermost-post-intercept/4982)
  * 上記YouTube動画の書き起こし

また、Mattermostプラグインの開発方法については下記のドキュメントを参考になります。
[Plugins \(Beta\)](https://developers.mattermost.com/extend/plugins/)

## (E10/E20)高度な権限設定

Enterprise版の機能として、チーム/チャンネル/投稿の作成・修正・削除に関する権限設定を柔軟に行えるようになりました。

<img width="929" alt="スクリーンショット 2018-06-17 15.31.45.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/5f96776b-8c19-0633-e5d2-4b1716c6854b.png">


権限設定はJSONファイルとしてエクスポート/インポートすることもできます。

また、E20プランの場合、追加で下記のような権限設定も可能となります。

- システム全体の設定だけでなくチームごとに権限設定可能
- 追加のチャンネル権限設定(読取専用チャンネルの作成、@all/@channel/@hereの使用制限、メンバー管理の制限)
- ユーザー単位やLDAPのグループ単位の権限付与

この機能についての開発チーム/ビジネスチームに対するデモの様子がYouTubeに上がっています。

[Mattermost \| Permission Schemes demo \- YouTube](https://www.youtube.com/watch?v=aKF_xYWD9LU&t=429s)
[Mattermost 5\.0 \| Advanced Permissions Demo and Q&A \- YouTube](https://www.youtube.com/watch?v=tRJWxK4_2u8)

また、権限管理機能に関する今後のロードマップについては、下記で触れられています。
[Upcoming Permissions changes: Schemes, Roles, Guest Accounts, Channel Permissions \- Announcements \- Mattermost Peer\-to\-Peer Forum](https://forum.mattermost.org/t/upcoming-permissions-changes-schemes-roles-guest-accounts-channel-permissions/4971)

```
Phase 1 (v4.9, April 2018): Backend work already implemented. No visible changes for end users or Admins.
Phase 2 (v5.0, June 2018): Permission Schemes.
Phase 3 (Q4 2018): Channel Permissions and Guest Accounts.
Phase 4 (Q1 2019): Supplementary Roles to grant individuals extra permissions.
Phase 5 (TBD): Supplementary Roles that can be synced with and granted to AD/LDAP groups.
```

チャンネル管理の権限設定や、ゲストアカウントの作成等が2018/4Qに予定されています。
ゆくゆくはAD/LDAPグループとMattermostの権限管理の同期も挙げられており、エンタープライズユーザーに対して親和性のあるコミュニティ基盤を目指していることが見て取れます。


## 投稿の最大文字数の拡大
今まで一つの投稿に含められる最大の文字数は4,000文字でしたが、今回のバージョンから最大文字数が16,383文字に引き上げられました。
これにより、大きなMarkdownのテーブルなども投稿可能になります。

この機能を有効にするには使用しているデータベースのマイグレーションが必要になります。

使用しているデータベースがMySQLの場合は

```
ALTER TABLE Posts MODIFY COLUMN Message TEXT;
```

PostgreSQLの場合、

```
ALTER TABLE Posts ALTER COLUMN Message TYPE VARCHAR(65535)
```

を実行し、Mattermostを再起動することで有効になります。
`Posts`テーブルが巨大な場合、パフォーマンスに影響が出る可能性があるため、ユーザーの少ない時間帯に実行することが推奨されています。


## Enterprise版の機能がTeam Edition(OSS版)に

Enterprise版でしか使用できなかった下記の機能がOSS版であるTeam Editionでも使用可能となりました。

### カスタムブランディング

- [Custom Branding Tools — Mattermost 5\.0 documentation](https://docs.mattermost.com/administration/branding.html)
* ログイン画面に組織独自の画像やテキストを設定可能に
* 設定: **システムコンソール** > **カスタマイズ** > **独自ブランド設定**

### チーム毎のテーマ設定
* チーム毎にMattermost画面のテーマカラーを変更可能に
* 設定: **アカウント設定** > **表示** > **テーマ** > (テーマ変更後`全ての自分のチームに新しいテーマを適用する`のチェックを外して更新) 

<img width="788" alt="スクリーンショット 2018-06-17 15.45.03.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/05ef9dc3-0b74-ce76-2569-01da2aaf5572.png">

### パスワードルールの設定
* パスワードのルール(パスワードの最低の長さ、英数字記号を必須とするかなど)を設定可能に
* 設定: **システムコンソール** > **セキュリティ** > **パスワード**

<img width="934" alt="スクリーンショット 2018-06-17 15.39.24.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/f3cec8e0-ceb5-7bb7-caee-d61008fbb6c8.png">

## 参加・脱退メッセージの結合

今まではチャンネル参加/脱退メッセージが一つ一つ投稿されていましたが、これらのメッセージが連続した場合、まとめて表示されるようになりました。

![Image Pasted at 2018-6-13 22-10.png](https://qiita-image-store.s3.amazonaws.com/0/9891/2281ef70-42e8-5cf9-534d-968d7149f51a.png)

## モバイルアプリの改善
* Androidユーザー向けの初回ロード時間の改善
* プッシュ通知画面の改善

その他の変更点については[mattermost\-mobile/CHANGELOG\.md at master · mattermost/mattermost\-mobile](https://github.com/mattermost/mattermost-mobile/blob/master/CHANGELOG.md#v190-release)を参照してください。

## デスクトップアプリの改善
* メモリキャッシュ削除間隔の見直しによるパフォーマンス改善
* 新しい"Enable GPU Hardware acceleration"オプションによる安定性の向上、および動作可能時間の改善
* UI更新によるUXの改善

その他の変更点については[desktop/CHANGELOG\.md at master · mattermost/desktop](https://github.com/mattermost/desktop/blob/master/CHANGELOG.md#release-v412)を参照してください。

## その他

### Extended Support Release

いくつかの顧客からの要望を受け、Mattermost v4.10について１年間のセキュリティバックポートが行われます。
毎月のアップデートが負担になっている場合、作業負荷を軽減することができるようになります。



# おわりに

次回の`v5.1`のリリースは2018/7/16を予定しています。

