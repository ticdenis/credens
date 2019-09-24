package service

import (
	pb "credens/apps/grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func NewGRPCAPIServiceServer(port int) (*grpc.Server, net.Listener) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterAPIServiceServer(server, &GRPCAPIService{})

	return server, listener
}
