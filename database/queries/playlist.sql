-- name: CreatePlaylist :one
INSERT INTO playlists (
  name,
  image,
  numberofweeks,
  dayofweek,
  food_items,
  is_active
) 
VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetPlaylist :one
SELECT * FROM playlists
WHERE id = $1 LIMIT 1;

-- name: GetAllPlaylists :many
SELECT * FROM playlists
ORDER BY id;

-- name: GetFoodItemsFromPlaylists :many
SELECT food_items FROM playlists
WHERE id = $1 LIMIT 1;

-- name: UpdateIsActive :one
UPDATE playlists
SET 
  is_active = $2
WHERE 
  id = $1
RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE id = $1;