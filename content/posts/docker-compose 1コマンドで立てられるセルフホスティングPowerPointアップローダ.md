
---
title: "docker-compose 1コマンドで立てられるセルフホスティングPowerPointアップローダ"
date: 2016-04-06T23:00:26+09:00
draft: false
toc: true
tags: ["PowerPoint","docker","docker-compose"]
---

# これは何？

[kaakaa/ppt-museum: ppt/pptx file uploader](https://github.com/kaakaa/ppt-museum)

`docker-compose up -d`だけで起動するPowerPoint資料アップローダを作っています。

トップ画面
<img width="480px" ald="top" src="https://qiita-image-store.s3.amazonaws.com/0/9891/eb1c9dde-803c-d94c-4f55-1a5f93b230c3.png">

スライド画面
<img width="480px" alt="スライド画面" src="https://qiita-image-store.s3.amazonaws.com/0/9891/425bdf28-7cfd-9f1a-7489-83a9856edb4a.png">



# 使い方

## 必須環境

* docker
* docker-compose

## 実行コマンド

```
$ git clone https://github.com/kaakaa/ppt-museum.git
$ cd ppt-museum

// docker v1.10.0+ / docker-compose v1.6.0+ の場合 
$ docker-compose up -d

// 上記のバージョン以外の場合
$ docker-compose -f docker-compose-v1.yml up -d

```

`http://localhost:4567/` にアクセスするとトップページが表示されます。

# 内部構造

## アーキテクチャ

4つのDockerコンテナを起動しています。

* [nginx](https://hub.docker.com/_/nginx/)
* [mongo](https://hub.docker.com/_/mongo/)
* [jodconverter](https://hub.docker.com/r/kaakaa/jodconverter/)
  * => [kaakaa/jodconverter-container: Dockerfile for JODConverter](https://github.com/kaakaa/jodconverter-container)
  * PowerPointファイルをPDFファイルに変換するコンテナ
* [ppt-museum-web](https://hub.docker.com/r/kaakaa/ppt-museum-web/)
  * => [kaakaa/ppt-museum-webapp: Webapp for ppt-museum](https://github.com/kaakaa/ppt-museum-webapp)
  * Javaのマイクロフレームワーク [Spark Framework - A tiny Java web framework](http://sparkjava.com/) ベースのWebアプリコンテナ

<img width="640px" alt="architecture" src="https://qiita-image-store.s3.amazonaws.com/0/9891/7d045e14-fd44-e7c2-b925-640bfccfa938.png">

## 既知の問題

* DB構造が変わったときのマイグレーションパスを用意していない
  * アップロード資産をZipダウンロードして、[kaakaa/ppt-batch](https://github.com/kaakaa/ppt-batch)でアップロードし直すという最底辺の人力運用でカバーの現状
* UI周りが要改善
  * フロントエンド力が圧倒的に足りない
* テストコードを全然書いていない

他にもたくさんありますが。。。


# 背景

以降は蛇足ですが、これを作ろうと思った背景について。

* * *

社内勉強会の資料をアップロードするために作りました。

社内勉強会の資料には、社外NGの話が含まれていたり、「外に公開するほどの…」という方もいたりするので、そういう資料を個人で貯めこむことの無いようにと作りました。また、勉強会開始当初は勉強会用に立ててるWikiに資料を添付するという運用を取っていましたが、Wikiに添付された資料をわざわざダウンロードして開くことなんて無いので、閲覧機会の損失を防ぎたかったという想いもあります。

* * *

その昔に[kaakaa/pptgallery](https://github.com/kaakaa/pptgallery)を作って今まで運用していました。
しかし、コレには

* LibreOffice / ImageMagick などのミドルに依存しており、二度と環境が作れる気がしない
* ruby / jsを手探りで組んでいたのでメンテできる気がしない

などの辛い面が多くあったため、匙を投げることにしました。
そして作りなおしたのが [kaakaa/ppt-museum](https://github.com/kaakaa/ppt-museum) です。

[kaakaa/pptgallery](https://github.com/kaakaa/pptgallery)の反省から

* 環境はどこでも立てられるようにdocker-composeを使う
* ruby/jsに夢を見るのをやめて、メンテする気になれるよう一番慣れているJavaで書く

という方針で作っています。

* * *

また、作りなおそうと思ったきっかけとなったのが、[azu/pdf.js-controller](https://github.com/azu/pdf.js-controller)を見つけたことでした。スライド表示部分では[azu/pdf.js-controller](https://github.com/azu/pdf.js-controller)をそのまま使っています。

PPTGalleryの頃にも、PowerPointファイルを[JODConverter | Art of Solving](http://www.artofsolving.com/opensource/jodconverter.html)でPDFファイルに変換することはできていましたが、そのPDFファイルをスライド形式で表示する方法が分からず、ImageMagickを使って画像ファイルに変換したものを表示するという茨の道を選択してしまいました。

[azu/pdf.js-controller](https://github.com/azu/pdf.js-controller)により、PowerPointファイルをPDFファイルに変換するところまでを考えれば良くなったため、Javascript方面へ深く足を踏み入れなくて済みました。

# おわりに

dockerを使って開発してみるのは初めてでしたが、一度動かせればインフラを簡単に立て
たり壊したりできるので、アプリを書くことに集中できるようになっていると思います。

* * *

これを作る時に試した、Gradleを使ったnpm/gruntビルドや、Dockerイメージの作成について、下記にまとめています。
[Gradleでbowerを利用したフロントエンドのビルドまでを行う方法 (1/6) - Qiita](http://qiita.com/kaakaa_hoe/items/0e3c622ae8bde3d116ad)

