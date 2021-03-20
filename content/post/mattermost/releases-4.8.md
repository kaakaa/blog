
---
title: "Mattermost4.8の新機能"
date: 2018-03-18T09:55:08+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

# はじめに

2018/03/16にMattermost 4.8.0がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4\.8: Faster loading assets with Amazon CloudFront and simplified SAML migration – Mattermost](https://about.mattermost.com/releases/mattermost-4-8/)
[Mattermost Changelog — Mattermost 4\.8 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-8)

# アップデート内容

## Amazon CloudFrontサポート

画像やJavaScriptファイルなどの静的リソースをAmazon CloudFrontでホストできるようになりました。
これにより、キャッシュのパフォーマンスが改善され、Mattermostのチームメンバーが世界中に散らばっていても短いロード時間でMattermostを使用することができます。

設定方法については以下のドキュメントを参照してください。
[Configuring CloudFront to host Mattermost static assets — Mattermost 4\.8 documentation](https://docs.mattermost.com/install/config-cloudfront.html)

## SAMLマイグレーションコマンド(E20)

ユーザーを簡単にSAML 2.0へマイグレートできるコマンドラインツールが追加されました。
認証にOkta, OneLogin, Microsoft's Active Directyory Federation Serviceを使用している会社でシングルサインオンが可能になります。

## プラットフォームの改善

* DBクエリやWebSockeイベントの最適化によりロード時間が短縮されました
* 20MB以上のファイルをアップロードが可能となるiOSのエンドポイントが作成されました
* タブレット画面でチャンネルヘッダーを見やすくの検索バーの変更など、Webアプリにいくつかの修正が行われました
* プロフィール写真への属性追加など508complianceの改善が行われました
  * `508compliance`とは、身体障害を持つ人に対するアクセシビリティを向上することを義務付けた米国の法律だそうです
  * [リハビリテーション法第508条 \- Wikipedia](https://ja.wikipedia.org/wiki/%E3%83%AA%E3%83%8F%E3%83%93%E3%83%AA%E3%83%86%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E6%B3%95%E7%AC%AC508%E6%9D%A1)
* 内向きのウェブフックがマルチパートやフォームを受け入れられるようになりました

## Extended Support Release

現在、Mattermostは毎月16日に新しいバージョンをリリースし、３リリース前までセキュリティ修正のバックポートを行うというポリシーでリリースが行われていますが、
Mattermostの顧客やコミュニティメンバーからリリース頻度を落として欲しいという要望が上がっているようです。

そこで、最大で1年間セキュリティ修正などのクリティカルな修正をバックポートするExtended Support Releaseを作成する案が検討されており、それについて意見を募集しています。
[Extended Support Release Discussion](https://forum.mattermost.org/t/extended-support-release-discussion/4598)

## その他(APIv3について)

先月のリリースでAPIv3がdeprecatedとなりましたが、APIv3はMattermost v5.0で削除される予定です。
APIv3からAPIv4へのマイグレーションについては下記を参考にしてください。
[Mattermost API Reference](https://api.mattermost.com/#tag/APIv3-Deprecation)

気になるMattermost v5のリリース時期ですが、「v5で削除される機能のアナウンスは２ヶ月前の6月にするのが良いだろう」という話がされているので8/16になるのかな、と思っています。
![スクリーンショット 2018-03-17 23.12.02.png](https://qiita-image-store.s3.amazonaws.com/0/9891/2c878db8-188d-fd7b-9854-48c693c4c2e8.png)


# おわりに

次回の`v4.9.0`のリリースは2018/4/16を予定しています。

