package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	assert.Nil(t, configuration)
	InitConfig()
	assert.NotNil(t, configuration)
}

func TestInitConfiguration(t *testing.T) {
	InitConfig()
	assert.Empty(t, configuration.Database.User)
	InitConfiguration()
	assert.NotEmpty(t, configuration.Database.User)
}

func TestGetConfig(t *testing.T) {
	InitConfig()
	assert.NotNil(t, GetConfig())
}
