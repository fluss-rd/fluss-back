package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	env "github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

var (
	// ErrCouldNotReadFile could not read file
	ErrCouldNotReadFile = errors.New("could not read file")
	// ErrMissingValue missing value
	ErrMissingValue = errors.New("missing value")
)

// AppConfig defines the configuration
type AppConfig struct {
	Environment string `yaml:"environment" validate:"required" env:"APPCONFIG_ENVIRONMENT,required"`
	Port        string `yaml:"port" validate:"required" env:"APPCONFIG_PORT,required"`

	DatabaseConfig struct {
		DatabaseType string `yaml:"databaseType" validate:"required" env:"APPCONFIG_REPOSITORYCONFIG_DATABASETYPE,required"`
		Connection   string `yaml:"connection" validate:"required" env:"APPCONFIG_REPOSITORYCONFIG_CONNECTION,required"`
		DatabseName  string `yaml:"databaseName" validate:"required" env:"APPCONFIG_REPOSITORYCONFIG_DATABASENAME,required"`
	} `yaml:"databaseConfig"`
}

func GetConfig(filename string) (*AppConfig, error) {
	if filename == "" {
		return GetConfigFromEnv()
	}

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%w:%s", ErrCouldNotReadFile, err.Error())
	}

	return getConfigFromFileContent(fileContent)
}

func getConfigFromFileContent(fileContent []byte) (*AppConfig, error) {
	fileContent = []byte(os.ExpandEnv(string(fileContent)))

	config := &AppConfig{}
	err := yaml.Unmarshal(fileContent, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfigFromEnv() (*AppConfig, error) {
	var config AppConfig

	err := env.Parse(&config)
	if err != nil && strings.Contains(err.Error(), "env: required environment variable") && strings.Contains(err.Error(), "is not set") {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &config, nil
}
