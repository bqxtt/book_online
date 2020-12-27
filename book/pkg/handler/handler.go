package handler

import (
	"context"
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
		}, nil
	}
	modelBooks, err := bh.bookService.GetAllBooks()
	if err != nil {
		return &bookpb.GetAllBooksResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service get all books error, err: %v", err),
		}, nil
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
		}, nil
	}
	page, pageSize := request.Page, request.PageSize
	if page < 0 || pageSize < 0 {
		return &bookpb.GetBooksByPageResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "page or page size is incorrect"),
		}, nil
	}
	modelBooks, err := bh.bookService.GetBooksByPage(page, pageSize)
	if err != nil {
		return &bookpb.GetBooksByPageResponse{
			Books:     nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service get book by page error, err: %v", err),
		}, nil
	}
	var rpcBooks []*bookpb.Book
	for _, modelBook := range modelBooks {
		rpcBooks = append(rpcBooks, adapter.ModelBookToRpcBook(modelBook))
	}
	return &bookpb.GetBooksByPageResponse{
		Books:     rpcBooks,
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}

func (bh *BookHandler) GetBookById(ctx context.Context, request *bookpb.GetBookByIdRequest) (*bookpb.GetBookByIdResponse, error) {
	if request == nil {
		return &bookpb.GetBookByIdResponse{
			Book:      nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get book by id request is nil"),
		}, nil
	}
	modelBook, err := bh.bookService.GetBookById(request.BookId)
	if err != nil {
		return &bookpb.GetBookByIdResponse{
			Book:      nil,
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "get book by id error, err: %v", err),
		}, nil
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
		}, nil
	}
	if request.Book == nil {
		return &bookpb.CreateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "create book request book is nil"),
		}, nil
	}
	modelBook := adapter.RpcBookToModelBook(request.Book)
	result, err := bh.bookService.CreateBook(modelBook)
	if err != nil {
		return &bookpb.CreateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service create book error, err: %v", err),
		}, nil
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
		}, nil
	}
	bookId, book := request.Id, request.Book
	if book == nil || bookId < 0 {
		return &bookpb.UpdateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "update book request book is nil or book id is incorrect"),
		}, nil
	}
	book.Id = bookId
	modelBook := adapter.RpcBookToModelBook(book)
	result, err := bh.bookService.UpdateBook(modelBook)
	if err != nil {
		return &bookpb.UpdateBookResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service update book error, err: %v", err),
		}, nil
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
		}, nil
	}
	bookId := request.Id
	if bookId < 0 {
		return &bookpb.DeleteBookByIdResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "delete book by id book id is incorrect"),
		}, nil
	}
	err := bh.bookService.DeleteBookById(bookId)
	if err != nil {
		return &bookpb.DeleteBookByIdResponse{
			BaseReply: utils.PbReplyf(base.REPLY_STATUS_FAILURE, "book service delete book error. err: %v", err),
		}, nil
	}
	return &bookpb.DeleteBookByIdResponse{
		BaseReply: utils.NewDefaultSuccessReply(),
	}, nil
}
