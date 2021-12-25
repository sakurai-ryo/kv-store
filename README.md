# kv-store
インメモリのkey-valueストア
書籍Cloud Native Go4章を参孝
- [x] CRUDの実装
- [x] ロックを取らない方法を考える
  - mapを使用している限りはスレッドセーフにならない
  - [スレッドセーフなHashMap実装を使う](https://github.com/cornelk/hashmap)
- [x] ホットリロードを開発環境で使えるようにする
  - `make up`
- [ ] バリデーション
    - 特にリクエストからメモリにデータを展開しているのでBodyのサイズは要確認
- [ ] トランザクションログをファイルに書き出す
  - 起動時にデータを復元出来るようにしたい

# Method
```shell
# Put
$ curl -X PUT -d 'Hello, key-value store!' -v http://localhost:8080/key1

# Get
$ curl -v http://localhost:8080/v1/key1

# Delete
$ curl -X DELETE -v http://localhost:8080/key1
```
