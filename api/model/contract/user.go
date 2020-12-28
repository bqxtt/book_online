package contract

import "github.com/bqxtt/book_online/api/model/entity"

type RegisterRequest struct {
	User *entity.User `form:"user" json:"user" binding:"required"`
}

type RegisterResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type LoginRequest struct {
	User *entity.User `form:"user" json:"user" binding:"required"`
}

type LoginResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
	Role         string        `json:"role"`
	Token        string        `json:"token"`
}

type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
}

type UpdateUserInfoRequest struct {
}

type UpdateUserInfoResponse struct {
}

type ListAllUsersRequest struct {
}

type ListAllUsersResponse struct {
}

type DeleteUserRequest struct {
}

type DeleteUserResponse struct {
}
