package rpc_rent

import (
	"github.com/bqxtt/book_online/rpc/clients/rpc_rent/rentpb"
	"google.golang.org/grpc"
	"log"
)

var RentServiceClient rentpb.RentServiceClient

const (
	address = "localhost:50003"
)

func Init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("rent service connect...")
	RentServiceClient = rentpb.NewRentServiceClient(conn)
}
