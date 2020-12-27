package dao

import (
	"github.com/bqxtt/book_online/book/pkg/model"
)

type DBDirectDAO struct {
}

func NewDBDirectDAO() (*DBDirectDAO, error) {
	return &DBDirectDAO{}, nil
}

func (DB *DBDirectDAO) GetBooksByPage(page int32, pageSize int32) (books []*model.Book, err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) GetBookById(id int64) (book *model.Book, err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) CreateBook(book *model.Book) (resultBook *model.Book, err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) UpdateBook(book *model.Book) (resultBook *model.Book, err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) DeleteBookById(bookId int64) (err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) GetAllBooks() (books []*model.Book, err error) {
	panic("implement me")
}
