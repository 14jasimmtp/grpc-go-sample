package main

import (
	"context"
	pb "github.com/14jasimmtp/grpc-go-sample/proto"

)

func (s *helloServer) SayHello(ctx context.Context,req *pb.NoParam) (*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "hello",
	},nil
}