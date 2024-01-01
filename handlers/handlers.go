package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// GET /todo
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add todo\n")
}

// POST /todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add todo\n")
}

// PUT /todo
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update todo\n")
}

// DELETE /todo
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "delete todo\n")
}
