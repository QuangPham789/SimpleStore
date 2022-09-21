-- name: CreateProduct :one
INSERT INTO products (
   code, name, description, category, price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products 
SET username = $3,
password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products 
WHERE id = $1;