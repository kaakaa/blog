---
title: "Mattermost 5.38の新機能"
emoji: "🎉"
published: true
date: 2021-08-21T00:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/08/16 に Mattermost `v5.38.0` がリリースされました。
その後、2021/08/18 に 同梱されているMattermost Playbook (旧: Mattermost Incident Collaboration) v1.16.0 がインストールできない問題に対応した `v5.38.1` がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.38 is now available \- Upgrade today](https://mattermost.com/blog/mattermost-v5-38/)
  - [Mattermost 5\.38\.1 released \- Upgrade today if your deployment is affected](https://mattermost.com/blog/mattermost-5-38-1-released/)
- [Mattermost Changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v5-38-feature-release)

---

## [アップグレード時の注意事項](https://docs.mattermost.com/administration/changelog.html#important-upgrade-notes)

* 返信スレッドの折りたたみ機能(ベータ版)導入によりスレッドやチャンネルのメンション数や未読数が不正確となってしまう問題が発生していました。この改善が`v5.38`にて行われていますが、大規模なデータベースの場合、`v5.38`アップグレード時のDBマイグレーションに数分かかることがあります。
* Mattermost v5.38.0と同時にリリースされたFocalboard v0.8.2 は、新たに追加されたデータベースコネクションシステムの影響でMattermost v5.37以降のバージョンでないと動作しなくなっています

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## はじめてログインした際の案内画面の改善

Matttermostにアカウントを作成し、初めてログインした時に表示される「はじめに」に当たる画面の内容が改善されました。この画面から、プロフィールの入力、デスクトップ通知の設定、ユーザーの招待を行うことができます。
![getting started](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.38/gettingstarted.png)

## `Mattermost Incident Collaboration` が `Mattermost Playbooks` へリブランディング

最近、活発に開発が行われていた`Mattermost Incident Collaboration Plugin`ですが、インシデント管理に限らず、あらゆる種類のチーム活動に訴求できるということから`Mattermost Playbooks`にリブランディングされました。

https://github.com/mattermost/mattermost-plugin-playbooks

また、以下のような機能追加も行われています。

* 右サイドバーの表示内容改善
* Mattermostプロジェクトで培われたベストプラクティスのテンプレート化
* Playbookチャンネル参加時、左サイドバーが自動で整理されるトリガーの追加
* Playbookが終了した際に自動でチャンネルエクスポート(E20/Cloud)が実行されるトリガーの追加

機能追加の詳細については [公式ブログ](https://mattermost.com/blog/mattermost-v5-38/#playbooks) を参照ください。

## `Focalboard`の更新

Mattermost v5.38のリリースに合わせて [Focalboard](https://www.focalboard.com/)も`v0.8.2`に更新されています。

* `作成者`プロパティの追加
* DBコネクションのパフォーマンス改善

## (E20) 詳細なデータ保持ポリシー設定

本バージョンからデータ保持ポリシーを特定のチーム及びチャンネルごとに設定できるようになりました。これによりコンプライアンスの維持とストレージの最適化を両立することができます。

(公式ブログでの本機能の紹介と、動作確認をしている[Mattermost Cloud](https://mattermost.com/mattermost-cloud/)の画面が一致していないように見えるため、詳細については割愛します)

## その他のトピック

### Mattermost Dockathon

ドキュメント改善をテーマにしたハッカソンである Mattermost Dockathon が 2021/07/26~08/06に行われました。
約2週間の間に **26** コントリビューターから **178** のコントリビュートがあったようです。

結果については以下のエントリにまとめられています。

[Highlights from the first\-ever Mattermost Docathon](https://mattermost.com/blog/mattermost-docathon-2021-highlights/)

## おわりに

3年半ぶりのメジャーバージョンアップとなる、次の`v6.0`のリリースは 2021/09/15(Web)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
