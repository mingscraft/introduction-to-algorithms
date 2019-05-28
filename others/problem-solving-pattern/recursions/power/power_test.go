package power

import (
	"testing"
)

// power(2,0) // 1
// power(2,2) // 4
// power(2,4) // 16

var testData = [][]int{
	{2,0,1},
	{2,2,4},
	{2,4,16},
}

func TestPower(t *testing.T) {
	for _, v := range testData {
		res := Power(v[0], v[1])
		if res != v[2] {
			t.Errorf("%d - %d: expect %d, got %d", v[0], v[1], v[2], res)
		}
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(testData[0][0], testData[0][1])
	}
}
