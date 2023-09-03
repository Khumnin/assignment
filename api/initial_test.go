package api

import "testing"

// Test first time initialization
// Result : Should be success both the initialization status and the value
func TestInit(t *testing.T) {
	count := 4
	want := true
	status := InitTable(count)
	if status.IsSuccess != want {
		t.Errorf("Init Table is failed. reason %v", status.Message)
	}

	// Clean up
	tableInstant = nil
}

// Test another initialization
// Result : Should be failed since we already initiated in the previous test
func TestMultipleInit(t *testing.T) {
	count := 5
	want := true

	// First initialization
	status := InitTable(count)
	if status.IsSuccess != want {
		t.Errorf(status.Message)
	}

	count = 2
	want = false

	// Second initialization
	status = InitTable(count)
	if status.IsSuccess != want {
		t.Errorf(status.Message)
	}

	// Clean up
	tableInstant = nil
}
