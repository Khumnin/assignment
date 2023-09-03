package main

import (
	"testing"

	"example.com/assignment/reservation/api"
)

// Integration test
// Scenario 1
//
//	Golden PATH : This scenario is the expected behavior without any complicated jouney
func TestScenario1(t *testing.T) {
	// Init with 4 Tables
	res := api.InitTable(4)
	want := true

	if res.IsSuccess != want {
		t.Errorf("The first initialization should pass")
	}

	// Reserve be 7 customers
	bookingId, bookedCount, remaining, result := api.Reserve(7)
	want = true

	if result.IsSuccess != want {
		t.Error("The first reservation should pass")
	}

	if bookedCount != 2 {
		t.Errorf("7 customers need 2 tables but got %d", bookedCount)
	}

	if remaining != 2 {
		t.Errorf("Should get 2 table left but got %d", remaining)
	}

	// Cancel the previos booking
	freedCount, remaining, result := api.Cancelation(bookingId)
	want = true

	if result.IsSuccess != want {
		t.Errorf("This cancelation is correct. It should pass but got %v\n", result.IsSuccess)
		t.Errorf("Reason : %v \n", result.Message)
	}

	if freedCount != 2 {
		t.Errorf("This booking ID reserved 2 tables. Should get 2 but got %d\n", freedCount)
	}

	if remaining != 4 {
		t.Errorf("Should get 4 but got %d\n", remaining)
	}
	api.CleanTable()
}

// Integration test
// Scenario 2
//
//	Over reservation PATH : This scenario has over reservation behavior
func TestScenario2(t *testing.T) {
	// Init with 4 Tables
	res := api.InitTable(4)
	want := true

	if res.IsSuccess != want {
		t.Errorf("The first initialization should pass")
	}

	// Reserve be 7 customers
	bookingId, bookedCount, remaining, result := api.Reserve(7)
	want = true

	if result.IsSuccess != want {
		t.Error("The first reservation should pass")
	}

	if bookedCount != 2 {
		t.Errorf("7 customers need 2 tables but got %d", bookedCount)
	}

	if remaining != 2 {
		t.Errorf("Should get 2 table left but got %d", remaining)
	}

	if len(bookingId) != 6 {
		t.Errorf("Should be 6 lenght but got %d in booking ID", len(bookingId))
	}

	// Another reservation
	bookingId, bookedCount, remaining, result = api.Reserve(10)
	want = false

	if result.IsSuccess != want {
		t.Error("This reservation should fail")
	}

	if bookedCount != 0 {
		t.Errorf("No reservation successfully, should get 0 but got %d", bookedCount)
	}

	if remaining != 2 {
		t.Errorf("No reservation successfully, should get 2 but got %d", remaining)
	}
	api.CleanTable()
}
