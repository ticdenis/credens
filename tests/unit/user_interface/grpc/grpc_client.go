package main

import (
	"context"
	pb "credens/src/user_interface/grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	host := "localhost:4040"

	fmt.Printf("Listening insecure grpc server at %s port...\n", host)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewAPIServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	fmt.Println("Calling SayHello function on grpc server...")
	req := &pb.SayHelloRequest{To: "World"}
	res, err := client.SayHello(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result from SayHello function called from grpc server: %v", res.Result)
}
