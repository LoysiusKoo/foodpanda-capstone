package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
	"gopkg.in/guregu/null.v4"
)

type createPlaylistRequest struct {
	Name      string `json:"name" binding:"required"`
	ImageUrl  string `json:"image"`
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
		Image:     null.NewString(req.ImageUrl, true),
		FoodItems: null.NewString(req.FoodItems, true),
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
