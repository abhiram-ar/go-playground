package main

import "fmt"

func main() {
	// slices
	// most common collection type
	// dynamic and can grow
	// []type{...values}

	results := []string{"Abhiram", "Sajeev"}
	fmt.Println(results)
	fmt.Println(results[0], results[len(results)-1])

	results[1] = "AR"
	fmt.Println(results)

	// declare an slice and use it later
	// note: zero value of decalred array is nil -> len = 0 and capacity = 0
	var nums []int
	nums = append(nums, 1, 2, 3, 4)
	fmt.Println("Nums", nums)
}
