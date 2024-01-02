package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/expsh13/go-todo/handlers"
	"github.com/gorilla/mux"
)

type TodoList struct {
	ID        int       `json:"todo_id"`
	Title     string    `json:"todo_title"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handlers.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handlers.AddTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.DeleteTodoHandler).Methods(http.MethodDelete)

	todo := TodoList{
		ID:        1,
		Title:     "title",
		CreatedAt: time.Now(),
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", r))
}
