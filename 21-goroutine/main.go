package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := range 5 {
		time.Sleep(time.Second * 5)
		fmt.Println(i, s)
	}
}

func main() {
	go say("go")
	fmt.Printf("hello")
}
