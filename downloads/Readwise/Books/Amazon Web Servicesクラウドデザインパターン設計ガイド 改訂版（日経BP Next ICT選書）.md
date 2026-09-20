# Amazon Web Servicesクラウドデザインパターン設計ガイド 改訂版（日経BP Next ICT選書）

![](https://m.media-amazon.com/images/I/91TnadJ0lWL._SY160.jpg)

### Metadata

- Author: 玉川 憲, 片山暁雄, and 瀬戸島 敏宏
- Full Title: Amazon Web Servicesクラウドデザインパターン設計ガイド 改訂版（日経BP Next ICT選書）
- Category: #books

### Highlights

- 論理的に整合性を取った状態でスナップショットを作成する必要がある。 ([Location 305](https://readwise.io/to_kindle?action=open&asin=B0126HZGP8&location=305))
- vmstatやリソースモニター、CloudWatchなどでリソース利用量を把握し、スペック不足（または過剰）な場合は、一旦EC2インスタンスを停止し、AWS Management ConsoleのChange Instance Typeメニューからインスタンスタイプを変更後、再度起動する。 ([Location 372](https://readwise.io/to_kindle?action=open&asin=B0126HZGP8&location=372))
- サーバースペックを変更するときは、EC2インスタンスを一旦停止する必要がある。その際、数十秒～数分（サーバーのディスク量や設定によって変わる）のオフライン状態が発生する。 ([Location 380](https://readwise.io/to_kindle?action=open&asin=B0126HZGP8&location=380))
- HTTPセッション管理やSSL処理などをELBに任せるのか、配下のサーバーで処理するのかを考慮する。 ([Location 434](https://readwise.io/to_kindle?action=open&asin=B0126HZGP8&location=434))
- ディスクは年々容量単価が下がる傾向にあるので、必要になったときに確保するというのは費用面のメリットが大きい。 ([Location 463](https://readwise.io/to_kindle?action=open&asin=B0126HZGP8&location=463))
