package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	HttpPort       int        `mapstructure:"HTTP_PORT"`
	GoogleClientId string     `mapstructure:"GOOGLE_CLIENT_ID"`
	Db             *DbConfig  `mapstructure:"DB"`
	Log            *LogConfig `mapstructure:"LOG"`
	Migration      *Migration `mapstructure:"MIGRATION"`
	Environment    string     `mapstructure:"ENVIRONMENT"`
}

type DbConfig struct {
	User     string `mapstructure:"USER"`
	DbName   string `mapstructure:"NAME"`
	Host     string `mapstructure:"HOST"`
	Password string `mapstructure:"PASSWORD"`
	Port     int    `mapstructure:"PORT"`
	SslMode  string `mapstructure:"SSL_MODE"`
}

type LogConfig struct {
	Level string `mapstructure:"LEVEL"`
}

type Migration struct {
	FilePath string `mapstructure:"FILE_PATH"`
}

var Config *Configuration

func InitConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		Logger.Panicw("cannot read config file", "error", err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		Logger.Panicw("Cannot unmarshal config", "error", err)
	}
}
