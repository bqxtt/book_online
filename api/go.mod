module github.com/bqxtt/book_online/api

go 1.14

require (
	github.com/bqxtt/book_online/rpc v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)

replace github.com/bqxtt/book_online/rpc => ../rpc
