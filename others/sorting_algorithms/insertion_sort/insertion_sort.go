package insertion_sort

func InsertionSort(arr []int) []int {

	var tmp = make([]int, len(arr))
	copy(tmp, arr)

	for i := 1; i < len(tmp); i ++ {
		val := tmp[i]

		j := i - 1

		for j >= 0 && tmp[j] > val {
			tmp[j+1] = tmp[j]
			j --
		}

		tmp[j + 1] = val
	}

	return tmp

}


