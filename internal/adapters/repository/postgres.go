package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kamalenoch/AtomicLedger/internal/core/domain"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(host, port, user, password, dbname string) (*PostgresRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

// CreateAccount (Implemented)
func (r *PostgresRepository) CreateAccount(ctx context.Context, acc *domain.Account) error {
	query := `
		INSERT INTO accounts (owner_id, balance, currency, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := r.db.QueryRowContext(ctx, query, acc.OwnerID, acc.Balance, acc.Currency, acc.CreatedAt).Scan(&acc.ID)
	if err != nil {
		log.Printf("Error creating account: %v", err)
		return err
	}
	return nil
}

// GetAccountByID (Placeholder to satisfy Interface)
func (r *PostgresRepository) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	// We will implement this later
	return nil, fmt.Errorf("method GetAccountByID not implemented yet")
}

// CreateTransaction executes the money move with ACID guarantees
func (r *PostgresRepository) CreateTransaction(ctx context.Context, txReq domain.Transaction) error {
	// 1. Start a Database Transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Safety net: Rollback if anything goes wrong (unless we Commit)
	defer tx.Rollback()

	// 2. FETCH & LOCK the Source Account ("SELECT FOR UPDATE")
	// This stops any race conditions immediately.
	var currentBalance int64
	queryLock := `SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`
	
	err = tx.QueryRowContext(ctx, queryLock, txReq.FromAccount).Scan(&currentBalance)
	if err != nil {
		return fmt.Errorf("failed to lock source account: %v", err)
	}

	// 3. Check Balance (The Business Rule)
	if currentBalance < txReq.Amount {
		return fmt.Errorf("insufficient funds")
	}

	// 4. Update Balances (Deduct from Source, Add to Destination)
	// Note: We should technically lock the destination too to prevent deadlocks, 
	// but for this MVP, we update directly.
	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, txReq.Amount, txReq.FromAccount)
	if err != nil { return err }

	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, txReq.Amount, txReq.ToAccount)
	if err != nil { return err }

	// 5. Record the Transaction in the Ledger
	queryInsert := `
		INSERT INTO transactions (from_account_id, to_account_id, amount, status)
		VALUES ($1, $2, $3, 'COMPLETED')
		RETURNING id`
	
	err = tx.QueryRowContext(ctx, queryInsert, txReq.FromAccount, txReq.ToAccount, txReq.Amount).Scan(&txReq.ID)
	if err != nil { return err }

	// 6. Commit (Make it permanent)
	return tx.Commit()
}
