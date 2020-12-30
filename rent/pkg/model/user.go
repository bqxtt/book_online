package model

import (
	"github.com/bqxtt/book_online/user/pkg/common"
	"time"
)

type User struct {
	ID        int64
	UserID    int64
	Name      string
	AvatarURL string
	Role      common.UserRole

	CreatedAt time.Time
}

type UserAuth struct {
	UserID    int64
	PwdDigest string
}
