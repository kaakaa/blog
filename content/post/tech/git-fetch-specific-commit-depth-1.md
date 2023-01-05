---
title: "go-gitで特定のコミットをDepth=1でCloneする"
emoji: "💡"
published: true
date: 2023-01-05T23:40:00+09:00
toc: true
tags: ["git", "tips", "go-git"]
---


## 背景

大きなGitリポジトリをCloneする際に `--depth <depth>` オプションを使用すると、指定した数だけの履歴を持つShallow Copyを取得するようになり、Cloneにかかる時間を短縮することができる。

しかし、 `HEAD` ではないある特定のコミット時点のリポジトリを取得したい場合、 `git clone` コマンドにはCommit Hashを指定してリポジトリを取得するオプションがないため、結局 `git clone` の `--depth` オプションは使用できず、全ての履歴を取得した上で `git checkout <sha1>` などでリポジトリの履歴を戻す必要がある。(例えば、リポジトリマイニング的な事をするためにGitHub APIでリポジトリの各Release時点のコミットIDは分かっている、という状況を想定)

## 解決方法

手順は増えるが、 `git fetch` ならば `--depth` を使用でき、かつCommit Hashを指定して特定のコミット時点のリポジトリの情報を取得できるため、やりたい事を実現することができる。

```sh
$ git init
$ git remote add origin https://github.com/$OWNER/$NAME
$ git fetch --depth 1 origin $SHA1
$ git checkout FETCH_HEAD
```

参考: [20211018: git \- commit hash 値で shallow clone \- PIB](https://seesaawiki.jp/w/kou1okada/d/20211018%3A%20git%20-%20commit%20hash%20%C3%CD%A4%C7%20shallow%20clone)

## go-gitを利用したコード

`go-git` を使用してこれを実現しようとする場合、 `git fetch` に直接Commit Hashだけを指定することはできない(?)ため、Refspecを指定する形に変更する必要があった。

```go:main.go
package main

import (
    "log"

    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/config"
    "github.com/go-git/go-git/v5/plumbing"
)

const (
    url        = "https://github.com/matterpoll/matterpoll"
    sha string = "a62dbe0410aa0836cd4b26e75f9319a952ff153d"
)

func main() {
    // git init
    repo, _ := git.PlainInit("work", false)

    // git remote add origin https://github.com/owner/name
    remote, _ := repo.CreateRemote(&config.RemoteConfig{
        Name: "origin",
        URLs: []string{url},
    })

    // git fetch --depth 1 origin $SHA1:refs/heads/target
    var target = plumbing.NewBranchReferenceName("target")
    _ = remote.Fetch(&git.FetchOptions{
        RemoteName: "origin",
        RefSpecs: []config.RefSpec{
            config.RefSpec(plumbing.ReferenceName(sha) + ":" + target),
        },
        Depth: 1,
    })

    // git checkout target
    tree, _ := repo.Worktree()
    _ = tree.Checkout(&git.CheckoutOptions{Branch: target})
}
```

上記のGoコードを実行すると、1コミットのみを持つGitリポジトリが取得できている。(`go.mod` 等の記述は割愛)

```sh
$ go run main.go
$ cd work
$ git log
commit a62dbe0410aa0836cd4b26e75f9319a952ff153d (grafted, HEAD -> target)
Author: Yusuke Nemoto <kaakaa@users.noreply.github.com>
Date:   Mon Apr 29 05:31:07 2019 +0900

    Release 1.1.0 (#150)
$
```