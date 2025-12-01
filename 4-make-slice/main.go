package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	c[1] = 10
	printSlice("c", c)

	d := c[1:5]
	printSlice("d", d)
}

func printSlice(idx string, slice []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", idx, len(slice), cap(slice), slice)
}
