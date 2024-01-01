package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid method というレスポンスを、405 番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
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
