package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentMethod struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CardNumber string    `json:"card_number"`
	CardExpiry string    `json:"card_expiry"`
	CreatedAt  time.Time `json:"created_at"`
}
