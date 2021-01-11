package clients

import (
	"context"
	"github.com/bqxtt/book_online/rent/pkg/sdk/include/bookpb"
	"github.com/bqxtt/book_online/rent/pkg/utils"
	"google.golang.org/grpc"
)

type BookClient struct {
	ctx       context.Context
	rpcClient bookpb.BookServiceClient
}

func NewBookClient() (*BookClient, error) {
	conn, err := grpc.Dial("139.9.140.136:50002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	rpcClient := bookpb.NewBookServiceClient(conn)
	ctx := context.Background()

	return &BookClient{
		ctx:       ctx,
		rpcClient: rpcClient,
	}, nil
}

func (b *BookClient) GetBookById(bookId int64) (*bookpb.Book, error) {
	reply, err := b.rpcClient.GetBookById(b.ctx, &bookpb.GetBookByIdRequest{
		BookId: bookId,
	})
	if err != nil {
		return nil, err
	}
	if err = utils.ParsePbReplyError(reply.BaseReply); err != nil {
		return nil, err
	}

	return reply.Book, nil
}

func (b *BookClient) CheckBookBorrowed(bookId int64) (bool, error) {
	reply, err := b.rpcClient.CheckBookBorrowed(b.ctx, &bookpb.CheckBookBorrowedRequest{
		BookId: bookId,
	})
	if err != nil {
		return false, err
	}
	if err = utils.ParsePbReplyError(reply.BaseReply); err != nil {
		return false, err
	}

	return reply.Borrowed, nil
}

func (b *BookClient) SetBookBorrowed(bookId int64) error {
	reply, err := b.rpcClient.SetBookBorrowed(b.ctx, &bookpb.SetBookBorrowedRequest{
		BookId: bookId,
	})
	if err != nil {
		return err
	}
	if err = utils.ParsePbReplyError(reply.BaseReply); err != nil {
		return err
	}

	return nil
}

func (b *BookClient) SetBookReturned(bookId int64) error {
	reply, err := b.rpcClient.SetBookReturned(b.ctx, &bookpb.SetBookReturnedRequest{
		BookId: bookId,
	})
	if err != nil {
		return err
	}
	if err = utils.ParsePbReplyError(reply.BaseReply); err != nil {
		return err
	}

	return nil
}
