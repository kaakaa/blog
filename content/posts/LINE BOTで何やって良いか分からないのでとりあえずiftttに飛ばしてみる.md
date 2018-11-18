
---
title: "LINE BOTで何やって良いか分からないのでとりあえずiftttに飛ばしてみる"
date: 2016-04-13T22:23:13+09:00
draft: false
toc: true
tags: ["Heroku","ifttt","Line","linebot"]
---

LINE BOTを始めBot界隈が盛り上がってきてますが、BotKit公開されても何すれば良いか分からない人間なので、とりあえずiftttへのパス作っとけば夢は広がるだろうというアレです。

下記の情報の組み合わせでできました。

[LINE BOT をとりあえずタダで Heroku で動かす - Qiita](http://qiita.com/yuya_takeyama/items/0660a59d13e2cd0b2516)
[AWS LambdaとIFTTTでお手軽にpush通知を実現する - Qiita](http://qiita.com/kawahiro311/items/41dca04fe899d4d142d9)

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://dashboard.heroku.com/new?button-url=https%3A%2F%2Fgithub.com%2Fkaakaa%2Fline-ifttt&template=https%3A%2F%2Fgithub.com%2Fkaakaa%2Fline-ifttt)

動かすには下記が必要
* [LINE BOT API のトライアルアカウン](https://business.line.me/services/products/4/introduction)
* [iftttのMakerチャンネルのSecretKey](https://ifttt.com/maker)


* * *

流れとしてはLINE BOTからHerokuへCallbackし、HerokuからiftttのMakerチャンネルへリクエストを飛ばしてます。
[Connect Maker to hundreds of apps - IFTTT](https://ifttt.com/maker)

Makerチャンネルへリクエストを飛ばすと`Congratulations! You've ～` という定型文が返されるので、それをまたLINEへつぶやくようにしています。
（リクエストがiftttへ届いたことを表しているだけで、レシピがエラーとなったかどうかは判別できなそうです）

Makerチャンネルはiftttのレシピ毎に`EventName`を指定できるため、LINEでのメッセージの先頭でその`EventName`を指定できるようにしており、Herokuにアプリ一つ立てておくだけで、iftttの複数のレシピに対応できるはずです。

![line-ifttt-7.png](https://qiita-image-store.s3.amazonaws.com/0/9891/20b58a7c-f787-9d79-6a0c-c3d07da018b5.png)


LINE上で`ifttt-test hogehoge`とつぶやけば、`EventName`が`ifttt-test`のレシピが実行されるはずです。

```rb:app.rb
# Mostly taken from http://qiita.com/masuidrive/items/1042d93740a7a72242a3
# And https://github.com/yuya-takeyama/line-echo

require 'sinatra/base'
require 'json'
require 'rest-client'

class App < Sinatra::Base
  SECRET_KEY = ENV['MAKER_SECRET_KEY'] # Maker channel's secret key
  EVENT_NAME_DEFAULT = ENV['MAKER_EVENT_NAME_DEFAULT'] # Maker channel's default event name
  
  post '/linebot/callback' do
    params = JSON.parse(request.body.read)
    
    RestClient.proxy = ENV['FIXIE_URL'] if ENV['FIXIE_URL']
    params['result'].each do |msg|
      event, message = msg['content']['text'].split(' ', 2)
      response = request_to_ifttt(event, message)
      
      request_to_line([msg['content']['from']], response)
    end

    "OK"
  end

  helpers do
    def request_to_ifttt(event = EVENT_NAME, message)
      # Maker channel's extra data
      request_content = {
        value1: message,
        value2: 'from',
        value3: 'Line Bot'
      }

      endpoint_uri = "https://maker.ifttt.com/trigger/#{event}/with/key/#{SECRET_KEY}"
      content_json = request_content.to_json
      
      response = RestClient.post(endpoint_uri, content_json, {
        'Content-Type' => 'application/json; charset=UTF-8',
        'Content-Length' => content_json.length
      })
    end

    def request_to_line(send_to, message)
      request_content = {
        to: send_to,
        toChannel: 1383378250, # Fixed  value
        eventType: "138311608800106203", # Fixed value
        content: { 
            "contentType": 1,  # Text type message
            "toType": 1,
            "text": message
        }
      }

      endpoint_uri = 'https://trialbot-api.line.me/v1/events'
      content_json = request_content.to_json

      RestClient.post(endpoint_uri, content_json, {
        'Content-Type' => 'application/json; charset=UTF-8',
        'X-Line-ChannelID' => ENV["LINE_CHANNEL_ID"],
        'X-Line-ChannelSecret' => ENV["LINE_CHANNEL_SECRET"],
        'X-Line-Trusted-User-With-ACL' => ENV["LINE_CHANNEL_MID"],
      })
    end  
  end
end
```

* * *

Makerチャンネルのリクエストにはパラメータを３つ（`Value1`, `Value2`, `Value3`）まで指定できます。
![line-ifttt-8.png](https://qiita-image-store.s3.amazonaws.com/0/9891/a3e22ceb-38d6-506b-f897-0ed1c6de704e.png)
`EventName`はそのままの意味、`OccurredAt`はリクエストを受け付けた日時になります。


ifttt上で指定する場合は`Value1`のように先頭を大文字の`V`にしないとエラーになりますが、リクエストパラメータでは先頭小文字の`value1`でないとダメみたいです。

```rb
      # Maker channel's extra data
      request_content = {
        value1: message,
        value2: 'from',
        value3: 'Line Bot'
      }
```

* * *

ソースは下記に置いてます。 (masuidrive さんに怒られたら消す)
[kaakaa/line-ifttt](https://github.com/kaakaa/line-ifttt)


とりあえず作ったもののやりたいことがない。

