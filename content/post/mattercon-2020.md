---
date: "2020-02-28"
slug: "mattercon2020"
tags: ["mattermost", "mattercon"]
title: "MatterCon2020 Report"
---

Mattermost公式ブログでも公開されました。  
https://mattermost.com/blog/going-to-mattercon-as-a-mattermost-contributor/

## Summary

- 概要:
  - 中南米・バハマで開催されたMattermost社のコミュニティイベント **MatterCon2020** に招待されたので行ってきました
- 目的:
  - 建前: OSSネイティブな企業の開発プロセス・営業の仕方などに関する知見の収集
  - 本音: バハマとか人生で行くことないだろうし、せっかく招待されたのだから行ってみたい
- 感想:
  - MatterConはMattermost社のメンバ、コミュニティメンバともに気軽に話しかけられる貴重な機会。CEO/CTOとも気軽に話ができる。
  - Mattermostはコミュニティを大事にする会社でありゲストとして至れり尽せりでした
  - 技術の話が主体ということはなく、参加者同士のコミュニケーションを重視しており、Mattermost社員、コントリビュータに関係なく様々な人と話すことができました
  - Fun Cultureを作り様々な背景を持つ人々をコミュニティとして巻き込んで規模を大きくしていくことが、OSSネイティブな企業の特性を活かしたビジネスの仕方だと感じた

## 背景

まず、MattermostおよびMatterConについてと、私がMatterConに参加するまでの経緯を述べます。

