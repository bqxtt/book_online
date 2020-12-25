package contract

import "github.com/bqxtt/book_online/api/model/entity"

type RegisterRequest struct {
	User *entity.User `form:"user" json:"user" binding:"required"`
}

type RegisterResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}
