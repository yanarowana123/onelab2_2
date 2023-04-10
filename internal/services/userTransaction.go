package services

import (
	"context"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"github.com/yanarowana123/onelab2_1/internal/repositories"
)

type IUserTransactionService interface {
	CreateTransaction(ctx context.Context, userTransaction models.CreateUserTransactionRequest) error
}

type UserTransactionService struct {
	repository repositories.Manager
}

func NewUserTransactionService(repository repositories.Manager) *UserTransactionService {
	return &UserTransactionService{
		repository: repository,
	}
}

func (u UserTransactionService) CreateTransaction(ctx context.Context, userTransaction models.CreateUserTransactionRequest) error {
	return u.repository.UserTransaction.CreateTransaction(ctx, userTransaction)
}
