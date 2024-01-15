package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/expsh13/go-todo/models"
	"github.com/expsh13/go-todo/repositories"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail 関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		testTitle string
		expected  models.TodoList
	}{
		{
			testTitle: "test1",
			expected: models.TodoList{
				ID:    7,
				Title: "firstPost",
			},
		},
		{
			testTitle: "test2",
			expected: models.TodoList{
				ID:    8,
				Title: "update tes",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectTodoDetail(db, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
		})
	}

}
