package capitalize_words

import (
	"reflect"
	"testing"
)

// let words = ["i", "am", "learning", "recursion"];
// capitalizedWords(words); // ["I", "AM", "LEARNING", "RECURSION"]

func TestCapitalizeWords(t *testing.T) {
	expect := []string{"I", "AM", "LEARNING", "RECURSION"} 
	got := CapitalizeWords([]string{"i", "am", "learning", "recursion"})
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}

func BenchmarkCapitalizeWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CapitalizeWords([]string{"i", "am", "learning", "recursion"})
	}
}
