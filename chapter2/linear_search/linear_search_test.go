package linear_search

import (
	"fmt"
	"testing"
)

var TestData = []int{1,2,3,4,5}

func TestLinearSearch(t *testing.T) {
	expect := 3
	got := LinearSearch(TestData, 4)

	if expect != got {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	expect := -1
	got := LinearSearch(TestData, 0)

	if expect != got {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func ExampleLinearSearch() {
	fmt.Println(LinearSearch(TestData, 4))
	// output
	// 3
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		LinearSearch(TestData, 4)
	}
}

