package model

import (
	"github.com/bqxtt/book_online/user/pkg/common"
	"time"
)

type User struct {
	ID         int64
	UserID     int64 `gorm:"primary_key"`
	Name       string
	AvatarUrl  string
	Department string
	Class      string
	Motto      string
	Role       common.UserRole

	CreatedAt time.Time
}

type UserAuth struct {
	UserID    int64
	PwdDigest string
}
