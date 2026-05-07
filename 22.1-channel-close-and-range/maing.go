package main

import "fmt"

func calc(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go calc(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
