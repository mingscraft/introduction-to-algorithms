package power

func Power(n1 int, n2 int) int {
	if n2 <= 0 {
		return 1
	}

	return n1 * Power(n1, n2 - 1)
}
