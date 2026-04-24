package main

import "fmt"

func main() {
	views1 := 1000
	views2 := 2000
	totalViews := views1 + views2

	likes := 10
	likes++
	likes++

	avgViess := totalViews / 2

	fmt.Println(totalViews, likes, avgViess)

	rating1 := 4.5
	ratint2 := 5.1

	avgRating := (rating1 + ratint2) / 2
	fmt.Println(avgRating)
}
