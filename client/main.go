package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/14jasimmtp/grpc-go-sample/proto"

)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect ,", err)
	}
	defer conn.Close()

	client:= pb.NewGreetServiceClient(conn)

	names:=&pb.NamesList{
		Names: []string{"Akhil","Alice","Bob"},
	}
	callSayHello(client)
	callSayHelloClientStream(client,names)
	callSayHelloBidirectionalStream(client,names)
	callSayHelloServerStream(client,names)

}
