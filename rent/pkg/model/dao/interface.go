package dao

import (
	"github.com/bqxtt/book_online/rent/pkg/common"
	"github.com/bqxtt/book_online/rent/pkg/model"
)

type RentDAOInterface interface {
	// TODO: add proxy/gORM conns
	CreateRentRecord(record *model.RentRecord) (result *model.RentRecord, err error)
	GetRentRecordById(recordId int64) (result *model.RentRecord, err error)
	UpdateRentRecord(record *model.RentRecord) (result *model.RentRecord, err error)
	ListBorrowedRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error)
	ListReturnedRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error)
	ListRecordByUserId(userId int64, limit int64, offset int64) ([]*model.RentRecord, error)
	ListAllBorrowedRecord(limit int64, offset int64) ([]*model.RentRecord, error)
	ListAllReturnedRecord(limit int64, offset int64) ([]*model.RentRecord, error)
	ListAllRecord(limit int64, offset int64) ([]*model.RentRecord, error)
	TransactionRent(errChan <-chan error, doneChan chan<- bool, record *model.RentRecord) (result *model.RentRecord, err error)
	TransactionReturn(errChan <-chan error, doneChan chan<- bool, record *model.RentRecord) (result *model.RentRecord, err error)
}

func NewRentDAO(daoType common.DAOType) (RentDAOInterface, error) {
	switch daoType {
	// TODO:
	default:
		return NewDBDirectDAO()
	}
}
