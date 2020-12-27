package main

import (
	"fmt"
	"github.com/bqxtt/book_online/book/pkg/common"
	book "github.com/bqxtt/book_online/book/pkg/handler"
	"github.com/bqxtt/book_online/book/pkg/sdk/bookpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", common.DefaultUserGRPCHandlerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := book.NewBookHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	bookpb.RegisterBookServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
