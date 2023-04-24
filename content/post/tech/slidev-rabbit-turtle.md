---
title: "Slidevでプレゼン時間管理のためにウサギとカメを表示する"
date: 2023-04-23T23:30:00+09:00
draft: false
toc: true
tags: ["slidev", "rabbit", "presentation"]
---

# はじめに

Rubyistの方々御用達(?)のプレゼンテーションツールに[Rabbit](https://rabbit-shocker.org/ja/)というものがある。  
Rabbit自体は使ったことはないんだけど、プレゼン画面にうさぎとかめを表示して、最初に設定したプレゼン想定時間と残りページ数の状況を可視化するという方式は何度か目にしたことがある。

最近、tmtmsさんがruby.wasmを使ってFirefoxでPDFを表示する際にうさぎとかめを表示するというものを作っているのをみて、自分が最近よく使っているSlidevでも同じようなことができないかと思い立ち、ちょっと作ってみた。

[Rabbit on Firefox \- tmtms のメモ](https://blog.tmtms.net/entry/202303-rabbit-on-firefox)

# できたもの

(2023/04/24追記)
[Slidev Addon](https://sli.dev/addons/use)としてnpm公開した。  
https://www.npmjs.com/package/slidev-addon-rabbit

---

[kaakaa/slidev\-rabbit\-turtle: Presentation time management for slidev inspired rabbit](https://github.com/kaakaa/slidev-rabbit-turtle)

ビルドしたスライド資料をGitHub Pagesで公開している。  
https://kaakaa.github.io/slidev-rabbit-turtle/?time=1

![screen shot](https://blog.kaakaa.dev/images/posts/tech/slidev-rabbit-turtle/screen.gif)

プレゼンを開始すると(表紙から2ページ目以降へ移動すると)、スライド資料下部にあるウサギとカメのアイコンが右に向かって移動する。  
ウサギはページを遷移するごとに移動し、最後のスライドが表示されるとウサギは一番右まで到達する。  
カメは時間経過と共に移動し、URLクエリ`time`で指定した時間(分)が経過すると一番右まで到達する。  
ウサギが右端に到達する前にカメが到達すると、時間オーバーということになる。

![description](https://blog.kaakaa.dev/images/posts/tech/slidev-rabbit-turtle/description.png)

# 使い方

SlidevはMarkdownからスライドを作成できるツール。  
公式ドキュメントはコミュニティによって日本語化されているし、ググれば解説資料はたくさん出てくるので、Slidevに関する説明は省略する。

* [Home \| Slidev](https://sli.dev/)
* [Slidevを導入してMarkdownで美しいスライドを書こう \- Qiita](https://qiita.com/loftkun/items/2fbeddc9449eb5d85dfd)

ウサギとカメの表示には、Slidevの[Global Layers](https://sli.dev/custom/global-layers.html)を利用している。これは、ヘッダーやフッターなどのスライド全ページに表示されるコンテンツを指定する機能。ウサギとカメはスライド下部に表示するので、フッターを表示するための`globa-bottom.vue`に記述している。

ウサギとカメの表示に必要なのは `components/`, `global-bottom.vue`, `style.css`。

```
.
├── components
│   ├── Flag.vue
│   ├── Rabbit.vue
│   └── Turtle.vue
├── global-bottom.vue
├── package-lock.json
├── package.json
├── pages
│   └── multiple-entries.md
├── slides.md
└── style.css
```

`global-bottom.vue`でスライドのフッターを宣言し、フッターに表示する内容は`components/`内のVueコンポーネントで定義している。また、`style.css`は表示位置の指定等に利用している。　　
Slidevプロジェクトにこれらのファイルを追加すれば、表示できるようになると思う。  

ファイルを追加した後に`slidev`コマンド等でスライドを開始し、スライドURLの末尾に`?time=10`等のプレゼン時間(分)を設定すると、ウサギとカメが表示される。`time`クエリがないとウサギとカメを表示しないようにしている。  
また、`slidev export`でPDF出力する際はアイコンがエクスポートされないようにしている。PDFでウサカメしたい場合は、冒頭で紹介したtmtmsさんの[Rabbit on Firefox](https://blog.tmtms.net/entry/202303-rabbit-on-firefox)などを使うと良いと思う。

# おわりに
毎回Slidevプロジェクトを作るたびにファイルをコピーするのは面倒なので、[Slidev Addon](https://sli.dev/addons/use.html)にできると使いやすくなるのかもしれない。(Addon機能を今初めて知ったので、そういうことに使えるのかどうかもわからないけど。)


