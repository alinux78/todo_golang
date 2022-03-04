package model

import "time"

//DISCUSS JSON

type TodoItem struct {
	ID          string `json:"id"`
	Title       string
	Description string
	CreatedAt   time.Time
	Done        bool
}

type TodoItemInput struct {
	Title       string
	Description string
}
