package repositories

import (
	"context"
	"github.com/yanarowana123/onelab2_1/configs"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"github.com/yanarowana123/onelab2_1/internal/repositories/pgsql"
)

type IUserTransactionRepository interface {
	CreateTransaction(ctx context.Context, userTransaction models.CreateUserTransactionRequest) error
}

type Manager struct {
	UserTransaction IUserTransactionRepository
}

func NewManager(config configs.Config) *Manager {
	connection := pgsql.ConnectDB(config.PgSqlDSN)
	userTransactionRepository := pgsql.NewUserTransactionRepository(connection)

	return &Manager{
		UserTransaction: userTransactionRepository,
	}
}
