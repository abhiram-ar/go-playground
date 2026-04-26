package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// returning multiple values
func sumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	sum1 := add(10, 20)
	fmt.Println(sum1)

	sum2, product2 := sumAndProduct(10, 20)
	fmt.Println(sum2, product2)
}
