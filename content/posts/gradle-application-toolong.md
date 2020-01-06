---
title: "Gradle Application プラグインで生成したスクリプトが実行できないことがある"
date: 2020-01-06T18:05:18+09:00
draft: false
toc: true
tags: ["Gradle", "Application plugin"]
---

# 要約
* Windowsのcmd.exeの制約により、GradleのApplicationプラグインで生成したBatファイルが実行不可能になることがあった
* Applicationプラグインが出力するBatファイル内でのクラスパスの指定を、**jarファイル指定でなくディレクトリ指定となるよう**設定を追加することにより回避可能となった
* 検証用プロジェクト: https://github.com/kaakaa/gradle-application-too-long-input-line

# 背景
[GradleのApplicationプラグイン](https://docs.gradle.org/current/userguide/application_plugin.html)を利用すると、CLIツールとして作成したJavaアプリを実行するためのスクリプトを生成することができる。

```build.gradle
plugins {
    id 'application'
}

application.mainClassName = 'sample.Main'
```

`main` メソッドを持つクラスを `application.mainClassName` として指定し、[Gradle Applicationプラグイン](https://docs.gradle.org/current/userguide/application_plugin.html)の `installDist` タスクを実行することで、`$buildDir/install/${APPLICATION_NAME}` 配下に実行可能な形式のアプリケーションが出力される。

```
$ ./gradlew installDist
$ tree build/install/
build/install/
└── app
    ├── bin
    │   ├── app
    │   └── app.bat
    └── lib
        ├── deps.jar
        ...
```

`bin` ディレクトリ配下にはUnix実行用のスクリプトファイルとWindows実行用のBATファイルが出力される。アプリケーションが依存するライブラリは `lib` 配下にまとめられ、実行用スクリプトファイルから `lib` 配下の依存ライブラリが参照されるようになっている。

# 問題
Windows環境において、Applicationプラグインから出力されたBATファイルを実行すると、エラーとなることがあった。

```
> build/install/gradle-application-too-long-input-line/bin/gradle-application-too-long-input-line.bat
The input line is too long.
The syntax of the command is incorrect.
```

[GitHub Actionsでの実行結果](https://github.com/kaakaa/gradle-application-too-long-input-line/commit/db4f53c31faabe65d4682ef091002abf82c8e754/checks?check_suite_id=385831791#step:5:7)


# 原因
Windowsのcmd.exeでは１行が**8191文字**を超えるようなコマンドは実行できない。

[コマンド プロンプト \(cmd\.exe\) のコマンドライン文字列の制限](https://support.microsoft.com/ja-jp/help/830473/command-prompt-cmd-exe-command-line-string-limitation)

> Microsoft Windows XP 以降の Windows を実行しているコンピューターでは、コマンド プロンプトで使用できる文字列の最大長は 8191 文字です。

GradleのApplicationプラグインは、実行スクリプトファイル内で依存ライブラリを `set CLASSPATH=%APP_HOME%\lib\deps1.jar;%APP_HOME%\lib\deps2.jar;...` のようにjarファイル単位でクラスパス文字列に結合しているため、依存ライブラリが多すぎると **8191文字制約** に引っかかることがある。

[サンプルBATファイル](https://github.com/kaakaa/gradle-application-too-long-input-line/blob/9aa12fa00228b8240679b32a5f80e3514a1f2b9d/build/install/gradle-application-too-long-input-line/bin/gradle-application-too-long-input-line.bat#L82)


# 解決方法
実行用スクリプトファイルを生成した後、クラスパス指定に関する行をディレクトリ参照1つで済むように文字列置換することでこのエラーは回避可能。

```build.gradle
...
tasks.withType(CreateStartScripts).each { task ->
    task.doLast {
        String text = task.windowsScript.text
        text = text.replaceFirst(/(set CLASSPATH=%APP_HOME%\\lib\\).*/, { "${it[1]}*" })
        task.windowsScript.write text
    }
}
...
```
参考: [Workaround for gradle application plugin 'the input line is too long' error on Windows](https://gist.github.com/jlmelville/2bfe9277e9e2c0ff79b6)

上記の設定を `build.gradle` に追加し、再度BATファイルを生成するとエラーとなっていた行は下記のように置換されるため、文字数制限に抵触しなくなる。

```
...
set CLASSPATH=%APP_HOME%\lib\*
...
```
[サンプルBATファイル](https://github.com/kaakaa/gradle-application-too-long-input-line/blob/master/build/install/gradle-application-too-long-input-line/bin/gradle-application-too-long-input-line.bat#L82
)

この設定を加えることで、エラーが解消できた。

[GitHub Actionsでの実行結果](https://github.com/kaakaa/gradle-application-too-long-input-line/commit/0eb5385d110c4e56ff7a288ac8fba120c73d2b1a/checks?check_suite_id=385886281#step:5:7)


---

(別の解決方法として、fatjarを作成してそのfatjarのみをクラスパスで参照するよう書き換えれば良いかとも思いましたが、`lib`ディレクトリ内のその他の依存関係を排除する方法が分からなかったので、そちらの方法は断念しました)

