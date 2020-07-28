---
title: "*.drawio.png と *.drawio.svg とはなんなのか"
date: 2020-07-28T23:16:05+09:00
draft: false
toc: true
tags: ["draw.io", "vs-code", "png"]
---

## VS Code Draw.io Integrationの画像エクスポート機能

数ヶ月前に、VS CodeのDraw.io Integrationが話題になりました。

[VSCodeでDraw\.ioが使えるようになったらしい！ \- Qiita](https://qiita.com/riku-shiru/items/5ab7c5aecdfea323ec4e)

ただ、上記の記事にもあるようにエクスポート機能は`v0.4.0`から無効化されており、エクスポート機能を利用するにはオンライン版のDraw.ioに繋がないといけないという本末転倒な状況になっているようです。

せっかくエディタ内でドキュメントと作図が完結しそうという素敵な体験にあと一歩というところなので、`*.drawio`ファイルを一括で画像ファイルに変換するスタンドアロンのツールを作ってみたり（[kaakaa/dio\-exporter](https://github.com/kaakaa/dio-exporter)）もしましたが、もっと良い方法があったようです。

## `*.drawio.png` と `*.drawio.svg`

[Visual Studio Code \- \*\.drawio\.svg や \*\.drawio\.png の衝撃 \- anfangd's blog](https://blog.anfangd.me/entry/2020/07/08/220628)

こちらの記事で知ったのですが、ファイルを`*.drawio.svg`や`*.drawio.png`という拡張子で保存すると、Draw.ioで開くことができ、かつSVG/PNGファイルとして扱うことができるようです。

公式リポジトリの方にも記述がありました。

https://github.com/hediet/vscode-drawio#editing-drawiosvgdrawiopng-files
> You can directly edit and save .drawio.svg and .drawio.png files. These files are perfectly valid svg/png-images that contain an embedded Draw.io diagram. Whenever you edit such a file, the svg/png part of that file is kept up to date.

なぜ、そんなことができるのかを調べてみました。

## なぜ画像ファイルとしてもDraw.ioとしても開けるのか

### `*.drawio.png` or `*.dio.png`

Draw.ioには、`extractGraphModelFromPng`というメソッドがあり、このメソッドにPNGファイルからDraw.ioのグラフモデルを取得する処理が書かれているようです。

https://github.com/jgraph/drawio/blob/a579fe9c094510093db631283166f35588848113/src/main/webapp/js/diagramly/Editor.js#L1547

```js
  Editor.extractGraphModelFromPng = function(data)
  {
    var result = null;

        ...  (snip) ...
      
      EditorUi.parsePng(binary, mxUtils.bind(this, function(pos, type, length)
      {
        var value = binary.substring(pos + 8, pos + 8 + length);
        
        if (type == 'zTXt')
        {
          var idx = value.indexOf(String.fromCharCode(0));
          
          if (value.substring(0, idx) == 'mxGraphModel')
          {
            // Workaround for Java URL Encoder using + for spaces, which isn't compatible with JS
            var xmlData = pako.inflateRaw(value.substring(idx + 2),
              {to: 'string'}).replace(/\+/g,' ');
            
            if (xmlData != null && xmlData.length > 0)
            {
              result = xmlData;
            }
          }
        }
        // Uncompressed section is normally not used
        else if (type == 'tEXt')
        {
          var vals = value.split(String.fromCharCode(0));
          
          if (vals.length > 1 && (vals[0] == 'mxGraphModel' ||
            vals[0] == 'mxfile'))
          {
            result = vals[1];
          }
        }
        
        if (result != null || type == 'IDAT')
        {
          // Stops processing the file as our text chunks
          // are always placed before the data section
          return true;
        }
      }));

        ...  (snip) ...
```

PNGファイルの`zTXt`チャンク、もしくは`tEXt`チャンクからグラフモデルを読み出しているようです。

実際に、適当な`*.dio.png`ファイルを作って中身をみてみます。 以下のコードは [PNGを読む \- Qiita](https://qiita.com/kouheiszk/items/17485ccb902e8190923b) を参考にしました。

```go
package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "io"
    "net/url"
    "os"
)

func parse(r io.Reader) (err error) {
    buffer := new(bytes.Buffer)
    _, err = buffer.ReadFrom(r)
    if err != nil {
        return err
    }
    if string(buffer.Next(8)) != "\x89PNG\r\n\x1a\n" {
        return fmt.Errorf("not a PNG")
    }

    data := make([]byte, 0, 32)
    loop := true
    for loop {
        length := int(binary.BigEndian.Uint32(buffer.Next(4)))
        chunkType := string(buffer.Next(4))

        switch chunkType {
        case "zTXt":
            fmt.Println("chunk: zTXt")
            data = append(data, buffer.Next(length)...)
            _ = buffer.Next(4) // CRC
        case "tEXt":
            fmt.Println("chunk: tEXt")
            data = append(data, buffer.Next(length)...)
            _ = buffer.Next(4) // CRC
        case "IEND":
            fmt.Println("chunk: IEND")
            loop = false
        default:
            fmt.Println("chunk:", chunkType)
            _ = buffer.Next(length) // chunk data
            _ = buffer.Next(4) // CRC
        }
    }
    d, _ := url.QueryUnescape(string(data))
    fmt.Println(d)

    return nil
}

func main() {
    imageFile := "test.dio.png"
    file, err := os.Open(imageFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    err = parse(file)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Complete")
}
```

実行結果は下記になります。

```
$ go run main.go
chunk: IHDR
chunk: tEXt
chunk: IDAT
chunk: IEND
mxfile<mxfile host="8ce946a9-c68b-479a-941a-3275fc70c066" modified="2020-07-28T13:37:16.115Z" agent="5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Code/1.47.0 Chrome/78.0.3904.130 Electron/7.3.2 Safari/537.36" version="13.1.3" etag="9Ipt2ayCl9tVzXnOHPW3" pages="2"><diagram id="TjNWy2FeEwpBci7nIAZN" name="ページ1"><mxGraphModel dx="630" dy="380" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0"><root><mxCell id="0"/><mxCell id="1" parent="0"/><mxCell id="2" value="TEST" style="rounded=0;whiteSpace=wrap;html=1;" vertex="1" parent="1"><mxGeometry x="260" y="170" width="120" height="60" as="geometry"/></mxCell></root></mxGraphModel></diagram><diagram id="R911WMiv1Lcx-Wbf2GqR" name="ページ2"><mxGraphModel dx="630" dy="380" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0"><root><mxCell id="fRIcQMW-JuJCGEe4Us9i-0"/><mxCell id="fRIcQMW-JuJCGEe4Us9i-1" parent="fRIcQMW-JuJCGEe4Us9i-0"/><mxCell id="fRIcQMW-JuJCGEe4Us9i-2" value="HOGE" style="ellipse;whiteSpace=wrap;html=1;" vertex="1" parent="fRIcQMW-JuJCGEe4Us9i-1"><mxGeometry x="260" y="160" width="120" height="80" as="geometry"/></mxCell></root></mxGraphModel></diagram></mxfile>
Complete
```

実行結果を見ると、Draw.ioのグラフモデルは `zTXt`チャンク ではなく `tEXt`チャンク として保存されているようです。
`tEXt`チャンクの中身は、XML要素ではない`mxfile`というワードで始まっていることと、パーセントエンコードされていること以外は、Draw.ioのデータ保存形式と同じように見えます。実際、先頭の`mxfile`を外したXMLをファイルに保存し、Draw.io Integrationで開いてみたところ問題なく開けました。

---

Draw.ioのコード内では、`tEXt`チャンクは基本的に使われないだろうというコメントが書かれています。

```
// Uncompressed section is normally not used
else if (type == 'tEXt')
```

しかし、VS Code Draw.io Integrationで`*.drawio.png`ファイルを作成した場合は、その`tEXt`チャンクにデータが保存されています。
個人的な予想ですが、エディタでの操作だと保存回数が多くなるため、保存のたびに圧縮をかける必要のある`zTXt`チャンクでは保存操作に対するレスポンスが悪くなる恐れがあるため、`tEXt`チャンクを採用しているのではないかと思っています。

### `*.drawio.svg` or `*.dio.svg`

SVGファイルの中身はXMLファイルなので、テキストファイルとして開けば `<svg>` タグの `content`属性の値として、HTMLエンコードされたDraw.ioのグラフモデルが格納されていることがわかります。

```
<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  version="1.1"
  width="121px"
  height="81px"
  viewBox="-0.5 -0.5 121 81"
  content="&lt;mxfile host=&quot;40cbe15d-9605-4d22-a796-b079811d4635&quot; modified=&quot;2020-07-28T13:50:29.983Z&quot; agent=&quot;5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Code/1.47.0 Chrome/78.0.3904.130 Electron/7.3.2 Safari/537.36&quot; etag=&quot;dN67jIN7lHvydTadAscc&quot; version=&quot;13.1.3&quot; pages=&quot;2&quot;&gt;&lt;diagram id=&quot;6hGFLwfOUW9BJ-s0fimq&quot; name=&quot;Page-1&quot;&gt;jZJNb8MgDIZ/DfcQtq69Luu2y06ptDMKXkAicUTJkuzXjwyTD1WVdsJ+bIP9YiaKZnxzstMfqMCyPFMjEy8szznPsnDMZIpEnB4iqJ1RlLSC0vwAQaqre6Pgukv0iNabbg8rbFuo/I5J53DYp32h3b/ayRpuQFlJe0s/jfI60mP+tPJ3MLVOL/PDKUYamZJpkquWCocNEmcmCofoo9WMBdhZvKRLrHu9E10ac9D6/xTkseBb2p5mu5zLC/XmpzSww75VMNdkTDwP2ngoO1nN0SF8cWDaNzZ4PJh0JTgP4922+DJs2BLABrybQgoViEfShxaEH8kfVrl50lBvpD4Qk/TD9XL1KkIwSIfkrnr/xTZbK86/&lt;/diagram&gt;&lt;diagram id=&quot;KO23mSEetFSPKnzUMudc&quot; name=&quot;ページ2&quot;&gt;rZNNT4QwEIZ/DcdNKNX9OK4sLhdjIgejt0pnaZNCSbcrsL/eYocv92BMPNF5Zvq2fWcIaFy2R8Nq8aQ5qCAKeRvQQxBFhISh+/Sk84Tu7jwojORYNIFMXgEh7isuksN5UWi1VlbWS5jrqoLcLhgzRjfLspNWy1NrVsANyHKmbumr5FZ4uo02E09BFmI4max3PlOyoRhfchaM62aGaBLQ2Ght/apsY1C9eYMvody/vL1/HNr0tE6S+CpWe7nyYo9/2TI+wUBl/1eaeulPpi7oV/p8TPC9thtMdEquXy54aIS0kNUs7zONGxnHhC2Vi4hbohwYC+0P/3+5PBkddaMIugRrOrcPVeg9NgGnkGwwbqaekqFRYtbPLTKGY1SM0pNVboFuDeHU1O/c7NegyRc=&lt;/diagram&gt;&lt;/mxfile&gt;"
>
    <defs/>
    <g>
        <ellipse cx="60" cy="40" rx="60" ry="40" fill="#ffffff" stroke="#000000" pointer-events="all"/>
        <g transform="translate(-0.5 -0.5)">
            <switch>
      ... (snip) ...

```

## 懸念点

`*.drawio.png`、`*.drawio.svg`形式で保存することで、画像ファイルとしても扱えますし、Draw.ioで開けば追加で編集も加えられるとても便利なファイルを作成することができます。しかし、いくつか懸念すべき点があります。

### ファイルサイズが大きくなる

`*.drawio.png`や`*.drawio.svg`は、画像ファイルとしてのデータとDraw.ioとしてのデータの両方を持つことになるため、どうしてもファイルサイズは大きくなってしまいます。また、VS Code Draw.io Integrationで作成すると、Draw.ioのデータは圧縮されずに `tEXt` チャンクに格納されるため、その点でもファイルサイズを低減できていません。
Webページに使用するなど、容量にシビアな環境では画像ファイルとしてエクスポートし直す必要がありそうです。

### 複数ページを扱うことはできない

Draw.ioでは一つのファイルに複数ページ(タブ)を作成することができますが、画像として参照できるのは１ページのみです。（おそらく最初のページ）  複数ページの存在するDraw.ioファイルを画像表示するには、各ページを別ファイルに分割する必要があります。
（[kaakaa/dio\-exporter](https://github.com/kaakaa/dio-exporter)は複数ページのDraw.ioファイルから複数の画像ファイルを出力できるようになっているはずです）

## おわりに

`*.drawio.png`や`*.drawio.svg`というファイル形式にはいくつか懸念事項はありますが、作っているツールの概要をサクッと書きたいときなどは、とても役に立つと思います。

また、どうやら本体のDraw.io Integrationの方でも、PNGエクスポートの機能は追加される予定のようです。

* [Support PNG Export · Issue \#56 · hediet/vscode\-drawio](https://github.com/hediet/vscode-drawio/issues/56)
> As you can see, this feature has been released to insiders:

執筆時点(2020/07/28)での最新版である `v0.7.2` では使えませんが、もしかしたらPNGエクスポートの機能は復活するのかもしれません。