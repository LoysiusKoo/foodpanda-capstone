
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
SELECT d.id, d.name, d.price, d.image_url, p.date_to_be_delivered, d.cuisine, d.diet_type, p.playlist_id, r.rating, r.num_of_reviews AS number_of_reviews
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
JOIN restaurants r ON d.restaurant_id = r.id
WHERE p.playlist_id = $1;