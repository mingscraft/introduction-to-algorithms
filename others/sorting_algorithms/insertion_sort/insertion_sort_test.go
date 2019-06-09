package insertion_sort

import (
	"math/rand"
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

var bigTestData = RandomBigDataset()

func RandomBigDataset() []int {
	var arr = make([]int, 0)
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(1000)	)
	}

	return arr
}

func TestBubbleSort(t *testing.T) {
	for _, v := range testData {
		res := InsertionSort(v.arr)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v, expect %v, got %v", v.arr, v.res, res)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// intentionally only run one, as it would kill your machine. because insertion sort is too slow.
	for i := 0; i < b.N; i ++ {
		InsertionSort(bigTestData)
	}
}



