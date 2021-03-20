
---
title: "JavaでHTTPSのサイトに認証付きプロキシ経由でアクセスしようとすると407エラー"
date: 2016-12-14T12:36:42+09:00
draft: false
toc: true
tags: ["Java","proxy"]
---

# はじめに

Java 8 Update 111から`java.net`パッケージのHTTPS接続の際のトンネリングにBasic認証を使用できない設定がデフォルトになったようです。
回避するには`jdk.http.auth.tunneling.disabledSchemes`に空文字列を設定する必要があるそうです。

# 事象

## 背景
Javaの`java.net.HttpURLConnection`を利用して、認証付きプロキシ内の環境からhttpsのサイトにアクセスしようとした時に、どんな方法でプロキシの認証情報を付与しても`407 Proxy Authentication Required`エラーを解消できなかった。

```java:Sample.java
import java.util.*;
import java.net.*;

public class Sample {
        private static String PROXY_HOST = "example.proxy.com";
        private static String PROXY_PORT = "8080";
        private static String PROXY_USER = "user";
        private static String PROXY_PASSWORD = "password";

        public static void main(String[] args) throws Exception {
                String url = "https://news.google.com";
                System.setProperty("proxySet", "true");
                System.setProperty("proxyHost", PROXY_HOST);
                System.setProperty("proxyPort", PROXY_PORT);

                Authenticator.setDefault(new Authenticator() {
                        @Override
                        protected PasswordAuthentication getPasswordAuthentication() {
                                return new PasswordAuthentication(PROXY_USER, PROXY_PASSWORD.toCharArray());
                        }
                });

                HttpURLConnection urlconn = (HttpURLConnection) new URL(url).openConnection();
                urlconn.setRequestMethod("GET");
                urlconn.connect();
        }
}
```

## 現象

上記のJavaコードを実行すると下記のエラーが発生する。

```text
$ java Sample
Exception in thread "main" java.io.IOException: Unable to tunnel through proxy. Proxy returns "HTTP/1.1 407 Proxy Authentication Required"
        at sun.net.www.protocol.http.HttpURLConnection.doTunneling(HttpURLConnection.java:2124)
        at sun.net.www.protocol.https.AbstractDelegateHttpsURLConnection.connect(AbstractDelegateHttpsURLConnection.java:183)
        at sun.net.www.protocol.https.HttpsURLConnectionImpl.connect(HttpsURLConnectionImpl.java:153)
        at Sample.main(Sample.java:25)
```

コード内での認証プロキシ設定だけでなく、下記も試しましたが結果はかわらず。

1. Java起動オプションにproxyの設定をする
  * 例: `java -Dhttp.proxyHost="example.proxy.com" -Dhttp.proxyPort=8080 -Dhttp.proxyUser=user -Dhttp.proxyPassword=password (httpsに関しても同様に設定) Sample`
2. Systemデフォルトを参照するよう設定
  * 例: `java -Djava.net.useSystemProxies=true Sample`

## 環境

OSはUbuntu 16.04、Javaのバージョンは下記です。

```
$ java -version
java version "1.8.0_111"
Java(TM) SE Runtime Environment (build 1.8.0_111-b14)
Java HotSpot(TM) 64-Bit Server VM (build 25.111-b14, mixed mode)
```

# 結論

2016/10にリリースされた8u111以降から、デフォルトの設定としてHTTPS接続する際のトンネリングにBasic認証を利用することができなくなりました。

[Java 8リリースの変更](https://www.java.com/ja/download/faq/release_changes.xml)

> HTTPSトンネリングのBasic認証の無効化
> 一部の環境で、HTTPSプロキシ時に特定の認証スキームが不適切である場合があります。そのため、Basic認証スキームは、デフォルトではOracle Java Runtimeでjdk.http.auth.tunneling.disabledSchemesネットワーク・プロパティにBasicを追加することで非アクティブ化されています。HTTPSのトンネルの設定時にBasic認証を必要とするプロキシは、デフォルトでは成功しないようになりました。必要な場合、jdk.http.auth.tunneling.disabledSchemesネットワーク・プロパティからBasicを削除するか、コマンド行で同じ名前のシステム・プロパティを"" (空)に設定することで、この認証スキームを再度アクティブ化できます。また、jdk.http.auth.tunneling.disabledSchemesおよびjdk.http.auth.proxying.disabledSchemesネットワーク・プロパティと同じ名前のシステム・プロパティは、それぞれHTTPSのトンネルの設定時またはプレーンなHTTPのプロキシ時に、アクティブ化されている可能性があるその他の認証スキームを無効化するために使用できます。JDK-8160838 (非公開)

[英語版のリリースノート](http://www.oracle.com/technetwork/java/javase/8u111-relnotes-3124969.html)を見ると、`core-libs/java.net`と書いてあるので、`java.net`パッケージ内のクラスの機能を使ってHTTPS接続する際のデフォルトの挙動の変更だと思われます。


# 回避方法

1. JVMオプション`jdk.http.auth.tunneling.disabledSchemes=""`を設定する。
  * 実行時: `java -Djdk.http.auth.tunneling.disabledSchemes="" Sample`
  * コード内: `System.setProperty("jdk.http.auth.tunneling.disabledSchemes", "");`
2. 他のライブラリを使う
  * ApacheのHTTPClientなど

# 補足

`System.getProperty("jdk.http.auth.tunneling.disabledSchemes")`をsysoutしてみましたが、結果は`null`でした。

