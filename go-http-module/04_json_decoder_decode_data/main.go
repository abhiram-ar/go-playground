package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

type TestRequest struct {
	Name string `json:"name"`
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only post is allowed",
		})
		return
	}

	defer r.Body.Close()

	var payload TestRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "invalid json format",
		})
		return
	}

	payload.Name = strings.TrimSpace(payload.Name)

	if payload.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "name must not be empty!",
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      payload,
		"timestamp": time.Now().UTC(),
	})
}

func main() {

	http.HandleFunc("/test", handleTest)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
