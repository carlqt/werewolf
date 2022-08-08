package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/carlqt/internal/entities"
	"github.com/carlqt/internal/models"
	"github.com/gorilla/handlers"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	// entities entities.Entities
	models models.Models
}

func main() {
	app := NewApp()
	app.Start()
}

func NewApp() *App {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=werewolf password=postgres sslmode=disable port=5433")
	if err != nil {
		panic(err)
	}

	entities := entities.NewEntities(db)

	app := &App{
		models: models.NewModels(entities),
	}

	return app
}

func (a *App) Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := "8000"

	loggedRouter := handlers.LoggingHandler(os.Stdout, a.NewRouter())

	log.Printf("starting at port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), loggedRouter)
}
