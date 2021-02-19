---
title: "Mattermost5.32の新機能"
date: 2021-02-20T06:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

2021/02/16 に Mattermost v5.32.0 がリリースされました。  
また、02/17 に PostgreSQL を使用している際の検索機能に関する修正を行った v5.32.1 がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。（なので、Enterprise 版限定の機能などについてはリリースノート書いてあることの紹介程度となっています）

変更内容の詳細については公式のリリースを確認してください。

- [Mattermost v5\.32 is now available \- Upgrade to the new release today](https://mattermost.com/blog/mattermost-release-v5-32/)
- [Mattermost Changelog — Mattermost 5\.32 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-32-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- `Cloud`: [Mattermost Cloud](https://mattermost.com/pricing-cloud/)
- `E20/E10`: [Enterprise E20/E10](https://mattermost.com/pricing-self-managed/)

見出しの前に何もない場合、Team Edition(OSS 版)でも利用可能な機能です。

## カスタムチャンネルカテゴリ

以前より実験的な機能として公開されていた、チャンネルリストのカスタムカテゴリ作成機能が一般公開されました。このアップデートにより、チャンネルサイドバーで以下のような操作を行えるようになりました。

- サイドバーのチャンネルをグループ化するための任意のカテゴリを作成できます。カテゴリ内のチャンネルのソート順として、手動・アルファベット順・更新履歴順をカテゴリごとに設定できます。また、カテゴリを折り畳んだり、カテゴリ単位でミュートしたりできます。
- 1 クリックで未読チャンネルのみを表示するよう切り替えることができます。その他のオプションとして、未読チャンネルをサイドバー上部にグループ化するよう設定することもできます。
- チャンネルに対する共通の操作（ミュート、既読にする、移動、お気に入り）をサイドバーから行うことができます。

チャンネルサイドバーの改善について、詳しくは下記を参照ください。

[Organize your sidebar with custom, collapsible channel categories](https://mattermost.com/blog/custom-collapsible-channel-categories/)

{{< youtube Jfr1ryP7-og >}}

## (Cloud/E20) Incident Management プラグインの改善

v5.30 から導入された Mattermost 上でインシデント管理ができる Incident Mangement プラグインについて、登録したインシデントのステータスを細かく設定できるようになりました。以前までのバージョンでは`Ongoing`(進行中)か`Ended`(完了)の 2 つのステータスを設定できましたが、ステータスが以下の 4 つに増えました。

- `Reported`: インシデントが発生したが、まだ調査用のリソース等が登録されていない状態
- `Active`: 報告された事象が責任者によりインシデントと認められ、調査が進行している状態
- `Resolved`: インシデントによるリスクは緩和されたが、まだ作業が残っている状態
- `Archived`: インシデント対応が完了し、関連するチャンネルや履歴がアーカイブされた状態。

また、完了済み・進行中の全インシデントを、リストとして一覧表示できるようになりました。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.32/incident-plugin-list.webp)
(画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-32/)より)

## (Cloud/E20) Microsoft Teams ミーティングプラグイン

Microsoft Teams と連携し、1 クリックで Teams ミーティングを作成・参加することができるプラグインが一般公開されました。

## (E10/E20) サブスクリプション更新のオンライン化

E10/E20 を利用している場合、サブスクリプションの更新手続きをオンラインで行えるようになりました。
更新作業はカスタマーポータルで行うことができ、カスタマーポータルでは、過去のサブスクリプション購入履歴も確認できます。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-5.32/renew-subscription.webp)
(画像は[公式ブログ](https://mattermost.com/blog/mattermost-release-v5-32/)より)

更新手続きについて、以下の動画で手順が説明されています。

{{< youtube Sz_1nhVufHY >}}

## その他の変更点

### 破壊的変更

- 既存のチャンネルサイドバー関連の設定は `EnableLegacySidebar` の設定を `true` にしない限り、有効にはならなくなりました。

  - `ExperimentalChannelOrganization`
  - `EnableXToLeaveChannelsFromLHS`
  - `CloseUnusedDirectMessages`
  - `ExperimentalHideTownSquareinLHS`

- Go 言語 API クライアントの以下の API の必須パラメータ `collapsedThread bool` が追加されました。以下の API を利用している場合、修正が必要です。
  - `GetPostThread`
  - `GetPostsForChannel`
  - `GetPostsSince`
  - `GetPostsAfter`
  - `GetPostsBefore`
  - `GetPostsAroundLastUnread`

破壊的変更について詳しくは [Changelog](https://docs.mattermost.com/administration/changelog.html#release-v5-32-feature-release) を参照してください。

## その他のトピック

### Sponsored Rocky Linux Project

CentOS に関する RedHad の方針変更を受けて生まれた Rocky Linux プロジェクトについて、Mattermost がスポンサーとしてコミュニケーション基盤を提供しています。

[Rocky Linux](https://rockylinux.org/)

### What Matterst - A Podcast from Mattermost

Mattermost 主催の Podcast にエピソードが追加されています。最新のエピソードでは、前述の Rocky Linux のコミュニティマネージャーがゲストです。

[What Matters \- A Podcast from Mattermost](https://mattermost.libsyn.com/)

## Mattermost MVP

昨年 12 月に書いていた[Mattermost Integration Advent Calendar](https://qiita.com/advent-calendar/2020/mattermost-integrations)を評価していただき、2 度目の MVP を受賞いたしました。

> Thanks for all community contributions this month and, in particular, our v5.32 Most Valued Professional (MVP), Yusuke Nemoto, who published 25 blog posts as a “Mattermost Integrations Advent Calendar” in December 2020. Thank you for your continued contributions, Yusuke Nemoto!

## おわりに

次の`v5.33`のリリースは 2021/03/16(Tue)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
