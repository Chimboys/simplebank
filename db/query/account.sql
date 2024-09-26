-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccountByOwner :one
SELECT * FROM accounts WHERE owner = $1;

-- name: GetAccounts :many
SELECT * FROM accounts;

-- name: GetAccountWithBalanceLowerThan :many
SELECT * FROM accounts WHERE balance < $1;

-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING *;


-- name: DeleteAccountByOwner :one
DELETE FROM accounts WHERE owner = $1 RETURNING *;
