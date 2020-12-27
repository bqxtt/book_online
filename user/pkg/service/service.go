package service

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/common"
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/model/dao"
)

type UserService struct {
	userDAO dao.UserDAOInterface
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

func (s *UserService) CreateUser(user *model.User, userAuth *model.UserAuth) (*model.User, error) {
	result, _, err := s.userDAO.CreateUser(user, userAuth)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return result, nil
}

func (s *UserService) GetUserByUserId(userId int64) (*model.User, error) {
	user, exists, err := s.userDAO.GetUserByUserId(userId)
	if err != nil {
		return nil, err
	}
	if (user == nil) || !exists {
		return nil, fmt.Errorf("user with user id %v not found", userId)
	}

	return user, nil
}

func (s *UserService) AuthenticateUser(userAuth *model.UserAuth) error {
	ua, exists, err := s.userDAO.GetUserAuthByUserId(userAuth.UserID)
	if err != nil {
		return err
	}
	if (ua == nil) || !exists {
		return fmt.Errorf("user with user id %v not found", userAuth.UserID)
	}

	if ua.PwdDigest != userAuth.PwdDigest {
		return fmt.Errorf("failed to authenticate user with user id %v: wrong password", userAuth.UserID)
	}

	return nil
}

func (s *UserService) UpdateUser(user *model.User) (*model.User, error) {
	resultUserInfo, err := s.userDAO.UpdateUserByUserId(user.UserID, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	if resultUserInfo == nil {
		// not reach
		return nil, fmt.Errorf("unexpected error when updating user")
	}

	return resultUserInfo, nil
}
