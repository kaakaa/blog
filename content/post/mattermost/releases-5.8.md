---
title: "Mattermost5.8の新機能"
date: 2019-02-18T15:17:06+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2019/2/15にMattermost v5.8.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

* [Mattermost 5\.8: MFA for Team Edition, LDAP group sync, improved image performance and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-8/)
* [Mattermost Changelog — Mattermost 5\.8 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-8)

---

# アップデート内容

## Team Editionに多要素認証機能が追加
[コミュニティからの要望](https://mattermost.uservoice.com/forums/306457-general/suggestions/19636537-make-mfa-part-of-the-team-edition)として、今までEnterprise Edition限定機能であったGoogle Authenticatorによる多要素認証機能がOSS版であるTeam Editionでも利用できるようになりました。
これにより、悪意のある第三者によるなりすましなどのリスクを低減することができます。
今回Team Editionに追加された多要素認証機能はユーザーが自ら設定を行うことで有効になるものであり、[サーバー内すべてのユーザーに多要素認証を強制する設定](https://docs.mattermost.com/deployment/auth.html#enforcing-mfa-e10)はEnterprise Edition限定機能のままです。

多要素認証の設定を行うには、まず **システムコンソール** > **認証** > **多要素認証** > **多要素認証を有効にする** を有効にします。
![スクリーンショット 2019-02-18 10.06.43.png](https://qiita-image-store.s3.amazonaws.com/0/9891/0d2beab9-64f9-1db8-7409-5bbac5542e78.png)


多要素認証を有効にすると **アカウント設定** > **セキュリティー** > **多要素認証** のメニューが現れるため、**多要素認証をあなたのアカウントに追加する** をクリックしGoogle Authenticatorの設定を行います。
![スクリーンショット 2019-02-18 10.08.54.png](https://qiita-image-store.s3.amazonaws.com/0/9891/68144a0b-d32d-5dd3-9215-b82521bbd977.png)


画面の指示に従い、QRコードをスキャンするかシークレットコードを入力し、アカウントをGoogle Authenticatorに登録します。すると、登録したアカウントの多要素認証コードがGoogle Authenticatorに表示されるので、そのコードを画面下の"多要素認証コード"欄に入力し、「保存」ボタンを押します。
![スクリーンショット 2019-02-18 10.10.14.png](https://qiita-image-store.s3.amazonaws.com/0/9891/376f306b-e18d-7c06-25e7-31dc161f9db6.png)

多要素認証コードが正しく入力されると、多要素認証の登録完了画面が表示されます。
![スクリーンショット 2019-02-18 10.10.58.png](https://qiita-image-store.s3.amazonaws.com/0/9891/8b22dc8f-d6d6-e852-0ace-a07a0a8c5ad2.png)


以降、このMattermostにログインする際にはユーザー名/パスワードによるサインインの後、多要素認証コードの入力を求める画面が表示されるようになります。ここには登録時と同様、Google Authenticatorで表示される多要素認証コードを入力することでログインを完了することができます。
![スクリーンショット 2019-02-18 10.11.35.png](https://qiita-image-store.s3.amazonaws.com/0/9891/84893711-59a9-79be-a055-6ed633dae5f7.png)



## (E20) LDAPグループ機能の追加
LDAPのグループに基づいてデフォルトのチームやチャンネルのメンバーシップを自動同期できる機能がE20版のベータ機能として追加されています。
グループ同期機能は[今後も機能追加が行われる予定](https://forum.mattermost.org/t/ldap-group-sync-alpha-release/6351)であり、E10版でも利用可能になる予定とのことです。

## 画像プロキシのデフォルト化

v5.8.0より画像プロキシがデフォルトで有効になりました。今後は画像プロキシによりキャッシュが行われるため、画像を含む投稿の表示が高速になります。また、Mattermost内で通信が完了するため、サードパーティーのサービスを使用するよりプライバシー性の向上が期待できます。

## その他の変更
* @channel, @all, @hereなどによるチャンネル全体へのメンションを通知として受け取るかどうかを設定できるようになりました
![スクリーンショット 2019-02-18 10.37.59.png](https://qiita-image-store.s3.amazonaws.com/0/9891/a06286ff-afe3-467d-2fdc-1df23f545022.png)

* スラッシュコマンドを実行した際、スラッシュコマンドを実行したチャンネル以外にもメッセージを投稿できるようになりました
  * [`model.CommandResponse`](https://github.com/mattermost/mattermost-server/blob/release-5.8/model/command_response.go#L24)に追加された`ChannelId`フィールドを指定することで実現できます。（[サンプルコード](https://gist.github.com/kaakaa/b040bcd0a6fdccc18b3daaea214c6112#file-mm_v5-8_response_to_other_channel-go)）
  * コマンド実行ユーザーが所属しているチャンネル以外のチャンネルに投稿しようとするとエラーとなります


# おわりに

次回の`v5.9`のリリースは2019/3/15(Fri)を予定しています。
そして機能追加が行われる`v5.10`は恐らく2019/4/16(Tue)のリリースになるかと思います。

---

[Mattermost 日本語\(@mattermost\_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) でMattermostに関する日本語の情報を提供しています。

