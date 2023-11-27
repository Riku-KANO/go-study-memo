package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hello "github.com/Riku-KANO/grpc-gateway/tutorial/proto/hello"
)

type server struct {
	hello.UnimplementedSayServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Hello(ctx context.Context, in *hello.Request) (*hello.Response, error) {
	return &hello.Response{Msg: in.Name + " world\n"}, nil
}

func run() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	hello.RegisterSayServer(s, &server{})

	log.Println("Serving gRPC on 0.0.0.0:9090")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		":9090",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return err
	}
	
	mux := runtime.NewServeMux()
	err = hello.RegisterSayHandler(context.Background(), mux, conn)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)


	return gwServer.ListenAndServe()

}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}