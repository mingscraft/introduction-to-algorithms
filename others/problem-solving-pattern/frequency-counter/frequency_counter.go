package frequency_counter

import "fmt"

// ValidAnagrams check two string has same characters
func ValidAnagrams(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Freq := map[rune]int{}
	str2Freq := map[rune]int{}

	for _, s := range str1 {
		if _, exist := str1Freq[s]; !exist {
			str1Freq[s] = 0
		}
		str1Freq[s] = str1Freq[s] + 1
	}

	for _, s := range str2 {
		if _, exist := str2Freq[s]; !exist {
			str2Freq[s] = 0
		}
		str2Freq[s] = str2Freq[s] + 1
	}

	for v, f := range str1Freq {
		if f2, exist := str2Freq[v]; !exist || f2 != f {
			return false
		}
	}

	return true
}

func ValidAnagrams2(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	lookup := map[rune]int{}

	for _, v := range str1 {
		if _, exist := lookup[v]; !exist {
			lookup[v] = 0
		}

		lookup[v] = lookup[v] + 1
	}

	for _, v := range str2 {
		if f, exist := lookup[v]; !exist || f == 0 {
			return false
		}

		lookup[v] = lookup[v] - 1
	}

	return true
}

func SameFrequency(num1 int, num2 int) bool {

	numStr1 := fmt.Sprintf("%d", num1)
	numStr2 := fmt.Sprintf("%d", num2)

	if len(numStr1) != len(numStr2) {
		return false
	}

	lookup := map[rune]int{}

	for _, v := range numStr1 {
		if _, exist := lookup[v]; !exist {
			lookup[v] = 0
		}
		lookup[v] ++
	}

	for _, v := range numStr2 {
		if val, exist := lookup[v]; !exist || val == 0  {
			return false
		}
		lookup[v] = lookup[v] - 1
	}

	return true
}