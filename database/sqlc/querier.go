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
	CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error)
	CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (Restaurant, error)
	DeleteDish(ctx context.Context, id int64) error
	DeletePlaylist(ctx context.Context, id int64) error
	DeleteRestaurant(ctx context.Context, id int64) error
	GetDish(ctx context.Context, id int64) (Dish, error)
	GetDishes(ctx context.Context) ([]Dish, error)
	GetDishesByCuisine(ctx context.Context, dollar_1 sql.NullString) ([]Dish, error)
	GetPlaylist(ctx context.Context, id int64) (Playlist, error)
	GetRestaurantByID(ctx context.Context, id int64) (Restaurant, error)
	GetRestaurants(ctx context.Context) ([]Restaurant, error)
	ListDishes(ctx context.Context, arg ListDishesParams) ([]Dish, error)
	ListPlaylists(ctx context.Context, arg ListPlaylistsParams) ([]Playlist, error)
	ListRestaurants(ctx context.Context, arg ListRestaurantsParams) ([]Restaurant, error)
	UpdateDish(ctx context.Context, arg UpdateDishParams) (Dish, error)
	UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) (Playlist, error)
	UpdateRestaurant(ctx context.Context, arg UpdateRestaurantParams) (Restaurant, error)
}

var _ Querier = (*Queries)(nil)
