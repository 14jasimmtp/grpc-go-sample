package main

import (
	"context"
	"io"
	"log"
	pb "github.com/14jasimmtp/grpc-go-sample/proto"

)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names:%v", err)
	} 

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error while streamiin", err)
		}
		log.Println(message)
	}
	log.Println("streaming finished")
}
