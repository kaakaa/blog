---
title: "Mattermost EmojiGen plugin"
date: 2021-03-20T15:05:21+09:00
draft: false
toc: true
tags: ["mattermost", "plugin", "emoji"]
---

## Mattermost EmojiGen

Mattermost 上でテキストから絵文字を作成できるプラグインを作ってた。  
下記にスクリーンキャストを載せてる。

[kaakaa/mattermost\-emojigen: Creating custom emoji from slash command](https://github.com/kaakaa/mattermost-emojigen#usage)

2 年前ぐらいに作ってたやつだけど、久しぶりに引っ張り出してゴリゴリと。以前、何かができなくて諦めてた気がするけど、そんなに大きな問題なく手当て完了。

主な変更点は、今まで 2x2 の 4 文字までの絵文字しか作れなかったけど、上限文字数を撤廃。3x3 の`圧倒的当事者意識`的な絵文字も作れるようにした。とはいえ、視認性的に 3x3 でも厳しいぐらいだと思うけど。

Go 言語で文字画像作るのは、文字数と描画位置をどのように決めるかとの戦い。テストコード書いて画像出力させながらゴリゴリ調整。脳筋。

https://github.com/kaakaa/mattermost-emojigen/commit/21fd13ada28106ca1b7fcd3d0279cfef1e90ed11#diff-285e170643d9e80db83915466e06d2269d66a6cd59a3140e9db79cf5b6d7f834R47

### No bot...

Mattermost プラグインでは、Plugin 起動時にプラグイン専用の Bot アカウントを作成して、その Bot から[Plugin API](https://developers.mattermost.com/extend/plugins/server/reference/#API)を叩かせて処理をやらせるというのが一般的だけど、悲しいことに絵文字登録 API が[Plugin API](https://developers.mattermost.com/extend/plugins/server/reference/#API)として実装されていなかったので、プラグインの設定画面にアクセストークンを入力して、絵文字登録は REST API 経由で実行するという方式を取っている。

そのため、セットアップ手順がちょっと面倒臭くなっている。REST API のなかで、Plugin API に実装されているものとされていないものの違いがわからないんだよな。Contribute しますって言えばやらせてもらえるのかな。

### CircleCI

久しぶりに CircleCI 回したらビルドがコケる。

```
...
Warning: Permanently added the RSA host key for IP address '140.82.112.3' to the list of known hosts.
Permission denied (publickey).
fatal: Could not read from remote repository.
...
```

Deploy Key を作り直せば良かったらしい。

[CircleCI で checkout 時に「fatal: Could not read from remote repository\.」出る（長く放置していた時に発生） \- nwtgck / Ryo Ota](https://scrapbox.io/nwtgck/CircleCI%E3%81%A7checkout%E6%99%82%E3%81%AB%E3%80%8Cfatal:_Could_not_read_from_remote_repository.%E3%80%8D%E5%87%BA%E3%82%8B%EF%BC%88%E9%95%B7%E3%81%8F%E6%94%BE%E7%BD%AE%E3%81%97%E3%81%A6%E3%81%84%E3%81%9F%E6%99%82%E3%81%AB%E7%99%BA%E7%94%9F%EF%BC%89)

### GitHub Release

tag 打っても Release Job が走らない。

基本、Mattermost Plugin 関連のファイルは [mattermost\-plugin\-starter\-template](https://github.com/mattermost/mattermost-plugin-starter-template) から取って来るけど、このリポジトリにある Circle の設定ファイルには Release Job が定義されていなかった。

使い慣れた？ Matterpoll の方の設定ファイルを持ってくる。

https://github.com/matterpoll/matterpoll/blob/master/.circleci/config.yml

`mattermost/plugin-ci`の Circle CI Orb は、最新バージョン`v0.1.6`まで進んでるんだけど、新しめのやつ(`v0.1.2`とかだったかな?)になると Mattermost 社の S3 インスタンスとかにアクセスしようとするので、認証情報設定しておかないとエラーになるのよね。なので一般コントリビュータは`v0.1.0`を使うしかない（と思う）。
ここら辺、ちゃんとドキュメント化されてないから手探りなのよね。

## まとめ

Go で画像生成できるのが面白くて作り始めた気がするけど、まぁとりあえず形にはなったかな。ちゃんとテスト書いてないし、ちょっと操作性が微妙な気はするけど。

言葉の絵文字が流行るのは日本的、というか漢字文化だよなぁ。
