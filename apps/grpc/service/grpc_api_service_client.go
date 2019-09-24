package service

import (
	"context"
	pb "credens/apps/grpc/proto"
	"google.golang.org/grpc"
	"time"
)

func NewGRPCAPIServiceClient(host string) pb.APIServiceClient {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	return pb.NewAPIServiceClient(conn)
}

func MakeContextWithTimeout(duration time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), duration*time.Millisecond)
	defer cancel()
	return ctx
}
