package sliding_window

import "testing"

type TestData struct {
	arr []int
	num int
	maxSum int
}

var testData = []TestData{
	TestData{[]int{1,2,5,2,8,1,5}, 2, 10},
	TestData{[]int{1,2,5,2,8,1,5}, 4, 17},
	TestData{[]int{4,2,1,6}, 1, 6},
	TestData{[]int{4,2,1,6,2}, 4, 13},
	TestData{[]int{}, 4, 0},
}

func TestMaxSubArraySum(t *testing.T) {
	for _, v := range testData {
		res := MaxSubArraySum(v.arr, v.num)
		if res != v.maxSum {
			t.Errorf("%v - %d : Expect %d, got %d", v.arr, v.num, v.maxSum, res)
		}
	}
}

func BenchmarkMaxSubArraySum(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		MaxSubArraySum(testData[1].arr, testData[1].num)
	}
}
