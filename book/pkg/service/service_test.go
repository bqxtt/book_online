package service

import (
	"fmt"
	"testing"
)

func TestBookService_CheckBookBorrowed(t *testing.T) {
	svc, _ := NewBookService()
	borrowed, err := svc.CheckBookBorrowed(2)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(borrowed)
}

func TestBookService_SetBookBorrowed(t *testing.T) {
	svc, _ := NewBookService()
	_ = svc.SetBookBorrowed(2)
}
