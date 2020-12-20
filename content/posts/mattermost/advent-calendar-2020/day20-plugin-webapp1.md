---
title: "[Mattermost Integrations] Plugin (Webapp)"
date: 2020-12-20T00:00:00+09:00
draft: false
toc: true
tags: ["mattermost", "integration", "plugin", "webapp"]
---

Mattermost記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

## 本記事について

[Mattermostの統合機能アドベントカレンダー](https://qiita.com/advent-calendar/2020/mattermost-integrations)の第20日目の記事です。

本記事では、Mattermost上の様々な操作に対応した処理を追加できるMattermost Pluginの**Webapp**サイドの機能を紹介していきます。

記事が長いので2日に分けて紹介していきます。それでも長いです。

Mattermost Pluginについての公式ドキュメントは下記になります。
https://developers.mattermost.com/extend/plugins/overview/

サンプルコードは下記リポジトリにコミットしています。
https://github.com/kaakaa/mattermost-plugin-api-sample

## Mattermost Plugin (Webapp) について

Mattermost Plugin (Webapp) の本体について紹介します。

Mattermostの**Webapp**サイドのPluginを実装する場合、プラグイン本体は`registry`と`store`を引数にとる`initialize`メソッドを持つクラスです。`registry`はMattermost Plugin (Webapp)の機能を呼び出すためのメソッドを持つオブジェクトで、`store`はMattermostのデータにアクセスするためのインターフェースです。

公式ドキュメントでは、以下のような`PluginClass`が紹介されています。
https://developers.mattermost.com/extend/plugins/webapp/reference/

```js
class PluginClass {
    /**
    * initialize is called by the webapp when the plugin is first loaded.
    * Receives the following:
    * - registry - an instance of the registry tied to your plugin id
    * - store - the Redux store of the web app.
    */
    initialize(registry, store)

    /**
    * uninitialize is called by the webapp if your plugin is uninstalled
    */
    uninitialize()
}
```

この`PluginClass`のインスタンスを`window`オブジェクトが持つ`registerPlugin`関数にPlugin IDと共に与えることで、Pluginが有効になります。

```js
window.registerPlugin('myplugin', new PluginClass());
```

## Webapp Plugin API

### [registerRootComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerRootComponent)
`registerRootComponent`は、画面全体に表示されるモーダルのReact Componentを登録するAPIです。Componentが登録されると、Componentを一意に識別するためのIDが返却されます。

ここで登録したRoot Componentは、登録するだけで表示されるわけではなく、別のアクションから起動されることで表示されます。

登録されたComponentで使用できるpropsは特にありません。[channel_controller.jsx](https://github.com/mattermost/mattermost-webapp/blob/61a4de5e36f994955bacc47c8d56734f9a699a5a/components/channel_layout/channel_controller.jsx#L86)

今回は、チャンネルのヘッダー部分にボタンを追加する`registerChannelHeaderButtonAction` APIを使ってRoot Componentを表示する例を以下に示します。全てのプログラムを記述すると非常に長くなってしまうため、Reduxによるデータのやりとりなど部分については本記事では割愛します。プログラム全体について知りたい場合は、[kaakaa/mattermost-plugin-api-sample](https://github.com/kaakaa/mattermost-plugin-api-sampleを参照してください。)
また、Mattermost公式チームにより作成されているデモ用プラグイン[mattermost-plugin-demo](https://github.com/mattermost/mattermost-plugin-demo/tree/master/webapp)の実装も参考になります。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/root-component.gif)

```js
...

import {openRootModal} from './actions'
import reducer from './reducer';

import Root from './components/root';

...

export default class Plugin {
    initialize(registry, store) {
        // (1) Root Componentの登録
        registry.registerRootComponent(Root)

        // (2) Root Componentを呼び出すアクションの登録
        registry.registerChannelHeaderButtonAction(
            channelHeaderButtonIcon,
            () => store.dispatch(openRootModal()),
            "Open Root modal","Open Root modal"
        );

        registry.registerReducer(reducer);
    }
}
```

上記の例では、`(1)`でRoot ComponentとなるReact Componentを登録し、`(2)`でRoot Componentを表示するアクションをチャンネルヘッダーボタンに与えています。`(2)`の`store.dispatch(openRootModal())`がRoot Componentを表示する処理ですが、これはRoot Componentを表示するstateを流す処理になります。

Root Componentは以下のようなコードであり、`visible = true`となることで画面に表示されます。

```js
import React from 'react';

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
                <p>Root modal</p>
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

Root ComponentはReact Componentのため、様々な要素を表示することができますが、基本的には通知のために利用し、ユーザーに入力を促すようなユースケースの場合は[`Interactive Dialog`](https://docs.mattermost.com/developer/interactive-dialogs.html)を使用すべきだと思います。


### [registerPopoverUserAttributesComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPopoverUserAttributesComponent)

`registerPopoverUserAttributesComponent`は、ユーザーアイコンをクリックすると表示されるpopover内に表示されるComponentを登録します。

Mattermost外からユーザーの情報を取得して表示する際に使用できます。所属の名称がコロコロ変わる会社などでLDAPから所属を引いて表示しておくと便利です。

登録したComponentは以下のpropsを受け取れます。[profile_popover.jsx](https://github.com/mattermost/mattermost-webapp/blob/61a4de5e36f994955bacc47c8d56734f9a699a5a/components/profile_popover/profile_popover.jsx#L356)
* `user`: popoverで表示しているユーザーのオブジェクト
  * アクセスできるのは[User](https://github.com/mattermost/mattermost-server/blob/master/model/user.go#L68)モデルで定義されている情報になると思います
  * `User`モデルで`omitempty`となっているものは空の場合除外されるため、値がある場合のみアクセスできます
* `hide`: popoverを閉じる関数
* `status`: popoverで表示しているユーザーのstatus
  * `online`, `offline`などの単なる文字列のようです

ここでは、固定文字列を表示する例を示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/user-attribute.png)

```js:index.js
...

import UserAttributes from './components/user_attributes'

...

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // (1) User Popoverに説明を追加するComponentの登録
		registry.registerPopoverUserAttributesComponent(UserAttributes)
		...
    }
}
```

```js:components/user_attributes.jsx
import React from 'react';
import PropTypes from 'prop-types';

const UserAttributes = ({user, hide, status}) => {
    console.log("user", user.id);
    return (
        <div>
            <p>{'UserId: ' + user.id}</p>
        </div>
    );
}

UserAttributes.propTypes = {
    user: PropTypes.object.isRequired,
    hide: PropTypes.func.isRequired,
    status: PropTypes.object.isRequired,
};

export default UserAttributes;
```


### [registerPopoverUserActionsComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPopoverUserActionsComponent)

`registerPopoverUserActionsComponent`は、ユーザーアイコンをクリックすると表示されるpopover内に表示されるComponentを登録します。`registerPopoverUserAttributesComponent`と似ていますが、表示される位置が異なるだけの違いだと思います。

登録したComponentは、`registerPopoverUserAttributesComponent`と同じく以下のpropsを受け取れます。[profile_popover.jsx](https://github.com/mattermost/mattermost-webapp/blob/61a4de5e36f994955bacc47c8d56734f9a699a5a/components/profile_popover/profile_popover.jsx#L483)
* `user`: popoverで表示しているユーザーのオブジェクト
* `hide`: popoverを閉じる関数
* `status`: popoverで表示しているユーザーのstatus


`registerPopoverUserActionsComponent`で、RootComponentを表示する例は以下のようになります。（処理を一部割愛しているため、プログラム全体は https://github.com/kaakaa/mattermost-plugin-api-sample から確認してください）

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/user-action.gif)

```js:index.js
...
import UserAction from './components/user_action';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // User Popoverにアクションを追加するComponentの登録
		registry.registerPopoverUserActionsComponent(UserAction)
		...
	}
}
```

```js:components/user_action/user_action.jsx
import React from 'react';
import PropTypes from 'prop-types';

