---
title: "Mattermost AI Pluinを試してみた"
emoji: "🤖"
published: true
date: 2023-07-23T23:30:00+09:00
toc: true
tags: ["mattermost", "plugin", "ai", "LLM"]
---

# Mattermot AI Plugin

Mattermostから生成的AIサービスを呼び出すためのプラグインとして、Mattermost社により**Mattermost AI Plugin**の開発が進められています。

https://github.com/mattermost/mattermost-plugin-ai

**Mattermost AI Plugin**はMattermostプラグインとして開発されているため、単にMattermost上でAI Botに質問するという使い方だけでなく、Mattermost本体の機能やプラグインと連携した形で利用することができるようです。この記事では、**Mattermost AI Plugin**で実行できる機能を一通り触ってみた記録です。

※ Mattermost AI Pluginは、まだ開発が始まったばかりのプロジェクトであり、この記事に書いた内容にも今後変更が加わる可能性がある点に注意して下さい。この記事では [Commit: 4b53767d6acadb290be1266f288e9d4b755fc6c4](https://github.com/mattermost/mattermost-plugin-ai/commit/4b53767d6acadb290be1266f288e9d4b755fc6c4) (2023/07/20)時点のプラグインを使用しています。

## **Mattermost AI Plugin**インストール

Mattermost AIは開発段階のため、手動でインストールする必要があります。  
インストールするには、以下のリンクから現在の最新版である`mattermost-ai-0.2.0.tar.gz`をダウンロードして、Mattermostへアップロードします。

https://github.com/mattermost/mattermost-plugin-ai/releases/tag/v0.2.0

初めてMattermostプラグインを手動でアップロードする場合は、以下の記事が参考になるかと思います。  
[\[Mattermost Integrations\] Plugin基本編](https://zenn.dev/kaakaa/articles/qiita-20201217-c190c104f44199da02ad#mattermost-plugin-%E3%81%AE%E4%BD%BF%E7%94%A8)

### 注意

**Mattermost AI Plugin**は、データベースにMySQLを使用したMattermostでは使用することができません。
過去にMySQLを使用した大規模なMattermostインスタンスで多くのパフォーマンス問題が発生していたことを受け、Mattermostでは将来的にPostgreSQLのみのサポートを行なっていく方針となったようで、この方針が**Mattermost AI Plugin**でMySQLをサポートしない判断の元となっているようです。その辺りの背景については以下のIssueで述べられています。

https://github.com/mattermost/mattermost-plugin-ai/issues/19

**Mattermost AI Plugin**を使用したサンドボックス環境を簡単に構築するため[mattermost/openops](https://github.com/mattermost/openops)が公開されているため、試してみるだけであればこちらを使用する方法もあります。


## **Mattermost AI Plugin**の設定

**Mattermost AI Plugin**のインストールができたら、プラグインから生成的AIサービス呼び出すための設定を行います。設定は、**システムコンソール > プラグイン > Mattermost AI Plugin**から行います。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-settings.png)

今回は、生成的AIサービスとして OpenAI API を利用し、モデルは`gpt-3.5-turbo`を使いました。

| Name | Value |
|:-----|:------|
| プラグインを有効にする | `有効` |
| OpenAI API Key | `sk-.....` |
| OpenAI Default Model | `gpt-3.5-turbo` |
| AI Large Language Model service | `Open AI` |

上記の他にも、別の生成的AIサービスを利用する場合の設定や、**Mattermost AI Plugin**の機能に関する設定項目等がありますが、とりあえず上記のみ設定すれば一通りの機能は使えるはずです。

## **Mattermost AI Plugin**の機能

ここからは、**Mattermost AI Plugin**の機能について紹介していきます。


### AIとのチャット (Chat anywhere)

AIに対して質問をして、AIから回答を得るというChatGPTなどでお馴染みの機能です。

Mattermost上のチャンネル、DM、グループメッセージなど、どこででも `@ai` にメンションすることで、AIとのチャットを開始することができますです。また、スレッド内で `@ai` によるメンションを行うと、スレッドの会話に基づいたやり取りを行うことができます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-chat.png)

上記の例では、一つ目の回答内容について、箇条書きの項番のみを指定する形で追加の質問を行いましたが、ちゃんと指定した項番の内容を理解して回答をしてくれました。

### スレッドの要約 (Thread Summarization)

Mattermost内の会話スレッドの内容を要約してくれる機能です。

サンプルとして以下のMattermost Forumの質問をMattermost上のスレッドに転記し、そのスレッドの要約を生成させてみました。  

[Font of Mattermost app changed accidentally \- Troubleshooting \- Mattermost Discussion Forums](https://forum.mattermost.com/t/font-of-mattermost-app-changed-accidentally/16347)

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary1-post.png)

質問内容は、MattermostデスクトップアプリでMattermostを開いた場合のみ、投稿のフォントが変わってしまい直し方がわからないというもので、それに対する回答として「`C:\windows\fonts`フォルダ内を確認し、怪しいものがあればリネームして治らないか確認してほしい」というコメントが付き、それを試したら解決した、という質問→回答→解決の3メッセージのみの流れです。

このスレッドの要約として、以下のような文章を得ました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary1-answer.png)

