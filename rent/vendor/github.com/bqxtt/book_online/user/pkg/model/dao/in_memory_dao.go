package dao

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/model"
)

type InMemoryDao struct {
	userMap     map[int64]*model.User
	userAuthMap map[int64]*model.UserAuth
}

// randomly ordered...
func (in *InMemoryDao) ListUsersByLimitOffset(limit int64, offset int64) (results []*model.User, err error) {
	var cnt int64 = 1
	for _, user := range in.userMap {
		if cnt <= offset {
			continue
		}
		if cnt > offset+limit {
			break
		}

		results = append(results, user)
		cnt++
	}

	return results, nil
}

func (in *InMemoryDao) CountUsers() (result int64, err error) {
	return int64(len(in.userMap)), nil
}

func NewInMemoryDao() (*InMemoryDao, error) {
	return &InMemoryDao{
		userMap:     make(map[int64]*model.User),
		userAuthMap: make(map[int64]*model.UserAuth),
	}, nil
}

func (in *InMemoryDao) CreateUser(user *model.User, userAuth *model.UserAuth) (*model.User, *model.UserAuth, error) {
	_, ok := in.userMap[user.UserID]
	if ok {
		return nil, nil, fmt.Errorf("user id %v already exists", user.UserID)
	}
	// TODO: if we need to check both the tables?
	_, ok = in.userAuthMap[user.UserID]
	if ok {
		return nil, nil, fmt.Errorf("user id %v already exists", user.UserID)
	}

	in.userMap[user.UserID] = user
	in.userAuthMap[user.UserID] = userAuth
	return in.userMap[user.UserID], in.userAuthMap[user.UserID], nil
}

func (in *InMemoryDao) GetUserByUserId(userId int64) (result *model.User, exists bool, err error) {
	user, ok := in.userMap[userId]
	if ok {
		return user, true, nil
	}

	return nil, false, nil
}

func (in *InMemoryDao) GetUserAuthByUserId(userId int64) (result *model.UserAuth, exists bool, err error) {
	ua, ok := in.userAuthMap[userId]
	if ok {
		return ua, true, nil
	}

	return nil, false, nil
}

func (in *InMemoryDao) UpdateUserByUserId(userId int64, user *model.User) (result *model.User, err error) {
	_, ok := in.userMap[userId]
	if !ok {
		return nil, fmt.Errorf("user with user id %v not found", userId)
	}

	in.userMap[userId] = user
	return in.userMap[userId], nil
}

func (in *InMemoryDao) DeleteUserByUserId(userId int64) error {
	in.userMap[userId] = nil
	return nil
}
