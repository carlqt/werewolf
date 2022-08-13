package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/carlqt/internal/entities"
)

// playerParams is the schema of request params made for GamesJoin
// type playerParams struct {
// 	name      string `json:"name"`
// 	ChannelID string `json:"channel_id"`
// }

type errorResponse struct {
	Message string `json:"message"`
}

func ResponseHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func (a *App) GamesCreate(w http.ResponseWriter, r *http.Request) {
	params, _ := requestParams(r.Body)

	game, err := a.models.Games.NewGame(params["channel_id"])

	if err != nil {
		renderError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(game)
	}
}

func (a *App) GamesJoin(w http.ResponseWriter, r *http.Request) {
	params, err := requestParams(r.Body)
	if err != nil {
		renderError(err, w)
		return
	}

	player := entities.Player{
		ExternalID: params["external_id"],
		Name:       params["name"],
	}

	err = a.models.Players.JoinGame(params["channel_id"], &player)

	if err != nil {
		renderError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(player)

	}
}

func renderError(err error, w http.ResponseWriter) {
	errorResp := errorResponse{Message: err.Error()}

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(errorResp)
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
