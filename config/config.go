package config

import (
	"github.com/spf13/viper"
	"log"
)

var configuration *Configuration

type Database struct {
	User     string
	DbName   string
	Host     string
	Password string
	Port     int
}

type Configuration struct {
	Database Database
}

func InitConfig() {
	configuration = new(Configuration)
}

func GetConfig() *Configuration {
	return configuration
}

func InitConfiguration() {
	viper.SetConfigName("config.yml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file" + err.Error())
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal(err)
	}
}
