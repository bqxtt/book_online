package rpc_user

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
	"testing"
)

func TestUserService(t *testing.T) {
	Init()
	resp, err := UserServiceClient.Register(context.Background(), &userpb.RegisterRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    321,
			PwdDigest: "333",
		},
		User: &userpb.User{
			UserId: 321,
			Name:   "bqx",
		},
	})
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(resp)
	}
}
