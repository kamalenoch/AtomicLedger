package domain

import (
	"time"
)

// User's Bank Account
type Account struct {
	ID        string    `json:"id"`
	OwnerId   string    `json:"owner_id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

// Using Double Entry
type Transaction struct {
	ID          string    `json:"id"`
	FromAccount string    `json:"from_account"`
	ToAccount   string    `json:"to_account"`
	Amount      int64     `json:"amount"`
	Type        string    `json:"type"`   // TRANSFER, DEPOSIT, WITHDRAWAL
	Status      string    `json:"status"` // PENDING, COMPLETED, FAILED
	CreatedAt   time.Time `json:"created_at"`
}
