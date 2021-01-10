package dao

import (
	"github.com/bqxtt/book_online/book/pkg/common"
	"github.com/bqxtt/book_online/book/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBDirectDAO struct {
	DB *gorm.DB
}

func NewDBDirectDAO() (*DBDirectDAO, error) {
	db, err := gorm.Open("mysql", "root:19981108@tcp(139.9.140.136:3307)/book_online?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		return nil, err
	}
	return &DBDirectDAO{
		DB: db,
	}, nil
}

func (DB *DBDirectDAO) GetBooksByPage(page int32, pageSize int32) ([]*model.Book, error) {
	var books []*model.Book
	result := DB.DB.Where("status = ?", common.BOOK_AVALIABLE).Offset((page - 1) * pageSize).Limit(pageSize).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (DB *DBDirectDAO) GetBookById(id int64) (*model.Book, error) {
	var book = &model.Book{}
	err := DB.DB.Where("id = ?", id).Find(book).Error
	if err != nil {
		return nil, err
	}
	//fmt.Println(book)
	return book, nil
}

func (DB *DBDirectDAO) CreateBook(book *model.Book) (*model.Book, error) {
	result := DB.DB.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (DB *DBDirectDAO) UpdateBook(book *model.Book) (*model.Book, error) {
	result := DB.DB.Model(book).Update(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (DB *DBDirectDAO) DeleteBookById(bookId int64) error {
	result := DB.DB.Model(&model.Book{ID: bookId}).Update("status", common.BOOK_DELETED)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (DB *DBDirectDAO) GetAllBooks() ([]*model.Book, error) {
	panic("implement me")
}

func (DB *DBDirectDAO) GetBooksCount() (int32, error) {
	var count int32
	result := DB.DB.Model(&model.Book{}).Where("status = ?", common.BOOK_AVALIABLE).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (DB *DBDirectDAO) SetBookRentStatus(bookId int64, rentStatus common.BookRentStatus) error {
	result := DB.DB.Model(&model.Book{ID: bookId}).Update("rent_status", rentStatus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
