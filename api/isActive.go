package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
)

type updateIsActiveRequest struct {
	PlaylistID int64 `uri:"id" binding:"required,min=0"`
	IsActive   bool  `uri:"is_active"`
}

func (server *Server) updateIsActive(ctx *gin.Context) {
	var req updateIsActiveRequest
	var arg db.UpdateIsActiveParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if req.IsActive {
		req.IsActive = false
	} else {
		req.IsActive = true
	}

	arg = db.UpdateIsActiveParams{
		ID:       req.PlaylistID,
		IsActive: req.IsActive,
	}
	playlist, err := server.store.UpdateIsActive(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}
