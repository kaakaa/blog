---
title: "Mattermost 8.0の新機能"
emoji: "🎉"
published: true
date: 2023-07-18T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2023/07/14 に Mattermost の約一年ぶりのメジャーアップデートとなる `v8.0.0` がリリースされました。  
2023/04/14にMattermost v7.10がリリースされてから、5/16リリース予定だった[v7.11のリリースがキャンセルされ](https://docs.mattermost.com/install/self-managed-changelog.html#release-v7-11-feature-release)、v8.0のリリースも当初の予定より1ヶ月延伸されたため、3ヶ月ぶりの新バージョンのリリースとなります。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Expanding the Mattermost ecosystem with new partnerships and integrations](https://mattermost.com/blog/mattermost-v8-0-is-now-available/)
- [Mattermost self\-hosted changelog](https://docs.mattermost.com/install/self-managed-changelog.html#release-v8-0-major-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

また、各見出しにPrefixとしてMattermostの機能分類を記述しています。

- [Channels](https://docs.mattermost.com/guides/channels.html): 従来のチャット機能
- [Playbook](https://docs.mattermost.com/guides/playbooks.html): Mattermost v6.0から追加されたインシデント管理機能
- [Boards](https://docs.mattermost.com/guides/boards.html): Mattermost v6.0から追加されたKanbanボード機能 ([Focalboard](https://www.focalboard.com/))
- [Calls](https://docs.mattermost.com/channels/make-calls.html): Mattermost上で音声通話と画面共有を行うことができるプラグイン
- Platform: 上記機能を含むMattermost全体

## (Platform) プライベートクラウドでのLLM, Azure AI, OpenAIとの連携

MatttermostでもGenerativeAIの活用は[ホットなトピック](https://venturebeat.com/ai/mattermost-brings-generative-ai-power-to-open-source-collaboration/)であり、最近、Mattermostと生成AIを組み合わせたサンドボックス環境を構築することを目的とした[OpenOps](https://github.com/mattermost/openops)が公開されました(以前[`mattermost/ai-framework`](https://github.com/mattermost/mattermost-ai-framework)として公開されていたもの)。[OpenOpsのリポジトリ](https://github.com/mattermost/openops)では、MattermostとLLMを連携する機能一式をセットアップするためのスクリプトが公開されており、手順通りに環境を構築するだけで、MattermostとGenerative AIとを連携させる環境を構築することができます。

既にMattermostのインスタンスを持っている場合、追加で必要なのは[Mattermost AI Plugin](https://github.com/mattermost/mattermost-plugin-ai)と、Mattermost AI Pluginのバックエンドとして呼び出されるGenerative AI系のサービスになるかと思います。バックエンドとして利用できるサービスはOpenAI APIや[Azure OpenAI Service](https://azure.microsoft.com/ja-jp/products/cognitive-services/openai-service)を始め、[Anthoropic](https://console.anthropic.com/login?selectAccount&returnTo=https%3A%2F%2Fconsole.anthropic.com%2Faccount%2Fkeys)、[Ask Sage](https://www.asksage.ai/)などが利用でき、さらにOpenAI APIと互換性のあるAPIであればなんでも利用することができるようです。OpenOpsでは、OpanAIと互換性のあるAPIとして[LocalAI](https://localai.io/)が利用されています。

MattermostからAIを呼び出すポイントは[Mattermost AI Plugin](https://github.com/mattermost/mattermost-plugin-ai)に実装されており、以下のような場面でAIを呼び出すことができます。(詳しい説明は[公式リポジトリのREADME](https://github.com/mattermost/mattermost-plugin-ai#usage)等を参照ください)

| 機能 | 概要 |
|:------|:---|
| `Streaming Conversation` | (これは利用シーンではないですが) AI Botからの返信は、ChatGPTのUIのように文字列ストリームとして受信した単語/ワードごとにMattermost上に表示されていきます |
| `Thread Summarization` | Mattermostのスレッドの内容の要約をAI BotからのDMとして受け取ることができます |
| `Answer questions about Threads` | AI Botからの返信に対して、スレッドから追加の質問を行うことができます |
| `Chat anywhere` | AI Botに@メンションすることで、どこからでもAIに質問することができます　|
| `Create meeting summary` | Mattermost上で音声通話ができる[Calls Plugin](https://github.com/mattermost/mattermost-plugin-calls)と連携することで、ミーティングの要約を生成することができます |
| `React for me` | AI Botがメッセージの内容に沿った絵文字リアクションをつけてくれます (Just for fun!) |
| `RLHF Feedback Collection` | (これは利用シーンではないですが)AI Botからの投稿には👍 👎のアイコンが付いており、fine tuningのためのフィードバックを集めることができます  |

Mattermost AI PluginとLocalAIを組み合わせた環境を試してみましたが、まだそのまま利用するには動作に難がありそうな点と、日本語に関するサポートがまだ十分ではない点には注意が必要かと思いました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/platform-openops-japanese.png)

とりあえずMattermostからOpenAI API(ChatGPT)に質問を投げたいというだけであれば、以前紹介されていた単にMattermostとOpenAI APIを繋ぐ手段も利用できるかと思います。
[Mattermost 7\.10の新機能 - ChatGPT Bot](https://blog.kaakaa.dev/post/mattermost/releases-7.10/#chatgpt-bot)

## (Channels) Apps Barの有効化

Mattermost v7.0でベータ機能として導入された[Apps Bar機能](https://blog.kaakaa.dev/post/mattermost/releases-7.0/#integrations-apps-bar-%E3%83%99%E3%83%BC%E3%82%BF%E7%89%88)がデフォルトで有効化されました。  

Apps Barは、チャンネル画面の右端に常に表示される領域のことで、ここにはインストール済みのプラグインやAppsのアイコンが表示されており、1クリックで各機能の詳細を呼び出すことができます。また、プラグインインストールの権限を持つユーザーには、Apps Barの下部にMattermostアプリマーケットプレースへのリンクアイコンが表示されており、ここからプラグインやAppsを簡単にインストールすることができます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-appsbar.png)

今までのバージョンでは、**システムコンソール > 実験的機能 > 機能 > Apps Barを有効にする** から設定を有効にした場合のみApps Barが表示されていましたが、v8からは **Disable Apps Bar** の設定を有効にしない限りApps Barが表示されます。この設定について、詳細については以下を参照してください。
* [Experimental configuration settings](https://docs.mattermost.com/configure/experimental-configuration-settings.html#disable-apps-bar)
* [Channel Header Plugin Changes \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/channel-header-plugin-changes/13551)

## (Channels) Enterprise/Professional: 優先度:緊急 のメッセージに対する永続的な通知

Enterprise版もしくはProfessional版限定の機能です。  
Mattermost v7.7で追加された[メッセージの優先度を指定する機能](https://blog.kaakaa.dev/post/mattermost/releases-7.7/#channels-%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8%E3%81%AE%E5%84%AA%E5%85%88%E5%BA%A6%E8%A8%AD%E5%AE%9A%E3%81%A8%E6%97%A2%E8%AA%AD%E7%A2%BA%E8%AA%8D)で、優先度が `緊急(Urgent)` に設定された投稿内に @メンション(`@channel`、`@all`、`@here`は対象外) が含まれていた場合、メンションされたユーザーが内容を確認するまで通知が送信され続ける機能(Persistent notifications/永続的な通知)が追加されました。

メッセージ優先度に `緊急(Urgent)` を指定すると、`Send persistent notifications` の設定が表示されます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-persistent-notifications.png)

`Send persistent notifications`を有効にし、@メンションを含むメッセージを投稿しようとすると、以下のような永続的な通知を送るかどうかを確認するダイアログが表示されます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-confirm-notifications.png)

`Send`をクリックすると、@メンションされたユーザーには定期的に通知音が鳴り続けます(一度既読状態にしたメッセージが再度未読状態になったり、DMなどで通知されたりはしないようでした)。この通知は、通知を受け取ったユーザーが確認処理(Acknowledge、返信、リアクションのいずれか)を行うまで通知され続けます。

永続的な通知の頻度や通知の最大送信回数は[システムコンソールから設定することができます](https://docs.mattermost.com/configure/site-configuration-settings.html#persistent-notifications)。デフォルトでは、通知頻度が5分間隔、通知の再度大送信数が6回となっているため、計30分間、5分間隔で通知が実行されることになります。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-persistent-notification-settings.png)

## (Channels) チャンネル内のスレッドの自動フォロー設定

> Added a new option to auto-follow all threads in the channel Notification Preference settings.

Mattermost v7.0でGAとなった[返信スレッドの折り畳み](https://blog.kaakaa.dev/post/mattermost/releases-7.0/#channels-%E8%BF%94%E4%BF%A1%E3%82%B9%E3%83%AC%E3%83%83%E3%83%89%E3%81%AE%E6%8A%98%E3%82%8A%E3%81%9F%E3%81%9F%E3%81%BF%E6%A9%9F%E8%83%BD%E3%81%8Cgageneral-availability%E3%81%B8%E6%98%87%E6%A0%BC)機能ですが、自分がメンションされていないスレッドや自分が投稿したことがないスレッドなどは例え重要な返信があったとしても通知されないため、重要な投稿を見逃してしまうことがありました。それを防ぐために通知を受け取るには、手動でスレッドをフォローする必要がありました。  
今回のバージョンから、手動でフォローしていないスレッドであっても、返信があった場合に更新通知を受け取るかどうかをチャンネル毎に指定できるようになりました。

この設定は、**チャンネルメニュー > 通知の設定 > Auto-follow all new threads in this channel** から設定できます。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-notification-settings.png)

**Auto-follow all new threads in this channel** を有効にすると、自分が参加していないスレッドであっても、誰かから追加の返信があった場合に、左サイドバーの**スレッド**に変更があったことが表示されるようになります。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-auto-follow-thraeds.png)

## その他の変更

### (Channels) チャンネル名に日本語を指定する場合の注意点

現在公開されている`v8.0.0`では、チャンネル作成ダイアログから新規にチャンネルを作成する際に、チャンネル名入力欄で`Enter`キーを押すとチャンネル作成処理が走ってしまうようです。チャンネル名に日本語を使おうとする場合、入力文字の確定のために`Enter`キーを押す必要があると思うのですが、この`Enter`キーを押した時点でチャンネルが作成されてしまう場合があるため、日本語話者としては難のある状況になっているかと思います。

チャンネル名欄に日本語のみ入力している場合は別のエラー(後述)によりチャンネル作成が停止しますが、例えば`team_`のように既に英数字が入力されている際に追加で日本語を入力して`Enter`キーを押すと、その時点のチャンネルURLが既存チャンネルと重複していない限り即座にチャンネル作成処理が実行されてしまいます。

本件についてはIssueを立てていますが、英語話者には問題点が伝わりづらい部分にも思うため、リアクションなりで反応いただけると助かります。

[Unexpectedly fire the handleEnterKeyPress in NewChannelModal with Japanese IME input mode · Issue \#23967 · mattermost/mattermost](https://github.com/mattermost/mattermost/issues/23967)

回避策としては、日本語入力後、確定のための`Enter`キーは押さず、チャンネル名欄以外の箇所をクリックすることで、チャンネルを作成することなく入力文字を確定することができます。

---

チャンネル名欄と日本語については、もう一つ怪しい動きがあります。  
今回のバージョンからチャンネル名に`URL-safe`な文字が1文字も存在しなかった場合に、自動で`チャンネルURL`を入力する機能が実装されました。

> * Implemented URL auto generation on channel creation for when there’s no URL-safe characters on its name.

しかし、試しに日本語のみのチャンネル名を指定してみると、以下のようなエラーとなり、チャンネルが作成できないようです。

> Channel URL must be 1 or more lowercase alphanumeric characters
> チャンネルURLは1文字以上の小文字の英数字にしてください

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-create-failed.png)

自動で入力されている`チャンネルURL`自体は正しい文字列らしく、手動で1文字削除してから同じ文字を再度入力するだけでチャンネルが作成できるようになるので、チャンネル名欄で`Enter`キーを押した場合のチャンネルURLのエラーチェックのタイミングがおかしいようです。  
こちらの回避策も、日本語を入力後、チャンネル名欄以外の箇所をクリックすることで、正しいURLとして認識されます。

### (Channels) 時差の表示

ユーザーアイコンをクリックした時に開くポップアップに、そのユーザーに対して設定されているタイムゾーンと、自分が設定したタイムゾーンの時差が表示されるようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-hours-ahead.png)

上記の例では、タイムゾーンが `UTC+09:00` のユーザーから `UTC±0` のユーザーのプロフィールポップアップを開いた時の画面であり、タイムゾーン的に9時間遅れであることが表示されています。

### (Channels) `CTRL/CMD + K` によるリンクの指定

> * CTRL/CMD + K shortcut can now be used to insert link formatting when text is selected.

メッセージ入力中、入力中のテキストの一部を選択した状態で`CTRL/CMD + K`を押すと、選択部分のテキストをMarkdownのリンク記法に変換できるようになりました。
(テキストを選択していない状態で同じショートカットを実行すると、チャンネル検索ダイアログが開きます。)

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-insert-link.gif)

### (Channels) チャンネル毎に通知音を変更することが可能に

> * Added support to specify different desktop notification sounds per channel.

チャンネル毎に異なる通知音を指定することができるようになりました。これにより、業務上重要なチャンネルのみ通知音を特別なものにするなどの設定ができるようになります。

チャンネルの通知音は、**チャンネルメニュー > 通知の設定 > Desktop Notifications > Notification sound**から設定できます。  
![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-notification-settings.png)

また、ユーザー毎のデフォルトの通知音は **設定 > 通知 > デスクトップ通知 > 通知音** から設定できます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-8.0/channels-notification-default.png)

### (Platform) Teamsとの統合

Microsoft Teams向けのConnectorの更新が行われ、Teamsアプリ内でMattermostを表示することができるようになったようです。  
動作の様子や、設定方法などは以下の動画で確認することができます。

{{< youtube Mg-stF7_Bjk >}}

{{< youtube MxegqqfErEg >}}

### (Platform) Atlassian製品との連携
Jira、Confluence、Trello等のAtlassian社サービスとの連携が強化されたようです。  
こちらについてはどのような点が更新されたかの説明はありませんが、以下の専用ページが紹介されています。

[Mattermost for Atlassian Suite](https://mattermost.com/atlassian/)

### (Platform) PostgreSQLが推奨に

> * To simplify management and scalability challenges, Mattermost 8.0 recommends deploying PostgreSQL over MySQL.

Mattermost v8では、パフォーマンスや効率の観点からMySQLよりPostgreSQLの方を推奨する方針のようです。

## アップグレード時の注意事項

ここではアップグレード時の注意事項の一部について触れます。すべての注意事項を確認する場合は、公式ドキュメントを確認ください。　 
[Important Upgrade Notes](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

### インサイト機能が非推奨に
Mattermost v7.1から導入された[ワークスペース内の活動状況を確認できるインサイト機能](https://blog.kaakaa.dev/post/mattermost/releases-7.1/#professionalenterprise-platform-%E3%82%A4%E3%83%B3%E3%82%B5%E3%82%A4%E3%83%88-%E3%83%99%E3%83%BC%E3%82%BF%E7%89%88)が非推奨機能となりました。この決定は、インサイト機能にパフォーマンス上の問題が見つかったことが理由と説明があり、環境変数`MM_FEATUREFLAGS_INSIGHTSENABLED`に`false`にセットして機能を無効化することが推奨されています。  
[Proposal to revise our Insights feature due to known performance issues \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/proposal-to-revise-our-insights-feature-due-to-known-performance-issues/16212)

今後、インサイト機能から派生した機能として、AI framework ([mattermost/openops](https://github.com/mattermost/openops)?)を活用した機能の開発を検討しているそうです。

### プリインストールされていたいくつかのプラグインがデフォルトで無効に

Mattermostにプリインストールされていた以下のプラグインが、デフォルトでは無効化されて起動するようになりました。有効化するには、**システムコンソール > プラグイン設定** から起動する必要があります。

* Focalboardプラグイン (Mattermost Board)
* Channel Exportプラグイン
* Appsプラグイン

### Go開発向けのPublicサブモジュールの公開

Go言語を使ってMattermost関連の機能開発を行う際に利用できる、公開APIをまとめた`public`サブモジュールが公開されました。  
[mattermost/server/public at master · mattermost/mattermost · GitHub](https://github.com/mattermost/mattermost/tree/master/server/public)

ただ、こちらはまだ安定したAPIではなく、破壊的な変更が追加される可能性があり、バージョンニングもMattermost Serverとは連動していないようなので、利用の際には注意が必要です。  
Mattermost v8に対して`public`サブモジュールは v0.5.0 が設定されており、追加のリファクタリング等を加えた後に、今年の後半にv1としてリリース予定のようです。

### PostgreSQL v10のサポート廃止

PostgreSQL v10のサポートが廃止され、最低バージョンとしてv11以上が必要となりました。

### 「プラグインが使用するReact DOMにパッチを当てる」設定の削除

Mattermost v7.7でフロントエンドで利用しているReactのバージョンを17にあげた影響で、古いバージョンのReactDOMを利用してビルドされているプラグインを利用するとクラッシュする恐れがあり、Mattermost側で自動でパッチを当ててくれる `ExperimentalSettings.PatchPluginsReactDOM`という設定が存在しました。今回のバージョンからこの設定が削除されるため、Mattermostアップグレード前に使用しているプラグインが最新バージョンにアップデートされているか確認する必要があります。

## その他のトピック

### Mattermost Academy

MattermostではMattermostに関する知識を学べるプラットフォームとして [Mattermost Academy]()というサービスを提供しています。  
ここに以下の二つのコースが追加されたというアナウンスがありました。

**[Mattermost End User Onboarding](https://academy.mattermost.com/p/mattermost-end-user-onboarding)**

Mattermostのオンボード用の動画コースで、Mattermostの基本的な使い方から各種Tips等について30分程度で学ぶことができます。

**[Use case training](https://academy.mattermost.com/courses/category/use-case-training)**

こちらは受講できていませんが、"Agile software development"や"Cybersecurity incident management"などの様々なユースケースにおいて、Mattermostをどのように活用すれば良いかということについて学べるコースになっています。こちらは動画コースではないようです。(動画コースもあるかもしれませんが)

## おわりに
次の`v8.1`のリリースは 2023/08/16(Wed)を予定しています。  

`v8.1`は延長サポート(Extended Support Release)対象のバージョンであり、2024/05/15までサポートされる予定です。  
[Release Lifecycle](https://docs.mattermost.com/upgrade/release-lifecycle.html)
