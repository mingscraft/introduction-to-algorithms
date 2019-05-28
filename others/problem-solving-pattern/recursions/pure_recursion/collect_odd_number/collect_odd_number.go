package collect_odd_number

func CollectOddNumber(arr []int) []int {
	var res = make([]int, 0)

	if len(arr) <= 0 {
		return res
	}

	if arr[0] % 2 != 0 {
		res = append(res, arr[0])
	}

	res = append(res, CollectOddNumber(arr[1:])...)

	return res
}
