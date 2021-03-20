---
title: "新しくなったGitHub ActionsでHugoブログのデプロイをしてみた"
date: 2019-08-31T10:50:21+09:00
draft: false
toc: true
tags: ["GitHub", "GitHubActions", "github-pages", "Hugo"]
---

# はじめに

CI/CD機能が付いて新しくなったGitHub Actionsが使えるようになったので、[Travisでビルド/デプロイしていたブログ](https://github.com/kaakaa/blog)をActionsに移行してみました。

GitHub Actionsのベータプログラムには下記サイトから応募できます。
https://github.com/features/actions

# 前提

今まで利用していたTravisの設定ファイルはこちら。

```yaml
notifications:
  email: false

sudo: false
dist: trusty
language: go

env:
  - HUGO_VERSION=0.51

go:
- '1.10'

install: true

before_script:
  - wget https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.deb
  - sudo dpkg -i hugo_${HUGO_VERSION}_Linux-64bit.deb

script:
  - git submodule update --init
  - hugo -v

deploy:
  provider: pages
  local-dir: public
  skip-cleanup: true
  github-token: $GITHUB_TOKEN
  keep-history: true
  on:
    branch: master
```

Hugoのインストールとビルドを行い、`pages`のproviderを使って`gh-pages`ブランチにデプロイしています。(GoとHugoのバージョンが古いですね...むしろGoはいらないんじゃないような...)

# 移行

## GitHub Actionsの作成

Actionsを使いたいリポジトリのタブの部分に `Actions` というメニューが生えているため、そこをクリックします。

<img width="645" alt="スクリーンショット 2019-08-31 10.01.26.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/b1276252-5a63-2823-aeb7-33c41590be69.png">

まだActionsが設定されていない場合は、Workflowのサンプルがいくつか表示されています。
自分は何も分からなかったので右上の `Set up a workflow yourself` のボタンを選択しました。

<img width="1188" alt="スクリーンショット 2019-08-31 10.05.34.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/db8fde68-bc9c-f430-aa60-20efbf10f693.png">

すると、 `.github/workflows/main.yml` というファイルにサンプルのワークフローの定義ファイルを作成する画面に移ります。CircleCIの設定ファイルによく似てますね。

<img width="1012" alt="スクリーンショット 2019-08-31 10.07.06.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/b6c65c5b-fae2-72a7-af04-33c273b18684.png">


右ペインに簡単な説明が書かれているのでそれを見たり、[ドキュメント](https://help.github.com/en/categories/automating-your-workflow-with-github-actions)を見たりしながら編集していきます。


## 完成品 (main.yml)

とりあえず動くものが出来ました。

```yaml
name: Deploy Pages

on:
  push:
    branches:
    - master

jobs:
  deploy:
    runs-on: ubuntu-latest    
    steps:
    - name: Checkout
      uses: actions/checkout@v1
    - name: Build and Deploy
      uses: JamesIves/github-pages-deploy-action@master
      env:
        ACCESS_TOKEN: ${{ secrets.ACCESS_TOKEN }}
        BRANCH: gh-pages
        FOLDER: public
        HUGO_VERSION: 0.57.2
        BUILD_SCRIPT: |
          wget https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.deb
          dpkg -i hugo_${HUGO_VERSION}_Linux-64bit.deb
          git submodule update --init
          hugo -v

```

`Build and Deploy` というステップで、[JamesIves/github\-pages\-deploy\-action](https://github.com/JamesIves/github-pages-deploy-action)という他の人の作ったGitHub PagesをデプロイするActionsを使っています。CircleCIで言うところのOrbでしょうか（あまり知らない）。

### Secrets

`env: ACCESS_TOKEN:` のところで `secrets.ACCESS_TOKEN` と言う変数を使っていますが、これはリポジトリのSettingsメニューから設定できる変数を使っているようです。
**Settings > Secrets > Add a new secret**を選択し、変数名とその値を入れるとActionsの定義ファイル(`main.yml`)でその変数を使えるようになります。

詳しくは[公式ドキュメント](https://help.github.com/en/articles/virtual-environments-for-github-actions#creating-and-using-secrets-encrypted-variables)を。

<img width="1009" alt="スクリーンショット 2019-08-31 10.29.54.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/2910d27e-fd10-c161-d9e1-cfa1c7b98c2f.png">


### 微妙な点

このステップの `BUILD_SCRIPT` のところでHugoでのビルドを行っているのですが、ここでHugoのインストールまでここでやっているのがちょっと微妙です。別のステップを作って、その中でHugoのインストールを行ってみたりしたのですが、`hugo`コマンドが見つからないと言うエラーとなり、別のステップでの変更は持ち込めないようでした。
なんかやり方あるのかな？

## 動作確認

コミットを契機にビルドが開始される設定となっているので、`.github/workflows/main.yml` をコミットすると、Actionsが開始されます。結果はActionsタブから確認できます。ビルドログも見れます。

<img width="1003" alt="スクリーンショット 2019-08-31 10.33.42.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/9891/cd34afd7-6a2e-149a-f64b-7d648436a199.png">


# おわりに

深くは使ってないので判断はできませんが、GitHubだけで完結できるのは楽ですね。

今年の初めぐらいに[TravisがIderaに買収](https://blog.travis-ci.com/2019-01-23-travis-ci-joins-idera-inc)されたり、[Travisのエンジニアが解雇された](https://news.ycombinator.com/item?id=19218036)りとTravis周りはキナ臭い話が多く、それを機にCircleCIに移行するプロジェクトも多かったと思いましたが、そこでビッグウェーブに乗らずに移行しなかったために難なくGitHub Actionsに移れてよかったです。果報は寝て待てですね（動きが遅かっただけ）。
