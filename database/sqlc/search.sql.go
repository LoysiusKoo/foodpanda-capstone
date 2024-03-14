// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: search.sql

package db

import (
	"context"
	"database/sql"

	null "gopkg.in/guregu/null.v4"
)

const createSearch = `-- name: CreateSearch :one
INSERT INTO searches (
  user_id,
  keyword
) 
VALUES (
  $1, $2
)
RETURNING id, user_id, keyword
`

type CreateSearchParams struct {
	UserID  int64       `json:"user_id"`
	Keyword null.String `json:"keyword"`
}

func (q *Queries) CreateSearch(ctx context.Context, arg CreateSearchParams) (Search, error) {
	row := q.db.QueryRowContext(ctx, createSearch, arg.UserID, arg.Keyword)
	var i Search
	err := row.Scan(&i.ID, &i.UserID, &i.Keyword)
	return i, err
}

const getSearch = `-- name: GetSearch :one
SELECT id, user_id, keyword FROM searches
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSearch(ctx context.Context, id int64) (Search, error) {
	row := q.db.QueryRowContext(ctx, getSearch, id)
	var i Search
	err := row.Scan(&i.ID, &i.UserID, &i.Keyword)
	return i, err
}

const listSearches = `-- name: ListSearches :many
SELECT id, user_id, keyword FROM searches
WHERE id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3
`

type ListSearchesParams struct {
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSearches(ctx context.Context, arg ListSearchesParams) ([]Search, error) {
	rows, err := q.db.QueryContext(ctx, listSearches, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Search{}
	for rows.Next() {
		var i Search
		if err := rows.Scan(&i.ID, &i.UserID, &i.Keyword); err != nil {
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

const searchDishes = `-- name: SearchDishes :many


SELECT d.id AS dish_id,
       d.name AS dish_name,
       d.description AS dish_description,
       d.price AS dish_price,
       d.image_url AS dish_imageURL,
       r.name AS restaurant_name,
       r.id AS restaurant_id
FROM dishes d
JOIN restaurants r ON d.restaurant_id = r.id
WHERE d.name ILIKE '%'||$1||'%' OR d.description ILIKE '%'||$1||'%'
`

type SearchDishesRow struct {
	DishID          int64       `json:"dish_id"`
	DishName        string      `json:"dish_name"`
	DishDescription null.String `json:"dish_description"`
	DishPrice       float64     `json:"dish_price"`
	DishImageurl    null.String `json:"dish_imageurl"`
	RestaurantName  string      `json:"restaurant_name"`
	RestaurantID    int64       `json:"restaurant_id"`
}

// -- name: UpdateSearch :one
// UPDATE searches
// SET
//
//	user_id = $2,
//	keyword = $3
//
// WHERE
//
//	id = $1
//
// RETURNING *;
// -- name: DeleteSearch :exec
// DELETE FROM searches
// WHERE id = $1;
func (q *Queries) SearchDishes(ctx context.Context, dollar_1 sql.NullString) ([]SearchDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, searchDishes, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchDishesRow{}
	for rows.Next() {
		var i SearchDishesRow
		if err := rows.Scan(
			&i.DishID,
			&i.DishName,
			&i.DishDescription,
			&i.DishPrice,
			&i.DishImageurl,
			&i.RestaurantName,
			&i.RestaurantID,
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
