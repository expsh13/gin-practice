package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-todo/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/todoGet", handlers.GetTodoHandler)
	http.HandleFunc("/todoAdd", handlers.AddTodoHandler)
	http.HandleFunc("/todoUpdate", handlers.UpdateTodoHandler)
	http.HandleFunc("/todoDelete", handlers.DeleteTodoHandler)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
