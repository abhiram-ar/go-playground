package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("case1: sucess")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}

	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}

}

func doWork(sucess bool) error {
	// image a resource
	// start message => resource acquired
	// cleanup message => resource released

	fmt.Println("start: resource acquired")

	// defer wil guratee this run at eht end of this func
	// regardless if the fuction completes or errors out
	defer fmt.Println("cleaup: resource released")

	if !sucess {
		return errors.New("Something went wrong. I am returning early")
	}

	fmt.Println("work: doing something imp")
	fmt.Println("work: this work is done")

	return nil
}
