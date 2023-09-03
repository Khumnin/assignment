package utility

import "testing"

func TestGetKey(t *testing.T) {
	table1 := map[int]string{1: "aa", 2: "aa", 3: "bb", 4: "cc", 5: "cc", 6: "cc"}

	value, res := GetKeysByValue(table1, "aa")
	want := true
	count := 2

	if want != res {
		println("The given value doesn't exist")
	}

	if len(value) != count {
		println("getKeysByValue is wrong. Got %d, want %d", len(value), count)
	}

	value, res = GetKeysByValue(table1, "bb")
	want = true
	count = 1

	if want != res {
		println("The given value doesn't exist")
	}

	if len(value) != count {
		println("getKeysByValue is wrong. Got %d, want %d", len(value), count)
	}

	value, res = GetKeysByValue(table1, "cc")
	want = true
	count = 3

	if want != res {
		println("The given value doesn't exist")
	}

	if len(value) != count {
		println("getKeysByValue is wrong. Got %d, want %d", len(value), count)
	}

	value, res = GetKeysByValue(table1, "dd")
	want = false
	count = 0

	if want != res {
		println("The given value shouldn't exist but found")
	}

	if len(value) != count {
		println("getKeysByValue is wrong. Got %d, want %d", len(value), count)
	}
}
