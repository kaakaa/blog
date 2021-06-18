
---
title: "ChatGopsで本物のGopherを、すぐそばに..."
date: 2016-08-29T21:19:05+09:00
draft: false
toc: true
tags: ["Go","gopher","Mattermost"]
---

# ChatOps

DevOpsの流れで、開発・運用における情報共有やビルド実行・通知などをコミュニケーションツールであるチャットから行うChatOpsというブームが起こりました。ChatOpsを採用することで開発チームと運用チームのリアルタイムな連携や、ビルド・デプロイの失敗に対する迅速な対応が可能となり、開発スピードの向上を図ることができるようになります。

その中核をなすのが、チャットに書き込まれたメッセージによる処理実行やチャットへの通知を管理するBotツールです。
ChatOpsの流行り始めはGithub社のHubotがデファクトスタンダートでしたが、昨今ではSlackのBotkitやMicrosoft Bot Framework、またLINEまでもBotを提供するなど、Bot / ChatOpsはシステム開発における一つの定番プラクティスとなってきました。

でも、何かが・・・何かが足りない・・

# ChatGops

そう、Gopherくんが足りない。

![bopher.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/bf9f9fef-de6c-41e1-50c9-f76368da4e1b.gif)

[kaakaa/bopher](https://github.com/kaakaa/bopher)

[本物の golang を\.\.\. 本物の Gopher を、お見せしますよ。 \- Qiita](http://qiita.com/mattn/items/b7889e3c036b408ae8bd)を全面的に利用しています。
（そのためWindows Onlyです）

## 必要なもの

* Mattermost 3.3～
* [mattn/gopher](https://github.com/mattn/gopher)をビルドした`gopher.exe`
* [kaakaa/bopher: Bot \+ Gopher = Bopher](https://github.com/kaakaa/bopher)

## 手順

1. Mattermost(v3.3以降)を立てる
2. [mattn/gopher](https://github.com/mattn/gopher)eをビルドする
3. [bopher.exe](https://github.com/kaakaa/bopher/releases/tag/v0.0.1)をダウンロードする
4. [bopher/config\.json](https://github.com/kaakaa/bopher/blob/master/config.json)に設定を書いて、bopher.exeと同じディレクトリに置く
5. bopher.exeを起動する

# 注意

* Core i7, メモリ24GBを積んだマシンでもGopherくんを5,60人同時に発生させたところ操作が効かなくなりました。設定ファイルでGopherくんの最大数を設定できるようにしてあります。用法・容量を守って正しく服用ください。

# きっかけ

* Mattermost3.3から、GoでBotを書けるようになったので、なんか作ってみようと思いました
  * [mattermost/mattermost\-bot\-sample\-golang](https://github.com/mattermost/mattermost-bot-sample-golang)

