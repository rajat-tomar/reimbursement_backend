package util

import (
	"github.com/spf13/cobra"
	"reimbursement_backend/config"
	"reimbursement_backend/database"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "This is the root command",
	Long:  "It is the root command of cobra",
}

var migrateUpCommand = &cobra.Command{
	Use:   "migrate-up",
	Short: "It will run all the up migrations",
	Long:  "It will run all the up migrations from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		database.RunMigrationUp()
	},
}

var migrateDownCommand = &cobra.Command{
	Use:   "migrate-down",
	Short: "It will run all the down migrations",
	Long:  "It will run all the down migrations from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		database.RunMigrationDown()
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCommand)
	rootCmd.AddCommand(migrateDownCommand)
}

func ExecuteCommands() {
	err := rootCmd.Execute()
	if err != nil {
		config.SugarLogger.Error(err)
	}
}
