package initiate

import (
	"go-boilerplate/apis"
	"go-boilerplate/config"
	"go-boilerplate/pkg/apm"
	"go-boilerplate/pkg/clients/db"
	grpcPkg "go-boilerplate/pkg/clients/grpc"
	httpPkg "go-boilerplate/pkg/clients/http"
	log "go-boilerplate/pkg/utils/logger"
	"go-boilerplate/shared"
)

// Initialize will initialize all the dependencies and the servers.
// Dependencies include config, external connections(grpc, http)
func Initialize() error {

	env, err := Env()
	if err != nil {
		return err
	}

	log.Info("Enviroment: " + env)

	// Sets apm env
	SetApmEnv(env)

	// Initializes APM based on environment
	if env != config.ENVDevelopment && env != config.ENVDocker && env != config.ENVStage {
		// init apm here TODO
	}

	// Gets config
	conf, err := config.NewConfig(env)
	if err != nil {
		return err
	}

	// Initializes the DB connections
	dbInstances, err := db.NewInstance(conf)
	if err != nil {
		return err
	}

	// Initializes the GRPC connections
	grpcCons, err := grpcPkg.NewConnections(conf)
	if err != nil {
		return err
	}

	// Initializes apm Handler
	handler := apm.NewApmHandler()

	// loads all common dependencies
	dependencies := shared.Deps{
		Config:        conf,
		Database:      dbInstances,
		GrpcConn:      grpcCons,
		HTTPRequester: httpPkg.NewRequest(),
		Apm:           handler,
	}

	// Initializes servers
	err = apis.InitServers(&dependencies)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
