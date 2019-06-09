package basic_binary_search

func BasicBinarySearch(haystack []int, needle int) int {

	left := 0
	right := len(haystack) - 1

	for right - left >= 0 {
		middle := (right + left) / 2
		if haystack[middle] == needle {
			return middle
		}

		if haystack[middle] > needle {
			right = middle - 1
		}
		if haystack[middle] < needle {
			left = middle + 1
		}
	}

	return -1
}