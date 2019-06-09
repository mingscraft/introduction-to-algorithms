package selection_sort

import (
	"reflect"
	"testing"
)

type TestData struct {
	arr []int
	res []int
}

var testData = []TestData{
	{[]int{10,2,5,7,2,10,13}, []int{2,2,5,7,10,10,13}},
	{[]int{10,200,53,43,29,17,19}, []int{10,17,19,29,43,53,200}},
	{[]int{0,2,34,22,10,19,17}, []int{0,2,10,17,19,22,34}},
}

func TestBubbleSort(t *testing.T) {
	for _, v := range testData {
		res := SelectionSort(v.arr)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v, expect %v, got %v", v.arr, v.res, res)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		SelectionSort(testData[1].arr)
	}
}


