package database

import (
	"github.com/stretchr/testify/assert"
	"reimbursement_backend/config"
	"testing"
)

func TestConnectToDatabase(t *testing.T) {
	config.InitConfig()
	config.InitConfiguration()
	assert.Nil(t, db)
	ConnectToDatabase()
	assert.NotNil(t, db)
}

func TestGet(t *testing.T) {
	config.InitConfig()
	config.InitConfiguration()
	ConnectToDatabase()
	assert.NotNil(t, GetDatabase())
}
