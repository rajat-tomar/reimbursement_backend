package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"reimbursement_backend/api"
	"reimbursement_backend/config"
	"reimbursement_backend/util"
)

func runServer() {
	const DefaultPort = 80
	port := config.Configuration.Server.HTTPPort
	if port == 0 {
		port = DefaultPort
	}
	address := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	handler := cors.Default().Handler(router)
	controllers := api.NewControllers()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	router.HandleFunc("/expense", controllers.ExpenseController.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses", controllers.ExpenseController.GetExpenses).Methods("GET")
	router.HandleFunc("/expense/{id}", controllers.ExpenseController.DeleteExpense).Methods("DELETE")

	http.ListenAndServe(address, handler)
}

func main() {
	config.InitConfiguration()
	config.InitLogger()
	config.InitDb()
	defer config.CloseDb()
	util.ExecuteCommands()
	runServer()
}
