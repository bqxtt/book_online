package common

// TODO: move to envs
const (
	DefaultUserGRPCHandlerPort = 50001
	DefaultUserAvatarURL       = "http://localhost:8080/image/avatar.jpg"
	DefaultUserPageSize        = 10
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
