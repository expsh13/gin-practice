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
	http.HandleFunc("/", helloHandler)

	log.Println("server start")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
