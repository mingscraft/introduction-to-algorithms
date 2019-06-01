package collect_strings

import (
	"reflect"
	"testing"
)

//const obj = {
//stuff: "foo",
//data: {
//val: {
//thing: {
//info: "bar",
//moreInfo: {
//evenMoreInfo: {
//weMadeIt: "baz"
//}
//}
//}
//}
//}
//}
//
//collectStrings(obj) // ["foo", "bar", "baz"])

var input = Obj{
	"stuff": "foo",
	"data": Obj{
		"val": Obj{
			"thing": Obj{
				"info": "bar",
				"moreInfo": Obj{
					"evenMoreInfo": Obj{
						"weMadeIt": "baz",
					},
				},
			},
		},
	},
}

func TestCollectString(t *testing.T) {
	expect := []string{"foo", "bar", "baz"}
	got := CollectString(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkCollectString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CollectString(input)
	}
}
