package flatten

func Flatten(arr interface{}) []int {
	return doFlatten([]int{}, arr)
}

func doFlatten(acc []int, arr interface{}) []int {

	switch v := arr.(type) {
	case []int:
		acc = append(acc, v...)
	case int:
		acc = append(acc, v)
	case []interface{}:
		for i := range v {
			acc = doFlatten(acc, v[i])
		}
	}

	return acc
}

