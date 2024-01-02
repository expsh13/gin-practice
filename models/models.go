package models

import "time"

type TodoList struct {
	ID        int       `json:"todo_id"`
	Title     string    `json:"todo_title"`
	CreatedAt time.Time `json:"created_at"`
}
