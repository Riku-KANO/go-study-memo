# tutorial

## go-micro
RPCやイベント駆動通信などといった分散システム開発の主要な機能を提供する。簡単に設計が始められて、尚且つすべてのシステムを代替できてしまうような pluggable なアーキテクチャをデフォルトで提供することを思想としている。RPCに主眼を置いている [go-kit](https://github.com/go-kit/kit) と比較すると、RPCだけでなくPubSubなどといったイベント駆動の非同期通信にも力を入れている。GitHubのスター数を比較すると go-micro はスター数が 21k であるのに対して go-kitは 25k もある（2023年11月時点）。個数で比較するとgo-kitがよく利用されているライブラリだと思われるが、コミュニティの活発度やGitHubのソースコードのアップデート状況を鑑みると go-micro が今後人気のあるフレームワークになりうる可能性もある。

## protoのビルド
~/.local/bin/protoc --proto_path=. --micro_out=. --go_out=. proto/greeter.proto

## フレームワーク仕様
以下の `Service` インターフェースを実装していくことによってマイクロサービスのパーツを構成していくことになる。`NewService(opt ...Options)` メソッドでServiceを構築し、引数にはオプションを指定する。そのあとに、 `Init` メソッドを呼び出すことによって `Service` インターフェースの実装が完了する。
```go
// Service is an interface that wraps the lower level libraries
// within go-micro. Its a convenience method for building
// and initializing services.
type Service interface {
	// The service name
	Name() string
	// Init initializes options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}

// Event is used to publish messages to a topic.
type Event interface {
	// Publish publishes a message to the event topic
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}
```

```go 
// client/client.go

// Client is the interface used to make requests to services.
// It supports Request/Response via Transport and Publishing via the Broker.
// It also supports bidirectional streaming of requests.
type Client interface {
	Init(...Option) error
	Options() Options
	NewMessage(topic string, msg interface{}, opts ...MessageOption) Message
	NewRequest(service, endpoint string, req interface{}, reqOpts ...RequestOption) Request
	Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
	Stream(ctx context.Context, req Request, opts ...CallOption) (Stream, error)
	Publish(ctx context.Context, msg Message, opts ...PublishOption) error
	String() string
}
```

```go
// server/server.go

// Server is a simple micro server abstraction.
type Server interface {
	// Initialize options
	Init(...Option) error
	// Retrieve the options
	Options() Options
	// Register a handler
	Handle(Handler) error
	// Create a new handler
	NewHandler(interface{}, ...HandlerOption) Handler
	// Create a new subscriber
	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
	// Register a subscriber
	Subscribe(Subscriber) error
	// Start the server
	Start() error
	// Stop the server
	Stop() error
	// Server implementation
	String() string
}
```