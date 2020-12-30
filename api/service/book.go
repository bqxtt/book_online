package service

import (
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/book/pkg/sdk/bookpb"
)

type IBookService interface {
	ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, error)
	CreateBook(book *entity.Book) error
	UpdateBook(book *entity.Book) error
	DeleteBookById(bookId int64) error
}

type bookServiceImpl struct{}

var BookService IBookService = &bookServiceImpl{}

func (bs *bookServiceImpl) ListBooksByPage(page int32, pageSize int32) ([]*entity.Book, error) {
	request := &bookpb.GetBooksByPageRequest{
		Page:     page,
		PageSize: pageSize,
	}

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
