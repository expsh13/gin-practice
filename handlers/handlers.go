package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// GET /api/todo
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("get List (page %d)\n", page)
	io.WriteString(w, resString)

}

// POST /api/todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "add todo\n")
}

// PUT /api/todo/{id}
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("update No.%d\n", listID)
	io.WriteString(w, resString)
}

// DELETE /api/todo/{id}
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "delete todo\n")
}
