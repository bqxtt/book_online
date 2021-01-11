package rpc_rent

import (
	"github.com/bqxtt/book_online/rpc/model/rentpb"
	"google.golang.org/grpc"
	"log"
)

var RentServiceClient rentpb.RentServiceClient

const (
	address = "localhost:50005"
)

func Init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("rent rpc service connect...")
	RentServiceClient = rentpb.NewRentServiceClient(conn)
}
