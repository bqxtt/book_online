package service

import (
	"errors"
	"fmt"
	"github.com/bqxtt/book_online/rent/pkg/common"
	"github.com/bqxtt/book_online/rent/pkg/model"
	"github.com/bqxtt/book_online/rent/pkg/model/dao"
	"github.com/bqxtt/book_online/rent/pkg/service/clients"
	"github.com/bqxtt/book_online/rent/pkg/utils"
	"time"
)

type RentService struct {
	rentDAO dao.RentDAOInterface

	bookClient *clients.BookClient
}

func NewRentService() (*RentService, error) {
	rd, err := dao.NewRentDAO(common.DAOTypeDefault)
	if err != nil {
		return nil, err
	}

	bookClient, err := clients.NewBookClient()
	if err != nil {
		return nil, err
	}

	return &RentService{
		rentDAO:    rd,
		bookClient: bookClient,
	}, nil
}

func (r *RentService) ListBorrowedBook(userId int64, pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListBorrowedRecordByUserId(userId, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) ListReturnedBook(userId int64, pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListReturnedRecordByUserId(userId, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) ListAllBorrowedBook(pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListAllBorrowedRecord(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) ListAllReturnedBook(pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListAllReturnedRecord(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) ReturnBook(recordId int64) error {
	record, err := r.rentDAO.GetRentRecordById(recordId)
	if err != nil {
		return err
	}
	if record == nil {
		return fmt.Errorf("nil record with id %v", recordId)
	}
	record.ReturnedAt = time.Now()
	record.Finished = true

	// 1. check book availability
	borrowed, err := r.bookClient.CheckBookBorrowed(record.BookID)
	if err != nil {
		return err
	}
	if !borrowed {
		return fmt.Errorf("invalid record. book with id %v not borrowed", record.BookID)
	}

	// 2. set book availability in another goroutine
	errChan := make(chan error, 1)
	doneChan := make(chan bool, 1)
	go func() {
		shouldContinue := <-doneChan
		if !shouldContinue {
			errChan <- errors.New("update record failed")
			return
		}

		errChan <- r.bookClient.SetBookReturned(record.BookID)
	}()

	if _, err = r.rentDAO.TransactionReturn(errChan, doneChan, record); err != nil {
		return err
	}

	return nil
}

func (r *RentService) ListRecordByUserId(userId int64, pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListRecordByUserId(userId, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) ListAllRecord(pageNo int64, pageSize int64) ([]*model.RentRecordDetail, int64, int64, error) {
	limit, offset := utils.CalculateLimitOffset(pageNo, pageSize)

	records, total, err := r.rentDAO.ListAllRecord(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	details, err := r.fillOutRecordDetails(records)
	if err != nil {
		return nil, 0, 0, err
	}

	return details, total, (total + pageSize - 1) / pageSize, nil
}

func (r *RentService) BorrowBook(userId int64, bookId int64) (deadline time.Time, err error) {
	borrowedAt, deadline := time.Now(), time.Now().Add(common.DefaultBookKeptInterval*24*time.Hour)
	// set default
	returnedAt := time.Unix(0, 0)

	record := &model.RentRecord{
		UserID:     userId,
		BookID:     bookId,
		BorrowedAt: borrowedAt,
		ReturnedAt: returnedAt,
		Deadline:   deadline,
		Finished:   false,
	}

	// 1. check book availability
	borrowed, err := r.bookClient.CheckBookBorrowed(record.BookID)
	if err != nil {
		return time.Time{}, err
	}
	if borrowed {
		return time.Time{}, fmt.Errorf("book with id %v already borrowed", record.BookID)
	}

	// 2. set book availability in another goroutine
	errChan := make(chan error, 1)
	doneChan := make(chan bool, 1)
	go func() {
		shouldContinue := <-doneChan
		if !shouldContinue {
			errChan <- errors.New("create record failed")
			return
		}

		errChan <- r.bookClient.SetBookBorrowed(record.BookID)
	}()

	// 3. create rent record
	record, err = r.rentDAO.TransactionRent(errChan, doneChan, record)
	if err != nil {
		return time.Time{}, err
	}

	return record.Deadline, nil
}

func (r *RentService) fillOutRecordDetails(records []*model.RentRecord) ([]*model.RentRecordDetail, error) {
	details := make([]*model.RentRecordDetail, 0, len(records))
	for _, record := range records {
		book, err := r.bookClient.GetBookById(record.BookID)
		if err != nil {
			return nil, err
		}

		details = append(details, &model.RentRecordDetail{
			RentRecord: record,
			Book:       book,
		})
	}

	return details, nil
}
