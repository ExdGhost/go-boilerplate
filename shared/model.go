package shared

import (
	"go-boilerplate/config"
	"go-boilerplate/pkg/apm"
	"go-boilerplate/pkg/clients/db"
	grpcPkg "go-boilerplate/pkg/clients/grpc"
	httpPkg "go-boilerplate/pkg/clients/http"
)

// VERSION keeps the version no. (commit id) for global use
var VERSION string

// Deps ... is a shared dependencies struct that contains common singletons
type Deps struct {
	Config        config.IConfig
	Database      *db.Instances
	GrpcConn      grpcPkg.IGrpcConnections
	HTTPRequester httpPkg.IRequest
	Apm           apm.HandlerInterface
}
