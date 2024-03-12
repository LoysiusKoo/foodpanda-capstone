package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
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
	cuisine sql.NullString `uri:"cuisinename"`
}

func (server *Server) getDishesByCuisine(ctx *gin.Context) {
	var req getDishesByCuisineRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	dishes, err := server.store.GetDishesByCuisine(ctx, req.cuisine)
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