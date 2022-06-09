package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reimbursement_backend/config"
	"reimbursement_backend/util"
)

func httpServer() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This Is Our Golang Hello World Server")
	})

	configuration := config.GetConfig()
	var port interface{}
	port = configuration.Server.Port
	if port == 0 {
		port = 80
	}

	address := fmt.Sprintf("0.0.0.0:%d", port)

	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config.InitConfig()
	config.InitConfiguration()
	config.InitLogger()
	util.ExecuteCommands()
	httpServer()
}
