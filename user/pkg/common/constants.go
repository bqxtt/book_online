package common

const (
	DefaultUserGRPCHandlerPort = 50001
)

type DAOType string
const (
	DAOTypeDefault DAOType = "default"
	DAOTypeDBDirectConn DAOType = "straight-db"
	DAOTypeDBProxy DAOType = "db-proxy"
	// mainly for tests
	DAOTypeInMemory DAOType = "in-memory"
	DAOTypeUnknown DAOType = "unknown"
)
