---
title: "[Mattermost Integrations] Plugin (Webapp 2)"
date: 2020-12-21T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "webapp"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第21日目の記事です。

昨日に引き続き、Mattermost上の様々な操作に対応した処理を追加できるMattermost Pluginの**Webapp**サイドの機能を紹介していきます。。

Mattermost Pluginについての公式ドキュメントは下記になります。
https://developers.mattermost.com/extend/plugins/overview/

サンプルコードは下記リポジトリにコミットしています。
https://github.com/kaakaa/mattermost-plugin-api-sample

## Webapp Plugin API

### [registerWebSocketEventHandler](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerWebSocketEventHandler)
`registerWebSocketEventHandler`は、Mattermost Serverから送信されるWebSocketイベントを処理するHandlerを登録します。

`registerWebSocketEventHandler`は2つの引数を取ります。

* `event`: 処理するWebSocketイベントの種別
* `handler`: WebSocketイベントのデータを引数に取る関数。引数のデータはイベントによって異なります。


Serverプラグインの[PublishWebSocketEvent](https://developers.mattermost.com/extend/plugins/server/reference/#API.PublishWebSocketEvent)と組み合わせて使用すると強力ですが、その辺りの例については22日目以降の記事で紹介します。

ここでは、MattermostデフォルトのWebSocketイベントである投稿が作成された際に送信される`posted`イベントを受信した際に、投稿内容に`open modal`という文言が入っていた場合にモーダルを開く例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/websocket-event-handler.gif)

```js:index.js
...
import {openRootModal, createPluginPost} from './actions';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 'open modal'を含む投稿を受信するとモーダルを開く
        registry.registerWebSocketEventHandler(
            'posted',
            (event) => {
                const post = JSON.parse(event.data.post);
                if (post && post.message && post.message.includes('open modal')) {
                    store.dispatch(openRootModal());
                }
            }
        );
    }
}
```

### [unregisterWebSocketEventHandler](https://developers.mattermost.com/extend/plugins/webapp/reference/#unregisterWebSocketEventHandler)
`unregisterWebSocketEventHandler`は、`registerWebSocketEventHandler`によってWebSocketイベントに対して設定されたHandlerを登録から除外します。

例は省略します。

### [registerReconnectHandler](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerReconnectHandler)

`registerReconnectHandler`は、一度インターネット接続が失われた後に再びMattermostへ接続した際に実行されるHandlerです。

`registerReconnectHandler`は引数のない関数を引数に取ります。

* `handler`: Mattermostへ再接続した際に実行される関数

例は省略します。

### [unregisterReconnectHandler](https://developers.mattermost.com/extend/plugins/webapp/reference/#unregisterReconnectHandler)
`unregisterReconnectHandler`は、`registerReconnectHandler`で登録したHandlerを登録から除外します。引数は取りません。

こちらも例は省略します。

### [registerMessageWillBePostedHook](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerMessageWillBePostedHook)

`registerMessageWillBePostedHook`は、ユーザーが投稿を送信した際、その投稿がサーバーに送信される前に実行される処理を登録します。

`registerMessageWillBePostedHook`は、引数を1つ取ります。

* `hook`: ユーザーによって投稿処理が実行された際、その投稿がサーバーに送信される直前に実行される処理

`hooks`は、引数を一つ取ります。

* `post`: 投稿の情報

`hook`の返り値として、投稿情報を持つ`error`フィールドを含む値を返却した場合、投稿はrejectされます。投稿情報を持つ`post`フィールドのを含むオブジェクトを返却した場合は、`post`フィールドの値で投稿が作成されます。

