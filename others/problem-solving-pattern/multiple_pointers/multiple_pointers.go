package multiple_pointers

// CountUniqueValues Count number of unique value in an array
func CountUniqueValues(vals []int) int {
	if len(vals) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(vals); j ++ {
		if vals[i] != vals[j] {
			i += 1
			vals[i] = vals[j]
		}
	}

	return i + 1
}
