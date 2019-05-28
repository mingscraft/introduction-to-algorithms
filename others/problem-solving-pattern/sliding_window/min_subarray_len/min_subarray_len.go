package min_subarray_len

import "math"

// MinSubarrayLen accept two parameters - an array of positive integers and a positive integer.
// return the minimal length of a contiguous subarray.
func MinSubArrayLen(arr []int, num int) int {
	total := 0
	start := 0
	end := 0
	minLen := math.MaxInt64

	for start < len(arr) {
		if total < num && end < len(arr) {
			total += arr[end]
			end ++
		} else if total >= num {
			if end-start < minLen {
				minLen = end - start
			}
			total -= arr[start]
			start ++
		} else {
			break
		}
	}

	if minLen == math.MaxInt64 {
		return 0
	}

	return minLen
}
