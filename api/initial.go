package api

import (
	"fmt"
	"sync"

	res "example.com/assignment/reservation/model"
)

// Initialize singleton
var lock = &sync.Mutex{}

type singleton map[int]string

var tableInstant singleton

var seatPerTable int
var duplicateError = "The table instant is already initialized."

// Initial Table API
// To initial the first use with the number of table
// Param
//
//	num : The amount of table that is going to be initialized
//
// Retrun
//
//	response : response model that contains status and error message (if any)
func InitTable(num int) res.Response {

	// Create singleton table instant for the first time
	if tableInstant == nil {
		lock.Lock()
		defer lock.Unlock()
		if tableInstant == nil {
			fmt.Println("Creating single instance now.")
			tableInstant = make(singleton, num)

			if tableInstant == nil {
				return res.Response{IsSuccess: false, Message: "Failed to create table instant"}
			}

		} else {
			fmt.Println(duplicateError)
			return res.Response{IsSuccess: false, Message: duplicateError}
		}
	} else {
		fmt.Println(duplicateError)
		return res.Response{IsSuccess: false, Message: duplicateError}
	}

	// Assign default value to map
	for i := 0; i < num; i++ {
		tableInstant[i+1] = ""
	}

	seatPerTable = 4
	return res.Response{IsSuccess: true}
}

func CleanTable() bool {
	tableInstant = nil
	return true
}
