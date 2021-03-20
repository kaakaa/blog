
---
title: "MattermostのLocalizationフロー"
date: 2016-08-29T18:59:39+09:00
draft: false
toc: true
tags: ["Go","golang","localization","Mattermost"]
---

社内勉強会用に作成した資料だったものに手を加えて公開。

# はじめに
Mattermostのローカライズに協力し始めたので、そのフローの記録(2016/12月現在)

# Mattermost開発参加の背景

## チャットツールとの出会い

外部の勉強会にはよく参加しており、DevOps/ChatOpsなんかの言葉も認識していました。
少し前のHubotによるChatOpsブームの際に、お試しとして社内のLinuxサーバにKandanを入れたのがChatツールとの出会いです。

その後、一緒にチャットを使ってくれる方々を巻き込みつつ、Kandanは機能が足りないのでLet's chatへ、Let's Chatの開発滞ってきたところに同期がDockerでMattermost立ててくれたので移行したり。
そんな感じでMattermostを使い始めてからは、機能的に不満もないためバージョンアップを繰り返しながら、そろそろ１年使い続けています。

最近ようやくSlackを使い始めたのですが、Mattermostに体が慣れすぎて違和感を感じ、プロキシという壁に阻まれた生活を続けたことで身体がガラパゴス化しているのを痛感しています。


## ローカライズ参加までの流れ

元々、リモートワーク的な開発を経験するために、そこそこ大きなOSSにコントリビュートしてみたいと思っていました。参加するなら、やはり日常的に使うツールの方が良いと思っており、また、MattermostはサーバサイドGo言語・フロントエンドReact.jsと、まさに興味ピンポイントな技術を採用していたため非常に興味を持っていました。

そんな中、普通にMattermostを使う中で誤翻訳を一つ発見。

