package services

import (
	"github.com/yanarowana123/onelab2_1/configs"
	"github.com/yanarowana123/onelab2_1/internal/repositories"
)

type Manager struct {
	Transaction ITransactionService
	Auth        IAuthService
}

func NewManager(repository repositories.Manager, config configs.Config) *Manager {
	transactionService := NewTransactionService(repository)
	authService := NewAuthService(config.SecretKey)
	return &Manager{Transaction: transactionService, Auth: authService}
}
