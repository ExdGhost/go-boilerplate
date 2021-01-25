package config

import (
	log "go-boilerplate-api/pkg/utils/logger"
	"io/ioutil"

	"go-boilerplate-api/pkg/utils/helpers"

	"gopkg.in/yaml.v3"
)

var configModel *Config

// NewConfig gets the configuration based on the environment passed
func NewConfig(env string) (IConfig, error) {

	configFile := "config/tier/" + env + ".yml"
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	if env != "development" && env != "staging" && env != "docker" && env != "testing" {
		cfg, err := InitRemoteKvStore(env)
		if err != nil {
			return nil, err
		}

		// anonymous function to fetch from remote k-v store
		getRemoteValue := func(key string) (string, error) {
			value, err := cfg.GetKey(key)
			if err != nil {
				log.Fatal("CONFIG.KEY.NOT.FOUND", "Key Not Found", log.Priority1, nil, map[string]interface{}{key: err.Error()})
				return "", err
			}
			return value, err
		}

		// Binds config based on ccms and present values
		err = helpers.BindConfig(bytes, &configModel, "config", getRemoteValue)
		if err != nil {
			return nil, err
		}

		return &IConfigModel{model: configModel}, nil
	}

	err = yaml.Unmarshal(bytes, &configModel)
	if err != nil {
		return nil, err
	}

	// Returns
	return &IConfigModel{model: configModel}, nil
}

// Get implements the interface function for IConfig
func (ic *IConfigModel) Get() *Config {
	return ic.model
}

// InitRemoteKvStore ...
func InitRemoteKvStore(env string) (RemoteStoreInterface, error) {
	remoteConfig:= newRemoteStore(env)
	err := remoteConfig.Init()
	if err != nil {
		return nil, err
	}
	return remoteConfig, nil
}

// dummyFuncs (should be a seperate package)

// RemoteStoreInterface ...
type  RemoteStoreInterface interface{
	Init() error
	GetKey(string) (string, error)
}

// RemoteStore ...
type RemoteStore struct {
 Env string
}

func newRemoteStore(env string) RemoteStoreInterface {
	return &RemoteStore{Env : env}
}

// Init ...
func (rs *RemoteStore) Init() error{

	return nil
}

// GetKey ...
func (rs *RemoteStore) GetKey(key string)(string, error){

	return "",nil
}
