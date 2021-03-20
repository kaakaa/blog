---
title: "[Mattermost Integrations] Plugin (Server hooks)"
date: 2020-12-18T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "server", "hooks"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第18日目の記事です。

本記事では、Mattermost上の様々な操作に対応した処理を追加できるMattermost Pluginの**Server**サイドの機能である**Server Hooks**について紹介します。

Mattermost Pluginについての公式ドキュメントは下記になります。
https://developers.mattermost.com/extend/plugins/overview/

## Mattermost Plugin (Server) について

**Server Hooks**の説明の前に、Mattermost Pluginの本体について紹介します。

Mattermostの**Server**サイドのPluginを実装する場合、[`plugin.MattermostPlugin`](https://github.com/mattermost/mattermost-server/blob/f2571099539ee6e432d59e5bb4bc390652426a0e/plugin/client.go#L39)を埋め込んだ構造体がPlugin本体となります。`plugin.MattermostPlugin`を埋め込んだ構造体は、`API`と`Helper`というフィールドを持ち、これらのフィールドを経由してMattermostのリソースを処理する様々なメソッドを呼び出すことができます。また、`plugin.MattermostPlugin`を埋め込んだ構造体に対して**Server Hooks**と同じインターフェースを持つメソッドを実装することで、**Server Hooks**を利用することができるようになります。

```go
package main

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type SamplePlugin struct {
	plugin.MattermostPlugin
}

// OnActivate Hooksの実装
func (p *SamplePlugin) OnActivate() error {
    // `API`フィールドを通じたPlugin APIの呼び出し
	if err := p.API.RegisterCommand(&model.Command{
		Trigger: "sample-command",
	}); err != nil {
		return errors.Wrap(err, "failed to register  command")
	}
    return nil
}
```

`main`メソッドで`plugin.MattermostPlugin`を埋め込んだ構造体を引数として`plugin.ClientMain`メソッドを呼ぶことで、プラグインを起動することができます。

```go
package main

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func main() {
	plugin.ClientMain(&SamplePlugin{})
}
```

## Server Hooks

Mattermost PluginのServer Hooksは、Mattermost上でユーザーがチャンネルに参加したときや、ユーザーがMattermostにログインしたときなど、何かのアクションに応じて実行される処理を追加できる機能です。
`plugin.MattermostPlugin`を埋め込んだ構造体に、Server Hooksと同じインターフェースを持つメソッドを実装することで利用可能になります。

Server Hooksの一覧は下記から確認できます。

https://developers.mattermost.com/extend/plugins/server/reference/#Hooks

### OnActivate

`OnActivate`はPluginが起動したときに呼ばれるHookです。Botを使うPluginの場合はこのHook内でBotユーザーを作成したり、Slash Commandを使うPluginならSlashCommandの登録などを行います。(Pluginから登録したSlash Commandは、通常の統合機能として作成したSlash Commandと違い外部アプリケーションにリクエストは送信されません。Slash Commandが実行された時の処理は、Server Hooksの`ExecuteCommand`で実装します。)

`error``が返却された場合は、プラグインが起動されません。

```go
func (p *SamplePlugin) OnActivate() error {
	// Botの登録
	bot := &model.Bot{
		Username:    "test-bot",
		DisplayName: "Sample Bot",
	}
	botUserID, appErr := p.Helpers.EnsureBot(bot)
	if appErr != nil {
		return errors.Wrap(appErr, "failed to ensure bot user")
	}
	p.botUserID = botUserID

	// Slash Commandの登録
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:      "sample",
		AutoComplete: true,
	}); err != nil {
		return errors.Wrap(err, "failed to register  command")
	}

    return nil
}
```

### Implemented

`Implemented`は、Pluginが実装しているHookの名前を返すためのHooksです。しかし、実装されているのを見たことがないので、用途はないかもしれません。

```go
func (p *MatterpollPlugin) Implemented() ([]string, error) {
    return []string{"OnActivate", "Implemented"}, nil
}
```

### OnDeactivate
`OnDeactivate`はプラグインが停止された時に実行されます。

```go
func (p *MatterpollPlugin) OnDeactivate() error {
    p.clean()
    return nil
}
```

### OnConfigurationChange
Plugin専用の設定が変更された際に実行されます。
Mattermost Pluginの[Manifestファイル](https://developers.mattermost.com/extend/plugins/manifest-reference/)に`settings`を記述することで、Plugin専用の設定画面を持つことができます。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/configuration-page.png)

`OnConfigurationChange`周りの処理は下記のStarterテンプレートのコードを流用すると良いです。
https://github.com/mattermost/mattermost-plugin-starter-template/blob/master/server/configuration.go

```go
type configuration struct {
	SampleSetting string
}

