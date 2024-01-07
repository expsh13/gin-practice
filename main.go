package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// r := mux.NewRouter()

	// r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	// r.HandleFunc("/api/todo", handlers.GetTodoHandler).Methods(http.MethodGet)
	// r.HandleFunc("/api/todo", handlers.AddTodoHandler).Methods(http.MethodPost)
	// r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.UpdateTodoHandler).Methods(http.MethodPut)
	// r.HandleFunc("/api/todo/{id:[0-9]+}", handlers.DeleteTodoHandler).Methods(http.MethodDelete)

	// log.Println("server start")
	// log.Fatal(http.ListenAndServe(":8080", r))

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}
