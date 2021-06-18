
---
title: "打ち合わせ みんな多忙で 草生えるwww"
date: 2017-12-21T13:49:26+09:00
draft: false
toc: true
tags: ["Node.js","ChatOps","会議","Mattermost","ews"]
---

字（草）余り。

---

この記事は[FUJITSU Advent Calendar 2017 \- Qiita](https://qiita.com/advent-calendar/2017/fujitsu) 22日目の記事です。
記事の内容は全て個人の見解であり、執筆内容は執筆者自身の責任です。所属する組織は関係ありません。

# はじめに

打ち合わせの調整を依頼されて、みんな多忙で入れる隙間も無いと草生えますよね。

**生やしましょう。**

# Mattermostでスケジュール可視化&調整

[去年に引き続き](https://qiita.com/kaakaa_hoe/items/7fd95171f38dc3178276)Mattermostネタです。
(去年は技術記事じゃなかったことを反省しながら)

MattermostはMattermost Inc.社が開発しているSlack AlternativeなOSSのチャットツールです。
https://about.mattermost.com/
https://github.com/mattermost

---

Mattermostの[Interactive Message Button](https://docs.mattermost.com/developer/interactive-message-buttons.html)使って何か作りたくて、スケジュール可視化&チャットから会議予約なモノを作りかけました。

[kaakaa/matter\-meeting: \[PoC\] Sending meeting request to Exchange from Mattermost](https://github.com/kaakaa/matter-meeting)

実用的なところまで行き着けなかったので表現弱めに。コンセプト実装とお受け取りください。

## 概要

こんな感じで動きます。

![user_genarated_image.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/ef605519-f446-3103-3d22-9c6dfbaf9366.gif)


Mattermostから打ち合わせ参加者のメアドを指定すると、現在から1週間後までの参加者全員のスケジュールをGitHubの草風に表示。さらにはInteractive Message Buttonから会議予約を飛ばすこともできたりします。

Exchangeから予定を取得にはEWSのGetUserAvailabilityというAPIを利用しており、このAPIが全員の予定の中から空いている時間を`quality`というフィールド付きで返してくれます。[SuggestionQuality \| Ews JavaScript Api](http://ews-javascript-api.github.io/api/enums/enumerations_suggestionquality.suggestionquality.html)。
`quality`はどうも0～3までの整数で表されるようなので、その値をそのまま草の濃淡として表示しています。

草の色は白に近いほど皆の予定が空いており、緑が濃いほど予定が詰まっていることを示しています。

つまり、**みんな忙しいと草生えるwww**
（深夜帯が草だらけなのは実装の都合上です。要検討事項。）


## 構成

ざっくりこんな構成で動いてます。

![Architecture.png](https://qiita-image-store.s3.amazonaws.com/0/9891/0caf8739-e4cd-1b17-cf15-7eab5c953f3f.png)

Slash Commandからリクエストを受けると、Exchange Serverへ問い合わせをし、そのレスポンスからMattermostへのBot投稿と草用のSVG画像を生成。
Bot投稿のボタンを押すと、Exchangeへ会議予約のリクエストが送られるという感じ。

Exchangeサーバーに繋ぐ部分にews-javascript-apiを使ってるのに引きずられて、慣れないNode.js使ってるのでPromiseの使い方がとても怪しい。
（そして型がない言語だからか、コード書いててもmodelが生えない…草は生えるのに…）

# 要素技術の紹介

## Mattermost

### Custom Slash Command

Mattermostでは自作の[Slash Commands](https://docs.mattermost.com/developer/slash-commands.html#custom-slash-command)を登録できます。
少し前にこれを使って[投票機能](https://qiita.com/kaakaa_hoe/items/b2605ce3816cfc517ecd)みたいなのも作ってました。

今回は自作Slash Commandsサーバーのレスポンスに[Messsage Attachments](https://docs.mattermost.com/developer/message-attachments.html)を使用し、ちょっとリッチなフォーマットで反応できるようにしています。

Message Attachmentsは、Slash Commandsのレスポンスとして下記のようなJSONをMattermostに返すと

```json
{
  "attachments": [
    {
      "fallback": "test",
      "color": "#FF8000",
      "pretext": "This is optional pretext that shows above the attachment.",
      "text": "This is the text of the attachment. It should appear just above an image of the Mattermost logo. The left border of the attachment should be colored orange, and below the image it should include additional fields that are formatted in columns. At the top of the attachment, there should be an author name followed by a bolded title. Both the author name and the title should be hyperlinks.",
      "author_name": "Mattermost",
      "author_icon": "http://www.mattermost.org/wp-content/uploads/2016/04/icon_WS.png",
      "author_link": "http://www.mattermost.org/",
      "title": "Example Attachment",
      "title_link": "http://docs.mattermost.com/developer/message-attachments.html",
      "fields": [
        {
          "short": false,
          "title":"Long Field",
          "value":"Testing with a very long piece of text that will take up the whole width of the table. And then some more text to make it extra long."
        },
        {
          "short":true,
          "title":"Column One",
          "value":"Testing"
        },
        {
          "short":true,
          "title":"Column Two",
          "value":"Testing"
        },
        {
        "short":false,
        "title":"Another Field",
        "value":"Testing"
        }
      ],
      "image_url": "http://www.mattermost.org/wp-content/uploads/2016/03/logoHorizontal_WS.png"
    }
  ]
}
```

こんな感じの投稿を作ってくれます。

<img width="723" alt="attachments-example.png" src="https://qiita-image-store.s3.amazonaws.com/0/9891/b05ada9b-1dcd-2278-30df-50019b8194cd.png">

(参考: [Message Attachments — Mattermost 4\.5 documentation](https://docs.mattermost.com/developer/message-attachments.html#example-message-attachment))

メッセージが5行以上の時は省略表示もしてくれるのも良い感じです。
最近、Mattermostのコアチームを真似てGHEのアクティビティをMattermostに流したりしてるのですが、コレを使うと長いIssueでも良い感じに表示してくれるので気に入ってます。

### Interactive Message Button

ChatOps黎明期のBotはメッセージを投稿するだけでしたが、最近のBotはメッセージに選択肢を付けてくれます。
Slackには昔からある機能みたいで、[GolangでSlack Interactive Messageを使ったBotを書く \- Mercari Engineering Blog](http://tech.mercari.com/entry/2017/05/23/095500)を見ながら面白そうだな～なんて思ってたら、Mattermostにもやってきました。（[PLT\-6403: Interactive messages by ccbrown](https://github.com/mattermost/mattermost-server/pull/7274)）。
Mattermostでは今のところボタンのみの対応ですが。

![poll.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/20500a8d-92bf-f3cc-13fd-8607185bb859.gif)
（参考: [Interactive Message Buttons \(Beta\) — Mattermost 4\.5 documentation](https://docs.mattermost.com/developer/interactive-message-buttons.html)）

今回は、調整した予定の中でもっとも参加可能人数が多い時間をInteractive Message Buttonとして表示しています。多すぎると訳分からなくなりそうだったので最大１０個まで表示。
ボタンを押すと会議出席依頼が飛ばせるというアレ。べんり。

![meeting2_0.png](https://qiita-image-store.s3.amazonaws.com/0/9891/f193c0f2-fd3d-6eaa-21ba-b62b395864af.png)

Interactive Message Buttonの作り方は、先述のMessage Attachmentsのレスポンスの中に[`integration`](https://docs.mattermost.com/developer/interactive-message-buttons.html#button-options)というフィールドとして指定することで作れます。

んで、ボタンを押すとhttpリクエストが飛ばされるというアレ。夢が広がりますね。はい。

ボタンだけでも十分遊べそうですが、セレクトボックスとかチェックボックスなんかも加わると更に夢が広がりそうだなぁ。まとまった時間とってコントリビュートしたいなぁ。

## Exchange

予定の取得や会議予約を飛ばしたりするのは[ews-javascript-api](https://github.com/gautamsi/ews-javascript-api)を使っています。
昔は[ews-java-api](https://github.com/OfficeDev/ews-java-api)を使って遊んでいましたが、歳のせいかJavaを書くのが辛くなってきました。

ews-javascript-apiはユースケース情報が転がってない気がするので、こんな感じで増やしていけると良いですね。

ただ、これ便利なんですが、気を許すと全社員宛のリクエストを飛ばすことになりかねなかったりするので、指定できるメアドにはホワイトリスト方式を採用しています。**マジでここだけ注意な。**

一応設定ファイルで正規表現指定できるけど、仲間内のメアドだけリストで指定しておくのが安全です。ews-javascript-api使って自分の部門だけに制限するとかも出来る…のかな？

## 草生やす

サーバー側でnunchucks テンプレートにスケジュール情報を流し込み、SVG作ってOSSのS3互換オブジェクトストレージである[Minio](https://minio.io/)に投げてます。
ここで生成した草画像はMattermostのMessage Attachmentsの[`image_url`](https://docs.mattermost.com/developer/message-attachments.html#images)に指定して表示させてます。

最初はsvg2png使って画像ファイル化してたけど、Message Attachmentsで指定したURLは[`img`タグで読み込んでる](https://github.com/mattermost/mattermost-webapp/blob/master/components/post_view/post_attachment.jsx#L303)ようなので、SVGのままでも行けるじゃん、と後で気づきました。

## Minio

生成したSVGは、上述のようにMinioに格納してURL指定で取得できるようにしています。
スラッシュコマンドを実行するごとに[shortid](https://www.npmjs.com/package/shortid)でユニークIDらしき物を作って、そのIDをファイル名としてSVGを保存してます。

SVGはローカルファイルとして格納しても良いかとも思ったのですが、なんとなくMinioを使ってみたかったので。
一定時間経ったSVGは削除するソリューションも欲しくなってます。

Minioだと少しヘビーな気がするのと、ランダム文字列とはいえ誰でも見えるところに調整結果の画像を置くのは少し忍びない気もしていますが、まぁとりあえず。

## テストは？

まだ無い

## コード品質は？

嘲笑するならパッチくれ

## 問題

実用に向けた最大の課題は、ews-javascript-api使うために指定したアカウントからしか会議出席依頼を飛ばせないところ。みんなで使うと私のスケジュールが凄いことになりそう。
Bot社員が求められている。

あと細々した改善点もたくさん。

* 30分間の予定しか想定されてない
* 現在から1週間後までの間しか調べてない(日時指定できるようにするとコマンドが複雑になるのがアレ)
* メアド毎回入力するのが面倒 (Mattermostの@channel使うとチャンネル内全員対象にする、とか出来ると理想)

などなど。
今回用のネタとして開発したものなので突き詰めるつもりもないですが(たぶん)

## Mattermostについて

最近、UberがSlack/HipChatを諦めてMattermostベースのコミュニケーションプラットフォームである`uChau`というのを作り始めているというニュースがあり。まだまだユーザー数を伸ばしていきそうな感じがします。
[The Road to uChat: Building Uber’s Internal Chat Solution](https://eng.uber.com/uchat/)

機能的にも、この記事でも取り扱ったInteractive Message Buttonや、[Plugins機能](https://docs.mattermost.com/administration/plugins.html)なんかも追加されてきているので、まだまだいろいろ遊べるようになっていく予感。

開発も活発なのでこれからも期待しています。

## おわりに

お仕事ではみんなで大草原回避しましょう。

