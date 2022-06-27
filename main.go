package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"reimbursement_backend/api"
	"reimbursement_backend/config"
	"reimbursement_backend/util"
)

func runServer() {
	const DefaultPort = 80
	r := mux.NewRouter()
	controllers := api.NewControllers()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health check says ok!")
	})
	r.HandleFunc("/expense", controllers.ExpenseController.CreateExpense).Methods("POST")
	r.HandleFunc("/expense", controllers.ExpenseController.GetExpenseById).Methods("GET")

	port := config.Configuration.Server.HTTPPort
	if port == 0 {
		port = DefaultPort
	}
	address := fmt.Sprintf("0.0.0.0:%d", port)
	if err := http.ListenAndServe(address, r); err != nil {
		config.Logger.Error(err)
	}
}

func main() {
	config.InitConfiguration()
	config.InitLogger()
	config.InitDb()
	defer config.CloseDb()
	util.ExecuteCommands()
	runServer()
}
