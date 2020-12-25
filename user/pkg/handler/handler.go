package handler

import (
	"context"
	"github.com/bqxtt/book_online/user/pkg/handler/adapter"
	"github.com/bqxtt/book_online/user/pkg/sdk/base"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"github.com/bqxtt/book_online/user/pkg/service"
)

var _ userpb.UserServiceServer = &UserHandler{}

type UserHandler struct {
	userpb.UnimplementedUserServiceServer

	userService *service.UserService
}

func (h *UserHandler) Register(ctx context.Context, request *userpb.RegisterRequest) (*userpb.RegisterReply, error) {
	r := adapter.RpcUser{User: request.User, UserAuth: request.UserAuth}
	user, userAuth := r.ToModelUser()
	err := h.userService.CreateUser(user, userAuth)
	if err != nil {
		return nil, err
	}
	return &userpb.RegisterReply{
		Reply: &base.BaseReply{
			Status:  base.REPLY_STATUS_SUCCESS,
			Message: "register success",
		},
		UserId: user.UserID,
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginReply, error) {
	panic("implement me")
}

func (h *UserHandler) UpdateInfo(ctx context.Context, request *userpb.UpdateInfoRequest) (*userpb.UpdateInfoReply, error) {
	panic("implement me")
}

func NewHandler() (*UserHandler, error) {
	us, err := service.NewUserService()
	if err != nil {
		return nil, err
	}

	return &UserHandler{
		userService: us,
	}, nil
}
