package common

// TODO: move to envs
const (
	DefaultUserGRPCHandlerPort = 50003
	DefaultRecordPageSize      = 10
	DefaultWaitBookRPCTime     = 3  // time unit: second
	DefaultBookKeptInterval    = 30 // time unit: day
)

type DAOType string

const (
	DAOTypeDefault      DAOType = "default"
	DAOTypeDBDirectConn DAOType = "straight-db"
	DAOTypeDBProxy      DAOType = "db-proxy"
	// mainly for tests
	DAOTypeInMemory DAOType = "in-memory"
	DAOTypeUnknown  DAOType = "unknown"
)

type UserRole int32

const (
	UserRoleUnknown = 0
	UserRoleNormal  = 1
	UserRoleAdmin   = 2
)
