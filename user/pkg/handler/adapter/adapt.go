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
		UserID:     user.UserId,
		Name:       user.Name,
		Department: user.Department,
		Class:      user.Class,
		Motto:      user.Motto,
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
		UserId:     user.UserID,
		Name:       user.Name,
		Department: user.Department,
		Class:      user.Class,
		Motto:      user.Motto,
		CreatedAt:  createdAt,
	}, nil
}

func AdaptToPbUsers(users []*model.User) ([]*userpb.User, error) {
	pbUsers := make([]*userpb.User, 0, len(users))
	for _, user := range users {
		pbUser, err := AdaptToPbUser(user)
		if err != nil {
			return nil, err
		}

		pbUsers = append(pbUsers, pbUser)
	}

	return pbUsers, nil
}
