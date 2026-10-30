package main

import (
	"encoding/json/v2"
	"net/http"
)

func handleListTodos(w http.ResponseWriter, r *http.Request) {

	json.MarshalWrite(w,
		map[string]any{
			"mgs":   "Active todos",
			"todos": todosList,
		})
}
