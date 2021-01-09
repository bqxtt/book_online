package dao

import (
	"fmt"
	"testing"
)

func TestBookDB(t *testing.T) {
	db, err := NewDBDirectDAO()
	if err != nil {
		fmt.Println(err)
		return
	}
	book, err := db.GetBookById(1)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(book)
}
