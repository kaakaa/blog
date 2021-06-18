---
title: "ConcurrentとParallel"
date: 2019-07-06T22:36:52+09:00
draft: false
toc: true
tags: ["programming","go","misc"]
---

# はじめに
[O'Reilly Japan \- Go言語による並行処理](https://www.oreilly.co.jp/books/9784873118468/)

を読んでいると、気になるフレーズが出てきました。

> 「並行性はコードの性質を指し、並列性は動作しているプログラムの性質を指します。」

「そうなの？」と思い某所で呟いてみると、色々情報が集まったのでまとめます。

## Go言語におけるConcurrent/Parallelの解釈

最初に、`「並行性はコードの性質を指し、並列性は動作しているプログラムの性質を指します。」`のについて。

原文では

> Concurrency is a property of the code; parallelism is a property of the running program.

となっているようです。(参照: [O'Reilly Concurrency in Goの読書メモ \- Qiita](https://qiita.com/ymgyt/items/420eaf2bcf7bee4ae152#difference-between-concurrency-and-parallelism))

また、Go言語に関する書籍である背景を鑑みると、Rob Pike先生の下記の言葉に基づいた解釈のような気がします。

### [Concurrency is not parallelism \- The Go Blog](https://blog.golang.org/concurrency-is-not-parallelism)

> In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations.

ここではハッキリと**Concurrency is the composition**、**Parallelism is the simultaneous execution**と言っています。それぞれの言葉を具象化すると**the composition ⊃ code**、**the simultaneous execution ⊃ running program**と言い換えられる気がするので、書籍の言葉はここからきているのでは無いかと思われます。

## その他のConcurrent/Parallelに関する話

ここからは余談になりますが、Go言語に限らずConcurrent/Parallelについて語られている情報について触れていきます。

### [並列\(Parallel\)と並行\(Concurrent\)の違いについて \- Togetter](https://togetter.com/li/159203)

このまとめ自体は2011年のものなので古い（とはいえ並行/並列という概念についていえば古いということは無い）ですが、ここから参照されている下記の記事が個人的にはわかりやすかったです。

### [parallel と concurrent、並列と並行の違い \- 本当は怖い情報科学](http://freak-da.hatenablog.com/entry/20100915/p1)

このエントリではこのように説明されています。

> * Concurrent（並行）は「複数の動作が、論理的に、順不同もしくは同時に起こりうる」こと
> * Parallel（並列）は、「複数の動作が、物理的に、同時に起こること」

元の`「並行性はコードの性質を指し、並列性は動作しているプログラムの性質を指します。」`というフレーズに照らして考えてみると、

* Concurrent(並行)は「複数の動作が、**論理的**に、順不同もしくは同時に起こりうる」ということですので、同時実行されるという事実を記述したコードがConcurrentであるという説明は合っている気がします
* また、Parallelも同様に「複数の動作が、**物理的**に、同時に起こること」ということなので、「（複数のコアを使うなどして）**動作している**プログラムの性質を刺す」という説明と合うと思います

`Concurrent = 論理的 = code`、`Parallel = 物理的 = 動作するプログラム`と考えると、しっくりくるような気がします。
(ただし、後で説明しますがConcurrent/Parallelは対義語などではなく、また並べて比較されるような言葉では無いため、**Concurrent/Parallel**を**論理的/物理的**と単純に解釈するのは危険です)

### [1\. Introduction \- Parallel and Concurrent Programming in Haskell \[Book\]](https://www.oreilly.com/library/view/parallel-and-concurrent/9781449335939/ch01.html)

> A parallel program is one that uses a multiplicity of computational hardware (e.g., several processor cores) to perform a computation more quickly. The aim is to arrive at the answer earlier, by delegating different parts of the computation to different processors that execute at the same time.
> 
> By contrast, concurrency is a program-structuring technique in which there are multiple threads of control. Conceptually, the threads of control execute “at the same time”; that is, the user sees their effects interleaved. Whether they actually execute at the same time or not is an implementation detail; a concurrent program can execute on a single processor through interleaved execution or on multiple physical processors.

次に、Haskellに関する書籍でのConcurrent/Parallelの説明。

この説明でもConcurrencyはプログラムの構造化技法、Parallel Programは多数のハードウェアを同時に使うものといっており、今までの理解と相違無いように読めます。

また、Concurrent(並行)については

> Whether they actually execute at the same time or not is an implementation detail

と、「実際に同時に実行されるかどうかは実施の詳細による」と、実際の動作についてはConcurrentの範疇ではないと言っています（ここでimplementation=`実装`と訳すと混乱するため`実施`と訳しています）。

また、Parallel(並列)については、

> The aim is to arrive at the answer earlier, by delegating different parts of the computation to different processors that execute at the same time.

「処理高速化を目的に複数の異なるプロセッサで同時に処理を実行すること」と、実施形態について意味していることが分かります。

ここまでで、**Concurrent**はプログラム技法やコードなど**実施形態の構成**を表すために使われることが多く、**Parallel**は実際に複数のプロセッサを使って処理を行うなど**実施形態**について意味する言葉だと理解することができ、そもそも並べて比べるような言葉でないことが分かってきます。

### [<=: Serial vs Parallel, Sequential vs Concurrent](http://s1l3n0.blogspot.com/2013/04/serial-vs-parallel-sequential-vs.html)

最後にこちらの記事。

タイトルからConcurrent/Parallelの対義語はそれぞれSequential/Serialだと言っています。
このエントリではConcurrencyをSequentialの対義語、つまり**同時発生**のような意味合いで使っているため**Concurrent is the compositon**という捉え方で読むと少し違和感を覚えますが、どのように同時発生的な事象を扱うかという方法論だと捉えると、今までのエントリで言っている事と変わらない気がします。

そして何より、Concurrent/Parallelそれぞれの対義語を意識することでその違いがより明確になります。

また、このエントリの最後の方の例で面白い例があり、並列な実行基盤の上で動作しているSequentianなプロセスは**deterministic**であり、並列な実行基盤の上で動作しているConcurrentなプロセスは**non-deterministic**であるということです。この概念についてはまだ消化しきれていませんが、Concurrentが同時発生的な概念であると考えると正しい気がします。

# おわりに
本エントリでは、Concurrent/Parallelについて調べました。
ここで挙げた情報源は公式な定義というわけでは無い（あったら教えてください）ため、どこかの解釈に誤りがある可能性があります。ただ、各所で語られている論理をまとめたことで、個人的にはそれぞれの言葉の意味するところが明確になったような気がします。
