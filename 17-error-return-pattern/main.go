package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go dont use exception for normal failures
	// functions return errors as normal return values

	// val, error := something()
	// if err != nill {handle error}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "a"
	level, error := parseLevel(input)

	if error != nil {
		return error
	}

	fmt.Println("selected level", level)

	return nil
}

func parseLevel(s string) (int, error) {
	//  nil error => success
	// non nil => failure

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 and 5")
	}

	return n, nil
}
