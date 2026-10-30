package main

import "time"

type TodoItem struct {
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

var todosList = []TodoItem{}
