package service

import (
	"github.com/bqxtt/book_online/user/pkg/common"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/model/dao"
)

type UserService struct {
	userDAO dao.UserDAOInterface
}

func (s *UserService) CreateUser(user *model.User, userAuth *model.UserAuth) error {
	return s.userDAO.CreateUser(user, userAuth)
}

func (s *UserService) GetUserByUserId(userId int64) (*model.User, error) {
	return s.userDAO.GetUserByUserId(userId)
}

func NewUserService() (*UserService, error) {
	ud, err := dao.NewUserDAO(common.DAOTypeDefault)
	if err != nil {
		return nil, err
	}
	return &UserService{
		userDAO: ud,
	}, nil
}