良い機会なので、公式サイトの[Localizationガイド](https://docs.mattermost.com/developer/localization-process.html)を見ながら開発チームのチャット（Mattermost）に参加し、翻訳のサジェスションを行ってみました。（翻訳を確定するには権限が必要。権限のないユーザはサジェスションまでしか行えない）
その時、未翻訳だと思われるメッセージが残っていたため、併せて翻訳作業を続けていました。

そうこうしている内に、新しいバージョンのリリースが近づいたようで翻訳サーバーに翻訳対象メッセージ大量投入されました。が、その当時の公式レビュワーの方が多忙で連絡つかないようで、開発チームのチャットの方でコアメンバーから個別に日本語公式レビュワーにならないかと打診され、本格的に日本語への翻訳を開始することになりました。

* * *

![matran8.png](https://qiita-image-store.s3.amazonaws.com/0/9891/4808ed14-a47a-8694-97e5-c68d3bfaac12.png)
* jason 「他に日本語翻訳してるひといないから公式レビュワーならない？」
* kaakaa 「( ﾟдﾟ )ﾎﾟｶｰﾝ..Okay!!」


#  Mattermostローカライズの進め方

ここからは、実際の翻訳作業について記述します。

## ローカライズ対象

Mattermostのローカライズ対象メッセージはサーバー/クライアントそれぞれ別ファイルで管理されています。

サーバーサイドはGo言語で書かれており、[`utils.T`](https://github.com/mattermost/platform/blob/master/utils/i18n.go#L14)もしくは[`Context.T`](https://github.com/mattermost/platform/blob/master/api/context.go#L44)を呼び出している部分で言語ごとのメッセージに変換しています。

```go
func NewServer() {
	l4g.Info(utils.T("api.server.new_server.init.info"))

	Srv = &Server{}
}
```
[platform/server.go at master · mattermost/platform](https://github.com/mattermost/platform/blob/master/api/server.go#L40)

引数のメッセージIDに対応するメッセージはJSON形式で[`i18n/{lang}.json`](https://github.com/mattermost/platform/tree/master/i18n)に存在します。

![localize_server.jpg](https://qiita-image-store.s3.amazonaws.com/0/9891/c5c5fc24-6a16-4aeb-b0f7-b306de5ab8d8.jpeg)
[Mattermost Developer Brown Bag - Localization - YouTube](https://www.youtube.com/watch?v=cVRmzXjpp7Y)

* * *


クライアントサイドはReact.jsで書かれており、下記関数を使って言語の変換を行っています。

* react-intoの`FormattedMessage`, `FormattedTime`, `FormattedDate`
* [loclizeMessage関数](https://github.com/mattermost/platform/blob/master/webapp/utils/utils.jsx#L1166)
 

```xml
<FormattedMessage
  id='help.learnMore'
  defaultMessage='Learn more about:'
/>
```
[platform/messaging.jsx・mattermost/platform](https://github.com/mattermost/platform/blob/bb69e98631b2541954b6ae465b8ba5f788b9dc49/webapp/components/help/components/messaging.jsx#L30)


引数のメッセージIDに対応するメッセージは[`webapp/i18n/{lang}.json`](https://github.com/mattermost/platform/tree/master/webapp/i18n)に存在します。

![localize_webapp.jpg](https://qiita-image-store.s3.amazonaws.com/0/9891/0bcb6e46-e8c8-0532-1eb0-bcb8ba22d34a.jpeg)
[Mattermost Developer Brown Bag - Localization - YouTube](https://www.youtube.com/watch?v=cVRmzXjpp7Y)

* * *

各言語メッセージがかかれている`{lang}.json`ファイルのうち、実際にPullRequestなどで編集されるのは`en.json`だけで、その他の言語ファイルについては後述の翻訳サーバーにより生成されます。

## Pootle

ローカライズ作業は[Pootle](http://pootle.translatehouse.org/)を利用して行っています。

  * [Welcome \| Mattermost Translation Server](http://translate.mattermost.com/)

日本語翻訳についてのコミュニケーションは[開発者チャット](https://pre-release.mattermost.com/login "Mattermost")上の`i18n - Japanese`や`i18n - Localization`あたりで行われています。
開発者チャットでは`Release Discussion`や`Developers`などの開発上のコミュニケーションなども行われているため、開発に参加する場合はアカウントを作成しておいたほうが良いです。

### Pootleとは?

* [Pootle \- Wikipedia](https://ja.wikipedia.org/wiki/Pootle)
  * `Pootle は翻訳画面を備えたオンラインの翻訳支援ツールである。ツールは Python で書かれたフリーソフトウェアであり、2004年から Translate.org.za において開発、リリースされている。`
  * `Pootle はフリーの翻訳支援ソフトウェアからの利用を想定しているが、それに限定されてはいない。また、文書の翻訳よりもソフトウェア・アプリケーションの GUI の国際化を行うことを目的として開発されている。`
  * `Pootle は、翻訳作業の様々な局面で利用することができる。もっとも単純には、サーバーにある翻訳文の統計データを表示することができる。これにより他のユーザーに修正案を示したり、翻訳を修正してレビューに回したりでき、一種のバグ・レポート管理システムとして使うことができる。また多数の翻訳者に担当を割り当てることができ、さらにオフラインで行われた翻訳をまとめ、翻訳作業全体のワークフローを管理することができる。`
  * `Pootle は OpenOffice.org[1]、OLPCプロジェクト[2]、その他[3]で利用されている。また Mozilla プロジェクトの翻訳インフラを構築しているVerbatim project でも基盤として採用されている。`

## ローカライズフロー

実際の開発で`en.json`内のメッセージが追加・変更されると、毎日0:00(GMT)に自動でPootleサーバーへ翻訳対象のメッセージの追加が行われます。
また、Pootleサーバー上で翻訳が完了したメッセージは、毎週月曜の22:00(GMT)に本体へPull Requestが作成され、それがマージされることにより翻訳作業が完了します。
[translations PR 20161208 by enahum · Pull Request \#4740 · mattermost/platform](https://github.com/mattermost/platform/pull/4740)

![mattermost_translatin.png](https://qiita-image-store.s3.amazonaws.com/0/9891/1d31a769-2eb7-8acc-3718-edb3ea8e0da6.png)
[Mattermost Developer Brown Bag \- Localization \- YouTube](https://www.youtube.com/watch?v=cVRmzXjpp7Y)


### Pootleサーバーでの実際の作業

全ての言語のローカライズがここで行われている
![matran1.png](https://qiita-image-store.s3.amazonaws.com/0/9891/07a4bdf9-48d4-b664-a751-95b4bacade98.png)

* * *

メッセージはサーバー用とフロントエンド用に分かれている
![matran2.png](https://qiita-image-store.s3.amazonaws.com/0/9891/0b2664a4-7aec-850c-cd73-f9dc24ed3ca1.png)

* * *

翻訳対象メッセージのステータスは`Translated(翻訳済み)`、`Fuzzy(機械翻訳)`、`Untranslated(未翻訳)`に分類される。
![matran3.png](https://qiita-image-store.s3.amazonaws.com/0/9891/793d6697-768e-9023-7eb4-01a6362439b6.png)

* * *
 
`Fuzzy(機械翻訳)`は、翻訳大将メッセージがサーバに投入されると同時にPootleシステムが自動で翻訳を行ってくれたメッセージ。
レビュワーによる承認作業が必要。
![matran4.png](https://qiita-image-store.s3.amazonaws.com/0/9891/f157c264-af82-5266-73d0-56a0e0a56e21.png)

* * *

翻訳作業は１メッセージごとに行う
![matran5.png](https://qiita-image-store.s3.amazonaws.com/0/9891/1d34bd10-65c2-edc8-d770-fdd58138ce7b.png)

* * *

翻訳を行うとこのように見える。レビュワーがReject(拒否)/Submit(承認)を判断する。
![matran6.png](https://qiita-image-store.s3.amazonaws.com/0/9891/2ed9ee6f-dea8-b914-15cb-dc99ac5f2ccc.png)

### 自分で翻訳ファイルを適用する場合

翻訳ファイルは毎週masterにマージされ、最新版が含まれた状態でリリースされるため、基本的に自分で動作しているMattermostに最新の翻訳を適用することは無いと思いますが、緊急で特定のメッセージを翻訳しなきゃいけない場合のために。

Pootleサーバーにて翻訳済みのメッセージはPootle画面からダウンロードできます。
![matran7.png](https://qiita-image-store.s3.amazonaws.com/0/9891/80e311a1-ca41-f732-0a83-3a2e5035ee9d.png)

Pootleサーバからダウンロードできるファイルは`.po`というファイルなのですが、実際のシステムで使用するためには`.po`ファイルを`.json`形式に変換する必要があります。
変換には[rodrigocorsi2/mattermosti18n](https://github.com/rodrigocorsi2/mattermosti18n)を使用するようです。

なんとなく、TravisCIで`.json`ファイルまで生成できるリポジトリ[kaakaa/mattermosti18n](https://github.com/kaakaa/mattermosti18n/tree/release)を作ってみましたが、最新の翻訳データを使いたいと思ったことが無いので使っていません。
（英語なら英語のままで良いやと思ってしまう）


### 翻訳する時に気をつけていること

* 可能な限り既存の訳し方を踏襲する
  * 自分の中で言葉が見つかっても、まずは同じ英単語で訳を検索して同じかを確認
  * なるべく同じ概念を違う言葉で表現することが無いように（難しい。。。）
* ソースコードを検索して、そのメッセージがどこで利用されているかを確認する
  * パッと見た時に訳し方を悩む場合はソースコードを確認して、近くで出現するメッセージを確認する
  * 設定画面で使われるのか・注意書きとして使われるか...など

気をつけてはいますが、なかなか統一できないことも多々。。

# ローカライズに参加してみて

## 成果

* [MattermostのTシャツ](http://www.zazzle.com/knights_of_mattermost_shirt_black-235102741116730031)もらえた
  * ポロシャツとかパーカーとかいっぱいある（デザイン同じ）[Knights of Mattermost Shirt \(Black\)](http://www.zazzle.com/pd/styles/pd-235102741116730031?dp=&qs=&tl=)
* 公式サイトに名前が載った
  * [Localization — Mattermost documentation](https://docs.mattermost.com/developer/localization-process.html)
* Pootle使った翻訳手法を知れた
* チャット上での英語でのコミュニケーションもなんとかなることが分かった（返信に時間かかるけど）

## 感想

* 開発プラットフォームが分散してる(Github/JIRA/チャットなど)ので慣れるまでは大変そう
  * コアメンバによるバグfixなどは結構スピーディーなので、コード修正とかに参加するのはちょっと壁がある感
* コアメンバはチャットへの返信も早くて親切
  * どんな質問にも答えてくれるけど、英語はやっぱり必要
* Mattermostチームは楽しそうに開発してるなーという印象を受けました

