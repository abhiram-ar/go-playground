package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error while fetching", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("statusCode", resp.StatusCode)

}
