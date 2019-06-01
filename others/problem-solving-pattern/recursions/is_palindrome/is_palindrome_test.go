package is_palindrome

import "testing"

// isPalindrome('awesome') // false
// isPalindrome('foobar') // false
// isPalindrome('tacocat') // true
// isPalindrome('amanaplanacanalpanama') // true
// isPalindrome('amanaplanacanalpandemonium') // false

type TestData struct {
	str string
	res bool
}

var testData = []TestData{
	{"awesome", false},
	{"foobar",false},
	{"tacocat", true},
	{"amanaplanacanalpanama", true},
	{"amanaplanacanalpandemonium", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, v := range testData {
		res := IsPalindrome(v.str)

		if res != v.res {
			t.Errorf("%s: expect %v, got %v", v.str, v.res, res)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		IsPalindrome(testData[4].str)
	}
}
