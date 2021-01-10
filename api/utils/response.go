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

func NewDefaultBook() *entity.Book {
	return &entity.Book{
		BookId:      2,
		ImgUrl:      "http://e.library.sh.cn/wread2017/_cover/cx/ShlibEpub_1487928690.jpg",
		BookName:    "昆曲日记",
		ISBN:        "9787302356288",
		Publisher:   "中央编译出版社",
		Category:    "艺术",
		Author:      "张允和",
		Description: "《昆曲日记》记录了北京昆曲研习社从创办以来的日常活动，记录了自1956年以来诸多社会文化名流对昆曲传承的支持，包括文化部、文化局以及各剧院团对昆曲事业做出的不懈努力，也记录了海内外曲友对于昆曲事业的执著。《昆曲日记》内容丰富，史料翔实，对于研究以及热爱、支持昆曲的各方人士来说，《昆曲日记》都是不可多得的必藏之书。此次修订，本书修正了上一版在知识上及文字上的讹误，重新校订了附录中曲人名录的详细信息，并将附录中北京昆曲研习社大事记更新到2012年，对于研究及热爱昆曲的各方人士来说，这一次修订使本书更具收藏价值。",
		RentStatus:  1,
	}
}
func NewDefaultRecords(userId int64, status entity.RecordStatus) *entity.Record {
	now := time.Now()
	return &entity.Record{
		UserId:     userId,
		Book:       NewDefaultBook(),
		BorrowedAt: now.Format("2006-01-02"),
		ReturnedAt: now.Format("2006-01-02"),
		Deadline:   now.Format("2006-01-02"),
		Status:     status,
	}
}
