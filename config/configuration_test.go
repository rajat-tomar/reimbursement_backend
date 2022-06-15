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
		Logger.Error(err)
	}
	Configuration.Db.User = os.Getenv("DB_USER")
	Configuration.Db.Password = os.Getenv("DB_PASSWORD")
	Configuration.Db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Configuration.Db.DbName = os.Getenv("DB_NAME")
	Configuration.Db.Host = os.Getenv("DB_HOST")
	Configuration.Logging.Level = os.Getenv("LOG_LEVEL")
}
func TestInitConfig(t *testing.T) {
	assert.Nil(t, Configuration)
	assert.NotNil(t, Configuration)
}

func TestInitConfiguration(t *testing.T) {
	assert.Empty(t, Configuration.Db.User)
	initConfigurationFromENV()
	assert.NotEmpty(t, Configuration.Db.User)
}

//func TestGetConfig(t *testing.T) {
//	assert.NotNil(t, GetConfig())
//}
