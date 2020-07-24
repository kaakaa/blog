---
title: "子供の見守りカメラをちょいホームオートメーション"
date: 2020-07-23T12:19:05+09:00
draft: false
toc: true
tags: ["automation","life"]
---

## 背景

数ヶ月前ぐらいから息子を寝室に寝かせるようになったのですが、寝室は2階にあり、リビングなどは1階にあるという状況なため、リビングにいても息子が起きたら気付けるようにと見守りカメラ的な物を購入して使っていました。

寝室にはあまり物がないため、[専用のクリップ](https://amzn.to/3eTiYmy)を一緒に購入し、窓枠に噛ませる形で設置をしていたのですが、子供が掴まりをし始めるとカメラの設置位置まで手が届い、設置位置を見直す必要が出てきました。

設置位置としては、子供が届かない高さであり、かつ電源が供給できる位置ということで、照明用のダクトレールのところに移動しようと思っていたのですが、せっかくなのでホームオートメーション的なことをしてみようと思いました。

## デバイス

### 見守りカメラ

カメラは[PlanexのCS-QS10](https://www.planex.co.jp/products/cs-qs10/index.shtml)を自宅のネットワークに参加させて、公式スマホアプリで映像を見る感じ。電気消して暗くなると自動で赤外線モードになってくれ、赤外線モードでも子供の動きが観測できる程度の明るさになるので便利。ただ、映像見る時にデバイスの音声が奪われてしまうので、音楽聴きながら映像見たりできない点がちょっとネック。

<!-- スマカメ -->
{{< iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=kaakaahoe1193-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B07J4F1GCR&linkId=32d1f01fa5d9cbbc0f2acb125185b5a1" >}}


### 電源 (SwitchBotプラグ)

電源部分については、まず[適当なダクトレール用のコンセント](https://amzn.to/3g7n5Nj)を購入し、そのコンセントに以前購入して使っていなかった[SwitchBot スマートプラグ](https://amzn.to/3fVim10)を挿しました。

<!-- SwitchBotプラグ-->
{{<iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=kaakaahoe1193-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B07TVSGJT4&linkId=265b9a044b9553f66025110a4b89f586" >}}

これで、[SwitchBotアプリ](https://apps.apple.com/jp/app/switchbot/id1087374760)経由で電源のON/OFFやスケジューリングなどができるようになります。設置方法としてはこんな感じ。

![スマカメ](https://blog.kaakaa.dev/images/posts/life/home-automation-smacame.png)

レールを這わせることでコードを見えなくすることもできそうですが、家族以外入らない寝室なので横着しました。

とりあえずココまででもスマホからカメラのON/OFFができ、かつ高さも取れているので十分なのですが、お遊びでiPhoenixのNFC機能を使ってみることにしまいた。

### NFCタグ

少し前に↓の記事を読み、興味を持っていました。(Androidではだいぶ昔からできていたようですが)
[iPhoneとNFCタグを使った家電制御が便利すぎる！【Siriショートカット】 \- CHASUKE\.com](https://chasuke.com/nfctag/)

[SwitchBot スマートプラグ](https://amzn.to/3fVim10)のON/OFF機能を持ったNFCタグを寝室の扉に貼っておき、子供を寝かせた後、出る時にピッとすると見守りカメラの電源をONにし、寝室にいくときにまたピッとするとカメラの電源がOFFになるイメージ。

とりあえずNFCタグを購入。

<!-- NFC タグシール -->
{{< iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=kaakaahoe1193-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B00GXSGL5G&linkId=656babae51c4d8104e6663d30c01c013" >}}

パッケージには "**Android**" としか書いておらず、「アレ？」と思って[公式サイト](https://www.sanwa.co.jp/product/syohin.asp?code=MM-NFCT)を見に行ったらハッキリと「iPhoneは非対応です」と書いてあり一瞬焦りましたが、まだiPhoneにNFC機能が付く前の話だったようで一安心。

> ※iPhone、iPadは非対応です。（2014年6月現在）

(2014年...相当昔の商品だったのか...)

設定後、ドアノブのところにNFCタグを貼り、さて動作確認と行こうとしたら動かず...どうやら金属製の部分に貼ると反応してくれないようで。[公式サイト](https://www.sanwa.co.jp/product/syohin.asp?code=MM-NFCT)にもハッキリ書いてありましたね...。

> ※アルミなどの金属面には対応していません。

貼りなおそうとしたらシール部分とNFC部分？が分離してしまい使い物にならなくなったため、また新しいNFCタグで設定し直し。今度はドアの木の部分に貼り動作することを確認。

![NFCタグ](https://blog.kaakaa.dev/images/posts/life/home-automation-nfc.png)

とりあえずこれで見守りカメラのON/OFFの操作をiPhoneをかざすだけでできるようになりました。NFCタグの使い勝手も分かり、ショートカットアプリ使用の実績も解除しました。

## おわりに

照明用のダクトレールに見守りカメラを設置すると、広い範囲を見れるので便利。（ダクトレールに取り付ける関係上、カメラの上下が逆転してしまってる）

![shot](https://blog.kaakaa.dev/images/posts/life/home-automation-shot.png)

NFCタグとショートカットの組み合わせは便利そうだけど、今のところ他に良いユースケースは思い付かず。まぁ、思いついたらすぐ試せる材料が揃えられたのは良かったかな。
