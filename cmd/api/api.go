package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kripesh12/my-notes/internal/routes"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", app.greetHandler)
		r.Get("/health", app.healthCheckHandler)
		r.Mount("/todos", routes.TodoRoutes())
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	server := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started on the port %s", app.config.addr)

	return server.ListenAndServe()
}
