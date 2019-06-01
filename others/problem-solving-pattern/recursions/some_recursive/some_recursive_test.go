package some_recursive

import (
	"reflect"
	"runtime"
	"testing"
)

// SAMPLE INPUT / OUTPUT
// const isOdd = val => val % 2 !== 0;

// someRecursive([1,2,3,4], isOdd) // true
// someRecursive([4,6,8,9], isOdd) // true
// someRecursive([4,6,8], isOdd) // false
// someRecursive([4,6,8], val => val > 10); // false

type TestData struct {
	arr []int
	callback func(int) bool
	res bool
}

var testData = []TestData{
	{[]int{1,2,3,4}, isOdd, true},
	{[]int{4,6,8,9}, isOdd, true},
	{[]int{4,6,8}, isOdd,false},
	{[]int{4,6,8}, greaterThan10, false},
}

var isOdd = func(i int) bool {
	return i % 2 != 0
}

var greaterThan10 = func(i int) bool {
	return i > 10
}

func TestSomeRecursive(t *testing.T) {
	for _, v := range testData {
		res := SomeRecursive(v.arr, v.callback)
		if res != v.res {
			t.Errorf("%v - %s: expect %v, got %v", v.arr, GetFunctionName(v.callback), v.res, res)
		}
	}
}

func BenchmarkSomeRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeRecursive(testData[1].arr, testData[1].callback)
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