func (p *SamplePlugin) OnConfigurationChange() error {
	var configuration = new(configuration)

	// Load the public configuration fields from the Mattermost server configuration.
	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return errors.Wrap(err, "failed to load plugin configuration")
	}

    p.setConfiguration(configuration)
    
	return nil
}
```

### ServeHTTP

Mattermost Plugin専用のエンドポイントに対してリクエストが送信された時に実行されます。

Mattermost PluginにはPluginごとにエンドポイントが存在します。Mattermostが`https://example.com:8065`で起動していたとすると、`https://example.com:8065/plugins/{PLUGING_ID}`がPlugin専用のエンドポイントになります。ここに送られたリクエストを処理するのが`ServeHTTP`です。

```go
func (p *SamplePlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
```

[Interactive Message](https://docs.mattermost.com/developer/interactive-messages.html)などのリクエスト送信先をPlugin用のエンドポイントにするなどの利用方法があります。

### ExecuteCommand
Plugin APIの[`RegisterCommand`](https://developers.mattermost.com/extend/plugins/server/reference/#API.RegisterCommand)で登録されたコマンドが実行された時に処理されます。

```go
func (p *SamplePlugin) OnActivate() error {
	// Slash Commandの登録
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:      "sample",
		AutoComplete: true,
	}); err != nil {
		return errors.Wrap(err, "failed to register  command")
	}

	return nil
}

func (p *SamplePlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	return &model.CommandResponse{Text: "Hello by plugin"}, nil
}
```

### UserHasBeenCreated
`UserHasBeenCreated`は、ユーザーが新規に作成された場合に実行されます。
新しく参加したユーザーにBotからメッセージを送る場合などに使用できます。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/user-has-been-created.png)

```go
func (p *SamplePlugin) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	channel, appErr := p.API.GetDirectChannel(p.botUserID, user.Id)
	if appErr != nil {
		p.API.LogWarn("failed to get direct channel", "user_id1", p.botUserID, "user_id2", user.Id, "details", appErr.Error())
		return
	}
	if _, appErr := p.API.CreatePost(&model.Post{
		ChannelId: channel.Id,
		UserId:    p.botUserID,
		Message:   "Welcome to Our Mattermost!",
	}); appErr != nil {
		p.API.LogWarn("failed to create welcome post.", "channel_id", channel.Id, "details", appErr.Error())
	}
}
```

### UserWillLogIn 
`UserWillLogIn`は、ユーザーがログインする直前に実行されます。

空文字以外の文字列を返すとログインが取り消されますが、現在のバージョンでは返却した文字列はユーザーには表示されず、ユーザー名とパスワードが合っていても`Enter a valid email or username and/or password`というメッセージが表示されてしまうようです。さらに、プラグインを管理できるユーザーがログアウトしてしまうと、前記のエラーでログインできず、利用可能な状態に戻すのが困難になってしまうかもしれないため、使用には注意が必要そうです。

```go
unc (p *SamplePlugin) UserWillLogIn(c *plugin.Context, user *model.User) string {
	if err := p.check(); err != nil {
		return err.Error()
	}
	return ""
}
```

### UserHasLoggedIn
`UserHasLoggedIn`は、ユーザーがログインした直後に実行されます。

前回ログアウト(オフライン)してから、7日以上経過していた場合にBotからメッセージを送信する場合などに利用できます。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/user-has-logged-in.png)

```go
func (p *SamplePlugin) UserHasLoggedIn(c *plugin.Context, user *model.User) {
	status, appErr := p.API.GetUserStatus(user.Id)
	if appErr != nil {
		p.API.LogWarn("failed to get user status", "user_id", user.Id, "details", appErr.Error())
		return
	}
	t := time.Unix(status.LastActivityAt/1000, status.LastActivityAt%1000)
	if status.Status == model.STATUS_OFFLINE && time.Now().After(t.AddDate(0, 0, 7)) {
		channel, appErr := p.API.GetDirectChannel(p.botUserID, user.Id)
		if appErr != nil {
			p.API.LogWarn("failed to get direct channel", "user_id1", p.botUserID, "user_id2", user.Id, "details", appErr.Error())
			return
		}
		if _, appErr := p.API.CreatePost(&model.Post{
			ChannelId: channel.Id,
			UserId:    p.botUserID,
			Message:   "Hi! :wave:",
		}); appErr != nil {
			p.API.LogWarn("failed to create post.", "channel_id", channel.Id, "details", appErr.Error())
		}
	}
}
```

