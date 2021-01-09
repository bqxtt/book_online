package service

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/api/adapter"
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_book"
	"github.com/bqxtt/book_online/rpc/clients/rpc_book/bookpb"
)

type IBookService interface {
	ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, *entity.PageInfo, error)
	CreateBook(book *entity.Book) error
	UpdateBook(book *entity.Book) error
	DeleteBookById(bookId int64) error
}

type bookServiceImpl struct{}

var BookService IBookService = &bookServiceImpl{}

func (bs *bookServiceImpl) ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, *entity.PageInfo, error) {
	request := &bookpb.GetBooksByPageRequest{
		Page:     page,
		PageSize: pageSize,
	}
	resp, err := rpc_book.BookServiceClient.GetBooksByPage(context.Background(), request)
	if err != nil {
		return nil, nil, fmt.Errorf("rpc book service error, err: %v", err)
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("rpc book service resp is nil")
	}
	rpcBooks := resp.Books
	entityBooks := make([]*entity.Book, 0)
	for _, v := range rpcBooks {
		entityBooks = append(entityBooks, adapter.RpcBookToEntityBook(v))
	}
	return entityBooks, &entity.PageInfo{
		TotalPages: resp.TotalPages,
		TotalCount: resp.TotalCount,
	}, nil
}

func (bs *bookServiceImpl) CreateBook(book *entity.Book) error {
	request := &bookpb.CreateBookRequest{
		Book: adapter.EntityBookToRpcBook(book),
	}
	resp, err := rpc_book.BookServiceClient.CreateBook(context.Background(), request)
	if err != nil {
		return fmt.Errorf("rpc book servcice error, err: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("rpc book service resp is nil")
	}
	if resp.BaseReply.Status != 1 {
		return fmt.Errorf("rpc book service error, err: %v", resp.BaseReply.Message)
	}
	return nil
}

func (bs *bookServiceImpl) UpdateBook(book *entity.Book) error {
	request := &bookpb.UpdateBookRequest{
		Id:   book.BookId,
		Book: adapter.EntityBookToRpcBook(book),
	}
	resp, err := rpc_book.BookServiceClient.UpdateBook(context.Background(), request)
	if err != nil {
		return fmt.Errorf("rpc book service error, err: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("rpc book serivce resp is nil")
	}
	if resp.BaseReply.Status != 1 {
		return fmt.Errorf("rpc book service error, err: %v", resp.BaseReply.Message)
	}
	return nil
}

func (bs *bookServiceImpl) DeleteBookById(bookId int64) error {
	request := &bookpb.DeleteBookByIdRequest{
		Id: bookId,
	}
	resp, err := rpc_book.BookServiceClient.DeleteBookById(context.Background(), request)
	if err != nil {
		return fmt.Errorf("rpc book service error, err: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("rpc book service resp is nil")
	}
	if resp.BaseReply.Status != 1 {
		return fmt.Errorf("rpc book service error, err: %v", err)
	}
	return nil
}
