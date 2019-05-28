package binary_sum

import (
	"fmt"
	"reflect"
	"testing"
)

var TestData = map[string][]int{
	"a": {1,0,1,1,0,1},
	"b": {0,1,1,0,1,1},
}
func TestSum(t *testing.T) {
	expect := []int{1,1,0,0,0,1,1}
	got := Sum(TestData["a"], TestData["b"])

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(TestData["a"], TestData["b"])
	}
}

func ExampleSum() {
	fmt.Printf("%v", Sum(TestData["a"], TestData["b"]))
	// output
	// [1 1 0 0 0 1 1]
}
