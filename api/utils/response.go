package utils

import (
	"fmt"
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/model/entity"
	"time"
)

func NewSuccessResponse(format string, args ...interface{}) *contract.BaseResponse {
	statusInfo := &contract.StatusInfo{
		Time:    time.Now().Unix(),
		Message: fmt.Sprintf(format, args...),
	}
	return &contract.BaseResponse{
		StatusCode: contract.SUCCESS,
		StatusInfo: statusInfo,
	}
}

func NewFailureResponse(format string, args ...interface{}) *contract.BaseResponse {
	statusInfo := &contract.StatusInfo{
		Time:    time.Now().Unix(),
		Message: fmt.Sprintf(format, args...),
	}
	return &contract.BaseResponse{
		StatusCode: contract.FAILURE,
		StatusInfo: statusInfo,
	}
}

// just for test
func NewDefaultUser() *entity.User {
	return &entity.User{
		UserId: "10175101201",
		Name:   "tcg",
	}
}
