package config

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
