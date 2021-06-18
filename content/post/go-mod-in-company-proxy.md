---
title: "Go Mod in Company Proxy"
date: 2020-04-03T00:04:58+09:00
draft: false
---


### 追記

単純に`http_proxy`と`https_proxy`を設定してないだけだった...。

(追記ここまで)

Goをv1.14.1にアップデートし、`go mod tidy`を実行したところエラーが出た。

```
$ go mod tidy
go: github.com/mattermost/mattermost-server/v5@v5.21.0: Get "https://proxy.golang.org/github.com/mattermost/mattermost-server/v5/@v/v5.21.0.mod": dial tcp 172.217.26.113:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```

メッセージから会社のプロキシに引っかかってるんだろうなーという予想はたったが、対処法がわからない。ググって `GOPROXY=direct`を付ければよさそうと思い、再度実行してみたがダメ。

```
$ GOPROXY=direct go mod tidy
go: github.com/mattermost/mattermost-server/v5@v5.21.0/go.mod: verifying module: github.com/mattermost/mattermost-server/v5@v5.21.0/go.mod: Get "https://sum.golang.org/lookup/github.com/mattermost/mattermost-server/v5@v5.21.0": dial tcp 216.58.196.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```

ただ、エラ〜メッセージは変わっているのでもう少し調べるとこんな情報。

[Go モジュールのミラーリング・サービス【正式版】 — プログラミング言語 Go \| text\.Baldanders\.info](https://text.baldanders.info/golang/mirror-index-and-checksum-database-for-go-module/)
> チェックサム・データベース

なるほど、確かに二番目のログは `sum.golang.org` に接続しようとしている。というわけで、`GOSUMDB=off` も付けてチェックサムデータベースへのアクセスも取りやめれば良い、と。

```
$ GOPROXY=direct GOSUMDB=off go mod tidy
go: github.com/mattermost/mattermost-server/v5@v5.21.0 requires
go.uber.org/atomic@v1.5.1: unrecognized import path "go.uber.org/atomic": https fetch: Get "https://go.uber.org/atomic?go-get=1": dial tcp 216.58.199.243:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```

はい。`go.user.org`に接続できないとなりました。コレはプロキシ経由しないと無理だなぁ...。どうすれば...。
