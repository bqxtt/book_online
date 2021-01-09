package adapter

import (
	"github.com/bqxtt/book_online/book/pkg/model"
	"github.com/bqxtt/book_online/book/pkg/sdk/bookpb"
)

func RpcBookToModelBook(rpcBook *bookpb.Book) *model.Book {
	return &model.Book{
		ImgUrl:      rpcBook.ImgUrl,
		ISBN:        rpcBook.ISBN,
		BookName:    rpcBook.BookName,
		Publisher:   rpcBook.Publisher,
		Category:    rpcBook.Category,
		Author:      rpcBook.Author,
		Description: rpcBook.Description,
		RentStatus:  int8(rpcBook.RentStatus),
	}
}

func ModelBookToRpcBook(modelBook *model.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:          modelBook.ID,
		ISBN:        modelBook.ISBN,
		BookName:    modelBook.BookName,
		Publisher:   modelBook.Publisher,
		Category:    modelBook.Category,
		Author:      modelBook.Author,
		ImgUrl:      modelBook.ImgUrl,
		Description: modelBook.Description,
		RentStatus:  int32(modelBook.RentStatus),
	}
}
