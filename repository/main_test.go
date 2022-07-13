package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"reimbursement_backend/config"
	"reimbursement_backend/db"
	"testing"
)

var testDb *sql.DB

func setup() {
	migration := config.Migration{FilePath: "../db/migration"}
	databaseConfiguration := config.DbConfig{
		User:     "reimbursement",
		DbName:   "reimbursement_test",
		Host:     "localhost",
		Password: "password",
		Port:     5432,
		SslMode:  "disable",
	}
	config.Config = &config.Configuration{
		Environment: "test",
		Db:          &databaseConfiguration,
		Log:         &config.LogConfig{Level: "debug"},
		Migration:   &migration,
	}
	config.InitDb()
	db.RunDbMigrationUp()
}

func TestMain(m *testing.M) {
	setup()
	testDb = config.GetDb()
	code := m.Run()
	teardown()
	defer config.CloseDb()
	os.Exit(code)
}

func teardown() {
	db.RunDbMigrationDown()
}
