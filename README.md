# echo-docker

## setup
```
git clone git@github.com:You-saku/echo-docker.git
cd echo-docker
cp server/.env.example server/.env
docker-compose up -d --build

docker-compose exec echo sh
go run main.go
```

## Official Doc
 * [echo](https://echo.labstack.com/)

## Todo
* connect front-end.
* manage migration.
* transaction manage rdb.

## 個人メモ

### APIの動作確認
 * postman-schemaディレクトリにpostmanでAPI叩けるようにするファイルを用意
 * importして使ってください(import方法は調べよう)
### ファイル分割
 * [ファイル分割したときに参考になる](https://qiita.com/fetaro/items/31b02b940ce9ec579baf#%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A2%E3%83%BC%E3%83%89%E3%81%A7%E3%81%AE%E5%86%85%E9%83%A8%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%81%AEimport)

### gorm
 * 参考
   * [公式Doc](https://gorm.io/ja_JP/)
   * リレーション
      * [1対多](https://gorm.io/ja_JP/docs/has_many.html) 

### validate
 * 参考
   * [公式Doc](https://echo.labstack.com/guide/request/)
   * [ライブラリの公式Doc](https://github.com/go-playground/validator)
   * https://qiita.com/nanamen/items/c824a2c8f2e1767f90f8
   
 * やってみたい
   * https://codehex.hateblo.jp/entry/echo-context
