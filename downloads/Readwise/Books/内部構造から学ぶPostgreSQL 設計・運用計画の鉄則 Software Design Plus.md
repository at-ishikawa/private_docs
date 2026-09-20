# 内部構造から学ぶPostgreSQL 設計・運用計画の鉄則 Software Design Plus

![](https://m.media-amazon.com/images/I/81ecktsrmVL._SY160.jpg)

### Metadata

- Author: 勝俣智成, 佐伯昌樹, and 原田登志
- Full Title: 内部構造から学ぶPostgreSQL 設計・運用計画の鉄則 Software Design Plus
- Category: #books

### Highlights

- 異なるメジャーバージョン間では、データベースを構成するファイルに互換性がない ([Location 408](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=408))
- 設定項目と設定値は、SHOW文で確認することができます（図3-2）。 ([Location 593](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=593))
- ファイル名のhbaはhost-based authentication（ホストベース認証）を意味します。 ([Location 627](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=627))
- リードアンコミッティド（READ UNCOMMITTED）レベルを指定した場合もリードコミッティドと同じ挙動となるので、事実上PostgreSQLではダーティリードは発生しません。 ([Location 922](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=922))
- タイムゾーンを含む場合は、time with time zoneよりtimestamp with time zoneを使用するほうが望ましいです。これはtime型のみで日付情報を持たないために、夏時間への対応が不十分となることと、データ格納領域の観点からもtime with time zoneよりtimestamp with time zoneのほうが有利だからです。 ([Location 1052](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1052))
- 基本的には8,192バイトのページと呼ばれる固定長領域が連続して配置されたものになります。 ([Location 1162](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1162))
- 例えば、検索時の条件として複数列インデックス作成時に最初に指定した列が含まれない場合には、作成した複数列インデックスは使用されなません（図6-8）。 ([Location 1406](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1406))
- 部分インデックスを使うことで、出現頻度が低いデータのみをインデックス化の対象にできます。 ([Location 1425](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1425))
- 部分インデックスを構築するためには、CREATE INDEX文で指定する列にWHERE条件を付加します。このとき、WHERE条件には出現頻度の高い値を除外する条件を指定するとよいでしょう。 ([Location 1428](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1428))
- エンコーディングはサーバ側とクライアント側で設定できますが、それぞれのエンコーディングが異なる場合、エンコーディングの変換が発生するため、性能上のロスが発生します。 ([Location 1458](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1458))
- これらの項目については、基本的にはOSやストレージに付属するツールによる監視を行います。具体的には、「sar」「iostat」「vmstat」「top」「netstat」コマンドなどを用いて、定期的に情報を収集します。 ([Location 1593](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1593))
- Linuxにおいて、制限を受ける可能性のあるカーネルパラメータは、共有メモリセグメントの最大容量を制限する「shmmax」と使用可能な共有メモリの総量を制限する「shmall」です（表9-2）。 ([Location 1721](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1721))
- また、sysctlで設定した値はサーバを再起動するとデフォルト値に戻ってしまうため、/etc/sysctl.confファイルに設定値を保存することが強く推奨されます。 ([Location 1732](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1732))
- PostgreSQLのメモリ設定の中でも特に重要なパラメータは「shared_buffers」で、PostgreSQLが共有バッファのために確保する共有メモリのサイズを設定します。初期値は128MBと比較的小さな値が設定されているため、ほとんどの場合で設定変更を行うことが推奨されます。目安はメモリを1GB以上搭載したサーバであれば、その25%程度を設定するとよいでしょう。 ([Location 1740](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1740))
- クライアント接続情報とプリペアドトランザクションを管理する領域の共有メモリ概算 ([Location 1759](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1759))
- PostgreSQLでは、実際のI/O要求を行うプロセスが、データ書き込みプロセス（bgwriter）やWAL書き込みプロセス（wal writer）といった少数のプロセスがI/O要求の大半を占め、データアクセスもランダムアクセスが多いために、deadlineに設定することが推奨されてきました。 ([Location 1795](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1795))
- WALレコードを出力しないHashインデックスやunloggedテーブルなどは、ストリーミングレプリケーションではレプリケーションできません。 ([Location 1889](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1889))
- 「同期」といってもスタンバイの古いデータが読まれる可能性があるので注意が必要です。なぜなら、WALの適用は常に非同期であるためです。 ([Location 1940](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1940))
- pg_ctl promoteを実行することで「昇格」の処理がなされます。 ([Location 1987](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=1987))
- シーク時間の差は、ディスクアクセスがランダムになるほど顕著となります。PostgreSQLでは、インデックススキャンの際にはランダムアクセスが発生します。 ([Location 3318](https://readwise.io/to_kindle?action=open&asin=B00NM7VROQ&location=3318))
