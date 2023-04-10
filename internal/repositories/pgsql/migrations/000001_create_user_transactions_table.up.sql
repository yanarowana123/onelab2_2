CREATE TABLE IF NOT EXISTS user_transactions
(
    user_id       uuid,
    book_id     uuid,
    money_amount     int,
    created_at timestamp default now()
)


