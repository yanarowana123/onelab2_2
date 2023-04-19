CREATE TABLE IF NOT EXISTS transactions
(
    user_id       uuid,
    book_id     uuid,
    checkout_id     uuid,
    money_amount     float,
    created_at timestamp default now()
)


