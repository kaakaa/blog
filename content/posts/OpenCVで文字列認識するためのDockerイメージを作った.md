
---
title: "OpenCVで文字列認識するためのDockerイメージを作った"
date: 2016-07-02T16:05:00+09:00
draft: false
toc: true
tags: ["OpenCV","docker","tesseract-ocr"]
---

# 作ったもの

[https://hub.docker.com/r/kaakaa/opencv-contrib-python3/](https://hub.docker.com/r/kaakaa/opencv-contrib-python3/)

# OpenCVで文字列認識

画像処理ライブラリのOpenCVで文字列認識をする場合、OpenCV Contorib[Itseez/opencv_contrib](https://github.com/Itseez/opencv_contrib)のtextモジュールを通じて[tesseract-ocr](https://github.com/tesseract-ocr)を利用するのが定番のようです。
[opencvで文字認識その１　Tesseractラッパ - whoopsidaisies's diary](http://whoopsidaisies.hatenablog.com/entry/2014/11/12/003100)

OpenCVの機能で画像中の文字列の存在する位置を特定し、その特定したそれぞれの位置に対してTesseract-ocrの文字列認識を実行する模様。

日本語を認識する場合、Tesseractで用意されている日本語トレーニングデータが使えそうです（まだ使ったことは無い）。
[tesseract-ocr の言語データ(jpn.traineddata)について（その1） - 古都のIT職人Blog](http://a244.hateblo.jp/entry/2015/08/24/001916)


日本語文字列認識ならNHocrというのもあるらしい。
[日本語OCRライブラリNHocrを利用してみる - Qiita](http://qiita.com/awakia/items/3e1c7eb7da39e64de3a6)

# OpenCV Contrib付きDockerイメージ

OpenCV自体は[公式でプリビルドバイナリが配布されていたりする](http://opencv.org/downloads.html)ので、環境構築自体はすぐにできる。

ただ、Contribのモジュールを使用する場合、公式ではバイナリが配布されておらず、さらに文字列認識を利用する場合はTesseractをインストールする必要があったりするので、自分でソースからビルドする必要がありそう。


そこら辺の手順が面倒なので、OpenCV + OpenCV Contrib + Tesseract + Python3環境付きのDockerイメージを作りました。

[https://hub.docker.com/r/kaakaa/opencv-contrib-python3/](https://hub.docker.com/r/kaakaa/opencv-contrib-python3/)
[kaakaa/docker-opencv-contrib-python3: Dockerfile for OpenCV, OpenCV contrib and Python 3.5](https://github.com/kaakaa/docker-opencv-contrib-python3)

元々、こちらのDockerイメージ[austriker/docker-opencv-contrib-python3](https://hub.docker.com/r/austriker/docker-opencv-contrib-python3/)を使おうとしましたが、Tesseractが入ってなかったので、そのあたりを修正しました。

# 試しに実行

[docker-opencv-contrib-python3/compose-sample at master · kaakaa/docker-opencv-contrib-python3](https://github.com/kaakaa/docker-opencv-contrib-python3/tree/master/compose-sample)の方に、文字列認識用のWebアプリが立ち上がるdocker-composeを用意しました。
（Webアプリとしても文字列認識としても、まだまだ使い物レベルにはならないですが）

実際に実行してみた結果が下記画面。

![sample.png](https://qiita-image-store.s3.amazonaws.com/0/9891/4a39539e-8c03-cfb0-48bd-4b2d5dbc8dac.png)


下の画像のピンクで囲まれた部分が、OpenCVによって特定された文字列位置なのですが、画像全体として認識されてしまっているよう。
一行一行認識してほしいのですが。

そこらへんはいろいろチューニングしていく必要がありそうです。

