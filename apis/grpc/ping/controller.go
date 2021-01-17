package ping

import (
	"context"
	pb "go-boilerplate-api/apis/grpc/generated/ping"
	"go-boilerplate-api/shared"
)

// PingInterface ...
type PingInterface interface {
	Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error)
}

// Service contains the methods required to perfom operation's on ping (proto definition)
type Service struct{}

// NewPingService ...
func NewPingService() PingInterface {
	return &Service{}
}

// pongResponse is a constant used for sending response message.
const pongResponse string = "PONG"

// Ping ...
func (*Service) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	res := pb.PingResponse{Message: pongResponse, App: "go-boilerplate-api", CommitId: shared.VERSION}
	return &res, nil
}
