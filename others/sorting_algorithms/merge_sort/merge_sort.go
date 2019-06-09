package merge_sort

func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	mid := l / 2
	leftHalf := arr[0: mid]
	rightHalf := arr[mid:]

	return Merge(MergeSort(leftHalf), MergeSort(rightHalf))
}

func Merge(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0

	var tmp = make([]int, 0)

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			tmp = append(tmp, arr1[i])
			i ++
		} else {
			tmp = append(tmp, arr2[j])
			j ++
		}
	}
	tmp = append(append(tmp, arr1[i:]...), arr2[j:]...)

	return tmp
}



