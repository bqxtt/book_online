module github.com/bqxtt/book_online/rpc

go 1.14

require (
	github.com/bqxtt/book_online/book v0.0.0-20210110154121-986ea162dee3
	github.com/bqxtt/book_online/rent v0.0.0-00010101000000-000000000000
	github.com/bqxtt/book_online/user v0.0.0-20210110154121-986ea162dee3
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/bqxtt/book_online/rent => ../rent
