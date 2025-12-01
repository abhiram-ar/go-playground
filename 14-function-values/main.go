package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)

	}

	fmt.Println("result: ", hypot(3, 4))
	fmt.Println("result: ", compute(hypot))
	fmt.Println("result: ", compute(math.Pow))
}
