package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Abhiram"
	lastName := "Sajeev"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println("upper", strings.ToUpper(fullName))

}
