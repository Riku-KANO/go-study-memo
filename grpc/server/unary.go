package main

import (
	"context"
	
	pb "github.com/Riku-KANO/go-study-memo/grpc/proto"
)
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}