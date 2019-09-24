package service

import (
	"context"
	pb "credens/apps/grpc/proto"
	"fmt"
)

type GRPCAPIService struct{}

func (_ *GRPCAPIService) SayHello(ctx context.Context, request *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Result: fmt.Sprintf("Hello %s!", request.GetTo())}, nil
}
