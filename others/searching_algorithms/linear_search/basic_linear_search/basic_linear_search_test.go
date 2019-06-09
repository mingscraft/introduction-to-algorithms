package basic_linear_search

import "testing"

type TestData struct {
	arr []int
	needle int
	res int
}

var testData = []TestData{
	{[]int{10,15,20,25,30}, 15, 1},
	{[]int{9,8,7,6,5,4,3,2,1,0}, 4, 5},
	{[]int{100}, 100, 0},
	{[]int{1,2,3,4,5}, 6, -1},
	{[]int{9,8,7,6,5,4,3,2,1,0}, 10, -1},
	{[]int{100}, 200, -1},
}

func TestBasicLinearSearch(t *testing.T) {
	for _, v := range testData {
		res := BasicLinearSearch(v.arr, v.needle)
		if res != v.res {
			t.Errorf("%v - %d: expect %d, got %d", v.arr, v.needle, v.res, res)
		}
	}
}

func BenchmarkBasicLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicLinearSearch(testData[1].arr, testData[1].needle)
	}
}


