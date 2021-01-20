package http

import (
	"sync"

	"go-boilerplate-api/apis/http/ping"
	"go-boilerplate-api/apis/http/swagger"
	httpUser "go-boilerplate-api/apis/http/user"
	"go-boilerplate-api/apis/middleware"
	"go-boilerplate-api/pkg/apm"
	log "go-boilerplate-api/pkg/utils/logger"
	"go-boilerplate-api/shared"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer starts the http server using the dependencies passed to it.
// It also initializes the routes
func StartServer(deps *shared.Deps, wg *sync.WaitGroup, fatalError chan error) error {
	address := deps.Config.Get().Server.HTTP.Address

	gin.SetMode(gin.DebugMode) // ToDo chage to release mode before prod
	router := gin.Default()
	// Injects apm to trace http requests in gin
	router.Use(middleware.ApmMiddleware(apm.APM))
	// adds cors support
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowMethods("GET", "DELETE", "POST", "PUT")
	router.Use(cors.New(config))
	// Adds panic handler as a middleware
	router.Use(middleware.HandlePanic)

	// Initializes Ping routes
	ping.NewPingRoute(router)
	// Initialize all the routes
	httpUser.NewUserRoute(router, deps)
	// Initialize all swagger routes
	swagger.NewSwaggerRoutes(router)

	// Logs server start
	log.Debug("HTTP Server listening on : " + address + ", Version: " + shared.VERSION)

	// Start the server
	err := router.Run(address)
	if err != nil {
		fatalError <- err
	}

	// Go routine finished (can also be deferred)
	wg.Done()

	return nil
}
