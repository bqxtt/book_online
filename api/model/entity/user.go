package entity

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	UserId   int64  `form:"user_id" json:"user_id" binding:"required"`
}
