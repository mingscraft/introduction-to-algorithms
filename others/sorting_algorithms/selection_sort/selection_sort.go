package selection_sort

func SelectionSort(arr []int) []int {
	var tmp = make([]int, len(arr))
	copy(tmp, arr)

	for i := 0; i < len(tmp) - 1; i ++ {
		minIndex := i
		for j := i + 1; j < len(tmp); j ++ {
			if tmp[minIndex] > tmp[j] {
				minIndex = j
			}
		}

		if minIndex != i  {
			tmp[i], tmp[minIndex] = tmp[minIndex], tmp[i]
		}
	}

	return tmp
}