const UserActionComponent = ({openRootModal, user, hide, status}) => {
	const onClick = () => {
		openRootModal();
		hide();
	};
    return (
        <div>
            <p>User Action <button onClick={onClick}>Action</button></p>
        </div>
    )
};

UserActionComponent.PropTypes = {
	openRootModal: PropTypes.func.isRequired,
    user: PropTypes.object.isRequired,
    hide: PropTypes.func.isRequired,
    status: PropTypes.object.isRequired,
};

export default UserActionComponent;
```

### [registerLeftSidebarHeaderComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerLeftSidebarHeaderComponent)

`registerLeftSidebarHeaderComponent`は、Mattermost画面の左サイドバーに表示するComponentを登録します。

登録されたComponentで使用できるpropsは特にありません。[sidebar.tsx](https://github.com/mattermost/mattermost-webapp/blob/61a4de5e36f994955bacc47c8d56734f9a699a5a/components/sidebar/sidebar.tsx#L179)

常に表示されている部分のため、職場の室温やCO2濃度などを表示するなどの利用方法があるかと思います。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/left-sidebar-header.png)

```js:index.js
...
import LeftSidebarHeader from './components/left_sidebar_header';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // User Popoverにアクションを追加するComponentの登録
		registry.registerLeftSidebarHeaderComponent(LeftSidebarHeader)
		...
	}
}
```

```js:components/left_sidebar_header.jsx
import React from 'react';

