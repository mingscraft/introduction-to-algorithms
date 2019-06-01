package capitalize_words

import "strings"

func CapitalizeWords(arr []string) []string {
	if len(arr) == 1 {
		return []string{strings.ToUpper(arr[0])}
	}

	return  append([]string{strings.ToUpper(arr[0])}, CapitalizeWords(arr[1:])...)
}