> The thread discusses an issue where the font of the Mattermost app was accidentally changed. The user is seeking information on how to change it back. They mention that the issue only occurs in the Mattermost app on Windows and not in any other apps or in the web browser. Reinstalling the Mattermost app does not resolve the issue. Another user suggests checking the C:\windows\fonts folder for the font causing the issue and renaming it or removing it to see if that fixes the problem. The issue was resolved by deleting the font and reinstalling the Mattermost app.

発生事象〜提案内容〜解決までが要約されており、要約としては及第点のように個人的には見えます。  
(一点、要約内の提案内容に、フォントファイルの削除指示(`...renaming it or removing it..`)がありますが、[元の投稿](https://forum.mattermost.com/t/font-of-mattermost-app-changed-accidentally/16347/2)には削除に関する言及は無いので、その辺りは少し気になりました。)

> Another user suggests checking the C:\windows\fonts folder for the font causing the issue and renaming it or removing it to see if that fixes the problem.

---

次に、もう少し複雑なやり取りの例として以下の質問の要約を生成させてみました。  

[Upgrading from v5\.1 to v7\.1 \- Troubleshooting / Docs \- Help Wanted \- Mattermost Discussion Forums](https://forum.mattermost.com/t/upgrading-from-v5-1-to-v7-1/13741)

この質問は、Mattermost v5.1からv7.1アップグレードする際に注意すべき点やバックアップの方法について尋ねています。  
質問事項に対して一旦は回答者が注意点について答えていますが、その後、実は元のMattermost Serverのバージョンがv6.3.7だと分かったり、Docker環境のアップデート方法に話が移ったりと、スレッド内で質問の状況が変更されています。

この質問の要約は以下のようになりました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary2-answer.png)

> The thread is about upgrading Mattermost from version 5.1 to version 7.1. The user asks for points to consider beyond the documentation, whether they should upgrade to another older ESR before jumping to v7.1, and the meaning of backing up the database. Another user suggests reading the Important Upgrade Notes for the big upgrade and mentions schema changes starting from version 5.22. They recommend upgrading in multiple steps and provide download links for older versions. They also explain how to back up the database using tar and mysqldump commands. The user acknowledges the advice and mentions failing once but restoring from a virtual machine snapshot. They also mention that their Mattermost version is actually 6.3.7 and not 5.1.1. They tried upgrading using a deprecated script and encountered errors. They ask for guidance on what to do next and provide some proposed steps. The other user suggests getting rid of the old installation and starting fresh with a supported environment. They mention backing up the current setup and deploying a new setup in parallel with the same version. They provide detailed steps for restoring the data and suggest testing everything before running the upgrade to 7.1.2. The user thanks for the advice and mentions using the Mattermost Docker setup from the mattermost/docker repository. They mention the issue they faced with the upgrade-postgres.sh script and ask for further guidance. The other user explains that the mattermost/docker setup does not require the upgrade-postgres.sh script and suggests changing the release tag in the .env file and restarting the setup. The user shares the docker ps output and asks for confirmation on the steps to take. The other user confirms the steps to follow and provides the necessary commands to update the .env file and restart the setup. The user follows the instructions and confirms that it worked. They express their gratitude for the help. The other user responds with a welcome message and expresses happiness that the user's Mattermost is up and running.

質問に対する結論から要約を導く訳ではなく、質問内の全ての回答に対する要約を生成しているようで、スレッド内のコメント数が多くなる程、単純に要約も長くなってしまうようでした。プロンプトとして、スレッドの結論を指定して要約ができたりすると、もう少し結論を元にした端的な要約が生成できるようになるのかなと思いました。

---

最後に、一つ目の質問を日本語に訳したスレッドに対して要約を実行してみましたが、回答は英語で返ってきました。英語のスレッドの要約とは若干表現が違う部分がありますが、概ね同じ内容のことを言っているように見えます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary3-post.png)

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary3-answer.png)

