package flatten

import (
	"reflect"
	"testing"
)

// flatten([1, 2, 3, [4, 5] ]) // [1, 2, 3, 4, 5]
// flatten([1, [2, [3, 4], [[5]]]]) // [1, 2, 3, 4, 5]
// flatten([[1],[2],[3]]) // [1,2,3]
// flatten([[[[1], [[[2]]], [[[[[[[3]]]]]]]]]]) // [1,2,3
type TestData struct {
	val interface{}
	res []int
}

var testData = []TestData{
	{[]interface{}{1, 2, 3, []int{4, 5} }, []int{1, 2, 3, 4, 5}},
	{[]interface{}{1, []interface{}{ 2, []int{3,4}, []interface{}{ []int{5} }, } }, []int{1, 2, 3, 4, 5}},
	{[]interface{}{[]int{1}, []int{2}, []int{3}}, []int{1,2,3}},
	{[]interface{} { []interface{} { []interface{}{  []interface{}{ []int{1} }, []interface{}{ []int{2} }, []interface{}{ []int{3} },} } }, []int{1,2,3}},
}

func TestFlatten(t *testing.T) {
	for _, v := range testData {
		res := Flatten(v.val)
		if  !reflect.DeepEqual(res, v.res) {
			t.Errorf("%v: expect %v, got %v", v.res, v.res, res)
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Flatten(testData[3].val)
	}
}