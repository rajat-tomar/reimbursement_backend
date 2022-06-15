package config

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func InitDb() {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		Configuration.Db.User, Configuration.Db.Password, Configuration.Db.Host, Configuration.Db.Port,
		Configuration.Db.DbName, Configuration.Db.SslMode)
	var err error
	gormDb, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		Logger.Panicw("cannot initialize database", "error", err)
	}
	db, err := gormDb.DB()
	if err = db.Ping(); err != nil {
		Logger.Panicw("cannot connect to the database", "error", err)
	}
	Logger.Infow("successfully connected to the database")
}

func CloseDb() error {
	db, err := gormDb.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func GetDb() *gorm.DB {
	return gormDb
}
