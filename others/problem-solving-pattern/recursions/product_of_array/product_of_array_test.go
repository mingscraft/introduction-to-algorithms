package product_of_array

import "testing"

type TestData struct {
	arr []int
	res int
}

var testData = []TestData{
	{[]int{1,2,3}, 6},
	{[]int{1,2,3,10}, 60},
}

func TestProductOfArray(t *testing.T) {
	for _, v := range testData {
		res := ProductOfArray(v.arr)
		if res != v.res {
			t.Errorf("product of %v: expect %d, got %d", v.arr, v.res, res)
		}
	}
}

func BenchmarkProductOfArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductOfArray(testData[1].arr)
	}
}

