// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

// Stores cuisine
type Cuisine struct {
	ID      int64          `json:"id"`
	Cuisine sql.NullString `json:"cuisine"`
}

// Stores dishes
type Dish struct {
	ID           int64          `json:"id"`
	RestaurantID int64          `json:"restaurant_id"`
	IsAvailable  bool           `json:"is_available"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Price        string         `json:"price"`
	Cuisine      sql.NullString `json:"cuisine"`
	ImageUrl     sql.NullString `json:"image_url"`
}

// Stores restaurants
type Restaurant struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description sql.NullString  `json:"description"`
	Address     sql.NullString  `json:"address"`
	Rating      sql.NullFloat64 `json:"rating"`
	Cuisine     sql.NullString  `json:"cuisine"`
	ImageUrl    sql.NullString  `json:"image_url"`
}
