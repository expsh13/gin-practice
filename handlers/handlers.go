package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/expsh13/go-todo/models"
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

	fmt.Println(page)

	todoList := []models.TodoList{models.Todo1, models.Todo2}
	json.NewEncoder(w).Encode(todoList)

}

// POST /api/todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	var reqTodo models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&reqTodo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	todo := reqTodo
	// todo := models.Todo1
	json.NewEncoder(w).Encode(todo)
}

// PUT /api/todo/{id}
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(listID)

	var reqTodo models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&reqTodo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	todo := models.Todo1

	json.NewEncoder(w).Encode(todo)
}

// DELETE /api/todo/{id}
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(listID)

	todo := models.Todo1
	json.NewEncoder(w).Encode(todo)
}
