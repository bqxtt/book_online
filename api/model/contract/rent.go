package contract

type BorrowBookRequest struct {
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
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListReturnedBookRequest struct {
}

type ListReturnedBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListBookRecordsRequest struct {
}

type ListBookRecordsResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListAllBorrowedBookRequest struct {
}

type ListAllBorrowedBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListAllReturnedBookRequest struct {
}

type ListAllReturnedBookResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}

type ListAllBookRecordsRequest struct {
}

type ListAllBookRecordsResponse struct {
	BaseResponse *BaseResponse `json:"base_response"`
}
