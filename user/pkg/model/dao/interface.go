package dao

import (
	"github.com/bqxtt/book_online/user/pkg/common"
	"github.com/bqxtt/book_online/user/pkg/model"
)

type UserDAOInterface interface {
	// TODO: add proxy/gORM conns
	CreateUser(user *model.User, userAuth *model.UserAuth) error
	GetUserByUserId(userId int64) (*model.User, error)
}

func NewUserDAO(daoType common.DAOType) (UserDAOInterface, error) {
	switch daoType {
	// TODO:
	default:
		return NewInmemoryDao()
	}
}
