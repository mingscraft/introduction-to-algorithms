package insertion_sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	expect := []int{1,2,3,4,5,6}
	got := InsertionSort([]int{5,2,4,6,1,3})

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort([]int{5,2,4,6,1,3})
	}
}

func ExampleInsertionSort() {
	fmt.Printf("%v", InsertionSort([]int{5,2,4,6,1,3}))
	// Output:
	// [1 2 3 4 5 6]
}