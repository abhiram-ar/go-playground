package main

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
	"time"
)

type TodoReqBody struct {
	Text string `json:"text"`
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reqData := TodoReqBody{}
	err := json.UnmarshalRead(r.Body, &reqData, json.OmitZeroStructFields(true))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("malfored request"))
		return
	}

	todo := TodoItem{Text: reqData.Text, CreatedAt: time.Now(), Completed: false}

	todosList = append(todosList, todo)

	err = json.MarshalWrite(w, todo)
	if err != nil {
		fmt.Print("Error writing to response")
	}

}
