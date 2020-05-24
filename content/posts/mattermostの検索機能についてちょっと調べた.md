---
title: "Mattermostの検索機能についてちょっと調べた"
date: 2020-05-24T15:16:57+09:00
draft: false
toc: true
tags: ["mattermost", "search"]
---

# Mattemrost の検索機能について少し調べた

この前、Mattermostnの検索機能について質問を受けました。
検索欄に `in:channel_name` と入力して、チャンネル指定だけして検索を実行するとチャンネル内の全てのメッセージが検索できるけど、そこから除外文字だけ設定する(`in:channel_name -excluded_term`)と検索結果が0件になってしまうとのこと。

2020/05/24時点の最新コミット (v5.24開発中)を元に見ていきます。
https://github.com/mattermost/mattermost-server/commit/2135096d886dad9055640edf881661004312c69b

## Mattermostの投稿検索

`search`機能の本体は[`store/sqlstore/post_store.go`内の`search`メソッド](https://github.com/mattermost/mattermost-server/blob/aad76a13e8381fc2a7c79d37d842d530c450a728/store/sqlstore/post_store.go#L1134)な模様。
[`model.SearchParams`](https://github.com/mattermost/mattermost-server/blob/master/model/search_params.go#L15)が検索条件を持つモデルで、この中の`Terms`、`ExcludedTerms`あたりが検索ワードに関するパラメータっぽい。

```go
type SearchParams struct {
	Terms                  string
	ExcludedTerms          string
	IsHashtag              bool
	InChannels             []string
	ExcludedChannels       []string
	FromUsers              []string
	ExcludedUsers          []string
	AfterDate              string
	ExcludedAfterDate      string
	BeforeDate             string
	ExcludedBeforeDate     string
	OnDate                 string
	ExcludedDate           string
	OrTerms                bool
	IncludeDeletedChannels bool
	TimeZoneOffset         int
	// True if this search doesn't originate from a "current user".
	SearchWithoutUserId bool
}
```

それぞれ一度変数に読み出しているので、こちらの変数を追っていく。

https://github.com/mattermost/mattermost-server/blob/aad76a13e8381fc2a7c79d37d842d530c450a728/store/sqlstore/post_store.go#L1200
```go
terms := params.Terms
excludedTerms := params.ExcludedTerms
```

[1217行目](https://github.com/mattermost/mattermost-server/blob/aad76a13e8381fc2a7c79d37d842d530c450a728/store/sqlstore/post_store.go#L1217)あたりで、これらの変数で分岐しているところがあったのでここらへんが怪しい。

```go
if terms == "" && excludedTerms == "" {
	// we've already confirmed that we have a channel or user to search for
	searchQuery = strings.Replace(searchQuery, "SEARCH_CLAUSE", "", 1)
} else if s.DriverName() == model.DATABASE_DRIVER_POSTGRES {
    ...
} else if s.DriverName() == model.DATABASE_DRIVER_MYSQL {
    ...
}
```

まず検索ワードが無い場合の処理をして、その後は使用しているDBごとに処理を分けている。

### 検索ワード/検索除外ワードが共に指定されていない場合

チャンネルだけ指定した検索(`in:channel_name`)はこれにあたりそう。
やっている処理は変数 `searchQuery` 内の `SEARCH_CLAUSE` という文字列を空文字に置換して消している。変数`searchQuery`は少し前の[1160行目](https://github.com/mattermost/mattermost-server/blob/aad76a13e8381fc2a7c79d37d842d530c450a728/store/sqlstore/post_store.go#L1160)で宣言されているSQL文。SQLテンプレートの内部を置換しながらSQLを組み立てて行っている模様。検索ワードの有無は`SEARCH_CLAUSE`というパートで扱われていることがわかる。

検索ワードが指定されている場合は使用しているDBごとにクエリの作り方を分岐させている。

### 検索ワード有り かつ PostgreSQLの場合

1. `terms`, `excludedTerms`のワイルドカードを `*` -> `:*` に置換
2. 除外句(`excludedClause`)の構築
  * `& !(exTerm1|exTerm2|exTerm3))`
3. 検索句全体(`:Terms`)の構築
  * OrTerms?
    * YES -> `(Term1|Term2|Term3) & !(exTerm1|exTerm2|exTerm3))`
    * No  -> `(Term1&Term2&Term3) & !(exTerm1|exTerm2|exTerm3))`

途中、SearchParams内のbool型変数`OrTerms`で分岐しているが、APIドキュメントによると`OrTerms`というのはOR検索かAND検索かを判別フラグっぽい。

https://api.mattermost.com/#tag/posts/paths/~1teams~1{team_id}~1posts~1search/post
> is_or_search
> Set to true if an Or search should be performed vs an And search.

以上からPostgreSQLの場合、検索除外ワードのみが設定されていた場合、検索句全体(`:Terms`)は `& !(exTerm1|exTerm2|exTerm3))` となるはずです。検索句が `&` で始まっているのでクエリ的に怪しい気が(SQLについては詳しくないので予想ですが)。
ExcludedTermsしか設定されていない場合は、検索対象の語句をワイルドカードで与えてあげる (`:* & !(exTerm1|exTerm2|exTerm3))`) と良いような気がします。(PostgreSQLは使ってないので未検証)

### 検索ワード有り かつ MySQLの場合

今回知りたかった部分の処理。

1. 除外句(`excludedClause`)構築 
  * `-($excludedTerms)`
2. 検索句全体 (`:Terms`) の構築
  * OrTerms?
    * YES -> `$terms -($excludedTerms)`
    * NO  -> `+term1 +term2 +term3 -($excludedTerms)`

MySQLの場合、`OrTerms`が設定されていると与えられた（除外）検索句をそのまま利用するようです。
スペース区切りで与えられるからそのままOr検索になるのかな？ ただそうだとすると、除外句の部分 (`-($excludedTerms)`) を見ると、除外句の方が `or` にならなそうな気がしますが...。

以上から、MySQLでは、検索除外話ーどのみが設定されていた場合、検索句全体(`:Terms`)は `-($excludedTerms)` となるはずです。

`-` で開始する検索句が有効なのか分からないのでMySQLのドキュメントを見てみる。

するど、どうやら検索除外ワードのみのクエリが与えられた場合、MySQLは空の結果を返すようです。

https://dev.mysql.com/doc/refman/5.7/en/fulltext-boolean.html

```
* -
 A leading or trailing minus sign indicates that this word must not be present in any of the rows that are returned. InnoDB only supports leading minus signs.
 
 Note: The - operator acts only to exclude rows that are otherwise matched by other search terms. Thus, a boolean-mode search that contains only terms preceded by - returns an empty result. It does not return “all rows except those containing any of the excluded terms.”
```

Mattermostは`boolean-mode`で検索を行っているため、検索除外ワードのみで検索ワードが与えられていないと検索結果が空となってしまい、空の検索結果から何を除外しても空ということのようです。
検索ワード、検索除外ワードともに空の場合は、PostgreSQL/MySQLごとに処理を分岐する以前に検索クエリ部分をSQL文から削除しているため、全ての投稿が検索結果として返ってきます。

つまり、MySQLの場合も検索除外ワードのみが設定されているような場合は、検索対象の語句をワイルドカードで与えてあげる (`* -($excludedTerms)`) と良い気がします。
が、ワイルドカード(`*`) を使った検索は `innodb_ft_min_token_size` の影響を受けてしまい、1文字のワイルドカード(`*`)はクエリとして無効となってしまうらしい...。(`AB*` のように2文字 + ワイルドカード(`*`) ならOK）。
Mattermostはこのサイズが最低でも`2`にしかできなかったため、単純にワイルドカードを付与する方法ではダメそうです...。

## おわりに

どうしよ。とりあえずIssueを立てた。
https://github.com/mattermost/mattermost-server/issues/14641


Issue立ててる間に調べててわかったけど、[Mattermostのコミュニティサーバー](https://community.mattermost.com/core)だと検索できるな。Elasticsearch使ってるからかな。お金を払ってちゃんと使おう。
