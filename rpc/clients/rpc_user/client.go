package rpc_user

import (
	"github.com/bqxtt/book_online/rpc/model/userpb"
	"google.golang.org/grpc"
	"log"
)

var UserServiceClient userpb.UserServiceClient

const (
	//address = "localhost:50001"
	address = "101.200.155.166:30001"
)

func Init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("user rpc service connect...")
	UserServiceClient = userpb.NewUserServiceClient(conn)
}
