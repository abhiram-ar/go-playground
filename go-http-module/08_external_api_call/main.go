package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TodoResponseStructure struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchTodos() (TodoResponseStructure, error) {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error while fetching", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TodoResponseStructure{}, fmt.Errorf("external api non-ok status code (%s)", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return TodoResponseStructure{}, fmt.Errorf("error reading response body")
	}

	var data TodoResponseStructure
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return TodoResponseStructure{}, fmt.Errorf("error unmarshelling response body")
	}

	return data, nil

}

func externalHander(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":      false,
			"message": "only get method is allows",
		})
		return
	}

	data, err := fetchTodos()
	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"ok":      false,
			"message": "failed to fetch data",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      data,
		"timestamp": time.Now().UTC(),
	})
}

func main() {

	http.HandleFunc("/external", externalHander)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
