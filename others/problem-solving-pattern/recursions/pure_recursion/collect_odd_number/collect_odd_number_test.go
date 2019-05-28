package collect_odd_number

import (
	"reflect"
	"testing"
)

func TestCollectOddNumber(t *testing.T) {
	expect := []int{1,3,5,7}
	got := CollectOddNumber([]int{1,2,3,4,5,6,7})

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}