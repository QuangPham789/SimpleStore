-- name: CreateItem :one
INSERT INTO items (
   code, product_id, unit, total_price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetItem :one
SELECT * FROM item
WHERE id = $1 LIMIT 1;

-- name: ListItem :many
SELECT * FROM item
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteProduct :exec
DELETE FROM item 
WHERE id = $1;