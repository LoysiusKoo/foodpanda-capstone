
-- name: CreatePlaylistDish :one
INSERT INTO playlist_dishes (
  playlist_id,
  dish_id,
  date_to_be_delivered
) 
VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: GetPlaylistDishes :many
SELECT d.id, d.name, d.price, d.image_url, p.date_to_be_delivered, d.cuisine, d.diet_type, p.playlist_id as playlist_id
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
WHERE p.playlist_id = $1;