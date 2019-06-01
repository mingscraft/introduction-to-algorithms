package capitalize_first

import (
	"reflect"
	"testing"
)

// capitalizeFirst(['car','taco','banana']); // ['Car','Taco','Banana']

func TestCapitalizeFirst(t *testing.T) {
	expect := []string{"Car","Taco","Banana"}
	got := CapitalizeFirst([]string{"car","taco","banana"})
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}
