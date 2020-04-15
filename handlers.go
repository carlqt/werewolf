package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func NewGame(db *sqlx.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO games DEFAULT VALUES")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func ResponseHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func GamesCreate(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastID, err := NewGame(app.db)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
		} else {
			textResponse := fmt.Sprintf("%d game is starting", lastID)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(textResponse))
		}
	})
}
