
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
SELECT *
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
WHERE p.id = $1
ORDER BY d.id;
