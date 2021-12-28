# kv-store
インメモリのkey-valueストア
[Cloud Native Go5章](https://github.com/cloud-native-go/examples/tree/main/ch05) を参孝

時間があるときにやりたいこと↓
- [x] CRUDの実装
- [x] ロックを取らない方法を考える
  - mapを使用している限りはスレッドセーフにならない
  - [スレッドセーフなHashMap実装を使う](https://github.com/cornelk/hashmap)
- [x] ホットリロードを開発環境で使えるようにする
  - `make up`
- [x] トランザクションログをファイルに書き出す
  - 起動時にデータを復元出来るようにしたい
- [ ] バリデーション
  - [x] 特にリクエストからメモリにデータを展開しているのでBodyのサイズは要確認
      - `MaxBytesReader`を使用する
  - [ ] string以外をBodyに入れるとpanicする
- [ ] ログファイルのClose処理
- [ ] ユニットテスト
- [ ] ログファイルの形式
  - 空行があるとエラーになる
  - ただのテキストなのでディスク効率がよくない
  - ログを削除する仕組みがない
    - 削除済みのキーなどの不要なデータを削除する必要がある

# Method
```shell
# Put
$ curl -X PUT -d 'Hello, key-value store!' -v http://localhost:8080/key1

# Get
$ curl -v http://localhost:8080/v1/key1

# Delete
$ curl -X DELETE -v http://localhost:8080/key1
```
