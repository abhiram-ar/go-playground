package main

import "fmt"

func main() {
	a := [...]string{"hello", "world"}
	// type of a = [2]int

	fmt.Printf("%v len=%d cap=%d", a, len(a), cap(a))
}
