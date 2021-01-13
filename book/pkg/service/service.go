package service

import (
	"fmt"
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
	"github.com/bqxtt/book_online/book/pkg/model/dao"
	"time"
)

type BookService struct {
	bookDao dao.IBookDao
}

func NewBookService() (*BookService, error) {
	bd, err := dao.NewBookDao(common.DAOTypeDBDirectConn)
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

func (bs *BookService) GetBooksByPage(page int32, pageSize int32) ([]*model.Book, *common.PageInfo, error) {
	count, err := bs.bookDao.GetBooksCount()
	if err != nil {
		return nil, nil, fmt.Errorf("fail to get books count, err: %v", err)
	}
	result, err := bs.bookDao.GetBooksByPage(page, pageSize)
	if err != nil {
		return nil, nil, fmt.Errorf("fail to get books by page, err: %v", err)
	}
	return result, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (bs *BookService) GetBookById(id int64) (*model.Book, error) {
	result, err := bs.bookDao.GetBookById(id)
	if err != nil {
		return nil, fmt.Errorf("fail to get book by id, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	book.CreateTime = time.Now()
	book.UpdateTime = time.Now()
	book.RentStatus = 1
	book.Status = int8(common.BOOK_AVALIABLE)
	result, err := bs.bookDao.CreateBook(book)
	if err != nil {
		return nil, fmt.Errorf("fail to create book, err: %v", err)
	}
	return result, nil
}

func (bs *BookService) UpdateBook(bookId int64, book *model.Book) (*model.Book, error) {
	book.UpdateTime = time.Now()
	result, err := bs.bookDao.UpdateBook(bookId, book)
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

func (bs *BookService) CheckBookBorrowed(bookId int64) (bool, error) {
	book, err := bs.bookDao.GetBookById(bookId)
	if err != nil {
		return false, err
	}
	return common.BookRentStatus(book.RentStatus) == common.BOOK_BORROWED, nil
}

func (bs *BookService) SetBookBorrowed(bookId int64) error {
	err := bs.bookDao.SetBookRentStatus(bookId, common.BOOK_BORROWED)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) SetBookReturned(bookId int64) error {
	err := bs.bookDao.SetBookRentStatus(bookId, common.BOOK_RETURNED)
	if err != nil {
		return err
	}
	return nil
}
