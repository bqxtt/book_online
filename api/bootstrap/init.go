package bootstrap

import (
	"github.com/bqxtt/book_online/api/router"
	"github.com/bqxtt/book_online/rpc/clients/rpc_book"
	"github.com/bqxtt/book_online/rpc/clients/rpc_user"
)

func Init() {
	initRpc()
	router.Init()
}

func initRpc() {
	rpc_user.Init()
	rpc_book.Init()
}
