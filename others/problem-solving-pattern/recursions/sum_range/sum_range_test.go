package sum_range

import "testing"

// recursiveRange(6) // 21
// recursiveRange(10) // 55

var testData = [][]int{
	{6, 21},
	{10, 55},
}

func TestSumRange(t *testing.T) {
	for _, v := range testData {
		res := sumRange(v[0])

		if res != v[1] {
			t.Errorf("sumRange(%d): expect %d, got %d", v[0], v[1], res)
		}
	}
}

func BenchmarkSumRange(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		sumRange(testData[1][1])
	}
}


