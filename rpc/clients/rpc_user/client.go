package rpc_user

import (
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
	"google.golang.org/grpc"
	"log"
)

var UserServiceClient userpb.UserServiceClient

const (
	address = "localhost:50001"
)

func Init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserServiceClient = userpb.NewUserServiceClient(conn)
}
