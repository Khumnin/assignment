package handler

import (
	"net/http"
	"strconv"

	"example.com/assignment/reservation/api"
	"example.com/assignment/reservation/model"
	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
}

func (h *ReservationHandler) InitialTable(c *gin.Context) {
	numOfTable, err := strconv.Atoi(c.Param("num"))

	if err != nil {
		c.JSON(http.StatusOK, model.Response{IsSuccess: false, Message: err.Error()})
	} else {
		res := api.InitTable(numOfTable)
		c.JSON(http.StatusOK, res)
	}
}

func (h *ReservationHandler) ReserveTable(c *gin.Context) {
	numOfCustomer, err := strconv.Atoi(c.Param("customerCount"))

	if err != nil {
		c.JSON(http.StatusOK, model.Response{IsSuccess: false, Message: err.Error()})
	} else {
		bookingId, bookedCount, remaining, res := api.Reserve(numOfCustomer)

		if res.IsSuccess == false {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusOK, model.Reservation{IsSuccess: true, BookingId: bookingId, NumOfBooked: bookedCount, Remaining: remaining})
		}

	}
}

func (h *ReservationHandler) CancelReservation(c *gin.Context) {
	freedCount, remaining, res := api.Cancelation(c.Param("bookingId"))

	if res.IsSuccess == false {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, model.Cancelation{IsSuccess: true, NumOfFreed: freedCount, Remaining: remaining})
	}
}
