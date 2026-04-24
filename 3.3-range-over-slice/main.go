package main

import "fmt"

func main() {
	views := []int{10, 20, 45, 50, 60}

	// for range => loop over collections
	total := 0
	for idx, val := range views {
		fmt.Println("day", idx, val)
		total += val
	}
	fmt.Println("total", total)
}
