package utils

import (
	"github.com/spf13/cobra"
	"os"
	"reimbursement_backend/config"
	"reimbursement_backend/db"
)

var root = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  "It is the root command of cobra",
}

var runServer = &cobra.Command{
	Use:   "server",
	Short: "run http server",
	Long:  "It will run the http server",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}

var createMigration = &cobra.Command{
	Use:   "migrate-create",
	Short: "create migrations",
	Long:  "It will create migrations from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := os.Args[2]
		if fileName == "" {
			config.Logger.Error("file name is required")
		}
		if err := db.CreateMigration(fileName); err != nil {
			config.Logger.Fatalw("Could not create migrations successfully: ", "error", err)
		}
	},
}

var migrateUp = &cobra.Command{
	Use:   "migrate-up",
	Short: "run all up migrations",
	Long:  "It will run all the up migration from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationUp(); err != nil {
			config.Logger.Fatalw("Could not run up migrations successfully: ", "error", err)
		}
	},
}

var migrateDown = &cobra.Command{
	Use:   "migrate-down",
	Short: "run all the down migrations",
	Long:  "It will run all the down migrations from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RunDbMigrationDown(); err != nil {
			config.Logger.Fatalw("Could not run down migrations successfully: ", "error", err)
		}
	},
}

var rollbackLatestMigration = &cobra.Command{
	Use:   "migrate-rollback",
	Short: "rollback latest migration",
	Long:  "It will rollback latest migration from CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.RollbackLatestMigration(); err != nil {
			config.Logger.Fatalw("could not rollback latest migrations successfully: ", "error", err)
		}
	},
}

func init() {
	root.AddCommand(runServer)
	root.AddCommand(createMigration)
	root.AddCommand(migrateUp)
	root.AddCommand(migrateDown)
	root.AddCommand(rollbackLatestMigration)
}

func ExecuteCommands() {
	if err := root.Execute(); err != nil {
		config.Logger.Error(err)
	}
}
