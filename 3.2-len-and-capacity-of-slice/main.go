package main

import "fmt"

func main() {
	// len => how many elements list have
	// cap => how many it can hold

	scores := make([]int, 0, 5)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 1, 2, 3)
	fmt.Println(scores, len(scores), cap(scores))

	// if slice exceeds the  its capacity,
	// go grows the backing array and usually doubles
	scores = append(scores, 4, 5, 6)
	fmt.Println(scores, len(scores), cap(scores))

	// spreding one slice on other
	todos := []string{"do youtube", "do something"}
	more := []string{"learn go lang"}
	todos = append(todos, more...)
	fmt.Println(todos)
}
