package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	api "example.com/assignment/reservation/api"
	model "example.com/assignment/reservation/model"
	"github.com/gin-gonic/gin"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	r := gin.New()

	r.POST("/initial/:num", func(c *gin.Context) {
		numOfTable, err := strconv.Atoi(c.Param("num"))

		if err != nil {
			c.JSON(http.StatusOK, model.Response{IsSuccess: false, Message: err.Error()})
		} else {
			res := api.InitTable(numOfTable)
			c.JSON(http.StatusOK, res)
		}

	})

	r.POST("/reserve/:customerCount", func(c *gin.Context) {
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
	})

	r.POST("/cancel/:bookingId", func(c *gin.Context) {

		freedCount, remaining, res := api.Cancelation(c.Param("bookingId"))

		if res.IsSuccess == false {
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusOK, model.Cancelation{IsSuccess: true, NumOfFreed: freedCount, Remaining: remaining})
		}

	})

	r.Run()
	//handleRequests()
}
