package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
	"gopkg.in/guregu/null.v4"
)

type searchRequest struct {
	UserID        int64  `uri:"userid" binding:"required,min=0"`
	SearchKeyword string `uri:"search"`
}

func (server *Server) searchDishes(ctx *gin.Context) {

	var req searchRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	searchArg := db.CreateSearchParams{
		UserID:  req.UserID,
		Keyword: null.NewString(strings.ToLower(req.SearchKeyword), true),
	}

	_, err := server.store.CreateSearch(ctx, searchArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	searchDishArg := sql.NullString{
		String: req.SearchKeyword,
		Valid:  true}

	searchDishesDB, err := server.store.SearchDishes(ctx, searchDishArg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := server.maptoModelSearchDish(searchDishesDB)

	ctx.JSON(http.StatusOK, resp)
}

type searchDish struct {
	DishID         int64   `json:"dishId"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	ImageURL       string  `json:"imageUrl"`
	RestaurantName string  `json:"restaurantName"`
	RestaurantID   int64   `json:"restaurantId"`
}

func (server *Server) maptoModelSearchDish(searchDishesRow []db.SearchDishesRow) []searchDish {

	var searchDishes []searchDish
	var searchDish searchDish

	for _, searchDishRow := range searchDishesRow {

		searchDish.DishID = searchDishRow.DishID
		searchDish.Name = searchDishRow.DishName
		searchDish.Description = searchDishRow.DishDescription.String
		searchDish.Price = searchDishRow.DishPrice
		searchDish.ImageURL = searchDishRow.DishImageurl.String
		searchDish.RestaurantName = searchDishRow.RestaurantName
		searchDish.RestaurantID = searchDishRow.RestaurantID

		searchDishes = append(searchDishes, searchDish)
	}

	return searchDishes
}
