package utils

// SliceEqual compare two slice are equal.
func SliceEqual(s1 []interface{}, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
