package selection_sort

import (
	"fmt"
	"github.com/ray.vhatt/instroduction-to-algorithms/chapter2/randon_slice"
	"reflect"
	"testing"
)

var TestData = []int{4,3,5,2,5,6,1}
var BenchmarkTestData = randon_slice.RandomIntSlice(1000)
func TestSelectionSort(t *testing.T) {
	expect := []int{1,2,3,4,5,5,6}
	got := SelectionSort(TestData)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(BenchmarkTestData)
	}
}

func ExampleSelectionSort() {
	fmt.Printf("%v", SelectionSort(TestData))
	// output
	// [1 2 3 4 5 5 6]
}
