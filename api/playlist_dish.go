package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deletePlaylistDishRequest struct {
	ID int `uri:"id"`
}

func (server *Server) deletePlaylistDish(ctx *gin.Context) {
	var req deletePlaylistDishRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	deletedDish, err := server.store.DeletePlaylistDish(ctx, int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, deletedDish)
}
