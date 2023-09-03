package api

import (
	"math"

	util "example.com/assignment/reservation/api/utility"
	res "example.com/assignment/reservation/model"
)

// Reservation API
// To reserve the table by amount of customer
// Param
//
//	num : Amount of customer
//
// Return
//
//	BookingId : The booking ID of current reservation
//	BookedCount : The number of reserved table
//	Remaining : The number of available table
//	response : response model that contains status and error message (if any)
func Reserve(num int) (string, int, int, res.Response) {

	// Check if the table instant is already initialized
	if tableInstant == nil {
		return "", 0, 0, res.Response{IsSuccess: false, Message: "The table instant hasn't initialized yet."}
	}

	// Search for available table
	availableTables, status := util.GetKeysByValue(tableInstant, "")

	if !status && availableTables == nil {
		return "", 0, 0, res.Response{IsSuccess: false, Message: "No Available tables"}
	}

	// Calculate the amount of table that is need to be reserved in this reservation
	tableCountForThisGroup := int(math.Ceil(float64(num) / float64(seatPerTable)))

	if tableCountForThisGroup > len(availableTables) {
		return "", tableCountForThisGroup, len(availableTables), res.Response{IsSuccess: false, Message: "Available tables is not enough"}
	}

	// Generate booking ID
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
		// Cancel booking by booking ID if the reservation is not success
		Cancelation(bookingId)

		availableTables, status = util.GetKeysByValue(tableInstant, "")
		return "", 0, len(availableTables), res.Response{IsSuccess: false, Message: "Reservation failed"}
	}

}
