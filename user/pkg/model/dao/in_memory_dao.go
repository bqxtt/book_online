package dao

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/model"
)

type InMemoryDao struct {
	userMap map[int64]*model.User
}

func NewInMemoryDao() (*InMemoryDao, error) {
	return &InMemoryDao{
		userMap: make(map[int64]*model.User),
	}, nil
}

func (in *InMemoryDao) CreateUser(user *model.User, userAuth *model.UserAuth) error {
	_, ok := in.userMap[user.UserID]
	if ok {
		return fmt.Errorf("user id %v already used", user.UserID)
	}

	in.userMap[user.UserID] = user
	fmt.Printf("user created, user: %v", user)
	return nil
}

func (in *InMemoryDao) GetUserByUserId(userId int64) (*model.User, error) {
	user, ok := in.userMap[userId]
	if ok {
		return user, nil
	}

	return nil, nil
}
