package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//Service Logic here later.

type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

// CreateAccountRequest defines what the user must send us
type CreateAccountRequest struct {
	OwnerID  string `json:"owner_id" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}


func (h *HTTPHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
	
	// Validate the JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// We return a Mock response to prove the API works
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Account created successfully",
		"account_id": "acc_mock_12345", // Mock ID
		"owner_id":   req.OwnerID,
		"balance":    0,
	})
}