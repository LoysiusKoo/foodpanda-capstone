
-- name: GetUpcomingDelivery :one
SELECT d.id AS playlistdish_id, i.id AS dish_id, i.name AS dish_name, i.price AS dish_price, d.date_to_be_delivered, i.cuisine, i.diet_type, p.id AS playlist_id, i.image_url AS dish_image, r.name AS restaurant_name
FROM playlist_dishes d
JOIN playlists p ON d.playlist_id = p.id
JOIN dishes i ON d.dish_id = i.id
JOIN restaurants r ON i.restaurant_id = r.id
WHERE p.is_active = true AND d.date_to_be_delivered = $1 LIMIT 1;


-- name: GetDeliveryDates :many
SELECT d.date_to_be_delivered 
FROM playlist_dishes d
JOIN playlists p ON d.playlist_id = p.id
WHERE p.is_active = true;
