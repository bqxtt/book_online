package dao

import (
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
)

type InMemoryDAO struct {
	BookMap map[int64]*model.Book
}

func NewInMemoryDAO() (*InMemoryDAO, error) {
	return &InMemoryDAO{
		BookMap: make(map[int64]*model.Book),
	}, nil
}

func (in *InMemoryDAO) GetBooksByPage(page int32, pageSize int32) (books []*model.Book, err error) {
	panic("implement me")
}

func (in *InMemoryDAO) GetBookById(id int64) (book *model.Book, err error) {
	panic("implement me")
}

func (in *InMemoryDAO) CreateBook(book *model.Book) (resultBook *model.Book, err error) {
	panic("implement me")
}

func (in *InMemoryDAO) UpdateBook(book *model.Book) (resultBook *model.Book, err error) {
	panic("implement me")
}

func (in *InMemoryDAO) DeleteBookById(bookId int64) (err error) {
	panic("implement me")
}

func (in *InMemoryDAO) GetAllBooks() (books []*model.Book, err error) {
	panic("implement me")
}

func (in *InMemoryDAO) GetBooksCount() (int32, error) {
	panic("implement me")
}

func (in *InMemoryDAO) SetBookRentStatus(bookId int64, rentStatus common.BookRentStatus) error {
	panic("implement me")
}
