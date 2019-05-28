package randon_slice

import "math/rand"

func RandomIntSlice(n int) []int {
	var result []int
	for i:=0; i < n; i ++ {
		rn := rand.Intn(n)
		result = append(result, rn)
	}

	return result
}
