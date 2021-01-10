package dao

import (
	"fmt"
	"github.com/bqxtt/book_online/user/pkg/model"
	"testing"
	"time"
)

func TestDBDirectDAO_CreateUser(t *testing.T) {
	db, err := NewDBDirectDAO()
	if err != nil {
		fmt.Println(err)
		return
	}
	user := &model.User{
		ID:         0,
		UserID:     1234,
		Name:       "tcg",
		AvatarURL:  "sdaf",
		Department: "dfa",
		Class:      "dsf",
		Motto:      "的撒法",
		Role:       1,
		CreatedAt:  time.Now(),
	}
	userAuth := &model.UserAuth{
		UserID:    1234,
		PwdDigest: "1234",
	}
	_, _, err = db.CreateUser(user, userAuth)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

func TestDBDirectDAO_GetUserByUserId(t *testing.T) {
	db, _ := NewDBDirectDAO()
	user, exist, err := db.GetUserByUserId(1234)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user, exist)
	}
}
