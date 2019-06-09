package merge_sort

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

func TestMerge(t *testing.T) {
	got := Merge([]int{1,5,9,9,10,50}, []int{2,5, 5, 9, 14, 99, 100})
	expect := []int{1,2,5,5,5,9,9,9,10,14,50,99,100}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expect %v, got %v", expect, got)
	}

}

func RandomBigDataset() []int {
	var arr = make([]int, 0)
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(1000)	)
	}

	return arr
}

func TestMergeSort(t *testing.T) {
	for _, v := range testData {
		res := MergeSort(v.arr)
		if !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v, expect %v, got %v", v.arr, v.res, res)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// intentionally only run one, as it would kill your machine. because insertion sort is too slow.
	for i := 0; i < b.N; i ++ {
		MergeSort(bigTestData)
	}
}