> A Mattermost app user reported that the font in the app was unintentionally changed and asked how to revert it back. The user is using Mattermost App version 5.4.0 on Windows 10 Enterprise. They mentioned that the font change issue only occurred in the Mattermost app and not in other apps or the web browser. Reinstalling the app did not resolve the issue. Another user suggested checking the C:\windows\fonts folder for the font or changing the font's file name if found. However, the original user resolved the issue by deleting the font and reinstalling the Mattermost app.

#### (実験)内部で使われているプロンプトを日本語に変更する

**Mattermost AI Plugin**は、各ユースケースごとに使用されるプロンプトを内部に保持しています。現在、プロンプトは英語のみに対応しているため、日本語のスレッドに対しても英語の要約が生成されてしまうのだと思います。  
試しに、このプロンプトの内容を日本語にして、日本語のスレッドの要約を生成させてみました。

スレッド要約時のプロンプトは、以下のものが使われます。  
https://github.com/mattermost/mattermost-plugin-ai/blob/4b53767d6acadb290be1266f288e9d4b755fc6c4/server/ai/prompts/summarize_thread.tmpl

このプロンプトの内容を以下のように変更しました。

```
{{define "summarize_thread.system"}}
{{template "standard_personality.tmpl" .}}
あなたはスレッドを要約する専門家です。
ユーザーがスレッドを与えると、スレッドの要約を返します。
スレッドについては言及せず、ただ要約のみを提供して下さい。
{{end}}
{{define "summarize_thread.user"}}
スレッドは以下の通りです:

---- Thread Start ----
{{.PromptParameters.Thread}}
---- Thread End ----
{{end}}
```

このテンプレートから使用されている [`standard_personality.tmpl`](https://github.com/mattermost/mattermost-plugin-ai/blob/4b53767d6acadb290be1266f288e9d4b755fc6c4/server/ai/prompts/standard_personality.tmpl)の方も以下のように変更しました。

```
あなたは利用者の手助けをするAIアシスタントです。Mattermostのサーバー名は{{.ServerName}}で、{{.CompanyName}}が所有しています。
利用者の現在の時刻と日付は{{.Time}}です。
利用者の名前は'{{.RequestingUser.Username}}'です。
{{if .RequestingUser.FirstName}}
フルネームは{{.RequestingUser.FirstName}} {{.RequestingUser.LastName}}です。
{{end}}
{{if .RequestingUser.Position}}
彼らの役職は'{{.RequestingUser.Position}}'です。彼らの役職に基づいて、利用者に対する応答をカスタマイズしてください。
{{end}}

{{if and (ne .Channel nil) (ne .Channel.Type "D")}}
あなたがやり取りをしているチャンネルは'{{.Channel.Name}}'で、表示名は'{{.Channel.DisplayName}}'です。チャンネルはチーム'{{.Team.Name}}'に属しており、表示名は'{{.Team.DisplayName}}'です。
{{end}}
```

プロンプトを変更した後にプラグインをビルドし直し、再度インストールした環境で、日本語スレッドの要約を生成させると、以下の出力を得ました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary4-answer.png)

