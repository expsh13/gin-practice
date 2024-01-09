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

	const sqlStr = `
	select *
 	from todo;
 	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	todoArray := make([]models.TodoList, 0)
	for rows.Next() {
		var todo models.TodoList
		var createdTime sql.NullTime

		err := rows.Scan(&todo.ID, &todo.Title, &createdTime)
		if createdTime.Valid {
			todo.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			todoArray = append(todoArray, todo)
		}
	}

	fmt.Printf("%+v\n", todoArray)
}
