package utils

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
	"reimbursement_backend/api"
	"reimbursement_backend/config"
	"strings"
)

func filterContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.ServeHTTP(w, r)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}

		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(reqToken, "Bearer ")
		ctx := context.Background()

		if r.URL.Path == "/login" {
			payload, err := idtoken.Validate(ctx, tokenString, config.Config.GoogleClientId)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			name := payload.Claims["name"].(string)
			email := payload.Claims["email"].(string)
			generatedJwt, err := generateJWT(email, name)
			ctx = context.WithValue(ctx, "name", name)
			ctx = context.WithValue(ctx, "email", email)
			ctx = context.WithValue(ctx, "token", generatedJwt)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			name, email, err := validateToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, "name", name)
			ctx = context.WithValue(ctx, "email", email)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func RunServer() {
	const DefaultPort = 80
	port := config.Config.HttpPort
	if port == 0 {
		port = DefaultPort
	}
	address := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()
	controllers := api.NewControllers()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://reimbursement.gaussb.io", "https://reimbursement.staging.gaussb.io"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: true,
	})

	//router.Use(filterContentTypeMiddleware)
	router.Use(authenticationMiddleware)
	handler := c.Handler(router)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, world! This is the Reimbursement Backend."))
		if err != nil {
			config.Logger.Panicf("Error writing to response writer: %v", err)
		}
	}).Methods("GET")
	router.HandleFunc("/login", controllers.UserController.Login).Methods("POST")
	router.HandleFunc("/users", controllers.UserController.GetUsers).Methods("GET")
	router.HandleFunc("/expense", controllers.ExpenseController.GetExpenseById).Methods("GET")
	router.HandleFunc("/expenses", controllers.ExpenseController.GetExpenses).Methods("GET")
	router.HandleFunc("/expense", controllers.ExpenseController.CreateExpense).Methods("POST")
	router.HandleFunc("/expense", controllers.ExpenseController.DeleteExpense).Methods("DELETE")
	router.HandleFunc("/expense", controllers.ExpenseController.UpdateExpense).Methods("PUT")
	router.HandleFunc("/reimbursement", controllers.ReimbursementController.CreateReimbursement).Methods("POST")
	router.HandleFunc("/reimbursements", controllers.ReimbursementController.GetReimbursements).Methods("GET")
	router.HandleFunc("/reimbursements/process", controllers.ReimbursementController.ProcessReimbursements).Methods("PUT")

	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(address, handler))
}