`忙しい`という文言を含む投稿を作成するとrejectされ、`帰りたい`という文言を含む投稿を作成すると`仕事したい`に変換される例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/message-will-be-posted-hook.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 投稿がサーバーに送信される前にrejectしたり内容を変換したりする
        registry.registerMessageWillBePostedHook(
            (post) => {
                if (post.message && post.message.includes('忙しい')) {
                    return {error: {message: '忙しくはないはずです'}};
                }
                post.message = post.message.replace(/帰りたい/gi, '仕事したい');
                return {post: post};
            }
        );
    }
}
```

### [registerSlashCommandWillBePostedHook](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerSlashCommandWillBePostedHook)

`registerSlashCommandWillBePostedHook`は、ユーザーがSlash Commandを実行した際、その投稿がサーバーに送信される前に実行される処理を登録します。

`registerSlashCommandWillBePostedHook`は、引数を1つ取ります。

* `hook`: Slash Commandが実行された際に、その内容がサーバーに送信される直前に実行される処理

`hook`は、引数を2つ取ります。

* `message`: 投稿されたメッセージ
* `args`: Slash Command実行情報(`channel_id`, `team_id`, `root_id`, `parent_id`)

`/away`をreject、`/help`を`/shrug`に書き換え、`/leave`をエラーメッセージなしでrejectするような処理を実行する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/slash-command-will-be-posted-hook.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // Slash Commandがサーバーに送信される前に実行される処理を追加する
        registry.registerSlashCommandWillBePostedHook(
            (message, args) => {
                console.log(message);
                if (message.startsWith('/away')) {
                    return {error: {message: 'rejected'}};
                }
                if (message.startsWith('/help')) {
                    console.log('help');
                    return {message: '/shrug converted from help command', args};
                }
                if (message.startsWith('/leave')) {
                    console.log('leave');
                    return {};
                }
            }
        );
    }
}
```

### [registerMessageWillFormatHook](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerMessageWillFormatHook)
`registerMessageWillFormatHook`は、投稿したメッセージがMarkdownテキストとして変換される直前に実行される処理を登録します。

`registerMessageWillFormatHook`は、引数を1つ取ります。

* `hook`: 投稿がMarkdownテキストとして変換される直前に実行される処理

`hook`は、引数を2つ取ります。

* `post`: 変換されていない投稿情報
* `message`: 投稿されたメッセージ（プラグインによって変換されている可能性あり）

`hook`の返却値として返された文字列が投稿として表示されます。

`registerMessageWillBePostedHook`は、投稿がサーバーに送信される前に投稿内容を編集するものでしたが、この`registerMessageWillFormatHook`は、投稿がサーバーに送信・保存された後にレンダリングされる際にメッセージを編集するものだと思われます。

良い利用方法が思いつかないので例は省略します。

### [registerFilePreviewComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerFilePreviewComponent)

`registerFilePreviewComponent`は、ファイルプレビュー用の独自のComponentを登録します。

`registerFilePreviewComponent`は、2つの関数を引数に取ります。

* `override`: 独自のファイルプレビューComponentを使用するかどうかを決定するための関数。以下の2つを引数に取ります。
* `component`: ファイルプレビュー用のComponent

`override`は、引数を2つ取ります。

* `fileInfo`: ファイルの情報
* `post`: 投稿の情報


`debug`で始まるメッセージを持つ投稿の添付ファイルをプレビューする際に、独自のコンポーネントを使用する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/file-preview-component.gif)


```js:index.js
...
import CustomFilePreview from './components/custom_file_preview';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // `debug`で始まるメッセージを持つ投稿の添付ファイルを独自コンポーネントでプレビューする
        registry.registerFilePreviewComponent(
            (fileInfo, post) => { return post.message && post.message.startsWith('debug'); },
            CustomFilePreview
        );
    }
}
```

```js:components/custom_file_preview.jsx
import React from 'react';
import PropTypes from 'prop-types';

const {formatText, messageHtmlToComponent} = window.PostUtils;

const CustomFilePreviewComponent = ({fileInfo, post}) => {
    const formattedText = messageHtmlToComponent(formatText(post.message));

    return (
        <div style={{backgroundColor: '#ffcccc'}}>
            {formattedText}
            <pre>
                {JSON.stringify(fileInfo, null, 4)}
            </pre>
        </div>
    )
}

CustomFilePreviewComponent.propTypes = {
    fileInfo: PropTypes.object.isRequired,
    post: PropTypes.object.isRequired,
};

export default CustomFilePreviewComponent;
```

### [registerTranslations](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerTranslations)
`registerTranslations`は、プラグイン内で使用しているメッセージの翻訳データを登録します。

`registerTranslations`は、`locale`を引数にとり、その`locale`に対する翻訳データを返す関数を引数に取ります。

