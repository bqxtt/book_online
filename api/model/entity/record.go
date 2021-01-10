package entity

type Record struct {
	UserId     int64        `json:"user_id"`
	Book       *Book        `json:"book"`
	BorrowedAt string       `json:"borrowed_at"`
	ReturnedAt string       `json:"returned_at"`
	Deadline   string       `json:"deadline"`
	Status     RecordStatus `json:"status"`
}

type RecordStatus int32

const (
	BOOK_BORROWED RecordStatus = 1
	BOOK_RETURNED RecordStatus = 2
)
