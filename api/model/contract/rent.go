package contract

import "github.com/bqxtt/book_online/api/model/entity"

type BorrowBookRequest struct {
	BookId int64 `form:"book_id" json:"book_id" binding:"required"`
}

type BorrowBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ReturnBookRequest struct {
}

type ReturnBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListBorrowedBookRequest struct {
}

type ListBorrowedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}

type ListReturnedBookRequest struct {
}

type ListReturnedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}

type ListBookRecordsRequest struct {
}

type ListBookRecordsResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}

type ListAllBorrowedBookRequest struct {
}

type ListAllBorrowedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}

type ListAllReturnedBookRequest struct {
}

type ListAllReturnedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}

type ListAllBookRecordsRequest struct {
}

type ListAllBookRecordsResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
}
