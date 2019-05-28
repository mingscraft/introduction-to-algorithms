package find_longest_substring

import "testing"

type TestData struct {
	str string
	res int
}

var testData = []TestData{
	{"", 0},
	{"rithmschool", 7},
	{"thisisawesome", 6},
	{"thecatinthehat", 7},
	{"bbbbbb", 1},
	{"longestsubstring", 8},
	{"thisishowwedoit", 6},
}

func TestFindLongestSubstring(t *testing.T) {
	for _, v := range testData {
		res := FindLongestSubstring(v.str)
		if res != v.res {
			t.Errorf("%s : expect %d, got %d", v.str, v.res, res)
		}
	}
}

func BenchmarkFindLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		FindLongestSubstring(testData[5].str)
	}
}
