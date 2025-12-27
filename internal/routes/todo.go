package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kripesh12/my-notes/internal/handlers"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.GetTodos)
	r.Post("/create", handlers.CreateTodo)

	return r
}
