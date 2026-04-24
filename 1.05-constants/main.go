package main

import "fmt"

func main() {
	const appName = "Go basics"

	// typed constants
	const maxUpload int = 25
	const discountedPrice float64 = 25.5

	fmt.Println(appName, maxUpload, discountedPrice)
}
