package collect_odd_number

// CollectOddNumber find all odd number within an array
// helper method recursion is a function that call another function recursively
func CollectOddNumber(arr []int) []int {
	var res = make([]int, 0)
	helper(arr, &res)

	return res
}

func helper(arr []int, accuArr *[]int) {
	if len(arr) <= 0 {
		return
	}

	if arr[0] % 2 != 0 {
		*accuArr = append(*accuArr, arr[0])
	}

	helper(arr[1:], accuArr)
}