const LeftSidebarHeaderComponent = () => {
	const style = {
		color: 'white',
	};
    return (
        <div style={style}>
            <p>left sidear header</p>
        </div>
    );
}

export default LeftSidebarHeader;
```


### [registerBottomTeamSidebarComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerBottomTeamSidebarComponent)

`registerBottomTeamSidebarComponent`は、チーム選択サイドバーの下部に表示されるComponentを登録します。１つのチームにしか参加していない場合、チーム選択サイドバー自体が表示されないため、ここで登録したComponentも表示されません。

登録されたComponentで使用できるpropsは特にありません。[legacy_team_sidebar_controller.tsx](https://github.com/mattermost/mattermost-webapp/blob/61a4de5e36f994955bacc47c8d56734f9a699a5a/components/legacy_team_sidebar/legacy_team_sidebar_controller.tsx#L283)

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/bottom-team-sidebar.png)

```js:index.js
...
import BottomTeamSidebar from './components/bottom_team_sidebar';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // User Popoverにアクションを追加するComponentの登録
		registry.registerLeftSidebarHeaderComponent(BottomTeamSidebar)
		...
	}
}
```

```js:components/bottom_team_sidebar.jsx
import React from 'react';

const BottomTeamSidebarComponent = () => (
    <>
        <a href="https://github.com/">
            <i
                className='icon fa fa-github'
                style={{color: 'white'}}
            />
        </a>
    </>
);

export default BottomTeamSidebarComponent;
```

### [registerLinkTooltipComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerLinkTooltipComponent)

`registerLinkTooltipComponent`は、リンクをhoverした時に表示されるtooltipのComponentを登録します。

登録したComponentは以下のpropsを受け取れます。[link_tooltip.tsx](https://github.com/mattermost/mattermost-webapp/blob/cc4ba4576ad978d2a29e6c51c3f12966d121f0d6/components/link_tooltip/link_tooltip.tsx#L113)
* `href`: hoverしているリンクのURL
* `show`: tooltipが表示されているかどうかのフラグ

以下の例では単にhoverしているリンクのURLを表示しているだけですが、リンク先から情報を取得して表示するようなこともできます。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/link-tooltip.png)

```js:index.js
...
import LinkTooltip from './components/link_tooltip';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // User Popoverにアクションを追加するComponentの登録
		registry.registerLinkTooltipComponent(LinkTooltip)
		...
	}
}
```


```js:components/link_tooltip.jsx
import React from 'react';
import PropTypes from 'prop-types';

const LinkTooltipComponent = ({href}) => {
    return (
        <>
            <p style={{
                backgroundColor: 'white',
                color: 'black',
                padding: '25px',
                border: '1px solid red',
            }}>
                {href}
            </p>
        </>
    )
}

LinkTooltipComponent.propTypes = {
    href: PropTypes.string.isRequired,
};

