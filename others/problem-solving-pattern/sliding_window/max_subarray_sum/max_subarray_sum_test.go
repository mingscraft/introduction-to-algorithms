package max_subarray_sum

import "testing"

type TestData struct {
	arr []int
	num int
	res int
}

var testData = []TestData{
	{[]int{100, 200, 300, 400}, 2, 700},
	{[]int{1,4,2,10,23,3,1,0,20}, 4, 39},
	{[]int{-3,4,0,-2,6,-1}, 2, 5},
	{[]int{3,-2,7,-4,1,-1,4,-2,1}, 2, 5},
	{[]int{2,3},3, 0},
}

func TestMaxSubarraySum(t *testing.T) {
	for _, v := range testData {
		res := MaxSubarraySum(v.arr, v.num)

		if res != v.res {
			t.Errorf("%v - %d : expect %d, got %d", v.arr, v.num, v.res, res)
		}
	}
}

func BenchmarkMaxSubarraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubarraySum(testData[2].arr, testData[2].num)
	}
}
