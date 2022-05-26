package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"reimbursement_backend/config"
)

func runMigrationUp() {
	driver, err := postgres.WithInstance(GetDatabase(), &postgres.Config{})
	if err != nil {
		config.SugarLogger.Error(err)
	}
	mig, err := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if err != nil {
		config.SugarLogger.Error(err)
	}
	result := mig.Up()
	if result == nil {
		config.SugarLogger.Error(err)
	}
}

func runMigrationDown() {
	driver, err := postgres.WithInstance(GetDatabase(), &postgres.Config{})
	if err != nil {
		config.SugarLogger.Error(err)
	}
	mig, err := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if err != nil {
		config.SugarLogger.Error(err)
	}
	result := mig.Down()
	if result == nil {
		config.SugarLogger.Error(err)
	}
}
