package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

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

func (server *Server) getDishesByParams(ctx *gin.Context) {
	var req getDishesByParamsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
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

	ctx.JSON(http.StatusOK, dishes)
}
