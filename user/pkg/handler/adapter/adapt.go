package adapter

import (
	"github.com/bqxtt/book_online/user/pkg/model"
	"github.com/bqxtt/book_online/user/pkg/sdk/userpb"
)

type RpcUser struct {
	User     *userpb.User
	UserAuth *userpb.UserAuth
}

func (rpcUser *RpcUser) ToModelUser() (*model.User, *model.UserAuth) {
	return &model.User{
			UserID: rpcUser.User.UserId,
			Name:   rpcUser.User.Name,
		}, &model.UserAuth{
			UserID:    rpcUser.User.UserId,
			PwdDigest: rpcUser.UserAuth.PwdDigest,
		}
}
