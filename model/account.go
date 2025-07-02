package model

import (
	"database/sql"
	"errors"
	"internal-transfers/database"
)

type Account struct {
	AccountID int64   `json:"account_id"`
	Balance   float64 `json:"balance"`
}

// CreateAccount creates a new account with the given account ID and initial balance.
// It returns an error if the account already exists or if there is a database error.
func CreateAccount(accountID int64, initialBalance float64) error {
	_, err := db.Conn.Exec(`INSERT INTO accounts (account_id, balance) VALUES ($1, $2)`, accountID, initialBalance)
	return err
}

// GetAccount retrieves an account by its ID.
func GetAccount(accountID int64) (*Account, error) {
	var acc Account
	err := db.Conn.QueryRow(`SELECT account_id, balance FROM accounts WHERE account_id = $1`, accountID).Scan(&acc.AccountID, &acc.Balance)
	if err == sql.ErrNoRows {
		return nil, errors.New("account not found")
	}
	return &acc, err
}
