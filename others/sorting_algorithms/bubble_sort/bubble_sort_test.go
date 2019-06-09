package bubble_sort

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
}

func TestBubbleSort(t *testing.T) {
	for _, v := range testData {
		res := BubbleSort(v.arr)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v, expect %v, got %v", v.arr, v.res, res)
		}
	}
}

func TestBubbleSortBetter(t *testing.T) {
	for _, v := range testData {
		res := BubbleSortBetter(v.arr)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v, expect %v, got %v", v.arr, v.res, res)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		BubbleSort(testData[1].arr)
	}
}

func BenchmarkBubbleSortBetter(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		BubbleSortBetter(testData[1].arr)
	}
}
