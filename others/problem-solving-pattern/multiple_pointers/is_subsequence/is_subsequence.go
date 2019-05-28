package is_subsequence

func IsSubSequence(str1 string, str2 string) bool {
	str1Ptr := 0

	for _, v := range str2 {
		if str1Ptr == len(str1) - 1 {
			break
		}
		if v == rune(str1[str1Ptr]) {
			str1Ptr ++
		}
	}

	return str1Ptr == len(str1) - 1
}
