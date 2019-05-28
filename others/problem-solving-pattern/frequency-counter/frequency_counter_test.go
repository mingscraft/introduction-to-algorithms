package frequency_counter

import "testing"

var testData = [][]string{
	{"", ""},
	{"aaz", "zza"},
	{"anagram", "nagaram"},
	{"rat", "cat"},
	{"awesome", "awesom"},
	{"qwerty", "qeywrt"},
	{"texttwisttime", "timetwisttext"},
}

var sameFreqTestData = [][2]int{
	{182, 281},
	{34, 14},
	{3589578, 5879385},
	{22,222},
}

var sameFreqTestResults = []bool{true, false, true, false}

var expectedResult = []bool{true, false, true, false, false, true, true}

func TestValidAnagrams(t *testing.T) {
	for i, td := range testData {
		res := ValidAnagrams(td[0], td[1]);
		if res != expectedResult[i] {
			t.Errorf("ValidAngagrams(%s, %s) expect %v, got %v", td[0], td[1], expectedResult[i], res)
		}
	}
}

func TestValidAnagrams2(t *testing.T) {
	for i, td := range testData {
		res := ValidAnagrams2(td[0], td[1])
		if res != expectedResult[i] {
			t.Errorf("ValidAngagrams2(%s, %s) expect %v, got %v", td[0], td[1], expectedResult[i], res)
		}
	}
}

func TestSameFrequency(t *testing.T) {
	for i, v := range sameFreqTestData {
		res := SameFrequency(v[0], v[1])
		if res != sameFreqTestResults[i] {
			t.Errorf("SameFrequency(%d, %d) expect %v, got %v", v[0], v[1], sameFreqTestResults[i], res)
		}
	}
}

func BenchmarkValidAnagrams(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		ValidAnagrams("texttwisttime", "timetwisttext")
	}
}

func BenchmarkValidAnagrams2(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		ValidAnagrams2("texttwisttime", "timetwisttext")
	}
}

func BenchmarkSameFrequency(b *testing.B) {
	for i:=0; i < b.N; i++ {
		SameFrequency(sameFreqTestData[2][0], sameFreqTestData[2][1])
	}
}
