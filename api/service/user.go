package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
	"strconv"
)

type IUserService interface {
	Register(userAuth *entity.UserAuth) error
	Login(userAuth *entity.UserAuth) error
}

type userServiceImpl struct{}

var UserService IUserService = &userServiceImpl{}

func (svc *userServiceImpl) Register(user *entity.UserAuth) error {
	userId, err := strconv.Atoi(user.UserId)
	if err != nil {
		return fmt.Errorf("user id error, err: %v", err)
	}
	request := &userpb.RegisterRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    int64(userId),
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
	userId, err := strconv.Atoi(userAuth.UserId)
	if err != nil {
		return fmt.Errorf("user id error, err: %v", err)
	}
	request := &userpb.LoginRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    int64(userId),
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
