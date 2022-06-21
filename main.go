package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"reimbursement_backend/config"
	"reimbursement_backend/util"
)

func httpServer() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This Is Our Golang Hello World ServerConfiguration")
	})

	var port interface{}
	port = config.Configuration.Server.HTTPPort
	if port == "" {
		port = "80"
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)

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
	httpServer()
}
