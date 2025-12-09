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

// CreateTransaction (Placeholder to satisfy Interface)
func (r *PostgresRepository) CreateTransaction(ctx context.Context, tx domain.Transaction) error {
	// We will implement this later
	return fmt.Errorf("method CreateTransaction not implemented yet")
}
