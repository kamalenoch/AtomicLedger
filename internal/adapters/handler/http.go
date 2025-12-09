package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamalenoch/AtomicLedger/internal/core/domain"
	"github.com/kamalenoch/AtomicLedger/internal/core/ports"
)

// HTTPHandler holds the Repository Interface so it can save to DB
type HTTPHandler struct {
	repo ports.BankingRepository
}

// NewHTTPHandler injects the database dependency
func NewHTTPHandler(repo ports.BankingRepository) *HTTPHandler {
	return &HTTPHandler{
		repo: repo,
	}
}

type CreateAccountRequest struct {
	OwnerID  string `json:"owner_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (h *HTTPHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest

	// 1. Validate the JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Create the domain object
	// We set Balance to 0 and CreatedAt to now
	newAcc := domain.Account{
		OwnerID:   req.OwnerID,
		Balance:   0,
		Currency:  req.Currency,
		CreatedAt: time.Now(),
	}

	// 3. Save to Real Database
	// We pass &newAcc (pointer) so the Repo can update the ID
	err := h.repo.CreateAccount(c, &newAcc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	// 4. Return success with the REAL ID from the database
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Account created successfully",
		"account_id": newAcc.ID, // This will be a real UUID 
		"owner_id":   newAcc.OwnerID,
		"balance":    newAcc.Balance,
	})
}
