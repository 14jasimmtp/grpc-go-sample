package main

import (
	"log"
	"net"

	pb "github.com/14jasimmtp/grpc-go-sample/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("failed to start the server ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
