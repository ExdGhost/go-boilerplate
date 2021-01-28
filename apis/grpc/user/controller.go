package user

import (
	"context"

	pb "go-boilerplate/apis/grpc/generated/user"
	"go-boilerplate/apis/grpc/utils"
	"go-boilerplate/config"
	"go-boilerplate/internal/favourite"
	"go-boilerplate/internal/rating"
	"go-boilerplate/internal/user"
	userRepo "go-boilerplate/internal/user/repo"
	"go-boilerplate/pkg/apm"
	"go-boilerplate/pkg/clients/db"
	grpcPkg "go-boilerplate/pkg/clients/grpc"
	httpPkg "go-boilerplate/pkg/clients/http"
	pkgUtils "go-boilerplate/pkg/utils"
)

// Service contains the methods required to perfom operation's on users (proto definition)
type Service struct {
	user user.UsersInterface
}

// NewUserService Create a new instance of a Service with the given dependencies.
func NewUserService(conf config.IConfig, db *db.Instances, apm apm.HandlerInterface, httpReq httpPkg.IRequest, grpcConn grpcPkg.IGrpcConnections) *Service {
	userRating := rating.NewRating(conf, httpReq)
	userFavourites := favourite.NewFavourite(conf, grpcConn)
	userRepo := userRepo.NewUserRepo(conf, db)
	userService := user.NewUser(conf, userRepo, apm, userRating, userFavourites)

	return &Service{user: userService}
}

// GetAll gets all users
func (service *Service) GetAll(ctx context.Context, req *pb.UserGetRequest) (res *pb.Users, err error) {
	defer utils.HandleError(&err)

	users, err := service.user.GetAll()
	if err != nil {
		return nil, err
	}

	res = &pb.Users{}
	err = pkgUtils.Bind(users, &res.Users)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne gets one users
func (service *Service) GetOne(ctx context.Context, req *pb.UserGetRequest) (res *pb.User, err error) {
	var userReq = user.User{}
	defer utils.HandleError(&err, &userReq)

	// Need to decode to user.User since User is an embedded struct
	err = pkgUtils.Bind(req, &userReq)
	if err != nil {
		return nil, err
	}

	users, err := service.user.GetOne(userReq.ID)
	if err != nil {
		return nil, err
	}

	res = &pb.User{}
	err = pkgUtils.Bind(users, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Insert stores a user in the datastore
func (service *Service) Insert(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	var userReq = user.User{}
	defer utils.HandleError(&err, &userReq)

	// Need to decode to user.User since User is an embedded struct
	err = pkgUtils.Bind(req, &userReq)
	if err != nil {
		return nil, err
	}

	err = service.user.Insert(userReq)
	if err != nil {
		return nil, err
	}

	err = pkgUtils.Bind(userReq, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetWithInfo gets a user from the database along with rating and favourites
func (service *Service) GetWithInfo(ctx context.Context, req *pb.UserGetRequest) (res *pb.User, err error) {
	var userReq = user.User{}
	defer utils.HandleError(&err, &userReq)

	// Need to decode to user.User since User is an embedded struct
	err = pkgUtils.Bind(req, &userReq)
	if err != nil {
		return nil, err
	}

	users, err := service.user.GetWithInfo(userReq.ID)
	if err != nil {
		return nil, err
	}

	res = &pb.User{}
	err = pkgUtils.Bind(users, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
