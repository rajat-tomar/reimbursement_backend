package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"reimbursement_backend/config"
	"time"
)

func RunDbMigrationUp() error {
	dbConn, err := sql.Open("postgres", getDatabaseURL())
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		config.Logger.Error(err)
	}
	mig, err := migrate.NewWithDatabaseInstance("file://"+config.Configuration.Migration.FilePath, "postgres", driver)
	if err != nil {
		config.Logger.Error(err)
	}
	err = mig.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}
	return err
}

func RollbackLatestMigration() error {
	mig, err := migrate.New("file://"+config.Configuration.Migration.FilePath, getDatabaseURL())
	if err != nil {
		config.Logger.Error(err)
	}
	err = mig.Steps(-1)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}
	return err
}

func RunDbMigrationDown() error {
	dbConn, err := sql.Open("postgres", getDatabaseURL())
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		config.Logger.Error(err)
	}
	mig, err := migrate.NewWithDatabaseInstance("file://"+config.Configuration.Migration.FilePath, "postgres", driver)
	if err != nil {
		config.Logger.Error(err)
	}
	err = mig.Down()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}
	return err
}

func CreateMigration(filename string) error {
	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", config.Configuration.Migration.FilePath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", config.Configuration.Migration.FilePath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		config.Logger.Error(err)
	}
	fmt.Printf("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		if err := os.Remove(upMigrationFilePath); err != nil {
			config.Logger.Error(err)
		}
	}
	fmt.Printf("created %s\n", downMigrationFilePath)
	return nil
}

func getDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Configuration.Db.User, config.Configuration.Db.Password, config.Configuration.Db.Host, config.Configuration.Db.Port,
		config.Configuration.Db.DbName, config.Configuration.Db.SslMode)
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		config.Logger.Error(err)
	}
	err = f.Close()
	return err
}
