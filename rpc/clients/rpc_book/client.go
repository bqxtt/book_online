package rpc_book

import (
	"github.com/bqxtt/book_online/rpc/model/bookpb"
	"google.golang.org/grpc"
	"log"
)

var BookServiceClient bookpb.BookServiceClient

const (
	address = "localhost:50002"
)

func Init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("book rpc service connect...")
	BookServiceClient = bookpb.NewBookServiceClient(conn)
}
