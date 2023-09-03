package api

import (
	"math"

	util "example.com/assignment/reservation/api/utility"
	res "example.com/assignment/reservation/model"
)

// Return BookingId, BookedCount, Remaining, Response
func Reserve(num int) (string, int, int, res.Response) {

	if tableInstant == nil {
		return "", 0, 0, res.Response{IsSuccess: false, Message: "The table instant hasn't initialized yet."}
	}

	availableTables, status := util.GetKeysByValue(tableInstant, "")

	if !status {
		return "", 0, 0, res.Response{IsSuccess: false, Message: "Failed to get available tables"}
	}

	if len(availableTables) == 0 {
		return "", 0, 0, res.Response{IsSuccess: false, Message: "No Available tables"}
	}

	tableCountForThisGroup := int(math.Ceil(float64(num) / float64(seatPerTable)))

	if tableCountForThisGroup > len(availableTables) {
		return "", tableCountForThisGroup, len(availableTables), res.Response{IsSuccess: false, Message: "Available tables is not enough"}
	}

	bookingId := util.GenerateBookingId(6)

	// Assign booking Id value
	for i := 0; i < int(tableCountForThisGroup); i++ {
		tableInstant[availableTables[i]] = bookingId
	}

	validateReserved := len(availableTables) - tableCountForThisGroup
	availableTables, status = util.GetKeysByValue(tableInstant, "")

	// Validate if the reservation is success
	if validateReserved == len(availableTables) {
		return bookingId, tableCountForThisGroup, validateReserved, res.Response{IsSuccess: true}
	} else {
		// Calcel booking by booking ID

		availableTables, status = util.GetKeysByValue(tableInstant, "")
		return "", 0, len(availableTables), res.Response{IsSuccess: false, Message: "Reservation failed"}
	}

}
