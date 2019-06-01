package some_recursive

func SomeRecursive(arr []int, cb func(int) bool) bool {
	if len(arr) <= 1 {
		return cb(arr[0])
	}

	return cb(arr[0])|| SomeRecursive(arr[1:], cb)
}
