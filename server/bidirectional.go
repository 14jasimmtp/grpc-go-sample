package main

import (
	"log"

	pb "github.com/14jasimmtp/grpc-go-sample/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error{
	for {
		req,err:=stream.Recv()
		if err != nil{
			return nil
		}
		if err != nil{
			return err
		}
		log.Printf("Got request with name : %v",req.Name)
		res:=&pb.HelloResponse{
			Message: "Hello"+req.Name,
		}
		if err :=stream.Send(res);err != nil{
			return err
		}
	}
}