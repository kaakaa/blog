---
title: "Mattermostの日本語検索を改善するためにBleveSearchを少し調査した"
emoji: "🎉"
published: true
date: 2022-08-16T17:00:00+09:00
toc: true
tags: ["mattermost", "bleve", "japanese", "search"]
---

## 背景: Mattermost Bleve検索

2021/06にリリースされたMattermost v5.24から、[Bleveによるメッセージの日本語検索が利用できるようになった](https://zenn.dev/kaakaa/articles/qiita-20200620-574972591f6b0b0f642f#bleve-%E3%82%92%E5%88%A9%E7%94%A8%E3%81%97%E3%81%9F%E5%85%A8%E6%96%87%E6%A4%9C%E7%B4%A2%E3%81%A7%E6%97%A5%E6%9C%AC%E8%AA%9E%E6%A4%9C%E7%B4%A2%E3%82%82%E5%8F%AF%E8%83%BD%E3%81%AB%EF%BC%81%EF%BC%88%E5%AE%9F%E9%A8%93%E7%9A%84%E6%A9%9F%E8%83%BD%EF%BC%89)。  

[Bleve search \(experimental\)](https://docs.mattermost.com/deploy/bleve-search.html)

それ以前のバージョンでは、[Mattermostセットアップ時にDB(MySQL, PostgreSQL)の設定に手を加える](https://docs.mattermost.com/configure/enabling-chinese-japanese-korean-search.html#japanese)必要があったが、Bleveによる検索機能が公式にサポートされたことで、システムコンソール画面からいくつか設定を変更するだけで日本語検索が実現できるようになった。嬉しい。

## 問題: 意図せぬ投稿が検索される

Bleve検索自体はとても便利な機能だが、その検索結果に少し問題がある。

2文字以上の単語で構成される日本語クエリによるBleve検索をおこなった際、本来ならばその単語が出現する投稿のみが検索結果として表示して欲しいところだが、実際の検索結果には、その単語内で使用されている文字が全て現れる投稿を検索してしまう。

例えば、「メロスは激怒した。必ず、かの邪智暴虐の王を除かなければならぬと決意した。」という投稿があったとする。

「激怒」「邪智暴虐」などの単語で検索を行うと、想定通りこの投稿を検索することができる。しかし、「激暴」という単語で検索を行った場合も、この単語自体は対象の投稿に存在しないにも関わらず、単語を構成する「激」「暴」の両方の文字が投稿に含まれているため、対象の投稿が検索結果として表示されてしまう。

![problem-bleve-search](https://blog.kaakaa.dev/images/posts/tech/bleve-search/problem-bleve-search.png)

この問題を改善するため、調査を行った。

## 調査1: Bleveのcjk言語向けAnalyzer

まず、Mattermost側でBleveをどのように使用しているかを調べる。

投稿内容のインデックスは、以下で定義されている。

```go
...
import (
    ...
	"github.com/blevesearch/bleve/v2/analysis/analyzer/standard"
    ...
)
...
func init() {
    ...
	standardMapping = bleve.NewTextFieldMapping()
	standardMapping.Analyzer = standard.Name
    ...
}
...
func getPostIndexMapping() *mapping.IndexMappingImpl {
	postMapping := bleve.NewDocumentMapping()
    ...
	postMapping.AddFieldMappingsAt("Message", standardMapping)
    ...
}
```
https://github.com/mattermost/mattermost-server/blob/5ea2ca8a3a25cb89b4cf52012ae3897d03e0a64f/services/searchengine/bleveengine/bleve.go#L79

Mattermostの投稿(`Post`)の内容(`Message`)は、Bleveの`standard.Name`で表されるAnalyzerによって解析されるようだ。Bleveについてあまり詳しくはないが、`standard.Name`で表されるAnalyzerの内容を見ると、`en.StopName`が使われており、英文を前提としたAnalyzerだと考えられるので、この部分が問題なのではないかと思う。  
https://github.com/blevesearch/bleve/blob/master/analysis/analyzer/standard/standard.go

BleveのAnalyzerについて調べていると、`cjk`向けのAnalyzerもあることが分かった。  
https://github.com/blevesearch/bleve/tree/master/analysis/lang/cjk

単純にMattermostの投稿内容のインデックスをBleveの`cjk` Analyzerを使って行うよう変更してみる。

```diff
@@ -14,7 +14,7 @@ import (
 
        "github.com/blevesearch/bleve/v2"
        "github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
-       "github.com/blevesearch/bleve/v2/analysis/analyzer/standard"
+       "github.com/blevesearch/bleve/v2/analysis/lang/cjk"
        "github.com/blevesearch/bleve/v2/mapping"
 
        "github.com/mattermost/mattermost-server/v6/model"
@@ -49,7 +49,7 @@ func init() {
        keywordMapping.Analyzer = keyword.Name
 
        standardMapping = bleve.NewTextFieldMapping()
-       standardMapping.Analyzer = standard.Name
+       standardMapping.Analyzer = cjk.AnalyzerName
 
```

この変更を有効にしてMattermostを再起動し、Bleveインデックスを破棄→再構築すると、「激暴」という単語で当該の投稿が検索されなくなった。

![bleve-search-improve1-1](https://blog.kaakaa.dev/images/posts/tech/bleve-search/bleve-search-improve1-1.png)

![bleve-search-improve1-2](https://blog.kaakaa.dev/images/posts/tech/bleve-search/bleve-search-improve1-2.png)

めでたしめでたし。

...とはならなかった。

上記の変更を加えたことで、今度は1語による検索ができなくなった。

![problem-bleve-search2](https://blog.kaakaa.dev/images/posts/tech/bleve-search/problem-bleve-search2.png)

「検索は2語以上で」という制限を設ければ問題は解決だが、この制限はあまり好ましくない。

## 調査2: Bleve `cjk` Analyzerを使用した際に1語検索ができるようにする

Bleveの`cjk` AnalyzerはTokenFilterとしてBigramを使用している。つまり、2語ずつインデックスを構築しているので、1語での検索ができないようだ。

```go
...
const AnalyzerName = "cjk"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
    ...
	bigramFilter, err := cache.TokenFilterNamed(BigramName)
	if err != nil {
		return nil, err
	}
	rv := analysis.Analyzer{
		Tokenizer: tokenizer,
		TokenFilters: []analysis.TokenFilter{
			widthFilter,
			toLowerFilter,
			bigramFilter,
		},
	}
    ...
}
...
```
https://github.com/blevesearch/bleve/blob/d89c6c0a6873fcca1673fb6e7e3128d39bc6494d/analysis/lang/cjk/analyzer_cjk.go#L40

もう一度、Bleveの`cjk`パッケージを見てみると`output_unigram`という設定があることに気づく。名前から、この設定を`On`にしてインデックスを構築することで、1語のインデックスも出力されるようになるのではないかと予想する。(それによるインデックスサイズの増加についてはここでは考慮しないことにする)

```go
...
func CJKBigramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	outputUnigram := false
	outVal, ok := config["output_unigram"].(bool)
	if ok {
		outputUnigram = outVal
	}
	return NewCJKBigramFilter(outputUnigram), nil
}
...
```
https://github.com/blevesearch/bleve/blob/d89c6c0a6873fcca1673fb6e7e3128d39bc6494d/analysis/lang/cjk/cjk_bigram.go#L185

`output_unigram`の設定をOnにする方法について色々調べてみたが、どうやってもOnにする方法が見つからない(これがこの記事を書くきっかけだったりする)。

以下の`ItemNamed`関数内で実行している`build(name, nil, cache)`の第二引数が前述の`CJKBigramFilterConstructor`関数の第一引数 (`config`)になり、ここに`output_unigram: true`が設定されていれば有効にできるらしいが、`nil`でハードコードされているのでどうしようもない。

```go
func (c *ConcurrentCache) ItemNamed(name string, cache *Cache, build CacheBuild) (interface{}, error) {
    ...
	// try to build it
	newItem, err := build(name, nil, cache)
    ...
}
```
https://github.com/blevesearch/bleve/blob/d89c6c0a6873fcca1673fb6e7e3128d39bc6494d/registry/cache.go#L47


自分が気づいていない方法があるのかも知れないが、調べていてもわからないので、`output_unigram`を`true`にしたAnalyzerを新たに登録し、それを使用してみる。

```diff
@@ -13,9 +13,13 @@ import (
        "time"
 
        "github.com/blevesearch/bleve/v2"
+       "github.com/blevesearch/bleve/v2/analysis"
        "github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
-       "github.com/blevesearch/bleve/v2/analysis/analyzer/standard"
+       "github.com/blevesearch/bleve/v2/analysis/lang/cjk"
+       "github.com/blevesearch/bleve/v2/analysis/token/lowercase"
+       "github.com/blevesearch/bleve/v2/analysis/tokenizer/unicode"
        "github.com/blevesearch/bleve/v2/mapping"
+       "github.com/blevesearch/bleve/v2/registry"
 
        "github.com/mattermost/mattermost-server/v6/model"
        "github.com/mattermost/mattermost-server/v6/shared/mlog"
@@ -44,12 +48,50 @@ var keywordMapping *mapping.FieldMapping
 var standardMapping *mapping.FieldMapping
 var dateMapping *mapping.FieldMapping
 
+const customCJKAnalyzerName = "custom_cjk_analyzer"
+const customCJKTokenFilterName = "custom_cjk_filter"
+
+func CustomCJKAnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
+       tokenizer, err := cache.TokenizerNamed(unicode.Name)
+       if err != nil {
+               return nil, err
+       }
+       widthFilter, err := cache.TokenFilterNamed(cjk.WidthName)
+       if err != nil {
+               return nil, err
+       }
+       toLowerFilter, err := cache.TokenFilterNamed(lowercase.Name)
+       if err != nil {
+               return nil, err
+       }
+       bigramFilter, err := cache.TokenFilterNamed(customCJKTokenFilterName)
+       if err != nil {
+               return nil, err
+       }
+       rv := analysis.Analyzer{
+               Tokenizer: tokenizer,
+               TokenFilters: []analysis.TokenFilter{
+                       widthFilter,
+                       toLowerFilter,
+                       bigramFilter,
+               },
+       }
+       return &rv, nil
+}
+
+func CustomCJKBigramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
+       return cjk.NewCJKBigramFilter(true), nil
+}
+
 func init() {
+       registry.RegisterTokenFilter(customCJKTokenFilterName, CustomCJKBigramFilterConstructor)
+       registry.RegisterAnalyzer(customCJKAnalyzerName, CustomCJKAnalyzerConstructor)
+
        keywordMapping = bleve.NewTextFieldMapping()
        keywordMapping.Analyzer = keyword.Name
 
        standardMapping = bleve.NewTextFieldMapping()
-       standardMapping.Analyzer = standard.Name
+       standardMapping.Analyzer = customCJKAnalyzerName
 
        dateMapping = bleve.NewNumericFieldMapping()
 }
```

冗長だ。

再びMattermostを再起動、Bleveインデックスの破棄→再構築を行い検索を実行すると、1語でも検索ができるようになった。

![bleve-search-improve2-1](https://blog.kaakaa.dev/images/posts/tech/bleve-search/bleve-search-improve2-1.png)

![bleve-search-improve2-2](https://blog.kaakaa.dev/images/posts/tech/bleve-search/bleve-search-improve2-2.png)

![bleve-search-improve2-3](https://blog.kaakaa.dev/images/posts/tech/bleve-search/bleve-search-improve2-3.png)


## おわりに

というわけで、MattermostのBleveによる日本語検索の改善について調べてみた。  
とりあえず小さなサンプルでは実現可能なことが分かったが、`output_unigram`の設定のせいで修正量が多くなりそうなので、もう少し調査してからIssueで報告しようと思う。

調査内容としては以下のあたりかな。

* 検索内容の正当性
  * 日本語検索のテストセットが欲しいな...
* 検索インデックスの容量増加
  * 日本語の投稿がたくさんあるMattermostインスタンスのデータが欲しいな...
* `blevex`のkagomeを使用した検索
  * `blevex`のREADMEを見ると、Mattermost本体に組み込んでメンテし続けるのはちょっと怖い気がするな...

Mattermostで実装するときは、Bleveの設定画面でAnalyzer選ぶような感じになるのかな。CJK以外でAnalyzer変更が必要な言語が無かったりすると、ちょっと独自改造っぽい感じになるので受け入れられにくそうな感じもするな。
