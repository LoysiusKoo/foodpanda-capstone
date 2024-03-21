// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	null "gopkg.in/guregu/null.v4"
)

// Stores cuisine
type Cuisine struct {
	ID      int64       `json:"id"`
	Cuisine null.String `json:"cuisine"`
}

// Stores dishes
type Dish struct {
	ID           int64       `json:"id"`
	RestaurantID int64       `json:"restaurant_id"`
	IsAvailable  bool        `json:"is_available"`
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	Price        float64     `json:"price"`
	DietType     string      `json:"diet_type"`
	Cuisine      string      `json:"cuisine"`
	ImageUrl     string      `json:"image_url"`
}

// Stores user playlist
type Playlist struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	FoodItems string    `json:"food_items"`
	IsActive  bool      `json:"is_active"`
	CreatedOn time.Time `json:"created_on"`
}

type PlaylistDish struct {
	ID                int64     `json:"id"`
	PlaylistID        int64     `json:"playlist_id"`
	DishID            int64     `json:"dish_id"`
	DateToBeDelivered string    `json:"date_to_be_delivered"`
	ImageUrl          string    `json:"image_url"`
	CreatedAt         time.Time `json:"created_at"`
	AddedAt           time.Time `json:"added_at"`
}

// Stores restaurants
type Restaurant struct {
	ID             int64       `json:"id"`
	Name           string      `json:"name"`
	Description    null.String `json:"description"`
	Address        null.String `json:"address"`
	Rating         float64     `json:"rating"`
	RestaurantType string      `json:"restaurant_type"`
	NumOfReviews   int32       `json:"num_of_reviews"`
	ImageUrl       null.String `json:"image_url"`
}

type Search struct {
	ID      int64       `json:"id"`
	UserID  int64       `json:"user_id"`
	Keyword null.String `json:"keyword"`
}
