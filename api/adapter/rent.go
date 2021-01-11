package adapter

import (
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/model/rentpb"
)

func RpcRecordStatusToEntityRecordStatus(status rentpb.RentStatus) entity.RecordStatus {
	returned := entity.BOOK_RETURNED
	borrowed := entity.BOOK_BORROWED
	if status == rentpb.RentStatus_FINISHED {
		return returned
	} else {
		return borrowed
	}
}

func RpcRecordToEntityRecord(record *rentpb.RentRecord) *entity.Record {
	return &entity.Record{
		RecordId:   record.Id,
		UserId:     record.UserId,
		Book:       RpcBookToEntityBook(record.Book),
		BorrowedAt: record.BorrowedAt.AsTime().Format("2006/01/02"),
		ReturnedAt: record.ReturnedAt.AsTime().Format("2006/01/02"),
		Deadline:   record.Deadline.AsTime().Format("2006/01/02"),
		Status:     RpcRecordStatusToEntityRecordStatus(record.Status),
	}
}
