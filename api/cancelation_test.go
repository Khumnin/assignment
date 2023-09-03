package api

import (
	"fmt"
	"testing"
)

func TestCancelationWithoutInit(t *testing.T) {

	freedCount, remaining, res := Cancelation("xxx")
	want := false

	if res.IsSuccess != want {
		t.Errorf("The cancelation shouldn't pass without initialization but got pass")
	}
	fmt.Printf("Result --- FreedCount : %d , Remaining : %d , Message : %v", freedCount, remaining, res.Message)
}

func TestCancelation(t *testing.T) {

	InitTable(4)
	tableInstant[1] = "ddd"
	tableInstant[2] = "bbb"
	tableInstant[3] = "ccc"
	tableInstant[4] = "bbb"

	freedCount, remaining, res := Cancelation("xxx")
	want := false

	if res.IsSuccess != want {
		t.Errorf("Cancelation with wrong booking id should be fail")
	}

	freedCount, remaining, res = Cancelation("bbb")
	want = true

	if res.IsSuccess != want {
		t.Errorf("Cancelation with correct booking id should be pass")
	}

	if freedCount != 2 {
		t.Errorf("The release table should be equal 2 but got %d", freedCount)
	}

	if remaining != 2 {
		t.Errorf("The release table should be equal 2 but got %d", remaining)
	}

	freedCount, remaining, res = Cancelation("ccc")
	want = true

	if res.IsSuccess != want {
		t.Errorf("Cancelation with correct booking id should be pass")
	}

	if freedCount != 1 {
		t.Errorf("The release table should be equal 1 but got %d", freedCount)
	}

	if remaining != 3 {
		t.Errorf("The release table should be equal 3 but got %d", remaining)
	}

	tableInstant = nil
}
