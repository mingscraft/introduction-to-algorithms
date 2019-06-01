package is_palindrome

func IsPalindrome(str string) bool {
	l := len(str)
	if l <= 1 {
		return true
	}

	return (str[0] == str[l-1]) && IsPalindrome(str[1:l-1])
}
