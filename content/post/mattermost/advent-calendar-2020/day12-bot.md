---
title: "[Mattermost Integrations] Botアカウント"
date: 2020-12-12T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "bot"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第12日目の記事です。

本記事では、Mattermostで特定の用途向けのBotアカウントを作成する機能について紹介します。


## Botアカウントの概要
Mattermostの外部からMattermostのリソースを操作するようなアプリケーションについて考えてみます。

Mattermost外部から投稿を作成するだけであれば、WebHookやSlash Commandを利用すれば実現することができますが、これらの機能では投稿作成以外の処理を行うことができません。Mattermost REST APIを使うことでMattermost上のチャンネル・チーム・投稿など様々なリソースをMattermost外部から操作することができますが、REST APIを実行するにはユーザーに紐づいたアクセストークンが必要なため、統合機能用のサーバーアプリケーションなどを構築する場合、ユーザー情報の不正利用を防ぐためにトークンの厳密な管理などを意識する必要が出てきます。

このような場合に有効なのが**Botアカウント**機能です。

**Botアカウント**機能により特定の処理専用のユーザーアカウントを作成することができ、**Botアカウント**のアクセストークンを利用することで、ユーザー情報の不正な利用を心配することなくREST APIの実行などを行うことができます。

また、Enterpriseを利用している場合、Botアカウントは登録ユーザーとしてカウントされないため、統合機能専用のユーザーを作りやすくなります。

### 設定

Mattermostの画面上からBotアカウントを作成するには、**システムコンソール > Botアカウント > Botアカウントの作成を有効にする** の設定が`有効`になっている必要があります。

![system console](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/config-bot.png)

### 作成

**メインメニュー > 統合機能 > Botアカウント > Botアカウントを追加する**からBotアカウントを作成できます。

![bot menu](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/bot-menu.png)

Botアカウントの作成画面で入力する情報は下記の通りです。

![create bot](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/create-bot.png)

* **ユーザー名**: Botアカウントのユーザー名です。既存のアカウントと同じユーザー名は使用できません。
* **Botアイコン**: Botアカウントのアイコンをアップロードできます。
* **表示名**: Botアカウントの表示名です。
* **説明**: Botアカウント一覧画面に表示されるアカウントの説明です。
* **役割**: Botアカウントに割り当てられる役割を **メンバー**(一般ユーザー) と **システム管理者** から選択します。
* **投稿:全て**: 公開チャンネルだけでなく、非公開チャンネルやダイレクトメッセージチャンネルなどにもBotアカウントによる投稿を作成する場合、こちらにチェックします。
* **投稿:チャンネル**: Botアカウントによる投稿の作成が公開チャンネルだけで良い場合、こちらにチェックします。

**役割**、**投稿:全て**、**投稿:チャンネル**が実際にどのような効果があるかが分かりづらいですが、例えば、**役割**が**メンバー**で、**投稿:全て**にも**投稿:チャンネル**にもチェックがない場合、そのBotはチャンネルにメンバーとして追加されない限り、そのチャンネルに投稿を作成できないということになります。**投稿:チャンネル**が設定された場合は、公開チャンネルであればチャンネルにメンバーとして参加していなくてもそのチャンネルに投稿は作成できますが、非公開チャンネルやダイレクトメッセージチャンネルにはメンバーとして追加されるまで投稿を作成する権限がないということになります。

Botアカウントの作成が完了すると、このBotアカウントに割り当てられたアクセストークンが表示されます。このトークンはこの画面を閉じると二度と表示することができませんが、Botアカウントにはいくつもアクセストークンを追加することができるため、アクセストークンを忘れてしまった場合は再生しすることになります。

![complete bot](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/complete-bot.png)

作成されたBotアカウントはBotアカウント一覧画面に表示されます。

![list bot](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/list-bot.png)

### 実行

Botアカウントに割り当てられたアクセストークンを使用してREST APIを実行すると、BotアカウントによってMattermostのリソースを操作することができます。

```bash
BODY='{
  "channel_id": "89xmji6bibbn9eqpe1okx8j8fe",
  "message": "Create post by bot account"
}'

curl -i \
  -H 'Authorization: Bearer 3uhox91m6bdbt8pqsczouy9mny' \
  -H 'Content-Type: application/json' \
  -d "$BODY" \
  http://localhost:8065/api/v4/posts
```

![post from bot](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/post-from-bot.png)

また、Botアカウントはユーザーと同じようにチャンネルやチームに参加させることもできます。

![join channel](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/join-bot.png)

Botアカウント作成時に**投稿:全て**や**投稿:チャンネル**にチェックを入れなかった場合、このようにBotをチャンネルに参加させない限り、Botアカウントはそのチャンネルにアクセスできない（権限エラーで投稿が作成できない）ことになります。

また、Botアカウントのアクセストークンを使ってWebSocket APIをListenしている場合、Botを参加せているチャンネルのイベントだけが取得できるため、Botアカウントが参加しているチャンネルのみを対象とした処理を実装することもできます。Outgoing Webhookで指定できたのは1つのチャンネルのみ、Custom Slash Commandでは作成したチーム内の全チャンネルを対象としていましたが、Botアカウントの場合、Botアカウントをチャンネルさせるかどうかで柔軟に統合機能の対象チャンネルを設定できるようになります。

![websocket](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/websocket-bot.gif)

### Botアカウントの削除
Botアカウントを利用しなくなった場合、無効化はBotアカウントの一覧画面から実行できます。

![deactivate bot](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day12/deactivate-bot.png)

しかし、Botアカウントの削除はMattermost画面からは実行できません（できないはずです）。Botアカウントの削除を行うにはMattermostのCLIツールである `mattermost` コマンドを使用する必要があります。`mattermost`コマンドはMattermostを起動しているサーバーのMattermostインストールディレクトリ内の`bin/mattermost`という場所に存在するはずです。
Linux用のバイナリであればMattermostのダウンロードページからダウンロードできるMattermost本体に含まれていますが、`mattermost`コマンドはMattermostサーバーの設定ファイル`config.json`を読み込んで動作するため、Mattermostサーバー上で実行するのが無難です。

```bash
$ bin/mattermost user delete sample_bot
```

と、`mattermost user delete`コマンドの引数としてBotアカウントのユーザー名を指定することで、Botアカウントの削除を行うことができます。

## さいごに

本日は、Botアカウントについて紹介しました。
明日はMattermostの機能から一旦離れて、今までMattermostから受け取ったSwag(Contributors Gift)を紹介します。
