package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
)

type IUserService interface {
	Register(user *entity.User) error
}

type userServiceImpl struct{}

var UserService IUserService = &userServiceImpl{}

func (svc *userServiceImpl) Register(user *entity.User) error {
	request := &userpb.RegisterRequest{
		User: &userpb.User{
			UserId: user.UserId,
			Name:   user.Name,
		},
		UserAuth: &userpb.UserAuth{
			UserId:    user.UserId,
			PwdDigest: user.Password,
		},
	}
	resp, err := rpc_user.UserServiceClient.Register(context.Background(), request)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("response is empty")
	}
	if resp.Reply.Status != 1 {
		return fmt.Errorf("%v", resp.Reply.Message)
	}
	return nil
}
