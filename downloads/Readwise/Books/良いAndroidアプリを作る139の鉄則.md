# 良いAndroidアプリを作る139の鉄則

![](https://m.media-amazon.com/images/I/7117UQZ9CdL._SY160.jpg)

### Metadata

- Author: 木田学, おかじゅん, 渡辺考裕, 奈良進, 荒川祐一郎, 兒島友三郎, 石立宏志, 小林正興, and テックファーム
- Full Title: 良いAndroidアプリを作る139の鉄則
- Category: #books

### Highlights

- 回避策は、launchMode=“singleTop”を設定することです。singleTopを設定すると、2回目以降のIntent呼び出しでインスタンスが生成されなくなります。 ([Location 916](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=916))
- exported属性のデフォルト値はintent-filterタグの有無で変わります。intent-filterを含んでいる場合はexported="true"になり、intent-filterを含んでいない場合はexported="false"になります。 ([Location 1021](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=1021))
- エラー番号を表示し「サポートにこの番号をお伝えください」と指示を強調する ● エラー内容を端末内に保存し、後から確認できるようにする ● エラー内容を端末内に保存し、アプリが正常起動した際に自動的にサーバに送信する エラーを保存するケースのサンプルソースを見てみましょう。 UncaughtExceptionHandler() ([Location 2451](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=2451))
- collapse_keyの指定は最大4つまで ([Location 3613](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=3613))
- リトライ間隔が30秒であれば、次は60秒、120秒、240秒……というように間隔をあけていきます。 ([Location 3628](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=3628))
- 上記を考慮して「アプリ名＋_日付＋_時間＋_リビジョン番号＋.apk」という命名規則を適用するとバージョンが管理しやすくなります。ファイル名が冗長になってしまいますが、apkの管理で混乱を避けることができます。 ([Location 5250](https://readwise.io/to_kindle?action=open&asin=B00KX5FKCA&location=5250))
