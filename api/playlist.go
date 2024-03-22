package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
)

type createPlaylistRequest struct {
	Name      string `json:"name" binding:"required"`
	Image     string `json:"image"`
	FoodItems string `json:"food_items"`
	IsActive  bool   `json:"is_active"`
}

func (server *Server) createPlaylist(ctx *gin.Context) {
	var req createPlaylistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreatePlaylistParams{
		Name:      req.Name,
		Image:     req.Image,
		FoodItems: req.FoodItems,
		IsActive:  req.IsActive,
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
	ctx.JSON(http.StatusOK, playlist)
}

type createPlaylistParamsRequest struct {
	Cuisine  string  `uri:"cuisinename"`
	Type     string  `uri:"type"`
	SmallNum float64 `uri:"smallnum"`
	BigNum   float64 `uri:"bignum"`
	Rating   float64 `uri:"rating"`
}

func (server *Server) createPlaylistWithParams(ctx *gin.Context) {
	var req createPlaylistParamsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
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

	dishesString, err := json.Marshal(dishes)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	dishesAsStr := string(dishesString)

	arg := db.CreatePlaylistParams{
		Name:      dishes[0].Cuisine,
		Image:     dishes[0].ImageUrl,
		FoodItems: dishesAsStr,
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
	ctx.JSON(http.StatusOK, playlist)

}

type getAllPlaylistRequest struct {
}

func (server *Server) getAllPlaylist(ctx *gin.Context) {
	var req getAllPlaylistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	playlists, err := server.store.GetAllPlaylists(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

type getPlaylistRequest struct {
	ID int `uri:"id"`
}

func (server *Server) getPlaylist(ctx *gin.Context) {
	var req getPlaylistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	playlist, err := server.store.GetPlaylistDishes(ctx, int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}
