package basic_binary_search

import "testing"

type TestData struct {
	arr []int
	needle int
	res int
}

var testData = []TestData{
	{[]int{5,6,10,13,14,18,30,34,35,37,40,44,64,79,84,86,95,96,98,99}, 10, 2},
	{[]int{5,6,10,13,14,18,30,34,35,37,40,44,64,79,84,86,95,96,98,99},  95, 16},
	{[]int{5,6,10,13,14,18,30,34,35,37,40,44,64,79,84,86,95,96,98,99},  100, -1},
	{[]int{1,2}, 2, 1},
	{[]int{1,2}, 1, 0},
	{[]int{1,2,3}, 1, 0},
	{[]int{1,2,3}, 2, 1},
	{[]int{1,2,3}, 3, 2},
	{[]int{1,2,3}, 4, -1},
	{[]int{1}, 4, -1},
	{[]int{1}, 1, 0},
}

func TestBinarySearch(t *testing.T) {
	for _, v := range testData {
		res := BasicBinarySearch(v.arr, v.needle)
		if res != v.res {
			t.Errorf("%v - %d: expect %d, got %d", v.arr, v.needle, v.res, res)
		}
	}
}

func BenchmarkBasicBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicBinarySearch(testData[1].arr, testData[1].needle)
	}
}
