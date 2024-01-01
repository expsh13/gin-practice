package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "add todo")
	}

	updateTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "update todo")
	}

	deleteTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "delete todo")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", addTodoHandler)
	http.HandleFunc("/update", updateTodoHandler)
	http.HandleFunc("/delete", deleteTodoHandler)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
