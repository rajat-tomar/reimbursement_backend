package util

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"reimbursement_backend/config"
	"reimbursement_backend/db"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "This is the root command",
	Long:  "It is the root command of cobra",
}

var migrateUpCommand = &cobra.Command{
	Use:   "migrate-up",
	Short: "It will run all the up migration",
	Long:  "It will run all the up migration from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationUp(); err != nil {
			config.Logger.Error(err)
		}
	},
}

var migrateDownCommand = &cobra.Command{
	Use:   "migrate-down",
	Short: "It will run all the down migration",
	Long:  "It will run all the down migration from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationDown(); err != nil {
			config.Logger.Error(err)
		}
	},
}

var createMigrationCommand = &cobra.Command{
	Use:   "create-migration",
	Short: "It will create migration",
	Long:  "It will create migration from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createMigrations(); err != nil {
			config.Logger.Fatalw("Migrations creation failed", "error", err)
		}
	},
}

var rollbackLatestMigrationCommand = &cobra.Command{
	Use:   "rollback-latest",
	Short: "It will rollback latest migration",
	Long:  "It will rollback latest migration from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RollbackLatestMigration(); err != nil {
			config.Logger.Fatalw("Migrations rollback failed", "error", err)
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
