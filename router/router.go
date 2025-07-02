package router

import (
	"github.com/gorilla/mux"
	"internal-transfers/handler" // Importing the handler package to access the HTTP handlers
)
// SetupRouter initializes the router and defines the routes for the application.
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{account_id:[0-9]+}", handler.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", handler.SubmitTransaction).Methods("POST")

	return r
}
