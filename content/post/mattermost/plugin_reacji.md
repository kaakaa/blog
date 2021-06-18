---
title: "Mattermost Reacji Channeler plugin"
date: 2021-05-02T23:05:21+09:00
draft: false
toc: true
tags: ["mattermost", "plugin", "reacji"]
---

## Mattermost Reacji Channeler Plugin

Mattermost 上で動く[Reacji Channeler](https://slack.com/intl/ja-jp/help/articles/360000482666-Slack-%E7%94%A8%E3%83%AA%E3%82%A2%E3%82%AF%E5%AD%97%E3%83%81%E3%83%A3%E3%83%B3%E3%83%8D%E3%83%A9%E3%83%BC)的なプラグインを作った。
まだ動くところまで作ったというレベルで、まだまだ破壊的な変更も必要そうな気がするけど、作ってる最中にいろいろ悩んでた部分を書き残しておきたいということで書いておく。

使い方等については README で。どういう風に動作するかについても、README にスクリーンキャストを載せてる。

[kaakaa/mattermost\-plugin\-reacji](https://github.com/kaakaa/mattermost-plugin-reacji#mattermost-plugin-reacji)

Slack の Reacji Channeler が数年前に日本で話題になってた時から面白そうな機能だなと思っていて、Mattermost v5.30 から[ReactionHasBeenAdded](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ReactionHasBeenAdded)というリアクションが付与された時に処理を差し込める Hook が追加されたので、勢いで作ってみた。（その少し前から WebSocket API 使った形である程度の形はできてたけど）

## Overview

![](https://blog.kaakaa.dev/images/posts/mattermost/plugin-reacji/overview.dio.png)

悩んだ箇所は下記で、それぞれのセクションで細かなところを書き残している。

- **(1) SlashCommand** - Plugin KeyValue Store の Key 長制約について
- **(2) (PluginHook) ReactionHasBeenAdded** - 絵文字エイリアスについて
- **(3) Store shared info** - Plugin KeyValue Store の Key 長制約について

### (1) SlashCommand

Reacji Channeler Plugin を有効にすると、`/reacji`というスラッシュコマンドが使えるようになる。
プラグインごとのスラッシュコマンドの実装は、[`API.reagisterCommand`](https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand)でコマンドを登録して、コマンド実行時の処理を[`Hooks.ExecuteCommand`](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ExecuteCommand)に記述する。

`/reacji add` というサブコマンドが実行されると、新たな Reacji を`ReacjiList`というインスタンスに格納する。その他にも参照や削除系のサブコマンドがあるが、使い方の話は[README](https://github.com/kaakaa/mattermost-plugin-reacji)で。

スラッシュコマンドのサブコマンドの実装についてはは`switch`文で地道に書く必要がある。  
以下のコードのようにスラッシュコマンドにサブコマンド的な実装を追加することもできるけど、これはスラッシュコマンド実行時の候補提示用の機能のため、実行時の処理についてはサポートしてくれない。

```go
func createAutoCompleteData() *model.AutocompleteData {
	suggestions := model.NewAutocompleteData("reacji", "[command]", "Available commands: add, list, remove, remove-all, help")
	suggestions.AddCommand(model.NewAutocompleteData("add", ":EMOJI: ~CHANNEL", "Register new reacji. If you attach EMOJI to the post in any channels except for DM/GM, the post will share to CHANNEL."))
	suggestions.AddCommand(model.NewAutocompleteData("add-from-here", ":EMOJI: ~CHANNEL", "Register new reacji. If you attach EMOJI to the post in the channel where this command is executed, the post will share to CHANNEL."))
	suggestions.AddCommand(model.NewAutocompleteData("list", "[--all]", "List reacjis in this channel. With `--all` list all registered reacjis in this server."))
	suggestions.AddCommand(model.NewAutocompleteData("remove", "[DeleteKey...]", "[CREATOR or SYSTEM_ADMIN only] Remove reacji by DeleteKey. You can see `DeleteKey` by `/reacji list`"))
	suggestions.AddCommand(model.NewAutocompleteData("remove-all", "", "[SYSTEM_ADMIN only] Remove all reacjis in this server."))
	suggestions.AddCommand(model.NewAutocompleteData("refresh-caches", "", "[SYSTEM_ADMIN only] Delete all caches. Reacji plugin caches data about shared post for a certain period in order to prevent duplicate sharing."))
	suggestions.AddCommand(model.NewAutocompleteData("help", "", "Show help"))
	return suggestions
}
```

https://github.com/kaakaa/mattermost-plugin-reacji/blob/3fe45b50563c912b783a0f120c18faec82e47ae9/server/plugin/command.go#L418

登録された Reacji をどう管理するかで、いろいろ悩んだ。  
Reacji は`ReacjiList`というまとまった形で管理するより、登録された Reacji 1 つが KeyValue Store の 1 データの対応するように作るのが無難だとは思うけど、Plugin 用に用意されている KeyValue Store は[Key の最大長が 50 文字](https://github.com/mattermost/mattermost-server/blob/df695115be82e09c27d9b54de754889599b0bf20/model/plugin_key_value.go#L13)までという制限があるため、Mattermost のチャンネル ID が 26 文字、[カスタム絵文字名の最大長が 64 文字](https://github.com/mattermost/mattermost-server/blob/df695115be82e09c27d9b54de754889599b0bf20/model/emoji.go#L14)という 2 つの情報を上手く 1 つの Key で表現することができなかった。
具体的なシーンとしては、あるチャンネルの投稿に絵文字が付与された場合、この情報を元に KeyValue Store からデータを取り出したいが、その場合、Key が最短でも`${CHANNEL_ID}${EMOJI_NAME} (26+64=90文字)`となってしまい、Key の最大長を超えてしまう。チャンネル ID と絵文字名からハッシュ値を作って管理することも考えたが、チャンネル ID と絵文字が一緒でも、共有先のチャンネルが違う Reacji というのも存在することがあり、そのようなデータの CRUD をどう管理すれば良いかいい案がなかったため、それならプラグイン起動時に全 Reacji を保持する`ReacjiList`のようなインスタンスを作って管理した方が I/O 少なくなって良いかと思った。その分、メモリを圧迫するので、プラグインの設定として登録できる Reacji の最大数を指定できるようにしてある。

そんな作りなのでマルチクラスタ構成で複数インスタンスで Mattermost が動いていると意図しない動きをしそうな気がする。マルチクラスタの時に KeyValue Store はどういう同期方法になっているんだろうか....

### (2) (Plugin Hook) ReactionHasBeenAdded

冒頭でも触れたけど、Mattermost v5.30 から追加された投稿にリアクションが付けられた時に実行される処理を登録することができる Hook。
[ReactionHasBeenAdded](https://developers.mattermost.com/extend/plugins/server/reference/#Hooks.ReactionHasBeenAdded)

この Hook 自体はわかりやすいものだったけど、Mattermost の絵文字エイリアスの扱いにちょっとハマったので少し書く
Mattermost デフォルトの絵文字リアクションには、1 つの絵文字に対して本来の絵文字名とは別の名前(絵文字エイリアス) を持っている絵文字がいくつか存在する。例えば、Mattermost の絵文字リアクション選択ダイアログで、`:uk:`と入力すると、`:gb: :uk:` という 2 つの名前で同じ絵文字を参照できることがわかる。ここで、この絵文字を選択すると `:gb:` でリアクションされたことになる。

![](https://blog.kaakaa.dev/images/posts/mattermost/plugin-reacji/emoji-alias.png)

Reacji の観点からすると、`/reacji add :uk: ~CHANNEL` という感じで`:uk:`を使った Reacji を登録し、絵文字選択ダイアログから`:uk:`を入力しても、実際に指定されるのは`:gb:`のため、`:uk:`で登録された Reacji は反応しないことになる。

絵文字エイリアスのデータは[mattermost-webapp に登録されており](https://github.com/mattermost/mattermost-webapp/blob/master/utils/emoji.json)、このデータから絵文字エイリアスについてまとめたのが下記のドキュメント。`aliases`列にある絵文字はあまり使わない方が良さそう。

https://github.com/kaakaa/mattermost-plugin-reacji/blob/master/notes-alias.md

絵文字エイリアスは Mattermost デフォルト絵文字のみの機能で、カスタム絵文字では使えないのでカスタム絵文字の方は気にする必要はない。

---

さらに話をややこしくすると、Mattermost 上で`+:uk:`というような投稿を行うと、直前の投稿に`:uk:`の絵文字リアクションを行うという機能があり、これで絵文字リアクションをすると、絵文字エイリアスは機能せず、`:uk:`という絵文字でリアクションできる。

![](https://blog.kaakaa.dev/images/posts/mattermost/plugin-reacji/emoji-alias-reaction.png)

また、Mattermost の画面上からは存在が確認できないが、実際は使用できる絵文字というのもあり、[Mattermost 上で本当に使用できる全絵文字のデータは mattermost-server 内に書かれている](https://github.com/mattermost/mattermost-server/blob/master/model/emoji_data.go)。例えば、`:older_man_medium_light_skin_tone:`などは Mattermost 上で補完などが効かないため存在を確認することはできないが、実際に投稿してみると存在することがわかる。

![](https://blog.kaakaa.dev/images/posts/mattermost/plugin-reacji/blind-emoji.png)

とりあえず絵文字エイリアスには気をつけてというぐらいしかない。

### (3) Store shared info

1 度 Reacji によって共有されたメッセージを再度同じチャンネルに共有しないようにするという話。絵文字が付けられるたびに同じ投稿が何度もチャンネルに共有されるのは邪魔ですからね。

[Clicking on the emoticon re\-triggers the shared from · Issue \#1 · kaakaa/mattermost\-plugin\-reacji](https://github.com/kaakaa/mattermost-plugin-reacji/issues/1)

すでに Reacji によって投稿が行われたかどうかは、Reacji の情報、投稿の PostID などから判断する必要があるため、これも`(1)`と同様に Plugin KeyValue Store の Key 長の制約が絡んでくる。ただ、今回はある条件を満たすデータが存在するかどうか(投稿が共有済みかどうか)を判断するだけでよく、Key に一意性もあるため、共有済みかどうかチェックするのに必要な情報からハッシュ値を求めて管理することにした。

Reacji によって共有されたという情報は KeyValue Store 上に格納されるが、デフォルトでは 30 日間経ったら格納された情報が消去されるようにしてある。30 日も前の投稿にリアクションすることはほぼ無いだろうという想定。この 30 日という期間については、プラグインの設定で変更することができる。また、システム管理者限定で共有されたという情報を削除する`reacji refresh-caches`というサブコマンドも入れてある。

とりあえず機能としては達成できたように見えるけど、もっと細かな要求があるような気がしている。

## まとめ

リアクション扱う Hook が実装されたので勢いで作ってみたけど、やっぱり作り始めるといろんな制約によって窮屈になってくる。KeyValue Store の Key 長さの制約とか絵文字エイリアスとかは、実際に不具合に当たらないと真剣に見ようとは思わないところなので、やっぱりいろいろ動かしてみないと見えてこない部分はたくさんありますね。
