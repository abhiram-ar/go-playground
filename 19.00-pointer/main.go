package main

import "fmt"

func main() {
	// pointer: store the memrory address of any val

	// operators
	// &x -> address of x
	// *x -> dereferencing (gto that address and read/write)

	score := 10
	fmt.Println("before:", score)

	addScore(&score)

	fmt.Println("after:", score)

}

func addScore(score *int) {
	*score = *score + 5
}