export default LinkTooltipComponent;
```

### [registerChannelHeaderButtonAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerChannelHeaderButtonAction)

`registerChannelHeaderButtonAction`は、`registerRootComponent`のところでも使用しましたが、チャンネルのヘッダ部分にアクション付きのボタンを登録します。

`registerChannelHeaderButtonAction`は、`(icon, action, dropdownText, tooltipText)`の４つの引数を取ります。
* `icon`: ボタンのアイコンを表すReact Element
* `action`: ボタンがクリックされたときに実行されるアクション
  * ボタンが押されたチャンネルの情報を持つ[`channel`](https://github.com/mattermost/mattermost-server/blob/master/model/channel.go#L36)と、ボタンを押したユーザーの[`channelMember`](https://github.com/mattermost/mattermost-server/blob/master/model/channel_member.go#L44)を引数に取ることができます
* `dropdown_text`: ボタンが複数登録されると一つのアイコンにまとめられ、ドロップダウンメニューから選ぶ形式になりますが、その時のドロップダウンメニューに表示される文字列
* `tooltip_text`: ボタンをhoverした時にtooltipとして表示される文字列

**ボタン形式**

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/channel-header-button-button.png)

**ドロップダウン形式**

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/channel-header-button-dropdown.png)

ここでは、先ほどのRoot Comopnentを開くButtonと、クリックすると投稿を作成するButtonの2つのButtonを表示する例を示します。投稿の作成は[mattermost-redux](https://github.com/mattermost/mattermost-redux)を使って投稿を作成するactionsを作成し、Channel Header Buttonのactionから呼び出しています。

![movie](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/channel-header-button.gif)

```js:index.js
import {openRootModal, createPluginPost} from './actions';
import reducer from './reducer';
import Root from './components/root';

...

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
        // Root Componentの登録
        const rootComponentId = registry.registerRootComponent(Root)

        // Root Componentを呼び出すアクションの登録
        registry.registerChannelHeaderButtonAction(
            () => (<i className='icon fa fa-commenting-o' style={{fontSize: '15px', position: 'relative', top: '-1px'}}/>),
            () => store.dispatch(openRootModal()),
            "Open Root modal","Open Root modal"
		);
		
		...

        // 投稿を作成するチャンネルヘッダボタンを登録
        registry.registerChannelHeaderButtonAction(
            () => (<i className='icon fa fa-commenting-o' style={{fontSize: '15px', position: 'relative', top: '-1px'}}/>),
            (channel, channelMembers) => store.dispatch(createPluginPost(channel.id)),
            "Create Sample Post", "Create Sample Post"
        );

        registry.registerReducer(reducer);
    }
}
```

```js:actions.js
import {createPost} from 'mattermost-redux/actions/posts';
import {getCurrentUserId} from 'mattermost-redux/selectors/entities/users';

...

export function createPluginPost(channelId) {
    return async (dispatch, getState) => {
        const state = getState();
        const userId = getCurrentUserId(state)
        const post = {
            channel_id: channelId,
            user_id: userId,
            message: "Post from webapp plugin",
        }
        return await dispatch(createPost(post));
    }
}
```

### [registerPostTypeComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostTypeComponent)

`registerPostTypeComponent`は、特定のtypeを持つ投稿をレンダリングする時に使用されるComponentを登録します。投稿（`Post`）に独自のtypeを指定する場合、そのtypeは`custom_`で始まる必要があります。

登録したComponentは以下のpropsを受け取れます。
* `post`: 投稿の内容 [post.go](https://github.com/mattermost/mattermost-server/blob/master/model/post.go#L72)
* `theme`: Mattermost画面の配色テーマ [constants.jsx](https://github.com/mattermost/mattermost-webapp/blob/0b47780407e1f36e703df45e7fe694c7cf750244/utils/constants.jsx#L1047)

typeに`custom_sample_post`を持つ投稿が作成された場合、背景色が赤の投稿をレンダリングする例を以下に示します。今回は、Incoming Webhookで投稿を作成する際のパラメータとして`custom_sample_post`の独自typeを指定しています。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"type": "custom_sample_post", "text": "Sample Message", "props": {"card": "## Sample\n* foo\n* bar"}}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-type.png)

```js:index.js
...
import CustomPost from './components/custom_post';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // type: custom_sample_post を持つ投稿をレンダリングするComponentの登録
		registry.registerPostTypeComponent('custom_sample_post', CustomPost);
	}
}
...
```

```js:components/custom_post.jsx
import React from 'react';
import PropTypes from 'prop-types';

const {formatText, messageHtmlToComponent} = window.PostUtils;

