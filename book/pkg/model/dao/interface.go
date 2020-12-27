package dao

import (
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
)

type IBookDao interface {
	GetAllBooks() (books []*model.Book, err error)
	GetBooksByPage(page int32, pageSize int32) (books []*model.Book, err error)
	GetBookById(id int64) (book *model.Book, err error)
	CreateBook(book *model.Book) (resultBook *model.Book, err error)
	UpdateBook(book *model.Book) (resultBook *model.Book, err error)
	DeleteBookById(bookId int64) (err error)
}

func NewBookDao(daoType common.DAOType) (IBookDao, error) {
	switch daoType {
	//todo
	default:
		return NewInMemoryDAO()
	}
}
