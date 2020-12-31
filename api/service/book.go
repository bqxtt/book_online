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
	ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, int32, error)
	CreateBook(book *entity.Book) error
	UpdateBook(book *entity.Book) error
	DeleteBookById(bookId int64) error
}

type bookServiceImpl struct{}

var BookService IBookService = &bookServiceImpl{}

func (bs *bookServiceImpl) ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, int32, error) {
	request := &bookpb.GetBooksByPageRequest{
		Page:     page,
		PageSize: pageSize,
	}
	resp, err := rpc_book.BookServiceClient.GetBooksByPage(context.Background(), request)
	if err != nil {
		return nil, 0, fmt.Errorf("rpc book service error, err: %v", err)
	}
	if resp == nil {
		return nil, 0, fmt.Errorf("rpc book service resp is nil")
	}
	rpcBooks := resp.Books
	entityBooks := make([]*entity.Book, 0)
	for _, v := range rpcBooks {
		entityBooks = append(entityBooks, adapter.RpcBookToEntityBook(v))
	}
	return entityBooks, resp.TotalPages, nil
}

func (bs *bookServiceImpl) CreateBook(book *entity.Book) error {
	panic("implement me")
}

func (bs *bookServiceImpl) UpdateBook(book *entity.Book) error {
	panic("implement me")
}

func (bs *bookServiceImpl) DeleteBookById(bookId int64) error {
	panic("implement me")
}
