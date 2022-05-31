package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"reimbursement_backend/config"
)

var db *sql.DB

func ConnectToDatabase() {
	configuration := config.GetConfig()
	postgresConnect := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", configuration.Database.User, configuration.Database.Password, configuration.Database.DbName)
	database, err := sql.Open("postgres", postgresConnect)
	if err != nil {
		config.SugarLogger.Error("Cannot connect to database.", err)
	}
	db = database
	err = database.Ping()
	if err != nil {
		config.SugarLogger.Error("Error in database ping.", err)
	}
}

func CloseDatabase() {
	defer db.Close()
}

func GetDatabase() *sql.DB {
	return db
}
