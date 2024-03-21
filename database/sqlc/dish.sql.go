// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: dish.sql

package db

import (
	"context"
	"database/sql"

	null "gopkg.in/guregu/null.v4"
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
RETURNING id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url
`

type CreateDishParams struct {
	RestaurantID int64       `json:"restaurant_id"`
	IsAvailable  bool        `json:"is_available"`
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	Price        float64     `json:"price"`
	Cuisine      string      `json:"cuisine"`
	ImageUrl     string      `json:"image_url"`
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
		&i.DietType,
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

const getD = `-- name: GetD :many
SELECT id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url
FROM dishes
WHERE cuisine = $1 
AND diet_type = $2
AND price BETWEEN $3 AND $4
ORDER BY id
`

type GetDParams struct {
	Cuisine  string  `json:"cuisine"`
	DietType string  `json:"diet_type"`
	Price    float64 `json:"price"`
	Price_2  float64 `json:"price_2"`
}

func (q *Queries) GetD(ctx context.Context, arg GetDParams) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getD,
		arg.Cuisine,
		arg.DietType,
		arg.Price,
		arg.Price_2,
	)
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
			&i.DietType,
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

const getDish = `-- name: GetDish :one
SELECT id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url FROM dishes
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
		&i.DietType,
		&i.Cuisine,
		&i.ImageUrl,
	)
	return i, err
}

const getDishes = `-- name: GetDishes :many
SELECT id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url FROM dishes
ORDER BY id
`

func (q *Queries) GetDishes(ctx context.Context) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getDishes)
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
			&i.DietType,
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

const getDishesByCuisine = `-- name: GetDishesByCuisine :many
SELECT id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url FROM dishes
WHERE cuisine ILIKE '%'||$1||'%'
ORDER BY id
`

func (q *Queries) GetDishesByCuisine(ctx context.Context, dollar_1 sql.NullString) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getDishesByCuisine, dollar_1)
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
			&i.DietType,
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

const getDishesByParams = `-- name: GetDishesByParams :many
SELECT d.id, restaurant_id, is_available, d.name, d.description, price, diet_type, cuisine, d.image_url, r.id, r.name, r.description, address, rating, restaurant_type, num_of_reviews, r.image_url
FROM dishes d JOIN restaurants r ON d.restaurant_id = r.id
WHERE d.cuisine ILIKE '%'||$1||'%'
AND d.diet_type ILIKE '%'||$2||'%'
AND d.price BETWEEN $3 AND $4
AND r.rating >= $5
ORDER BY r.id
LIMIT 10
`

type GetDishesByParamsParams struct {
	Column1 sql.NullString `json:"column_1"`
	Column2 sql.NullString `json:"column_2"`
	Price   float64        `json:"price"`
	Price_2 float64        `json:"price_2"`
	Rating  float64        `json:"rating"`
}

type GetDishesByParamsRow struct {
	ID             int64       `json:"id"`
	RestaurantID   int64       `json:"restaurant_id"`
	IsAvailable    bool        `json:"is_available"`
	Name           string      `json:"name"`
	Description    null.String `json:"description"`
	Price          float64     `json:"price"`
	DietType       string      `json:"diet_type"`
	Cuisine        string      `json:"cuisine"`
	ImageUrl       string      `json:"image_url"`
	ID_2           int64       `json:"id_2"`
	Name_2         string      `json:"name_2"`
	Description_2  null.String `json:"description_2"`
	Address        null.String `json:"address"`
	Rating         float64     `json:"rating"`
	RestaurantType string      `json:"restaurant_type"`
	NumOfReviews   int32       `json:"num_of_reviews"`
	ImageUrl_2     null.String `json:"image_url_2"`
}

func (q *Queries) GetDishesByParams(ctx context.Context, arg GetDishesByParamsParams) ([]GetDishesByParamsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDishesByParams,
		arg.Column1,
		arg.Column2,
		arg.Price,
		arg.Price_2,
		arg.Rating,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetDishesByParamsRow{}
	for rows.Next() {
		var i GetDishesByParamsRow
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
			&i.Name_2,
			&i.Description_2,
			&i.Address,
			&i.Rating,
			&i.RestaurantType,
			&i.NumOfReviews,
			&i.ImageUrl_2,
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

const listDishes = `-- name: ListDishes :many
SELECT id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url FROM dishes
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
			&i.DietType,
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
RETURNING id, restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url
`

type UpdateDishParams struct {
	ID           int64       `json:"id"`
	RestaurantID int64       `json:"restaurant_id"`
	IsAvailable  bool        `json:"is_available"`
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	Price        float64     `json:"price"`
	Cuisine      string      `json:"cuisine"`
	ImageUrl     string      `json:"image_url"`
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
		&i.DietType,
		&i.Cuisine,
		&i.ImageUrl,
	)
	return i, err
}
