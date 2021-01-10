package common

const (
	DefaultBookGRPCHandlerPort = 50002
)

type DAOType string

const (
	DAOTypeDefault      DAOType = "default"
	DAOTypeDBDirectConn DAOType = "db-direct"
	DAOTypeDBProxy      DAOType = "db-proxy"
	// mainly for tests
	DAOTypeInMemory DAOType = "in-memory"
	DAOTypeUnknown  DAOType = "unknown"
)

type BookStatus int8

const (
	BOOK_AVALIABLE BookStatus = 1
	BOOK_DELETED   BookStatus = 2
)

type BookRentStatus int8

const (
	BOOK_RETURNED BookRentStatus = 1
	BOOK_BORROWED BookRentStatus = 2
)
