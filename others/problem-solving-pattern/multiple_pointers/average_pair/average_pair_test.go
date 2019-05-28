package average_pair

import "testing"

type TestData struct {
	arr []int
	avg float32
	res bool
}

var testData = []TestData{
	{[]int{1,2,3}, 2.5, true},
	{[]int{1,3,3,5,6,7,10,12,19}, 8, true},
	{[]int{-1,0,3,4,5,6}, 4.1, false},
	{[]int{}, 4, false},
}

func TestAveragePair(t *testing.T) {
	for _, v := range testData {
		res := AveragePair(v.arr, v.avg)
		if res != v.res {
			t.Errorf("%v - %v : expect %v, got %v", v.arr, v.avg, v.res, res)
		}
	}
}

func BenchmarkAveragePair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AveragePair(testData[1].arr, testData[1].avg)
	}
}


