package main

import (
	"reimbursement_backend/config"
	"reimbursement_backend/utils"
)

func main() {
	config.InitConfiguration()
	config.InitLogger()
	config.InitDb()
	defer config.CloseDb()
	utils.ExecuteCommands()
}
