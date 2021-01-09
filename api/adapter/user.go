package adapter

import (
	"github.com/bqxtt/book_online/api/model/entity"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user/userpb"
	"strconv"
)

func RpcUserToEntityUser(rpcUser *userpb.User) *entity.User {
	role := ""
	if rpcUser.Role == 1 {
		role = "user"
	} else if rpcUser.Role == 2 {
		role = "admin"
	} else {
		role = "unknown"
	}
	return &entity.User{
		UserId:     strconv.FormatInt(rpcUser.UserId, 10),
		Name:       rpcUser.Name,
		AvatarUrl:  rpcUser.AvatarUrl,
		Role:       role,
		Department: rpcUser.Department,
		Class:      rpcUser.Class,
		Motto:      rpcUser.Motto,
	}
}

func EntityUserToRpcUser(entityUser *entity.User) *userpb.User {
	userId, _ := strconv.ParseInt(entityUser.UserId, 10, 64)
	var role userpb.USER_ROLE
	if entityUser.Role == "user" {
		role = userpb.USER_ROLE_NORMAL
	} else if entityUser.Role == "admin" {
		role = userpb.USER_ROLE_ADMIN
	}
	return &userpb.User{
		UserId:     userId,
		Name:       entityUser.Name,
		AvatarUrl:  entityUser.AvatarUrl,
		Role:       role,
		Department: entityUser.Department,
		Class:      entityUser.Class,
		Motto:      entityUser.Motto,
	}
}
