package main

import (
	"context"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &userpb.RegisterRequest{
		User: &userpb.User{
			Name:   "bqx",
			UserId: 12,
		},
		UserAuth: &userpb.UserAuth{
			PwdDigest: "123456",
			UserId:    12,
		},
	})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("register message: %v", r.GetUserId())
}