RootModal内のメッセージを日本語に翻訳する例を作ってみましたが、どうやら翻訳が正常に動作していない模様？

```js:index.js
...
import ja from 'i18n/';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 翻訳メッセージを登録する
        registry.registerTranslations((locale) => {
            switch (locale) {
            case 'en':
                return en;
            case 'ja':
                return ja;
            }
            return {};
        });
    }
}
```

```json:i18n/ja.json
{
    "rootModal.message": "ルートモーダル"
}
```

```js:components/root/root.jsx

import React from 'react';

import {FormattedMessage} from 'react-intl';

const Root = ({visible, close}) => {
    if (!visible) {
        return null;
    }

    const style = getStyle();

    return (
        <div
            style={style.backdrop}
            onClick={close}
        >
            <div style={style.modal}>
                <FormattedMessage
                    id='rootModal.message'
                    defaultMessage='Root Modal2'
                />
            </div>
        </div>
    );
};

Root.propTypes = {
    visible: PropTypes.bool.isRequired,
    close: PropTypes.func.isRequired,
};

const getStyle = () => ({
    backdrop: {
        position: 'absolute',
        display: 'flex',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        backgroundColor: 'rgba(0, 0, 0, 0.50)',
        zIndex: 2000,
        alignItems: 'center',
        justifyContent: 'center',
    },
    modal: {
        height: '300px',
        width: '500px',
        padding: '1em',
        color: 'black',
        backgroundColor: 'white',
    },
});

export default Root;
```

### [registerAdminConsolePlugin](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerAdminConsolePlugin)
`registerAdminConsolePlugin`は、AdminConsole(システムコンソール?)の内容を上書きするための関数を登録します。

Mattermost内部での利用が主目的であり、ユーザープラグインによる使用は推奨されていないようなので例は省略します。（使い方がよく分からない）

### [unregisterAdminConsolePlugin](https://developers.mattermost.com/extend/plugins/webapp/reference/#unregisterAdminConsolePlugin)
`unregisterAdminConsolePlugin`は、`registerAdminConsolePlugin`で登録した　AdminConsole上書き用関数を登録から除外します。
`registerAdminConsolePlugin`がMattermost内部での利用が主目的のため、こちらも説明、例は省略します。

### [registerAdminConsoleCustomSetting](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerAdminConsoleCustomSetting)
`registerAdminConsoleCustomSetting`は、プラグイン用の設定画面に独自のコンポーネントを追加します。
このプラグインAPIについては、公式ドキュメントの[Best Practices](https://developers.mattermost.com/extend/plugins/best-practices/#how-can-a-plugin-define-its-own-setting-type)のページに詳細にまとめられています。


Mattermost Pluginでは、マニフェストファイルの`settings_schema`という項目を指定することで、プラグイン専用の設定項目を作成することができます。

https://developers.mattermost.com/extend/plugins/manifest-reference/

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/admin-console-custom-settings-default.png)

デフォルトでは、下記の`type`を持つ設定項目を追加することができます。

* "bool": ラジオボタン (2値)
* "dropdown": ドロップダウンメニュー
* "generated": 自動生成テキスト
* "radio": ラジオボタン (任意)
* "text": インプットボックス
* "longtext": 複数行テキスト
* "number": 数値入力欄
* "username":ユーザー名入力欄

デフォルトの`type`以外の設定項目を指定したい場合に`registerAdminConsoleCustomSetting`を使用します。

`registerAdminConsoleCustomSetting`は、3つの引数を取ります。

* `key`: 上書きする設定項目の`key`。`key`は、マニフェストファイルに設定項目ごとに任意で指定する値です。
* `component`: 設定画面に表示されるComponent。
* `options`: 設定項目の表示方法についてのオプション
  * `showTitle`: `true`を指定した場合、マニフェストファイルの`display_name`が設定項目の左側に表示されます

パスワードなどを入力する際に、入力項目をUI上に表示しないような設定項目を追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/admin-console-custom-settings.png)

```json:plugin.json
{
    ...
    "settings_schema": {
        "settings": [
            {
                "key": "SampleSetting",
                "display_name": "Sample Setings Value",
                "type": "text",
                "help_text": "Sample",
                "default": "sample"
            }
        ]
    }
}
```

