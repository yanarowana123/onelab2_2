package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2_1/configs"
	"github.com/yanarowana123/onelab2_1/internal/models"
	"github.com/yanarowana123/onelab2_1/internal/repositories/pgsql"
)

type ITransactionRepository interface {
	Create(ctx context.Context, userTransaction models.CreateTransactionRequest) error
	GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error)
}

type Manager struct {
	Transaction ITransactionRepository
}

func NewManager(config configs.Config) *Manager {
	connection := pgsql.ConnectDB(config.PgSqlDSN)
	transactionRepository := pgsql.NewTransactionRepository(connection)

	return &Manager{
		Transaction: transactionRepository,
	}
}
