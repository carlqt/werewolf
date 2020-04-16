package main

import (
	"fmt"
	"net/http"

	"github.com/carlqt/werewolf/models"
)

func ResponseHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func GamesCreate(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := models.NewGame(app.db)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
		} else {
			textResponse := fmt.Sprintf("The game is starting")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(textResponse))
		}
	})
}

// Join game - query Games table by an active open game by channel id
func GamesJoin(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a player
		// join the player to the game

		// err := models.NewGame(app.db)

		// if err != nil {
		// 	w.WriteHeader(http.StatusUnprocessableEntity)
		// 	w.Write([]byte(err.Error()))
		// } else {
		// 	textResponse := fmt.Sprintf("The game is starting")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(textResponse))
		// }
	})
}