```js:index.js
...
import CustomSettings from './components/custom_settings';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 独自の設定画面項目を追加する
        registry.registerAdminConsoleCustomSetting(
            'SampleSetting',
            CustomSettings,
            {showTitle: true}
        );
    }
}
```

```js:components/custom_settings.jsx
import React from 'react';
import PropTypes from 'prop-types';

const CustomSettingsComponent = ({helpText, id, onChange, value}) => {
    const handleChange = (e) => {
        onChange(id, e.target.value);
    }
    return (
        <div style={{backgroundColor: '#ffcccc'}}>
            <input
              type={'password'}
              value={value}
              onChange={handleChange}
            />
            <pre>
                {JSON.stringify(helpText.props, null, 4)}
            </pre>
        </div>
    )
}

CustomSettingsComponent.propTypes = {
    helpText: PropTypes.shape ({
        props: PropTypes.object,
    }),
    id: PropTypes.string,
    onChange: PropTypes.func,
    value: PropTypes.any,
};

export default CustomSettingsComponent;
```


### [registerRightHandSidebarComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerRightHandSidebarComponent)
`registerRightHandSidebarComponent`は、Mattermostの右サイドバーに表示する独自のComponentを登録します。

`registerRightHandSidebarComponent`は、2つの引数を取ります。

* `component`: 右サイドバーに表示されるComponent
* `title`: 右サイドバーのタイトル部分に表示されるテキスト

また、`registerRightHandSidebarComponent`は4つの引数を返却します。

* `id`: 登録されたComponentのID
* `showRHSPlugin`: 登録したComponentを表示するためのアクション
* `hideRHSPlugin`: 登録したComponentを非表示にするためのアクション
* `toggleRHSPlugin`: 登録したComponentの表示/非表示を切り替えるアクション

登録したComponentは、他のプラグインAPIから`shorRHSPlugin`、`hideRHSPlugin`、`toggleRHSPlugin`のアクションを実行することで表示されるようになります。`RHS`は`RightHandSideber`の略です。

右サイドバーに独自のComponentを表示するためのメニューをメインメニューに追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/right-hand-sidebar.gif)

```js:index.js
...
import CustomRightHandSideber from './components/custom_rhs';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 右サイドバーに表示される独自Componentを登録する
        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(CustomRightHandSideber, "Sample RHS")
        // 右サイドバーを表示するためのメインメニューを追加する
        registry.registerMainMenuAction(
            'Open RHS',
            () => store.dispatch(toggleRHSPlugin),
            () => (<i/>)
        );
    }
}
```


```js:components/custom_rhs.jsx
import React from 'react';
import PropTypes from 'prop-types';

const ComponentRightHandSidebar = ({theme}) => {
    return (
        <div style={{backgroundColor: '#ffcccc'}}>
            <pre>
                {JSON.stringify(theme, null, 4)}
            </pre>
        </div>
    )
}

ComponentRightHandSidebar.propTypes = {
    PluggableId: PropTypes.string.isRequired,
    theme: PropTypes.object.isRequired,
};

export default ComponentRightHandSidebar;
```

### [registerNeedsTeamRoute](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerNeedsTeamRoute)
`registerNeedsTeamRoute`は、チームごとにプラグイン専用のRouteを追加します。プラグイン専用のエラー画面を作成したい場合などに使用するものだと思います。

`registerNeedsTeamRoute`は、2つの引数を取ります。

* `route`: ルート文字列
* `comopnent`: `route`にアクセスされた際に呼び出されるComponent

`http://localhost:8065`でMattermostが起動していて、`test`というチーム名のチームがあり、`sample.plugin`というプラグインIDを持つプラグインがインストールされており、その中で`registerNeedsTeamRoute`の引数として`route="/subpath"`が指定されていた場合、`http://localhost:8065/test/sample.plugin/subpath`にアクセスすると、`component`に指定したコンポーネントが呼び出されます。

以下に例を示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day21/needs-team-route.gif)

```js:index.js
...
import CustomTeamRoute from './components/custom_team_route';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 独自のRouteを追加する
        registry.registerNeedsTeamRoute('/', CustomTeamRoute)
    }
}
```


