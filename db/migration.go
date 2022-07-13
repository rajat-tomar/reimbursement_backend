package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"reimbursement_backend/config"
	"time"
)

func RunDbMigrationUp() error {
	db, err := sql.Open("postgres", getDSN())
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		config.Logger.Error(err)
		return err
	}
	mig, err := migrate.NewWithDatabaseInstance("file://"+config.Config.Migration.FilePath, "postgres", driver)
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	err = mig.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func RollbackLatestMigration() error {
	mig, err := migrate.New("file://"+config.Config.Migration.FilePath, getDSN())
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	err = mig.Steps(-1)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func RunDbMigrationDown() error {
	db, err := sql.Open("postgres", getDSN())
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		config.Logger.Error(err)
		return err
	}
	mig, err := migrate.NewWithDatabaseInstance("file://"+config.Config.Migration.FilePath, "postgres", driver)
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	err = mig.Down()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func CreateMigration(filename string) error {
	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", config.Config.Migration.FilePath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", config.Config.Migration.FilePath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		config.Logger.Error(err)
		return err
	}
	log.Printf("Created migration file: %s", upMigrationFilePath)
	if err := createFile(downMigrationFilePath); err != nil {
		if err := os.Remove(upMigrationFilePath); err != nil {
			config.Logger.Error(err)
			return err
		}
	}
	log.Printf("Created migration file: %s", downMigrationFilePath)

	return nil
}

func getDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Db.User, config.Config.Db.Password, config.Config.Db.Host, config.Config.Db.Port,
		config.Config.Db.DbName, config.Config.Db.SslMode)
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	err = f.Close()

	return err
}
