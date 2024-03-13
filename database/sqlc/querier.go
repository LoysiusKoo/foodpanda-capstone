// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateDish(ctx context.Context, arg CreateDishParams) (Dish, error)
	CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (Restaurant, error)
	DeleteDish(ctx context.Context, id int64) error
	DeleteRestaurant(ctx context.Context, id int64) error
	GetDish(ctx context.Context, id int64) (Dish, error)
	GetDishes(ctx context.Context) ([]Dish, error)
	GetDishesByCuisine(ctx context.Context, dollar_1 sql.NullString) ([]Dish, error)
	GetRestaurantByID(ctx context.Context, id int64) (Restaurant, error)
	GetRestaurants(ctx context.Context) ([]Restaurant, error)
	ListDishes(ctx context.Context, arg ListDishesParams) ([]Dish, error)
	ListRestaurants(ctx context.Context, arg ListRestaurantsParams) ([]Restaurant, error)
	UpdateDish(ctx context.Context, arg UpdateDishParams) (Dish, error)
	UpdateRestaurant(ctx context.Context, arg UpdateRestaurantParams) (Restaurant, error)
}

var _ Querier = (*Queries)(nil)
