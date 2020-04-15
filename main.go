package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	db *sqlx.DB
}

const (
	dbName   = "werewolf"
	host     = "localhost"
	user     = "postgres"
	password = ""
	dbPort   = 5433
)

func newApp() App {
	db := initDB()
	return App{db: db}
}

func initDB() *sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		dbPort,
		user,
		password,
		dbName,
	)

	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return db
}

func (app *App) ping() {
	err := app.db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("PONG")
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func main() {
	app := newApp()
	defer app.db.Close()
	app.ping()

	router := mux.NewRouter()
	router.HandleFunc("/play", PlayHandler).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("starting at port 8000")
	http.ListenAndServe(":8000", loggedRouter)

}
