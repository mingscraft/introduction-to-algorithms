package stringify_numbers

import (
	"reflect"
	"testing"
)

var input = Obj{
	"num":  1,
	"test": []interface{}{},
	"data": Obj{
		"val": 4,
		"info": Obj{
			"isRight": true,
			"random":  66,
		},
	},
}

func TestStringifyNumber(t *testing.T) {
	expect := Obj{
		"num":  "1",
		"test": []interface{}{},
		"data": Obj{
			"val": "4",
			"info": Obj{
				"isRight": true,
				"random":  "66",
			},
		},
	}

	got := StringifyNumber(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkStringifyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringifyNumber(input)
	}
}
