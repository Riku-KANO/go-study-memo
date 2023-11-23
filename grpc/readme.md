# simple gRPC
client/main.goのコメントを外してstreamingを指定。
- unary RPC
- client streaming RPC
- server streaming RPC
- bidirectional streaming RPC 

## コマンド
- サーバー側  
`server` ディレクトリ下で以下のコマンドをたたいてサーバーを起動。
```
go run *.go
```

- クライアント側  
`client` ディレクトリ下で以下のコマンドをたたく。簡単なメッセージを送る。
```
go run *.go
```
## protoファイルのビルド
著者は`~/.local/bin`に`protoc`を入れていたのでそこにあるバイナリを参照してビルド。（`/usr/local/bin`配下にも`protoc`があり、そのバージョンが低いため、わざわざバイナリを探して実行する。）実行すると`proto`ディレクトリ下に`greet_grpc.pb.go`と`greet.pb.go`が作成される。
```bash
~/.local/bin/protoc --go_out=. --go-proc_out=. proto/greet.proto
```