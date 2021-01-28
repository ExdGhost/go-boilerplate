package config

import (
	"io/ioutil"

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
