package bubble_sort

func BubbleSort(arr []int) []int {
	var tmp = make([]int, len(arr))
	copy(tmp, arr)

	for i:= 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp) - 1; j++ {
			if tmp[j] > tmp[j + 1] {
				tmp[j], tmp[j+1] = tmp[j+1], tmp[j]
			}
		}
	}

	return tmp
}

func BubbleSortBetter(arr []int) []int {
	var tmp = make([]int, len(arr))
	copy(tmp, arr)

	nSwap := 1

	for i:= len(tmp) - 1; i >= 0 && nSwap != 0; i-- {
		nSwap = 0
		for j := 0; j <= i - 1; j++ {
			if tmp[j] > tmp[j + 1] {
				tmp[j], tmp[j+1] = tmp[j+1], tmp[j]
				nSwap ++
			}
		}
	}


	return tmp
}
