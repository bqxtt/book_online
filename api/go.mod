module github.com/bqxtt/book_online/api

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bqxtt/book_online/book v0.0.0-20201230134400-cebe5ed17b7a
	github.com/bqxtt/book_online/rpc v0.0.0-00010101000000-000000000000
	github.com/bqxtt/book_online/user v0.0.0-20201230134400-cebe5ed17b7a // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
)

replace github.com/bqxtt/book_online/rpc => ../rpc
