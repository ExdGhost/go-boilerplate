package user

import (
	"net/http"

	"go-boilerplate-api/apis/http/utils"
	"go-boilerplate-api/apm"
	"go-boilerplate-api/config"
	"go-boilerplate-api/pkg/clients/db"
	grpcPkg "go-boilerplate-api/pkg/clients/grpc"
	httpPkg "go-boilerplate-api/pkg/clients/http"
	user "go-boilerplate-api/pkg/user"
	"go-boilerplate-api/pkg/user/favourite"
	"go-boilerplate-api/pkg/user/rating"
	userRepo "go-boilerplate-api/pkg/user/repo"

	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-errors"
)

// Service contains the methods required to perfom operation's on users
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

// GetAll godoc
// @Summary Get all users
// @Description get user all objects
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} user.User
// @Failure 400 {object} utils.Error
// @Router /users [get]
func (service *Service) getAll(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	users, err := service.user.GetAll()
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetOne godoc
// @Summary Get an user
// @Description get user object by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} user.User
// @Failure 400 {object} utils.Error
// @Router /users/{id} [get]
func (service *Service) getOne(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	userID := ctx.Param("userID")
	users, err := service.user.GetOne(userID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetWithInfo godoc
// @Summary Get an user with rating info
// @Description get user object with rating info by user ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} user.User
// @Failure 400 {object} utils.Error
// @Router /users/{id}/rating [get]
func (service *Service) getWithInfo(ctx *gin.Context) {
	var err error
	var reference map[string]interface{}
	defer utils.HandleError(ctx, &err, &reference)

	userID := ctx.Param("userID")
	reference = map[string]interface{}{"userID": userID}
	users, err := service.user.GetWithInfo(userID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Insert godoc
// @Summary Insert user
// @Description Insert an user object
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.User true "Insert user"
// @Success 200 {object} user.User
// @Failure 400 {object} utils.Error
// @Router /users [post]
func (service *Service) insert(ctx *gin.Context) {
	var err error
	var user user.User
	defer utils.HandleError(ctx, &err, &user)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		err = errors.NewBadRequest("Could not bind request to model").SetCode("APIS.HTTP.USER.REQUEST_BIND_FAILED")
		return
	}

	err = service.user.Insert(user)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, user)
}
