package main

import "fmt"

func main() {
	rgn := 3
	fmt.Printf("sumRange of %d is: %d\n", rgn, sumRange(rgn))
}

func sumRange(num int) int {
	if num == 1 {
		return 1
	}

	return num + sumRange(num - 1)
}
