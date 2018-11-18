
---
title: "Mattermostの開発に参加してみた記録"
date: 2016-12-16T00:03:36+09:00
draft: false
toc: true
tags: ["OSS","Mattermost"]
---

この記事は[Fujitsu extended Advent Calendar 2016](http://qiita.com/advent-calendar/2016/fujitsu_ext) 16日目の記事です。

※この記事に描かれていることは個人の見解であり、所属する組織の公式見解ではありません。`
個人的な活動の経験から得た知見の共有を目的としています。

# はじめに

SI企業で働く傍ら社外のプログラミング関係の勉強会にも参加しており、その中で語られるOSS的な開発プロセスがどんなものであるか、実際に参加して体験してみたいと常々思っていました。
今年の夏ごろから[Mattermost](https://about.mattermost.com)というSlackクローンのOSSチャットツールの開発に参加し始めているので、Mattermost開発の進め方などを紹介したいと思います。
他の組織の開発を体験することで、日頃の開発の改善点が見つかるのでは無いかというのが参加のモチベーションです。

Mattermostを選んだ理由としては、今年の初め頃から日常的にMattermostを利用しており、Eat Your Own Dogfoodの精神で「より良くしたい」という意識を持ちやすいことです。
また、MattermostはサーバーサイドのGo言語、フロントエンドのReact.jsを始めとし、DockerfileやiOS,Androidアプリ(最近react-nativeで書き直し中)、[Electron製のデスクトップアプリ](https://github.com/mattermost/desktop)など、新し目の技術を取り入れているため、普段は使わない技術に触れられるところが魅力でした。

開発に参加してると言っても、枝葉のIssue対応ぐらいしかしていないので、もちろんMattermostチームを代表しての記事ではありません。
誤っている部分がありましたら、報告いただけると幸いです。

# チャットツール

## Mattermost

MatermostはPalo Altoにある[Mattermost, Inc.](https://about.mattermost.com/company/)が開発しているオープンソースのチャットツールです。
[公式サイト](https://about.mattermost.com/)にもあるように`Open source, self-hosted Slack-alternative`を謳っています。。
名前の由来は`achieve what matters most(最も重要なことを達成する)`だそうです。

機能的にはSlackに追従することを目的としつつ、オンプレミスで動作可能であることやオープンな場で開発されていることを強みとしています。
また、Slackにはない多言語対応や、Enterprise版ではAD/LDAP認証を利用できるなど独自の機能も追加しています。
[Pricing – Mattermost](https://about.mattermost.com/pricing/)

隔月の定期リリースを行っており、次回のv3.6のリリースは2017/1/16に予定されています。

## その他のチャットツール

Mattermost以外にもChatツールはプロプライエタリ・OSS含め様々あります。

|名前|公式サイト|Github|スター数|言語|
|:----:|:--------:|:--------:|:------------:|:--------:|
|Slack|[Slack: Be less busy](https://slack.com/)|SaaSのみ|-|PHP|
|HipChat|[HipChat - Atlassian](https://ja.atlassian.com/software/hipchat)| SaaSのみ|-|?|
|Mattermost|[Mattermost](https://www.mattermost.org/)|[github](https://github.com/mattermost/platform)|[![GitHub stars](https://img.shields.io/github/stars/mattermost/platform.svg)]()|go,js|
|RocketChat|[Rocket.Chat](https://rocket.chat/)|[github](https://github.com/RocketChat/Rocket.Chat)|[![GitHub stars](https://img.shields.io/github/stars/RocketChat/Rocket.Chat.svg)]()|js,coffee|
|Zulip|[Zulip](https://zulip.org/)|[github](https://github.com/zulip/zulip)|[![GitHub stars](https://img.shields.io/github/stars/zulip/zulip.svg)]()|python,js|
|Actor Messanger|[Actor Messenger](https://actor.im/)|[github](https://github.com/actorapp/actor-platform)|[![GitHub stars](https://img.shields.io/github/stars/actorapp/actor-platform.svg)]()|java,scala,swift|
|Let's Chat|[Let's Chat](http://sdelements.github.io/lets-chat/)|[github](https://github.com/sdelements/lets-chat)|[![GitHub stars](https://img.shields.io/github/stars/sdelements/lets-chat.svg)]()|js|
|Kandan|[KandanApp](http://getkandan.com/)|[github](https://github.com/kandanapp/kandan)|[![GitHub stars](https://img.shields.io/github/stars/kandanapp/kandan.svg)]()|js,coffee,ruby|

参考：[Kickball/awesome\-selfhosted](https://github.com/Kickball/awesome-selfhosted#custom)

基本無料かつサーバーのお守りの必要のないSaaSであるSlackがチャットツールの定番かと思いますが、やはり組織内で利用するには外部サービスということで利用に慎重になってしまう部分があったため、何も心配せず使えるOSSを選択していました。

私はOSSとしてはKandan, Let's Chat, Mattermostしか使ったことがないですが、Rocket.chatとZulipも開発が活発なため、OSSのチャットツールを検討する場合は候補となるかと思います。
Kandanは機能的に物足りなく、Let's Chatは快適に利用できてはいましたが最近ほとんどメンテナンスされていないため、Mattermostへ乗り換え、今でもコミュニケーションツールとして利用しています。

# Mattermost開発ツール

Mattermostチームは多くのサービスを使って開発されており、そのほとんどが誰でも閲覧できるところに公開されています。（各サービスのアカウントを作成する必要はありますが）

## [Mattermost](https://pre-release.mattermost.com/)

Mattermost開発におけるコミュニケーションの多くはMattermost上で行われています。
バグのトリアージからPull RequestやGithu Issueへの更新の通知、ユーザーフォーラムの質問への回答依頼など、Mattermost本体に関わることから、ドキュメンテーション、ローカライズなどなど開発に関することが日頃から議論されています。
ほぼすべて英語です。

「どこどこのチャンネルに属さなくてはいけない」という指定は無いと思いますが、開発に参加する場合は`Contributors`と、開発対象のチャンネル(`Documentation`、`i18n - localization`など)の動向はチェックしておいても良いかと思います。

開発チャットとして使われているMattermostは最新のコードから作成された環境のため、新機能をいち早く試すことができます。
最近ではSlackにあるような、投稿に対するReactionの機能などが追加されています。

最初にログインした時にデフォルトで参加することになるチャンネルである`Reception`の参加者数は1184名(2016/12/12時点)となっています。
![__Reception_-_Contributors_Mattermost.png](https://qiita-image-store.s3.amazonaws.com/0/9891/4243baac-8c38-fe80-96d3-2e35d276a7a0.png)

## [Github](https://github.com/mattermost/platform)

Githubではソースコード、Issue、Pull Requestなどが管理されています。
最近追加されたGithubの新機能である[Projects](https://github.com/mattermost/platform/projects)やReviewers指定なども積極的に採用しています。

また、Mattermost本体のソースコードだけでなく、Mattermostのorganizationでは[ドキュメント(Sphinx製)](https://github.com/mattermost/docs)や[Dockerfile](https://github.com/mattermost/mattermost-docker)、[デスクトップ](https://github.com/mattermost/desktop)・[Android](https://github.com/mattermost/android)・[iOS](https://github.com/mattermost/ios)アプリなどのソースコードも公開されています。

最近では[React Nativeによるスマホアプリ](https://github.com/mattermost/mattermost-mobile)の開発も始まっているようです。


## [JIRA Cloud](https://mattermost.atlassian.net/secure/Dashboard.jspa)

JIRAでもIssueの管理が行われています。
GithubのIssue管理と被る部分もありそうですが、おそらくJIRAがメインのIssue Trackerで使用されており、GithubはGithubに慣れている人へのIssue報告の場やHelp Wanted Issue置き場として利用しているように思います。

## [Jenkins](https://build.mattermost.com)

CIサーバーとしてJenkinsを利用しています。
PullRequestが作成されると、テストとチェックスタイルが実行され、CIの結果はPullRequestへ通知されます。

![PLT-1555_Message_Editing_and_Deleting_permissions_by_ayadav_·_Pull_Request__4692_·_mattermost_platform.png](https://qiita-image-store.s3.amazonaws.com/0/9891/9e9f6309-b681-e391-0609-6c3958f6b2b0.png)

## [Mattermost Peer-to-Peer Forum](https://forum.mattermost.org)

ユーザーによるPeer-to-Peerのフォーラムも用意されています。
回答の必要があるトピックについては、定期的に開発者チャットの方でキャッチアップされているのため、回答がつかないということはあまり無いようです。

![Contributors_-_Contributors_Mattermost_と_マイファイル.png](https://qiita-image-store.s3.amazonaws.com/0/9891/448f8a43-0755-7167-1acf-cb85287ad9a7.png)


## [Mattermost Translation Server](http://translate.mattermost.com)

翻訳作業用のPootleサーバーです。
翻訳については、別エントリに詳しく書いています。
[MattermostのLocalizationフロー - Qiita](http://qiita.com/kaakaa_hoe/items/ac79289d65f7a3e3d9ea)

# Code Contribution

コードを送る際に必要なプロセスは、公式の[Development Process](https://docs.mattermost.com/guides/developer.html#development-process)にまとめてあるので、詳細についてはこちらを参照ください。
もちろん上記のドキュメントについてもオープンソースとして公開されており、PullRequestも受け付けています。https://github.com/mattermost/docs

ここでは、メインリポジトリである[mattermost/platform](https://github.com/mattermost/platform)へのコントリビューションのざっくりとした流れを紹介します。
（GithubでのPullRequestの作り方などの基本的なことには触れません）

### チケット選択

まず、対応するチケットを下記から選びます。

* [Help Wanted Issues - Github](https://github.com/mattermost/platform/issues?utf8=✓&q=is%3Aissue%20is%3Aopen%20%5BHelp%20Wanted%5D)
  * 2016/12/8〜2017/1/8まで、4つプルリクを送ると[Tシャツ(デザイン違うかも)](http://www.zazzle.com/knights_of_mattermost_shirt_black-235102741116730031)が貰える[Mattermost Holiday Hackfest](https://about.mattermost.com/mattermost-holiday-hackfest-2016/)が開催中
  * 2016年10月の[Hacktoberfest](https://hacktoberfest.digitalocean.com)以降、GithubにHelp Wanted Issueが作成されるようになった
* [Accepted Pull Request - JIRA](https://mattermost.atlassian.net/browse/PLT-4950?filter=10101)
  * [Good First Contribution - JIRA](https://mattermost.atlassian.net/browse/PLT-3485?filter=10206)というのもあります。が、最近あまり追加はされてないようです。

### CLA

対応するチケットが決まったら[Mattermost Contributor Agreement](https://www.mattermost.org/mattermost-contributor-agreement/)を結んでおきます。
これをやらずにPullRequestを送ると、Mattermostチームの方からPull Requestのコメントとして上記CLAを結ぶよう言われます。

### Developer Machine Setup

修正したソースコードの動作を確認するために、ローカルでMattermostを起動するための環境を作ります。
[Developer Machine Setup](https://docs.mattermost.com/developer/developer-setup.html#developer-machine-setup)を参考に。(DockerやGo環境が必要です)

### コーディング

MattermostはGithub Flowを採用していると思われます。
PullRequest用の作業をする際は、masterから開発用ブランチを作成し、開発が終わったらmasterに対してPullRequestを作成します。

ブランチ名はJIRAチケットなら`plt-394`のようにJIRAのチケット番号を使用します。
JIRAチケットがなく、Github Issueのみある場合は`gh-555`のようにIssue番号を使用しています。

---

実際のコードの修正部分ですが、リポジトリ内のファイル構造の概要は[Repository Structure](https://docs.mattermost.com/developer/developer-flow.html#repository-structure)に書かれています。
個人的な感想としては、サーバーサイドのGoの部分はディレクトリ構造から責務が理解しやすいですが、フロントエンドのReact Component入れ子を読み解くのが難儀でした。

---

動作確認は、先の[Developer Machine Setup](https://docs.mattermost.com/developer/developer-setup.html#developer-machine-setup)を済ませた環境であれば、ルートディレクトリで`make run`を実行することで、`http://localhost:8065`にMattermostが立ち上がるので、そちらで実機確認が出来ます。
また、Pull Requestを作成する前に`make test`、`make check-style`でエラーがないか確認しておきます。(私のMacBookAir 2011モデルでテスト完了まで10分程度かかるのが難点・・)

Pull Requestのテンプレートは下記のようになっているので指示通りに書き直します。

```
Please make sure you've read the [pull request](http://docs.mattermost.com/developer/contribution-guide.html#preparing-a-pull-request) section of our [code contribution guidelines](http://docs.mattermost.com/developer/contribution-guide.html).

When filling in a section please remove the help text and the above text.

#### Summary
[A brief description of what this pull request does.]

#### Ticket Link
[Please link the GitHub issue or Jira ticket this PR addresses.]

#### Checklist
[Place an '[x]' (no spaces) in all applicable fields. Please remove unrelated fields.]
- [ ] Added or updated unit tests (required for all new features)
- [ ] Added API documentation (required for all new APIs)
- [ ] All new/modified APIs include changes to the drivers
- [ ] Has enterprise changes (please link)
- [ ] Has UI changes
- [ ] Includes text changes and localization file ([.../i18n/en.json](https://github.com/mattermost/platform/blob/master/i18n/en.json) and [.../webapp/i18n/en.json](https://github.com/mattermost/platform/tree/master/webapp/i18n/en.json)) updates
- [ ] Touches critical sections of the codebase (auth, upgrade, etc.)
```

この辺りの詳細については[Workflow](https://docs.mattermost.com/developer/developer-flow.html#workflow)に記述があります。


### レビュー

Pull Requestを作成後、mergeされるためにはPMによるレビューと二人以上ののDeveloperによるレビューを通す必要があります。
PullRequestのレビューステータスはGithub上のラベル`1: PM Review`、`2: PM Review`などで表現されています。

また、`Setup Test Server`のラベルが貼られると、自動でPull Requestのコードを含んだMattermostを立ち上げ、その環境のURLをPull Requestにコメントしてくれる機能があるようです。
これらラベルの操作はコアチームしか操作できません。

レビュー完了までにConflictが発生するとrebaseを求められるので、対応します。
私は[こちらのエントリ](http://d.hatena.ne.jp/wyukawa/touch/20140603/1401797290)を参考にしながら対応しています。

---

上記がCode Contributionの流れです。

# 所感

今回、大きめと思われるオープンソースへのコントリビューションを体験してみましたが、基本的には一般的に良いと言われているプラクティスをしっかり遂行しているという感じで、大きな驚きはありませんでした。
ツールを選べば社内でも遂行できますし、実際に社内でもWebベースのレビューツールやチャット、Jenkinsを利用したCIなども行っていたこともあるため、キーワードだけ見れば同じ開発スタイルを遂行できると思います。

ただ、開発に参加するときは事務作業など考えず、コードだけをに集中できるため気が楽でした。社内との一番の違いはそこかと思います。

また、Mattermostチームはコントリビューターに対してとても親切だと感じました。PullRequestに対する対応も(リリース前後でなければ)基本的には早いですし、「Issue対応したいんだけど、どこから手をつけたらいい？」というような初歩的な質問にも丁寧に回答してくれます。
きちんと組織立ったOSSであるMattermostに参加できたことは、OSSに参加する上での不安や躊躇いを払拭する意味で有意義だったと感じます。

# おわりに

[二日目の高橋さんのエントリ](http://tnaoto.hatenablog.com/entry/2016/12/02/114611)にもあったように
> 自分たちの体験こそが知見ノウハウであり、お客様のビジネスをよりよくするための糧になると私は考えています。（個人の見解です）
ということだと思っています。

また、こちらの前半部分については個人的にとても同意できます。
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">「OSSはコストで採用するのではない。自立&gt;社員の成長&gt;事業の成長につながるものだ。」<br>わかる。「責任取りたくないから ORACLE」ってシステムが近所にいくつかあるけど、そういうベンダーが作ったシステムは DB 以外もメタクソ。エンドユーザが苦労している。 <a href="https://twitter.com/hashtag/elasticon?src=hash">#elasticon</a></p>&mdash; Toshihiko Nozaki (@bukaz54) <a href="https://twitter.com/bukaz54/status/809266226337746944">2016年12月15日</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

OSSを選択・OSSに参加するかどうかは個人個人考え方は違うでしょうが、他の組織の開発にも気軽に参加出来る時代なので、いろいろ見て回っていきたいです。

