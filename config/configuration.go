package config

import (
	"github.com/spf13/viper"
)

var Configuration *Configurations

type Configurations struct {
	Environment string
	Db          *DatabaseConfiguration `mapstructure:"DB"`
	Logging     *LoggingConfiguration  `mapstructure:"LOG"`
	Server      *ServerConfiguration   `mapstructure:"SERVER"`
	Migration   *Migration             `mapstructure:"MIGRATION"`
	OAuth       *OAuth2Configuration   `mapstructure:"OAUTH"`
}

type DatabaseConfiguration struct {
	User     string `mapstructure:"USER"`
	DbName   string `mapstructure:"NAME"`
	Host     string `mapstructure:"HOST"`
	Password string `mapstructure:"PASSWORD"`
	Port     int    `mapstructure:"PORT"`
	SslMode  string `mapstructure:"SSL_MODE"`
}

type GoogleOAuth2Configuration struct {
	ClientID     string   `mapstructure:"CLIENT_ID"`
	ClientSecret string   `mapstructure:"CLIENT_SECRET"`
	RedirectURL  string   `mapstructure:"REDIRECT_URL"`
	Scopes       []string `mapstructure:"SCOPES"`
}

type OAuth2Configuration struct {
	Google *GoogleOAuth2Configuration `mapstructure:"GOOGLE"`
}

type LoggingConfiguration struct {
	Level string `mapstructure:"LEVEL"`
}

type ServerConfiguration struct {
	HTTPPort int `mapstructure:"HTTP_PORT"`
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
