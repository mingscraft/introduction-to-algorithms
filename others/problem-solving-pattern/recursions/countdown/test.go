package main

import "fmt"

func main() {
	countDown(10)
}

func countDown(n int) {
	if n <= 0 {
		fmt.Println("All done!")
		return
	}

	fmt.Printf("Count down: %d\n", n)
	n --
	countDown(n)
}
