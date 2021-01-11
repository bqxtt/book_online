package contract

import "github.com/bqxtt/book_online/api/model/entity"

type RegisterRequest struct {
	UserAuth *entity.UserAuth `form:"user_auth" json:"user_auth" binding:"required"`
}

type RegisterResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type LoginRequest struct {
	UserAuth *entity.UserAuth `form:"user_auth" json:"user_auth" binding:"required"`
}

type LoginResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
	UserInfo     *entity.User  `json:"user_info"`
}

type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
	User         *entity.User  `json:"user"`
}

type UpdateUserInfoRequest struct {
	User *entity.User `form:"user" json:"user" binding:"required"`
}

type UpdateUserInfoResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
	User         *entity.User  `json:"user"`
}

type ListAllUsersRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListAllUsersResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Users        []*entity.User   `json:"users"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type DeleteUserRequest struct {
}

type DeleteUserResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}
