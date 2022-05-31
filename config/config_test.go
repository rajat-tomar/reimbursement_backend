package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func initConfigurationFromENV() {
	err := godotenv.Load("../.env")
	if err != nil {
		SugarLogger.Error(err)
	}
	configuration.Database.User = os.Getenv("DB_USER")
	configuration.Database.Password = os.Getenv("DB_PASSWORD")
	configuration.Database.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	configuration.Database.DbName = os.Getenv("DB_NAME")
	configuration.Database.Host = os.Getenv("DB_HOST")
	configuration.Log.Level = os.Getenv("LOG_LEVEL")
}
func TestInitConfig(t *testing.T) {
	assert.Nil(t, configuration)
	InitConfig()
	assert.NotNil(t, configuration)
}

func TestInitConfiguration(t *testing.T) {
	InitConfig()
	assert.Empty(t, configuration.Database.User)
	initConfigurationFromENV()
	assert.NotEmpty(t, configuration.Database.User)
}

func TestGetConfig(t *testing.T) {
	InitConfig()
	assert.NotNil(t, GetConfig())
}
