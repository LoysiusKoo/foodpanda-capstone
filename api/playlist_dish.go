package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
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

	//Move dishes with later delivery dates up

	playlistID, err := server.store.GetPlaylistFromPlaylistDishID(ctx, int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	playlistDishes, err := server.store.GetPlaylistDishes(ctx, int64(playlistID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	lengthOfPlaylistDishes := len(playlistDishes)

	for i, dish := range playlistDishes {
		if dish.PlaylistDishesid != int64(req.ID) {
			lengthOfPlaylistDishes--
		}
		if dish.PlaylistDishesid == int64(req.ID) {
			for j := i; j < lengthOfPlaylistDishes+1; j++ {
				arg := db.UpdateDeliveryDateParams{
					ID:                playlistDishes[j+1].PlaylistDishesid,
					DateToBeDelivered: playlistDishes[j].DateToBeDelivered,
				}

				_, err := server.store.UpdateDeliveryDate(ctx, arg)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errResponse(err))
					return
				}
			}
		}
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
