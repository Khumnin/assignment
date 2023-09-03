package utility

// Get key from map's value
// Param
//	m : map object
//	value : value to search in map object
// Return
//	keys : int array that contains key which is matched value

func GetKeysByValue(m map[int]string, value string) (keys []int, ok bool) {
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
			ok = true
		}
	}
	return
}
