package model

import (
	"errors"
	"internal-transfers/database"
)

// SubmitTransaction processes a transaction between two accounts.
// It locks the source account, checks if it has sufficient funds,
func SubmitTransaction(sourceID, destID int64, amount float64) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// Lock accounts
	var sourceBalance float64
	err = tx.QueryRow(`SELECT balance FROM accounts WHERE account_id=$1 FOR UPDATE`, sourceID).Scan(&sourceBalance)
	if err != nil {
		return errors.New("source account not found")
	}
	// Check if source has enough balance
	if sourceBalance < amount {
		return errors.New("insufficient funds")
	}
	// Check if destination account exists
	var destExists bool
	err = tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM accounts WHERE account_id=$1)`, destID).Scan(&destExists)
	if err != nil || !destExists {
		return errors.New("destination account not found")
	}
	// Update balances of source account
	_, err = tx.Exec(`UPDATE accounts SET balance = balance - $1 WHERE account_id = $2`, amount, sourceID)
	if err != nil {
		return err
	}
	// Update balances of destination account
	_, err = tx.Exec(`UPDATE accounts SET balance = balance + $1 WHERE account_id = $2`, amount, destID)
	if err != nil {
		return err
	}
	// Record the transaction
	_, err = tx.Exec(`INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)`,
		sourceID, destID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
