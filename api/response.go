package api

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
// )

// type playlistDish struct {
// 	ID                int64     `form:"id" json:"id"`
// 	PlaylistID        int64     `form:"playlistId" json:"playlistId"`
// 	DishID            int64     `form:"dishId" json:"dishId"`
// 	DateToBeDelivered time.Time `form:"dateToBeDelivered" json:"dateToBeDelivered"`
// 	CreatedAt         time.Time `form:"createdAt" json:"createdAt"`
// 	UpdatedAt         time.Time `form:"updatedAt" json:"updatedAt"`
// }

// type playlistv2 struct {
// 	ID          int64  `form:"id" json:"id"`
// 	Name        string `form:"name" json:"name"`
// 	ImageURL    string `form:"imageUrl" json:"imageUrl"`
// 	DeliveryDay string `form:"deliveryDay" json:"deliveryDay"`
// }

// func (server *Server) maptoModelV2(ctx *gin.Context, playlistsDB []db.Playlist) ([]playlistv2, error) {
// 	var playlists []playlistv2

// 	for _, playlistDB := range playlistsDB {
// 		var playlist playlistv2
// 		var dish dish
// 		var cost float64

// 		playlistdishesDB, err := server.store.getDishesByParams(ctx, playlistDB.ID)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				ctx.JSON(http.StatusNotFound, errResponse(err))
// 				return nil, err
// 			}
// 			ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 			return nil, err
// 		}

// 		for _, playlistdishDB := range playlistdishesDB {
// 			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
// 			if err != nil {
// 				if err == sql.ErrNoRows {
// 					ctx.JSON(http.StatusNotFound, errResponse(err))
// 					return nil, err
// 				}
// 				ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 				return nil, err
// 			}

// 			dish.ID = dishDB.ID
// 			dish.RestaurantID = dishDB.RestaurantID
// 			dish.IsAvailable = dishDB.IsAvailable
// 			dish.Name = dishDB.Name
// 			dish.Price = dishDB.Price
// 			dish.Cuisine = dishDB.Cuisine.String
// 			dish.ImageURL = dishDB.ImageUrl.String
// 			dish.PlaylistDish = playlistDish{
// 				ID:           playlistdishDB.ID,
// 				PlaylistID:   playlistdishDB.PlaylistID,
// 				DishID:       playlistdishDB.DishID,
// 				DishQuantity: playlistdishDB.DishQuantity,
// 				CreatedAt:    playlistdishDB.CreatedAt,
// 				UpdatedAt:    playlistdishDB.AddedAt,
// 			}

// 			cost += dishDB.Price * float64(playlistdishDB.DishQuantity)
// 		}

// 		playlist.ID = playlistDB.ID
// 		playlist.Name = playlistDB.Name
// 		playlist.ImageURL = playlistDB.ImageUrl.String
// 		playlist.IsPublic = playlistDB.IsPublic
// 		playlist.DeliveryDay = playlistDB.DeliveryDay.String
// 		playlist.Category = playlistDB.Category.String
// 		playlist.Cost = fmt.Sprintf("%.2f", cost)

// 		status, _ := server.store.ListStatusByPlaylistID(ctx, playlistDB.ID)
// 		playlist.Status = status.String

// 		playlists = append(playlists, playlist)
// 	}
// 	return playlists, nil
// }
