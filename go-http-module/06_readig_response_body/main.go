package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error while fetching", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("httpstatus", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return
	}

	bodyText := string(bodyBytes)

	// just limiting the printing to 250 characters max
	// if response has less than 250 characters and
	// we print string slice with assumed capacity 250
	// we will get a panic error
	max := 250
	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println("body data", bodyText[:max])

}
