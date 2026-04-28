package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TodoResponseStructure struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

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
		fmt.Println("read body failed", err)
		return
	}

	var data TodoResponseStructure

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("error while umamarchelling json", err)
	}

	fmt.Println("raw data", string(bodyBytes))
	fmt.Println("struct data", data)

}
