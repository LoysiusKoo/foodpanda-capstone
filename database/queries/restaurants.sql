-- name: CreateRestaurant :one
INSERT INTO restaurants (
  name,
  description,
  address,
  rating,
  restaurant_type,
  num_of_reviews,
  image_url
) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetRestaurants :many
SELECT * FROM restaurants
ORDER BY id;

-- name: GetRestaurantByID :one
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
  restaurant_type = $6,
  num_of_reviews = $7,
  image_url = $8
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants
WHERE id = $1;