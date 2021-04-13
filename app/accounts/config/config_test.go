package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfigFromFileContent(t *testing.T) {
	c := require.New(t)

	fileContent := "environment: dev\nport: 5000\ndatabaseConfig:\n  databaseType: mongoDB\n  connection: ${MONGO_CONNECTION}\n  databaseName: dbTest"

	appConfig, err := getConfigFromFileContent([]byte(fileContent))
	c.Nil(err)
	c.Equal("5000", appConfig.Port)
	c.Equal("dev", appConfig.Environment)
	c.Equal("mongoDB", appConfig.DatabaseConfig.DatabaseType)
}

func TestGetConfigFromEnv(t *testing.T) {
	c := require.New(t)

	os.Setenv("APPCONFIG_ENVIRONMENT", "test")
	os.Setenv("APPCONFIG_PORT", "5000")
	os.Setenv("APPCONFIG_REPOSITORYCONFIG_DATABASETYPE", "mongoDB")
	os.Setenv("APPCONFIG_REPOSITORYCONFIG_CONNECTION", "mongodb://mongodb0.example.com:27017")
	os.Setenv("APPCONFIG_REPOSITORYCONFIG_DATABASENAME", "test")

	defer func() {
		os.Unsetenv("APPCONFIG_ENVIRONMENT")
		os.Unsetenv("APPCONFIG_PORT")
		os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_DATABASETYPE")
		os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_CONNECTION")
		os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_DATABASENAME")
	}()

	config, err := GetConfig("")
	c.Nil(err)
	c.Equal("test", config.Environment)
	c.Equal("5000", config.Port)
}

func TestGetConfigMissingValue(t *testing.T) {
	c := require.New(t)

	os.Unsetenv("APPCONFIG_ENVIRONMENT")
	os.Unsetenv("APPCONFIG_PORT")
	os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_DATABASETYPE")
	os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_CONNECTION")
	os.Unsetenv("APPCONFIG_REPOSITORYCONFIG_DATABASENAME")

	_, err := GetConfig("")
	c.NotNil(err)
	c.True(errors.Is(err, ErrMissingValue))
}
