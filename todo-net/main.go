package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/create", handleCreateTodo)
	mux.HandleFunc("/list", handleListTodos)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
