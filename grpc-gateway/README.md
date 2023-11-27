# gRPC gateway
- [grpc-gateway repositry](https://github.com/grpc-ecosystem/grpc-gateway)  

protocの拡張プラグインの一つ。.protoファイルからゲートウェイ(リバースプロキシ)を作成することができる。外部からのRESTful JSONリクエストをゲートウェイにてgRPCリクエストへと変換する。そのリクエストをそれぞれのサービスへ転送する。
## installation
### grpc-gatewayのインストール
```bash
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### bufのインストール
protocを使う代わりに`buf`コマンドを使う。protocでもコンパイルはできるが、それを利用する際はgoogleapisなどの依存ライブラリをプロジェクトへダウンロードする必要がある。bufを使えばbuf.yamlに依存ライブラリを記入するだけで良い。  
[インストールページ](https://github.com/bufbuild/buf)

## サンプル
以下のコマンドでゲートウェイとサーバーの起動をする。
```bash
go run tutorial/main.go
```
立ち上がったゲートウェイのアドレス(localhost:8080)の`/greeter/hello`エンドポイントに対して`curl`コマンドをたたいてメッセージを送る。
```
curl -H "content-type: application/json" -d '{"name": "Kano"}' http://localhost:8080/greeter/hello
```