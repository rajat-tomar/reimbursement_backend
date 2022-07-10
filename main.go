package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"reimbursement_backend/api"
	"reimbursement_backend/config"
	"reimbursement_backend/utils"
)

func runServer() {
	const DefaultPort = 80
	port := config.Configuration.Server.HTTPPort
	if port == 0 {
		port = DefaultPort
	}
	address := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://accounts.google.com/", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	controllers := api.NewControllers()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte("{\"hello\": \"world\"}")); err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
	})
	router.HandleFunc("/login", controllers.OAuthController.GoogleLogin).Methods("GET")
	router.HandleFunc("/auth/google/callback", controllers.OAuthController.GoogleCallback).Methods("GET")
	router.HandleFunc("/expense", controllers.ExpenseController.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses", controllers.ExpenseController.GetExpenses).Methods("GET")
	router.HandleFunc("/expense", controllers.ExpenseController.DeleteExpense).Methods("DELETE")

	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(address, handler))
}

func main() {
	config.InitConfiguration()
	config.InitLogger()
	config.InitDb()
	defer config.CloseDb()
	utils.ExecuteCommands()
	runServer()
}
