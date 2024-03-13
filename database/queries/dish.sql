-- name: CreateDish :one
INSERT INTO dishes (
  restaurant_id,
  is_available,
  name,
  description,
  price,
  cuisine,
  image_url

) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetDish :one
SELECT * FROM dishes
WHERE id = $1 LIMIT 1;

-- name: GetDishesByCuisine :many
SELECT * FROM dishes
WHERE cuisine ILIKE '%'||$1||'%'
ORDER BY id;

-- name: GetDishes :many
SELECT * FROM dishes
ORDER BY id;

-- name: ListDishes :many
SELECT * FROM dishes
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateDish :one
UPDATE dishes
SET 
  restaurant_id = $2,
  is_available = $3,
  name = $4,
  description = $5,
  price = $6,
  cuisine = $7,
  image_url = $8
WHERE 
  id = $1
RETURNING *;

-- name: DeleteDish :exec
DELETE FROM dishes
WHERE id = $1;