### MessageWillBePosted
`MessageWillBePosted`は、投稿されたメッセージがデータベースに保存される前に実行されます。投稿を拒否したり、投稿内容を自動で編集したい場合などに利用できます。投稿作成時に拒否や編集以外の処理を実行する場合は、投稿がデータベースに保存された後に実行される`MessageHasBeenPosted`の利用が推奨されています。

投稿を拒否する場合は、２つ目のreturn値に空でない文字列を指定します。一つ目の返却値の`*model.Post`には内容を編集した後の`*model.Post`を指定します。`nil`を指定した場合でも、引数で与えられた`*model.Post`が指定されたものと解釈されます。

```go
func (p *SamplePlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	if strings.Contains(post.Message, "shit") || strings.Contains(post.Message, "💩") {
		return nil, "You can't use `shit` and 💩 on this server."
	}
	return nil, ""
}
```

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/message-will-be-posted.png)

このHookによって投稿が拒否された場合、ユーザーからはその拒否理由が見えないようなので、拒否基準を明文化したり、拒否理由をBotから通知するなどの対応が必要そうです。

### MessageWillBeUpdated 
`MessageWillBeUpdated`は、投稿済みのメッセージを編集した際、編集内容がデータベースに保存される直前に実行される処理です。
`MessageWillBePosted`とほぼ同じ内容のため、例は省略します。

```go
func (p *SamplePlugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
    ...
} 
```

### MessageHasBeenPosted

`MessageHasBeenPosted`は、投稿がデータベースに保存された直後に実行される処理です。

特定のキーワードを含むメッセージが作成された場合に、特定のチャンネルに通知するようなコードは下記のようになります。
Botが作成した投稿もこのHookで処理されるため、考慮が漏れると処理が無限ループしてしまうため注意が必要です。また、非公開チャンネルの投稿なども処理されてしまうため、その点を考慮する必要もあります。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/message-has-been-posted.png)

```go
func (p *SamplePlugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	postUrl := fmt.Sprintf("http://localhost:8065/_redirect/pl/%s", post.Id)
	if strings.Contains(post.Message, "mattermost") && post.UserId != p.botUserID {
		p.API.CreatePost(&model.Post{
			Message:   fmt.Sprintf("Post refered to `mattermost` is created. See [here](%s) ", postUrl),
			UserId:    p.botUserID,
			ChannelId: "su7w9z51atnspjufg1c73ijx8w",
		})
	}
}
```


### MessageHasBeedUpdated
`MessageHasBeenUpdated`は、投稿済みのメッセージを編集した際、編集内容がデータベースに保存された直後に実行される処理です。
こちらも`MessageHasBeenPosted`とほぼ同じ内容のため、例は省略します。

```go
func (p *SamplePlugin) MessageHasBeenUpdated(c *plugin.Context, newPost, oldPost *model.Post) {
    ...
}
```

### ChannelHasBeenCreated

`ChannelHasBeenCreated`は、チャンネルが作成された直後に実行されます。

チャンネルが作成されたことを`town-square`チャンネルに通知するコードは下記のようになります。このHookについても、非公開チャンネルが作成された場合の考慮が必要になります。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/channel-has-been-created.png)

```go
func (p *SamplePlugin) ChannelHasBeenCreated(c *plugin.Context, channel *model.Channel) {
	if channel.Type != model.CHANNEL_OPEN {
		return
	}

	u, appErr := p.API.GetUser(channel.CreatorId)
	if appErr != nil {
		p.API.LogError("Failed to get user", "details", appErr)
		return
	}
	townSquare, appErr := p.API.GetChannelByName(channel.TeamId, model.DEFAULT_CHANNEL, false)
	if appErr != nil {
		p.API.LogError("Failed to get channel", "details", appErr)
		return
	}

	if _, appErr := p.API.CreatePost(&model.Post{
		Type:      model.POST_DEFAULT,
		ChannelId: townSquare.Id,
		UserId:    p.botUserID,
		Message:   fmt.Sprintf("Channel ~%s has been created by %s.", channel.Name, u.GetDisplayName(model.SHOW_USERNAME)),
	}); appErr != nil {
		p.API.LogError("Failed to create post", "details", appErr)
	}
}
```


