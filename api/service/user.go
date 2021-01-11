package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/adapter"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user"
	"github.com/bqxtt/book_online/rpc/model/base"
	"github.com/bqxtt/book_online/rpc/model/userpb"
	"strconv"
)

type IUserService interface {
	Register(userAuth *entity.UserAuth) error
	Login(userAuth *entity.UserAuth) (*entity.User, error)
	GetUserInfo(userId string) (*entity.User, error)
	UpdateUserInfo(user *entity.User) (*entity.User, error)
	DeleteUser(userId string) error
	ListUsersByPage(page int32, pageSize int32) ([]*entity.User, *entity.PageInfo, error)
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

func (svc *userServiceImpl) Login(userAuth *entity.UserAuth) (*entity.User, error) {
	userId, err := strconv.Atoi(userAuth.UserId)
	if err != nil {
		return nil, fmt.Errorf("user id error, err: %v", err)
	}
	request := &userpb.LoginRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    int64(userId),
			PwdDigest: userAuth.Password,
		},
	}
	resp, err := rpc_user.UserServiceClient.Login(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("login response is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, fmt.Errorf("%v", resp.Reply.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User), nil
}

func (svc *userServiceImpl) GetUserInfo(userId string) (*entity.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("user id error, err: %v", err)
	}
	request := &userpb.GetInfoRequest{
		UserId: int64(id),
	}
	resp, err := rpc_user.UserServiceClient.GetInfo(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("get user info resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, fmt.Errorf("rpc resp message: %v", resp.Reply.Message)
	}
	return adapter.RpcUserToEntityUser(resp.User), nil
}

func (svc *userServiceImpl) UpdateUserInfo(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	request := &userpb.UpdateInfoRequest{
		UserInfo: adapter.EntityUserToRpcUser(user),
	}
	resp, err := rpc_user.UserServiceClient.UpdateInfo(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("update user info resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, fmt.Errorf("rpc resp message: %v", resp.Reply.Message)
	}
	return adapter.RpcUserToEntityUser(resp.ResultUserInfo), nil
}

func (svc *userServiceImpl) DeleteUser(userId string) error {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return fmt.Errorf("user id error, err: %v", err)
	}
	request := &userpb.DeleteUserRequest{
		UserId: int64(id),
	}
	resp, err := rpc_user.UserServiceClient.DeleteUser(context.Background(), request)
	if err != nil {
		return fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("update user info resp is nil")
	}
	if resp.Reply.Status != 1 {
		return fmt.Errorf("rpc resp message: %v", resp.Reply.Message)
	}
	return nil
}
func (svc *userServiceImpl) ListUsersByPage(page int32, pageSize int32) ([]*entity.User, *entity.PageInfo, error) {
	request := &userpb.ListUsersPagedRequest{Page: &base.Page{
		PageNo:   int64(page),
		PageSize: int64(pageSize),
	}}
	resp, err := rpc_user.UserServiceClient.ListUsersPaged(context.Background(), request)
	if err != nil {
		return nil, nil, fmt.Errorf("rpc user service error, err: %v", err)
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list user resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc resp message: %v", resp.Reply.Message)
	}
	var users []*entity.User
	for _, user := range resp.Users {
		users = append(users, adapter.RpcUserToEntityUser(user))
	}
	return users, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}
