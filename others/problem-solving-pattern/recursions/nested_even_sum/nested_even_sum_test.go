package nested_even_sum

import "testing"

type TestData struct {
	obj Obj
	sum int
}

var testData = []TestData{
	{Obj{
		"outer": 2,
		"obj": Obj{
			"inner": 2,
			"otherObj": Obj{
				"supperInner": 2,
				"notANumber": true,
				"alsoNotANumber": "yup",
			},
		},
	}, 6},
	{Obj{
		"a": 2,
		"b": Obj{
			"b": 2,
			"bb": Obj{
				"b": 3,
				"bb": Obj{
					"b": 2,
				},
			},
		},
		"c": Obj{
			"c": Obj{
				"c": 2,
			},
			"cc": "ball",
			"ccc": 5,
		},
		"d": 1,
		"e": Obj{
			"e": Obj{
				"e": 2,
			},
			"ee": "car",
		},
	}, 10},
}

func TestNestedEvenSum(t *testing.T) {
	for _, v := range testData {
		res := NestedEvenSum(v.obj)
		if res != v.sum {
			t.Errorf("%v: expect %d, got %d", v.obj, v.sum, res)
		}
	}
}

func BenchmarkNestedEvenSum(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		NestedEvenSum(testData[1].obj)
	}
}

