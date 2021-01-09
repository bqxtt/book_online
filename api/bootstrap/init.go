package bootstrap

import (
	"github.com/bqxtt/book_online/api/router"
	"github.com/bqxtt/book_online/rpc/clients/rpc_book"
)

func Init() {
	initRpc()
	router.Init()
}

func initRpc() {
	//rpcs := os.Getenv("rpc")
	//if rpcs == "" {
	//	return
	//}
	//rpc := strings.Split(rpcs, ".")
	//funcs := make(map[string]func(), 0)
	//funcs["user"] = rpc_user.Init
	//funcs["book"] = rpc_book.Init
	//for _, r := range rpc {
	//	funcs[r]()
	//}
	//rpc_user.Init()
	rpc_book.Init()
}
