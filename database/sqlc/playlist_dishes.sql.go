// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: playlist_dishes.sql

package db

import (
	"context"
)

const createPlaylistDish = `-- name: CreatePlaylistDish :one
INSERT INTO playlist_dishes (
  playlist_id,
  dish_id,
  date_to_be_delivered
) 
VALUES (
  $1, $2, $3
)
RETURNING id, playlist_id, dish_id, date_to_be_delivered, created_at, added_at
`

type CreatePlaylistDishParams struct {
	PlaylistID        int64  `json:"playlist_id"`
	DishID            int64  `json:"dish_id"`
	DateToBeDelivered string `json:"date_to_be_delivered"`
}

func (q *Queries) CreatePlaylistDish(ctx context.Context, arg CreatePlaylistDishParams) (PlaylistDish, error) {
	row := q.db.QueryRowContext(ctx, createPlaylistDish, arg.PlaylistID, arg.DishID, arg.DateToBeDelivered)
	var i PlaylistDish
	err := row.Scan(
		&i.ID,
		&i.PlaylistID,
		&i.DishID,
		&i.DateToBeDelivered,
		&i.CreatedAt,
		&i.AddedAt,
	)
	return i, err
}

const getPlaylistDishes = `-- name: GetPlaylistDishes :many
SELECT d.id, d.name, d.price, d.image_url, p.date_to_be_delivered, d.cuisine, d.diet_type, p.playlist_id as playlist_id, r.rating, r.num_of_reviews AS numberOfReviews
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
JOIN restaurants r ON d.restaurant_id = r.id
WHERE p.playlist_id = $1
`

type GetPlaylistDishesRow struct {
	ID                int64   `json:"id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	ImageUrl          string  `json:"image_url"`
	DateToBeDelivered string  `json:"date_to_be_delivered"`
	Cuisine           string  `json:"cuisine"`
	DietType          string  `json:"diet_type"`
	PlaylistID        int64   `json:"playlist_id"`
	Rating            float64 `json:"rating"`
	Numberofreviews   int32   `json:"numberofreviews"`
}

func (q *Queries) GetPlaylistDishes(ctx context.Context, playlistID int64) ([]GetPlaylistDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPlaylistDishes, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPlaylistDishesRow{}
	for rows.Next() {
		var i GetPlaylistDishesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.ImageUrl,
			&i.DateToBeDelivered,
			&i.Cuisine,
			&i.DietType,
			&i.PlaylistID,
			&i.Rating,
			&i.Numberofreviews,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
