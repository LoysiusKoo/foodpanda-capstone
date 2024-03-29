package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getRestaurantsRequest struct {
}

func (server *Server) getRestaurants(ctx *gin.Context) {
	var req getRestaurantsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	restaurants, err := server.store.GetRestaurants(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, restaurants)
}

type getRestaurantByIDRequest struct {
	ID int64 `uri:"restaurantid" binding:"required,min=1"`
}

func (server *Server) getRestaurantByID(ctx *gin.Context) {
	var req getRestaurantByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	restaurant, err := server.store.GetRestaurantByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, restaurant)
}
