package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-todo/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handlers.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.GetTodoDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handlers.AddTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.DeleteTodoHandler).Methods(http.MethodDelete)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", r))

}
