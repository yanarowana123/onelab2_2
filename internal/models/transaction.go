package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateTransactionRequest struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	BookID      uuid.UUID `json:"book_id" validate:"required"`
	CheckoutID  uuid.UUID `json:"checkout_id" validate:"required"`
	MoneyAmount float64   `json:"money_amount" validate:"required"`
}
type TransactionResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	BookID      uuid.UUID `json:"book_id"`
	CheckoutID  uuid.UUID `json:"checkout_id"`
	MoneyAmount float64   `json:"money_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
