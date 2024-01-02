package models

import "time"

var (
	Todo1 = TodoList{
		ID:        1,
		Title:     "title",
		CreatedAt: time.Now(),
	}
	Todo2 = TodoList{
		ID:        2,
		Title:     "title2",
		CreatedAt: time.Now(),
	}
)
