package config

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		Configuration.Db.User, Configuration.Db.Password, Configuration.Db.Host, Configuration.Db.Port,
		Configuration.Db.DbName, Configuration.Db.SslMode)
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		Logger.Panicw("cannot initialize database", "error", err)
	}
	db = database
	if err = db.Ping(); err != nil {
		Logger.Panicw("cannot connect to the database", "error", err)
	}
	Logger.Infow("successfully connected to the database")
}

func CloseDb() {
	if err := db.Close(); err != nil {
		Logger.Panicw("cannot close the database", "error", err)
	}
}

func GetDb() *sql.DB {
	return db
}
