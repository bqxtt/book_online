package main

import (
	"fmt"
	"github.com/bqxtt/book_online/rent/pkg/common"
	"github.com/bqxtt/book_online/rent/pkg/handler"
	"github.com/bqxtt/book_online/rent/pkg/sdk/rentpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", common.DefaultUserGRPCHandlerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := handler.NewHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	rentpb.RegisterRentServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
