package multiple_pointers

import "sort"

type RuneSlice []rune

func (p RuneSlice) Len() int { return len(p)}
func (p RuneSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p RuneSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i]}

func AreThereDuplicate(arr ...rune) bool {
	sort.Sort(RuneSlice(arr))

	start := 0
	next := 1

	for next < len(arr) {
		if arr[start] == arr[next] {
			return true
		}
		start++
		next ++
	}

	return false
}