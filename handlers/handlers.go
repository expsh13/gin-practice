package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add todo")
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update todo")
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "delete todo")
}
