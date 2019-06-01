package fibonacci

import "testing"
// fib(4) // 3
// fib(10) // 55
// fib(28) // 317811
// fib(35) // 9227465
var testData = [][]int{
	{4, 3},
	{10, 55},
	{28, 317811},
	{35, 9227465},
}

func TestFibonacci(t *testing.T) {
	for _, v := range testData {
		res := Fibonacci(v[0])
		if res != v[1] {
			t.Errorf("fibonacci of %d: expect %d, got %d", v[0], v[1], res)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Fibonacci(testData[3][0])
	}
}
