package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) getUpcomingDelivery(ctx *gin.Context) {

	deliveryDates, err := server.store.GetDeliveryDates(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	earliestDate1, err := getEarliestDeliveryDate(deliveryDates)
	if err != nil {
		fmt.Println("Error parsing delivery date:", err)
		return
	}

	earliestDateStr := earliestDate1.Format("2 Jan, 03:04 pm")

	var req deletePlaylistDishRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	//GET upcoming delivery
	upcomingdelivery, err := server.store.GetUpcomingDelivery(ctx, earliestDateStr)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, upcomingdelivery)
}

func getEarliestDeliveryDate(deliveryDates []string) (time.Time, error) {
	// Define the date format
	dateFormat := "2 Jan, 03:04 pm"
	var parsedDates []time.Time

	// Parse the string delivery dates to time.Time objects for comparison
	for _, dateStr := range deliveryDates {
		parsedDate, err := time.Parse(dateFormat, dateStr)
		if err != nil {
			return time.Time{}, err
		}
		parsedDates = append(parsedDates, parsedDate)
	}

	// Find the earliest delivery date
	var earliestDate = parsedDates[0]
	for _, date := range parsedDates {
		if date.Before(earliestDate) {
			earliestDate = date
		}
	}

	return earliestDate, nil
}

func (server *Server) getDelivery(ctx *gin.Context) {

	deliveryDates, err := server.store.GetDeliveryDates(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, deliveryDates)

}
