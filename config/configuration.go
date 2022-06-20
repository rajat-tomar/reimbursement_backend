package config

import (
	"github.com/spf13/viper"
)

var Configuration *Configurations

type Configurations struct {
	Db         *DatabaseConfiguration `mapstructure:"DB"`
	Logging    *LoggingConfiguration  `mapstructure:"LOG"`
	Server     *ServerConfiguration   `mapstructure:"SERVER"`
	Migrations *Migration             `mapstructure:"MIGRATIONS"`
}

type DatabaseConfiguration struct {
	User     string `mapstructure:"USER"`
	DbName   string `mapstructure:"NAME"`
	Host     string `mapstructure:"HOST"`
	Password string `mapstructure:"PASSWORD"`
	Port     int    `mapstructure:"PORT"`
	SslMode  string `mapstructure:"SSL_MODE"`
}

type LoggingConfiguration struct {
	Level string `mapstructure:"LEVEL"`
}

type ServerConfiguration struct {
	HTTPPort string `mapstructure:"HTTP_PORT"`
}

type Migration struct {
	FilePath string `mapstructure:"FILE_PATH"`
}

func InitConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		Logger.Error(err)
	}
}

func InitTestConfiguration() {
	viper.SetConfigName("config.test")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		Logger.Error(err)
	}
}
