package config

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestInitLogger(t *testing.T) {
	assert.Nil(t, Logger)
	InitLogger()
	assert.NotNil(t, Logger)
}

func TestGetLogLevel(t *testing.T) {
	expectedLevel := zapcore.DebugLevel
	actualLevel := getLogLevel()
	assert.Equal(t, expectedLevel, actualLevel)
}
