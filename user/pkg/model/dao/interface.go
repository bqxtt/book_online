package dao

import (
	"github.com/bqxtt/book_online/user/pkg/common"
	"github.com/bqxtt/book_online/user/pkg/model"
)

type UserDAOInterface interface {
	// TODO: add proxy/gORM conns
	CreateUser(user *model.User, userAuth *model.UserAuth) (resultUser *model.User, resultUA *model.UserAuth, err error)
	GetUserByUserId(userId int64) (result *model.User, exists bool, err error)
	GetUserAuthByUserId(userId int64) (result *model.UserAuth, exists bool, err error)
	UpdateUserByUserId(userId int64, user *model.User) (result *model.User, err error)
}

func NewUserDAO(daoType common.DAOType) (UserDAOInterface, error) {
	switch daoType {
	// TODO:
	default:
		return NewInMemoryDao()
	}
}
