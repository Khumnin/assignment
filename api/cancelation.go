package api

import (
	util "example.com/assignment/reservation/api/utility"
	res "example.com/assignment/reservation/model"
)

// Return freedCount, Remaining, Response
func Cancelation(bookingId string) (int, int, res.Response) {

	if tableInstant == nil {
		return 0, 0, res.Response{IsSuccess: false, Message: "The table instant hasn't initialized yet."}
	}

	bookedTables, status := util.GetKeysByValue(tableInstant, bookingId)

	if status == true {
		// Release the table
		for i := 0; i < len(bookedTables); i++ {
			tableInstant[bookedTables[i]] = ""
		}
	} else {
		return 0, 0, res.Response{IsSuccess: false, Message: "BookingId doesn't exist"}
	}

	// Get available table after release
	availableTables, status := util.GetKeysByValue(tableInstant, "")

	return len(bookedTables), len(availableTables), res.Response{IsSuccess: true}
}
