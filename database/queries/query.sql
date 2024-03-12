-- name: GetRestaurant :one
SELECT * FROM restaurant WHERE id = $1 LIMIT 1;

-- name: ListRestaurants :many
SELECT * FROM restaurants
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;