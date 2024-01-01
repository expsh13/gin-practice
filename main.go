package main

import (
	"log"
	"net/http"

	"github.com/expsh13/go-todo/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/add", handlers.AddTodoHandler)
	http.HandleFunc("/update", handlers.UpdateTodoHandler)
	http.HandleFunc("/delete", handlers.DeleteTodoHandler)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
