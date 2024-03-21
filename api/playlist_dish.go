package api

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lib/pq"

// 	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
// )

// type dish struct {
// 	ID           int64        `form:"id" json:"id"`
// 	RestaurantID int64        `form:"restaurantId" json:"restaurantId"`
// 	IsAvailable  bool         `form:"isAvailable" json:"isAvailable"`
// 	Name         string       `form:"name" json:"name"`
// 	Description  string       `form:"description" json:"description"`
// 	Price        float64      `form:"price" json:"price"`
// 	Cuisine      string       `form:"cuisine" json:"cuisine"`
// 	ImageURL     string       `form:"imageUrl" json:"imageUrl"`
// 	PlaylistDish playlistDish `form:"PlaylistDish" json:"PlaylistDish"`
// }

// type createPlaylistDishRequest struct {
// 	PlaylistID        int64    `form:"playlist_id" json:"playlist_id" binding:"required"`
// 	DateToBeDelivered []string `form:"date_to_be_delivered" json:"date_to_be_delivered"`
// }

// type createPlaylistDishResponse struct {
// 	ID                int64  `json:"id"`
// 	PlaylistID        int64  `json:"playlist_id"`
// 	DishID            int64  `json:"dish_id"`
// 	DateToBeDelivered string `json:"date_to_be_delivered"`
// }

// func (server *Server) createPlaylistDishes(ctx *gin.Context) {
// 	//create playlist_dishes
// 	var req createPlaylistDishRequest
// 	var res createPlaylistDishResponse
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errResponse(err))
// 		return
// 	}

// 	dishes, err := server.store.GetFoodItemsFromPlaylists(ctx, req.PlaylistID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errResponse(err))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 		return
// 	}

// 	arg := db.CreatePlaylistDishParams{
// 		PlaylistID:        req.PlaylistID,
// 		DishID:            dishes.ID,
// 		DateToBeDelivered: req.DateToBeDelivered,
// 	}

// 	user, err := server.store.CreatePlaylistDish(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 		return
// 	}

// 	for _, date := range playlistDishReqs {
// 		playlistDishReq.DateToBeDelivered = date.DateToBeDelivered

// 		for _, dish := range dishes {

// 			params := db.CreatePlaylistDishParams{
// 				PlaylistID:        playlist.ID,
// 				DishID:            dish.ID,
// 				DateToBeDelivered: playlistDishReq.DateToBeDelivered,
// 			}

// 			_, err = server.store.CreatePlaylistDish(ctx, params)
// 			if err != nil {
// 				if pqErr, ok := err.(*pq.Error); ok {
// 					switch pqErr.Code.Name() {
// 					case "foreign_key_violation", "unique_violation":
// 						ctx.JSON(http.StatusForbidden, errResponse(err))
// 						return
// 					}
// 				}
// 				ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 				return
// 			}

// 			rsp := createUserResponse{
// 				ID:       user.ID,
// 				Username: user.Username,
// 				Email:    user.Email,
// 				Address:  user.Address,
// 			}

// 			ctx.JSON(http.StatusOK, rsp)
// 		}
// 	}

// }
