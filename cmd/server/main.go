package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamalenoch/AtomicLedger/internal/adapters/handler"
)

func main() {
	log.Println("Starting AtomicLedger Banking System...")

	// 1. Initialize the HTTP Handler
	myHandler := handler.NewHTTPHandler()

	// 2. Setup the Web Framework (Gin)
	router := gin.Default()

	// 3. Define Routes (API Endpoints)
	router.POST("/accounts", myHandler.CreateAccount)

	// 4. Start the Server on Port 8080
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