const CustomPostComponent = ({post, theme}) => {
    const formattedText = messageHtmlToComponent(formatText(post.message));

    console.log('theme', theme);
    return (
        <div style={{backgroundColor: '#ffcccc'}}>
            {formattedText}
            <pre>
                {JSON.stringify(post.props, null, 4)}
            </pre>
        </div>
    )
}

CustomPostComponent.propTypes = {
    post: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};

export default CustomPostComponent;
```

独自のtypeを指定した投稿を作成した後にPluginを無効にするなど、typeに対応したComponentが登録されていない場合は、通常の投稿と同じようにレンダリングされます。

### [registerPostCardTypeComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostCardTypeComponent)

`registerPostCardTypeComponent`は、特定のtypeを持つ投稿の`card`要素をレンダリングする時に使用されるComponentを登録します。投稿（`Post`）に独自のtypeを指定する場合、そのtypeは`custom_`で始まる必要があります。

登録したComponentは以下のpropsを受け取れます。
* `post`: 投稿の内容 [post.go](https://github.com/mattermost/mattermost-server/blob/master/model/post.go#L72)
* `theme`: Mattermost画面の配色テーマ [constants.jsx](https://github.com/mattermost/mattermost-webapp/blob/0b47780407e1f36e703df45e7fe694c7cf750244/utils/constants.jsx#L1047)

typeに`custom_sample_card`を持つ投稿が作成された場合、背景色が青のカード要素を持つ投稿をレンダリングする例を以下に示します。`custom_sample_card`という独自typeはIncoming Webhookのパラメータとして指定しています。

```bash
curl \
  -H "Content-Type: application/json" \
  -d '{"type": "custom_sample_card", "text": "Sample Message", "props": {"card": "## Sample\n* foo\n* bar"}}'  \
  http://localhost:8065/hooks/ucw5qjw86jgeum77o1uw8197jr
```

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-card-type.png)

作成する投稿に`props.card`プロパティがない場合、`card`要素を表示するための`i`ボタンが表示されないためカード要素を表示することができません。必ず`props.card`要素を指定する必要があります。

```js:index.js
...
import CustomCard from './components/custom_card';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // type: custom_sample_card を持つ投稿をレンダリングするComponentの登録
        registry.registerPostCardTypeComponent('custom_sample_card', CustomCard);
	}
}
...
```

```js:components/custom_card.jsx
import React from 'react';
import PropTypes from 'prop-types';

const {formatText, messageHtmlToComponent} = window.PostUtils;

const CustomCardComponent = ({post, theme}) => {
    const formattedText = messageHtmlToComponent(formatText(post.message));

    console.log('theme', theme);
    return (
        <div style={{backgroundColor: '#ccccff'}}>
            {formattedText}
            <pre>
                {JSON.stringify(post.props, null, 4)}
            </pre>
        </div>
    )
}

CustomCardComponent.propTypes = {
    post: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};

export default CustomCardComponent;
```

独自のtypeを指定した投稿を作成した後にPluginを無効にするなど、typeに対応したComponentが登録されていない場合は、通常の投稿と同じようにレンダリングされます。

### [registerPostWillRenderEmbedComponent ](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostWillRenderEmbedComponent )

`registerPostWillRenderEmbedComponent`は、投稿内で最初に出現するURLの内容をプレビュー表示する部分のコンポーネントを登録します。OpenGraphの情報を表示するアレです。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-will-render-embed-sample.png) 

`registerPostWillRenderEmbedComponent`は3つの引数を取ります。
* `match`: ここで指定したコンポーネントでプレビュー要素をレンダリングするかどうかを決定するための関数。プレビュー表示対象のURLや種別を使い、`true`を返すとここで指定したコンポーネントを使用する。
* `component`: プレビュー要素をレンダリングするコンポーネント
* `toggleable`: プレビュー要素を折り畳み可能にするかどうかのフラグ

また、登録した`Component`は以下のpropsを受け取れます。
* `embed`: 投稿の内容 [post.go](https://github.com/mattermost/mattermost-server/blob/master/model/post.go#L72)
  * `type`: `url`から取得できるプレビューの種別。`opengraph`や`image`などの値が入る
  * `url`: 投稿ないで最初に出現するURL
  * `data`: `url`から取得できるプレビュー内容に関する情報が格納されるらしいが、試している中では値が取得できなかった

`https://github.com/mattermost`で始まるURLが指定された場合に、青い背景色のプレビューをレンダリングする例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-will-render-embed.png) 

