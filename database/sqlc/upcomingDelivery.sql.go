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
SELECT id, playlist_id, dish_id, date_to_be_delivered, image_url, created_at, added_at FROM playlist_dishes
WHERE date_to_be_delivered = $1 LIMIT 1
`

func (q *Queries) GetUpcomingDelivery(ctx context.Context, dateToBeDelivered string) (PlaylistDish, error) {
	row := q.db.QueryRowContext(ctx, getUpcomingDelivery, dateToBeDelivered)
	var i PlaylistDish
	err := row.Scan(
		&i.ID,
		&i.PlaylistID,
		&i.DishID,
		&i.DateToBeDelivered,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.AddedAt,
	)
	return i, err
}
