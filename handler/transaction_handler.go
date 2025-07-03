package handler

import (
	"encoding/json"
	"internal-transfers/model"
	"net/http"
)

type transactionRequest struct {
	SourceAccountID      *int64   `json:"source_account_id"`
	DestinationAccountID *int64   `json:"destination_account_id"`
	Amount               float64 `json:"amount,string"`
}

// SubmitTransaction handles the submission of a transaction between two accounts.
// It validates the request, checks account balances, and processes the transaction.
func SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var req transactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.SourceAccountID == nil {
		http.Error(w, "Source Account ID is required", http.StatusBadRequest)
		return
	}

	if req.DestinationAccountID == nil {
		http.Error(w, "Destination Account ID is required", http.StatusBadRequest)
		return
	}
	// Check if the source and destination accounts exist
	if *req.SourceAccountID <= 0 || *req.DestinationAccountID <= 0 {
		http.Error(w, "account IDs must be greater than zero", http.StatusBadRequest)
		return
	}
	// Check if the amount is valid
	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than zero", http.StatusBadRequest)
		return
	}
	// Check if source and destination accounts are the same
	// This prevents a transaction from being submitted to the same account
	if req.SourceAccountID == req.DestinationAccountID {
		http.Error(w, "source and destination cannot be the same", http.StatusBadRequest)
		return
	}
	// Submit the transaction
	// This will lock the source account, check if it has sufficient funds,
	if err := model.SubmitTransaction(*req.SourceAccountID, *req.DestinationAccountID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
