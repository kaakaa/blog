
---
title: "Mattermostで投票機能もどき"
date: 2017-04-02T21:09:59+09:00
draft: false
toc: true
tags: ["Go","Mattermost"]
---

(2017/4/4 Golang Driverについての記述を追加しました)

## はじめに

Mattermostのスラッシュコマンドで簡単に投票メッセージを作成できるアプリを作りました。
[kaakaa/matterpoll-emoji: Poll server for Mattermost](https://github.com/kaakaa/matterpoll-emoji)

Mattermostで`poll`というスラッシュコマンドを作り、下記のようにコマンドを実行

```
/poll `What do you gys wanna grab for lunch?` :pizza: :sushi: :fried_shrimp: :spaghetti: :apple:
```

すると、下記のような投票メッセージを投稿してくれます。
![matterpoll-emoji.png](https://qiita-image-store.s3.amazonaws.com/0/9891/192794f0-6f90-8371-6f1b-352fcdbd54a0.png)

べんり。

Botに定期的に出欠確認メッセージを呟かせるなどすると良いかと思います。
Mattermost v3.7で動作確認しています。

## 使い方

1. `git clone`する
2. `config.json`を書く
3. matterpoll-emojiを起動する
4. Mattermostでカスタムスラッシュコマンドを作成する
5. スラッシュコマンドを実行する

です。
詳細は[README](https://github.com/kaakaa/matterpoll-emoji/blob/master/README.md)を参照ください。


## 問題

* Mattermostでは0票のReactionを付けられないため、必ず最初に一票は入った状態になってしまう
  * Botアカウント作成推奨です
* Mattermost API Version4未対応
  * API v4はまだβバージョンなので、API v3で作っています。2017/9月にAPI v3が使えなくなるので、それまでにAPI v4に対応する必要があります。

## スラッシュコマンドについて

Marrermostで`/`で始まる投稿はスラッシュコマンドだと解釈されます。
Mattermostの投稿画面で`/`を入力すると、内臓スラッシュコマンドの一覧が表示されます。

![slash_command.png](https://qiita-image-store.s3.amazonaws.com/0/9891/853cd433-efd1-7b20-ed13-5ebb749954ad.png)


### カスタムスラッシュコマンド

ユーザー独自のスラッシュコマンドを作成することもできます。
システムコンソールでカスタムスラッシュコマンドを有効にし、統合機能からスラッシュコマンドを作成できます。

![システムコンソール.png](https://qiita-image-store.s3.amazonaws.com/0/9891/3aca649f-8aeb-272e-e97e-de93b766fd85.png)

![統合機能.png](https://qiita-image-store.s3.amazonaws.com/0/9891/42c5d435-02de-f056-1188-73574dee532a.png)

### リクエスパラメータ

Mattermostのスラッシュコマンド実行時に送られるリクエストパラメーターは下記になります。

[Creating Integrations with Commands](https://docs.mattermost.com/developer/slash-commands.html#creating-integrations-with-commands)

> 1. Your integration should have a function for receiving HTTP POSTs or GETs from Mattermost that look like this example:
> 
> ```
Content-Length: 244
User-Agent: Go 1.1 package http
Host: localhost:5000
Accept: application/json
Content-Type: application/x-www-form-urlencoded
> 
> channel_id=cniah6qa73bjjjan6mzn11f4ie&
channel_name=town-square&
command=/somecommand&
response_url=not+supported+yet&
team_domain=someteam&
team_id=rdc9bgriktyx9p4kowh3dmgqyc&
text=hello+world&
token=xr3j5x3p4pfk7kk6ck7b4e6ghh&
user_id=c3a4cqe3dfy6dgopqt8ai3hydh&
user_name=somename
```

---

BodyがURLパラメータ形式で送られてくるので、今回作った`matterpoll-emoji`ではダミーのURLのURLパラメータとして結合し、Goの`url`パッケージで解析する方法を取っています。
[poll_request.go#18](https://github.com/kaakaa/matterpoll-emoji/blob/master/poll/poll_request.go#L18)

### レスポンスパラメータ

[Creating Integrations with Commands](https://docs.mattermost.com/developer/slash-commands.html#creating-integrations-with-commands)

リクエストを受け取ったサーバーは、下記のような`application/json`形式の情報を返す必要があります。

> 3 . If you want your integration to post a message back to the same channel, it can respond to the HTTP POST request from Mattermost with a JSON response body similar to this example:
> 
> ```
{
  "response_type": "in_channel",
  "text": "This is some response text.",
  "username": "robot",
  "icon_url": "https://www.mattermost.org/wp-content/uploads/2016/04/icon.png"
}
```

`response_type`が`in_channel`の場合、普通の投稿と同じように投稿が作成されます。`ephemeral`を設定した場合はシステムメッセージのような、スラッシュコマンドを実行した人にしか見えない投稿が作られます。

`text`にはMattermostに投稿するメッセージテキストを指定しますが、Mattermostでの投稿と同じようにMarkdownやEmojiを利用することができます。

`username`と`icon_url`は投稿のユーザー名とアイコンを上書きするオプション情報です。

### Golang Driver

本来、スラッシュコマンドからリクエストされるサーバーの処理は、上記のレスポンスパラメータを返す形になりますが、それでは投稿にReactionを付けることなどができないため`matterpoll-emoji`では[Golang Driver](https://docs.mattermost.com/developer/api.html#golang-driver)を利用して投稿の作成・Reactionの付与をしています。

[model.NewClient](https://github.com/mattermost/platform/blob/master/model/client.go#L75)によって生成されるclientオブジェクトに投稿、ユーザー、チャンネルなどに対する操作が定義されています。

`matterpoll-emoji`ではGolang Driverの[ログイン、投稿、リアクションの操作を利用しています](https://github.com/kaakaa/matterpoll-emoji/blob/master/poll/poll_func.go#L43)。

## 開発の背景

弊社では週に一回ゆるい感じで勉強会を行っているのですが、人数があまり多くないため、聴講者が０人の日があるかもしれないという恐怖に常にさらされています。MattermostやWikiなどで出席・欠席などを報告できる場はあるのです、毎週の報告など面倒であるため報告機能がうまく回っているとは言えず。「ちゃんと出欠入力しましょう」なんて言うのも野暮。

そんな時、Mattermost v3.6で[Emoji リアクション](http://qiita.com/kaakaa_hoe/items/8148638982b86b1304ce#emoji-%E3%83%AA%E3%82%A2%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3)の機能が追加され、Slackのように投稿に絵文字でのリアクションを付けることができるようになりました。これを使って、毎週勉強会の日にBotに出欠確認メッセージを投稿させれば出欠登録も確認も楽になるのではないか、というのがきっかけです。

また、投票機能はMattermostのFeature Ideaでも人気の高い機能であるため、なるべくにいろいろな用途で使えるようにしました。[Polls (like doodle or for decision making) – Mattermost Feature Proposal Forum](https://mattermost.uservoice.com/forums/306457-general/suggestions/11721996-polls-like-doodle-or-for-decision-making)

## 結果

前回の勉強会出席者は一人でした。
0人回避。やったね。

