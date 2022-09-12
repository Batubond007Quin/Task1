package main

import (
	"cmd/service/main.go/pkg/api/protobuf"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatal("Failed to listen port 9000: %v", err)
	}

	s := protobuf.Server{}
	grpcServer := grpc.NewServer()

	protobuf.RegisterGatewayServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve grpc server over port 9000: %v", err)
	}
}
