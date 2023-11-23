package main

import (
	"log"

	pb "github.com/Riku-KANO/go-study-memo/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Riku"},
	}

	// callSayHello(client) // unary
	// callSayHelloServerStream(client, names) // Server streaming
	// callSayHelloClientStream(client, names) // Clinet streaming
	callSayHelloBidirectionalStream(client, names) // Bidirectional streaming
}