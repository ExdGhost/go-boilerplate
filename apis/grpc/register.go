package grpc

import (
	pingpb "go-boilerplate-api/apis/grpc/generated/ping"
	pb "go-boilerplate-api/apis/grpc/generated/user"
	"go-boilerplate-api/apis/grpc/ping"
	"go-boilerplate-api/apis/grpc/user"
	"go-boilerplate-api/shared"

	"google.golang.org/grpc"
)

func registerService(server *grpc.Server, deps *shared.Deps) {

	userService := user.NewUserService(deps.Config, deps.Database, deps.Apm, deps.HTTPRequester, deps.GrpcConn)
	pingService := ping.NewPingService()

	// Bind the RPC services to the grpc server
	pb.RegisterUserServiceServer(server, userService)

	pingpb.RegisterPingServiceServer(server, pingService)
}
