package entity

type User struct {
	UserId    string `form:"user_id" json:"user_id" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url"`
	Role      string `form:"role" json:"role"`
	Token     string `form:"token" json:"token"`
}

type UserAuth struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
