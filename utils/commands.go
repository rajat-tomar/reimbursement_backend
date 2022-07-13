package utils

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"reimbursement_backend/config"
	"reimbursement_backend/db"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  "It is the root command of cobra",
}

var migrateUpCommand = &cobra.Command{
	Use:   "migrate-up",
	Short: "run all up migrations",
	Long:  "It will run all the up migration from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationUp(); err != nil {
			config.Logger.Fatalw("Could not run up migrations successfully: ", "error", err)
		}
	},
}

var migrateDownCommand = &cobra.Command{
	Use:   "migrate-down",
	Short: "run all the down migrations",
	Long:  "It will run all the down migrations from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationDown(); err != nil {
			config.Logger.Fatalw("Could not run down migrations successfully: ", "error", err)
		}
	},
}

var createMigrationCommand = &cobra.Command{
	Use:   "create-migration",
	Short: "create migrations",
	Long:  "It will create migrations from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createMigrations(); err != nil {
			config.Logger.Fatalw("Could not create migrations successfully: ", "error", err)
		}
	},
}

var rollbackLatestMigrationCommand = &cobra.Command{
	Use:   "rollback-latest",
	Short: "rollback latest migration",
	Long:  "It will rollback latest migration from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RollbackLatestMigration(); err != nil {
			config.Logger.Fatalw("could not rollback latest migrations successfully: ", "error", err)
		}
	},
}

func createMigrations() error {
	fileName := os.Args[2]
	if len(fileName) == 0 {
		return errors.New("filename is not provided")
	}
	return db.CreateMigration(fileName)
}

func init() {
	rootCmd.AddCommand(migrateUpCommand)
	rootCmd.AddCommand(migrateDownCommand)
	rootCmd.AddCommand(createMigrationCommand)
	rootCmd.AddCommand(rollbackLatestMigrationCommand)
}

func ExecuteCommands() {
	if err := rootCmd.Execute(); err != nil {
		config.Logger.Error(err)
	}
}
