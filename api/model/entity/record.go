package entity

import "time"

type Record struct {
	UserId     int64
	Book       *Book
	BorrowedAt time.Time
	ReturnedAt time.Time
	Deadline   time.Time
	Status     RecordStatus
}

type RecordStatus int32

const (
	BOOK_BORROWED RecordStatus = 1
	BOOK_RETURNED RecordStatus = 2
)
