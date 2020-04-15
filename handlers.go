package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func NewGame(db *sqlx.DB) (int64, error) {
	res, err := db.Exec("INSERT INTO games DEFAULT VALUES")
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func ResponseHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func GamesCreate(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rowsAffected, err := NewGame(app.db)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
		} else {
			textResponse := fmt.Sprintf("%d game is starting", rowsAffected)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(textResponse))
		}
	})
}
