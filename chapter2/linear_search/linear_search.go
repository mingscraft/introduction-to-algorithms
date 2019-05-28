package linear_search

func LinearSearch(stack []int, needle int) int {
	var val int
	for i := 0; i < len(stack); i ++ {
		val = stack[i]
		if val == needle {
			return i
		}
	}
	return -1
}
