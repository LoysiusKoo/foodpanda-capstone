// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: playlist_dishes.sql

package db

import (
	"context"
	"time"

	null "gopkg.in/guregu/null.v4"
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
SELECT d.id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url, p.id, playlist_id, dish_id, date_to_be_delivered, created_at, added_at
FROM dishes d JOIN playlist_dishes p ON d.id = p.dish_id
WHERE p.id = $1
ORDER BY d.id
`

type GetPlaylistDishesRow struct {
	ID                int64       `json:"id"`
	RestaurantID      int64       `json:"restaurant_id"`
	IsAvailable       bool        `json:"is_available"`
	Name              string      `json:"name"`
	Description       null.String `json:"description"`
	Price             float64     `json:"price"`
	DietType          string      `json:"diet_type"`
	Cuisine           string      `json:"cuisine"`
	ImageUrl          string      `json:"image_url"`
	ID_2              int64       `json:"id_2"`
	PlaylistID        int64       `json:"playlist_id"`
	DishID            int64       `json:"dish_id"`
	DateToBeDelivered string      `json:"date_to_be_delivered"`
	CreatedAt         time.Time   `json:"created_at"`
	AddedAt           time.Time   `json:"added_at"`
}

func (q *Queries) GetPlaylistDishes(ctx context.Context, id int64) ([]GetPlaylistDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPlaylistDishes, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPlaylistDishesRow{}
	for rows.Next() {
		var i GetPlaylistDishesRow
		if err := rows.Scan(
			&i.ID,
			&i.RestaurantID,
			&i.IsAvailable,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.DietType,
			&i.Cuisine,
			&i.ImageUrl,
			&i.ID_2,
			&i.PlaylistID,
			&i.DishID,
			&i.DateToBeDelivered,
			&i.CreatedAt,
			&i.AddedAt,
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