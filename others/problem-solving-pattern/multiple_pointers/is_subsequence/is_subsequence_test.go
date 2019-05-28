package is_subsequence

import "testing"

type TestData struct {
	str1 string
	str2 string
	res bool
}

var testData = []TestData{
	{"hello", "hello word", true},
	{"sing", "sting", true},
	{"abc", "abracadabra", true},
	{"abc", "acb", true},
}

func TestIsSubsequence(t *testing.T) {
	for _, v := range testData {
		res := IsSubSequence(v.str1, v.str2)
		if res != v.res {
			t.Errorf("%s - %s : expect %v, got %v", v.str1, v.str2, v.res, res)
		}
	}
}

func BenchmarkIsSubSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSubSequence(testData[2].str1, testData[2].str2)
	}
}
