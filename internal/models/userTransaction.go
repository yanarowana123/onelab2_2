package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserTransactionRequest struct {
	UserID      uuid.UUID
	BookID      uuid.UUID
	MoneyAmount int
}
type UserTransactionResponse struct {
	UserID      uuid.UUID
	BookID      uuid.UUID
	MoneyAmount int
	CreatedAt   time.Time
}