```js:index.js
...
import CustomEmbed from './components/custom_embed';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 投稿に含まれるURLのプレビューをレンダリングするComponentの登録
        registry.registerPostWillRenderEmbedComponent(
            (embed) => embed.url && embed.url.startsWith(`https://github.com/mattermost/`),
            CustomEmbed,
            true
        );
    }
}
...
```

```js:components/custom_embed.jsx
import React from 'react';
import PropTypes from 'prop-types';

const CustomEmbedComponent = ({embed}) => {
    const {type, url, data} = embed;
    return (
        <div style={{backgroundColor: '#ccffcc'}}>
            <p>{type}</p>
            <p>{url}</p>
            <pre>
                {JSON.stringify(data, null, 4)}
            </pre>
        </div>
    )
}

CustomEmbedComponent.propTypes = {
    embed: PropTypes.shape ({
        type: PropTypes.string.isRequired,
        url: PropTypes.string.isRequired,
        data: PropTypes.object.isRequired,
    })
};

export default CustomEmbedComponent;
```

以下のリポジトリの`oEmbed-plugin`ブランチに、Mattermost公式チームによる`registerPostWillRenderEmbedComponent`を使用したプラグインの実装が残っています。（作りかけで止まっているようですが）
https://github.com/mattermost/mattermost-plugin-oembed/tree/oEmbed-plugin

### [registerMainMenuAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerMainMenuAction)

`registerMainMenuAction`は、Mattermostのメインメニュー部分に独自のメニューを追加します。

`registerPostWillRenderEmbedComponent`は3つの引数を取ります。
* `text`: メインメニューに表示される文字列かReact Elementを指定します
* `action`: メインメニューが選択された際に実行されるアクションを登録します
* `mobileIcon`: モバイルアプリで表示されるアイコンを登録します（が、モバイルアプリだとメインメニュー表示されない...?）

モーダルを開くアクションが登録されたメインメニューを追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/main-menu.gif)


```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // メインメニューを追加する
        registry.registerMainMenuAction(
            'Sample Main Menu',
            () => store.dispatch(openRootModal()),
            () => (<i className='icon fa fa-plug' style={{fontSize: '15px', position: 'relative', top: '-1px'}}/>)
        );
    }
}
...
```

### [registerChannelHeaderMenuAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerChannelHeaderMenuAction)

`registerChannelHeaderMenuAction`は、チャンネル名の部分をクリックした時に表示されるメニューに独自のメニューを追加します。

`registerPostWillRenderEmbedComponent`は2つの引数を取ります。
* `text`: メニューに表示される文字列かReact Element
* `action`: メニューが選択された際に実行されるアクション。実行されたチャンネルのチャンネルIDが渡されます。

モーダルを開くアクションが登録されたチャンネルヘッダーメニューを追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/channel-header-menu-action.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // チャンネル名をクリックした際に表示されるメニューに独自メニューを追加する
        registry.registerChannelHeaderMenuAction(
            <i className='icon fa fa-plug' style={{fontSize: '15px', position: 'relative', top: '-1px'}}>{'Sample Menu'}</i>,
            (chnnelId) => store.dispatch(openRootModal())
        );
    }
}
...
```

### [registerPostDropdownMenuAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostDropdownMenuAction)
`registerPostDropdownMenuAction`は、投稿に対するドロップダウンメニューに独自のメニューを追加します。

`registerPostDropdownMenuAction`は3つの引数を取ります。

* `text`: メニューに表示される文字列かReact Element
* `action`: メニューが選択された際に実行されるアクション。メニューが実行された投稿のPostIDが渡されます。
* `filter`: メニューを表示するかどうかを決定する関数。実行された投稿のPostIDが渡されます。