### Mattermostとは
[Mattermost](https://mattermost.com/)はSlack Alternativeを標榜している、オープンソースのチャットツールです。SlackやTeamsと違い、データも含めたチャット環境すべてを自分たちの管理下に置けることを重視しており、SaaS展開などを行っていないのが特徴です。

昨年、TeamsのDAUがSlackを超えた[^about-mattermost-teamsdau]というニュースが流れるなど、ユーザー数的にはSlack/Teamsが頭ひとつ抜けている感じですが、チャットのような組織内のコミュニケーションに関するデータを組織外の環境に置くことに懸念を持つようなセキュリティ意識の高い組織などで、Mattermostのようなオンプレミス環境に構築できるチャットツールの採用が増えています。その需要の大きさを示すように、昨年、Mattermostは昨年$50Mの投資を受け Series B[^about-mattermot-seriesb] となりました。

[^about-mattermost-teamsdau]: https://jp.techcrunch.com/2019/07/12/2019-07-11-microsoft-says-its-slack-competitor-teams-now-has-13-million-daily-active-users/
[^about-mattermot-seriesb]: https://mattermost.com/blog/yc-leads-50m-series-b-in-mattermost-as-open-source-slack-alternative/

### MatterConとは?

MatterConとはMattermost社が主催するイベントであり、Mattermost社のハンドブックでは以下のように定義されています。

https://handbook.mattermost.com/contributors/mattercon
> We try to get together every 12 months or so to get face-time with one another, build community, reinforce our leadership principles, and get some work done! Since our team is scattered all over the globe, we try to plan a different location for each MatterCon.
> MatterCon is not a mandatory trip or a department offsite, nor is it a vacation or incentive trip. It's a chance for everyone to meet fellow Mattermost team-members across all departments and regions: part team building, part education, part community building, and hopefully all fun. This is a work trip with fun mixed in.

Mattermostはフルリモートの会社のため、12ヶ月ごとに一同に介して顔を合わせる機会を作っており、それを **MatterCon** と呼んでいます。
また、Mattermostは、会社のミッションの一つとして **オープンソースコラボレーション** を掲げており[^about-mattercon-mission]、その一環としてMatterConにコミュニティメンバーを招待しています。

一昨年までは参加者の規模が2,30人程度だったようですが、昨年のMatterConからは100人ほどの規模で開催されているようです（正確な人数については把握しておりません）。過去のMatterConの模様はYouTubeで公開されています。

* [Mattermost 2017 Core Committers Summit & Hackathon \- Las Vegas \- YouTube](https://www.youtube.com/watch?v=_RpmrM-5UFY&t=65s)
* [Mattermost \| MatterCon 2018 \| Community Meetup in Lisbon \- YouTube](https://www.youtube.com/watch?v=CZXaYttz3NA)
* [Mattermost \| MatterCon 2019 \| Community Meetup in Dominican Republic \- YouTube](https://www.youtube.com/watch?v=pMySvCfy7Bw)

[^about-mattercon-mission]: https://docs.mattermost.com/process/handbook.html#mission

### 参加経緯
私は現在、Mattermost公式のコミュニティメンバとしてMattermostの日本語化[^about-partition-translation-maintainer]、Twitter Mattermost日本語アカウント([@mattermost_jp](https://twitter.com/mattermost_jp))の管理などを行っています。また他にも、Mattermost関連の記事投稿(blog[^about-partition-blog], Qiita[^about-partition-qiita])やHelp Wantedなどの軽微なIssueの解決などを行っています。もう2年も前の話ですがMattermost v4.6がリリースされた際にはMVP[^about-partition-mvp]にも選出していただきました。

{{<figure src="/images/posts/mattercon-2020/mvp.jpeg" alt="MVP" height="640px">}}

Mattermostへの貢献を始めたきっかけは、2016年ごろに職場でMattermostを使い始めたことに起因します。
その頃、ChatOpsが流行りはじめていたということもあり、オンプレミスで運用できるようなOSSのチャットツールを探していました。Mattermostを使い始める前にも[Kandan](https://github.com/kandanapp/kandan)や[Let's Chat](https://github.com/sdelements/lets-chat)などのOSSのチャットツールを試していたのですが、それぞれ他システムとの連携機能が不十分であったり、開発が活発でなかったりと問題が見えていました。そんな中、会社の同期が見つけてきたMattermostは、当時(`v3.4`ぐらい？)から機能も豊富で開発も非常に活発であり、とても魅力的なソフトウェア/プロジェクトでした。Matterostでいろいろ遊んでいるうちに小さな翻訳ミスを見つけ、元々OSSへの貢献に興味があった私は公式ドキュメントで貢献の仕方を調べ、翻訳の修正を送ったことがMattermostへの貢献の始まりでした。

当時、一つ小さな修正を送るだけでコアメンバーの方が[コミュニティ用のMattermost](https://community.mattermost.com/)上ですぐにコメントをくれたことも非常に印象的でした。
さらに、MattermostはGolangやReact.jsなどの技術を利用していることも自分の興味を惹き、定期的にTシャツなどのSwagを定期的に送ってくれたりしたこともあり、段々と貢献の幅を広げていきました。

そして今回、MatterConへ招待いただくこととなりました。

[^about-partition-translation-maintainer]: https://docs.mattermost.com/developer/localization.html#translation-maintenance
[^about-partition-blog]: https://blog.kaakaa.dev
[^about-partition-qiita]: https://qiita.com/kaakaa_hoe
[^about-partition-mvp]: https://developers.mattermost.com/contribute/mvp/

## MatterCon出発前

ここでは、Mattermostの概要や出発前のやり取り、旅程について簡単に述べます。

### 概要

- 期間: 2020/2/22(Sat), 23(Sun)
- 場所: [Hotel Melia Nassau Beach \- All Inclusive, in Nassa](https://www.melia.com/en/hotels/bahamas/nassau/melia-nassau-beach/index.htm)
- 参加者: 100人ぐらい？
    - Mattermost Inc.
    - Partner / Contributor
    - SOs (Significant Others, 奥さんなど)

### 開催前

MatterCon開催前の動きとしては、おおよそ下記のような流れでした。

- (11月下旬) コミュニティMattermostで招待のDMが届く
- (1月上旬) 利用するフライトについて連絡 -> 予約されたフライト情報が送られてくる
- (1/24) MatterCon Guest用チャンネルが作成される
  - MatterCon開催前、開催期間中を通じて必要な情報はここに投稿される
- (2/4) Conference Agenda発表
  - MatterCon開催中のスケジュールなど
- (2/10) Guest Agenda発表
  - ゲスト向けの詳細ドキュメント(空港からのシャトルバスの乗り方など)

### 旅程

旅程は以下の通りでした。バハマまでは直行便がないため、帰りは乗り継ぎ14時間空いたりするなど足掛け30時間以上かかりました...(流石にホテルを予約しましたが)。

- 2/21: NRT - DFW - NAS
  - 乗り継ぎの待機時間(4h)含めて移動時間は20時間ほど
  - ナッソーの空港から会場のホテルまではMattermost社手配のシャトルバスで10分ぐらい
- 2/22: MatterCon Day1 (Unconference, Activitiesなど)
- 2/23: MatterCon Day2 (Hungout, Partner Session, Lightning Talks, Dinnerなど)
- 2/24: NSA - DFW
  - 乗り継ぎで 14時間空くのでDFW泊
- 2/25-26: DFW - NRT
  - DFWでの一泊も含めると、ナッソーから成田までは31時間ぐらい...

## MatterCon

ここでは、MatterConの中で私が体験したことを述べます。

### MatterCon Day 0 (Arrival)

自宅から成田空港へ行き、成田(NRT)からダラス・フォートワース空港(DFW)へ。DFWで4時間ほど時間を潰しAmerican Airlinesでバハマ・ナッソーにあるリンデン・ピンドリング空港(NAS)へ。
コロナウイルスの影響もあり検疫等で時間がかかることを予想していたましが、拍子抜けするほど何もなく。出入国時に過去14日以内の中国滞在歴を聞かれたぐらいでした。
DFW - NAS のフライトの途中、バハマの入国カードが配られたのですが、その紙は2枚綴りになっており、2枚目は出国時に提出するバハマのアンケートになっている模様。さすが観光立国。

ナッソーのリンデン・ピンドリング空港からはシャトルバス。到着ターミナル出口には観光会社の方だと思われる人たちが10人以上。出発前に送付されたGuest Agendaに記載の通り、現地の観光会社であるMajestic Toursのマークを掲げている人を見つけて声かけ。陽気なおっちゃんに誘われてシャトルバスへ。その間に、Buddyの Miguel ([@mgdelacroix](https://github.com/mgdelacroix)) からMattermostのDMで「ホテルに着いたら会おう」というメッセージが来ているのに気付き、「ちょうどバス乗ったところだから数十分ぐらいで着くよ」と返信。MatterConでは、私のようなゲスト一人一人に対して滞在中に分からないことがあった時の相談相手としてMattermost社の人をBuddyとして充てがってくれていました。

ホテルまではバスで10分ほど。バスから見える風景は、道路脇に椰子の木🌴が並ぶまさに南国リゾートのそれでした。(写真看板奥の建物は別のホテル)

{{<figure src="/images/posts/mattercon-2020/mattercon-day0-bahamas.jpeg" alt="bahamas" height="480px">}}

17:00ごろにホテルに到着してそのままチェックイン。受付の人も適度にフレンドリーで好印象。特に問題なくチェックインを済ませ、ロビーでBuddyのMiguelを探すとGitHubのアイコン通りの見た目だったのですぐに見つかり挨拶。ちなみにホテルのロビーはこんな感じでした。

{{<figure src="/images/posts/mattercon-2020/mattercon-day0-lobby.jpeg" alt="lobby" height="480px">}}

Miguelがホテル内を案内してくれるとのことなので、簡単に歩き回りながら案内してもらいました。案内してもらってる間、Miguelの持っているアタリのバッグばかりが目に入る...。聞いてみるとVideoゲームが好きとのこと（その後、日本のアニメ・漫画の知識も豊富であることが判明）。

案内の最後に部屋まで行く途中、「18:00から4Fの日本料理店でDinnerがあるけど、疲れてるなら休んでても良いし来たいなら来ても良いし自由にして」と労いの言葉をかけてくれました。とりあえず部屋に入って荷物を置いてシャワーを浴びていたら18:00を過ぎていましたが、せっかくの機会なのでDinnerの会場へ。

受付の人に何といえば良いか迷いましたが、「Mattermost社の人いますか？」という聞き方をしたら通じたようで案内してもらいました。ただ、案内してもらったテーブルにBuddyの人はおらず...テーブルの方の何人かがMattermostのTシャツを着ていたのでBuddyとは別テーブルに来てしまったと理解。少し戸惑っていると非常にフレンドリーにWelcomeしてくれたので空いてた席に着席。日本語を少し喋れる人がいたり、日本料理ということでちょっと会話に参加したり。英語が結構わからず戸惑いましたが楽しい雰囲気だったので楽しかった。2時間ほどしたら流石に眠くなってきたので先に退席して就寝。

部屋に戻ってから、部屋の外を見ると疑いようもないリゾート感。パシャリ。

{{<figure src="/images/posts/mattercon-2020/mattercon-day0-hotel.jpeg" alt="day0-hotel" height="480px">}}

Dinnerでは会えなかったけど、別テーブルで楽しんでたことをMiguelに連絡。

[^mattercon-day0-immigration-card]: https://www.toppantravel.com/files/ed/018.pdf

### MatterCon Day 1 (Unconference, Activities...)

昨日の連絡の続きでMiguelと朝食に行く流れになったので、ありがたく同行。朝食はバイキング。色んな種類があり、なかなか美味。食事の途中でシンガポールのコントリビュータ [@liusy182](https://github.com/liusy182) とも同席したので初めましての挨拶。

#### Greeting
朝食の後は、ホテル内の会議室でGreeting。

Greetingの前に会場前でノベルティを配っていたので、Tシャツとステッカー、ノートとペンをもらいました。また、その場に[@jasonblais](https://github.com/jasonblais)がいたので挨拶。彼は、私が初めてコントリビュートをした時にメッセージをくれたり、その後もよく気にかけてくれてコメントをくれたりしていたので、今回のMatterConで感謝を述べたい人の一人でした。彼がいなかったらここまでコントリビュートをしてなかったと思います。このタイミングではあまり時間もなかったので、後で改めて感謝の意を伝えました。

Greeting会場は、会議室らしく机と椅子が並べられている部屋。Mattermost社の人たちは数日前からプライベートなカンファレンスを行ってたらしいが、この日からコミュニティメンバーも参加するということで改めてCEO/CTOから挨拶。

{{<figure src="/images/posts/mattercon-2020/mattercon-day1-greeting.jpeg" alt="day1-greeting" height="480px">}}

#### Unconference

挨拶の後はUnconference。いくつかのトピックについて円座になってディスカッションすると言うもの。まず、Greetingと同じ部屋でディスカッションのトピックとそれを担当するMattermost社の開発者が紹介され、興味がありそうなところの議論に参加するという形。トピックはPlatform, Security, DevOps, Kubernetes, Mobile, Integrationsなどなど7,8トピックぐらい？ありました。Miguelの誘いもあり、Platformのディスカッションに参加。会話の内容としては、モバイルアプリでネットワーク切断中に投稿されたメッセージのWebSocketをどう扱うか、や、GraphQLの導入についての話でした。

#### Group Photo
ランチの前にMatterCon2020用の黄色いTシャツを着てグループフォト撮影。

{{<figure src="/images/posts/mattercon-2020/mattercon-day1-tshirts.jpeg" alt="MatterCon2020 T-shirt" height="480px">}}

#### Lunch
Unconference後はランチ。自分と同じ日本人コントリビュータの [@maruTA-bis5](https://github.com/maruTA-bis5)とご一緒させてもらう。
朝食と同じ場所と聞いていたけど、そのレストランが空いておらず、近くで掃除してた人に聞いてビーチ近くのレストランを紹介される。
ビーチのレストランで受付をしようとしたところ、昨年日本で行われたMattermostイベント [^mattercon-day1-nri]で知り合ったNRIの方々や、その時に来日していたMattermost営業の[Steve](https://twitter.com/stephenlgreen)がいたので同席させてもらう。ありがたい。
日本人コミュニティとして初めましてをしていたが、次のアクティビティの集合時間が迫っていたり、忙しいのか中々オーダーを取りに来てくれないということもあり慌ただしくランチを済ませることに...。

#### Activities
１日目の午後はActivity。
出発前のアンケートで下記の中から参加したいActivityに関するヒアリングがありました。

- [Ardastra Gardens, Zoo and Conservation Centre](http://www.majestictoursbahamas.com/tours/?id=03)
- [City and Country Tour](http://www.majestictoursbahamas.com/tours/?id=01)
- [Swimming with pigs](https://www.travelandleisure.com/trip-ideas/bahamas-swimming-pigs-big-major-cay?slide=443e1f7d-e5c1-4fa1-87e3-33cbaa112afb#443e1f7d-e5c1-4fa1-87e3-33cbaa112afb)
- matterconandchill at the resort, swim in the pools, paddleboard, snorkelling or kayak at the private beach

街中ブラブラするのが好きというのと、何かお土産になりそうなのを探せないかということでナッソーの街へ行くというCity and Country Tourを選択。

指定された時間にロビーへ行くと、朝食で一緒だった[@liusy182](https://github.com/liusy182)がいたので一緒に行動。
まずはシャーロット砦[^mattercon-day1-charlotte]へ。シャーロット砦には説明員の方がおり、その方に砦の中を案内してもらう。説明員の方が一々オーバーリアクションで面白かった。シャーロット砦自体はそんなに広くないので、2,30分ぐらいで一通り案内も終え、最後は屋上を散策しながら写真を撮ってバスへ。高台に建てられてる建物といこともあり眺めがとても良かった。

{{<figure src="/images/posts/mattercon-2020/mattercon-day1-activities1.jpeg" alt="charlotte" height="480px">}}

次はナッソー市街へ。市街をバスで走った後、ラムケーキ屋[^mattercon-day1-rumcake]さんで停車。バハマではラムケーキが有名らしい。中に入って試食。甘さとアルコール感が強め。さぁ、ここから自由時間...と思ったら、ナッソー市街で停車するのはココだけらしい。バスで回るだけだったのね...お土産どうしよ。そう言ってても始まらないのでバスに乗ってまた移動。写真はナッソー市街にあるParliament Squareの銅像。

{{<figure src="/images/posts/mattercon-2020/mattercon-day1-activities3.jpeg" alt="city" height="480px">}}

海岸線を20分ほど走った後、海岸に到着。綺麗な景色。ブラブラ歩いて写真を撮る。一通り歩き回って出発待ってると、Microsoftから数ヶ月前にMattermostにジョインしたという方が話しかけてくれたので少し会話。
そしてそのままバスに乗ってホテル着。
ナッソー市街歩けなかったのは残念だけど、色んな景色を見れたので満足でした。
{{<figure src="/images/posts/mattercon-2020/mattercon-day1-activities4.jpeg" alt="sea" height="480px">}}

[^mattercon-day1-nri]: https://www.sbbit.jp/eventinfo/52965/
[^mattercon-day1-charlotte]: https://www.bahamas.com/vendor/fort-charlotte
[^mattercon-day1-rumcake]: https://www.thebahamasrumcakefactory.com/


### MatterCon Day 2 (Hungout, Partner Session, Lightning Talks...)

#### Hungout
午前中はHungoutという名の自由時間。
[Matterpoll](https://github.com/matterpoll/matterpoll)というMattermostの投票機能プラグインを一緒に開発しているMattermost社の [@hanzei](https://github.com/hanzei) とIssueのトリアージ。[@hanzei](https://github.com/hanzei) とは前身の [matterpoll-emoji](https://github.com/matterpoll/matterpoll-emoji) の頃から3年近くGitHub上で開発をしており、その方とリアルで顔を合わせる機械などは想像していなかったので不思議な気持ちでした。

途中、Partner Sessionへ招待されたのでそちらへ。

#### Partner Session (NDA)
内容はNDAらしいのでノーコメント。

Partner Sessionの後は午前のHungoutで話をした問題について解決策を探る（結局うまく解決はできなかった）。部屋でコード書いていると、コントリビュータの[@cometkim](https://github.com/cometkim) からランチのお誘いが。面識はなかったけど、名前は知っていたので快諾。

#### Lunch
[@cometkim](https://github.com/cometkim)とBuddyの[@wiggin77](https://github.com/wiggin77)と一緒にランチビュッフェへ。途中でCTOの[@coreyhulen](https://github.com/coreyhulen)も一緒に。
[@cometkim](https://github.com/cometkim)に話を聞くと、彼が最初にMattermost本体にPullRequest(CJK関連)[^mattercon-day2-cjk]を送った時、そのPRに自分がコメントをしたのを覚えてくれていたらしく、今回ランチに誘ってくれたそう。その頃はちょうど自分もMattermost本体へのPRをはじめた頃で、CJK関係のPRだったので自分の環境で動かしてコメントした覚えがあります。コメントした時は、まさかこんな形で対面するなんて想像もしていませんでした...不思議なものです。

#### Lightning Talks
ランチの後は会議室でライトニングトーク。
参加者は開発者だけではないので、ライトニングトークの内容はツール/フレームワークの話から営業向けサービス、コミュニティに対する想いなど様々でした。途中、Zoomで発表しているコミュニティの方がおり、そのセッションは非常に盛り上がっていました。

{{<figure src="/images/posts/mattercon-2020/mattercon-day2-lt.jpeg" alt="day2-lighningtalks" height="480px">}}

| Topic | Presenter | Notes |
|:------|:----------|:------|
| Maybe somthing about Airplanes | Christopher Speller | |
| mattermost-govet (Keeping code clean and consistent) | Jesús Espino | |
| Designing for Failure in Engineering Systems | George Goldberg | |
| When Marketing Met Automation: A Love Story | Kendall Reicherter | Marketing automation at a high-level; best practices; common use cases |
| Security smells in code reviews | Juho Nurminen | |
| Current state of React Native | Miguel Alatzar | |
| git commit -m ""Contributing to Mattermost" | Rajat Varyani (Community Member) | |
| Impact of good open source projects on building experience and confidence | Allan Guwatudde (Community) | This talk will mainly be about my journey as a self-taught developer and how working on the mattermost project as a contributor has helped me learn, build experience and confidence. |
| Is there life after Redux/REST? POC Mattermost GraphQL | Eli Yukelzon	| |
| Matrix testing Mattermost permissions | Martin Kraft | |
| Default: Open? | Jesse Hallam | |
| Folded Reply Threads update | Adam Clarkson | this project is just getting started so ~10 min will be plenty to give folks an update on this highly anticipated feature |
| Introduction to Open Source | Shota Gvinepadze | University course build around contributing to Mattermost |
| Mattermost and Vulnerabilities | Aaron Rothschild | |

やはりオープンソースコミュニティの集まりであるからか、コミュニティやオープンさに関する発表が多かった印象があります。個人的には、Mattermostで最も要望の多い[^mattercon-day2-foldedreply]Folded Replyの機能に関する紹介が興味深かったです。


#### Dinner at beach
最終日のDinnerはビーチで。
最後にまたグループフォトを撮ってからディナーへ。

グループフォト撮影後、砂浜に設置されたテーブルに着席。
CEO/CTOなどから挨拶があり、貢献してくれた方々にプレゼントがあるということで名前を呼ばれ取りに行く。
貰ったのは、PatagoniaのダウンやTシャツ、ステッカー、ボールペンなどのSwagでした。Patagoniaのダウンをプレゼントできるとは強いな...。

{{<figure src="/images/posts/mattercon-2020/mattercon-day2-swag.jpeg" alt="swag" height="480px">}}

ビュッフェスタイルで食事を取り、食べ、飲み、眠くなったので就寝。夜のビーチは風も強くて少し肌寒く、貰ったPatagoniaのダウンを早速着たり、会場にあった焚き火の周りにどんどん人が集まっていました。

{{<figure src="/images/posts/mattercon-2020/mattercon-day2-dinner.jpeg" alt="day2-dinner" height="480px">}}

[^mattercon-day2-cjk]: https://github.com/mattermost/mattermost-server/pull/4555
[^mattercon-day2-foldedreply]: https://mattermost.uservoice.com/forums/306457-general/suggestions/19572469-make-threads-collapsible


### MatterCon Day 3 (Departure)
3日目は特に何も予定がなく、2日目にコミュニティMattermostのGuest用チャンネルで通知のあった各ゲストのシャトルバスの時間になったらロビーに集まって空港へ向かうだけ。
他の方の出発時間にロビーにいって挨拶したり、ロビーで作業しながら時間を潰し、バスで空港へ。

空港でチェックイン前に売店でラムケーキをいくつか買い、荷物に詰めてチェックイン（バハマは消費税が12%...）。
チェックイン時（だったかな？）に入国カードに付いてた出国カード（バハマについてのアンケート）の提出を求められたので渡す。

後は空港内で時間を潰してDFWへ。
DFWでは乗り継ぎに14時間空くのでBaggage Claimで荷物を受け取り、滞在予定のホテルであるHyatt Regency DFW[^mattercon-day3-hyatt]へ。
このホテルはターミナル内にあるのですが少し行き方が特殊で、C18あたりのゲートから駐車場に向かって直進していると"Hyatt Regency"と書かれた案内板があるので、階段を上り下りしながらそのままずっと直進することでたどり着けます。
ターミナルCのBaggage Claimからシャトルバスを依頼できる電話があるらしいのですが、歩いても10分は掛からないので荷物を持って階段を上り下りできるなら歩きで行った方が楽に感じました。

そのままホテルに宿泊し、翌日の飛行機で成田空港まで帰ってきました。時差の関係もあり、ナッソーからの帰りのフライトは2泊3日の長旅となりました...。

[^mattercon-day3-hyatt]: https://www.hyatt.com/ja-JP/hotel/texas/hyatt-regency-dfw-international-airport/dfwap

## おわりに

2日間と短い間でしたが非常に楽しく、また、自分の興味から始めたOSS活動が現実のイベント参加に結びつくという非常に興味深い体験でした。

イベント全体を通して、コミュニティメンバそれぞれにBuddyを付けてコミュニケーションを取りやすくしたり、一人でいるとMattermost社の人が積極的に話しかけてくれていたりと、Mattermostがコミュニティメンバをゲストとして大切に扱おうという意識を強く感じました。
話しかけてくれた方には数ヶ月前にMattermostにジョインしたと話している方も多く、会社の規模がまだまだ大きくなっていっているという点も実感できました。また、CEO/CTOなどは動画等ではよく拝見していましたが、普通に食事も一緒にすることができたりと簡単にアクセスできすぎて非常に不思議な感じでした。

イベントの雰囲気としては、OSSコミュニティの集まりということで技術方面の意識の強いイベントかと思っていましたが、開発職以外の方々も多く参加しており、また、Significant Othersとして家族を同伴している方も数多くいたため、全体としては開発以外の雰囲気も強かったように感じました。もちろんHungoutの時間などに開発者を捕まえて議論をすることもできるので、開発に関する議論を沢山したい場合はそのようにできたのだと思います。その点については、もう少し英語力を磨いておけばより深い議論もできたはずであり悔いの残る点でした。

ロケーションについては、リゾート地らしく基本的にはTシャツ一枚で過ごせる気候で過ごしやすかったです（雨が降ったあとなどは肌寒かったりもしましたが）。現地の人々もリゾート地らしく陽気な方々が多く、ホテルの近くを移動するだけならば治安的な心配は全くありませんでした。また、会場である [Melia Nassau Beach](https://www.melia.com/en/hotels/bahamas/nassau/melia-nassau-beach/index.htm) は All Inclusive ということで、ホテル内の飲食が自由であるためにホテル外に出る必要がなかったのも治安が良いと感じる一因かと思います。後、当然ですがホテルのビーチは非常に綺麗でした。

{{<figure src="/images/posts/mattercon-2020/mattercon-beach.jpeg" alt="Beach" height="480px">}}

また来年も招待していただけるよう、時間を見つけて貢献を積み重ねていければと思います。

## Photos

https://www.dropbox.com/sh/clhwtcu1q370trf/AAARbty5pp7lhVCfQgsVyXq9a?dl=0
