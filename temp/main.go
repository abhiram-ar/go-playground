package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello() // launches concurrently, doesn't block

	fmt.Println("Hello from main!")
	time.Sleep(1 * time.Second) // wait so the program doesn't exit immediately
}
