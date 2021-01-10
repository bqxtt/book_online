package dao

import (
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBDirectDAO struct {
	DB *gorm.DB
}

func NewDBDirectDAO() (*DBDirectDAO, error) {
	db, err := gorm.Open("mysql", "root:19981108@tcp(139.9.140.136:3307)/book_online?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		return nil, err
	}
	return &DBDirectDAO{
		DB: db,
	}, nil
}

func (DB *DBDirectDAO) CreateUser(user *model.User, userAuth *model.UserAuth) (resultUser *model.User, resultUA *model.UserAuth, err error) {
	result := DB.DB.Create(user)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	result = result.Create(userAuth)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	return user, userAuth, nil
}

func (DB *DBDirectDAO) GetUserByUserId(userId int64) (*model.User, bool, error) {
	user := &model.User{}
	result := DB.DB.Where("user_id = ?", userId).Find(user)
	if result.Error != nil {
		return nil, false, result.Error
	}
	//fmt.Println(user)
	return user, true, nil
}

func (DB *DBDirectDAO) GetUserAuthByUserId(userId int64) (*model.UserAuth, bool, error) {
	userAuth := &model.UserAuth{}
	result := DB.DB.Where("user_id = ?", userId).Find(userAuth)
	if result.Error != nil {
		return nil, false, result.Error
	}
	return userAuth, true, nil
}

func (DB *DBDirectDAO) UpdateUserByUserId(userId int64, user *model.User) (*model.User, error) {
	user.UserID = userId
	result := DB.DB.Model(user).Update(user)
	if result.Error != nil {
		return nil, result.Error
	}
	DB.DB.Model(user).Find(user)
	return user, nil
}

func (DB *DBDirectDAO) ListUsersByLimitOffset(limit int64, offset int64) ([]*model.User, error) {
	var users []*model.User
	result := DB.DB.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (DB *DBDirectDAO) CountUsers() (int64, error) {
	var count int64
	result := DB.DB.Model(&model.User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (DB *DBDirectDAO) DeleteUserByUserId(userId int64) error {
	result := DB.DB.Delete(&model.User{UserID: userId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
