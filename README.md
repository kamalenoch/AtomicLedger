# AtomicLedger

AtomicLedger is a high-concurrency, double-entry accounting system designed to eliminate race conditions in financial transactions. Built with **Go** and **PostgreSQL**.


## The Problem: Double Spending
In high-scale fintech environments (like UPI payments), concurrent requests can lead to "double spending" if database locking isn't handled correctly at the isolation level.

## Architecture & Tech Stack
* **Language:** Go (Golang) - for goroutine concurrency handling.
* **Database:** PostgreSQL - utilizing `SERIALIZABLE` isolation and `SELECT ... FOR UPDATE` row locking.
* **Pattern:** Hexagonal Architecture (Ports & Adapters) to decouple core banking logic from HTTP/DB layers.

## Database Schema Design
AtomicLedger uses a strict double-entry system. Money is never "created" or "destroyed," only moved.

**Table: Accounts**
- `id` (UUID)
- `balance` (BigInt - storing cents/paisa to avoid floating point errors)

**Table: Entries**
- `id` (UUID)
- `account_id` (FK)
- `amount` (BigInt) (+ for credit, - for debit)
- `transaction_id` (FK)

## Concurrency Strategy
To handle 1000+ concurrent requests:
1.  **Pessimistic Locking:** The system locks the specific rows involved in a transaction during the write operation.
2.  **Idempotency Keys:** Redis cache will store request IDs to prevent network retries from duplicating transactions.

## Roadmap
- [x] System Design & Schema
- [x] Core Ledger Logic (Go)
- [x] API Layer (Gin/Fiber)
- [x] Dockerization
- [ ] Load Testing (k6 script)
