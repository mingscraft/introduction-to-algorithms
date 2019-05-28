package find_longest_substring

func FindLongestSubstring(str string) int {
	longest := 0
	seen := map[rune]int{}
	start := 0

	for i := 0; i < len(str); i++ {
		c := rune(str[i])

		if v, exist := seen[c]; exist {
			if v > start {
				start = v
			}
		}

		l := i - start + 1
		if l > longest {
			longest = l
		}

		seen[c] = i + 1
	}

	return longest

}
