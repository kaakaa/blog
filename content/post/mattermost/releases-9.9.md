---
title: "Mattermost 9.9の新機能"
emoji: "🎉"
published: true
date: 2024-06-14T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/06/14 に Mattermost のアップデートとなる `v9.9.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [Mattermost changelog \- Mattermost documentation](https://docs.mattermost.com/deploy/mattermost-changelog.html#release-v9-9-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

{{< youtube W4e1ZyFI5IA >}}

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## メイン画面のデザイン変更

Mattermostのメイン画面のUIが少し変わりました。  

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.9/channels-core-ui.png)

パッと見た感じ、以下のような点が変更されているようです。

* 画面左上にMattermostアイコンと(Freeプランを利用中の場合のみ)利用中のプランが表示されるようになった
* 中央のチャット画面の外枠が丸みを帯びたデザインとなった
* 右端のAppBarの色が変更された


また、色味についても目の疲労を軽減するよう変更されたようです(微妙な違いですが)。  
特に、デフォルトテーマの中のダーク系のテーマで変化が顕著だそうです。以下は `Onyx` テーマについて、前バージョンと並べて比較したものです。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.9/channels-theme-color.png)


## 内向きのウェブフックでメッセージ優先度付きの投稿が可能に

Mattermost外のシステムからMattermostに対してメッセージを投稿する内向きのウェブフック機能で、メッセージ優先度付きの投稿を作成することができるようになりました。

メッセージ優先度は、投稿の重要度を目立たせるためにラベルを付けて投稿することができる機能です。全員に見て欲しい重要なメッセージを投稿する際などに使用します。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.9/channels-message-priority.png)

内向きのウェブフックでメッセージ優先度付きの投稿をする場合、以下のように `{"priority": {"priority":"urgent"}}` を含むリクエストを送信します。  

```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"text": "post with priority", "priority":{"priority":"urgent"}}' \
  http://192.168.11.99:8065/hooks/9z63d3i7838y3eqs1y9f3nej8c 
```

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.9/channels-post-with-priority.png)

`priority` として指定できるのは `urgent`, `important`, `standard`です。


内向きのウェブフックについて、詳しくは以下のドキュメントを参照してください。  
[Incoming webhooks](https://developers.mattermost.com/integrate/webhooks/incoming/)

また、外向きのウェブフックに対するレスポンスとして作成できる投稿にも、メッセージ優先度を設定できるようになりました。（こちらの説明は割愛します）  
[Outgoing webhooks](https://developers.mattermost.com/integrate/webhooks/outgoing/)

## その他のトピック

### Mattermost Copilot

これまでにも何度か紹介しているMattermostのAI系機能であるMattermost AI Copilotプラグインについて、新たな紹介動画が公開されています。  
JIRAやGitHubのリンクをMattermostに投稿することで、チケットの要約をMattermost上に投稿することなどもできるようです。 (動画 0:57~)

{{< youtube E3EGLxgNxNA >}}

また、以下のサイトからインタラクティブなデモを体験することもできるようになっています。  
[AI innovation meets unparalleled data control](https://mattermost.com/copilot/)

## おわりに
次の`v9.10`のリリースは 2024/7/16(Tue)を予定しています。  
