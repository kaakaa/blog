
---
title: "Mattermost4.5の新機能"
date: 2017-12-24T14:25:15+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

2017/12/16にMattermost 4.5.0がリリースされたので、追加された機能などについて簡単に紹介します。

詳細については公式のリリースを確認ください。
[Mattermost 4.5: Zoom plugin for video, audio and screen sharing, plus new mobile features - Mattermost](https://about.mattermost.com/blog/mattermost-4-5/)
[Mattermost Changelog — Mattermost 4.5 documentation](https://docs.mattermost.com/administration/changelog.html#release-v4-5)

# 新機能

## Zoomプラグイン

Mattermost上から簡単にビデオカンファレンスサービスである[Zoom](https://zoom.us/)を利用できる機能がプラグインとして追加されました。
このプラグインのソースコードも独立のリポジトリとして公開されているため、サーバープラグインのリファレンス実装として参考になりそうです。
[mattermost/mattermost-plugin-zoom: Zoom plugin for Mattermost](https://github.com/mattermost/mattermost-plugin-zoom)

ただ、このプラグインを利用するために必要なZoomのAPI KeyなどはZoomの有料のDevelopersアカウントを作成する必要がありそうです。
[My Account – Zoom Developers](https://developer.zoom.us/me/)

## E20: Actianceコンプライアンスエクスポート

企業ではコンプライアンス上、Mattermost上の投稿やデータなどを保持しておく必要があり、そのためのポリシーを自社で策定したりしています。
上記を満たすためのアーカイブ機能として、自動でMattermost上のデータをサードパーティシステムであるActianceにエクスポートするための機能が追加されたようです。
[Compliance Export Beta (E20 Add-On) — Mattermost 4.5 documentation](https://docs.mattermost.com/administration/compliance-export.html)

## モバイルアプリのアップデート

* VideoやPDF, MS Office製品のファイルなどを閲覧できるようになりました

タップするとダウンロードが始まり...

<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/19350852-b31a-39a2-7234-4c64d5f86a36.jpeg" height="360px">

PDFやPPTXファイルなどはこんな感じで表示されます。

<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/f8546717-b6fc-9d39-15c7-9b481debeb4a.jpeg" height="360px">

* モバイルアプリ上でもスラッシュコマンドを選択式に入力できるようになりました
<img src="https://qiita-image-store.s3.amazonaws.com/0/9891/f21f9ec2-f11a-0c41-660b-a8064b6a2aa9.jpeg" height="360px">

* iOSアプリの3Dタッチ機能を利用して既読操作を簡単に行えるようになったようです
* Markdownのイメージ記法を利用すると、投稿内でレンダリングされるようになりました。（以前はURLリンクを表示する形式だった模様）
* iPhone Xをサポートしたようです

# API Version 3 Deprecated Jan. 16, 2018

以前よりアナウンスされていましたが、Mattermost API v3が次回のリリース日である2018/1/16に廃止になります。

# オープンソースプロジェクト（コミュニティ機能）

## Giphy Integration on AWS Lambda

AWS Lambdaを利用した[Giphy](https://giphy.com)との連携機能のようです。
[pableu/mattermost-giphy-lambda: AWS Lambda Function to create a Slash-Command for Giphy-Integration in Mattermost](https://github.com/pableu/mattermost-giphy-lambda)

## MantisBT Plugin Integration

Issue Trackerである[MantisBT](http://www.mantisbt.org)との連携機能のようです。
[aalasolutions/MantisBT-Mattermost: Mantis integration plugin with Mattermost](https://github.com/aalasolutions/MantisBT-Mattermost)

## Laravel Web Framework for Mattermost PHP driver

PHPのフレームワークであるLaravelと連携する機能のようです。
[gnello/laravel-mattermost-driver: A Laravel integration for the package php-mattermost-driver](https://github.com/gnello/laravel-mattermost-driver)


## おわりに

次回の`v4.6.0`のリリースは2018/1/16を予定しています。

