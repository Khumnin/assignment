package main

import (
	"log"

	"example.com/assignment/reservation/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/healthcheck", handler.HealthCheckHandler)

	v1 := r.Group("/api/v1")
	{
		reservationHandler := handler.ReservationHandler{}
		v1.POST("/initial/:num", reservationHandler.InitialTable)
		v1.POST("/reserve/:customerCount", reservationHandler.ReserveTable)
		v1.POST("/cancel/:bookingId", reservationHandler.CancelReservation)
	}

	log.Fatal((r.Run(":8080")))
}
