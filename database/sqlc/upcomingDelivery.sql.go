// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: upcomingDelivery.sql

package db

import (
	"context"
)

const getDeliveryDates = `-- name: GetDeliveryDates :many
SELECT d.date_to_be_delivered 
FROM playlist_dishes d
JOIN playlists p ON d.playlist_id = p.id
WHERE p.is_active = true
`

func (q *Queries) GetDeliveryDates(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getDeliveryDates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var date_to_be_delivered string
		if err := rows.Scan(&date_to_be_delivered); err != nil {
			return nil, err
		}
		items = append(items, date_to_be_delivered)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUpcomingDelivery = `-- name: GetUpcomingDelivery :one
SELECT d.id AS playlistdish_id, i.id AS dish_id, i.name AS dish_name, i.price AS dish_price, p.image AS playlist_image, d.date_to_be_delivered, i.cuisine, i.diet_type, p.id AS playlist_id, i.image_url AS dish_image
FROM playlist_dishes d
JOIN playlists p ON d.playlist_id = p.id
JOIN dishes i ON d.dish_id = i.id
WHERE p.is_active = true AND d.date_to_be_delivered = $1 LIMIT 1
`

type GetUpcomingDeliveryRow struct {
	PlaylistdishID    int64   `json:"playlistdish_id"`
	DishID            int64   `json:"dish_id"`
	DishName          string  `json:"dish_name"`
	DishPrice         float64 `json:"dish_price"`
	PlaylistImage     string  `json:"playlist_image"`
	DateToBeDelivered string  `json:"date_to_be_delivered"`
	Cuisine           string  `json:"cuisine"`
	DietType          string  `json:"diet_type"`
	PlaylistID        int64   `json:"playlist_id"`
	DishImage         string  `json:"dish_image"`
}

func (q *Queries) GetUpcomingDelivery(ctx context.Context, dateToBeDelivered string) (GetUpcomingDeliveryRow, error) {
	row := q.db.QueryRowContext(ctx, getUpcomingDelivery, dateToBeDelivered)
	var i GetUpcomingDeliveryRow
	err := row.Scan(
		&i.PlaylistdishID,
		&i.DishID,
		&i.DishName,
		&i.DishPrice,
		&i.PlaylistImage,
		&i.DateToBeDelivered,
		&i.Cuisine,
		&i.DietType,
		&i.PlaylistID,
		&i.DishImage,
	)
	return i, err
}
