package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci2() func() int {
	first := 0
	last := 1
	return func() int {
		result := first
		first, last = last, first+last
		return result
	}
}

func main2() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
