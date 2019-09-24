package main

import (
	pb "credens/apps/grpc/proto"
	"credens/apps/grpc/service"
	"credens/libs/shared/infrastructure/logging/logrus"
	"fmt"
)

func main() {
	logger := logrus.NewLogger()
	address := fmt.Sprintf("%s:%d", "localhost", 4041)

	logger.Log(fmt.Sprintf("Listening insecure grpc server at %s...", address))
	client := service.NewGRPCAPIServiceClient(address)

	logger.Log("Calling SayHello function on grpc server...")
	req := &pb.SayHelloRequest{To: "World"}
	res, err := client.SayHello(service.MakeContextWithTimeout(20), req)
	if err != nil {
		panic(err)
	}

	logger.Log(fmt.Sprintf("Result from SayHello function called from grpc server: %v", res.Result))
}