モーダルを開くアクションが登録されたドロップダウンメニューを追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-dropdown-menu-action.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 投稿に対するドロップダウンメニューに独自メニューを追加する
        registry.registerPostDropdownMenuAction(
            <i className='icon fa fa-plug' style={{fontSize: '15px'}}>{'Sample Post Menu'}</i>,
            (postId) => store.dispatch(openRootModal()),
            (postId) => { return true; }
        );
    }
}
...
```

### [registerPostDropdownSubMenuAction](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostDropdownSubMenuAction)
`registerPostDropdownSubMenuAction`は、投稿に対するドロップダウンメニューにサブメニュー付きの独自のメニューを追加します。

まず、`registerPostDropdownSubMenuAction`は3つの引数を取ります。

* `text`: メニューに表示される文字列かReact Element
* `action`: メニューが選択された際に実行されるアクション。実行された投稿のPostIDが渡されます。(しかし、クリックしても何も実行されないようです)
* `filter`: メニューを表示するかどうかを決定する関数。メニューが実行された投稿のPostIDが渡されます。

`registerPostDropdownSubMenuAction`は、返り値として登録したアクションのIDと、サブメニューを登録するための関数を返します。ここで返される関数を使ってサブメニューを登録していきますが、この関数は2つの引数を取ります。

* `text`: メニューに表示される文字列かReact Element
* `action`: メニューが選択された際に実行されるアクション。メニューが実行された投稿のPostIDが渡されます。

（サブメニューを登録する関数の第3引数として`filter`となる関数を渡しても、特に効果はないようです）

モーダルを開くアクションが登録されたドロップダウンサブメニューを追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-dropdown-submenu.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 投稿に対するサブメニュー付きのドロップダウンメニューを追加する
        const {id, rootRegisterMenuItem} = registry.registerPostDropdownSubMenuAction(
            <i className='icon fa fa-plug' style={{fontSize: '15px'}}>{'Sample SubMenu'}</i>,
            (postId) => store.dispatch(openRootModal()),  // 実行されない
            (postId) => { return true; }
        );
        rootRegisterMenuItem(
            <i className='icon fa fa-plug' style={{fontSize: '15px'}}>{'SubMenu 1'}</i>,
            (postId) => store.dispatch(openRootModal()),
        );
        rootRegisterMenuItem(
            <i className='icon fa fa-plug' style={{fontSize: '15px'}}>{'SubMenu 2'}</i>,
            (postId) => store.dispatch(openRootModal())
        );
    }
}
...
```


### [registerPostDropdownMenuComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerPostDropdownMenuComponent)

`registerPostDropdownMenuComponent`は、投稿に対するドロップダウンメニューに独自のコンポーネントを追加します。

登録したComponentは以下のpropsを受け取れます。
* `postId`: アクションが実行された投稿のPostID
* `theme`: Mattermost画面の配色テーマ [constants.jsx](https://github.com/mattermost/mattermost-webapp/blob/0b47780407e1f36e703df45e7fe694c7cf750244/utils/constants.jsx#L1047)
* `dispatch`: actionをdispatchするための関数

モーダルを開くアクションが登録されたドロップダウンメニューを追加する例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/post-dropdown-component.gif)

```js:index.js
...
import CustomPostDropdown from './components/custom_post_dropdown';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 投稿に対するドロップダウンメニューにコンポーネントを登録する
        registry.registerPostDropdownMenuComponent(CustomPostDropdown);
		...
	}
}
```

```js:components/custom_post_dropdown.jsx
import React from 'react';
import PropTypes from 'prop-types';

import {openRootModal} from 'actions';

const CustomPostDropdownComponent = ({postId, theme, dispatch}) => {
    return (
        <div
            style={{backgroundColor: '#ffcccc'}}
            onClick={() => dispatch(openRootModal())}
        >
            {postId}
        </div>
    )
}

CustomPostDropdownComponent.propTypes = {
    postId: PropTypes.string.isRequired,
    theme: PropTypes.object.isRequired,
    dispatch: PropTypes.func.isRequired,
};

export default CustomPostDropdownComponent;
```

### [registerFileUploadMethod](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerFileUploadMethod)
`registerFileUploadMethod`は、ファイルアップロードメニューに独自のメニューを追加します。

`registerFileUploadMethod`は3つの引数を取ります。(公式ドキュメントは引数の順序が違っているので注意)

* `icon`: JSX形式のメニューに表示されるアイコン
* `action`: メニューが選択された際に実行されるアクション。ファイルをアップロードするための関数が渡されます。
* `text`: メニューに表示されるテキスト

