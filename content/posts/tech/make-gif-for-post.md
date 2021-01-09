---
title: "ブログ用のgifファイルを作るコマンドの備忘録"
date: 2021-01-09T22:15:00+09:00
draft: false
toc: true
tags: ["gif", "ffmpeg"]
---

何かの解説記事を書く時とかに、デスクトップをキャプチャした操作動画などを載せたいことがある。
そんな時は、普段Macを使っているので `Shift + Cmd + 5` のショートカットで動画を撮り、生成された.movファイルをffmpegで適当に.gifファイル化していたんだけど、どうも生成した.gifのサイズがMB単位になっていたようで、圧縮したいなと思った。

↓の記事を参考に試行。
[ffmpegでとにかく綺麗なGIFを作りたい \- Qiita](https://qiita.com/yusuga/items/ba7b5c2cac3f2928f040)

そしてできたのがこちらのコマンド。

```bash
$ NAME=incident
$ ffmpeg -i ${NAME}.mov -filter:v "setpts=PTS/2.0,scale=trunc(iw/2)*2:trunc(ih/2)*2" ${NAME}_refine.mov
$ ffmpeg -i ${NAME}_refine.mov -filter_complex "[0:v] fps=10,scale=640:-1 [r];[r] split [a][b];[b]fifo[bb]; [a] palettegen [p];[bb][p] paletteuse" ${NAME}.gif
```

1つ目の`ffmpeg`コマンドでは.movファイルに手を加えている。等速の操作動画だとモタモタしすぎなので `setpts=PTS/2.0` で倍速にしているのと、サイズを2の倍数にしている(はず)。サイズを2の倍数にするのは.gif化するときに **width not divisible by 2** というエラーで怒られたから（だったはず。記憶曖昧）。

2つ目の`ffmpeg`コマンドで、1つ目のコマンドで出力した.movファイルを.gifファイル化。

途中で **Buffer queue overflow, dropping.** という警告メッセージが出ている場合、最終成果物の.gifファイルが途中でフリーズした感じで生成されてしまうので注意。一応、これの対処も入ったコマンドのはずだけど。
