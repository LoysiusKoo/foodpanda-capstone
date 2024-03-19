-- name: CreatePlaylist :one
INSERT INTO playlists (
  name,
  image,
  food_items,
  is_active
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPlaylist :one
SELECT * FROM playlists
WHERE id = $1 LIMIT 1;

-- name: GetAllPlaylists :many
SELECT * FROM playlists
ORDER BY id;

-- name: UpdatePlaylist :one
UPDATE playlists
SET 
  name = $2,
  image = $3,
  food_items = $4,
  is_active = $5
WHERE 
  id = $1
RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE id = $1;