```js:components/custom_team_route.jsx
import React from 'react';
import PropTypes from 'prop-types';

import {Switch, Route} from 'react-router-dom';
import {getCurrentTeam} from 'mattermost-redux/selectors/entities/teams';
import {useSelector} from 'react-redux';

import {id} from 'src/manifest';

const CustomTeamRouteComponent = () => {
    const currentTeam = useSelector(getCurrentTeam);
    return (
        <Switch>
            <Route path={`/${currentTeam.name}/${id}/error`}>
                <h3>{'This is error page!'}</h3>
                <p>
                    <a href={`/${currentTeam.name}`}>Back to Top</a>
                </p>
            </Route>
            <Route>
                <h3>{'404 Not Found'}</h3>
                <p>
                    <a href={`/${currentTeam.name}`}>Back to Top</a>
                </p>
            </Route>
        </Switch>
    )
}

CustomTeamRouteComponent.propTypes = {
    pluggableId: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};

export default CustomTeamRouteComponent;
```

### [registerCustomRoute](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerCustomRoute)
`registerCustomRoute`は、プラグイン専用のRouteを追加します。

`registerNeedsTeamRoute`は、`/${team_name}/${plugin_id}/${route}`のようにチームごとにRouteを追加するAPIでしたが、`registerCustomRoute`は`/plug/${plugin_id}/${route}`のように、Mattermost全体としてプラグインごとに一つのRouteを追加するAPIです。

使い方は`registerNeedsTeamRoute`とほぼ同じのため、例は省略します。

## その他

Mattermost Plugin Webapp開発中に使えるPlugin開発用の便利機能がいくつかあります。その概要だけ紹介します。

### [Theme](https://developers.mattermost.com/extend/plugins/webapp/reference/#theme)
Mattermost Plugin APIの中でも何度か出てきましたが、Webapp PluginではMattermostのテーマカラーを参照することができます。Mattermostではユーザーごとにテーマカラーを変更することができるため、Webapp PluginでUIの色を指定する場合は、ユーザーごとに見え方が異なることを考慮に入れる必要があります。

参考: [Mattermostのテーマ集 \- Qiita](https://qiita.com/kaakaa_hoe/items/45c8857589ccd822ab1a)

Mattermostで扱われるテーマカラー一覧は以下で紹介されています。

https://developers.mattermost.com/extend/plugins/webapp/reference/#theme

### [Exported Libraries and Functions](https://developers.mattermost.com/extend/plugins/webapp/reference/#exported-libraries-and-functions)

Mattermost Webapp PluginはReact.jsを使用して開発しますが、React開発によく使われるいくつかのライブラリはMattermost本体から`window`オブジェクトを介して取得できるようになっています。
取得できるライブラリは以下で紹介されています。

https://developers.mattermost.com/extend/plugins/webapp/reference/#exported-libraries-and-functions

また、`window`オブジェクトから参照できるライブラリとして`window.PostUtils`というのがありますが、これはMattermostフォーマットのテキストを扱うための便利関数を持つオブジェクトです。

以下のようにすることで、@メンションなどを含むMarkdownテキスト(`text`)をフォーマットして扱うことができます。

```js
const {formatText, messageHtmlToComponent} = window.PostUtils;

const text = '...';
const formattedText = messageHtmlToComponent(formatText(text));
```

https://developers.mattermost.com/extend/plugins/webapp/reference/#post-utils

### [Redux Action](https://developers.mattermost.com/extend/plugins/webapp/actions/))
Webapp上で投稿やユーザー情報の取得などのMattermostに対する何かしらの処理を実行する場合、[mattermost-redux](https://github.com/mattermost/mattermost-redux)というReduxライブラリがあります。これはMattermost本体のWebappでも利用されている公式のJavascript APIのような位置付けのものです。

mattermost-reduxはもちろんMattermost Plugin開発でも使用することができ、下記のページで使い方について紹介されています。

https://developers.mattermost.com/extend/plugins/webapp/actions/

## さいごに

本日は、Mattermost Pluginの**Webapp**サイドの実装について紹介しました。
明日からは、Mattermost上で投票機能を使うことができるMatterPollプラグインの実装について紹介していきます。
