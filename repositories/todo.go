package repositories

import (
	"database/sql"

	"github.com/expsh13/go-todo/models"
)

const (
	todoNumPerPage = 5
)

// GET /api/todo
func SelectTodoList(db *sql.DB, page int) ([]models.TodoList, error) {
	const sqlStr = `
		select todo_id, title
		from todo
		limit ? offset ?;`

	rows, err := db.Query(sqlStr, todoNumPerPage, ((page - 1) * todoNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todoArray := make([]models.TodoList, 0)
	for rows.Next() {
		var todo models.TodoList
		rows.Scan(&todo.ID, &todo.Title)
		todoArray = append(todoArray, todo)
	}

	return todoArray, nil
}

// GET /api/todo/{id}
func SelectTodoDetail(db *sql.DB, todoId int) (models.TodoList, error) {
	const sqlStr = `
		select *
		from todo
		where todo_id = ?;`

	row := db.QueryRow(sqlStr, todoId)
	// if err := row.Err(); err != nil {
	// 	return models.TodoList{}, err
	// }

	var todo models.TodoList
	row.Scan(&todo.ID, &todo.Title)

	return todo, nil
}

// POST /api/todo
func InsertTodo(db *sql.DB, todo models.TodoList) (models.TodoList, error) {
	const sqlStrInsert = `
	insert into todo (title, created_at) values
	(?, now());
	`
	result, err := db.Exec(sqlStrInsert, todo.Title)
	if err != nil {
		return models.TodoList{}, err
	}

	id, _ := result.LastInsertId()

	var newTodo models.TodoList
	newTodo.Title = todo.Title
	newTodo.ID = int(id)

	return newTodo, nil
}

// PUT /api/todo/{id}
func UpdateTodo(db *sql.DB, todoId int, todo models.TodoList) error {
	const sqlStr = `
	update todo set title = ? where todo_id = ?;
	`
	_, err := db.Exec(sqlStr, todo.Title, todoId)
	if err != nil {
		return err
	}

	return nil
}

// DELETE /api/todo/{id}
func DeleteTodo(db *sql.DB, todoId int) error {
	const sqlStr = `
	delete from todo where todo_id = ?;
	`
	_, err := db.Exec(sqlStr, todoId)
	if err != nil {
		return err
	}
	return nil
}
