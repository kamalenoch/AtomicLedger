package main

import (
	"fmt"
	"log"

	"github.com/kamalenoch/AtomicLedger/internal/core/domain"
)

func main() {
	log.Println("Starting AtomicLedger Banking System...")

	testAccount := domain.Account{
		ID:       "acc_test_001",
		OwnerId:  "kamal_enoch",
		Balance:  50000,
		Currency: "INR",
	}

	fmt.Printf("System Initialized.\nStruct Test: Account [%s] loaded with Balance: %d %s\n",
		testAccount.ID, testAccount.Balance, testAccount.Currency)

}