### UserHasJoinedChannel
`UserHasJoinedChannel`は、ユーザーがチャンネルに参加した直後に実行されます。第３引数の`actor`は、他のユーザーがユーザーをチャンネルに追加した場合など、ユーザーをチャンネルに追加する処理を実行した人の情報が入ります。

チャンネルに新しく参加したユーザーに読んで欲しいリンクなどを通知する場合に利用できます。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/user-has-joined-channel.png)


```go
func (p *SamplePlugin) UserHasJoinedChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	if channelMember.ChannelId != TargetChannelID {
		return
	}
	p.API.SendEphemeralPost(actor.Id, &model.Post{
		ChannelId: channelMember.ChannelId,
		UserId:    p.botUserID,
		Message:   fmt.Sprintf("This chanels is for XXX user. You'd better to read [notes for this channel](%s).", UrlForNotes),
	})
}
```


### UserHasLeftChannel
`UserHasLeftChannel`は、ユーザーがチャンネルから脱退した直後に実行されます。
`UserHasJoinedChannel`とほぼ同じ内容のため、例は省略します。

```go
func (p *SamplePlugin) UserHasLeftChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
    ...
}
```

### UserHasJoinedTeam

`UserHasJoinedTeam`は、ユーザーがチームに参加した直後に実行されます。
`UserHasJoinedChannel`とほぼ同じ内容のため、例は省略します。

```go
func (p *SamplePlugin) UserHasJoinedTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User)  {
    ...
}
```

### UserHasLeftTeam
`UserHasLeftTeam`は、ユーザーがチャンネルから脱退した直後に実行されます。
`UserHasJoinedChannel`とほぼ同じ内容のため、例は省略します。

```go
func (p *SamplePlugin) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
    ...
}
```

### FileWillBeUploaded
`FileWillBeUploaded`は、メッセージ入力欄にファイルが添付された時に実行されます。ユーザーが投稿作成を実行する前にファイルが変換されます。

添付されたファイルの情報は第2引数の[`*model.FileInfo`](https://pkg.go.dev/github.com/mattermost/mattermost-server/v5/model#FileInfo)から、ファイルの内容は第3引数の`io.Reader`から取得できます。添付ファイルに変更を加えた場合は、第4引数の`io.Writer`に書き込みます。
また、ファイルの添付を拒否する場合は、2つ目の返却値に空でない文字列を指定します。1つ目の返却値の`*model.FileInfo`には内容を編集した後の`*model.FileInfo`を指定します。もしファイルを編集した場合、編集後のファイルサイズについては自動で更新されるため、`FileWillBeUploaded`内で計算する必要はありません。

画像にフィルタをかける例は下記のようになります。

![IMAGE](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day18/file-will-be-uploaded.png)

```go
func (p *SamplePlugin) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	if info.IsImage() {
		// Decode original image
		img, _, err := image.Decode(file)
		if err != nil {
			p.API.LogWarn("failed to decode uploaded image", "details", err.Error())
			return nil, ""
		}
		// Draw original image
		base := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
		draw.Draw(base, base.Bounds(), img, image.ZP, draw.Src)

		// Create green filter
		src := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
		draw.Draw(src, src.Bounds(), &image.Uniform{color.RGBA{255, 128, 255, 128}}, image.ZP, draw.Src)

		// Mask original image
		mask := image.Rect(25, 25, base.Bounds().Dx()-25, base.Bounds().Dy()-25)
		draw.DrawMask(base, base.Bounds(), src, image.ZP, mask, image.ZP, draw.Over)

		// Write masked image
		png.Encode(output, base)
	}
	return info, ""
}
```

画像ファイルのExifを削除するなどで利用することもできますが、処理自体はMattermostサーバーにファイルが送られた後に実行されるため、注意が必要です。


## さいごに

本日は、Mattermost PluginのServer Hooksについて紹介しました。
明日も、Mattermost Pluginの**Server**サイドで使用できるAPIとHelper関数について紹介します。
