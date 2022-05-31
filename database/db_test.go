package database

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"reimbursement_backend/config"
	"strconv"
	"testing"
)

func initDatabaseFromENV() {
	err := godotenv.Load("../.env")
	if err != nil {
		config.SugarLogger.Error(err)
	}
	config.GetConfig().Database.User = os.Getenv("DB_USER")
	config.GetConfig().Database.Password = os.Getenv("DB_PASSWORD")
	config.GetConfig().Database.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	config.GetConfig().Database.DbName = os.Getenv("DB_NAME")
	config.GetConfig().Database.Host = os.Getenv("DB_HOST")
}

func TestConnectToDatabase(t *testing.T) {
	config.InitConfig()
	assert.Nil(t, db)
	initDatabaseFromENV()
	ConnectToDatabase()
	assert.NotNil(t, db)
}

func TestGet(t *testing.T) {
	config.InitConfig()
	initDatabaseFromENV()
	ConnectToDatabase()
	assert.NotNil(t, GetDatabase())
}
