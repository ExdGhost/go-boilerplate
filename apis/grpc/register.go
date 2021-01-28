package grpc

import (
	pingpb "go-boilerplate/apis/grpc/generated/ping"
	pb "go-boilerplate/apis/grpc/generated/user"
	"go-boilerplate/apis/grpc/ping"
	"go-boilerplate/apis/grpc/user"
	"go-boilerplate/shared"

	"google.golang.org/grpc"
)

func registerService(server *grpc.Server, deps *shared.Deps) {

	userService := user.NewUserService(deps.Config, deps.Database, deps.Apm, deps.HTTPRequester, deps.GrpcConn)
	pingService := ping.NewPingService()

	// Bind the RPC services to the grpc server
	pb.RegisterUserServiceServer(server, userService)

	pingpb.RegisterPingServiceServer(server, pingService)
}
