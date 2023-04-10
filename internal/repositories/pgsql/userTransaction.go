package pgsql

import (
	"context"
	"database/sql"
	"github.com/yanarowana123/onelab2_1/internal/models"
)

type UserTransactionRepository struct {
	db *sql.DB
}

func NewUserTransactionRepository(db *sql.DB) *UserTransactionRepository {
	return &UserTransactionRepository{
		db,
	}
}

func (r UserTransactionRepository) CreateTransaction(ctx context.Context, userTransaction models.CreateUserTransactionRequest) error {
	_, err := r.db.ExecContext(ctx, "insert into user_transactions (user_id, book_id, money_amount) values ($1,$2,$3)",
		userTransaction.UserID, userTransaction.BookID, userTransaction.MoneyAmount)
	return err
}
