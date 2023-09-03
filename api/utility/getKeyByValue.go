package utility

// Use to get key from value
func GetKeysByValue(m map[int]string, value string) (key []int, ok bool) {
	for k, v := range m {
		if v == value {
			key = append(key, k)
			ok = true
		}
	}
	return
}
