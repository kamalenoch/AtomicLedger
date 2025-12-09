package ports

import (
	"context"

	"github.com/kamalenoch/AtomicLedger/internal/core/domain"
)

// BankingRepository defines the "Contract" for the database.
type BankingRepository interface {

	CreateAccount(ctx context.Context, account *domain.Account) error

	GetAccountByID(ctx context.Context, id string) (*domain.Account, error)

	// CreateTransaction performs the atomic money transfer
	// the Locking magic 
	CreateTransaction(ctx context.Context, tx domain.Transaction) error
}