> 以下は、与えられたスレッドの要約です：
> 
> Mattermostアプリのフォントが意図せず変更されてしまったが、元に戻す方法を尋ねる投稿がありました。再現手順や想定される動作についての詳細があります。他のユーザーがフォントの問題についてアドバイスをしましたが、結局はフォントを削除してアプリを再インストールした結果、問題が解決されました。ユーザーは感謝の気持ちを述べています。

英語の時よりだいぶスッキリした要約が生成されていますが、最も重要な部分であるフォントファイルの場所(`C:\windows\fonts`)が回答から抜けてしまっているため、要約としては少し厳しいものがある気がしました。

### スレッドに関する質問 (Answer questions about Threads)

スレッドの要約を生成した後、そのスレッドに関する追加の質問を行うこともできるようです。

一つ前のセクションで書いた日本語スレッド向けのプロンプトを導入した環境で、回答に含まれていなかったフォントファイルの場所について追加で質問してみたところ、以下の回答を得ました。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/mattermost-plugin-ai/ai-thread-summary-response1.png)

> はい、このスレッド内でフォントが存在するフォルダに関する議論があります。
> 
> 投稿者の"test-user"は、問題の解決策として、フォントが存在するフォルダであるC:\windows\fontsを確認するよう提案しています。フォルダの中にこのフォントが存在するか、またはファイルの最終更新日時をチェックすることで、問題のフォントを特定しようとしています。
> 
> 投稿者"kaakaa"は、フォントを削除してMattermostアプリを再インストールした結果、問題が解決されたと報告しています。フォントが存在するフォルダに関しては具体的な議論はありませんが、再インストールによって問題が解決されたことが伺えます。

追加の質問をすることで、詳しい状況を確認することができました(ただ、これはスレッドの内容を知っているからできる質問であって、スレッドを読んでいない状態でこの質問を思いつくかは少し微妙だなと思いました)。

また、元の要約ではメッセージを投稿したユーザー名等は出てきていませんでしたが、追加の質問ではユーザー名が明示されるようになっています。

### OpenAI限定の機能

OpenAIのFunction calling機能を使った機能も開発中のようです。  
現在は、`LookupMattermostUser`と`GetChannelPosts`、また、GitHub Pluginを有効化している場合のみ`GetGithubIssue`というFunctionが使えるようでです。詳細は[server/built_in_tools.go](https://github.com/mattermost/mattermost-plugin-ai/blob/4b53767d6acadb290be1266f288e9d4b755fc6c4/server/built_in_tools.go#L143)の定義を参照ください。

いくつか試してみましたが、思うような回答を得ることができなかったため、ここでは概要の紹介のみにとどめます。

# さいごに

**Mattermost AI Plugin**の機能をいくつか試してみましたが、単純にAIに対して質問をするだけでなく、スレッドの要約や、Function Callingの機能を使ったMattermostの情報との連携など、Mattermost上での生成的AIのユースケースを色々と模索している感じがしました。  
また、現時点では英語のプロンプトのみが用意されているため、まだ日本語でのチャットとは少し相性が悪いかもしれません(単純なAIチャットの機能だけなら大丈夫そうですが)。英語以外の言語への対応については、以下のIssueで方向性について意見を募集している段階のようです。

[Support multiple languages for input messages · Issue \#34 · mattermost/mattermost\-plugin\-ai](https://github.com/mattermost/mattermost-plugin-ai/issues/34)

対応言語に関する部分に限らず、Mattermostの情報をどのように生成的AIサービスへ連携させるかというところは試行錯誤となるであろう部分のため、今後どのような形に落ち着いていくのかは気になるところです。
