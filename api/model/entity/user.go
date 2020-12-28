package entity

type User struct {
	UserId int64  `form:"user_id" json:"user_id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
}

type UserAuth struct {
	UserId   int64  `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
