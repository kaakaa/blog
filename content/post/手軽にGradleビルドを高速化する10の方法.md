
---
title: "手軽にGradleビルドを高速化する10の方法"
date: 2016-07-05T22:32:49+09:00
draft: false
toc: true
tags: ["gradle"]
---

# はじめに

Gradleのビルドを高速化する方法について、下記エントリを参考に調べてみました。
[10 Tips to Improve Your Gradle Build Time — Medium](https://medium.com/@kenodoggy/10-tips-to-improve-your-gradle-build-time-1ddeca924759#.hsxgddism)

記事の冒頭でも書かれている通り、上記のエントリはAndroidプロジェクト向けに書かれていますが、AndroidプロジェクトでなくてもGradleプロジェクトならば適用できる方法ばかりです。
（一つAndroid限定のものもありますが）
特にマルチプロジェクト構成の際に効果が高いものばかりなので、大規模なGradleプロジェクトを構築する際の参考になりそうです。

ただ、最後のまとめにも書きましたが、ビルド時間の大半を占めるのがテスト実行時間であることが多いため、劇的にビルドが早くなるという訳では無いかと思いました。

## 0. ドキュメント

Gradleのコマンドラインオプションと実行時オプションを利用する方法が多いため、公式ドキュメントを載せておきます。

[Appendix D. Gradle Command Line](https://docs.gradle.org/current/userguide/gradle_command_line.html)
[The Build Environment - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/build_environment.html)


## 1. デーモン化

これは昔から推奨されてましたね。
Gradleプロセスを常駐化させることで、JVMの起動や必須ライブラリの読み込みを毎回行わなくて済むようにする方法。

[The Gradle Daemon - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/gradle_daemon.html)

### 指定方法

```gradle.properties
org.gradle.daemon=true
```

or

```command:cli
./gradlew　--daemon # 一度起動するとそのまま常駐します
./gradlew build
```


## 2. 並列ビルド (incubating)

**マルチプロジェクトビルド限定**です。
互いに参照しあっていないプロジェクト([Decoupled Project](https://docs.gradle.org/current/userguide/multi_project_builds.html#sec:decoupled_projects))のビルドを並列実行できます。

### 指定方法

```gradle.properties
org.gradle.parallel=true
org.gradle.workers.max=4  # 並列の最大数(デフォルトではプロセッサの数)
```

or

```command:cli
./gradlew build --parallel --max-workers=4
./gradlew build --parallel --parallel-threads=4
```

`--max-workers`と`--parallel-thread`はどちらも`org.gradle.workers.max`と同じく、並列の最大数を指定するオプションです。
同時に指定された場合、`--parallel-thread`が優先されるそうです。


## 3. オンデマンドConfigure (incubating)

[Multi-project Builds - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/multi_project_builds.html#sec:configuration_on_demand)

GradleビルドのConfigure Phaseを必要なときだけ行うということですかね。
ちょっとここらへん理解が浅いです。

### 指定方法

```gradle.properties
org.gradle.configureondemand=true
```

or

```command:cli
./gradlew build --configure-on-demand
```


## 4. グローバル設定

今まで紹介した手法と趣向が異なります。

1-3でも登場した`gradle.properteis`ファイルは、下記のいずれかの場所に置くことでビルド時に読み込まれます。

* Gradleプロジェクトのトップディレクトリ
* `$USER_HOME/.gradle/gradle.properties`
* `$GRADLE_USER_HOME/gradle.properties`

1-3に挙げているようなオプションは、プロジェクト関係なく有効な手法なのでユーザ設定にしてしてしまっても良さそうです。

[The Build Environment - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/build_environment.html#sec:gradle_configuration_properties)


## 5. 最小化

ビルド実行時のコマンドラインオプションで実行しないタスクを指定できます。
[Using the Gradle Command-Line - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/tutorial_gradle_command_line.html)

### 指定方法

```command:cli
./gradlew build -x test
./gradlew build --exclude-task test
```

## 6. minSdkVersionを設定してARTのみを対象に

**Androidプロジェクト限定**です。
Android SDKの最低バージョンを`23`に設定することでDalvik向けのdexing?をしないようにするみたいです。

Android Studioのドキュメントでも紹介されているやり方のようで。
[Configure Apps with Over 64K Methods | Android Studio](https://developer.android.com/studio/build/multidex.html#dev-build)

### 指定方法

```build.gradle
...

def minSdk = hasProperty('minSdk') ? minSdk : 16

android {
    compileSdkVersion 23
    buildToolsVersion "23.0.2"

    dexOptions {
        javaMaxHeapSize "4g"
        jumboMode true
    }

    defaultConfig {
        applicationid "com.myapplication.app"

        minSdkVersion minSdk
        targetSdkVersion 23
...

```

```command:cli
./gradlew assembleDebug -PminSdk=23
```

Android Studioでの指定方法については下記を参考に。
[How to make faster Android build without sacrificing new api lint check](https://gist.github.com/cesarferreira/8480ea6fd0b95ba57f98)


## 7. オフライン実行

Gradleがネットワークリソースへアクセスするのを抑止するオプションのようです。
一度、ビルドに必要な資産をダウンロードすれば、以降はオフラインでも良さそう。

### 指定方法

```command:cli
./gradlew build --offline
```


## 8. Gradleのバージョンを2.4まで上げる

公式ブログでもアナウンスがあったように、Gradle2.4では2.3の時よりconfiguration timeが34%も削減されているそうで。
[Gradle 2.4: The fastest yet - Gradle](https://gradle.org/blog/gradle-2-4-the-fastest-yet/)

現在、Gradle2.14までリリースされているので、2.3以下を使っている場合はちょっと古すぎるので上げておくべきかと。


## 9. jCenterの利用

MavenCentralよりjCenterを利用しましょうということらしいです。
確かにGradleのbuild-init pluginを使って生成したbuild.gradleでも、jCenterを利用するよう設定されています。

理由については、下記のエントリが紹介されています。
* [Android Studio – Migration from Maven Central to JCenter | Blog @Bintray](https://blog.bintray.com/2015/02/09/android-studio-migration-from-maven-central-to-jcenter/)
* [Issue 72061 - android - The Maven index is *huge*! - Android Open Source Project - Issue Tracker - Google Project Hosting](https://code.google.com/p/android/issues/detail?id=72061)
* [Feel secure with SSL? Think again. | Blog @Bintray](https://blog.bintray.com/2014/08/04/feel-secure-with-ssl-think-again/)


## 10. Profile!

ちゃんとプロファイリングして、自分たちのビルドの内容を把握しておこうということですね。
`--profile`オプションを付けてビルドを実行すると、ビルドプロファイリングのHTMLファイルが出力されます。

[Using the Gradle Command-Line - Gradle User Guide Version 2.14](https://docs.gradle.org/current/userguide/tutorial_gradle_command_line.html#sec:profiling_build)

また、Gradle Inc.は[Gradle Build Scans](https://scans.gradle.com/)というサービスを開始しているので、コレを利用するのも良いかもしれません。
このサービスについては、私のブログで少し紹介しています。
[Gradleがビルド結果解析サービス Gradle.com を開始していた - kaakaa Blog](http://kaakaa.hatenablog.com/entry/2016/06/15/000223)


### 指定方法

```command:cli
./gradlew build --profile
```


## さいごに

いくつかの高速化手法について、[Netflix/Hystrix](https://github.com/Netflix/Hystrix)を対象にビルド所要時間を測定してみましたが、やはりビルドで一番時間がかかるのはテスト実行のため、ほぼ高速化を体感できませんでした。
ただし、並列ビルドについてはサブプロジェクトのテストが並列に実行されるため、とても効果がありました。

