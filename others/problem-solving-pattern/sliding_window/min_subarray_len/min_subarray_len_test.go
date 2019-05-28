package min_subarray_len

import "testing"

type TestData struct {
	arr []int
	sum int
	res int
}

var testData = []TestData{
	{[]int{2,3,1,2,4,3}, 7, 2},
	{[]int{2,1,6,5,4}, 9, 2},
	{[]int{3,1,7,11,2,9,8,21,62,33,19}, 52, 1},
	{[]int{1,4,16,22,5,7,8,9,10}, 39, 3},
	{[]int{1,4,16,22,5,7,8,9,10}, 55, 5},
	{[]int{4,3,3,8,1,2,3}, 11, 2},
	{[]int{1,4,16,22,5,7,8,9,10}, 95, 0},
}

func TestMinSubArrayLen(t *testing.T) {
	for _, v := range testData {
		res := MinSubArrayLen(v.arr, v.sum)
		if res != v.res {
			t.Errorf("%v - %d : expect %d, got %d", v.arr, v.sum, v.res, res)
		}
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinSubArrayLen(testData[2].arr, testData[2].sum)
	}
}
