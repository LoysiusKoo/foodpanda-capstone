-- name: CreateRestaurant :one
INSERT INTO restaurants (
  name,
  description,
  address,
  rating,
  cuisine,
  image_url
) 
VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetRestaurant :one
SELECT * FROM restaurants
WHERE id = $1 LIMIT 1;

-- name: ListRestaurants :many
SELECT * FROM restaurants
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateRestaurant :one
UPDATE restaurants
SET 
  name = $2,
  description = $3,
  address = $4,
  rating = $5,
  cuisine = $6,
  image_url = $7
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants
WHERE id = $1;