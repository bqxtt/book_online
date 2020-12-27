package model

import "time"

type User struct {
	ID        int64
	UserID    int64
	Name      string
	CreatedAt time.Time
}

type UserAuth struct {
	UserID    int64
	PwdDigest string
}
