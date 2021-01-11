package common

// TODO: move to envs
const (
	DefaultUserGRPCHandlerPort = 50001
	DefaultUserAvatarURL       = "https://ftp.bmp.ovh/imgs/2021/01/bdd1426b103fd5d3.jpg"
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
