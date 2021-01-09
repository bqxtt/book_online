package contract

import "github.com/bqxtt/book_online/api/model/entity"

type ListBooksRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListBooksResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Books        []*entity.Book   `json:"books"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type CreateBookRequest struct {
	Book *entity.Book `form:"book" json:"book"`
}

type CreateBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type UpdateBookRequest struct {
	Book *entity.Book `form:"book" json:"book"`
}

type UpdateBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type DeleteBookRequest struct {
}

type DeleteBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}
