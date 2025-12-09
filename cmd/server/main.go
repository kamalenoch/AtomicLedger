package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kamalenoch/AtomicLedger/internal/adapters/handler"
	"github.com/kamalenoch/AtomicLedger/internal/adapters/repository"
)

func main() {
	log.Println("Starting AtomicLedger...")

	// 1. Connect to Docker Database
	// We use the same credentials defined in docker-compose.yml
	dbRepo, err := repository.NewPostgresRepository("localhost", "5432", "user", "password", "atomic_ledger")
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to PostgreSQL Database")

	// 2. Inject the DB connection into the Handler
	// THIS FIXES THE ERROR: We pass 'dbRepo' into the function
	myHandler := handler.NewHTTPHandler(dbRepo)

	// 3. Setup Router
	router := gin.Default()
	router.POST("/accounts", myHandler.CreateAccount)
	router.POST("/transfer", myHandler.TransferMoney) 

	// 4. Run
	log.Println("Server running on :8080")
	router.Run(":8080")
}
