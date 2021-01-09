package adapter

import (
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_book/bookpb"
)

func RpcBookToEntityBook(rpcBook *bookpb.Book) *entity.Book {
	return &entity.Book{
		BookId:      rpcBook.Id,
		ImgUrl:      rpcBook.ImgUrl,
		BookName:    rpcBook.BookName,
		ISBN:        rpcBook.ISBN,
		Publisher:   rpcBook.Publisher,
		Category:    rpcBook.Category,
		Author:      rpcBook.Author,
		Description: rpcBook.Description,
		RentStatus:  int8(rpcBook.RentStatus),
	}
}

func EntityBookToRpcBook(entityBook *entity.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:          entityBook.BookId,
		ISBN:        entityBook.ISBN,
		BookName:    entityBook.BookName,
		Publisher:   entityBook.Publisher,
		Category:    entityBook.Category,
		Author:      entityBook.Author,
		ImgUrl:      entityBook.ImgUrl,
		Description: entityBook.Description,
		RentStatus:  int32(entityBook.RentStatus),
	}
}
