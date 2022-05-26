package main

import (
	"reimbursement_backend/config"
	"reimbursement_backend/database"
	"reimbursement_backend/util"
)

func main() {
	config.InitConfig()
	config.InitConfiguration()
	config.InitLogger()
	database.ConnectToDatabase()
	util.ExecuteCommands()
	database.CloseDatabase()
}
