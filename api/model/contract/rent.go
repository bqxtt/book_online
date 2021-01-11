package contract

import "github.com/bqxtt/book_online/api/model/entity"

type BorrowBookRequest struct {
	BookId int64 `form:"book_id" json:"book_id" binding:"required"`
}

type BorrowBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
	Deadline     string        `json:"deadline"`
}

type ReturnBookRequest struct {
	RecordId int64 `form:"record_id" json:"record_id" binding:"required"`
}

type ReturnBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListBorrowedBookRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListBorrowedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type ListReturnedBookRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListReturnedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type ListBookRecordsRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListBookRecordsResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type ListAllBorrowedBookRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListAllBorrowedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type ListAllReturnedBookRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListAllReturnedBookResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}

type ListAllBookRecordsRequest struct {
	Page     int32 `form:"page" json:"page" binding:"required"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required"`
}

type ListAllBookRecordsResponse struct {
	BaseResponse *BaseResponse    `json:"base_response"`
	Records      []*entity.Record `json:"records"`
	PageInfo     *entity.PageInfo `json:"page_info"`
}
