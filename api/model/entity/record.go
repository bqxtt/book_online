package entity

type Record struct {
	UserId int64
	Book   *Book
	Status RecordStatus
}

type RecordStatus int32

const (
	BOOK_BORROWED RecordStatus = 1
	BOOK_RETURNED RecordStatus = 2
)
