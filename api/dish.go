package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
)

type getDishesRequest struct {
}

func (server *Server) getDishes(ctx *gin.Context) {
	var req getDishesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	dishes, err := server.store.GetDishes(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

type getDishesByCuisineRequest struct {
	Cuisine string `uri:"cuisinename"`
}

func (server *Server) getDishesByCuisine(ctx *gin.Context) {
	var req getDishesByCuisineRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	keyword := sql.NullString{
		String: strings.ToLower(req.Cuisine),
		Valid:  true}

	dishes, err := server.store.GetDishesByCuisine(ctx, keyword)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

type getDRequest struct {
	Cuisine  string  `uri:"cuisinename"`
	Type     string  `uri:"type"`
	SmallNum float64 `uri:"smallnum"`
	BigNum   float64 `uri:"bignum"`
}

func (server *Server) getD(ctx *gin.Context) {
	var req getDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	keyword := db.GetDParams{
		Cuisine:  req.Cuisine,
		DietType: req.Type,
		Price:    req.SmallNum,
		Price_2:  req.BigNum,
	}

	dishes, err := server.store.GetD(ctx, keyword)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

type getDishesByParamsRequest struct {
	Cuisine  string  `uri:"cuisinename"`
	Type     string  `uri:"type"`
	SmallNum float64 `uri:"smallnum"`
	BigNum   float64 `uri:"bignum"`
	Rating   float64 `uri:"rating"`
}

func (server *Server) CreatePlaylistByParams(ctx *gin.Context) {
	var req getDishesByParamsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if req.Cuisine == "Anything" {
		req.Cuisine = ""
	}

	keyword := db.GetDishesByParamsParams{
		Column1: sql.NullString{
			String: strings.ToLower(req.Cuisine),
			Valid:  true},
		Column2: sql.NullString{
			String: strings.ToLower(req.Type),
			Valid:  true},
		Price:   req.SmallNum,
		Price_2: req.BigNum,
		Rating:  req.Rating,
	}

	dishes, err := server.store.GetDishesByParams(ctx, keyword)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	//map dishes to string
	dishesString, _ := json.Marshal(dishes)
	str := string(dishesString)

	if dishes[0].ImageUrl == "placeholder_image_url" {
		dishes[0].ImageUrl = "https://sm.mashable.com/mashable_sea/photo/default/mm20220106-4-1_b67x.jpg"
	}

	if req.Cuisine == "" {
		req.Cuisine = "Anything"
	}

	//create playlist
	arg := db.CreatePlaylistParams{
		Name:      req.Cuisine,
		Image:     dishes[0].ImageUrl,
		FoodItems: str,
		IsActive:  true,
	}

	playlist, err := server.store.CreatePlaylist(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	//Create playlist_dishes
	type createPlaylistDishRequest []struct {
		DateToBeDelivered string `json:"date_to_be_delivered"`
	}

	var date createPlaylistDishRequest
	if err := ctx.ShouldBindJSON(&date); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	for i, dish := range dishes {

		if dish.ImageUrl == "placeholder_image_url" {
			dish.ImageUrl = "https://sm.mashable.com/mashable_sea/photo/default/mm20220106-4-1_b67x.jpg"
		}

		params := db.CreatePlaylistDishParams{
			PlaylistID:        playlist.ID,
			DishID:            dish.ID,
			DateToBeDelivered: date[i].DateToBeDelivered,
			ImageUrl:          dish.ImageUrl,
		}

		_, err = server.store.CreatePlaylistDish(ctx, params)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "foreign_key_violation", "unique_violation":
					ctx.JSON(http.StatusForbidden, errResponse(err))
					return
				}
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	}
	//return playlist dishes
	playlistDishes, err := server.store.GetPlaylistDishes(ctx, playlist.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlistDishes)
}
