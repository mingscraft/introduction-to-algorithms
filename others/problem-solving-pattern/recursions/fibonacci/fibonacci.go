package fibonacci

func Fibonacci(num int) int {

	if num <= 2 {
		return 1
	}

	return Fibonacci(num - 1) + Fibonacci(num - 2)
}

