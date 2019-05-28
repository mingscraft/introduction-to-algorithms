package multiple_pointers

import "testing"

var testData = [][]int{
	{1,1,1,1,1,2},
	{1,2,3,4,4,4,7,7,12,12,13},
	{},
	{-2,-1,-1,0,1},
}

var testResult = []int{2,7,0,4}

func TestCountUniqueValues(t *testing.T) {
	for i, test := range testData {
		got := CountUniqueValues(test)
		if got != testResult[i] {
			t.Errorf("%v: expect %d, got %d", test, testResult[i], got)
		}
	}
}

func BenchmarkCountUniqueValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountUniqueValues(testData[1])
	}
}
