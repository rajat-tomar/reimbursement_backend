package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func initDatabaseFromENV() {
	err := godotenv.Load("../.env")
	if err != nil {
		Logger.Error(err)
	}
	Configuration.Db.User = os.Getenv("DB_USER")
	Configuration.Db.Password = os.Getenv("DB_PASSWORD")
	Configuration.Db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Configuration.Db.DbName = os.Getenv("DB_NAME")
	Configuration.Db.Host = os.Getenv("DB_HOST")
}

func TestConnectToDatabase(t *testing.T) {
	InitConfiguration()
	assert.Nil(t, GetDb())
	initDatabaseFromENV()
	InitDb()
	assert.NotNil(t, GetDb())
}

func TestGet(t *testing.T) {
	InitConfiguration()
	initDatabaseFromENV()
	InitDb()
	assert.NotNil(t, GetDb())
}
