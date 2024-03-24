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
	if err == sql.ErrNoRows {
		fmt.Println("No upcoming deliveries found:", err)
		ctx.JSON(http.StatusNoContent, deliveryDates)
		return
	}

	earliestDate1, err := getEarliestDeliveryDate(deliveryDates)
	if err != nil {
		fmt.Println("Error parsing delivery date:", err)
		return
	}

	earliestDateStr := earliestDate1.Format("2 Jan, 03:04 pm")

	//GET upcoming delivery
	upcomingdelivery, err := server.store.GetUpcomingDelivery(ctx, earliestDateStr)
	if err == sql.ErrNoRows {
		fmt.Println("No upcoming deliveries found:", err)
		ctx.JSON(http.StatusNoContent, upcomingdelivery)
	}

	ctx.JSON(http.StatusOK, upcomingdelivery)
}

func getEarliestDeliveryDate(deliveryDates []string) (time.Time, error) {
	// Define the date format
	dateFormat := "2 Jan, 03:04 pm"
	var parsedDates []time.Time

	// Parse the string delivery dates to time.Time objects for comparison
	if len(deliveryDates) != 0 {
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
	} else {
		return time.Time{}, nil
	}
}
