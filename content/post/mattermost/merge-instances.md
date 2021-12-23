---
title: "Lessons learned from merging Mattermost instances"
emoji: "🎉"
published: true
date: 2021-12-20T12:00:00+09:00
toc: true
tags: ["mattermost", "bulk-export", "bulk-import", "advent-calendar", "tips"]
---

この記事は[FUJITSU Advent Calender 2021](https://qiita.com/advent-calendar/2021/fujitsu)の20日目の記事です。 

今回は、別々に運用されていた2つのMattermostインスタンスを合体させた事例について紹介します。

## 変更履歴
* 2021/12/22: **移行時に実行したコマンド** にて、Bulk Exportコマンドに`--attachments`オプションが付与されていなかったため追加
* 2021/12/22: **移行時に実行したコマンド** について説明を追加

## 背景

社内では、オンプレミスで運用できるチャット基盤としてOSSの[Mattermost](https://mattermost.com/)を利用しているところが数多くあります。もちろんコミュニケーション基盤は統一的なインスタンスで運用できているのが望ましいですが、MattermostはOSSで誰でも使えるということもあり、過去に同時発生的に有志によるMattermostの運用が始まってしまったため、それぞれのインスタンスでの運用が続いているという状況があります。

今回統合を行った2つのMattermostインスタンスは、幾度かの組織変更により両方のインスタンスにアカウントを持つユーザーが増えてきているということもあって、以前からユーザーの利便性のためにも統合したいねという話が出ていました。[今年の4月に大きなグループ再編の動き](https://pr.fujitsu.com/jp/news/2021/01/28-2.pdf)があったことから、この流れに乗ってMattermostインスタンスの統合に踏み切りました。

Mattermostのインスタンスを統合した事例は今まで目にしたことは無いため、公開されている事例としては世界初になるのではないかと思っています(あったら教えて欲しい)。(refs: [Mattermost Forum](https://forum.mattermost.org/t/consolidate-two-instances/5265/4))


### 免責事項
Mattermostは、2つのインスタンスの統合を公式にはサポートしていません。
本記事で述べる内容は非公式の手順であるため、データの損失や欠落を招く可能性があります。その点に留意してください。

最終的に統合作業を完了させていますが、以下のような制限事項がありました。

1. [統合機能（ウェブフックやスラッシュコマンド）](https://docs.mattermost.com/about/integrations.html#integrations-overview)は手動での移行作業が必要
   - API の URL が変わるので、内向きのウェブフックなどは連携アプリ側も書き換える必要がある
1. [ピン留めされた投稿](https://docs.mattermost.com/messaging/pinning-messages.html)のリストが消える
1. [保存済みの投稿](https://docs.mattermost.com/messaging/saving-messages.html)のリストが消える
1. チャンネルやダイレクトメッセージが全部未読になる
1. **プロフィール画像**が移行されない
1. 投稿に付けられた**絵文字リアクション**が消えることがある
1. 投稿に対する URL が変わるので、Mattermost内に書き込まれた**投稿へのリンク**が切れてしまう
1. **チャンネルの作成者と作成日時**の情報が消えてしまう
1. [カスタム絵文字](https://docs.mattermost.com/messaging/using-emoji.html#creating-custom-emojis)の登録者の情報が消えてしまう
1. 移行前と移行後でユーザー名が異なる場合、投稿内の **@メンション** がリンクからテキストになってしまう
1. チャンネルへの参加/脱退やヘッダー変更を伝える**システムメッセージ**がユーザーの投稿になってしまう ([対応中](https://github.com/mattermost/mattermost-server/pull/18992))
1. 投稿を後から編集した場合にメッセージに表示される **(編集済)** のフラグが消えてしまう ([対応中](https://github.com/mattermost/mattermost-server/pull/18992))

これらの点については、事前にユーザーに通知をし、制限事項を認識してもらったうえでインスタンスの統合を実施しました。

### Mattermostの機能について

Mattermostには、2つの異なるインスタンス間の共有や移行に関する以下のような機能があります。

#### [Shared Channel](https://docs.mattermost.com/onboard/shared-channels.html)

複数インスタンス間でチャンネルを共有する機能です。  
今回は、ユーザーの利便性のためにインスタンス自体を統合したいという要求であったため、この方法は採用しませんでした。(また、この機能は有償版限定のためそもそも採用できませんでした)

#### [Migration Guide](https://docs.mattermost.com/onboard/migrating-to-mattermost.html)

Mattermostは、稼働中のインスタンスを別サーバーを移行するための手順を公開しています(Slackなどの他のメッセージングサービスからからMattermostへ移行する手順の紹介などもあります)。  
ただし、この手順によるMattermostサーバーの移行は、データベースごと新しいインスタンスへ移行する方式のため、既にデータの存在するMattermostインスタンスへの移行には適用できませんでした。

#### [Bulk Export](https://docs.mattermost.com/manage/bulk-export-tool.html) / [Bulk Import(Loading)](https://docs.mattermost.com/onboard/bulk-loading-data.html)

**Bulk Export/Bulk Import(Loading)** 機能は、インスタンス上のデータをファイル形式でExportする機能と、Exportされたファイル形式を元にImport処理を行う機能です。  
これらの機能も、Mattermostインスタンス上のデータをアーカイブとして出力することを目的とした **Bulk Export** と、他のメッセージングシステム(Slack, HipChat等)からExportしたデータをMattermostへImportする際のデータ形式を定めた **Bulk Import** の機能という位置づけであり、2つのMattermostインスタンスの統合をサポートしているものではありません。単純に **Bulk Export** したデータを、そのまま既にデータの存在するMattermostインスタンスへ **Bulk Import** すると、様々なデータの不整合が発生する可能性があります。

しかし、その他にインスタンスを統合する有効な手段が無かったため、今回はこの[Bulk Export](https://docs.mattermost.com/manage/bulk-export-tool.html)/[Bulk Import(Loading)](https://docs.mattermost.com/onboard/bulk-loading-data.html)の機能をベースにインスタンス統合の作業を行いました。

本記事では、 **Bulk Export/Bulk Import** 機能を利用したインスタンス統合において発生した問題と、対処方法について紹介していきます。

## Mattermostインスタンス統合作業について

### はじめに

インスタンスの統合を行う際には、以下のことが重要です。

- 移行元と移行先のMattermostのバージョンを合わせること
- インスタンス統合結果が想定通りのものとなっていることを確認できるテスト環境を用意すること

今回は以下のような2つのインスタンスを統合しました。

|                    | 移行元     | 移行先    |
|:-------------------|:----------|:----------|
| Mattermost Version | v5.35.2   | v5.34.0   |
| ユーザー数          | 301       | 685       |
| 投稿数              | 数十万程度 | 100万程度 |

(早速バージョン合ってないですが、移行元からデータをExportする際にバージョン上げないと上手く動かない部分があり...)

### Bulk Export/Bulk Import機能について

Mattermostの [Bulk Export](https://docs.mattermost.com/manage/bulk-export-tool.html) 機能は、Mattermost内の各リソースを以下のようなJSONL形式で出力できる機能です。

```JSONL
{ "type": "version", ... }
{ "type": "team", "team": { "name": "TeamA", ...} }
{ "type": "team", "team": { "name": "TeamB", ...} }
{ "type": "channel", "channel": { "team": "TeamA", "name": "ChannelA1", ...} }
{ "type": "channel", "channel": { "team": "TeamA", "name": "ChannelA2", ...} }
{ "type": "user", "user": { "username": "user001", ...} }
{ "type": "user", "user": { "username": "user002", ...} }
{ "type": "user", "user": { "username": "user003", ...} }
.
.
.
{ "type": "post", { "team": "TeamA", "name": "ChannelA1", "user": "user001", ...} }
{ "type": "post", { "team": "TeamA", "name": "ChannelA1", "user": "user001", ...} }
.
.
.
```

Mattermost内のリソースが1行のJSONデータとして各行に記録されており、Bulk Import実行時も1行ずつ処理されていきます。  
投稿に対する添付ファイルや、カスタム絵文字のファイルの実体などは、このJSONLファイルとは別に出力されており、Bulk Import実行時に併せて読みだされてMattermostへ登録されていきます。

Exportされるデータ種別やその内容については、公式ドキュメントを参照してください。
https://docs.mattermost.com/manage/bulk-export-tool.html

### 移行時に実行したコマンド

統合実行時のコマンドは以下になります。

まず、移行元のサーバーでBulk Exportのコマンドを実行します。  
この時、`--config`オプションの値には起動しているMattermostサーバーの`config.json`を指定します。また、`mattermost-export-YYYY-MM-DD.txt`が出力ファイルとなるのですが、**このファイルの出力先フォルダにMattermostの添付ファイルを格納している`data/`フォルダが存在すると、その内容がすべて0byteのファイルで上書きされてしまいます。Bulk Exportコマンドの出力先には十分注意してください。**

```bash
# Export
$ ./bin/mattermost --config ~/config.json export bulk ~/mattermost-export-YYYY-MM-DD.txt --attachments
```

先にも述べましたが、既にデータの存在するインスタンスに対して、ExportしたJSONLファイルをそのままImportしようすると、データの不整合が発生する可能性があります。そのため、Exportされたデータに対してsedによる置換を実施しています。また、Bulk Import実行時にエラーとならないようファイルの拡張子を`.jsonl`に変更しています。sedによる置換内容については以降のセクションで紹介していきます。

```
# Edit exported data
$ sed -f migration.sed ~/mattermost-export-YYYY-MM-DD.txt > ~/mattermost-import-YYYY-MM-DD.jsonl
```

最後に、Bulk Exportによって出力された`data/`,`exported_emoji/`,`mattermost-import-YYYY-MM-DD.jsonl`を移行先のMattermostサーバーへコピーし、Bulk Importを実行します。この時、`mattermost-import-YYYY-MM-DD.jsonl`と`data/`,`exported_emoji/`フォルダは同じフォルダ内に置いておく必要があります。

```
# Validation
$ ./bin/mattermost --config ~/config.json import bulk ~/mattermost-import-YYYY-MM-DD.jsonl --import-path ~/data --validate

# Import
$ ./bin/mattermost --config ~/config.json import bulk ~/mattermost-import-YYYY-MM-DD.jsonl --import-path ~/data --apply
```

(余談ですが、今回統合を実施した環境では、Mattermostの管理CLIツールが`mattermost`コマンドから`mmctl`コマンドへ移行する過渡期であり、ところどころの調査では`mmctl`コマンドを利用していたりもします。Mattermost v6からは管理CLIコマンドの大部分が`mmctl`コマンドに集約されているため、Mattermost v6移行の環境で統合を実施する場合は実行コマンドが変わるかもしれません。)

## 統合実施時に注意が必要なデータ

### [Team object](https://docs.mattermost.com/manage/bulk-export-tool.html#team-object)

[Example Team object](https://docs.mattermost.com/onboard/bulk-loading-data.html#example-team-object)
```JSON
{
  "type": "team",
  "team": {
    "name": "team-name",
    "display_name": "Team Display Name",
    "type": "O",
    "description": "The Team Description",
    "allow_open_invite": true
  }
}
```

#### `name`

移行元/移行先で **チーム名** (`name`)が同じだった場合、移行元と移行先のチームがマージされる (移行先のチームに移行元の同名チーム内のチャンネルが作成される) ような動作になります。

今回の移行では、移行元/移行先で異なるチーム名で運用していましたが、**チームを統一するために移行元のチーム名を移行先のメインで使われているチーム名に変更**しました。

### [Channel object](https://docs.mattermost.com/manage/bulk-export-tool.html#channel-object)

[Example Channel object](https://docs.mattermost.com/onboard/bulk-loading-data.html#example-channel-object)
```json
{
  "type": "channel",
  "channel": {
    "team": "team-name",
    "name": "channel-name",
    "display_name": "Channel Name",
    "type": "O",
    "header": "The Channel Header",
    "purpose": "The Channel Purpose",
  }
}
```

#### `team`/`name`

移行元/移行先の両方に同じ　**チーム名/チャンネル名**　を持つチャンネルが存在する場合、2つのチャンネルの投稿が1つのチャンネルにマージされます（すごい）。
今回のインスタンス統合では、移行元/移行先で同名となるチャンネルについて事前に調査を行い、マージを希望しないチャンネルについては、移行元のインスタンスのチャンネル名に **特定のprefixを付ける** ことでマージが実施されないように対処しました。一部、マージを希望するチャンネルについては同名チャンネルのまま統合作業を行いましたが、**特に問題なく移行元/移行先のチャンネルのマージが行われました**(すごい)。

ただし、Mattermostでチームを作成した時に自動で作成される `town-square`、`off-topic` チャンネルには注意が必要です。  
この2つのチャンネルは、ユーザーがチームに新しく参加した時に自動で参加するチャンネルであり、事前にチャンネル名を変更してしまうと、新たにユーザーをチームに参加させたときにエラーが発生してしまいます(`off-topic`はUIから名前を変更できますが、`town-square`はシステム的な固有値だったはず)。この2つのチャンネルのチャンネル名を変更する場合はExportされたデータ上でチャンネル名を編集した方が良いです。

---

また、Mattermost内に日本語を含むチャンネル表示名がある場合も注意が必要です。

Mattermostではチャンネルの名前を表す属性として **チャンネル名**  (`name`) と **チャンネル表示名**  (`display_name`) の2つがあります。普段Mattermost画面上で見ている名前は **チャンネル表示名**  (`display_name`) で、URLに使われる文字列が **チャンネル名** (`name`) です。  
Mattermostで画面上からチャンネルを作成する際に入力する名前は **チャンネル表示名** (`display_name`)になりますが、表示名を入力すると自動で **チャンネル名** (`name`)の方も同じ文字列で設定されます（チャンネル名のみを変更することも可能）。英語のみで構成されている **チャンネル表示名** であれば同じ値が設定されますが、もし日本語などのマルチバイト文字が含まれる場合、Mattermostはそれらのマルチバイト文字を無視して英語部分だけで **チャンネル名** を設定します。そのため「○○PJ(プロジェクトの意味)」のような、日本語を含む短い **チャンネル表示名** を設定している場合、意図せず **チャンネル名**  (`name`) が重複することが有ります。

このため、日本語を含むチャンネル名が多く存在する場合、チャンネル名のConflictが発生しやすくなっており、事前の調査がとても重要になります。

---

また、[アーカイブ済みのチャンネル](https://docs.mattermost.com/messaging/managing-channels.html#archiving-a-channel)は **Bulk Exportの対象になりません**。
今回はアーカイブ済みのチャンネルも移行するために、Bulk Export実行前に移行元インスタンスのアーカイブされているチャンネルを復元し、移行先へImportした後に再度アーカイブを実行しました。

また、移行元のチャンネル名と同名のチャンネルが移行先にあり、かつ、移行先のチャンネルがアーカイブ済みだった場合、**Import処理が失敗します**。この場合は、当該チャンネルのデータを削除するか、アーカイブを解除する対処が必要です。

### [User object](https://docs.mattermost.com/manage/bulk-export-tool.html#user-object)

[Example User object](https://docs.mattermost.com/onboard/bulk-loading-data.html#example-user-object)
```json
{
  "type": "user",
  "user": {
    "profile_image": "avatar.png",
    "username": "username",
    "email": "email@example.com",
    "auth_service": "",
    "auth_data": "",
    "password": "passw0rd",
    "nickname": "bobuser",
    "first_name": "Bob",
    "last_name": "User",
    "position": "Senior Developer",
    "roles": "system_user",
    "locale": "pt_BR",
    "teams": [
      {
        "name": "team-name",
        "theme": "{\"awayIndicator\":\"#DBBD4E\",...}",
        "roles": "team_user team_admin",
        "channels": [
          {
            "name": "channel-name",
            "roles": "channel_user",
            "notify_props": {
              "desktop": "default",
              "mark_unread": "all"
            }
          }
        ]
      }
    ]
  }
}
```

#### `username`/`email`

移行元/移行先の両方のインスタンスにアカウントを持っているユーザーがいる場合、アカウントのユーザー名/メールアドレスに応じて対処が必要になります。

* 移行元/移行先のアカウントが...
  1. **同じ**ユーザー名 / **同じ**メールアドレス: 対処不要
  2. **同じ**ユーザー名 / **異なる**メールアドレス: 移行元のメールアドレスに**書き換えられる**
  3. **異なる**ユーザー名 / **同じ**メールアドレス: Import処理で**エラーが発生**するため対処が必要
  4. **異なる**ユーザー名 / **異なる**メールアドレス: 移行元のユーザー情報のユーザーが作成される

`2.`の事象については、もし、別々の人が移行元/移行先で同じユーザー名を使っていた場合、アカウントを乗っ取られたような動作になるので対処が必要になります。
`4.`の事象については、本来は正常動作ですが、同じユーザーがが移行元/移行先で**異なるユーザー名 / 異なるメールアドレス**のユーザーを持っていた場合、そのユーザーに対して2つのアカウントが生成されることになります。本来あまり起こらない事象だとは思いますが、過去に全社的にメールアドレスのドメイン名が微妙に変更になったことが有り(以前のメールアドレスも利用可能)、このような事象が発生しやすくなっていました。

これらの点についても、事前にユーザー名/メールアドレスの情報を突合せ、問題が発生する可能性のあるアカウントについては事前にユーザー名/メールアドレスの変更をお願いしたり、移行先のデータが書き換えられないようExportデータの内容を変更するなどして対処しました。

#### `auth_service`/`auth_data`

移行元/移行先で異なる認証サービスを利用していた場合、Export後のJSONLファイルに対する対処が必要です。

今回の統合においては、**移行元のMattermostではGitLab認証**を利用しており、**移行先ではMattermost自身のパスワード認証**を利用していました。ただ、移行元で利用していたGitLabは、[富士通研究所としてGitHub Enterpriseが導入された](https://dev.classmethod.jp/articles/github-constellation-conf-fujitsu/)こともあってMattermost認証のためのアカウント管理機能ぐらいしか使われていなかったため、統合に併せてアカウント管理も移行先のMattermostに寄せることにしました。  
認証関連の情報は、 User object が持つ `auth_service`/`auth_data` 要素として格納されているため、Import実施前にこれらの要素を`{..., "auth_service":null, ...}`のみ(`auth_data`要素は削除)に編集しました。これにより、Import実施時にシステム側でパスワードが生成・設定されるようになります。システム側で設定されたパスワードを知る術がないので、移行先のインスタンスでパスワードリセットの操作を行う必要がありました。

#### `roles`

ユーザーの権限に関するデータである`roles`は **非常に注意が必要です**。

移行元/移行先の両方に同じユーザー名のアカウントが存在した場合、そのユーザーの権限はExportデータ内の`roles`の値で上書きされてしまいます。もし、移行先のシステム管理者ユーザーが移行元システムの一般ユーザー権限で上書きされ、**移行先にシステム管理者がいなくなってしまった場合**、そのインスタンスの運用を続けることが困難になってしまいます。統合にあたって事前にテストするなどして、このような状況が発生しないよう十分に注意してください。

#### `locale`/`theme`/`notify_props`/...

その他、ユーザーのアカウントに関する設定（表示言語、テーマ、通知設定など）も、移行元/移行先で同じユーザー名のユーザーが存在した場合、移行元のデータで置き換えられてしまいます。

### [Post object](https://docs.mattermost.com/manage/bulk-export-tool.html#post-object)

[Example Post object](https://docs.mattermost.com/onboard/bulk-loading-data.html#example-post-object)
```json
{
  "type": "post",
  "post": {
    "team": "team-name",
    "channel": "channel-name",
    "user": "username",
    "message": "The post message",
    "props": {
      "attachments": [{
        "pretext": "This is the attachment pretext.",
        "text": "This is the attachment text."
      }]
    },
    "create_at": 140012340013,
    "flagged_by": [
      "username1",
      "username2",
      "username3"
    ],
    "replies": [{
      "user": "username4",
      "message": "The reply message",
      "create_at": 140012352049,
      "attachments": [{
          "path": "/some/valid/file/path/1"
      }],
    }, {
      "user": "username5",
      "message": "Other reply message",
      "create_at": 140012353057,
    }],
    "reactions": [{
      "user": "username6",
      "emoji_name": "+1",
      "create_at": 140012356032,
    }, {
      "user": "username7",
      "emoji_name": "heart",
      "create_at": 140012359034,
    }],
    "attachments": [{
      "path": "/some/valid/file/path/1"
    }, {
      "path": "/some/valid/file/path/2"
    }]
  }
}
```

#### Attachments

投稿に添付されたファイルの拡張子と実際に添付されたファイルの形式が異なる場合、**Import処理でエラー**が発生します。  
Import処理では、投稿に添付された画像ファイルなどに対してサムネイルを作成する処理が自動で実行されます。この時、Mattermostはファイル名の拡張子を元にサムネイル作成処理を実行しようとしますが、ファイル名の拡張子と実際のファイル形式が異なる場合、エラーが発生します。

今回の統合作業でも、移行元のインスタンスに、`.gif`という拡張子を持つJPEGファイルなどが存在したため、以下のようなエラーが発生しました。

```jsx
GetInfoForBytes: Could not decode gif., gif: can't recognize format "\xff\xd8\xff\xe0\x00\x10"
Error occurred on data file line 87634
{"level":"info","ts":1624515819.9272816,"caller":"app/server.go:841","msg":"Stopping Server..."}
{"level":"info","ts":1624515819.9273949,"caller":"app/web_hub.go:103","msg":"stopping websocket hub connections"}
{"level":"info","ts":1624515819.9294012,"caller":"app/server.go:914","msg":"Server stopped"}
Error: GetInfoForBytes: Could not decode gif., gif: can't recognize format "\xff\xd8\xff\xe0\x00\x10"
```

### [Emoji object](https://docs.mattermost.com/manage/bulk-export-tool.html#emoji-object)

[Example Emoji object](https://docs.mattermost.com/onboard/bulk-loading-data.html#example-emoji-object)

```json
{
"type": "emoji",
"emoji": {
  "name": "custom-emoji-troll",
  "image": "bulkdata/emoji/trollolol.png"
  }
}
```

#### `name`

移行元/移行先で同名のカスタム絵文字が登録されている場合、**移行元の画像でカスタム絵文字が上書きされて** しまいます。
今回の移行では、同名のカスタム絵文字については移行元のデータを削除することで対応しました。もし、Exportデータ内でカスタム絵文字の名前を変更するという対処を行う場合、投稿のデータにもリアクションとしてカスタム絵文字名が使用されている場合があるため、併せてそちらも修正する必要があると思います。

## OSSへの貢献

Mattermost Bulk Export/Importツールについてはあまり利用例がないのか、移行にあたっていくつか問題が発生しました。これらの問題について、MattermostへIssueによる報告や、PRによる改善要望を送っています。

### 解決済み
- [PR] [Bulk export fails due to missing user · Issue \#11659 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/11659)
  - 投稿にリアクションを付けていたユーザーをインスタンスから削除した場合、Bulk Export実行時にエラーとなるバグの報告と修正
  - 2年以上前のIssueですが、この頃から少しずつ移行に向けて進捗していました
- [PR] [Fix bulk export file when exporting reply threads by kaakaa · Pull Request \#17849 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/pull/17849)
  - メッセージスレッド内に多くのファイルが添付されている投稿をエクスポートする場合に、スレッド内の全添付ファイルを各投稿の情報として保持してしまうという問題の修正です。
- [PR] [Add bulk export options by kaakaa · Pull Request \#4688 · mattermost/docs](https://github.com/mattermost/docs/pull/4688)
  - 公式ドキュメントにて、Bulk Exportコマンドのオプションの説明が不足していたため、説明を追加。

### 未解決

- [Issue] [Bulk import makes the system message the user message · Issue \#18640 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/18640)
- [PR] [Added support for exporting and importing the type and edit\_at of a post by kaakaa · Pull Request \#18992 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/pull/18992)
  - Bulk Export/Importによるデータ移行を実施した際、システムメッセージが移行先でユーザーメッセージとして表示されてしまう問題
- [Issue] [Bulk export writes 0byte data to existing files · Issue \#18699 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/18699)
  - Bulk Exportコマンドを `--attachments` オプション（添付ファイルをエクスポートするオプション）付きで実行した際、コマンド実行ディレクトリにMattermostのデータファイル格納フォルダ(`data/`)が存在する場合にデータファイルを0バイトで上書きしてしまう問題
- [Issue] [Doing bulk export/import drops saved post data · Issue \#18641 · mattermost/mattermost\-server](https://github.com/mattermost/mattermost-server/issues/18641)
  - Bulk Export実行時、保存された投稿の情報がExport対象にならない問題
- [Issue] [There is no way to list archived private channels if the user isn't the member of its channel · Issue \#401 · mattermost/mmctl](https://github.com/mattermost/mmctl/issues/401)
  - Mattermostの新しいCLIツール `mmctl` では、privateかつarchivedなチャンネルの一覧を取得できない問題

## さいごに

今回は、Mattermostの[Bulk Export](https://docs.mattermost.com/manage/bulk-export-tool.html) / [Bulk Import(Loading)](https://docs.mattermost.com/onboard/bulk-loading-data.html)の機能を使い、2つのMattermostインスタンスの統合を行いました。そして、インスタンス統合時に発生するデータ不整合と、その不整合にどのように対応したかについて紹介しました。  
2つのMattermostを完全な形で移行することは出来ませんでしたが、制限事項の部分以外は移行後も大きな問題が起こることも無く、現時点で数か月運用が出来ています。

今回、インスタンス統合実施時に発生する問題に対処しつつ作業を完遂出来たのは、MattermostがOSSであったという点が大きいと思っています。問題が発生しても、データとコードを読み解くことで問題の原因を特定でき、原因に対してソフトウェアの方を修正するという対処も取りながら作業を進めることが出来ました。Mattermostは、**自分たちのコミュニケーションのデータを自分たちの自由にできないフラストレーション** から生まれたプロダクトであり、今回はその恩恵を十分に受けることができたと思います。(refs: [The Mattermost Origin Story \- The Craft of Open Source \- Flagsmith](https://flagsmith.com/podcast/ian-tien-mattermost/))

## 謝辞

今回、記事を公開しているのは私ですが、実際のインスタンス統合作業にはあまり関れておらず、普段からMattermostに関する記事を書いているということで、記事公開役を拝命しております。

実際のインスタンス統合作業は、[@zenjiro0123](https://twitter.com/zenjiro0123)さんが 移行元のMattermostインスタンスの管理者 兼 今回の統合のメインの実施者として進めてくれました。また、[@zenjiro0123](https://twitter.com/zenjiro0123)さんが途中から育休に入るということで、[岩田 聡](https://researchmap.jp/iwata-satoshi)さん(任天堂じゃない方)も、様々な面でサポートに回ってもらいました。また、統合作業のテストについては、移行先のMattermostインスタンスの管理者である[@lastis](https://qiita.com/lastis)さんに協力いただき、最終的な移行作業も[@lastis](https://qiita.com/lastis)さんに実行してもらいました。

皆様、お疲れさまでした。
