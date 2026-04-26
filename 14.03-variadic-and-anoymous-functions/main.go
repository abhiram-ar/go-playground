package main

import "fmt"

// varadic function
func sumAll(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func main() {
	ans := sumAll(1, 2, 3, 4, 5)
	fmt.Println("sum:", ans)

	numsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("sum of slice:", sumAll(numsSlice...))

	// anonymous function
	res := func(n int) int {
		return n * 2
	}
	fmt.Println(res(2))

	// IIFE
	res2 := func(a int, b int) int {
		return a + b
	}(5, 10)
	fmt.Println(res2)

}
