package main

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/common"
	user "github.com/bqxtt/book_online/user/pkg/handler"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", common.DefaultUserGRPCHandlerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := user.NewHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
