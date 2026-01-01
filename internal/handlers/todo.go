package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kripesh12/my-notes/internal/db"
	"github.com/kripesh12/my-notes/internal/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	queryString := "SELECT id, title, completed FROM todos"
	rows, err := db.DB.Query(context.Background(), queryString)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	if err = json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "cannot encode todos", http.StatusInternalServerError)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	json.NewDecoder(r.Body).Decode(&t)

	queryString := "INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id"

	err := db.DB.QueryRow(context.Background(), queryString, t.Title, t.Completed).Scan(&t.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
