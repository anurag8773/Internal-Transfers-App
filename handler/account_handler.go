package handler

import (
	"encoding/json"
	"internal-transfers/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type createAccountRequest struct {
	AccountID      int64   `json:"account_id"`
	InitialBalance float64 `json:"initial_balance,string"`
}

// CreateAccount handles the creation of a new account.
// It validates the request body, checks for negative initial balances,
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.InitialBalance < 0 {
		http.Error(w, "initial balance cannot be negative", http.StatusBadRequest)
		return
	}

	if err := model.CreateAccount(req.AccountID, req.InitialBalance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAccount retrieves an account by its ID.
// It returns a 404 error if the account does not exist.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["account_id"], 10, 64)

	acc, err := model.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(acc)
}
