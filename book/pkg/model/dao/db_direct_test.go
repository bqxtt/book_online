package dao

import (
	"fmt"
	"github.com/bqxtt/book_online/book/pkg/model"
	"testing"
	"time"
)

func TestBookDB(t *testing.T) {
	db, err := NewDBDirectDAO()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.GetBookById(1)
	if err != nil {
		fmt.Print(err)
		return
	}
	//fmt.Println(book)
}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:         1,
		Status:     2,
		RentStatus: 2,
		//ImgUrl:      "dsfa",
		//ISBN:        "gjlkfdjlgk",
		//BookName:    "kasjdklfj",
		//Publisher:   "iuwioej",
		//Category:    "diosajflk",
		//Author:      "lkajdlk",
		//Description: "kadsjflk",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	db, _ := NewDBDirectDAO()
	_, _ = db.UpdateBook(book)
}

func TestFindBookByPage(t *testing.T) {
	db, _ := NewDBDirectDAO()
	books, err := db.GetBooksByPage(2, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(books)
}

func TestBooksCount(t *testing.T) {
	db, _ := NewDBDirectDAO()
	count, err := db.GetBooksCount()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func TestDBDirectDAO_CreateBook(t *testing.T) {
	db, _ := NewDBDirectDAO()
	book := &model.Book{
		ID:          0,
		Status:      1,
		RentStatus:  1,
		ImgUrl:      "s",
		ISBN:        "s",
		BookName:    "s",
		Publisher:   "s",
		Category:    "s",
		Author:      "s",
		Description: "s",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	book, err := db.CreateBook(book)
	if err != nil {
		fmt.Println(err)
	}
}
