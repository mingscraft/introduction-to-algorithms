package insertion_sort

import (
	"fmt"
	"github.com/ray.vhatt/instroduction-to-algorithms/chapter2/randon_slice"
	"reflect"
	"testing"
)

var TestData = randon_slice.RandomIntSlice(1000)

func TestInsertionSortAscending(t *testing.T) {
	expect := []int{1,2,3,4,5,6}
	got := InsertionSort([]int{5,2,4,6,1,3}, Ascending)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func TestInsertionSortDescending(t *testing.T) {
	expect := []int{6,5,4,3,2,1}
	got := InsertionSort([]int{5,2,4,6,1,3}, Descending)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkInsertionSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(TestData, Ascending)
	}
}

func BenchmarkInsertionSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort([]int{5,2,4,6,1,3}, Descending)
	}
}

func ExampleInsertionSortAscending() {
	fmt.Printf("%v", InsertionSort([]int{5,2,4,6,1,3}, Ascending))
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleInsertionSortDescending() {
	fmt.Printf("%v", InsertionSort([]int{5,2,4,6,1,3}, Descending))
	// Output:
	// [6 5 4 3 2 1]
}