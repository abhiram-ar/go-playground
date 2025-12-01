package main

import "fmt"

func main() {
	a := make([]int, 0, 2)

	a = append(a, 3)
	printSlice("a", a)
	fmt.Println("")

	b := append(a, 2)
	printSlice("a", a)
	printSlice("b", b)
	fmt.Println("")

	c := append(b, 1)
	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)

}

func printSlice(idx string, slice []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", idx, len(slice), cap(slice), slice)
}
