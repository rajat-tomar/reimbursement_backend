package config

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		Config.Db.User, Config.Db.Password, Config.Db.Host, Config.Db.Port,
		Config.Db.DbName, Config.Db.SslMode)
	database, err := sql.Open("pgx", dsn)
	if err != nil {
		Logger.Panicw("cannot open the database", "error", err)
	}
	db = database
	if err = db.Ping(); err != nil {
		Logger.Panicw("cannot ping the database", "error", err)
	}
}

func CloseDb() {
	if err := db.Close(); err != nil {
		Logger.Panicw("cannot close the database", "error", err)
	}
}

func GetDb() *sql.DB {
	return db
}
