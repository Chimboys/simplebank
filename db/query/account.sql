-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: CreateAccount :exec
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;
