package binary_sum

// Sum work out the sum two binary numbers
func Sum(a []int, b []int) []int {
	// initialisation invariant: result[0..i] should be sum of a[0..i] and b[0...i]
	// maintain invariant: as we increase i = i + 1, result[0..i] should be sum of a[0..i] and b[0...i]
	// termination invariant: when i = n - 1, result[0..n] should be sum of a[0..n] and b[0...n]
	var result []int
	residual := 0
	for i := 0; i < len(a); i++ {
		total := a[i] + b[i] + residual
		residual = 0
		if total >= 2 {
			residual = 1
		}
		result = append(result, total % 2)
	}

	result = append(result, residual)

	return result
}
