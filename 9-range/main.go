package main

import "fmt"

var slice = []int{1, 2, 4, 16, 32, 64, 128}

func main() {
	for idx, val := range slice {
		fmt.Printf("%d %d\n", idx, val)
	}

	// skipping first value
	println("")
	for _, val := range slice {
		fmt.Printf("%d\n", val)
	}

	// avoiding second value
	println("")
	for idx := range slice {
		fmt.Printf("%d\n", idx)
	}
}
