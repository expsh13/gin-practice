package main

import (
	"database/sql"
	"fmt"

	"github.com/expsh13/go-todo/models"
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

	todoID := 7
	const sqlStr = `
	select *
 	from todo
	where todo_id = ?;
 	`
	row := db.QueryRow(sqlStr, todoID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var todo models.TodoList
	var createdTime sql.NullTime
	err = row.Scan(&todo.ID, &todo.Title, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if createdTime.Valid {
		todo.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", todo)
}
