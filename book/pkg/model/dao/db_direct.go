package dao

import (
	"github.com/bqxtt/book_online/book/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBDirectDAO struct {
	DB *gorm.DB
}

func NewDBDirectDAO() (*DBDirectDAO, error) {
	db, err := gorm.Open("mysql", "root:19981108@tcp(139.9.140.136:3307)/book_online?charset=utf8mb4&loc=Local")
	if err != nil {
		return nil, err
	}
	return &DBDirectDAO{
		DB: db,
	}, nil
}

func (DB *DBDirectDAO) GetBooksByPage(page int32, pageSize int32) ([]*model.Book, error) {
	panic("")
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

func (DB *DBDirectDAO) UpdateBook(book *model.Book) (resultBook *model.Book, err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) DeleteBookById(bookId int64) (err error) {
	panic("implement me")
}

func (DB *DBDirectDAO) GetAllBooks() (books []*model.Book, err error) {
	panic("implement me")
}
