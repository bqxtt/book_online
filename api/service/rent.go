package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/adapter"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_rent"
	"github.com/bqxtt/book_online/rpc/model/base"
	"github.com/bqxtt/book_online/rpc/model/rentpb"
	"time"
)

type IRentService interface {
	BorrowBook(userId int64, bookId int64) (time.Time, error)
	ReturnBook(recordId int64) error
	ListBorrowedBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
	ListReturnedBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
	ListBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
	ListAllBorrowedBook(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
	ListAllReturnedBook(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
	ListAllBookRecords(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error)
}

type RentServiceImpl struct{}

var RentService IRentService = &RentServiceImpl{}

func (rs *RentServiceImpl) BorrowBook(userId int64, bookId int64) (time.Time, error) {
	request := &rentpb.BorrowBookRequest{
		UserId: userId,
		BookId: bookId,
	}
	resp, err := rpc_rent.RentServiceClient.BorrowBook(context.Background(), request)
	if err != nil {
		return time.Time{}, err
	}
	if resp == nil {
		return time.Time{}, fmt.Errorf("borrow book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return time.Time{}, fmt.Errorf("rpc rent service error, err: %v", resp.Reply.Message)
	}
	return resp.Deadline.AsTime(), nil
}

func (rs *RentServiceImpl) ReturnBook(recordId int64) error {
	request := &rentpb.ReturnBookRequest{
		RecordId: recordId,
	}
	resp, err := rpc_rent.RentServiceClient.ReturnBook(context.Background(), request)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("return book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return fmt.Errorf("rpc rent service error, err: %v", resp.Reply.Message)
	}
	return nil
}

func (rs *RentServiceImpl) ListBorrowedBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListBorrowedBookRequest{
		UserId: userId,
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListBorrowedBook(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list borrowed book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc rent service error, err: %v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (rs *RentServiceImpl) ListReturnedBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListReturnedBookRequest{
		UserId: userId,
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListReturnedBook(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list returned book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc service error, err:%v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (rs *RentServiceImpl) ListBook(userId int64, page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListBookRequest{
		UserId: userId,
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListBook(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc service error, err:%v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (rs *RentServiceImpl) ListAllBorrowedBook(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListAllBorrowedBookRequest{
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListAllBorrowedBook(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list all borrowed book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc service error, err:%v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (rs *RentServiceImpl) ListAllReturnedBook(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListAllReturnedBookRequest{
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListAllReturnedBook(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list all returned book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc service error, err:%v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (rs *RentServiceImpl) ListAllBookRecords(page int64, pageSize int64) ([]*entity.Record, *entity.PageInfo, error) {
	request := &rentpb.ListAllBookRecordsRequest{
		Page: &base.Page{
			PageNo:   page,
			PageSize: pageSize,
		},
	}
	resp, err := rpc_rent.RentServiceClient.ListAllBookRecords(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("list all book resp is nil")
	}
	if resp.Reply.Status != 1 {
		return nil, nil, fmt.Errorf("rpc service error, err:%v", resp.Reply.Message)
	}
	var records []*entity.Record
	for _, r := range resp.Records {
		records = append(records, adapter.RpcRecordToEntityRecord(r))
	}
	return records, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}
