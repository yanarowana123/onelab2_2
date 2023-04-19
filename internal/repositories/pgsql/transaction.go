package pgsql

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2_1/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db,
	}
}

func (r TransactionRepository) Create(ctx context.Context, createTransactionRequest models.CreateTransactionRequest) error {
	_, err := r.db.ExecContext(ctx, "insert into transactions (user_id, book_id, checkout_id, money_amount) values ($1,$2,$3,$4)",
		createTransactionRequest.UserID, createTransactionRequest.BookID, createTransactionRequest.CheckoutID, createTransactionRequest.MoneyAmount)
	return err
}

func (r TransactionRepository) GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error) {
	var sum float64
	err := r.db.QueryRowContext(ctx, "select sum(money_amount) from transactions where book_id =$1 group by book_id", bookID).Scan(&sum)

	return sum, err
}
