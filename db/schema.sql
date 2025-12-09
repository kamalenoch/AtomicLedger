-- Enable UUID extension (Required for unique IDs)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Accounts Table: Stores the user and their balance
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0, -- Storing in cents/paisa
    currency VARCHAR(3) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Transactions Table: The Double-Entry Ledger
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_account_id UUID REFERENCES accounts(id),
    to_account_id UUID REFERENCES accounts(id),
    amount BIGINT NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);