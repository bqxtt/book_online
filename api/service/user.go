package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
)

type IUserService interface {
	Register(userAuth *entity.UserAuth) error
	Login(userAuth *entity.UserAuth) error
}

type userServiceImpl struct{}

var UserService IUserService = &userServiceImpl{}

func (svc *userServiceImpl) Register(user *entity.UserAuth) error {
	request := &userpb.RegisterRequest{
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
		return fmt.Errorf("register response is nil")
	}
	if resp.Reply.Status != 1 {
		return fmt.Errorf("%v", resp.Reply.Message)
	}
	return nil
}

func (svc *userServiceImpl) Login(userAuth *entity.UserAuth) error {
	request := &userpb.LoginRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    userAuth.UserId,
			PwdDigest: userAuth.Password,
		},
	}
	resp, err := rpc_user.UserServiceClient.Login(context.Background(), request)
	if err != nil {
		return fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("login response is nil")
	}
	if resp.Reply.Status != 1 {
		return fmt.Errorf("%v", resp.Reply.Message)
	}
	return nil
}