メニューが選択されると、テキストファイルが2つアップロードされる例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/file-upload-method.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // ファイルアップロードメニューを追加する
        registry.registerFileUploadMethod(
            <i className='icon fa fa-pencil-square-o' style={{fontSize: '15px'}}/>,
            (upload) => upload([
                new File(["test1"], "sample1.txt"),
                new File(["test2"], 'sample2.txt')
            ]),
            'Sample File Upload'
        );
    }
}
...
```

### [registerFilesWillUploadHook](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerFilesWillUploadHook)

`registerFilesWillUploadHook`は、ファイルアップロード時に実行される関数を登録できます。

`registerFilesWillUploadHook`は2つの引数を取る関数を引数に取ります。

* `files`: アップロードしようとしているファイルの配列
* `upload`: ファイルをアップロードするための関数
  * この関数を使わなくとも、アップロードしたいファイルを返却値として指定することでファイルをアップロードできます

返り値として以下のフィールドを持つオブジェクトを返却する必要があります。
* `message`: 処理結果として画面に表示されるメッセージを指定します
* `files`: 処理結果としてアップロードするファイルを配列として指定します。`null`を指定するとアップロードをrejectします。

2つ以上のファイルを同時にアップロードしようとするとアップロードをrejectする例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/files-will-upload-hook.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...

        // ファイルアップロード時の処理を登録する
        registry.registerFilesWillUploadHook((files, upload) => {
            let msg = '';
            if (files.length >= 2 ) {
                files = null;
                msg = 'Must upload one by one.';
            }
            return {
                message: msg,
                files: files,
            };
        });
    }
}
...
```

### [unregisterComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#unregisterComponent)

`unregisterComponent`は、プラグインとして登録したComponentやHookを登録から除外します。


`registerLeftSidebarHeaderComponent`によって登録したコンポーネントを登録から除外するメインメニューの例を以下に示します。

![](https://blog.kaakaa.dev/images/posts/advent-calendar-2020/day20/unregister-component.gif)

```js:index.js
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // 左サイドバーの上部に表示されるComponentの登録
        const leftSidebarHeaderComponentId = registry.registerLeftSidebarHeaderComponent(LeftSidebarHeader);
        ...
        // プラグインによって登録されたコンポーネントを登録から除外する
        registry.registerMainMenuAction(
            'Unregister LeftSideberHeader',
            () => registry.unregisterComponent(leftSidebarHeaderComponentId),
            () => (<i className='icon fa fa-plug' style={{fontSize: '15px', position: 'relative', top: '-1px'}}/>)
        );
    }
}
...
```

### [unregisterPostTypeComponent](https://developers.mattermost.com/extend/plugins/webapp/reference/#unregisterPostTypeComponent)

`unregisterPostTypeComponent`は、`unregisterComponent`と同様、独自の`PostType`に対して設定したComponentを登録から除外します。

例は省略します。

### [registerReducer](https://developers.mattermost.com/extend/plugins/webapp/reference/#registerReducer)

`registerReducer`は、プラグイン内で使用するReducerを登録します。

ここまでに紹介してきた例では、モーダルを開くアクションが実行された場合などに`registerReducer`によって登録されたreducerを利用しています。

```js:index.js
...
import reducer from './reducer';
...
export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
		...
        // Reducerを登録する
        registry.registerReducer(reducer);
    }
}
```

```js:reducer.js
import {combineReducers} from 'redux';

import {OPEN_ROOT_MODAL, CLOSE_ROOT_MODAL} from './action_types';

const rootModalVisible = (state = false, action) => {
    switch (action.type) {
    case OPEN_ROOT_MODAL:
        return true;
    case CLOSE_ROOT_MODAL:
        return false;
    default:
        return state;
    }
};

export default combineReducers({
    rootModalVisible,
});
```

```js:action_types.js
import {id as pluginId} from './manifest';

export const OPEN_ROOT_MODAL = pluginId + '_open_root_modal';
export const CLOSE_ROOT_MODAL = pluginId + '_close_root_modal';
```

## さいごに

本日は、Mattermost Pluginの**Webapp**サイドの実装について紹介しました。
明日は、残りの**Webapp**サイドのAPIを紹介していきます。
