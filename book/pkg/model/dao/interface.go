package dao

import (
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
)

type IBookDao interface {
	GetAllBooks() ([]*model.Book, error)
	GetBooksByPage(page int32, pageSize int32) ([]*model.Book, error)
	GetBooksCount() (int32, error)
	GetBookById(id int64) (*model.Book, error)
	CreateBook(book *model.Book) (*model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	//软删除，status
	DeleteBookById(bookId int64) error
}

func NewBookDao(daoType common.DAOType) (IBookDao, error) {
	switch daoType {
	case common.DAOTypeDBDirectConn:
		return NewDBDirectDAO()
	//todo
	default:
		return NewInMemoryDAO()
	}
}
