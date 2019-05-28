package factorial

import "testing"

var testData = [][]int {
	{1, 1},
	{2, 2},
	{4, 24},
	{7, 5040},
}

func TestFactorial(t *testing.T) {
	for _, v := range testData {
		res := factorial(v[0])

		if res != v[1] {
			t.Errorf("factorial of %d : expect %d, got %d", v[0], v[1], res)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(testData[3][0])
	}
}
