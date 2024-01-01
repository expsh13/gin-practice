package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-todo/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", handlers.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", handlers.AddTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo", handlers.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/todo", handlers.DeleteTodoHandler).Methods(http.MethodDelete)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", r))
}
