package handler

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/common"
	"github.com/bqxtt/book_online/user/pkg/handler/adapter"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/sdk/base"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"github.com/bqxtt/book_online/user/pkg/service"
	"github.com/bqxtt/book_online/user/pkg/utils"
)

const defaultSuccessMessage = "success"

type UserHandler struct {
	userpb.UnimplementedUserServiceServer

	userService *service.UserService
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

func (h *UserHandler) Register(ctx context.Context, request *userpb.RegisterRequest) (*userpb.RegisterReply, error) {
	if request == nil {
		return &userpb.RegisterReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil register request pb message"),
		}, nil
	}

	userAuth, err := adapter.AdaptToModelUserAuth(request.UserAuth)
	if err != nil {
		return &userpb.RegisterReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt pb: %v", err),
		}, nil
	}

	//todo set default user info
	defaultUser := &model.User{
		UserID:    userAuth.UserID,
		Name:      fmt.Sprint(userAuth.UserID),
		AvatarURL: common.DefaultUserAvatarURL,
		Role:      common.UserRole(common.UserRoleNormal),
	}

	_, err = h.userService.CreateUser(defaultUser, userAuth)
	if err != nil {
		return &userpb.RegisterReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to register user: %v", err),
		}, nil
	}

	return &userpb.RegisterReply{
		Reply: utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginReply, error) {
	if request == nil {
		return &userpb.LoginReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil login request pb message"),
		}, nil
	}

	userAuth, err := adapter.AdaptToModelUserAuth(request.UserAuth)
	if err != nil {
		return &userpb.LoginReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt pb: %v", err),
		}, nil
	}

	if err := h.userService.AuthenticateUser(userAuth); err != nil {
		return &userpb.LoginReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to login: %v", err),
		}, nil
	}

	user, err := h.userService.GetUserByUserId(userAuth.UserID)
	if err != nil {
		return &userpb.LoginReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to get user info: %v", err),
		}, nil
	}

	resultPb, err := adapter.AdaptToPbUser(user)
	if err != nil {
		return &userpb.LoginReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to convert to pb: %v", err),
		}, nil
	}

	return &userpb.LoginReply{
		Reply: utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		User:  resultPb,
	}, nil
}

func (h *UserHandler) UpdateInfo(ctx context.Context, request *userpb.UpdateInfoRequest) (*userpb.UpdateInfoReply, error) {
	if request == nil {
		return &userpb.UpdateInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil update info request pb message"),
		}, nil
	}

	user, err := adapter.AdaptToModelUser(request.UserInfo)
	if err != nil {
		return &userpb.UpdateInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to adapt pb: %v", err),
		}, nil
	}

	// NOTE: The UpdateUser must has filtered nil resultUserInfo.
	resultUserInfo, err := h.userService.UpdateUser(user)
	if err != nil {
		return &userpb.UpdateInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to update user info: %v", err),
		}, nil
	}

	resultPb, err := adapter.AdaptToPbUser(resultUserInfo)
	if err != nil {
		return &userpb.UpdateInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to convert to pb: %v", err),
		}, nil
	}

	return &userpb.UpdateInfoReply{
		Reply:          utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		ResultUserInfo: resultPb,
	}, nil
}

func (h *UserHandler) GetInfo(ctx context.Context, request *userpb.GetInfoRequest) (*userpb.GetInfoReply, error) {
	if request == nil {
		return &userpb.GetInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil get info request pb message"),
		}, nil
	}

	user, err := h.userService.GetUserByUserId(request.UserId)
	if err != nil {
		return &userpb.GetInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to get user info: %v", err),
		}, nil
	}

	resultPb, err := adapter.AdaptToPbUser(user)
	if err != nil {
		return &userpb.GetInfoReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to convert to pb: %v", err),
		}, nil
	}

	return &userpb.GetInfoReply{
		Reply: utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		User:  resultPb,
	}, nil
}

func (h *UserHandler) ListUserPaged(ctx context.Context, request *userpb.ListUsersPagedRequest) (*userpb.ListUsersPagedReply, error) {
	if request == nil {
		return &userpb.ListUsersPagedReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "nil list user paged request pb message"),
		}, nil
	}

	pageNo, pageSize := request.GetPage().GetPageNo(), request.GetPage().GetPageSize()
	if utils.IsValidPagedInfo(pageNo, pageSize) {
		return &userpb.ListUsersPagedReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE,
				"get invalid page info: pageNo=%v, pageSize=%v", pageNo, pageSize),
		}, nil
	}

	users, totalPages, err := h.userService.ListUserPaged(pageNo, pageSize)
	if err != nil {
		return &userpb.ListUsersPagedReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE,
				"failed to list users pageNo %v pageSize %v: %v", pageNo, pageSize, err),
		}, nil
	}

	resultPb, err := adapter.AdaptToPbUsers(users)
	if err != nil {
		return &userpb.ListUsersPagedReply{
			Reply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "failed to convert to pb: %v", err),
		}, nil
	}

	return &userpb.ListUsersPagedReply{
		Reply:      utils.PbReplyf(base.REPLY_STATUS_SUCCESS, defaultSuccessMessage),
		Users:      resultPb,
		TotalPages: totalPages,
	}, nil
}
