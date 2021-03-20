
---
title: "MattermostにRedashのチャートを貼り付ける"
date: 2018-01-03T15:52:35+09:00
draft: false
toc: true
tags: ["redash","Mattermost"]
---

# はじめに

Redashでアクセス数を可視化してるのですが、わざわざRedashを開くのが面倒になってきたのでMattermostからRedashのチャートを取得できる[kaakaa/matter-redash](https://github.com/kaakaa/matter-redash)というものを作りました。

SlackだとRedash連携は色々あるのにMattermostでは見当たらない = 作るしかない。
[re:dashにslack bot機能が追加されたよ！ - Qiita](https://qiita.com/toru-takahashi/items/e86acc292e0b3dc360e8)
[hakobera/redashbot: Slack Bot for re:dash](https://github.com/hakobera/redashbot)

# どんなもの？

MattermostのSlash Commandとして動きます。

![matter-redash.gif](https://qiita-image-store.s3.amazonaws.com/0/9891/2de5c4ed-2a28-d9e2-fe33-b4fac2b585c5.gif)

# どうやって動かす？

[kaakaa/matter-redash: Redash integrations for Mattermost](https://github.com/kaakaa/matter-redash)のREADMEに書いています。

# どういう風に動いてる？

* MattermostのSlash Commandで`matter-redash`へリクエスト飛ばす
* 引数で指定されたチャート画像のスクリーンショットを取得
* Mattermostへアップロード＆チャートを投稿

という流れです。

以下、いくつかつまづいたところ。

## Redash

Redashからチャート画像を取得するために`/embed` APIを使用しています。（RedashのAPI Specはどこかに無いだろうか・・）
[redash/embed.py at 1ef2238d65c0c87ed0bc74141cca44b7766b8453 · getredash/redash](https://github.com/getredash/redash/blob/1ef2238d65c0c87ed0bc74141cca44b7766b8453/redash/handlers/embed.py#L69)

このAPIを叩くとSVG入りのHTMLが取得でき、そのままではMattermostでレンダリングできないので[GoogleChrome/puppeteer](https://github.com/GoogleChrome/puppeteer)を使ってスクリーンショットを撮り、その画像ファイルをMattermostにアップロードしています。

```js:app/mattermost.js#L35
const embedUrl = util.format('%s://%s/embed/query/%d/visualization/%d?api_key=%s', protocol, redashHost, queryId, visualizationId, config.redash.apiKey);
console.log('Redash Embed URL: %s', embedUrl); // eslint-disable-line no-console |
const file = await webshot(embedUrl); |
```


```js:app/webshot.js#L4
const webshot = async (url) => {
    const tmpFile = tempfile('.png'); 
    const browser = await puppeteer.launch(); 
    const page = await browser.newPage();

    await page.goto(url); 
    await page.screenshot({path: tmpFile}); 
    await browser.close(); 

    return tmpFile;
};
```

そこら辺の実装方法については[hakobera/redashbot](https://github.com/hakobera/redashbot)を参考にさせていただきました。

## Mattermost

### Personal Access Token

MattermostにAPI経由で画像ファイルをアップロードするために投稿権限が必要で、今の所[Personal Access Token](https://docs.mattermost.com/developer/personal-access-tokens.html)のみサポートしてます。
なので、MattermostはPersonal Access Tokenが実装された`v4.1.2`以上である必要があります。

### Public Link

Mattermost上でチャート画像を見やすくするため、一旦チャート画像をメッセージの添付ファイルとして投稿し、[ファイルに対する公開リンク]([Sharing Files — Mattermost 4.5 documentation](https://docs.mattermost.com/help/messaging/attaching-files.html#sharing-public-links))を取得し、[Message Attachments](https://docs.mattermost.com/developer/message-attachments.html#images)として投稿するようにしています。

```js:app/mattermost.js#L39
    // Upload webshot image file
    const imageFormData = new FormData();
    imageFormData.append('files', fs.createReadStream(file));
    imageFormData.append('channel_id', req.body.channel_id);

    const uploadFile = await client.uploadFile(imageFormData, imageFormData.getBoundary());

    // Get public link of uploaded file
    const fileId = uploadFile.file_infos[0].id;
    const post = await client.createPost({
        channel_id: req.body.channel_id,
        message: 'This message is posted solely to get the public link and should be deleted immediately.',
        file_ids: [fileId]
    });
    const link = await client.getFilePublicLink(fileId);
    console.log('Mattermost Public Link: %s', link.link); // eslint-disable-line no-console

    // Response to Mattermost
    res.header({'content-type': 'application/json'});
    res.send(commandResponse(url, link.link));

    // Delete unnecessary post.
    client.deletePost(post.id);
```

最後に不要なPostを削除していますが、これはMattermostの仕様上、一度Postに対する添付ファイルとして投稿しないとファイルの公開リンクを取得できないためです。
なので、実行ごとに`(このメッセージは削除されました)`という投稿が発生していまうので少し気持ち悪いです。

### 注意点

Personal Access TokenもPublic Linkも設定で無効にできたりするので、システムコンソールから有効にしておく必要があります。

## おわりに

とりあえず運用してみて、使えそうなら[Share your Mattermost projects | Mattermost](https://www.mattermost.org/share-your-mattermost-projects/)からMattermostチームに報告する所存。

あと、[Mattermostプラグイン](https://docs.mattermost.com/administration/plugins.html)として作り直したいな。

