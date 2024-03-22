
-- name: GetUpcomingDelivery :one
SELECT * FROM playlist_dishes
WHERE date_to_be_delivered = $1 LIMIT 1;


-- name: GetDeliveryDates :many
SELECT d.date_to_be_delivered 
FROM playlist_dishes d
JOIN playlists p ON d.playlist_id = p.id
WHERE p.is_active = true;
