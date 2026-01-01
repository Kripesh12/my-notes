package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kripesh12/my-notes/internal/handlers"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/signup", handlers.Register)

	return r
}
