package api

import (
	util "example.com/assignment/reservation/api/utility"
	res "example.com/assignment/reservation/model"
)

// Cancelation API
// To cancel the reservation by booking ID
// Param
//
//	bookingId : Booking ID in string
//
// Return
//
//	freedCount : The number of table that is released
//	remaining : The number of available table
//	response : response model that contains status and error message (if any)
func Cancelation(bookingId string) (int, int, res.Response) {

	// Check if the table instant is already initialized
	if tableInstant == nil {
		return 0, 0, res.Response{IsSuccess: false, Message: "The table instant hasn't initialized yet."}
	}

	// Search for booked table by booking ID
	bookedTables, status := util.GetKeysByValue(tableInstant, bookingId)

	if status == true {
		// Release the table by setting the value to ""
		for i := 0; i < len(bookedTables); i++ {
			tableInstant[bookedTables[i]] = ""
		}
	} else {
		return 0, 0, res.Response{IsSuccess: false, Message: "BookingId doesn't exist"}
	}

	// Get the lastest available table after release
	availableTables, status := util.GetKeysByValue(tableInstant, "")

	return len(bookedTables), len(availableTables), res.Response{IsSuccess: true}
}
