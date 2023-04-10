package services

import (
	"github.com/yanarowana123/onelab2_1/internal/repositories"
)

type Manager struct {
	UserTransaction IUserTransactionService
}

func NewManager(repository repositories.Manager) *Manager {
	userTransactionService := NewUserTransactionService(repository)
	return &Manager{UserTransaction: userTransactionService}
}
