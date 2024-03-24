-- name: CreateSearch :one
INSERT INTO searches (
  user_id,
  keyword
) 
VALUES (
  $1, $2
)
RETURNING *;

-- name: GetSearch :one
SELECT * FROM searches
WHERE id = $1 LIMIT 1;

-- name: ListSearches :many
SELECT * FROM searches
WHERE id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;

SELECT d.id AS dish_id,
       d.name AS dish_name,
       d.description AS dish_description,
       d.price AS dish_price,
       d.image_url AS dish_imageURL,
       r.name AS restaurant_name,
       r.id AS restaurant_id
FROM dishes d
JOIN restaurants r ON d.restaurant_id = r.id
WHERE d.name ILIKE '%'||$1||'%' OR d.description ILIKE '%'||$1||'%';