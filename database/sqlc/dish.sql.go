// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: dish.sql

package db

import (
	"context"
	"database/sql"
)

const createDish = `-- name: CreateDish :one
INSERT INTO dishes (
  restaurant_id,
  is_available,
  name,
  description,
  price,
  cuisine,
  image_url

) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, restaurant_id, is_available, name, description, price, cuisine, image_url
`

type CreateDishParams struct {
	RestaurantID int64          `json:"restaurant_id"`
	IsAvailable  bool           `json:"is_available"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Price        string         `json:"price"`
	Cuisine      sql.NullString `json:"cuisine"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) CreateDish(ctx context.Context, arg CreateDishParams) (Dish, error) {
	row := q.db.QueryRowContext(ctx, createDish,
		arg.RestaurantID,
		arg.IsAvailable,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Cuisine,
		arg.ImageUrl,
	)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.RestaurantID,
		&i.IsAvailable,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Cuisine,
		&i.ImageUrl,
	)
	return i, err
}

const deleteDish = `-- name: DeleteDish :exec
DELETE FROM dishes
WHERE id = $1
`

func (q *Queries) DeleteDish(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteDish, id)
	return err
}

const getDish = `-- name: GetDish :one
SELECT id, restaurant_id, is_available, name, description, price, cuisine, image_url FROM dishes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDish(ctx context.Context, id int64) (Dish, error) {
	row := q.db.QueryRowContext(ctx, getDish, id)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.RestaurantID,
		&i.IsAvailable,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Cuisine,
		&i.ImageUrl,
	)
	return i, err
}

const listDishes = `-- name: ListDishes :many
SELECT id, restaurant_id, is_available, name, description, price, cuisine, image_url FROM dishes
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListDishesParams struct {
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDishes(ctx context.Context, arg ListDishesParams) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, listDishes, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dish{}
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.RestaurantID,
			&i.IsAvailable,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Cuisine,
			&i.ImageUrl,
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

const updateDish = `-- name: UpdateDish :one
UPDATE dishes
SET 
  restaurant_id = $2,
  is_available = $3,
  name = $4,
  description = $5,
  price = $6,
  cuisine = $7,
  image_url = $8
WHERE 
  id = $1
RETURNING id, restaurant_id, is_available, name, description, price, cuisine, image_url
`

type UpdateDishParams struct {
	ID           int64          `json:"id"`
	RestaurantID int64          `json:"restaurant_id"`
	IsAvailable  bool           `json:"is_available"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Price        string         `json:"price"`
	Cuisine      sql.NullString `json:"cuisine"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) UpdateDish(ctx context.Context, arg UpdateDishParams) (Dish, error) {
	row := q.db.QueryRowContext(ctx, updateDish,
		arg.ID,
		arg.RestaurantID,
		arg.IsAvailable,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Cuisine,
		arg.ImageUrl,
	)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.RestaurantID,
		&i.IsAvailable,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Cuisine,
		&i.ImageUrl,
	)
	return i, err
}
