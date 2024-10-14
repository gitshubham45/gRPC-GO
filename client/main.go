package main

import (
	"log"

	pb "github.com/gitshubham45/gRPC-GO/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connnect : %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Akhil", "John", "Alice", "Bob", "Peter"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
