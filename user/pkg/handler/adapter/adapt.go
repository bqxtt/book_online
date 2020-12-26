package adapter

import (
	"errors"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
	"github.com/golang/protobuf/ptypes"
)

func AdaptToModelUser(user *userpb.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("nil user pb message")
	}

	return &model.User{
		UserID: user.UserId,
		Name:   user.Name,
	}, nil
}

func AdaptToModelUserAuth(ua *userpb.UserAuth) (*model.UserAuth, error) {
	if ua == nil {
		return nil, errors.New("nil user auth pb message")
	}

	return &model.UserAuth{
		UserID:    ua.UserId,
		PwdDigest: ua.PwdDigest,
	}, nil
}

func AdaptToPbUser(user *model.User) (*userpb.User, error) {
	if user == nil {
		return nil, errors.New("nil user model")
	}

	createdAt, err := ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &userpb.User{
		UserId:    user.UserID,
		Name:      user.Name,
		CreatedAt: createdAt,
	}, nil
}
