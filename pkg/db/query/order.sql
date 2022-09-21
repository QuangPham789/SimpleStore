-- name: CreateOrder :one
INSERT INTO order (
   code, account_id, items, total_price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM order
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts 
SET username = $3,
password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = $1;