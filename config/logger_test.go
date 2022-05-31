package config

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestInitLogger(t *testing.T) {
	assert.Nil(t, SugarLogger)
	InitLogger()
	assert.NotNil(t, SugarLogger)
}

func TestGetLogLevel(t *testing.T) {
	InitConfig()
	initConfigurationFromENV()
	expectedLevel := zapcore.DebugLevel
	actualLevel := getLogLevel()
	assert.Equal(t, expectedLevel, actualLevel)
}
