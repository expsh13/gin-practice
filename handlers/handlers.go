package handlers

import (
	"encoding/json"
	"errors"
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
	jsonData, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

// POST /api/todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.Header.Get(("Content-Length")))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}

	reqBodybuffer := make([]byte, length)

	if _, err := r.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var reqTodo models.TodoList
	if err := json.Unmarshal(reqBodybuffer, &reqTodo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	todo := reqTodo
	// todo := models.Todo1
	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// PUT /api/todo/{id}
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(listID)

	todo := models.Todo1
	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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
	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
