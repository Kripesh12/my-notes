package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kripesh12/my-notes/internal/db"
	"github.com/kripesh12/my-notes/internal/env"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env file")
	}

	//database connection
	if err := db.Connect(); err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal("server stopped: ", err.Error())
	}

}
