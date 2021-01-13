package handler

import (
	"context"
	"fmt"
	"github.com/bqxtt/book_online/book/pkg/handler/adapter"
	"github.com/bqxtt/book_online/book/pkg/sdk/base"
	"github.com/bqxtt/book_online/book/pkg/sdk/bookpb"
	"github.com/bqxtt/book_online/book/pkg/service"
	"github.com/bqxtt/book_online/book/pkg/utils"
)

type BookHandler struct {
	bookpb.UnimplementedBookServiceServer
	bookService *service.BookService
}

func NewBookHandler() (*BookHandler, error) {
	bs, err := service.NewBookService()
	return &BookHandler{
		bookService: bs,
	}, err
}

func (bh *BookHandler) GetAllBooks(ctx context.Context, request *bookpb.GetAllBooksRequest) (*bookpb.GetAllBooksResponse, error) {
	if request == nil {
		return &bookpb.GetAllBooksResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get all books request is nil"),
		}, fmt.Errorf("get all books request is nil")
	}
	modelBooks, err := bh.bookService.GetAllBooks()
	if err != nil {
		return &bookpb.GetAllBooksResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service get all books error, err: %v", err),
		}, err
	}
	var rpcBooks []*bookpb.Book
	for _, modelBook := range modelBooks {
		rpcBooks = append(rpcBooks, adapter.ModelBookToRpcBook(modelBook))
	}
	return &bookpb.GetAllBooksResponse{
		Books:     rpcBooks,
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) GetBooksByPage(ctx context.Context, request *bookpb.GetBooksByPageRequest) (*bookpb.GetBooksByPageResponse, error) {
	if request == nil {
		return &bookpb.GetBooksByPageResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get books by page request is nil"),
		}, fmt.Errorf("get books by page request is nil")
	}
	page, pageSize := request.Page, request.PageSize
	if page < 1 || pageSize < 0 {
		return &bookpb.GetBooksByPageResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "page or page size is incorrect"),
		}, fmt.Errorf("page or page size is incorrect")
	}
	modelBooks, pageInfo, err := bh.bookService.GetBooksByPage(page, pageSize)
	if err != nil {
		return &bookpb.GetBooksByPageResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service get book by page error, err: %v", err),
		}, err
	}
	var rpcBooks []*bookpb.Book
	for _, modelBook := range modelBooks {
		rpcBooks = append(rpcBooks, adapter.ModelBookToRpcBook(modelBook))
	}
	return &bookpb.GetBooksByPageResponse{
		Books:      rpcBooks,
		TotalCount: pageInfo.TotalCount,
		TotalPages: pageInfo.TotalPages,
		BaseReply:  utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) GetBookById(ctx context.Context, request *bookpb.GetBookByIdRequest) (*bookpb.GetBookByIdResponse, error) {
	if request == nil {
		return &bookpb.GetBookByIdResponse{
			Book:      nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get book by id request is nil"),
		}, fmt.Errorf("get book by id request is nil")
	}
	modelBook, err := bh.bookService.GetBookById(request.BookId)
	if err != nil {
		return &bookpb.GetBookByIdResponse{
			Book:      nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get book by id error, err: %v", err),
		}, err
	}
	return &bookpb.GetBookByIdResponse{
		Book:      adapter.ModelBookToRpcBook(modelBook),
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) CreateBook(ctx context.Context, request *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {
	if request == nil {
		return &bookpb.CreateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "create book request is nil"),
		}, fmt.Errorf("create book request is nil")
	}
	if request.Book == nil {
		return &bookpb.CreateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "create book request book is nil"),
		}, fmt.Errorf("create book request book is nil")
	}
	modelBook := adapter.RpcBookToModelBook(request.Book)
	result, err := bh.bookService.CreateBook(modelBook)
	if err != nil {
		return &bookpb.CreateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service create book error, err: %v", err),
		}, err
	}
	return &bookpb.CreateBookResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
		Id:        result.ID,
	}, nil
}

func (bh *BookHandler) UpdateBook(ctx context.Context, request *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	if request == nil {
		return &bookpb.UpdateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "update book request is nil"),
		}, fmt.Errorf("update book request is nil")
	}
	bookId, book := request.Id, request.Book
	if book == nil || bookId < 0 {
		return &bookpb.UpdateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "update book request book is nil or book id is incorrect"),
		}, fmt.Errorf("update book request book is nil or book id is incorrect")
	}
	modelBook := adapter.RpcBookToModelBook(book)
	result, err := bh.bookService.UpdateBook(bookId, modelBook)
	if err != nil {
		return &bookpb.UpdateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service update book error, err: %v", err),
		}, err
	}
	return &bookpb.UpdateBookResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
		Id:        result.ID,
	}, nil
}

func (bh *BookHandler) DeleteBookById(ctx context.Context, request *bookpb.DeleteBookByIdRequest) (*bookpb.DeleteBookByIdResponse, error) {
	if request == nil {
		return &bookpb.DeleteBookByIdResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "delete book by id request is nil"),
		}, fmt.Errorf("delete book by id request is nil")
	}
	bookId := request.Id
	if bookId < 0 {
		return &bookpb.DeleteBookByIdResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "delete book by id book id is incorrect"),
		}, fmt.Errorf("delete book by id book id is incorrect")
	}
	err := bh.bookService.DeleteBookById(bookId)
	if err != nil {
		return &bookpb.DeleteBookByIdResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service delete book error. err: %v", err),
		}, err
	}
	return &bookpb.DeleteBookByIdResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) CheckBookBorrowed(ctx context.Context, request *bookpb.CheckBookBorrowedRequest) (*bookpb.CheckBookBorrowedResponse, error) {
	if request == nil {
		return &bookpb.CheckBookBorrowedResponse{
			Borrowed:  false,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "check book borrowed request is nil"),
		}, fmt.Errorf("check book borrowed request is nil")
	}
	borrowed, err := bh.bookService.CheckBookBorrowed(request.BookId)
	if err != nil {
		return &bookpb.CheckBookBorrowedResponse{
			Borrowed:  false,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service check book borrowed error, err: %v", err),
		}, err
	}
	return &bookpb.CheckBookBorrowedResponse{
		Borrowed:  borrowed,
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) SetBookBorrowed(ctx context.Context, request *bookpb.SetBookBorrowedRequest) (*bookpb.SetBookBorrowedResponse, error) {
	if request == nil {
		return &bookpb.SetBookBorrowedResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "set book borrowed request is nil"),
		}, fmt.Errorf("set book borrowed request is nil")
	}
	err := bh.bookService.SetBookBorrowed(request.BookId)
	if err != nil {
		return &bookpb.SetBookBorrowedResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service set book borrowed error, err: %v", err),
		}, err
	}
	return &bookpb.SetBookBorrowedResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) SetBookReturned(ctx context.Context, request *bookpb.SetBookReturnedRequest) (*bookpb.SetBookReturnedResponse, error) {
	if request == nil {
		return &bookpb.SetBookReturnedResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "set book returned request is nil"),
		}, fmt.Errorf("set book returned request is nil")
	}
	err := bh.bookService.SetBookReturned(request.BookId)
	if err != nil {
		return &bookpb.SetBookReturnedResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service set book returned error, err: %v", err),
		}, err
	}
	return &bookpb.SetBookReturnedResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}
