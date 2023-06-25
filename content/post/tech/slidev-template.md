---
title: "Slidev関連のファイルやスライド資料をGitHubリポジトリで管理する"
emoji: "👨‍💻"
published: true
date: 2023-05-21T15:30:00+09:00
toc: true
tags: ["slidev", "github","githubactions","githubpages","githubreleases"]
---


# はじめに

Markdownkからプレゼンテーションを作成できる[Slidev](https://sli.dev/)というVue/Viteベースのフレームワークがある。2年ほど前に日本でも少し話題になっていた。

* [開発者のためのスライド作成ツール Slidev がすごい](https://zenn.dev/ryo_kawamata/articles/introduce-slidev)
* [Slidevを導入してMarkdownで美しいスライドを書こう \- Qiita](https://qiita.com/loftkun/items/2fbeddc9449eb5d85dfd)

[reveal.js](https://marp.app/), [Marp](https://marp.app/)なども触ったことはあったが、[Slidev](https://sli.dev/)プロジェクトを作成した時に同時に生成されるサンプルスライドがスタイリッシュなので、今でも愛用している。

Slidevでいくつも資料を作っていると、Slidevの原稿やビルドしたスライド資料をどこかに管理しておきたくなる。Slidevは、基本的には1スライドが1Web Applicationのような立て付け（だと思う）ので、Slidev資料の原稿をGitHubで管理する場合、1スライド1リポジトリのような形になると思う。(npm workspaceでモノレポっぽく管理するのもアリだとは思うが)

Slidevの原稿といってもMarkdownと多少の画像ファイルがあるぐらいなので、GitHubの1つのリポジトリ内で複数スライドの資産を管理したいなと思っており、せっかくGitHubで管理するならActionsでSlidevのビルドまで実行し、GitHub Pages/Releasesでビルドしたスライド資料の管理もできるようにしてみたという話。

# できたもの

GitHubのTemplate Repositoryとして作成した。(個人用途のものを勉強がてらテンプレートリポジトリにして見たものなので、諸々不具合が出ることも予想される)

https://github.com/kaakaa/slidev-resources-template

このテンプレートリポジトリを使ってSlidev資産を管理しているリポジトリ↓。

https://github.com/kaakaa/slidev-resources


## ディレクトリ構成

スライドごとにディレクトリを作成し、Markdown原稿(`slides.md`)や画像ファイルなどをそのディレクトリ内で管理する形式。

```
.
├── .github/         # GitHub Actionsなど
├── example-slide  　# Slidev原稿
│   ├── pages/
│   ├── public/      # 画像ファイルなど
│   └── slides.md    # Markdownファイル名は`slides.md`固定
├── awesome-slide/   # 別のSlidev原稿
│   └── slides.md
└── package.json     # slidevへの依存等はここにまとめられている
```

この方式だと、全てのSlidev資料に対して共通の`package.json`を使ってバージョン管理・ビルドスクリプト管理をしているため、今後Slidevのバージョンをあげた際に過去のSlidev資産をビルドできなくなる可能性はある。  
[npm workspaces](https://docs.npmjs.com/cli/v7/using-npm/workspaces)を使うことで、スライドごとにSlidevのバージョンを指定してモノレポ的に管理することもできたとは思うが、そうするとchild packageを作成するたびにSlidevビルド用のnpm scriptを編集する必要が出てきそうだったので辞めた。(npm workspacesでは、parent packageで定義したscriptをchild packageから呼び出すことは出来なかったと認識している。)  
また、スライド作成が完了したら、ビルドして成果物はPages/Releaseで管理されるため、後から資産をビルドし直すということもあまりないかとも思っているため、今の管理方式でも問題にはならないと考えている。（最悪、Gitの履歴からビルドが成功した時に使っていたSlidevのバージョンを調べることもできるはずだし）

## ビルド実行

ディレクトリ名と同名のTagを打つと、指定したディレクトリ内のSlidev資産をビルドしてGitHub PagesとGitHub ReleasesにSPA形式とPDF形式のスライドをそれぞれデプロイしてくれる。

* [GitHub Pages(SPA)](https://kaakaa.github.io/slidev-resources/)
* [GitHub Releases(PDF)](https://github.com/kaakaa/slidev-resources/releases)

図にするとこんな感じ

![structure](https://raw.githubusercontent.com/kaakaa/slidev-resources-template/main/assets/structure.png)

## 使い方

[テンプレートリポジトリ](https://github.com/kaakaa/slidev-resources-template)から新たにリポジトリを作成するだけだが、GitHub PagesによるSPA管理を実施するには、いくつかステップが必要。(PDF形式のSlidev資料だけを管理する場合は、テンプレートからリポジトリを作成してTagを打つだけで良いと思う。ただ、GitHub ActionsのSPAビルド/デプロイ関連のステップでエラーが出ると思うので、不要なステップの削除は必要そう(試してない))

GitHub PagesでSPA形式のSlidev資料を管理する場合は、以下の2つの手順が必要。

### 1. gh-pagesブランチを用意する

SPA形式でビルドされたSlidev資料は`gh-pages`へデプロイするようGitHub Actionsが組んであるので、`gh-pages`ブランチを用意しておく必要がある。テンプレートリポジトリの方に`gh-pages`ブランチを用意してあるので、テンプレートリポジトリからリポジトリを作成する際に`Include all branches`にチェックを入れるだけで良い。

自分で`gh-pages`ブランチを作ってPushしても良いが、その場合は`gh-pages`ブランチをorphanブランチとして切っておいた方が良い。

### 2. slidev buildコマンドの`--base`オプションの値を変更する

`slidev build`でSPA形式のビルドを実行する際、`--base`オプションにより出力されるSPAのbaseパスを指定することができる。テンプレートリポジトリでは`--base`オプションの指定はしているが、この内容を自分のリポジトリ向けに書き換えておかないと、SPAをGitHub Pagesにデプロイしても、スライドを表示することができない。

実際に編集が必要なのは、[`package.json`](https://github.com/kaakaa/slidev-resources-template/blob/b04259db5c82da58e5b9dce3e22c8a5af062353a/package.json#L5)内の`build`スクリプトで指定されている `--base /slidev-resources-template/${npm_config_slide}` の`slidev-resources-template`部分を、自分のリポジトリ名に置き換えるだけ。

この修正を含む、リポジトリのセットアップを行うPRを作成するWorkflowを用意してあるので、テンプレートから作成したリポジトリでこれを実行する方が早い。(GitHub ActionsからPull Requestを生成するには、リポジトリの`Settings > Actions > General > Allow GitHub Actions to create and approve pull requests`にチェックを入れる必要がある)

![](https://raw.githubusercontent.com/kaakaa/slidev-resources-template/main/assets/run-setup-workflow.png)

Setup workflowを実行すると、こんなPull Requestが作成される。
[chore: setup repository by github\-actions · Pull Request \#1 · kaakaa/slidev\-resources](https://github.com/kaakaa/slidev-resources/pull/1/files#diff-7ae45ad102eab3b6d7e7896acd08c427a9b25b346470d7bc6507b6481575d519)

# おわりに

Slidevの資産を成果物まで含めて一つのリポジトリで管理できるようになったので、スライド資料の管理にあまり頭を使う必要がなくなったように感じる。  
スライドで使う画像ファイルもリポジトリで管理するところがちょっと気持ち悪いが、画像管理のために他サービス使い出すと頭への負荷が上がってしまうので、目下のところはそこが悩みどころ。

あと、Tag消した時に関連する資産を削除したりはしてくれないので、その辺りなんとかしたい。あとOGP周りも。
