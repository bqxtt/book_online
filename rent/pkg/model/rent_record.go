package model

import (
	"github.com/bqxtt/book_online/rent/pkg/sdk/include/bookpb"
	"time"
)

type RentRecord struct {
	ID         int64     `gorm:"AUTO_INCREMENT;primary_key"`
	UserID     int64     `gorm:"type:bigint;not null"`
	BookID     int64     `gorm:"type:bigint;not null"`
	BorrowedAt time.Time `gorm:"type:DATETIME(2);not null"`
	Deadline   time.Time `gorm:"type:DATETIME(2);not null"`
	ReturnedAt time.Time `gorm:"type:DATETIME(2);not null"`
	Finished   bool      `gorm:"type:bool;not null"`
}

// detailed rent record for query
type RentRecordDetail struct {
	*RentRecord
	Book *bookpb.Book
}
