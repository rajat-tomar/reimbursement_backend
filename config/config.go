package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database
	Log      Log
	Server   Server
}

type Database struct {
	User     string
	DbName   string
	Host     string
	Password string
	Port     int
}

type Log struct {
	Level string
}

type Server struct {
	Port int
}

var configuration *Configuration

func InitConfig() {
	configuration = new(Configuration)
}

func GetConfig() *Configuration {
	return configuration
}

func InitConfiguration() {
	viper.SetConfigName("config.yml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		SugarLogger.Error(err)
	}
}
