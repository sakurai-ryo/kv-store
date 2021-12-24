# kv-store
インメモリのkey-valueストア
書籍Cloud Native Go3章を参孝に実装
- [x] CRUDの実装
- [ ] バリデーション
    - 特にリクエストからメモリにデータを展開しているのでBodyのサイズは要確認
- [ ] トランザクションログをファイルに書き出す
  - 起動時にデータを復元出来るようにしたい
- [ ] ロックを取らない方法を考える
  - mapを使用している限りはGoroutineセーフにならないのでデータ構造から考える

# Method
```shell
# Put
$ curl -X PUT -d 'Hello, key-value store!' -v http://localhost:8080/key1

# Get
$ curl -v http://localhost:8080/v1/key1

# Delete
$ curl -X DELETE -v http://localhost:8080/key1
```
