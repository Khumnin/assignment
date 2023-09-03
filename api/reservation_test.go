package api

import (
	"fmt"
	"testing"
)

func TestReservationWithoutInit(t *testing.T) {
	count := 5
	// Reserve table without initialization shoule fail
	want := false

	bookingId, bookedCount, remainging, res := Reserve(count)
	if res.IsSuccess != want {
		t.Errorf("The reservation shouldn't pass without initialization but got pass")
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)
}

func TestReservation(t *testing.T) {

	InitTable(4)

	count := 1
	// Reserve first table should pass
	bookingId, bookedCount, remainging, res := Reserve(count)
	if res.IsSuccess == false {
		t.Errorf(res.Message)
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)

	count = 1
	// Reserve second table should pass
	bookingId, bookedCount, remainging, res = Reserve(count)
	if res.IsSuccess == false {
		t.Errorf(res.Message)
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)

	count = 2
	// Reserve third table should pass
	bookingId, bookedCount, remainging, res = Reserve(count)
	if res.IsSuccess == false {
		t.Errorf(res.Message)
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)

	count = 2
	// Reserve forth table should pass
	bookingId, bookedCount, remainging, res = Reserve(count)
	if res.IsSuccess == false {
		t.Errorf(res.Message)
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)

	count = 2
	// Reserve fifth table should fail
	bookingId, bookedCount, remainging, res = Reserve(count)
	if res.IsSuccess == true {
		t.Errorf("Shouldn't be able to reserve exceed the table count")
	}
	fmt.Printf("Reserve : %d --- Result : booking %v , bookedCount %d, remainging %d \n", count, bookingId, bookedCount, remainging)

}
