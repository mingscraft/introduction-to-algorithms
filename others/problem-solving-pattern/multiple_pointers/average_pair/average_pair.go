package average_pair

// AveragePair determine if there is a pair of values in the array where the average of the pair equals the target average.
func AveragePair(arr []int, avg float32) bool {
	i := 0
	j := len(arr) - 1

	for i < j {
		thisAvg := float32(arr[i] + arr[j]) / 2.0
		if thisAvg == avg {
			return true
		}

		if thisAvg < avg {
			i ++
		}

		if thisAvg > avg {
			j --
		}

	}

	return false
}
