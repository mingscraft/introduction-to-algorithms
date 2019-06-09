package basic_linear_search

func BasicLinearSearch(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}

	return -1
}
