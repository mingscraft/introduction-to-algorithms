package max_subarray_sum

func MaxSubarraySum(arr []int, num int) int {
	if len(arr) < num {
		return 0
	}

	maxSum := 0
	tmpSum := 0

	for i := 0; i < num; i ++ {
		maxSum += arr[i]
	}

	tmpSum = maxSum

	for i := num; i < len(arr); i++ {
		tmpSum = tmpSum - arr[i - num] + arr[i]

		if tmpSum > maxSum {
			maxSum = tmpSum
		}
	}

	return maxSum
}
