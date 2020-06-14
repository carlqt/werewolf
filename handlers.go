package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		params, _ := requestParams(r.Body)

		err := models.NewGame(app.db, params["channel_id"])
		// Create new game
		// when successful, move state to "waiting for players"
		// go routine to reply
		// go routine for a countdown timer and move to next state

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

// GamesJoin - join the game on the passed in channelID
func GamesJoin(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, err := requestParams(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		game, err := models.ActiveGameOnChannel(app.db, params["channel_id"])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		player := models.Player{
			Name:   params["name"],
			GameID: game.ID,
		}

		err = player.Create(app.db)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		response := fmt.Sprintf("%s is joining game %s", player.Name, game.ChannelID)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})
}

func requestParams(requestBody io.ReadCloser) (map[string]string, error) {
	var params map[string]string

	body, err := ioutil.ReadAll(requestBody)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
