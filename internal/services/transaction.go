package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"github.com/yanarowana123/onelab2_1/internal/repositories"
)

type ITransactionService interface {
	Create(ctx context.Context, userTransaction models.CreateTransactionRequest) error
	GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error)
}

type TransactionService struct {
	repository repositories.Manager
}

func NewTransactionService(repository repositories.Manager) *TransactionService {
	return &TransactionService{
		repository: repository,
	}
}

func (s *TransactionService) Create(ctx context.Context, userTransaction models.CreateTransactionRequest) error {
	return s.repository.Transaction.Create(ctx, userTransaction)
}

func (s *TransactionService) GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error) {
	return s.repository.Transaction.GetSumByBookID(ctx, bookID)
}
