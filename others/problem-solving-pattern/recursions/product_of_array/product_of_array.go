package product_of_array

func ProductOfArray(arr []int) int {
	for len(arr) == 0 {
		return 1
	}

	return arr[0] * ProductOfArray(arr[1:])
}
