package service

import (
	"fmt"
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
	"github.com/bqxtt/book_online/book/pkg/model/dao"
)

type BookService struct {
	bookDao dao.IBookDao
}

func NewBookService() (*BookService, error) {
	bd, err := dao.NewBookDao(common.DAOTypeInMemory)
	return &BookService{
		bookDao: bd,
	}, err
}
func (bs *BookService) GetAllBooks() ([]*model.Book, error) {
	result, err := bs.bookDao.GetAllBooks()
	if err != nil {
		return nil, fmt.Errorf("fail to get all books, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) GetBooksByPage(page int32, pageSize int32) ([]*model.Book, error) {
	result, err := bs.bookDao.GetBooksByPage(page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("fail to get books by page, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) GetBookById(id int64) (*model.Book, error) {
	result, err := bs.bookDao.GetBookById(id)
	if err != nil {
		return nil, fmt.Errorf("fail to get book by id, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	result, err := bs.bookDao.CreateBook(book)
	if err != nil {
		return nil, fmt.Errorf("fail to create book, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) UpdateBook(book *model.Book) (*model.Book, error) {
	result, err := bs.bookDao.UpdateBook(book)
	if err != nil {
		return nil, fmt.Errorf("fail to update book, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) DeleteBookById(bookId int64) error {
	err := bs.bookDao.DeleteBookById(bookId)
	if err != nil {
		return fmt.Errorf("fail to delete book by id, err: %v", err)
	}
	return nil
}
