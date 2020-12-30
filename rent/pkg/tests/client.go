package tests

// This package is used to test the user grpc server.

import (
	"context"
	"errors"
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/handler/adapter"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/sdk/base"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
)

type clientTyp struct {
	userpb.UserServiceClient

	ctx context.Context
}

const (
	address = "localhost:50001"
)

func (testClient *clientTyp) get(ctx context.Context, id int64) (*model.User, error) {
	r, err := testClient.GetInfo(ctx, &userpb.GetInfoRequest{
		UserId: id,
	})
	if err = parseError(r.GetReply(), err); err != nil {
		return nil, fmt.Errorf("failed to get info: %v", err)
	}

	return adapter.AdaptToModelUser(r.GetUser())
}

func (testClient *clientTyp) login(ctx context.Context, id int64, pwdDigest string) error {
	r, err := testClient.Login(ctx, &userpb.LoginRequest{
		UserAuth: &userpb.UserAuth{
			UserId:    id,
			PwdDigest: pwdDigest,
		},
	})
	if err = parseError(r.GetReply(), err); err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}

	return nil
}

func (testClient *clientTyp) update(ctx context.Context, id int64, name string) (*model.User, error) {
	r, err := testClient.UpdateInfo(ctx, &userpb.UpdateInfoRequest{
		UserInfo: &userpb.User{
			UserId: id,
			Name:   name,
		},
	})
	if err = parseError(r.GetReply(), err); err != nil {
		return nil, fmt.Errorf("failed to update: %v", err)
	}

	return adapter.AdaptToModelUser(r.GetResultUserInfo())
}

func (testClient *clientTyp) register(ctx context.Context, id int64, pwdDigest string) error {
	r, err := testClient.Register(ctx, &userpb.RegisterRequest{
		UserAuth: &userpb.UserAuth{
			PwdDigest: pwdDigest,
			UserId:    id,
		},
	})

	if err = parseError(r.GetReply(), err); err != nil {
		return fmt.Errorf("failed to register: %v", err)
	}

	return nil
}

func parseError(reply *base.BaseReply, err error) error {
	if err != nil {
		return err
	}

	if reply.GetStatus() != base.REPLY_STATUS_SUCCESS {
		return errors.New(reply.GetMessage())
	}

	return nil
}
