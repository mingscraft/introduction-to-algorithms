package capitalize_first

import "unicode"

func CapitalizeFirst(arr []string) []string {
	if len(arr) <= 0 {
		return []string{}
	}
	cap := unicode.ToUpper(rune(arr[0][0]))
	upper := string(cap) + arr[0][1:]
	return append([]string{upper}, CapitalizeFirst(arr[1:])...)
}
