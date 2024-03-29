
-- name: CreatePlaylistDish :one
INSERT INTO playlist_dishes (
  playlist_id,
  dish_id,
  date_to_be_delivered,
  image_Url
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetPlaylistDishes :many
SELECT d.id AS dish_id, p.id AS playlist_dishesID, p.playlist_id, d.name, d.price, p.image_url, r.name, p.date_to_be_delivered, d.cuisine, d.diet_type, r.rating, r.num_of_reviews AS number_of_reviews, l.name AS playlist_name, l.is_active, l.numberofweeks, l.dayofweek
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
JOIN restaurants r ON d.restaurant_id = r.id
JOIN playlists l ON p.playlist_id = l.id 
WHERE p.playlist_id = $1
ORDER BY p.id;

-- name: DeletePlaylistDish :one
DELETE FROM playlist_dishes
WHERE id = $1

RETURNING *;

-- name: UpdateDeliveryDate :one
UPDATE playlist_dishes
SET 
  date_to_be_delivered = $2
WHERE 
  id = $1
RETURNING *;