package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getRestaurantRequest struct {
	ID int64 `uri:"restaurantid" binding:"required,min=1"`
}

func (server *Server) getRestaurant(ctx *gin.Context) {
	var req getRestaurantRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	restaurant, err := server.store.GetRestaurant(ctx, req.ID)
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
