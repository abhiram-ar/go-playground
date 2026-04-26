package main

import "fmt"

func divide(a int, b int) (ans int) {
	ans = a / b

	// naked return will return the current value of ans
	return
}

func main() {
	fmt.Println("ans:", divide(4, 2))
}
